package azure

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/spiffe/spire/pkg/common/plugin/azure"
	"github.com/spiffe/spire/proto/agent/nodeattestor"
	"github.com/spiffe/spire/proto/common/plugin"
	"github.com/stretchr/testify/suite"
	jose "gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

func TestMSIAttestorPlugin(t *testing.T) {
	suite.Run(t, new(MSIAttestorSuite))
}

type MSIAttestorSuite struct {
	suite.Suite

	attestor *nodeattestor.BuiltIn

	expectedResource string
	token            string
	tokenErr         error
}

func (s *MSIAttestorSuite) SetupTest() {
	s.expectedResource = azure.DefaultMSIResourceID
	s.token = ""
	s.tokenErr = nil

	s.newAttestor()

	_, err := s.attestor.Configure(context.Background(), &plugin.ConfigureRequest{
		GlobalConfig: &plugin.ConfigureRequest_GlobalConfig{
			TrustDomain: "example.org",
		},
	})
	s.Require().NoError(err)
}

func (s *MSIAttestorSuite) TestFetchAttestationDataNotConfigured() {
	s.newAttestor()
	s.requireFetchError("azure-msi: not configured")
}

func (s *MSIAttestorSuite) TestFetchAttestationDataFailedToObtainToken() {
	s.tokenErr = errors.New("FAILED")
	s.requireFetchError("azure-msi: unable to fetch token: FAILED")
}

func (s *MSIAttestorSuite) TestFetchAttestationDataAccessTokenMalformed() {
	s.token = ""
	s.requireFetchError("azure-msi: unable to parse token")
}

func (s *MSIAttestorSuite) TestFetchAttestationDataAccessTokenHasBadClaims() {
	s.token = "e30.f32.baadf00d"
	s.requireFetchError("azure-msi: unable to parse token claims")
}

func (s *MSIAttestorSuite) TestFetchAttestationDataAccessTokenMissingPrincipalID() {
	s.token = s.makeAccessToken("", "TENANTID")
	s.requireFetchError("azure-msi: token missing subject claim")
}

func (s *MSIAttestorSuite) TestFetchAttestationDataAccessTokenMissingTenantID() {
	s.token = s.makeAccessToken("PRINCIPALID", "")
	s.requireFetchError("azure-msi: token missing tenant ID claim")
}

func (s *MSIAttestorSuite) TestFetchAttestationDataSuccess() {
	s.token = s.makeAccessToken("PRINCIPALID", "TENANTID")

	stream, err := s.attestor.FetchAttestationData(context.Background())
	s.Require().NoError(err)
	s.Require().NotNil(stream)

	resp, err := stream.Recv()
	s.Require().NoError(err)
	s.Require().NotNil(resp)

	// assert attestation data
	s.Require().Equal("spiffe://example.org/spire/agent/azure_msi/TENANTID/PRINCIPALID", resp.SpiffeId)
	s.Require().NotNil(resp.AttestationData)
	s.Require().Equal("azure_msi", resp.AttestationData.Type)
	s.Require().JSONEq(fmt.Sprintf(`{"token": %q}`, s.token), string(resp.AttestationData.Data))

	// node attestor should return EOF now
	_, err = stream.Recv()
	s.Require().Equal(io.EOF, err)
}

func (s *MSIAttestorSuite) TestConfigure() {
	// malformed configuration
	resp, err := s.attestor.Configure(context.Background(), &plugin.ConfigureRequest{
		Configuration: "blah",
		GlobalConfig:  &plugin.ConfigureRequest_GlobalConfig{},
	})
	s.requireErrorContains(err, "azure-msi: unable to decode configuration")
	s.Require().Nil(resp)

	resp, err = s.attestor.Configure(context.Background(), &plugin.ConfigureRequest{})
	s.requireErrorContains(err, "azure-msi: global configuration is required")
	s.Require().Nil(resp)

	// missing trust domain
	resp, err = s.attestor.Configure(context.Background(), &plugin.ConfigureRequest{GlobalConfig: &plugin.ConfigureRequest_GlobalConfig{}})
	s.Require().EqualError(err, "azure-msi: global configuration missing trust domain")
	s.Require().Nil(resp)

	// success
	resp, err = s.attestor.Configure(context.Background(), &plugin.ConfigureRequest{
		GlobalConfig: &plugin.ConfigureRequest_GlobalConfig{TrustDomain: "example.org"},
	})
	s.Require().NoError(err)
	s.Require().Equal(resp, &plugin.ConfigureResponse{})
}

func (s *MSIAttestorSuite) TestGetPluginInfo() {
	resp, err := s.attestor.GetPluginInfo(context.Background(), &plugin.GetPluginInfoRequest{})
	s.Require().NoError(err)
	s.Require().Equal(resp, &plugin.GetPluginInfoResponse{})
}

func (s *MSIAttestorSuite) newAttestor() {
	attestor := NewMSIAttestorPlugin()
	attestor.hooks.fetchMSIToken = func(ctx context.Context, httpClient azure.HTTPClient, resource string) (string, error) {
		if httpClient != http.DefaultClient {
			return "", errors.New("unexpected http client")
		}
		if resource != s.expectedResource {
			return "", fmt.Errorf("expected resource %s; got %s", s.expectedResource, resource)
		}
		s.T().Logf("RETURNING %v %v", s.token, s.tokenErr)
		return s.token, s.tokenErr
	}
	s.attestor = nodeattestor.NewBuiltIn(attestor)
}

func (s *MSIAttestorSuite) requireFetchError(contains string) {
	stream, err := s.attestor.FetchAttestationData(context.Background())
	s.Require().NoError(err)
	s.Require().NotNil(stream)

	resp, err := stream.Recv()
	s.requireErrorContains(err, contains)
	s.Require().Nil(resp)
}

func (s *MSIAttestorSuite) requireErrorContains(err error, contains string) {
	s.Require().Error(err)
	s.Require().Contains(err.Error(), contains)
}

func (s *MSIAttestorSuite) makeAccessToken(principalID, tenantID string) string {
	claims := azure.MSITokenClaims{
		Claims: jwt.Claims{
			Subject: principalID,
		},
		TenantID: tenantID,
	}

	signingKey := jose.SigningKey{Algorithm: jose.HS256, Key: []byte("KEY")}
	signer, err := jose.NewSigner(signingKey, nil)
	s.Require().NoError(err)

	token, err := jwt.Signed(signer).Claims(claims).CompactSerialize()
	s.Require().NoError(err)
	return token
}
