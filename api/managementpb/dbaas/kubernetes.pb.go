// Code generated by protoc-gen-go. DO NOT EDIT.
// source: managementpb/dbaas/kubernetes.proto

package dbaasv1beta1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/mwitkow/go-proto-validators"
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

// KubeAuth represents Kubernetes / kubectl authentication and authorization information.
type KubeAuth struct {
	// Kubeconfig file content.
	Kubeconfig           string   `protobuf:"bytes,1,opt,name=kubeconfig,proto3" json:"kubeconfig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KubeAuth) Reset()         { *m = KubeAuth{} }
func (m *KubeAuth) String() string { return proto.CompactTextString(m) }
func (*KubeAuth) ProtoMessage()    {}
func (*KubeAuth) Descriptor() ([]byte, []int) {
	return fileDescriptor_1118852c6b00fddc, []int{0}
}

func (m *KubeAuth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KubeAuth.Unmarshal(m, b)
}
func (m *KubeAuth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KubeAuth.Marshal(b, m, deterministic)
}
func (m *KubeAuth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubeAuth.Merge(m, src)
}
func (m *KubeAuth) XXX_Size() int {
	return xxx_messageInfo_KubeAuth.Size(m)
}
func (m *KubeAuth) XXX_DiscardUnknown() {
	xxx_messageInfo_KubeAuth.DiscardUnknown(m)
}

var xxx_messageInfo_KubeAuth proto.InternalMessageInfo

func (m *KubeAuth) GetKubeconfig() string {
	if m != nil {
		return m.Kubeconfig
	}
	return ""
}

type ListKubernetesClustersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListKubernetesClustersRequest) Reset()         { *m = ListKubernetesClustersRequest{} }
func (m *ListKubernetesClustersRequest) String() string { return proto.CompactTextString(m) }
func (*ListKubernetesClustersRequest) ProtoMessage()    {}
func (*ListKubernetesClustersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1118852c6b00fddc, []int{1}
}

func (m *ListKubernetesClustersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListKubernetesClustersRequest.Unmarshal(m, b)
}
func (m *ListKubernetesClustersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListKubernetesClustersRequest.Marshal(b, m, deterministic)
}
func (m *ListKubernetesClustersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListKubernetesClustersRequest.Merge(m, src)
}
func (m *ListKubernetesClustersRequest) XXX_Size() int {
	return xxx_messageInfo_ListKubernetesClustersRequest.Size(m)
}
func (m *ListKubernetesClustersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListKubernetesClustersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListKubernetesClustersRequest proto.InternalMessageInfo

type ListKubernetesClustersResponse struct {
	// Kubernetes cluster name.
	KubernetesClusterName string   `protobuf:"bytes,1,opt,name=kubernetes_cluster_name,json=kubernetesClusterName,proto3" json:"kubernetes_cluster_name,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *ListKubernetesClustersResponse) Reset()         { *m = ListKubernetesClustersResponse{} }
func (m *ListKubernetesClustersResponse) String() string { return proto.CompactTextString(m) }
func (*ListKubernetesClustersResponse) ProtoMessage()    {}
func (*ListKubernetesClustersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1118852c6b00fddc, []int{2}
}

func (m *ListKubernetesClustersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListKubernetesClustersResponse.Unmarshal(m, b)
}
func (m *ListKubernetesClustersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListKubernetesClustersResponse.Marshal(b, m, deterministic)
}
func (m *ListKubernetesClustersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListKubernetesClustersResponse.Merge(m, src)
}
func (m *ListKubernetesClustersResponse) XXX_Size() int {
	return xxx_messageInfo_ListKubernetesClustersResponse.Size(m)
}
func (m *ListKubernetesClustersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListKubernetesClustersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListKubernetesClustersResponse proto.InternalMessageInfo

func (m *ListKubernetesClustersResponse) GetKubernetesClusterName() string {
	if m != nil {
		return m.KubernetesClusterName
	}
	return ""
}

type RegisterKubernetesClusterRequest struct {
	// Kubernetes cluster name.
	KubernetesClusterName string `protobuf:"bytes,1,opt,name=kubernetes_cluster_name,json=kubernetesClusterName,proto3" json:"kubernetes_cluster_name,omitempty"`
	// Kubernetes auth.
	KubeAuth             *KubeAuth `protobuf:"bytes,2,opt,name=kube_auth,json=kubeAuth,proto3" json:"kube_auth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RegisterKubernetesClusterRequest) Reset()         { *m = RegisterKubernetesClusterRequest{} }
func (m *RegisterKubernetesClusterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterKubernetesClusterRequest) ProtoMessage()    {}
func (*RegisterKubernetesClusterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1118852c6b00fddc, []int{3}
}

func (m *RegisterKubernetesClusterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterKubernetesClusterRequest.Unmarshal(m, b)
}
func (m *RegisterKubernetesClusterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterKubernetesClusterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterKubernetesClusterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterKubernetesClusterRequest.Merge(m, src)
}
func (m *RegisterKubernetesClusterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterKubernetesClusterRequest.Size(m)
}
func (m *RegisterKubernetesClusterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterKubernetesClusterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterKubernetesClusterRequest proto.InternalMessageInfo

func (m *RegisterKubernetesClusterRequest) GetKubernetesClusterName() string {
	if m != nil {
		return m.KubernetesClusterName
	}
	return ""
}

func (m *RegisterKubernetesClusterRequest) GetKubeAuth() *KubeAuth {
	if m != nil {
		return m.KubeAuth
	}
	return nil
}

type RegisterKubernetesClusterResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterKubernetesClusterResponse) Reset()         { *m = RegisterKubernetesClusterResponse{} }
func (m *RegisterKubernetesClusterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterKubernetesClusterResponse) ProtoMessage()    {}
func (*RegisterKubernetesClusterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1118852c6b00fddc, []int{4}
}

func (m *RegisterKubernetesClusterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterKubernetesClusterResponse.Unmarshal(m, b)
}
func (m *RegisterKubernetesClusterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterKubernetesClusterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterKubernetesClusterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterKubernetesClusterResponse.Merge(m, src)
}
func (m *RegisterKubernetesClusterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterKubernetesClusterResponse.Size(m)
}
func (m *RegisterKubernetesClusterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterKubernetesClusterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterKubernetesClusterResponse proto.InternalMessageInfo

type UnregisterKubernetesClusterRequest struct {
	// Kubernetes cluster name.
	KubernetesClusterName string   `protobuf:"bytes,1,opt,name=kubernetes_cluster_name,json=kubernetesClusterName,proto3" json:"kubernetes_cluster_name,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *UnregisterKubernetesClusterRequest) Reset()         { *m = UnregisterKubernetesClusterRequest{} }
func (m *UnregisterKubernetesClusterRequest) String() string { return proto.CompactTextString(m) }
func (*UnregisterKubernetesClusterRequest) ProtoMessage()    {}
func (*UnregisterKubernetesClusterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1118852c6b00fddc, []int{5}
}

func (m *UnregisterKubernetesClusterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnregisterKubernetesClusterRequest.Unmarshal(m, b)
}
func (m *UnregisterKubernetesClusterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnregisterKubernetesClusterRequest.Marshal(b, m, deterministic)
}
func (m *UnregisterKubernetesClusterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnregisterKubernetesClusterRequest.Merge(m, src)
}
func (m *UnregisterKubernetesClusterRequest) XXX_Size() int {
	return xxx_messageInfo_UnregisterKubernetesClusterRequest.Size(m)
}
func (m *UnregisterKubernetesClusterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnregisterKubernetesClusterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnregisterKubernetesClusterRequest proto.InternalMessageInfo

func (m *UnregisterKubernetesClusterRequest) GetKubernetesClusterName() string {
	if m != nil {
		return m.KubernetesClusterName
	}
	return ""
}

type UnregisterKubernetesClusterResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnregisterKubernetesClusterResponse) Reset()         { *m = UnregisterKubernetesClusterResponse{} }
func (m *UnregisterKubernetesClusterResponse) String() string { return proto.CompactTextString(m) }
func (*UnregisterKubernetesClusterResponse) ProtoMessage()    {}
func (*UnregisterKubernetesClusterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1118852c6b00fddc, []int{6}
}

func (m *UnregisterKubernetesClusterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnregisterKubernetesClusterResponse.Unmarshal(m, b)
}
func (m *UnregisterKubernetesClusterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnregisterKubernetesClusterResponse.Marshal(b, m, deterministic)
}
func (m *UnregisterKubernetesClusterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnregisterKubernetesClusterResponse.Merge(m, src)
}
func (m *UnregisterKubernetesClusterResponse) XXX_Size() int {
	return xxx_messageInfo_UnregisterKubernetesClusterResponse.Size(m)
}
func (m *UnregisterKubernetesClusterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UnregisterKubernetesClusterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UnregisterKubernetesClusterResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*KubeAuth)(nil), "dbaas.v1beta1.KubeAuth")
	proto.RegisterType((*ListKubernetesClustersRequest)(nil), "dbaas.v1beta1.ListKubernetesClustersRequest")
	proto.RegisterType((*ListKubernetesClustersResponse)(nil), "dbaas.v1beta1.ListKubernetesClustersResponse")
	proto.RegisterType((*RegisterKubernetesClusterRequest)(nil), "dbaas.v1beta1.RegisterKubernetesClusterRequest")
	proto.RegisterType((*RegisterKubernetesClusterResponse)(nil), "dbaas.v1beta1.RegisterKubernetesClusterResponse")
	proto.RegisterType((*UnregisterKubernetesClusterRequest)(nil), "dbaas.v1beta1.UnregisterKubernetesClusterRequest")
	proto.RegisterType((*UnregisterKubernetesClusterResponse)(nil), "dbaas.v1beta1.UnregisterKubernetesClusterResponse")
}

func init() {
	proto.RegisterFile("managementpb/dbaas/kubernetes.proto", fileDescriptor_1118852c6b00fddc)
}

var fileDescriptor_1118852c6b00fddc = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xb5, 0x11, 0xaa, 0xda, 0xe1, 0xe3, 0xb0, 0x12, 0xb4, 0x98, 0x8f, 0x86, 0x0d, 0xa0,
	0x10, 0xd5, 0x5e, 0x92, 0x8a, 0x1e, 0x2a, 0x84, 0x44, 0xe0, 0x56, 0xc4, 0xc1, 0x08, 0xa9, 0xe2,
	0x12, 0xad, 0x93, 0x61, 0x63, 0x25, 0xde, 0x35, 0xde, 0x75, 0x72, 0xe7, 0x15, 0x78, 0x01, 0x4e,
	0x48, 0x9c, 0x39, 0xf2, 0x18, 0x3c, 0x00, 0x12, 0xe2, 0x41, 0x90, 0x3f, 0x1a, 0x43, 0xda, 0xc6,
	0xe1, 0xd0, 0x93, 0xad, 0x9d, 0x99, 0xff, 0xff, 0x37, 0xb3, 0x63, 0x43, 0x2b, 0x12, 0x4a, 0x48,
	0x8c, 0x50, 0xd9, 0x38, 0xe0, 0xa3, 0x40, 0x08, 0xc3, 0x27, 0x69, 0x80, 0x89, 0x42, 0x8b, 0xc6,
	0x8b, 0x13, 0x6d, 0x35, 0xbd, 0x9a, 0x9f, 0x7b, 0xb3, 0x6e, 0x80, 0x56, 0x74, 0x9d, 0x03, 0x19,
	0xda, 0x71, 0x1a, 0x78, 0x43, 0x1d, 0xf1, 0x68, 0x1e, 0xda, 0x89, 0x9e, 0x73, 0xa9, 0xdd, 0x3c,
	0xd7, 0x9d, 0x89, 0x69, 0x38, 0x12, 0x56, 0x27, 0x86, 0x2f, 0x5e, 0x0b, 0x19, 0xe7, 0xb6, 0xd4,
	0x5a, 0x4e, 0x91, 0x8b, 0x38, 0xe4, 0x42, 0x29, 0x6d, 0x85, 0x0d, 0xb5, 0x2a, 0x4d, 0x9c, 0xbd,
	0xfc, 0x31, 0x74, 0x25, 0x2a, 0xd7, 0xcc, 0x85, 0x94, 0x98, 0x70, 0x1d, 0xe7, 0x19, 0xa7, 0xb3,
	0x59, 0x0f, 0x36, 0x8f, 0xd2, 0x00, 0x9f, 0xa7, 0x76, 0x4c, 0x1f, 0x02, 0x64, 0xc8, 0x43, 0xad,
	0xde, 0x87, 0x72, 0x87, 0x34, 0x49, 0x7b, 0xab, 0xbf, 0xf1, 0xeb, 0xe7, 0x6e, 0xe3, 0x98, 0xf8,
	0x7f, 0x45, 0xd8, 0x2e, 0xdc, 0x79, 0x15, 0x1a, 0x7b, 0xb4, 0x68, 0xef, 0xc5, 0x34, 0x35, 0x16,
	0x13, 0xe3, 0xe3, 0x87, 0x14, 0x8d, 0x65, 0xc7, 0x70, 0xf7, 0xbc, 0x04, 0x13, 0x6b, 0x65, 0x90,
	0x1e, 0xc0, 0x76, 0x35, 0x9d, 0xc1, 0xb0, 0x08, 0x0f, 0x94, 0x88, 0xb0, 0xf0, 0xf5, 0xaf, 0x4f,
	0x96, 0x8b, 0x5f, 0x8b, 0x08, 0xd9, 0x67, 0x02, 0x4d, 0x1f, 0x65, 0x98, 0x1d, 0x9c, 0x92, 0x2f,
	0xed, 0xe9, 0xb3, 0x1a, 0xf1, 0x45, 0x53, 0x67, 0x9b, 0xd0, 0xa7, 0xb0, 0x95, 0x05, 0x06, 0x22,
	0xb5, 0xe3, 0x9d, 0x46, 0x93, 0xb4, 0x2f, 0xf7, 0xb6, 0xbd, 0x7f, 0xae, 0xce, 0x3b, 0x99, 0x59,
	0x21, 0xd5, 0x24, 0xfe, 0xe6, 0xa4, 0x3c, 0x61, 0x2d, 0xb8, 0xb7, 0x82, 0xb0, 0xe8, 0x9f, 0x8d,
	0x80, 0xbd, 0x55, 0xc9, 0x05, 0x37, 0xc2, 0x1e, 0x40, 0x6b, 0xa5, 0x4b, 0x01, 0xd3, 0xfb, 0x7a,
	0x09, 0xa0, 0x8a, 0xd2, 0x2f, 0x04, 0x6e, 0x9c, 0x7d, 0x7d, 0x74, 0x6f, 0x69, 0x0c, 0x2b, 0xd7,
	0xc0, 0x71, 0xd7, 0xcc, 0x2e, 0x67, 0xc2, 0x3f, 0xfe, 0xf8, 0xfd, 0xa9, 0xf1, 0x88, 0xdd, 0xe7,
	0xb3, 0x2e, 0xaf, 0x3e, 0x27, 0xfe, 0xb2, 0x2f, 0xc4, 0x1b, 0x5e, 0x15, 0xf2, 0x4c, 0xe7, 0x90,
	0x74, 0xe8, 0x37, 0x02, 0x37, 0xcf, 0x1d, 0x35, 0xe5, 0x4b, 0xee, 0x75, 0x6b, 0xe3, 0x3c, 0x5e,
	0xbf, 0xa0, 0x24, 0xde, 0xcf, 0x89, 0x5d, 0xd6, 0xae, 0x23, 0x3e, 0x91, 0xca, 0xa8, 0xbf, 0x13,
	0xb8, 0xb5, 0xe2, 0x56, 0x68, 0x77, 0x09, 0xa3, 0x7e, 0x4f, 0x9c, 0xde, 0xff, 0x94, 0x94, 0xec,
	0x4f, 0x72, 0x76, 0xce, 0x3a, 0x75, 0xec, 0x95, 0xd8, 0x21, 0xe9, 0xf4, 0xaf, 0xbd, 0xbb, 0x92,
	0x7b, 0x95, 0x56, 0xc1, 0x46, 0xfe, 0x1b, 0xd9, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x96,
	0x47, 0x54, 0x00, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KubernetesClient is the client API for Kubernetes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KubernetesClient interface {
	// ListKubernetesClusters returns a list of all registered Kubernetes clusters.
	ListKubernetesClusters(ctx context.Context, in *ListKubernetesClustersRequest, opts ...grpc.CallOption) (*ListKubernetesClustersResponse, error)
	// RegisterKubernetesCluster registers an existing Kubernetes cluster in PMM.
	RegisterKubernetesCluster(ctx context.Context, in *RegisterKubernetesClusterRequest, opts ...grpc.CallOption) (*RegisterKubernetesClusterResponse, error)
	// UnregisterKubernetesCluster removes a registered Kubernetes cluster from PMM.
	UnregisterKubernetesCluster(ctx context.Context, in *UnregisterKubernetesClusterRequest, opts ...grpc.CallOption) (*UnregisterKubernetesClusterResponse, error)
}

type kubernetesClient struct {
	cc grpc.ClientConnInterface
}

func NewKubernetesClient(cc grpc.ClientConnInterface) KubernetesClient {
	return &kubernetesClient{cc}
}

func (c *kubernetesClient) ListKubernetesClusters(ctx context.Context, in *ListKubernetesClustersRequest, opts ...grpc.CallOption) (*ListKubernetesClustersResponse, error) {
	out := new(ListKubernetesClustersResponse)
	err := c.cc.Invoke(ctx, "/dbaas.v1beta1.Kubernetes/ListKubernetesClusters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClient) RegisterKubernetesCluster(ctx context.Context, in *RegisterKubernetesClusterRequest, opts ...grpc.CallOption) (*RegisterKubernetesClusterResponse, error) {
	out := new(RegisterKubernetesClusterResponse)
	err := c.cc.Invoke(ctx, "/dbaas.v1beta1.Kubernetes/RegisterKubernetesCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClient) UnregisterKubernetesCluster(ctx context.Context, in *UnregisterKubernetesClusterRequest, opts ...grpc.CallOption) (*UnregisterKubernetesClusterResponse, error) {
	out := new(UnregisterKubernetesClusterResponse)
	err := c.cc.Invoke(ctx, "/dbaas.v1beta1.Kubernetes/UnregisterKubernetesCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KubernetesServer is the server API for Kubernetes service.
type KubernetesServer interface {
	// ListKubernetesClusters returns a list of all registered Kubernetes clusters.
	ListKubernetesClusters(context.Context, *ListKubernetesClustersRequest) (*ListKubernetesClustersResponse, error)
	// RegisterKubernetesCluster registers an existing Kubernetes cluster in PMM.
	RegisterKubernetesCluster(context.Context, *RegisterKubernetesClusterRequest) (*RegisterKubernetesClusterResponse, error)
	// UnregisterKubernetesCluster removes a registered Kubernetes cluster from PMM.
	UnregisterKubernetesCluster(context.Context, *UnregisterKubernetesClusterRequest) (*UnregisterKubernetesClusterResponse, error)
}

// UnimplementedKubernetesServer can be embedded to have forward compatible implementations.
type UnimplementedKubernetesServer struct {
}

func (*UnimplementedKubernetesServer) ListKubernetesClusters(ctx context.Context, req *ListKubernetesClustersRequest) (*ListKubernetesClustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKubernetesClusters not implemented")
}
func (*UnimplementedKubernetesServer) RegisterKubernetesCluster(ctx context.Context, req *RegisterKubernetesClusterRequest) (*RegisterKubernetesClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterKubernetesCluster not implemented")
}
func (*UnimplementedKubernetesServer) UnregisterKubernetesCluster(ctx context.Context, req *UnregisterKubernetesClusterRequest) (*UnregisterKubernetesClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterKubernetesCluster not implemented")
}

func RegisterKubernetesServer(s *grpc.Server, srv KubernetesServer) {
	s.RegisterService(&_Kubernetes_serviceDesc, srv)
}

func _Kubernetes_ListKubernetesClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKubernetesClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesServer).ListKubernetesClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbaas.v1beta1.Kubernetes/ListKubernetesClusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesServer).ListKubernetesClusters(ctx, req.(*ListKubernetesClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubernetes_RegisterKubernetesCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterKubernetesClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesServer).RegisterKubernetesCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbaas.v1beta1.Kubernetes/RegisterKubernetesCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesServer).RegisterKubernetesCluster(ctx, req.(*RegisterKubernetesClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubernetes_UnregisterKubernetesCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnregisterKubernetesClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesServer).UnregisterKubernetesCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbaas.v1beta1.Kubernetes/UnregisterKubernetesCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesServer).UnregisterKubernetesCluster(ctx, req.(*UnregisterKubernetesClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Kubernetes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dbaas.v1beta1.Kubernetes",
	HandlerType: (*KubernetesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListKubernetesClusters",
			Handler:    _Kubernetes_ListKubernetesClusters_Handler,
		},
		{
			MethodName: "RegisterKubernetesCluster",
			Handler:    _Kubernetes_RegisterKubernetesCluster_Handler,
		},
		{
			MethodName: "UnregisterKubernetesCluster",
			Handler:    _Kubernetes_UnregisterKubernetesCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/dbaas/kubernetes.proto",
}
