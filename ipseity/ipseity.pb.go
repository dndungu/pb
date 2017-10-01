// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipseity.proto

/*
Package ipseity is a generated protocol buffer package.

It is generated from these files:
	ipseity.proto

It has these top-level messages:
	Credentials
	Grant
	Role
	Identity
	Empty
	UpdateIdentitySecretRequest
	UpdateIdentityRolesRequest
	Id
	ListFilter
*/
package ipseity

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Credentials_Provider int32

const (
	Credentials_BASIC    Credentials_Provider = 0
	Credentials_PASSWORD Credentials_Provider = 1
	Credentials_FACEBOOK Credentials_Provider = 2
	Credentials_GOOGLE   Credentials_Provider = 3
	Credentials_TWITTER  Credentials_Provider = 4
)

var Credentials_Provider_name = map[int32]string{
	0: "BASIC",
	1: "PASSWORD",
	2: "FACEBOOK",
	3: "GOOGLE",
	4: "TWITTER",
}
var Credentials_Provider_value = map[string]int32{
	"BASIC":    0,
	"PASSWORD": 1,
	"FACEBOOK": 2,
	"GOOGLE":   3,
	"TWITTER":  4,
}

func (x Credentials_Provider) String() string {
	return proto.EnumName(Credentials_Provider_name, int32(x))
}
func (Credentials_Provider) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Credentials struct {
	Key      string               `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Secret   string               `protobuf:"bytes,2,opt,name=secret" json:"secret,omitempty"`
	Provider Credentials_Provider `protobuf:"varint,3,opt,name=provider,enum=ipseity.Credentials_Provider" json:"provider,omitempty"`
}

func (m *Credentials) Reset()                    { *m = Credentials{} }
func (m *Credentials) String() string            { return proto.CompactTextString(m) }
func (*Credentials) ProtoMessage()               {}
func (*Credentials) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Credentials) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Credentials) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *Credentials) GetProvider() Credentials_Provider {
	if m != nil {
		return m.Provider
	}
	return Credentials_BASIC
}

type Grant struct {
	Id    int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Notes string `protobuf:"bytes,3,opt,name=notes" json:"notes,omitempty"`
}

func (m *Grant) Reset()                    { *m = Grant{} }
func (m *Grant) String() string            { return proto.CompactTextString(m) }
func (*Grant) ProtoMessage()               {}
func (*Grant) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Grant) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Grant) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Grant) GetNotes() string {
	if m != nil {
		return m.Notes
	}
	return ""
}

type Role struct {
	Id     int32    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name   string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Tags   []string `protobuf:"bytes,3,rep,name=tags" json:"tags,omitempty"`
	Grants []*Grant `protobuf:"bytes,4,rep,name=grants" json:"grants,omitempty"`
}

func (m *Role) Reset()                    { *m = Role{} }
func (m *Role) String() string            { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()               {}
func (*Role) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Role) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Role) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Role) GetGrants() []*Grant {
	if m != nil {
		return m.Grants
	}
	return nil
}

type Identity struct {
	Id         int32        `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Credential *Credentials `protobuf:"bytes,2,opt,name=credential" json:"credential,omitempty"`
	RoleIds    []int32      `protobuf:"varint,3,rep,packed,name=roleIds" json:"roleIds,omitempty"`
}

func (m *Identity) Reset()                    { *m = Identity{} }
func (m *Identity) String() string            { return proto.CompactTextString(m) }
func (*Identity) ProtoMessage()               {}
func (*Identity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Identity) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Identity) GetCredential() *Credentials {
	if m != nil {
		return m.Credential
	}
	return nil
}

