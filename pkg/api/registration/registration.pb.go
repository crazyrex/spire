// Code generated by protoc-gen-go. DO NOT EDIT.
// source: registration.proto

/*
Package registration is a generated protocol buffer package.

It is generated from these files:
	registration.proto

It has these top-level messages:
	RegistrationEntryID
	ParentID
	SpiffeID
	UpdateEntryRequest
	FederatedBundle
	CreateFederatedBundleRequest
	ListFederatedBundlesReply
	FederatedSpiffeID
*/
package registration

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_api1 "google.golang.org/genproto/googleapis/api/annotations"
import spire_common "github.com/spiffe/spire/pkg/common"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// http from public import google/api/annotations.proto
var E_Http = google_api1.E_Http

// Empty from public import github.com/spiffe/spire/pkg/common/common.proto
type Empty spire_common.Empty

func (m *Empty) Reset()         { (*spire_common.Empty)(m).Reset() }
func (m *Empty) String() string { return (*spire_common.Empty)(m).String() }
func (*Empty) ProtoMessage()    {}

// AttestedData from public import github.com/spiffe/spire/pkg/common/common.proto
type AttestedData spire_common.AttestedData

func (m *AttestedData) Reset()          { (*spire_common.AttestedData)(m).Reset() }
func (m *AttestedData) String() string  { return (*spire_common.AttestedData)(m).String() }
func (*AttestedData) ProtoMessage()     {}
func (m *AttestedData) GetType() string { return (*spire_common.AttestedData)(m).GetType() }
func (m *AttestedData) GetData() []byte { return (*spire_common.AttestedData)(m).GetData() }

// Selector from public import github.com/spiffe/spire/pkg/common/common.proto
type Selector spire_common.Selector

func (m *Selector) Reset()           { (*spire_common.Selector)(m).Reset() }
func (m *Selector) String() string   { return (*spire_common.Selector)(m).String() }
func (*Selector) ProtoMessage()      {}
func (m *Selector) GetType() string  { return (*spire_common.Selector)(m).GetType() }
func (m *Selector) GetValue() string { return (*spire_common.Selector)(m).GetValue() }

// Selectors from public import github.com/spiffe/spire/pkg/common/common.proto
type Selectors spire_common.Selectors

func (m *Selectors) Reset()         { (*spire_common.Selectors)(m).Reset() }
func (m *Selectors) String() string { return (*spire_common.Selectors)(m).String() }
func (*Selectors) ProtoMessage()    {}
func (m *Selectors) GetEntries() []*Selector {
	o := (*spire_common.Selectors)(m).GetEntries()
	if o == nil {
		return nil
	}
	s := make([]*Selector, len(o))
	for i, x := range o {
		s[i] = (*Selector)(x)
	}
	return s
}

// RegistrationEntry from public import github.com/spiffe/spire/pkg/common/common.proto
type RegistrationEntry spire_common.RegistrationEntry

func (m *RegistrationEntry) Reset()         { (*spire_common.RegistrationEntry)(m).Reset() }
func (m *RegistrationEntry) String() string { return (*spire_common.RegistrationEntry)(m).String() }
func (*RegistrationEntry) ProtoMessage()    {}
func (m *RegistrationEntry) GetSelectors() []*Selector {
	o := (*spire_common.RegistrationEntry)(m).GetSelectors()
	if o == nil {
		return nil
	}
	s := make([]*Selector, len(o))
	for i, x := range o {
		s[i] = (*Selector)(x)
	}
	return s
}
func (m *RegistrationEntry) GetParentId() string {
	return (*spire_common.RegistrationEntry)(m).GetParentId()
}
func (m *RegistrationEntry) GetSpiffeId() string {
	return (*spire_common.RegistrationEntry)(m).GetSpiffeId()
}
func (m *RegistrationEntry) GetTtl() int32 { return (*spire_common.RegistrationEntry)(m).GetTtl() }
func (m *RegistrationEntry) GetFbSpiffeIds() []string {
	return (*spire_common.RegistrationEntry)(m).GetFbSpiffeIds()
}

// RegistrationEntries from public import github.com/spiffe/spire/pkg/common/common.proto
type RegistrationEntries spire_common.RegistrationEntries

