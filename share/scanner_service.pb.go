// Code generated by protoc-gen-go. DO NOT EDIT.
// source: scanner_service.proto

package share

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ScanImageRequest struct {
	Registry           string                 `protobuf:"bytes,1,opt,name=Registry" json:"Registry,omitempty"`
	Username           string                 `protobuf:"bytes,2,opt,name=Username" json:"Username,omitempty"`
	Password           string                 `protobuf:"bytes,3,opt,name=Password" json:"Password,omitempty"`
	Repository         string                 `protobuf:"bytes,4,opt,name=Repository" json:"Repository,omitempty"`
	Tag                string                 `protobuf:"bytes,5,opt,name=Tag" json:"Tag,omitempty"`
	Proxy              string                 `protobuf:"bytes,6,opt,name=Proxy" json:"Proxy,omitempty"`
	ScanLayers         bool                   `protobuf:"varint,7,opt,name=ScanLayers" json:"ScanLayers,omitempty"`
	ScanSecrets        bool                   `protobuf:"varint,8,opt,name=ScanSecrets" json:"ScanSecrets,omitempty"`
	BaseImage          string                 `protobuf:"bytes,9,opt,name=BaseImage" json:"BaseImage,omitempty"`
	RootsOfTrust       []*SigstoreRootOfTrust `protobuf:"bytes,10,rep,name=RootsOfTrust" json:"RootsOfTrust,omitempty"`
	Token              string                 `protobuf:"bytes,11,opt,name=Token" json:"Token,omitempty"`
	ScanTypesRequested *ScanTypeMap           `protobuf:"bytes,12,opt,name=ScanTypesRequested" json:"ScanTypesRequested,omitempty"`
}

func (m *ScanImageRequest) Reset()                    { *m = ScanImageRequest{} }
func (m *ScanImageRequest) String() string            { return proto.CompactTextString(m) }
func (*ScanImageRequest) ProtoMessage()               {}
func (*ScanImageRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *ScanImageRequest) GetRegistry() string {
	if m != nil {
		return m.Registry
	}
	return ""
}

func (m *ScanImageRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ScanImageRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *ScanImageRequest) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *ScanImageRequest) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *ScanImageRequest) GetProxy() string {
	if m != nil {
		return m.Proxy
	}
	return ""
}

func (m *ScanImageRequest) GetScanLayers() bool {
	if m != nil {
		return m.ScanLayers
	}
	return false
}

func (m *ScanImageRequest) GetScanSecrets() bool {
	if m != nil {
		return m.ScanSecrets
	}
	return false
}

func (m *ScanImageRequest) GetBaseImage() string {
	if m != nil {
		return m.BaseImage
	}
	return ""
}

func (m *ScanImageRequest) GetRootsOfTrust() []*SigstoreRootOfTrust {
	if m != nil {
		return m.RootsOfTrust
	}
	return nil
}

func (m *ScanImageRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ScanImageRequest) GetScanTypesRequested() *ScanTypeMap {
	if m != nil {
		return m.ScanTypesRequested
	}
	return nil
}