func (m *Identity) GetRoleIds() []int32 {
	if m != nil {
		return m.RoleIds
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type UpdateIdentitySecretRequest struct {
	Id     int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Secret string `protobuf:"bytes,2,opt,name=secret" json:"secret,omitempty"`
}

func (m *UpdateIdentitySecretRequest) Reset()                    { *m = UpdateIdentitySecretRequest{} }
func (m *UpdateIdentitySecretRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateIdentitySecretRequest) ProtoMessage()               {}
func (*UpdateIdentitySecretRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateIdentitySecretRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateIdentitySecretRequest) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

type UpdateIdentityRolesRequest struct {
	Id    int32   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Roles []int32 `protobuf:"varint,2,rep,packed,name=roles" json:"roles,omitempty"`
}

func (m *UpdateIdentityRolesRequest) Reset()                    { *m = UpdateIdentityRolesRequest{} }
func (m *UpdateIdentityRolesRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateIdentityRolesRequest) ProtoMessage()               {}
func (*UpdateIdentityRolesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UpdateIdentityRolesRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateIdentityRolesRequest) GetRoles() []int32 {
	if m != nil {
		return m.Roles
	}
	return nil
}

type Id struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *Id) Reset()                    { *m = Id{} }
func (m *Id) String() string            { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()               {}
func (*Id) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Id) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ListFilter struct {
	Offset int32 `protobuf:"varint,1,opt,name=offset" json:"offset,omitempty"`
	Limit  int32 `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
}

func (m *ListFilter) Reset()                    { *m = ListFilter{} }
func (m *ListFilter) String() string            { return proto.CompactTextString(m) }
func (*ListFilter) ProtoMessage()               {}
func (*ListFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ListFilter) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListFilter) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func init() {
	proto.RegisterType((*Credentials)(nil), "ipseity.Credentials")
	proto.RegisterType((*Grant)(nil), "ipseity.Grant")
	proto.RegisterType((*Role)(nil), "ipseity.Role")
	proto.RegisterType((*Identity)(nil), "ipseity.Identity")
	proto.RegisterType((*Empty)(nil), "ipseity.Empty")
	proto.RegisterType((*UpdateIdentitySecretRequest)(nil), "ipseity.UpdateIdentitySecretRequest")
	proto.RegisterType((*UpdateIdentityRolesRequest)(nil), "ipseity.UpdateIdentityRolesRequest")
	proto.RegisterType((*Id)(nil), "ipseity.Id")
	proto.RegisterType((*ListFilter)(nil), "ipseity.ListFilter")
	proto.RegisterEnum("ipseity.Credentials_Provider", Credentials_Provider_name, Credentials_Provider_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Ipseity service

type IpseityClient interface {
	CreateIdentity(ctx context.Context, in *Identity, opts ...grpc.CallOption) (*Id, error)
	UpdateIdentitySecret(ctx context.Context, in *UpdateIdentitySecretRequest, opts ...grpc.CallOption) (*Empty, error)
	UpdateIdentityRoles(ctx context.Context, in *UpdateIdentityRolesRequest, opts ...grpc.CallOption) (*Empty, error)
	FindIdentityById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Identity, error)
	FindIdentityByCredentials(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*Identity, error)
	DeleteIdentity(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error)
	CreateRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Id, error)
	FindRoleById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Role, error)
	ListRole(ctx context.Context, in *ListFilter, opts ...grpc.CallOption) (Ipseity_ListRoleClient, error)
	UpdateRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Empty, error)
	DeleteRole(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error)
}

type ipseityClient struct {
	cc *grpc.ClientConn
}

func NewIpseityClient(cc *grpc.ClientConn) IpseityClient {
	return &ipseityClient{cc}
}

func (c *ipseityClient) CreateIdentity(ctx context.Context, in *Identity, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := grpc.Invoke(ctx, "/ipseity.Ipseity/CreateIdentity", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ipseityClient) UpdateIdentitySecret(ctx context.Context, in *UpdateIdentitySecretRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/ipseity.Ipseity/UpdateIdentitySecret", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ipseityClient) UpdateIdentityRoles(ctx context.Context, in *UpdateIdentityRolesRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/ipseity.Ipseity/UpdateIdentityRoles", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ipseityClient) FindIdentityById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Identity, error) {
	out := new(Identity)
	err := grpc.Invoke(ctx, "/ipseity.Ipseity/FindIdentityById", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ipseityClient) FindIdentityByCredentials(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*Identity, error) {
	out := new(Identity)
	err := grpc.Invoke(ctx, "/ipseity.Ipseity/FindIdentityByCredentials", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ipseityClient) DeleteIdentity(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/ipseity.Ipseity/DeleteIdentity", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ipseityClient) CreateRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := grpc.Invoke(ctx, "/ipseity.Ipseity/CreateRole", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ipseityClient) FindRoleById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := grpc.Invoke(ctx, "/ipseity.Ipseity/FindRoleById", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ipseityClient) ListRole(ctx context.Context, in *ListFilter, opts ...grpc.CallOption) (Ipseity_ListRoleClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Ipseity_serviceDesc.Streams[0], c.cc, "/ipseity.Ipseity/ListRole", opts...)
	if err != nil {
		return nil, err
	}
	x := &ipseityListRoleClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Ipseity_ListRoleClient interface {
	Recv() (*Role, error)
	grpc.ClientStream
}

type ipseityListRoleClient struct {
	grpc.ClientStream
}

func (x *ipseityListRoleClient) Recv() (*Role, error) {
	m := new(Role)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ipseityClient) UpdateRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/ipseity.Ipseity/UpdateRole", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ipseityClient) DeleteRole(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/ipseity.Ipseity/DeleteRole", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Ipseity service

type IpseityServer interface {
	CreateIdentity(context.Context, *Identity) (*Id, error)
	UpdateIdentitySecret(context.Context, *UpdateIdentitySecretRequest) (*Empty, error)
	UpdateIdentityRoles(context.Context, *UpdateIdentityRolesRequest) (*Empty, error)
	FindIdentityById(context.Context, *Id) (*Identity, error)
	FindIdentityByCredentials(context.Context, *Credentials) (*Identity, error)
	DeleteIdentity(context.Context, *Id) (*Empty, error)
	CreateRole(context.Context, *Role) (*Id, error)
	FindRoleById(context.Context, *Id) (*Role, error)
	ListRole(*ListFilter, Ipseity_ListRoleServer) error
	UpdateRole(context.Context, *Role) (*Empty, error)
	DeleteRole(context.Context, *Id) (*Empty, error)
}

func RegisterIpseityServer(s *grpc.Server, srv IpseityServer) {
	s.RegisterService(&_Ipseity_serviceDesc, srv)
}

func _Ipseity_CreateIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Identity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IpseityServer).CreateIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ipseity.Ipseity/CreateIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IpseityServer).CreateIdentity(ctx, req.(*Identity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ipseity_UpdateIdentitySecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIdentitySecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IpseityServer).UpdateIdentitySecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ipseity.Ipseity/UpdateIdentitySecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IpseityServer).UpdateIdentitySecret(ctx, req.(*UpdateIdentitySecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ipseity_UpdateIdentityRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIdentityRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IpseityServer).UpdateIdentityRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ipseity.Ipseity/UpdateIdentityRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IpseityServer).UpdateIdentityRoles(ctx, req.(*UpdateIdentityRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ipseity_FindIdentityById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IpseityServer).FindIdentityById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ipseity.Ipseity/FindIdentityById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IpseityServer).FindIdentityById(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ipseity_FindIdentityByCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Credentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IpseityServer).FindIdentityByCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ipseity.Ipseity/FindIdentityByCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IpseityServer).FindIdentityByCredentials(ctx, req.(*Credentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ipseity_DeleteIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IpseityServer).DeleteIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ipseity.Ipseity/DeleteIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IpseityServer).DeleteIdentity(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ipseity_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Role)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IpseityServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ipseity.Ipseity/CreateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IpseityServer).CreateRole(ctx, req.(*Role))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ipseity_FindRoleById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IpseityServer).FindRoleById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ipseity.Ipseity/FindRoleById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IpseityServer).FindRoleById(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ipseity_ListRole_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IpseityServer).ListRole(m, &ipseityListRoleServer{stream})
}

type Ipseity_ListRoleServer interface {
	Send(*Role) error
	grpc.ServerStream
}

type ipseityListRoleServer struct {
	grpc.ServerStream
}

func (x *ipseityListRoleServer) Send(m *Role) error {
	return x.ServerStream.SendMsg(m)
}

func _Ipseity_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Role)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IpseityServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ipseity.Ipseity/UpdateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IpseityServer).UpdateRole(ctx, req.(*Role))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ipseity_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IpseityServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ipseity.Ipseity/DeleteRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IpseityServer).DeleteRole(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ipseity_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ipseity.Ipseity",
	HandlerType: (*IpseityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateIdentity",
			Handler:    _Ipseity_CreateIdentity_Handler,
		},
		{
			MethodName: "UpdateIdentitySecret",
			Handler:    _Ipseity_UpdateIdentitySecret_Handler,
		},
		{
			MethodName: "UpdateIdentityRoles",
			Handler:    _Ipseity_UpdateIdentityRoles_Handler,
		},
		{
			MethodName: "FindIdentityById",
			Handler:    _Ipseity_FindIdentityById_Handler,
		},
		{
			MethodName: "FindIdentityByCredentials",
			Handler:    _Ipseity_FindIdentityByCredentials_Handler,
		},
		{
			MethodName: "DeleteIdentity",
			Handler:    _Ipseity_DeleteIdentity_Handler,
		},
		{
			MethodName: "CreateRole",
			Handler:    _Ipseity_CreateRole_Handler,
		},
		{
			MethodName: "FindRoleById",
			Handler:    _Ipseity_FindRoleById_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _Ipseity_UpdateRole_Handler,
		},
		{
			MethodName: "DeleteRole",
			Handler:    _Ipseity_DeleteRole_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListRole",
			Handler:       _Ipseity_ListRole_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ipseity.proto",
}

func init() { proto.RegisterFile("ipseity.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 572 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5f, 0x8f, 0xd2, 0x4e,
	0x14, 0xa5, 0xa5, 0x05, 0xf6, 0xb2, 0x4b, 0xfa, 0xbb, 0x4b, 0x7e, 0xa9, 0x18, 0x13, 0x32, 0x1a,
	0x43, 0xfc, 0x83, 0xa6, 0xee, 0x8b, 0xbe, 0x01, 0x0b, 0xa4, 0x4a, 0x64, 0x33, 0x60, 0xf6, 0x19,
	0xb7, 0xc3, 0x66, 0xb4, 0x50, 0x6c, 0x47, 0x13, 0xbe, 0x9c, 0x9f, 0xca, 0x0f, 0x60, 0x66, 0xfa,
	0x67, 0xcb, 0xb6, 0x6c, 0x7c, 0x9b, 0x73, 0xe7, 0xcc, 0xb9, 0xe7, 0x9e, 0x99, 0x16, 0xce, 0xf8,
	0x2e, 0x62, 0x5c, 0xec, 0xfb, 0xbb, 0x30, 0x10, 0x01, 0xd6, 0x13, 0x48, 0x7e, 0x6b, 0xd0, 0x1c,
	0x85, 0xcc, 0x63, 0x5b, 0xc1, 0x57, 0x7e, 0x84, 0x16, 0x54, 0xbf, 0xb3, 0xbd, 0xad, 0x75, 0xb5,
	0xde, 0x09, 0x95, 0x4b, 0xfc, 0x1f, 0x6a, 0x11, 0xbb, 0x09, 0x99, 0xb0, 0x75, 0x55, 0x4c, 0x10,
	0xbe, 0x87, 0xc6, 0x2e, 0x0c, 0x7e, 0x71, 0x8f, 0x85, 0x76, 0xb5, 0xab, 0xf5, 0x5a, 0xce, 0x93,
	0x7e, 0xda, 0x24, 0xa7, 0xd8, 0xbf, 0x4a, 0x48, 0x34, 0xa3, 0x93, 0x8f, 0xd0, 0x48, 0xab, 0x78,
	0x02, 0xe6, 0x70, 0xb0, 0x70, 0x47, 0x56, 0x05, 0x4f, 0xa1, 0x71, 0x35, 0x58, 0x2c, 0xae, 0xe7,
	0xf4, 0xd2, 0xd2, 0x24, 0x9a, 0x0c, 0x46, 0xe3, 0xe1, 0x7c, 0xfe, 0xc9, 0xd2, 0x11, 0xa0, 0x36,
	0x9d, 0xcf, 0xa7, 0xb3, 0xb1, 0x55, 0xc5, 0x26, 0xd4, 0x97, 0xd7, 0xee, 0x72, 0x39, 0xa6, 0x96,
	0x41, 0x06, 0x60, 0x4e, 0xc3, 0xd5, 0x56, 0x60, 0x0b, 0x74, 0xee, 0x29, 0xe3, 0x26, 0xd5, 0xb9,
	0x87, 0x08, 0xc6, 0x76, 0xb5, 0x61, 0x89, 0x6b, 0xb5, 0xc6, 0x36, 0x98, 0xdb, 0x40, 0xb0, 0x48,
	0x19, 0x3e, 0xa1, 0x31, 0x20, 0x6b, 0x30, 0x68, 0xe0, 0xb3, 0x7f, 0x52, 0x40, 0x30, 0xc4, 0xea,
	0x56, 0x0a, 0x54, 0x65, 0x4d, 0xae, 0xf1, 0x39, 0xd4, 0x6e, 0xa5, 0x85, 0xc8, 0x36, 0xba, 0xd5,
	0x5e, 0xd3, 0x69, 0x65, 0x39, 0x28, 0x67, 0x34, 0xd9, 0x25, 0xdf, 0xa0, 0xe1, 0xaa, 0x58, 0xc4,
	0xbe, 0xd0, 0xeb, 0x02, 0xe0, 0x26, 0x0b, 0x4d, 0x75, 0x6c, 0x3a, 0xed, 0xb2, 0x3c, 0x69, 0x8e,
	0x87, 0x36, 0xd4, 0xc3, 0xc0, 0x67, 0xae, 0x17, 0x1b, 0x32, 0x69, 0x0a, 0x49, 0x1d, 0xcc, 0xf1,
	0x66, 0x27, 0xf6, 0x64, 0x0c, 0x8f, 0xbf, 0xec, 0xbc, 0x95, 0x60, 0x69, 0xeb, 0x85, 0xba, 0x3e,
	0xca, 0x7e, 0xfc, 0x64, 0x51, 0x31, 0xb5, 0x23, 0xb7, 0x4d, 0x86, 0xd0, 0x39, 0x94, 0x91, 0x89,
	0x45, 0xc7, 0x54, 0xda, 0x60, 0x4a, 0x23, 0x91, 0xad, 0x2b, 0x57, 0x31, 0x20, 0x6d, 0xd0, 0x5d,
	0xef, 0x3e, 0x97, 0x7c, 0x00, 0x98, 0xf1, 0x48, 0x4c, 0xb8, 0x2f, 0x58, 0x28, 0xfb, 0x07, 0xeb,
	0x75, 0xc4, 0x44, 0xc2, 0x48, 0x90, 0x54, 0xf4, 0xf9, 0x86, 0xc7, 0xb6, 0x4c, 0x1a, 0x03, 0xe7,
	0x8f, 0x01, 0x75, 0x37, 0xce, 0x08, 0x1d, 0x68, 0x8d, 0x42, 0x96, 0x73, 0x88, 0xff, 0x65, 0xf9,
	0xa5, 0xa5, 0x4e, 0x33, 0x57, 0x22, 0x15, 0xfc, 0x0c, 0xed, 0xb2, 0x70, 0xf0, 0x59, 0x46, 0x7b,
	0x20, 0xbb, 0xce, 0xdd, 0x3d, 0xc7, 0x51, 0x57, 0x70, 0x06, 0xe7, 0x25, 0x29, 0xe1, 0xd3, 0x23,
	0x72, 0xf9, 0x0c, 0x4b, 0xd4, 0x2e, 0xc0, 0x9a, 0xf0, 0xad, 0x97, 0xb2, 0x87, 0x7b, 0xd7, 0xc3,
	0xfc, 0x00, 0x9d, 0xe2, 0x80, 0xa4, 0x82, 0x13, 0x78, 0x74, 0x78, 0x2a, 0xff, 0x79, 0x97, 0x3e,
	0xa9, 0x72, 0x9d, 0x37, 0xd0, 0xba, 0x64, 0x3e, 0xcb, 0xe5, 0x79, 0xd0, 0xbb, 0x68, 0xf7, 0x05,
	0x40, 0x7c, 0x01, 0xea, 0x63, 0x3a, 0xcb, 0xf6, 0x25, 0xbc, 0x1f, 0xfc, 0x2b, 0x38, 0x95, 0x26,
	0xe5, 0x56, 0x71, 0xac, 0xc3, 0xa3, 0xa4, 0x82, 0x0e, 0x34, 0xe4, 0x13, 0x51, 0xba, 0xe7, 0xd9,
	0xe6, 0xdd, 0xab, 0x29, 0x9c, 0x78, 0xab, 0xe1, 0x6b, 0x80, 0x38, 0xec, 0x32, 0x37, 0x45, 0xf3,
	0x2f, 0x01, 0xe2, 0x69, 0x15, 0xfd, 0xe1, 0x49, 0xbf, 0xd6, 0xd4, 0x4f, 0xf4, 0xdd, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x48, 0xf7, 0x36, 0x81, 0x55, 0x05, 0x00, 0x00,
}
