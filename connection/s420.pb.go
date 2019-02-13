// Code generated by protoc-gen-go. DO NOT EDIT.
// source: connection/s420.proto

package s420con

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type NewObjectInBucket struct {
	Bucket               string         `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	FileName             string         `protobuf:"bytes,2,opt,name=fileName,proto3" json:"fileName,omitempty"`
	Data                 []byte         `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Options              *ObjectOptions `protobuf:"bytes,4,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *NewObjectInBucket) Reset()         { *m = NewObjectInBucket{} }
func (m *NewObjectInBucket) String() string { return proto.CompactTextString(m) }
func (*NewObjectInBucket) ProtoMessage()    {}
func (*NewObjectInBucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_5013b312c4546119, []int{0}
}

func (m *NewObjectInBucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewObjectInBucket.Unmarshal(m, b)
}
func (m *NewObjectInBucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewObjectInBucket.Marshal(b, m, deterministic)
}
func (m *NewObjectInBucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewObjectInBucket.Merge(m, src)
}
func (m *NewObjectInBucket) XXX_Size() int {
	return xxx_messageInfo_NewObjectInBucket.Size(m)
}
func (m *NewObjectInBucket) XXX_DiscardUnknown() {
	xxx_messageInfo_NewObjectInBucket.DiscardUnknown(m)
}

var xxx_messageInfo_NewObjectInBucket proto.InternalMessageInfo

func (m *NewObjectInBucket) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *NewObjectInBucket) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *NewObjectInBucket) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *NewObjectInBucket) GetOptions() *ObjectOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type NewObjectParams struct {
	Path                 string         `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Data                 []byte         `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Options              *ObjectOptions `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *NewObjectParams) Reset()         { *m = NewObjectParams{} }
func (m *NewObjectParams) String() string { return proto.CompactTextString(m) }
func (*NewObjectParams) ProtoMessage()    {}
func (*NewObjectParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_5013b312c4546119, []int{1}
}

func (m *NewObjectParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewObjectParams.Unmarshal(m, b)
}
func (m *NewObjectParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewObjectParams.Marshal(b, m, deterministic)
}
func (m *NewObjectParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewObjectParams.Merge(m, src)
}
func (m *NewObjectParams) XXX_Size() int {
	return xxx_messageInfo_NewObjectParams.Size(m)
}
func (m *NewObjectParams) XXX_DiscardUnknown() {
	xxx_messageInfo_NewObjectParams.DiscardUnknown(m)
}

var xxx_messageInfo_NewObjectParams proto.InternalMessageInfo