type SigstoreRootOfTrust struct {
	Name                 string              `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	RekorPublicKey       string              `protobuf:"bytes,2,opt,name=RekorPublicKey" json:"RekorPublicKey,omitempty"`
	RootCert             string              `protobuf:"bytes,3,opt,name=RootCert" json:"RootCert,omitempty"`
	SCTPublicKey         string              `protobuf:"bytes,4,opt,name=SCTPublicKey" json:"SCTPublicKey,omitempty"`
	Verifiers            []*SigstoreVerifier `protobuf:"bytes,5,rep,name=Verifiers" json:"Verifiers,omitempty"`
	RootlessKeypairsOnly bool                `protobuf:"varint,6,opt,name=RootlessKeypairsOnly" json:"RootlessKeypairsOnly,omitempty"`
}

func (m *SigstoreRootOfTrust) Reset()                    { *m = SigstoreRootOfTrust{} }
func (m *SigstoreRootOfTrust) String() string            { return proto.CompactTextString(m) }
func (*SigstoreRootOfTrust) ProtoMessage()               {}
func (*SigstoreRootOfTrust) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *SigstoreRootOfTrust) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SigstoreRootOfTrust) GetRekorPublicKey() string {
	if m != nil {
		return m.RekorPublicKey
	}
	return ""
}

func (m *SigstoreRootOfTrust) GetRootCert() string {
	if m != nil {
		return m.RootCert
	}
	return ""
}

func (m *SigstoreRootOfTrust) GetSCTPublicKey() string {
	if m != nil {
		return m.SCTPublicKey
	}
	return ""
}

func (m *SigstoreRootOfTrust) GetVerifiers() []*SigstoreVerifier {
	if m != nil {
		return m.Verifiers
	}
	return nil
}

func (m *SigstoreRootOfTrust) GetRootlessKeypairsOnly() bool {
	if m != nil {
		return m.RootlessKeypairsOnly
	}
	return false
}

type SigstoreVerifier struct {
	Name           string                  `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Type           string                  `protobuf:"bytes,2,opt,name=Type" json:"Type,omitempty"`
	KeypairOptions *SigstoreKeypairOptions `protobuf:"bytes,3,opt,name=KeypairOptions" json:"KeypairOptions,omitempty"`
	KeylessOptions *SigstoreKeylessOptions `protobuf:"bytes,4,opt,name=KeylessOptions" json:"KeylessOptions,omitempty"`
}

func (m *SigstoreVerifier) Reset()                    { *m = SigstoreVerifier{} }
func (m *SigstoreVerifier) String() string            { return proto.CompactTextString(m) }
func (*SigstoreVerifier) ProtoMessage()               {}
func (*SigstoreVerifier) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *SigstoreVerifier) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SigstoreVerifier) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *SigstoreVerifier) GetKeypairOptions() *SigstoreKeypairOptions {
	if m != nil {
		return m.KeypairOptions
	}
	return nil
}

func (m *SigstoreVerifier) GetKeylessOptions() *SigstoreKeylessOptions {
	if m != nil {
		return m.KeylessOptions
	}
	return nil
}

type SigstoreKeypairOptions struct {
	PublicKey string `protobuf:"bytes,1,opt,name=PublicKey" json:"PublicKey,omitempty"`
}

func (m *SigstoreKeypairOptions) Reset()                    { *m = SigstoreKeypairOptions{} }
func (m *SigstoreKeypairOptions) String() string            { return proto.CompactTextString(m) }
func (*SigstoreKeypairOptions) ProtoMessage()               {}
func (*SigstoreKeypairOptions) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *SigstoreKeypairOptions) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

type SigstoreKeylessOptions struct {
	CertIssuer  string `protobuf:"bytes,1,opt,name=CertIssuer" json:"CertIssuer,omitempty"`
	CertSubject string `protobuf:"bytes,2,opt,name=CertSubject" json:"CertSubject,omitempty"`
}

func (m *SigstoreKeylessOptions) Reset()                    { *m = SigstoreKeylessOptions{} }
func (m *SigstoreKeylessOptions) String() string            { return proto.CompactTextString(m) }
func (*SigstoreKeylessOptions) ProtoMessage()               {}
func (*SigstoreKeylessOptions) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *SigstoreKeylessOptions) GetCertIssuer() string {
	if m != nil {
		return m.CertIssuer
	}
	return ""
}

func (m *SigstoreKeylessOptions) GetCertSubject() string {
	if m != nil {
		return m.CertSubject
	}
	return ""
}

type ScanCacheStatRes struct {
	RecordCnt  uint64 `protobuf:"varint,1,opt,name=RecordCnt" json:"RecordCnt,omitempty"`
	RecordSize uint64 `protobuf:"varint,2,opt,name=RecordSize" json:"RecordSize,omitempty"`
	MissCnt    uint64 `protobuf:"varint,3,opt,name=MissCnt" json:"MissCnt,omitempty"`
	HitCnt     uint64 `protobuf:"varint,4,opt,name=HitCnt" json:"HitCnt,omitempty"`
}

