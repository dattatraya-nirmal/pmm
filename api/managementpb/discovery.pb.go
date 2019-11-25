// Code generated by protoc-gen-go. DO NOT EDIT.
// source: managementpb/discovery.proto

package managementpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// DiscoverRDSEngine describes supported RDS instance engines.
type DiscoverRDSEngine int32

const (
	DiscoverRDSEngine_DISCOVER_RDS_ENGINE_INVALID DiscoverRDSEngine = 0
	DiscoverRDSEngine_DISCOVER_RDS_MYSQL          DiscoverRDSEngine = 1
)

var DiscoverRDSEngine_name = map[int32]string{
	0: "DISCOVER_RDS_ENGINE_INVALID",
	1: "DISCOVER_RDS_MYSQL",
}

var DiscoverRDSEngine_value = map[string]int32{
	"DISCOVER_RDS_ENGINE_INVALID": 0,
	"DISCOVER_RDS_MYSQL":          1,
}

func (x DiscoverRDSEngine) String() string {
	return proto.EnumName(DiscoverRDSEngine_name, int32(x))
}

func (DiscoverRDSEngine) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a52e3d9a4900498f, []int{0}
}

type DiscoverRDSRequest struct {
	// AWS Access key. Optional.
	AwsAccessKey string `protobuf:"bytes,1,opt,name=aws_access_key,json=awsAccessKey,proto3" json:"aws_access_key,omitempty"`
	// AWS Secret key. Optional.
	AwsSecretKey         string   `protobuf:"bytes,2,opt,name=aws_secret_key,json=awsSecretKey,proto3" json:"aws_secret_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscoverRDSRequest) Reset()         { *m = DiscoverRDSRequest{} }
func (m *DiscoverRDSRequest) String() string { return proto.CompactTextString(m) }
func (*DiscoverRDSRequest) ProtoMessage()    {}
func (*DiscoverRDSRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a52e3d9a4900498f, []int{0}
}

func (m *DiscoverRDSRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoverRDSRequest.Unmarshal(m, b)
}
func (m *DiscoverRDSRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoverRDSRequest.Marshal(b, m, deterministic)
}
func (m *DiscoverRDSRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoverRDSRequest.Merge(m, src)
}
func (m *DiscoverRDSRequest) XXX_Size() int {
	return xxx_messageInfo_DiscoverRDSRequest.Size(m)
}
func (m *DiscoverRDSRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoverRDSRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoverRDSRequest proto.InternalMessageInfo

func (m *DiscoverRDSRequest) GetAwsAccessKey() string {
	if m != nil {
		return m.AwsAccessKey
	}
	return ""
}

func (m *DiscoverRDSRequest) GetAwsSecretKey() string {
	if m != nil {
		return m.AwsSecretKey
	}
	return ""
}

// DiscoverRDSInstance models an unique RDS instance for the list of instances returned by Discovery.
type DiscoverRDSInstance struct {
	// AWS region.
	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	// AWS instance ID.
	InstanceId string `protobuf:"bytes,2,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	// Address used to connect to it.
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	// Instance engine.
	Engine DiscoverRDSEngine `protobuf:"varint,4,opt,name=engine,proto3,enum=management.DiscoverRDSEngine" json:"engine,omitempty"`
	// Engine version.
	EngineVersion        string   `protobuf:"bytes,5,opt,name=engine_version,json=engineVersion,proto3" json:"engine_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscoverRDSInstance) Reset()         { *m = DiscoverRDSInstance{} }
func (m *DiscoverRDSInstance) String() string { return proto.CompactTextString(m) }
func (*DiscoverRDSInstance) ProtoMessage()    {}
func (*DiscoverRDSInstance) Descriptor() ([]byte, []int) {
	return fileDescriptor_a52e3d9a4900498f, []int{1}
}

func (m *DiscoverRDSInstance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoverRDSInstance.Unmarshal(m, b)
}
func (m *DiscoverRDSInstance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoverRDSInstance.Marshal(b, m, deterministic)
}
func (m *DiscoverRDSInstance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoverRDSInstance.Merge(m, src)
}
func (m *DiscoverRDSInstance) XXX_Size() int {
	return xxx_messageInfo_DiscoverRDSInstance.Size(m)
}
func (m *DiscoverRDSInstance) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoverRDSInstance.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoverRDSInstance proto.InternalMessageInfo

func (m *DiscoverRDSInstance) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *DiscoverRDSInstance) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *DiscoverRDSInstance) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *DiscoverRDSInstance) GetEngine() DiscoverRDSEngine {
	if m != nil {
		return m.Engine
	}
	return DiscoverRDSEngine_DISCOVER_RDS_ENGINE_INVALID
}

func (m *DiscoverRDSInstance) GetEngineVersion() string {
	if m != nil {
		return m.EngineVersion
	}
	return ""
}

type DiscoverRDSResponse struct {
	RdsInstances         []*DiscoverRDSInstance `protobuf:"bytes,1,rep,name=rds_instances,json=rdsInstances,proto3" json:"rds_instances,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *DiscoverRDSResponse) Reset()         { *m = DiscoverRDSResponse{} }
func (m *DiscoverRDSResponse) String() string { return proto.CompactTextString(m) }
func (*DiscoverRDSResponse) ProtoMessage()    {}
func (*DiscoverRDSResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a52e3d9a4900498f, []int{2}
}

func (m *DiscoverRDSResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoverRDSResponse.Unmarshal(m, b)
}
func (m *DiscoverRDSResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoverRDSResponse.Marshal(b, m, deterministic)
}
func (m *DiscoverRDSResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoverRDSResponse.Merge(m, src)
}
func (m *DiscoverRDSResponse) XXX_Size() int {
	return xxx_messageInfo_DiscoverRDSResponse.Size(m)
}
func (m *DiscoverRDSResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoverRDSResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoverRDSResponse proto.InternalMessageInfo

func (m *DiscoverRDSResponse) GetRdsInstances() []*DiscoverRDSInstance {
	if m != nil {
		return m.RdsInstances
	}
	return nil
}

func init() {
	proto.RegisterEnum("management.DiscoverRDSEngine", DiscoverRDSEngine_name, DiscoverRDSEngine_value)
	proto.RegisterType((*DiscoverRDSRequest)(nil), "management.DiscoverRDSRequest")
	proto.RegisterType((*DiscoverRDSInstance)(nil), "management.DiscoverRDSInstance")
	proto.RegisterType((*DiscoverRDSResponse)(nil), "management.DiscoverRDSResponse")
}

func init() { proto.RegisterFile("managementpb/discovery.proto", fileDescriptor_a52e3d9a4900498f) }

var fileDescriptor_a52e3d9a4900498f = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x5f, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x71, 0x0b, 0xa9, 0x3a, 0x49, 0xa3, 0xb2, 0x48, 0x55, 0x14, 0x02, 0x09, 0x01, 0xd4,
	0x50, 0x68, 0x2c, 0x82, 0x78, 0xe1, 0xcd, 0x60, 0x0b, 0x59, 0x0d, 0x41, 0xac, 0xa5, 0x48, 0xc0,
	0x83, 0xd9, 0xda, 0x53, 0xcb, 0xa2, 0xdd, 0x35, 0x3b, 0x6e, 0x22, 0xbf, 0x22, 0x4e, 0x00, 0x77,
	0xe2, 0x02, 0x5c, 0x81, 0x83, 0x54, 0xf5, 0x9f, 0xc6, 0x55, 0x95, 0x27, 0x7b, 0x66, 0x7e, 0xfa,
	0xbe, 0x6f, 0xec, 0x5d, 0xe8, 0x9d, 0x09, 0x29, 0x22, 0x3c, 0x43, 0x99, 0x26, 0xc7, 0x66, 0x18,
	0x53, 0xa0, 0x16, 0xa8, 0xb3, 0x71, 0xa2, 0x55, 0xaa, 0x18, 0xac, 0xa6, 0xdd, 0x5e, 0xa4, 0x54,
	0x74, 0x8a, 0xa6, 0x48, 0x62, 0x53, 0x48, 0xa9, 0x52, 0x91, 0xc6, 0x4a, 0x52, 0x41, 0x76, 0x5f,
	0xe4, 0x8f, 0xe0, 0x30, 0x42, 0x79, 0x48, 0x4b, 0x11, 0x45, 0xa8, 0x4d, 0x95, 0xe4, 0xc4, 0x4d,
	0x7a, 0xf8, 0x0d, 0x98, 0x5d, 0x5a, 0x71, 0xdb, 0xe3, 0xf8, 0xe3, 0x1c, 0x29, 0x65, 0x4f, 0xa0,
	0x2d, 0x96, 0xe4, 0x8b, 0x20, 0x40, 0x22, 0xff, 0x3b, 0x66, 0x1d, 0x63, 0x60, 0x8c, 0xb6, 0x79,
	0x4b, 0x2c, 0xc9, 0xca, 0x9b, 0x47, 0x98, 0x55, 0x14, 0x61, 0xa0, 0x31, 0xcd, 0xa9, 0x8d, 0x2b,
	0xca, 0xcb, 0x9b, 0x47, 0x98, 0x0d, 0xff, 0x1a, 0x70, 0xaf, 0x66, 0xe1, 0x4a, 0x4a, 0x85, 0x0c,
	0x90, 0xed, 0x41, 0x43, 0x63, 0x14, 0x2b, 0x59, 0x6a, 0x97, 0x15, 0xeb, 0x43, 0x33, 0x2e, 0x19,
	0x3f, 0x0e, 0x4b, 0x49, 0xa8, 0x5a, 0x6e, 0xc8, 0x3a, 0xb0, 0x25, 0xc2, 0x50, 0x23, 0x51, 0x67,
	0x33, 0x1f, 0x56, 0x25, 0x7b, 0x0d, 0x0d, 0x94, 0x51, 0x2c, 0xb1, 0x73, 0x7b, 0x60, 0x8c, 0xda,
	0x93, 0x07, 0xe3, 0xd5, 0x57, 0x1b, 0xd7, 0x32, 0x38, 0x39, 0xc4, 0x4b, 0x98, 0x3d, 0x85, 0x76,
	0xf1, 0xe6, 0x2f, 0x50, 0xd3, 0x65, 0xa2, 0x3b, 0xb9, 0xee, 0x4e, 0xd1, 0x9d, 0x17, 0xcd, 0xe1,
	0xd7, 0x6b, 0x7b, 0x70, 0xa4, 0x44, 0x49, 0x42, 0x66, 0xc3, 0x8e, 0x0e, 0xc9, 0xaf, 0x02, 0x52,
	0xc7, 0x18, 0x6c, 0x8e, 0x9a, 0x93, 0xfe, 0x1a, 0xef, 0x6a, 0x7f, 0xde, 0xd2, 0x21, 0x55, 0x05,
	0x1d, 0x4c, 0xe1, 0xee, 0x8d, 0x80, 0xac, 0x0f, 0xf7, 0x6d, 0xd7, 0x7b, 0xf7, 0x71, 0xee, 0x70,
	0x9f, 0xdb, 0x9e, 0xef, 0xcc, 0xde, 0xbb, 0x33, 0xc7, 0x77, 0x67, 0x73, 0x6b, 0xea, 0xda, 0xbb,
	0xb7, 0xd8, 0x1e, 0xb0, 0x6b, 0xc0, 0x87, 0xcf, 0xde, 0xa7, 0xe9, 0xae, 0x31, 0xf9, 0x65, 0xc0,
	0x76, 0x25, 0x97, 0xb1, 0x25, 0x34, 0x6b, 0xda, 0xec, 0xe1, 0x9a, 0x64, 0xe5, 0xcf, 0xef, 0xf6,
	0xd7, 0xce, 0x8b, 0x8d, 0x87, 0xfb, 0x3f, 0xff, 0xfd, 0xff, 0xb3, 0xf1, 0x68, 0xd8, 0x33, 0x17,
	0x2f, 0xcd, 0x15, 0x6b, 0x5e, 0x39, 0x9a, 0xdc, 0xf6, 0xde, 0x18, 0x07, 0x6f, 0xfd, 0xdf, 0xd6,
	0x8c, 0x4f, 0x61, 0x2b, 0xc4, 0x13, 0x71, 0x7e, 0x9a, 0x32, 0x0b, 0x98, 0x25, 0x07, 0xa8, 0xb5,
	0xd2, 0x03, 0x5d, 0xaa, 0x8d, 0xd9, 0x73, 0x78, 0xd6, 0xdd, 0x7f, 0x6c, 0x86, 0x78, 0x12, 0xcb,
	0xb8, 0x38, 0xa7, 0xf5, 0xab, 0xe0, 0x5c, 0xe2, 0x95, 0xf7, 0x97, 0x56, 0x7d, 0x74, 0xdc, 0xc8,
	0x0f, 0xf1, 0xab, 0x8b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x42, 0x21, 0x5f, 0x3c, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DiscoveryClient is the client API for Discovery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DiscoveryClient interface {
	// DiscoverRDS discovers RDS instances.
	DiscoverRDS(ctx context.Context, in *DiscoverRDSRequest, opts ...grpc.CallOption) (*DiscoverRDSResponse, error)
}

type discoveryClient struct {
	cc *grpc.ClientConn
}

func NewDiscoveryClient(cc *grpc.ClientConn) DiscoveryClient {
	return &discoveryClient{cc}
}

func (c *discoveryClient) DiscoverRDS(ctx context.Context, in *DiscoverRDSRequest, opts ...grpc.CallOption) (*DiscoverRDSResponse, error) {
	out := new(DiscoverRDSResponse)
	err := c.cc.Invoke(ctx, "/management.Discovery/DiscoverRDS", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscoveryServer is the server API for Discovery service.
type DiscoveryServer interface {
	// DiscoverRDS discovers RDS instances.
	DiscoverRDS(context.Context, *DiscoverRDSRequest) (*DiscoverRDSResponse, error)
}

// UnimplementedDiscoveryServer can be embedded to have forward compatible implementations.
type UnimplementedDiscoveryServer struct {
}

func (*UnimplementedDiscoveryServer) DiscoverRDS(ctx context.Context, req *DiscoverRDSRequest) (*DiscoverRDSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiscoverRDS not implemented")
}

func RegisterDiscoveryServer(s *grpc.Server, srv DiscoveryServer) {
	s.RegisterService(&_Discovery_serviceDesc, srv)
}

func _Discovery_DiscoverRDS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoverRDSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).DiscoverRDS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Discovery/DiscoverRDS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).DiscoverRDS(ctx, req.(*DiscoverRDSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Discovery_serviceDesc = grpc.ServiceDesc{
	ServiceName: "management.Discovery",
	HandlerType: (*DiscoveryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DiscoverRDS",
			Handler:    _Discovery_DiscoverRDS_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/discovery.proto",
}