func (m *NewObjectParams) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *NewObjectParams) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *NewObjectParams) GetOptions() *ObjectOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type ObjectFromBucket struct {
	Bucket               string   `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectFromBucket) Reset()         { *m = ObjectFromBucket{} }
func (m *ObjectFromBucket) String() string { return proto.CompactTextString(m) }
func (*ObjectFromBucket) ProtoMessage()    {}
func (*ObjectFromBucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_5013b312c4546119, []int{2}
}

func (m *ObjectFromBucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectFromBucket.Unmarshal(m, b)
}
func (m *ObjectFromBucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectFromBucket.Marshal(b, m, deterministic)
}
func (m *ObjectFromBucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectFromBucket.Merge(m, src)
}
func (m *ObjectFromBucket) XXX_Size() int {
	return xxx_messageInfo_ObjectFromBucket.Size(m)
}
func (m *ObjectFromBucket) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectFromBucket.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectFromBucket proto.InternalMessageInfo

func (m *ObjectFromBucket) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *ObjectFromBucket) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type ObjectPath struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectPath) Reset()         { *m = ObjectPath{} }
func (m *ObjectPath) String() string { return proto.CompactTextString(m) }
func (*ObjectPath) ProtoMessage()    {}
func (*ObjectPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_5013b312c4546119, []int{3}
}

func (m *ObjectPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectPath.Unmarshal(m, b)
}
func (m *ObjectPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectPath.Marshal(b, m, deterministic)
}
func (m *ObjectPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectPath.Merge(m, src)
}
func (m *ObjectPath) XXX_Size() int {
	return xxx_messageInfo_ObjectPath.Size(m)
}
func (m *ObjectPath) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectPath.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectPath proto.InternalMessageInfo

func (m *ObjectPath) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type ObjectOptions struct {
	ContentType          string   `protobuf:"bytes,1,opt,name=contentType,proto3" json:"contentType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectOptions) Reset()         { *m = ObjectOptions{} }
func (m *ObjectOptions) String() string { return proto.CompactTextString(m) }
func (*ObjectOptions) ProtoMessage()    {}
func (*ObjectOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_5013b312c4546119, []int{4}
}

func (m *ObjectOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectOptions.Unmarshal(m, b)
}
func (m *ObjectOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectOptions.Marshal(b, m, deterministic)
}
func (m *ObjectOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectOptions.Merge(m, src)
}
func (m *ObjectOptions) XXX_Size() int {
	return xxx_messageInfo_ObjectOptions.Size(m)
}
func (m *ObjectOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectOptions proto.InternalMessageInfo

func (m *ObjectOptions) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

type SaveResponse struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	FilePath             string   `protobuf:"bytes,2,opt,name=filePath,proto3" json:"filePath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveResponse) Reset()         { *m = SaveResponse{} }
func (m *SaveResponse) String() string { return proto.CompactTextString(m) }
func (*SaveResponse) ProtoMessage()    {}
func (*SaveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5013b312c4546119, []int{5}
}

func (m *SaveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveResponse.Unmarshal(m, b)
}
func (m *SaveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveResponse.Marshal(b, m, deterministic)
}
func (m *SaveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveResponse.Merge(m, src)
}
func (m *SaveResponse) XXX_Size() int {
	return xxx_messageInfo_SaveResponse.Size(m)
}
func (m *SaveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SaveResponse proto.InternalMessageInfo

func (m *SaveResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *SaveResponse) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

type GetResponse struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	ContentType          string   `protobuf:"bytes,2,opt,name=contentType,proto3" json:"contentType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5013b312c4546119, []int{6}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GetResponse) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func init() {
	proto.RegisterType((*NewObjectInBucket)(nil), "s420.NewObjectInBucket")
	proto.RegisterType((*NewObjectParams)(nil), "s420.NewObjectParams")
	proto.RegisterType((*ObjectFromBucket)(nil), "s420.ObjectFromBucket")
	proto.RegisterType((*ObjectPath)(nil), "s420.ObjectPath")
	proto.RegisterType((*ObjectOptions)(nil), "s420.ObjectOptions")
	proto.RegisterType((*SaveResponse)(nil), "s420.SaveResponse")
	proto.RegisterType((*GetResponse)(nil), "s420.GetResponse")
}

func init() { proto.RegisterFile("connection/s420.proto", fileDescriptor_5013b312c4546119) }

var fileDescriptor_5013b312c4546119 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xbf, 0x4f, 0xc2, 0x40,
	0x14, 0xc7, 0x69, 0x69, 0xf8, 0xf1, 0x00, 0x85, 0x47, 0xd0, 0x86, 0xa9, 0xb9, 0x89, 0x45, 0xc4,
	0xca, 0x64, 0x8c, 0x03, 0x24, 0x12, 0x17, 0x20, 0xc5, 0xc9, 0xad, 0xd4, 0x33, 0x20, 0x72, 0xd7,
	0xd0, 0x53, 0xe3, 0xe8, 0xea, 0x5f, 0x6d, 0x7a, 0x47, 0x8f, 0x5a, 0x30, 0xba, 0xbd, 0x5f, 0xbd,
	0xef, 0xe7, 0xbe, 0x7d, 0x07, 0xad, 0x80, 0x33, 0x46, 0x03, 0xb1, 0xe4, 0xec, 0x3c, 0xea, 0xbb,
	0xbd, 0x6e, 0xb8, 0xe1, 0x82, 0xa3, 0x15, 0xc7, 0xe4, 0xcb, 0x80, 0xc6, 0x98, 0xbe, 0x4f, 0xe6,
	0xcf, 0x34, 0x10, 0x77, 0x6c, 0xf0, 0x1a, 0xac, 0xa8, 0xc0, 0x13, 0x28, 0xcc, 0x65, 0x64, 0x1b,
	0x8e, 0xd1, 0x29, 0x7b, 0xdb, 0x0c, 0xdb, 0x50, 0x7a, 0x5a, 0xbe, 0xd0, 0xb1, 0xbf, 0xa6, 0xb6,
	0x29, 0x3b, 0x3a, 0x47, 0x04, 0xeb, 0xd1, 0x17, 0xbe, 0x9d, 0x77, 0x8c, 0x4e, 0xd5, 0x93, 0x31,
	0x9e, 0x41, 0x91, 0x87, 0xb1, 0x70, 0x64, 0x5b, 0x8e, 0xd1, 0xa9, 0xb8, 0xcd, 0xae, 0x24, 0x50,
	0x72, 0x13, 0xd5, 0xf2, 0x92, 0x19, 0xb2, 0x80, 0x63, 0xcd, 0x32, 0xf5, 0x37, 0xfe, 0x3a, 0x8a,
	0x4f, 0x0d, 0x7d, 0xb1, 0xd8, 0x72, 0xc8, 0x58, 0x2b, 0x99, 0x87, 0x95, 0xf2, 0xff, 0x50, 0xba,
	0x81, 0xba, 0xea, 0xdc, 0x6e, 0xf8, 0xfa, 0x8f, 0x4b, 0x27, 0x08, 0xe6, 0x0e, 0x81, 0x38, 0x00,
	0x09, 0xa6, 0x02, 0xca, 0x42, 0x92, 0x0b, 0xa8, 0xfd, 0xd0, 0x46, 0x07, 0x2a, 0x01, 0x67, 0x82,
	0x32, 0x71, 0xff, 0x11, 0xd2, 0xed, 0x6c, 0xba, 0x44, 0xae, 0xa0, 0x3a, 0xf3, 0xdf, 0xa8, 0x47,
	0xa3, 0x90, 0xb3, 0x88, 0xe2, 0x11, 0x98, 0x7c, 0x25, 0x07, 0x4b, 0x9e, 0xc9, 0x57, 0x89, 0xfb,
	0xd3, 0x1d, 0x8c, 0xce, 0xc9, 0x10, 0x2a, 0x23, 0x2a, 0xf4, 0xa7, 0x89, 0x45, 0x46, 0xca, 0xa2,
	0x0c, 0x80, 0xb9, 0x07, 0xe0, 0x7e, 0x9a, 0x60, 0xcd, 0xfa, 0x6e, 0x0f, 0x87, 0x80, 0x31, 0x49,
	0x66, 0x2b, 0x4e, 0x95, 0xa5, 0x7b, 0xeb, 0xd2, 0x46, 0xd5, 0x48, 0xc3, 0x93, 0x1c, 0x0e, 0xa0,
	0x39, 0xa2, 0x62, 0xdf, 0xe6, 0xf4, 0x8f, 0xd9, 0xd5, 0xdb, 0x0d, 0x55, 0x4f, 0xdd, 0x82, 0xe4,
	0xf0, 0x1a, 0x6a, 0xf1, 0xa9, 0x5a, 0x12, 0x5b, 0x19, 0x06, 0xb5, 0x26, 0xbf, 0x10, 0xb8, 0x50,
	0xd6, 0x04, 0x58, 0x4f, 0xeb, 0xc6, 0x9e, 0x1d, 0x54, 0x1c, 0x94, 0x1f, 0x8a, 0x71, 0x35, 0xe0,
	0x6c, 0x5e, 0x90, 0x0f, 0xe5, 0xf2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x4f, 0x8a, 0x2d, 0x41,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// S420Client is the client API for S420 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type S420Client interface {
	SaveObjectInBucket(ctx context.Context, in *NewObjectInBucket, opts ...grpc.CallOption) (*SaveResponse, error)
	GetObjectFromBucket(ctx context.Context, in *ObjectFromBucket, opts ...grpc.CallOption) (*GetResponse, error)
	SaveNewObject(ctx context.Context, in *NewObjectParams, opts ...grpc.CallOption) (*SaveResponse, error)
	GetObject(ctx context.Context, in *ObjectPath, opts ...grpc.CallOption) (*GetResponse, error)
}