func (m *RegistrationEntries) Reset()         { (*spire_common.RegistrationEntries)(m).Reset() }
func (m *RegistrationEntries) String() string { return (*spire_common.RegistrationEntries)(m).String() }
func (*RegistrationEntries) ProtoMessage()    {}
func (m *RegistrationEntries) GetEntries() []*RegistrationEntry {
	o := (*spire_common.RegistrationEntries)(m).GetEntries()
	if o == nil {
		return nil
	}
	s := make([]*RegistrationEntry, len(o))
	for i, x := range o {
		s[i] = (*RegistrationEntry)(x)
	}
	return s
}

// * A type that represents the id of an entry.
type RegistrationEntryID struct {
	// * RegistrationEntryID.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *RegistrationEntryID) Reset()                    { *m = RegistrationEntryID{} }
func (m *RegistrationEntryID) String() string            { return proto.CompactTextString(m) }
func (*RegistrationEntryID) ProtoMessage()               {}
func (*RegistrationEntryID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RegistrationEntryID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// * A type that represents a parent Id.
type ParentID struct {
	// * ParentId.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *ParentID) Reset()                    { *m = ParentID{} }
func (m *ParentID) String() string            { return proto.CompactTextString(m) }
func (*ParentID) ProtoMessage()               {}
func (*ParentID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ParentID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// * A type that represents a SPIFFE Id.
type SpiffeID struct {
	// * SpiffeId.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *SpiffeID) Reset()                    { *m = SpiffeID{} }
func (m *SpiffeID) String() string            { return proto.CompactTextString(m) }
func (*SpiffeID) ProtoMessage()               {}
func (*SpiffeID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SpiffeID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// * A type with the id with want to update plus values to modify.
type UpdateEntryRequest struct {
	// * Id of the entry to update.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// * Values in the RegistrationEntry to update.
	Entry *spire_common.RegistrationEntry `protobuf:"bytes,2,opt,name=entry" json:"entry,omitempty"`
}

func (m *UpdateEntryRequest) Reset()                    { *m = UpdateEntryRequest{} }
func (m *UpdateEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateEntryRequest) ProtoMessage()               {}
func (*UpdateEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpdateEntryRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateEntryRequest) GetEntry() *spire_common.RegistrationEntry {
	if m != nil {
		return m.Entry
	}
	return nil
}

// * A CA bundle for a different Trust Domain than the one used and managed by the Control Plane.
type FederatedBundle struct {
	// * A SPIFFE ID that has a Federated Bundle
	SpiffeId string `protobuf:"bytes,1,opt,name=spiffe_id,json=spiffeId" json:"spiffe_id,omitempty"`
	// * A trusted cert bundle that is not part of Control Planes trust domain but belongs to a different Trust Domain
	FederatedBundle []byte `protobuf:"bytes,2,opt,name=federated_bundle,json=federatedBundle,proto3" json:"federated_bundle,omitempty"`
	// * Time to live.
	Ttl int32 `protobuf:"varint,3,opt,name=ttl" json:"ttl,omitempty"`
}

func (m *FederatedBundle) Reset()                    { *m = FederatedBundle{} }
func (m *FederatedBundle) String() string            { return proto.CompactTextString(m) }
func (*FederatedBundle) ProtoMessage()               {}
func (*FederatedBundle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *FederatedBundle) GetSpiffeId() string {
	if m != nil {
		return m.SpiffeId
	}
	return ""
}

func (m *FederatedBundle) GetFederatedBundle() []byte {
	if m != nil {
		return m.FederatedBundle
	}
	return nil
}

func (m *FederatedBundle) GetTtl() int32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

// * It represents a request with a FederatedBundle to create.
type CreateFederatedBundleRequest struct {
	// * A trusted cert bundle that is not part of Control Planes trust domain but belongs to a different Trust Domain.
	FederatedBundle *FederatedBundle `protobuf:"bytes,1,opt,name=federated_bundle,json=federatedBundle" json:"federated_bundle,omitempty"`
}

func (m *CreateFederatedBundleRequest) Reset()                    { *m = CreateFederatedBundleRequest{} }
func (m *CreateFederatedBundleRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateFederatedBundleRequest) ProtoMessage()               {}
func (*CreateFederatedBundleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CreateFederatedBundleRequest) GetFederatedBundle() *FederatedBundle {
	if m != nil {
		return m.FederatedBundle
	}
	return nil
}

// * It represents a reply with a list of FederatedBundle.
type ListFederatedBundlesReply struct {
	// * A list of FederatedBundle.
	Bundles []*FederatedBundle `protobuf:"bytes,1,rep,name=bundles" json:"bundles,omitempty"`
}

func (m *ListFederatedBundlesReply) Reset()                    { *m = ListFederatedBundlesReply{} }
func (m *ListFederatedBundlesReply) String() string            { return proto.CompactTextString(m) }
func (*ListFederatedBundlesReply) ProtoMessage()               {}
func (*ListFederatedBundlesReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListFederatedBundlesReply) GetBundles() []*FederatedBundle {
	if m != nil {
		return m.Bundles
	}
	return nil
}

// * A type that represents a Federated SPIFFE Id.
type FederatedSpiffeID struct {
	// * FederatedSpiffeID
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *FederatedSpiffeID) Reset()                    { *m = FederatedSpiffeID{} }
func (m *FederatedSpiffeID) String() string            { return proto.CompactTextString(m) }
func (*FederatedSpiffeID) ProtoMessage()               {}
func (*FederatedSpiffeID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *FederatedSpiffeID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*RegistrationEntryID)(nil), "spire.api.registration.RegistrationEntryID")
	proto.RegisterType((*ParentID)(nil), "spire.api.registration.ParentID")
	proto.RegisterType((*SpiffeID)(nil), "spire.api.registration.SpiffeID")
	proto.RegisterType((*UpdateEntryRequest)(nil), "spire.api.registration.UpdateEntryRequest")
	proto.RegisterType((*FederatedBundle)(nil), "spire.api.registration.FederatedBundle")
	proto.RegisterType((*CreateFederatedBundleRequest)(nil), "spire.api.registration.CreateFederatedBundleRequest")
	proto.RegisterType((*ListFederatedBundlesReply)(nil), "spire.api.registration.ListFederatedBundlesReply")
	proto.RegisterType((*FederatedSpiffeID)(nil), "spire.api.registration.FederatedSpiffeID")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Registration service

type RegistrationClient interface {
	// * Creates an entry in the Registration table, used to assign SPIFFE IDs to nodes and workloads.
	CreateEntry(ctx context.Context, in *spire_common.RegistrationEntry, opts ...grpc.CallOption) (*RegistrationEntryID, error)
	// * Deletes an entry and returns the deleted entry.
	DeleteEntry(ctx context.Context, in *RegistrationEntryID, opts ...grpc.CallOption) (*spire_common.RegistrationEntry, error)
	// * Retrieve a specific registered entry.
	FetchEntry(ctx context.Context, in *RegistrationEntryID, opts ...grpc.CallOption) (*spire_common.RegistrationEntry, error)
	// * Updates a specific registered entry.
	UpdateEntry(ctx context.Context, in *UpdateEntryRequest, opts ...grpc.CallOption) (*spire_common.RegistrationEntry, error)
	// * Returns all the Entries associated with the ParentID value.
	ListByParentID(ctx context.Context, in *ParentID, opts ...grpc.CallOption) (*spire_common.RegistrationEntries, error)
	// * Returns all the entries associated with a selector value.
	ListBySelector(ctx context.Context, in *spire_common.Selector, opts ...grpc.CallOption) (*spire_common.RegistrationEntries, error)
	// * Return all registration entries for which SPIFFE ID matches.
	ListBySpiffeID(ctx context.Context, in *SpiffeID, opts ...grpc.CallOption) (*spire_common.RegistrationEntries, error)
	// * Creates an entry in the Federated bundle table to store the mappings of Federated SPIFFE IDs and their associated CA bundle.
	CreateFederatedBundle(ctx context.Context, in *CreateFederatedBundleRequest, opts ...grpc.CallOption) (*spire_common.Empty, error)
	// * Retrieves Federated bundles for all the Federated SPIFFE IDs.
	ListFederatedBundles(ctx context.Context, in *spire_common.Empty, opts ...grpc.CallOption) (*ListFederatedBundlesReply, error)
	// * Updates a particular Federated Bundle. Useful for rotation.
	UpdateFederatedBundle(ctx context.Context, in *FederatedBundle, opts ...grpc.CallOption) (*spire_common.Empty, error)
	// * Delete a particular Federated Bundle. Used to destroy inter-domain trust.
	DeleteFederatedBundle(ctx context.Context, in *FederatedSpiffeID, opts ...grpc.CallOption) (*spire_common.Empty, error)
}

type registrationClient struct {
	cc *grpc.ClientConn
}

func NewRegistrationClient(cc *grpc.ClientConn) RegistrationClient {
	return &registrationClient{cc}
}

func (c *registrationClient) CreateEntry(ctx context.Context, in *spire_common.RegistrationEntry, opts ...grpc.CallOption) (*RegistrationEntryID, error) {
	out := new(RegistrationEntryID)
	err := grpc.Invoke(ctx, "/spire.api.registration.Registration/CreateEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationClient) DeleteEntry(ctx context.Context, in *RegistrationEntryID, opts ...grpc.CallOption) (*spire_common.RegistrationEntry, error) {
	out := new(spire_common.RegistrationEntry)
	err := grpc.Invoke(ctx, "/spire.api.registration.Registration/DeleteEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationClient) FetchEntry(ctx context.Context, in *RegistrationEntryID, opts ...grpc.CallOption) (*spire_common.RegistrationEntry, error) {
	out := new(spire_common.RegistrationEntry)
	err := grpc.Invoke(ctx, "/spire.api.registration.Registration/FetchEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationClient) UpdateEntry(ctx context.Context, in *UpdateEntryRequest, opts ...grpc.CallOption) (*spire_common.RegistrationEntry, error) {
	out := new(spire_common.RegistrationEntry)
	err := grpc.Invoke(ctx, "/spire.api.registration.Registration/UpdateEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationClient) ListByParentID(ctx context.Context, in *ParentID, opts ...grpc.CallOption) (*spire_common.RegistrationEntries, error) {
	out := new(spire_common.RegistrationEntries)
	err := grpc.Invoke(ctx, "/spire.api.registration.Registration/ListByParentID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationClient) ListBySelector(ctx context.Context, in *spire_common.Selector, opts ...grpc.CallOption) (*spire_common.RegistrationEntries, error) {
	out := new(spire_common.RegistrationEntries)
	err := grpc.Invoke(ctx, "/spire.api.registration.Registration/ListBySelector", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationClient) ListBySpiffeID(ctx context.Context, in *SpiffeID, opts ...grpc.CallOption) (*spire_common.RegistrationEntries, error) {
	out := new(spire_common.RegistrationEntries)
	err := grpc.Invoke(ctx, "/spire.api.registration.Registration/ListBySpiffeID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationClient) CreateFederatedBundle(ctx context.Context, in *CreateFederatedBundleRequest, opts ...grpc.CallOption) (*spire_common.Empty, error) {
	out := new(spire_common.Empty)
	err := grpc.Invoke(ctx, "/spire.api.registration.Registration/CreateFederatedBundle", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationClient) ListFederatedBundles(ctx context.Context, in *spire_common.Empty, opts ...grpc.CallOption) (*ListFederatedBundlesReply, error) {
	out := new(ListFederatedBundlesReply)
	err := grpc.Invoke(ctx, "/spire.api.registration.Registration/ListFederatedBundles", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationClient) UpdateFederatedBundle(ctx context.Context, in *FederatedBundle, opts ...grpc.CallOption) (*spire_common.Empty, error) {
	out := new(spire_common.Empty)
	err := grpc.Invoke(ctx, "/spire.api.registration.Registration/UpdateFederatedBundle", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationClient) DeleteFederatedBundle(ctx context.Context, in *FederatedSpiffeID, opts ...grpc.CallOption) (*spire_common.Empty, error) {
	out := new(spire_common.Empty)
	err := grpc.Invoke(ctx, "/spire.api.registration.Registration/DeleteFederatedBundle", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Registration service

type RegistrationServer interface {
	// * Creates an entry in the Registration table, used to assign SPIFFE IDs to nodes and workloads.
	CreateEntry(context.Context, *spire_common.RegistrationEntry) (*RegistrationEntryID, error)
	// * Deletes an entry and returns the deleted entry.
	DeleteEntry(context.Context, *RegistrationEntryID) (*spire_common.RegistrationEntry, error)
	// * Retrieve a specific registered entry.
	FetchEntry(context.Context, *RegistrationEntryID) (*spire_common.RegistrationEntry, error)
	// * Updates a specific registered entry.
	UpdateEntry(context.Context, *UpdateEntryRequest) (*spire_common.RegistrationEntry, error)
	// * Returns all the Entries associated with the ParentID value.
	ListByParentID(context.Context, *ParentID) (*spire_common.RegistrationEntries, error)
	// * Returns all the entries associated with a selector value.
	ListBySelector(context.Context, *spire_common.Selector) (*spire_common.RegistrationEntries, error)
	// * Return all registration entries for which SPIFFE ID matches.
	ListBySpiffeID(context.Context, *SpiffeID) (*spire_common.RegistrationEntries, error)
	// * Creates an entry in the Federated bundle table to store the mappings of Federated SPIFFE IDs and their associated CA bundle.
	CreateFederatedBundle(context.Context, *CreateFederatedBundleRequest) (*spire_common.Empty, error)
	// * Retrieves Federated bundles for all the Federated SPIFFE IDs.
	ListFederatedBundles(context.Context, *spire_common.Empty) (*ListFederatedBundlesReply, error)
	// * Updates a particular Federated Bundle. Useful for rotation.
	UpdateFederatedBundle(context.Context, *FederatedBundle) (*spire_common.Empty, error)
	// * Delete a particular Federated Bundle. Used to destroy inter-domain trust.
	DeleteFederatedBundle(context.Context, *FederatedSpiffeID) (*spire_common.Empty, error)
}

func RegisterRegistrationServer(s *grpc.Server, srv RegistrationServer) {
	s.RegisterService(&_Registration_serviceDesc, srv)
}

func _Registration_CreateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spire_common.RegistrationEntry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).CreateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.registration.Registration/CreateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).CreateEntry(ctx, req.(*spire_common.RegistrationEntry))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registration_DeleteEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationEntryID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).DeleteEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.registration.Registration/DeleteEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).DeleteEntry(ctx, req.(*RegistrationEntryID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registration_FetchEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationEntryID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).FetchEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.registration.Registration/FetchEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).FetchEntry(ctx, req.(*RegistrationEntryID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registration_UpdateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).UpdateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.registration.Registration/UpdateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).UpdateEntry(ctx, req.(*UpdateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registration_ListByParentID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParentID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).ListByParentID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.registration.Registration/ListByParentID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).ListByParentID(ctx, req.(*ParentID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registration_ListBySelector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spire_common.Selector)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).ListBySelector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.registration.Registration/ListBySelector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).ListBySelector(ctx, req.(*spire_common.Selector))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registration_ListBySpiffeID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpiffeID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).ListBySpiffeID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.registration.Registration/ListBySpiffeID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).ListBySpiffeID(ctx, req.(*SpiffeID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registration_CreateFederatedBundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFederatedBundleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).CreateFederatedBundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.registration.Registration/CreateFederatedBundle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).CreateFederatedBundle(ctx, req.(*CreateFederatedBundleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registration_ListFederatedBundles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spire_common.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).ListFederatedBundles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.registration.Registration/ListFederatedBundles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).ListFederatedBundles(ctx, req.(*spire_common.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registration_UpdateFederatedBundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FederatedBundle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).UpdateFederatedBundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.registration.Registration/UpdateFederatedBundle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).UpdateFederatedBundle(ctx, req.(*FederatedBundle))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registration_DeleteFederatedBundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FederatedSpiffeID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).DeleteFederatedBundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.registration.Registration/DeleteFederatedBundle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).DeleteFederatedBundle(ctx, req.(*FederatedSpiffeID))
	}
	return interceptor(ctx, in, info, handler)
}

var _Registration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.api.registration.Registration",
	HandlerType: (*RegistrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEntry",
			Handler:    _Registration_CreateEntry_Handler,
		},
		{
			MethodName: "DeleteEntry",
			Handler:    _Registration_DeleteEntry_Handler,
		},
		{
			MethodName: "FetchEntry",
			Handler:    _Registration_FetchEntry_Handler,
		},
		{
			MethodName: "UpdateEntry",
			Handler:    _Registration_UpdateEntry_Handler,
		},
		{
			MethodName: "ListByParentID",
			Handler:    _Registration_ListByParentID_Handler,
		},
		{
			MethodName: "ListBySelector",
			Handler:    _Registration_ListBySelector_Handler,
		},
		{
			MethodName: "ListBySpiffeID",
			Handler:    _Registration_ListBySpiffeID_Handler,
		},
		{
			MethodName: "CreateFederatedBundle",
			Handler:    _Registration_CreateFederatedBundle_Handler,
		},
		{
			MethodName: "ListFederatedBundles",
			Handler:    _Registration_ListFederatedBundles_Handler,
		},
		{
			MethodName: "UpdateFederatedBundle",
			Handler:    _Registration_UpdateFederatedBundle_Handler,
		},
		{
			MethodName: "DeleteFederatedBundle",
			Handler:    _Registration_DeleteFederatedBundle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registration.proto",
}

func init() { proto.RegisterFile("registration.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 572 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x71, 0xa2, 0x96, 0x74, 0x1c, 0xd2, 0x76, 0x43, 0xaa, 0x60, 0x2a, 0x11, 0x8c, 0x10,
	0x69, 0x10, 0xb6, 0x28, 0x70, 0xe1, 0x46, 0x68, 0x2b, 0x55, 0x70, 0x88, 0x5c, 0x45, 0x48, 0x20,
	0xb5, 0x72, 0xe2, 0x49, 0xba, 0xc4, 0xf1, 0x9a, 0xf5, 0xe6, 0x10, 0x21, 0x2e, 0xbc, 0x02, 0x0f,
	0xc3, 0x83, 0xf0, 0x0a, 0x3c, 0x08, 0xca, 0x6e, 0x6c, 0xb9, 0xae, 0xdd, 0xa4, 0x07, 0x4e, 0x89,
	0x76, 0x66, 0xfe, 0x6f, 0x26, 0x33, 0xbf, 0x02, 0x84, 0xe3, 0x98, 0x46, 0x82, 0xbb, 0x82, 0xb2,
	0xc0, 0x0a, 0x39, 0x13, 0x8c, 0xec, 0x45, 0x21, 0xe5, 0x68, 0xb9, 0x21, 0xb5, 0xd2, 0x51, 0x63,
	0x7f, 0xcc, 0xd8, 0xd8, 0x47, 0xdb, 0x0d, 0xa9, 0xed, 0x06, 0x01, 0x13, 0xf2, 0x39, 0x52, 0x55,
	0xc6, 0x8b, 0x31, 0x15, 0x97, 0xb3, 0x81, 0x35, 0x64, 0x53, 0x3b, 0x0a, 0xe9, 0x68, 0x84, 0x76,
	0xc4, 0xa9, 0x1d, 0x4e, 0xc6, 0xf6, 0x90, 0x4d, 0xa7, 0x2c, 0x58, 0x7e, 0xa8, 0x74, 0xf3, 0x29,
	0xd4, 0x9d, 0x94, 0xf8, 0x71, 0x20, 0xf8, 0xfc, 0xf4, 0x88, 0xd4, 0xa0, 0x44, 0xbd, 0xa6, 0xd6,
	0xd2, 0xda, 0x5b, 0x4e, 0x89, 0x7a, 0xa6, 0x01, 0x95, 0x9e, 0xcb, 0x31, 0x10, 0xf9, 0xb1, 0x33,
	0x09, 0xca, 0x89, 0x7d, 0x01, 0xd2, 0x0f, 0x3d, 0x57, 0xa0, 0x14, 0x76, 0xf0, 0xdb, 0x0c, 0x23,
	0x91, 0xcd, 0x22, 0x6f, 0x60, 0x03, 0x17, 0xf1, 0x66, 0xa9, 0xa5, 0xb5, 0xf5, 0xc3, 0x47, 0x96,
	0x9a, 0x7c, 0xd9, 0xe8, 0xb5, 0xfe, 0x1c, 0x95, 0x6d, 0x4e, 0x60, 0xfb, 0x04, 0x3d, 0xe4, 0xae,
	0x40, 0xaf, 0x3b, 0x0b, 0x3c, 0x1f, 0xc9, 0x43, 0xd8, 0x52, 0x43, 0x5f, 0x24, 0x80, 0x8a, 0x7a,
	0x38, 0xf5, 0xc8, 0x01, 0xec, 0x8c, 0xe2, 0xfc, 0x8b, 0x81, 0x2c, 0x90, 0xc4, 0xaa, 0xb3, 0x3d,
	0xca, 0xe8, 0xec, 0x40, 0x59, 0x08, 0xbf, 0x59, 0x6e, 0x69, 0xed, 0x0d, 0x67, 0xf1, 0xd5, 0xe4,
	0xb0, 0xff, 0x9e, 0xa3, 0x2b, 0x30, 0x83, 0x8c, 0x67, 0x72, 0x72, 0xc4, 0x35, 0x39, 0xce, 0x33,
	0x2b, 0x7f, 0x91, 0x56, 0x56, 0x29, 0xdb, 0x85, 0x79, 0x0e, 0x0f, 0x3e, 0xd2, 0x48, 0x64, 0xf2,
	0x22, 0x07, 0x43, 0x7f, 0x4e, 0xde, 0xc1, 0x5d, 0x85, 0x89, 0x9a, 0x5a, 0xab, 0x7c, 0x1b, 0x4e,
	0x5c, 0x67, 0x3e, 0x81, 0xdd, 0x24, 0x56, 0xb4, 0xc2, 0xc3, 0xdf, 0x15, 0xa8, 0xa6, 0x57, 0x40,
	0x02, 0xd0, 0xd5, 0x2f, 0x21, 0x97, 0x41, 0x56, 0x6d, 0xcb, 0x78, 0x5e, 0xd4, 0x57, 0xce, 0xe1,
	0x99, 0xbb, 0x3f, 0xff, 0xfc, 0xfd, 0x55, 0xd2, 0xcd, 0x4d, 0x5b, 0xee, 0xf8, 0xad, 0xd6, 0x21,
	0x13, 0xd0, 0x8f, 0xd0, 0xc7, 0x98, 0x77, 0x1b, 0x39, 0x63, 0x55, 0x73, 0x66, 0x4d, 0xf2, 0x2a,
	0x9d, 0x25, 0x8f, 0x30, 0x80, 0x13, 0x14, 0xc3, 0xcb, 0xff, 0xc1, 0xaa, 0x4b, 0xd6, 0x3d, 0xa2,
	0x2b, 0x96, 0xfd, 0x9d, 0x7a, 0x3f, 0xc8, 0x57, 0xd0, 0x53, 0x0e, 0x21, 0x9d, 0x22, 0xe2, 0x75,
	0x1b, 0xad, 0x3d, 0x9c, 0x11, 0x0f, 0xd7, 0x87, 0xda, 0xe2, 0x9e, 0xba, 0xf3, 0xc4, 0xcb, 0xad,
	0x22, 0x5c, 0x9c, 0x61, 0x3c, 0xbe, 0x19, 0x42, 0x31, 0x22, 0x1f, 0x62, 0xd9, 0x33, 0xf4, 0x71,
	0x28, 0x18, 0x27, 0x7b, 0x57, 0x8b, 0xe2, 0xf7, 0x75, 0xc4, 0x92, 0x1e, 0x93, 0x83, 0x2c, 0xec,
	0x31, 0xce, 0x58, 0x47, 0x76, 0x00, 0x8d, 0x5c, 0xfb, 0x92, 0xd7, 0x45, 0xea, 0x37, 0xb9, 0xdd,
	0xa8, 0x5f, 0x25, 0x1e, 0x4f, 0x43, 0x31, 0x27, 0xe7, 0x70, 0x3f, 0xcf, 0xae, 0x24, 0x2f, 0xd9,
	0x78, 0x59, 0xc4, 0x2d, 0x76, 0x7c, 0x1f, 0x1a, 0xea, 0x0a, 0xb2, 0x33, 0xac, 0xeb, 0xfc, 0xfc,
	0xb6, 0x3f, 0x41, 0x43, 0xf9, 0x2b, 0x2b, 0x7b, 0xb0, 0x52, 0x36, 0xd9, 0x40, 0x9e, 0x70, 0xb7,
	0xf6, 0xb9, 0x9a, 0x2e, 0xeb, 0xdd, 0xe9, 0x69, 0x83, 0x4d, 0xf9, 0xb7, 0xf3, 0xea, 0x5f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x81, 0x6e, 0x5b, 0xfe, 0xf1, 0x06, 0x00, 0x00,
}