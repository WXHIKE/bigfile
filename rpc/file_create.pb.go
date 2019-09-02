// Code generated by protoc-gen-go. DO NOT EDIT.
// source: file_create.proto

package rpc

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// FileCreateRequest represent the file create request
type FileCreateRequest struct {
	Token  string                `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Path   string                `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Secret *wrappers.StringValue `protobuf:"bytes,4,opt,name=secret,proto3" json:"secret,omitempty"`
	Hidden *wrappers.BoolValue   `protobuf:"bytes,6,opt,name=hidden,proto3" json:"hidden,omitempty"`
	// Types that are valid to be assigned to Operation:
	//	*FileCreateRequest_Overwrite
	//	*FileCreateRequest_Rename
	//	*FileCreateRequest_Append
	//	*FileCreateRequest_CreateDir
	//	*FileCreateRequest_None
	Operation            isFileCreateRequest_Operation `protobuf_oneof:"operation"`
	Content              *wrappers.BytesValue          `protobuf:"bytes,12,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *FileCreateRequest) Reset()         { *m = FileCreateRequest{} }
func (m *FileCreateRequest) String() string { return proto.CompactTextString(m) }
func (*FileCreateRequest) ProtoMessage()    {}
func (*FileCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8a75d4c3ddc50ae, []int{0}
}

func (m *FileCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileCreateRequest.Unmarshal(m, b)
}
func (m *FileCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileCreateRequest.Marshal(b, m, deterministic)
}
func (m *FileCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileCreateRequest.Merge(m, src)
}
func (m *FileCreateRequest) XXX_Size() int {
	return xxx_messageInfo_FileCreateRequest.Size(m)
}
func (m *FileCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FileCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FileCreateRequest proto.InternalMessageInfo

func (m *FileCreateRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *FileCreateRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *FileCreateRequest) GetSecret() *wrappers.StringValue {
	if m != nil {
		return m.Secret
	}
	return nil
}

func (m *FileCreateRequest) GetHidden() *wrappers.BoolValue {
	if m != nil {
		return m.Hidden
	}
	return nil
}

type isFileCreateRequest_Operation interface {
	isFileCreateRequest_Operation()
}

type FileCreateRequest_Overwrite struct {
	Overwrite bool `protobuf:"varint,7,opt,name=overwrite,proto3,oneof"`
}

type FileCreateRequest_Rename struct {
	Rename bool `protobuf:"varint,8,opt,name=rename,proto3,oneof"`
}

type FileCreateRequest_Append struct {
	Append bool `protobuf:"varint,9,opt,name=append,proto3,oneof"`
}

type FileCreateRequest_CreateDir struct {
	CreateDir bool `protobuf:"varint,10,opt,name=create_dir,json=createDir,proto3,oneof"`
}

type FileCreateRequest_None struct {
	None bool `protobuf:"varint,11,opt,name=none,proto3,oneof"`
}

func (*FileCreateRequest_Overwrite) isFileCreateRequest_Operation() {}

func (*FileCreateRequest_Rename) isFileCreateRequest_Operation() {}

func (*FileCreateRequest_Append) isFileCreateRequest_Operation() {}

func (*FileCreateRequest_CreateDir) isFileCreateRequest_Operation() {}

func (*FileCreateRequest_None) isFileCreateRequest_Operation() {}

func (m *FileCreateRequest) GetOperation() isFileCreateRequest_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *FileCreateRequest) GetOverwrite() bool {
	if x, ok := m.GetOperation().(*FileCreateRequest_Overwrite); ok {
		return x.Overwrite
	}
	return false
}

func (m *FileCreateRequest) GetRename() bool {
	if x, ok := m.GetOperation().(*FileCreateRequest_Rename); ok {
		return x.Rename
	}
	return false
}

func (m *FileCreateRequest) GetAppend() bool {
	if x, ok := m.GetOperation().(*FileCreateRequest_Append); ok {
		return x.Append
	}
	return false
}

func (m *FileCreateRequest) GetCreateDir() bool {
	if x, ok := m.GetOperation().(*FileCreateRequest_CreateDir); ok {
		return x.CreateDir
	}
	return false
}

func (m *FileCreateRequest) GetNone() bool {
	if x, ok := m.GetOperation().(*FileCreateRequest_None); ok {
		return x.None
	}
	return false
}

func (m *FileCreateRequest) GetContent() *wrappers.BytesValue {
	if m != nil {
		return m.Content
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FileCreateRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FileCreateRequest_Overwrite)(nil),
		(*FileCreateRequest_Rename)(nil),
		(*FileCreateRequest_Append)(nil),
		(*FileCreateRequest_CreateDir)(nil),
		(*FileCreateRequest_None)(nil),
	}
}

// FileCreateResponse represent the response from creating file
type FileCreateResponse struct {
	RequestId            uint64   `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	File                 *File    `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileCreateResponse) Reset()         { *m = FileCreateResponse{} }
func (m *FileCreateResponse) String() string { return proto.CompactTextString(m) }
func (*FileCreateResponse) ProtoMessage()    {}
func (*FileCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8a75d4c3ddc50ae, []int{1}
}

func (m *FileCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileCreateResponse.Unmarshal(m, b)
}
func (m *FileCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileCreateResponse.Marshal(b, m, deterministic)
}
func (m *FileCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileCreateResponse.Merge(m, src)
}
func (m *FileCreateResponse) XXX_Size() int {
	return xxx_messageInfo_FileCreateResponse.Size(m)
}
func (m *FileCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FileCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FileCreateResponse proto.InternalMessageInfo

func (m *FileCreateResponse) GetRequestId() uint64 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *FileCreateResponse) GetFile() *File {
	if m != nil {
		return m.File
	}
	return nil
}

func init() {
	proto.RegisterType((*FileCreateRequest)(nil), "bigfile.file_create.FileCreateRequest")
	proto.RegisterType((*FileCreateResponse)(nil), "bigfile.file_create.FileCreateResponse")
}

func init() { proto.RegisterFile("file_create.proto", fileDescriptor_d8a75d4c3ddc50ae) }

var fileDescriptor_d8a75d4c3ddc50ae = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0x5e, 0x97, 0x90, 0xdd, 0x4c, 0x91, 0xd0, 0x9a, 0x1e, 0xac, 0x02, 0xdd, 0xaa, 0x87, 0xa5,
	0x27, 0x17, 0x15, 0x78, 0x81, 0x80, 0x10, 0x88, 0x4b, 0x15, 0x10, 0x48, 0x70, 0xa8, 0xd2, 0x64,
	0x36, 0xb5, 0x48, 0x6d, 0xaf, 0xe3, 0x52, 0xed, 0xeb, 0x70, 0xe4, 0x01, 0x11, 0x47, 0x14, 0xdb,
	0x69, 0x2b, 0x15, 0x89, 0x53, 0xf2, 0xfd, 0x8d, 0xed, 0x99, 0x81, 0xcb, 0x1b, 0x51, 0xe3, 0xb2,
	0x30, 0x98, 0x5b, 0xe4, 0xda, 0x28, 0xab, 0xe8, 0xa3, 0x95, 0xa8, 0x5a, 0x96, 0x1f, 0x49, 0x43,
	0x70, 0x8c, 0x33, 0x0c, 0x47, 0x95, 0x52, 0x55, 0x8d, 0x33, 0x87, 0x56, 0xdb, 0x9b, 0xd9, 0xce,
	0xe4, 0x5a, 0xa3, 0x69, 0xbc, 0x3e, 0xf9, 0xdd, 0x83, 0xcb, 0xb7, 0xa2, 0xc6, 0xd7, 0x2e, 0x9a,
	0xe1, 0xed, 0x16, 0x1b, 0x4b, 0x07, 0x70, 0xdf, 0xaa, 0xef, 0x28, 0x19, 0x19, 0x93, 0x69, 0x92,
	0x79, 0x40, 0x29, 0x44, 0x3a, 0xb7, 0x6b, 0xd6, 0x73, 0xa4, 0xfb, 0xa7, 0x2f, 0x21, 0x6e, 0xb0,
	0x30, 0x68, 0x59, 0x34, 0x26, 0xd3, 0xfe, 0xfc, 0x09, 0xf7, 0x07, 0xf2, 0xee, 0x40, 0xfe, 0xd1,
	0x1a, 0x21, 0xab, 0xcf, 0x79, 0xbd, 0xc5, 0x2c, 0x78, 0xe9, 0x1c, 0xe2, 0xb5, 0x28, 0x4b, 0x94,
	0x2c, 0x76, 0xa9, 0xe1, 0x49, 0x2a, 0x55, 0xaa, 0x0e, 0x19, 0xef, 0xa4, 0x23, 0x48, 0xd4, 0x0f,
	0x34, 0x3b, 0x23, 0x2c, 0xb2, 0xf3, 0x31, 0x99, 0x5e, 0xbc, 0x3b, 0xcb, 0x0e, 0x14, 0x65, 0x10,
	0x1b, 0x94, 0xf9, 0x06, 0xd9, 0x45, 0x10, 0x03, 0x6e, 0x95, 0xf6, 0xcd, 0xb2, 0x64, 0x49, 0xa7,
	0x78, 0x4c, 0xaf, 0x00, 0x7c, 0xcf, 0x96, 0xa5, 0x30, 0x0c, 0xba, 0xa2, 0x9e, 0x7b, 0x23, 0x0c,
	0x1d, 0x40, 0x24, 0x95, 0x44, 0xd6, 0x0f, 0x92, 0x43, 0xf4, 0x15, 0x9c, 0x17, 0x4a, 0x5a, 0x94,
	0x96, 0x3d, 0x70, 0xf7, 0x7f, 0x7c, 0x7a, 0xff, 0x3b, 0x8b, 0x8d, 0x7f, 0x40, 0xe7, 0x4d, 0xfb,
	0x90, 0x28, 0x8d, 0x26, 0xb7, 0x42, 0xc9, 0xc9, 0x37, 0xa0, 0xc7, 0x7d, 0x6f, 0xb4, 0x92, 0x0d,
	0xd2, 0xa7, 0x00, 0xc6, 0xcf, 0x60, 0x29, 0x4a, 0xd7, 0xfd, 0x28, 0x4b, 0x02, 0xf3, 0xbe, 0xa4,
	0xd7, 0x10, 0xb5, 0xb3, 0x75, 0x13, 0xe8, 0xcf, 0x29, 0x3f, 0x9e, 0x3e, 0x6f, 0xcb, 0x65, 0x4e,
	0x9f, 0xdf, 0x02, 0x1c, 0x8a, 0xd3, 0x02, 0xdc, 0x46, 0x04, 0x74, 0xcd, 0xff, 0xb1, 0x33, 0xfc,
	0x64, 0x07, 0x86, 0xcf, 0xfe, 0xeb, 0xf3, 0x77, 0x9e, 0x9c, 0x4d, 0xc9, 0x73, 0x92, 0x5a, 0x18,
	0x14, 0x6a, 0xb3, 0xcf, 0x74, 0x8d, 0x48, 0x1f, 0x1e, 0x12, 0x8b, 0x96, 0x5b, 0x90, 0xaf, 0xa3,
	0x4a, 0xd8, 0xf5, 0x76, 0xc5, 0x0b, 0xb5, 0x99, 0x05, 0xff, 0xfe, 0x6b, 0x74, 0xf1, 0x87, 0x90,
	0x9f, 0xbd, 0x7b, 0xe9, 0x22, 0xfb, 0xd5, 0xbb, 0x4a, 0x43, 0xb9, 0x45, 0xd7, 0xd7, 0x2f, 0x58,
	0xd7, 0x1f, 0xa4, 0xda, 0xc9, 0x4f, 0x77, 0x1a, 0x9b, 0x55, 0xec, 0xce, 0x79, 0xf1, 0x37, 0x00,
	0x00, 0xff, 0xff, 0x0b, 0x2a, 0x8c, 0x5e, 0x1b, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FileCreateClient is the client API for FileCreate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileCreateClient interface {
	FileCreate(ctx context.Context, opts ...grpc.CallOption) (FileCreate_FileCreateClient, error)
}

type fileCreateClient struct {
	cc *grpc.ClientConn
}

func NewFileCreateClient(cc *grpc.ClientConn) FileCreateClient {
	return &fileCreateClient{cc}
}

func (c *fileCreateClient) FileCreate(ctx context.Context, opts ...grpc.CallOption) (FileCreate_FileCreateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileCreate_serviceDesc.Streams[0], "/bigfile.file_create.FileCreate/fileCreate", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileCreateFileCreateClient{stream}
	return x, nil
}

type FileCreate_FileCreateClient interface {
	Send(*FileCreateRequest) error
	Recv() (*FileCreateResponse, error)
	grpc.ClientStream
}

type fileCreateFileCreateClient struct {
	grpc.ClientStream
}

func (x *fileCreateFileCreateClient) Send(m *FileCreateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileCreateFileCreateClient) Recv() (*FileCreateResponse, error) {
	m := new(FileCreateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileCreateServer is the server API for FileCreate service.
type FileCreateServer interface {
	FileCreate(FileCreate_FileCreateServer) error
}

// UnimplementedFileCreateServer can be embedded to have forward compatible implementations.
type UnimplementedFileCreateServer struct {
}

func (*UnimplementedFileCreateServer) FileCreate(srv FileCreate_FileCreateServer) error {
	return status.Errorf(codes.Unimplemented, "method FileCreate not implemented")
}

func RegisterFileCreateServer(s *grpc.Server, srv FileCreateServer) {
	s.RegisterService(&_FileCreate_serviceDesc, srv)
}

func _FileCreate_FileCreate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileCreateServer).FileCreate(&fileCreateFileCreateServer{stream})
}

type FileCreate_FileCreateServer interface {
	Send(*FileCreateResponse) error
	Recv() (*FileCreateRequest, error)
	grpc.ServerStream
}

type fileCreateFileCreateServer struct {
	grpc.ServerStream
}

func (x *fileCreateFileCreateServer) Send(m *FileCreateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileCreateFileCreateServer) Recv() (*FileCreateRequest, error) {
	m := new(FileCreateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _FileCreate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bigfile.file_create.FileCreate",
	HandlerType: (*FileCreateServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "fileCreate",
			Handler:       _FileCreate_FileCreate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "file_create.proto",
}