type s420Client struct {
	cc *grpc.ClientConn
}

func NewS420Client(cc *grpc.ClientConn) S420Client {
	return &s420Client{cc}
}

func (c *s420Client) SaveObjectInBucket(ctx context.Context, in *NewObjectInBucket, opts ...grpc.CallOption) (*SaveResponse, error) {
	out := new(SaveResponse)
	err := c.cc.Invoke(ctx, "/s420.S420/SaveObjectInBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s420Client) GetObjectFromBucket(ctx context.Context, in *ObjectFromBucket, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/s420.S420/GetObjectFromBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s420Client) SaveNewObject(ctx context.Context, in *NewObjectParams, opts ...grpc.CallOption) (*SaveResponse, error) {
	out := new(SaveResponse)
	err := c.cc.Invoke(ctx, "/s420.S420/SaveNewObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s420Client) GetObject(ctx context.Context, in *ObjectPath, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/s420.S420/GetObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// S420Server is the server API for S420 service.
type S420Server interface {
	SaveObjectInBucket(context.Context, *NewObjectInBucket) (*SaveResponse, error)
	GetObjectFromBucket(context.Context, *ObjectFromBucket) (*GetResponse, error)
	SaveNewObject(context.Context, *NewObjectParams) (*SaveResponse, error)
	GetObject(context.Context, *ObjectPath) (*GetResponse, error)
}

func RegisterS420Server(s *grpc.Server, srv S420Server) {
	s.RegisterService(&_S420_serviceDesc, srv)
}

func _S420_SaveObjectInBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewObjectInBucket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S420Server).SaveObjectInBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/s420.S420/SaveObjectInBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S420Server).SaveObjectInBucket(ctx, req.(*NewObjectInBucket))
	}
	return interceptor(ctx, in, info, handler)
}

func _S420_GetObjectFromBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectFromBucket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S420Server).GetObjectFromBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/s420.S420/GetObjectFromBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S420Server).GetObjectFromBucket(ctx, req.(*ObjectFromBucket))
	}
	return interceptor(ctx, in, info, handler)
}

func _S420_SaveNewObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewObjectParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S420Server).SaveNewObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/s420.S420/SaveNewObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S420Server).SaveNewObject(ctx, req.(*NewObjectParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _S420_GetObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectPath)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S420Server).GetObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/s420.S420/GetObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S420Server).GetObject(ctx, req.(*ObjectPath))
	}
	return interceptor(ctx, in, info, handler)
}

var _S420_serviceDesc = grpc.ServiceDesc{
	ServiceName: "s420.S420",
	HandlerType: (*S420Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveObjectInBucket",
			Handler:    _S420_SaveObjectInBucket_Handler,
		},
		{
			MethodName: "GetObjectFromBucket",
			Handler:    _S420_GetObjectFromBucket_Handler,
		},
		{
			MethodName: "SaveNewObject",
			Handler:    _S420_SaveNewObject_Handler,
		},
		{
			MethodName: "GetObject",
			Handler:    _S420_GetObject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "connection/s420.proto",
}