func (m *ScanCacheStatRes) Reset()                    { *m = ScanCacheStatRes{} }
func (m *ScanCacheStatRes) String() string            { return proto.CompactTextString(m) }
func (*ScanCacheStatRes) ProtoMessage()               {}
func (*ScanCacheStatRes) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *ScanCacheStatRes) GetRecordCnt() uint64 {
	if m != nil {
		return m.RecordCnt
	}
	return 0
}

func (m *ScanCacheStatRes) GetRecordSize() uint64 {
	if m != nil {
		return m.RecordSize
	}
	return 0
}

func (m *ScanCacheStatRes) GetMissCnt() uint64 {
	if m != nil {
		return m.MissCnt
	}
	return 0
}

func (m *ScanCacheStatRes) GetHitCnt() uint64 {
	if m != nil {
		return m.HitCnt
	}
	return 0
}

type ScanCacheDataRes struct {
	DataZb []byte `protobuf:"bytes,1,opt,name=DataZb,proto3" json:"DataZb,omitempty"`
}

func (m *ScanCacheDataRes) Reset()                    { *m = ScanCacheDataRes{} }
func (m *ScanCacheDataRes) String() string            { return proto.CompactTextString(m) }
func (*ScanCacheDataRes) ProtoMessage()               {}
func (*ScanCacheDataRes) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

func (m *ScanCacheDataRes) GetDataZb() []byte {
	if m != nil {
		return m.DataZb
	}
	return nil
}

func init() {
	proto.RegisterType((*ScanImageRequest)(nil), "share.ScanImageRequest")
	proto.RegisterType((*SigstoreRootOfTrust)(nil), "share.SigstoreRootOfTrust")
	proto.RegisterType((*SigstoreVerifier)(nil), "share.SigstoreVerifier")
	proto.RegisterType((*SigstoreKeypairOptions)(nil), "share.SigstoreKeypairOptions")
	proto.RegisterType((*SigstoreKeylessOptions)(nil), "share.SigstoreKeylessOptions")
	proto.RegisterType((*ScanCacheStatRes)(nil), "share.ScanCacheStatRes")
	proto.RegisterType((*ScanCacheDataRes)(nil), "share.ScanCacheDataRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ScannerService service

type ScannerServiceClient interface {
	ScanRunning(ctx context.Context, in *ScanRunningRequest, opts ...grpc.CallOption) (*ScanResult, error)
	ScanImageData(ctx context.Context, in *ScanData, opts ...grpc.CallOption) (*ScanResult, error)
	ScanImage(ctx context.Context, in *ScanImageRequest, opts ...grpc.CallOption) (*ScanResult, error)
	ScanAppPackage(ctx context.Context, in *ScanAppRequest, opts ...grpc.CallOption) (*ScanResult, error)
	Ping(ctx context.Context, in *RPCVoid, opts ...grpc.CallOption) (*RPCVoid, error)
	ScanAwsLambda(ctx context.Context, in *ScanAwsLambdaRequest, opts ...grpc.CallOption) (*ScanResult, error)
	ScanCacheGetStat(ctx context.Context, in *RPCVoid, opts ...grpc.CallOption) (*ScanCacheStatRes, error)
	ScanCacheGetData(ctx context.Context, in *RPCVoid, opts ...grpc.CallOption) (*ScanCacheDataRes, error)
}

type scannerServiceClient struct {
	cc *grpc.ClientConn
}

func NewScannerServiceClient(cc *grpc.ClientConn) ScannerServiceClient {
	return &scannerServiceClient{cc}
}

func (c *scannerServiceClient) ScanRunning(ctx context.Context, in *ScanRunningRequest, opts ...grpc.CallOption) (*ScanResult, error) {
	out := new(ScanResult)
	err := grpc.Invoke(ctx, "/share.ScannerService/ScanRunning", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scannerServiceClient) ScanImageData(ctx context.Context, in *ScanData, opts ...grpc.CallOption) (*ScanResult, error) {
	out := new(ScanResult)
	err := grpc.Invoke(ctx, "/share.ScannerService/ScanImageData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scannerServiceClient) ScanImage(ctx context.Context, in *ScanImageRequest, opts ...grpc.CallOption) (*ScanResult, error) {
	out := new(ScanResult)
	err := grpc.Invoke(ctx, "/share.ScannerService/ScanImage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scannerServiceClient) ScanAppPackage(ctx context.Context, in *ScanAppRequest, opts ...grpc.CallOption) (*ScanResult, error) {
	out := new(ScanResult)
	err := grpc.Invoke(ctx, "/share.ScannerService/ScanAppPackage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scannerServiceClient) Ping(ctx context.Context, in *RPCVoid, opts ...grpc.CallOption) (*RPCVoid, error) {
	out := new(RPCVoid)
	err := grpc.Invoke(ctx, "/share.ScannerService/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scannerServiceClient) ScanAwsLambda(ctx context.Context, in *ScanAwsLambdaRequest, opts ...grpc.CallOption) (*ScanResult, error) {
	out := new(ScanResult)
	err := grpc.Invoke(ctx, "/share.ScannerService/ScanAwsLambda", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scannerServiceClient) ScanCacheGetStat(ctx context.Context, in *RPCVoid, opts ...grpc.CallOption) (*ScanCacheStatRes, error) {
	out := new(ScanCacheStatRes)
	err := grpc.Invoke(ctx, "/share.ScannerService/ScanCacheGetStat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scannerServiceClient) ScanCacheGetData(ctx context.Context, in *RPCVoid, opts ...grpc.CallOption) (*ScanCacheDataRes, error) {
	out := new(ScanCacheDataRes)
	err := grpc.Invoke(ctx, "/share.ScannerService/ScanCacheGetData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ScannerService service

type ScannerServiceServer interface {
	ScanRunning(context.Context, *ScanRunningRequest) (*ScanResult, error)
	ScanImageData(context.Context, *ScanData) (*ScanResult, error)
	ScanImage(context.Context, *ScanImageRequest) (*ScanResult, error)
	ScanAppPackage(context.Context, *ScanAppRequest) (*ScanResult, error)
	Ping(context.Context, *RPCVoid) (*RPCVoid, error)
	ScanAwsLambda(context.Context, *ScanAwsLambdaRequest) (*ScanResult, error)
	ScanCacheGetStat(context.Context, *RPCVoid) (*ScanCacheStatRes, error)
	ScanCacheGetData(context.Context, *RPCVoid) (*ScanCacheDataRes, error)
}

func RegisterScannerServiceServer(s *grpc.Server, srv ScannerServiceServer) {
	s.RegisterService(&_ScannerService_serviceDesc, srv)
}

func _ScannerService_ScanRunning_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScanRunningRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScannerServiceServer).ScanRunning(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/share.ScannerService/ScanRunning",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScannerServiceServer).ScanRunning(ctx, req.(*ScanRunningRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScannerService_ScanImageData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScanData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScannerServiceServer).ScanImageData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/share.ScannerService/ScanImageData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScannerServiceServer).ScanImageData(ctx, req.(*ScanData))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScannerService_ScanImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScanImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScannerServiceServer).ScanImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/share.ScannerService/ScanImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScannerServiceServer).ScanImage(ctx, req.(*ScanImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScannerService_ScanAppPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScanAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScannerServiceServer).ScanAppPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/share.ScannerService/ScanAppPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScannerServiceServer).ScanAppPackage(ctx, req.(*ScanAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScannerService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCVoid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScannerServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/share.ScannerService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScannerServiceServer).Ping(ctx, req.(*RPCVoid))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScannerService_ScanAwsLambda_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScanAwsLambdaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScannerServiceServer).ScanAwsLambda(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/share.ScannerService/ScanAwsLambda",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScannerServiceServer).ScanAwsLambda(ctx, req.(*ScanAwsLambdaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScannerService_ScanCacheGetStat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCVoid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScannerServiceServer).ScanCacheGetStat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/share.ScannerService/ScanCacheGetStat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScannerServiceServer).ScanCacheGetStat(ctx, req.(*RPCVoid))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScannerService_ScanCacheGetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCVoid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScannerServiceServer).ScanCacheGetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/share.ScannerService/ScanCacheGetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScannerServiceServer).ScanCacheGetData(ctx, req.(*RPCVoid))
	}
	return interceptor(ctx, in, info, handler)
}

var _ScannerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "share.ScannerService",
	HandlerType: (*ScannerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ScanRunning",
			Handler:    _ScannerService_ScanRunning_Handler,
		},
		{
			MethodName: "ScanImageData",
			Handler:    _ScannerService_ScanImageData_Handler,
		},
		{
			MethodName: "ScanImage",
			Handler:    _ScannerService_ScanImage_Handler,
		},
		{
			MethodName: "ScanAppPackage",
			Handler:    _ScannerService_ScanAppPackage_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _ScannerService_Ping_Handler,
		},
		{
			MethodName: "ScanAwsLambda",
			Handler:    _ScannerService_ScanAwsLambda_Handler,
		},
		{
			MethodName: "ScanCacheGetStat",
			Handler:    _ScannerService_ScanCacheGetStat_Handler,
		},
		{
			MethodName: "ScanCacheGetData",
			Handler:    _ScannerService_ScanCacheGetData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scanner_service.proto",
}

func init() { proto.RegisterFile("scanner_service.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 742 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xcb, 0x6e, 0x13, 0x31,
	0x14, 0x55, 0xc8, 0x24, 0x4d, 0x9c, 0x10, 0x8a, 0xe9, 0x63, 0x08, 0x0f, 0x45, 0xb3, 0xa8, 0x22,
	0x16, 0x5d, 0xa4, 0x02, 0x24, 0x2a, 0x81, 0xda, 0x80, 0xa0, 0x6a, 0x4b, 0x23, 0x4f, 0xe8, 0xa2,
	0x1b, 0xe4, 0x4c, 0x6e, 0xd3, 0xa1, 0xc9, 0x78, 0xb0, 0x1d, 0x4a, 0x58, 0xf2, 0x35, 0x7c, 0x0a,
	0xdf, 0xc0, 0xb7, 0xb0, 0x40, 0x7e, 0x4c, 0xe3, 0x0c, 0xa1, 0xec, 0x7c, 0xce, 0xb9, 0xef, 0x3b,
	0xf6, 0xa0, 0x75, 0x11, 0xd1, 0x24, 0x01, 0xfe, 0x51, 0x00, 0xff, 0x12, 0x47, 0xb0, 0x9d, 0x72,
	0x26, 0x19, 0x2e, 0x89, 0x0b, 0xca, 0xa1, 0x59, 0x8f, 0xd8, 0x64, 0xc2, 0x12, 0x43, 0x36, 0x91,
	0xb2, 0x35, 0xe7, 0xe0, 0x47, 0x11, 0xad, 0x86, 0x11, 0x4d, 0x0e, 0x26, 0x74, 0x04, 0x04, 0x3e,
	0x4f, 0x41, 0x48, 0xdc, 0x44, 0x15, 0x02, 0xa3, 0x58, 0x48, 0x3e, 0xf3, 0x0b, 0xad, 0x42, 0xbb,
	0x4a, 0xae, 0xb1, 0xd2, 0x3e, 0x08, 0xe0, 0x09, 0x9d, 0x80, 0x7f, 0xcb, 0x68, 0x19, 0x56, 0x5a,
	0x8f, 0x0a, 0x71, 0xc5, 0xf8, 0xd0, 0x2f, 0x1a, 0x2d, 0xc3, 0xf8, 0x31, 0x42, 0x04, 0x52, 0x26,
	0x62, 0xc9, 0xf8, 0xcc, 0xf7, 0xb4, 0xea, 0x30, 0x78, 0x15, 0x15, 0xfb, 0x74, 0xe4, 0x97, 0xb4,
	0xa0, 0x8e, 0x78, 0x0d, 0x95, 0x7a, 0x9c, 0x7d, 0x9d, 0xf9, 0x65, 0xcd, 0x19, 0xa0, 0xe2, 0xa8,
	0x7a, 0x8f, 0xe8, 0x0c, 0xb8, 0xf0, 0x57, 0x5a, 0x85, 0x76, 0x85, 0x38, 0x0c, 0x6e, 0xa1, 0x9a,
	0x42, 0x21, 0x44, 0x1c, 0xa4, 0xf0, 0x2b, 0xda, 0xc0, 0xa5, 0xf0, 0x43, 0x54, 0xdd, 0xa7, 0x02,
	0x74, 0xc7, 0x7e, 0x55, 0xc7, 0x9e, 0x13, 0xf8, 0x25, 0xaa, 0x13, 0xc6, 0xa4, 0x38, 0x39, 0xef,
	0xf3, 0xa9, 0x90, 0x3e, 0x6a, 0x15, 0xdb, 0xb5, 0x4e, 0x73, 0x5b, 0x0f, 0x72, 0x3b, 0x8c, 0x47,
	0x42, 0x32, 0x0e, 0xca, 0xc4, 0x5a, 0x90, 0x05, 0x7b, 0x55, 0x75, 0x9f, 0x5d, 0x42, 0xe2, 0xd7,
	0x4c, 0xd5, 0x1a, 0xe0, 0x7d, 0x84, 0x55, 0x09, 0xfd, 0x59, 0x0a, 0xc2, 0x4e, 0x19, 0x86, 0x7e,
	0xbd, 0x55, 0x68, 0xd7, 0x3a, 0x38, 0x8b, 0x6d, 0x0d, 0x8e, 0x69, 0x4a, 0x96, 0x58, 0x07, 0xbf,
	0x0b, 0xe8, 0xde, 0x92, 0xfc, 0x18, 0x23, 0xef, 0xbd, 0xda, 0x86, 0xd9, 0x94, 0x3e, 0xe3, 0x2d,
	0xd4, 0x20, 0x70, 0xc9, 0x78, 0x6f, 0x3a, 0x18, 0xc7, 0xd1, 0x21, 0xcc, 0xec, 0xae, 0x72, 0xac,
	0xde, 0x34, 0x63, 0xb2, 0x0b, 0x5c, 0x66, 0x1b, 0xcb, 0x30, 0x0e, 0x50, 0x3d, 0xec, 0xf6, 0xe7,
	0x11, 0xcc, 0xce, 0x16, 0x38, 0xfc, 0x14, 0x55, 0x4f, 0x81, 0xc7, 0xe7, 0xb1, 0x5a, 0x46, 0x49,
	0x8f, 0x6a, 0x33, 0x37, 0xaa, 0x4c, 0x27, 0x73, 0x4b, 0xdc, 0x41, 0x6b, 0x2a, 0xcd, 0x18, 0x84,
	0x38, 0x84, 0x59, 0x4a, 0x63, 0x2e, 0x4e, 0x92, 0xb1, 0xd9, 0x74, 0x85, 0x2c, 0xd5, 0x82, 0x9f,
	0x05, 0xb4, 0x9a, 0x8f, 0xb9, 0xb4, 0x77, 0x8c, 0x3c, 0x35, 0x39, 0xdb, 0xb1, 0x3e, 0xe3, 0x37,
	0xa8, 0x61, 0x83, 0x9d, 0xa4, 0x32, 0x66, 0x89, 0xd0, 0xdd, 0xd6, 0x3a, 0x8f, 0x72, 0xc5, 0x2e,
	0x1a, 0x91, 0x9c, 0x93, 0x0d, 0xa3, 0x4a, 0xcb, 0xc2, 0x78, 0xff, 0x0a, 0xe3, 0x18, 0x91, 0x9c,
	0x53, 0xf0, 0x0c, 0x6d, 0x2c, 0x4f, 0xa8, 0xbe, 0xcd, 0xf9, 0xc0, 0x4d, 0x53, 0x73, 0x22, 0x38,
	0x5b, 0xf0, 0x73, 0x22, 0xaa, 0x5b, 0xa1, 0x76, 0x76, 0x20, 0xc4, 0x14, 0xb8, 0x75, 0x74, 0x18,
	0x75, 0x2b, 0x14, 0x0a, 0xa7, 0x83, 0x4f, 0x10, 0x49, 0x3b, 0x1a, 0x97, 0x0a, 0xbe, 0x17, 0xcc,
	0x43, 0xd0, 0xa5, 0xd1, 0x05, 0x84, 0x92, 0x4a, 0x02, 0xba, 0x1c, 0x02, 0x11, 0xe3, 0xc3, 0x6e,
	0x22, 0x75, 0x54, 0x8f, 0xcc, 0x09, 0x73, 0xa5, 0x15, 0x08, 0xe3, 0x6f, 0x66, 0xdc, 0x1e, 0x71,
	0x18, 0xec, 0xa3, 0x95, 0xe3, 0x58, 0x08, 0xe5, 0x5b, 0xd4, 0x62, 0x06, 0xf1, 0x06, 0x2a, 0xbf,
	0x8b, 0xa5, 0x12, 0x3c, 0x2d, 0x58, 0x14, 0x3c, 0x71, 0x6a, 0x78, 0x4d, 0x25, 0x55, 0x35, 0x6c,
	0xa0, 0xb2, 0x3a, 0x9e, 0x0d, 0x74, 0x01, 0x75, 0x62, 0x51, 0xe7, 0x57, 0x11, 0x35, 0x42, 0xf3,
	0xe8, 0x85, 0xe6, 0xcd, 0xc3, 0xbb, 0xe6, 0xee, 0x93, 0x69, 0x92, 0xc4, 0xc9, 0x08, 0xdf, 0x77,
	0x2e, 0x96, 0xe5, 0xec, 0x6d, 0x6a, 0xde, 0x75, 0x25, 0x10, 0xd3, 0xb1, 0xc4, 0x3b, 0xe8, 0xf6,
	0xf5, 0x43, 0xa8, 0x52, 0xe0, 0x3b, 0x8e, 0x8d, 0x22, 0x96, 0x39, 0x3d, 0x47, 0xd5, 0x6b, 0x27,
	0xbc, 0xe9, 0xe8, 0xee, 0x7b, 0xba, 0xcc, 0xf1, 0x85, 0x29, 0x7e, 0x2f, 0x4d, 0x7b, 0x34, 0xba,
	0x54, 0xde, 0xeb, 0x8e, 0xd1, 0x5e, 0x9a, 0xde, 0xe0, 0xbb, 0x85, 0xbc, 0x9e, 0xea, 0xaf, 0x61,
	0x25, 0xd2, 0xeb, 0x9e, 0xb2, 0x78, 0xd8, 0xcc, 0x61, 0xfc, 0xca, 0x74, 0xb4, 0x77, 0x25, 0x8e,
	0xe8, 0x64, 0x30, 0xa4, 0xf8, 0x81, 0x9b, 0x22, 0x63, 0x6f, 0x48, 0xb4, 0xeb, 0xac, 0xe3, 0x2d,
	0x48, 0xf5, 0x55, 0xfc, 0x95, 0xd4, 0x6d, 0x7a, 0xe1, 0xdb, 0xc9, 0x39, 0xeb, 0x91, 0xfe, 0xd7,
	0xd9, 0x2e, 0x7d, 0x50, 0xd6, 0x7f, 0xa7, 0x9d, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x64, 0xea,
	0x44, 0x52, 0xd7, 0x06, 0x00, 0x00,
}
