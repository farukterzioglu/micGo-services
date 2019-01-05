// Code generated by protoc-gen-go. DO NOT EDIT.
// source: review_service.proto

package reviewservice

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

type Review struct {
	Text                 string   `protobuf:"bytes,1,opt,name=Text,proto3" json:"Text,omitempty"`
	Star                 int32    `protobuf:"varint,2,opt,name=Star,proto3" json:"Star,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Review) Reset()         { *m = Review{} }
func (m *Review) String() string { return proto.CompactTextString(m) }
func (*Review) ProtoMessage()    {}
func (*Review) Descriptor() ([]byte, []int) {
	return fileDescriptor_e02e8736aac43c55, []int{0}
}

func (m *Review) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Review.Unmarshal(m, b)
}
func (m *Review) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Review.Marshal(b, m, deterministic)
}
func (m *Review) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Review.Merge(m, src)
}
func (m *Review) XXX_Size() int {
	return xxx_messageInfo_Review.Size(m)
}
func (m *Review) XXX_DiscardUnknown() {
	xxx_messageInfo_Review.DiscardUnknown(m)
}

var xxx_messageInfo_Review proto.InternalMessageInfo

func (m *Review) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Review) GetStar() int32 {
	if m != nil {
		return m.Star
	}
	return 0
}

type NewReviewRequest struct {
	Review               *Review  `protobuf:"bytes,1,opt,name=review,proto3" json:"review,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewReviewRequest) Reset()         { *m = NewReviewRequest{} }
func (m *NewReviewRequest) String() string { return proto.CompactTextString(m) }
func (*NewReviewRequest) ProtoMessage()    {}
func (*NewReviewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e02e8736aac43c55, []int{1}
}

func (m *NewReviewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewReviewRequest.Unmarshal(m, b)
}
func (m *NewReviewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewReviewRequest.Marshal(b, m, deterministic)
}
func (m *NewReviewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewReviewRequest.Merge(m, src)
}
func (m *NewReviewRequest) XXX_Size() int {
	return xxx_messageInfo_NewReviewRequest.Size(m)
}
func (m *NewReviewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewReviewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewReviewRequest proto.InternalMessageInfo

func (m *NewReviewRequest) GetReview() *Review {
	if m != nil {
		return m.Review
	}
	return nil
}

type ReviewId struct {
	ReviewId             string   `protobuf:"bytes,1,opt,name=reviewId,proto3" json:"reviewId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReviewId) Reset()         { *m = ReviewId{} }
func (m *ReviewId) String() string { return proto.CompactTextString(m) }
func (*ReviewId) ProtoMessage()    {}
func (*ReviewId) Descriptor() ([]byte, []int) {
	return fileDescriptor_e02e8736aac43c55, []int{2}
}

func (m *ReviewId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReviewId.Unmarshal(m, b)
}
func (m *ReviewId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReviewId.Marshal(b, m, deterministic)
}
func (m *ReviewId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReviewId.Merge(m, src)
}
func (m *ReviewId) XXX_Size() int {
	return xxx_messageInfo_ReviewId.Size(m)
}
func (m *ReviewId) XXX_DiscardUnknown() {
	xxx_messageInfo_ReviewId.DiscardUnknown(m)
}

var xxx_messageInfo_ReviewId proto.InternalMessageInfo

func (m *ReviewId) GetReviewId() string {
	if m != nil {
		return m.ReviewId
	}
	return ""
}

type GetTopReviewsRequest struct {
	Count                int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTopReviewsRequest) Reset()         { *m = GetTopReviewsRequest{} }
func (m *GetTopReviewsRequest) String() string { return proto.CompactTextString(m) }
func (*GetTopReviewsRequest) ProtoMessage()    {}
func (*GetTopReviewsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e02e8736aac43c55, []int{3}
}

func (m *GetTopReviewsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTopReviewsRequest.Unmarshal(m, b)
}
func (m *GetTopReviewsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTopReviewsRequest.Marshal(b, m, deterministic)
}
func (m *GetTopReviewsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTopReviewsRequest.Merge(m, src)
}
func (m *GetTopReviewsRequest) XXX_Size() int {
	return xxx_messageInfo_GetTopReviewsRequest.Size(m)
}
func (m *GetTopReviewsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTopReviewsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTopReviewsRequest proto.InternalMessageInfo

func (m *GetTopReviewsRequest) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*Review)(nil), "reviewservice.Review")
	proto.RegisterType((*NewReviewRequest)(nil), "reviewservice.NewReviewRequest")
	proto.RegisterType((*ReviewId)(nil), "reviewservice.ReviewId")
	proto.RegisterType((*GetTopReviewsRequest)(nil), "reviewservice.GetTopReviewsRequest")
}

func init() { proto.RegisterFile("review_service.proto", fileDescriptor_e02e8736aac43c55) }

var fileDescriptor_e02e8736aac43c55 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x51, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0xed, 0x8a, 0x09, 0x75, 0x4a, 0x40, 0x86, 0x88, 0x25, 0x17, 0xcb, 0x0a, 0x92, 0x83, 0x86,
	0x50, 0x7f, 0x81, 0x27, 0xed, 0x41, 0x85, 0x4d, 0xef, 0x52, 0xdb, 0x39, 0xe4, 0xe2, 0xc6, 0xdd,
	0x4d, 0xe2, 0x5f, 0xf7, 0x26, 0xee, 0xac, 0x4a, 0x42, 0x6e, 0xde, 0x66, 0x1e, 0x6f, 0xde, 0xc7,
	0x2e, 0xa4, 0x86, 0xba, 0x9a, 0xfa, 0x17, 0x4b, 0xa6, 0xab, 0xf7, 0x54, 0x34, 0x46, 0x3b, 0x8d,
	0x09, 0xa3, 0x01, 0x94, 0x25, 0xc4, 0xca, 0x03, 0x88, 0x70, 0xbc, 0xa5, 0x0f, 0xb7, 0x14, 0x2b,
	0x91, 0x9f, 0x28, 0x3f, 0x7f, 0x63, 0x95, 0xdb, 0x99, 0xe5, 0xd1, 0x4a, 0xe4, 0x91, 0xf2, 0xb3,
	0xbc, 0x83, 0xd3, 0x27, 0xea, 0xf9, 0x48, 0xd1, 0x7b, 0x4b, 0xd6, 0xe1, 0x0d, 0xc4, 0x2c, 0xeb,
	0xaf, 0x17, 0xeb, 0xb3, 0x62, 0xe0, 0x52, 0x04, 0x76, 0x20, 0xc9, 0x2b, 0x98, 0x33, 0xb2, 0x39,
	0x60, 0x06, 0x73, 0x13, 0xe6, 0x60, 0xfd, 0xbb, 0xcb, 0x6b, 0x48, 0xef, 0xc9, 0x6d, 0x75, 0xc3,
	0x6c, 0xfb, 0x63, 0x97, 0x42, 0xb4, 0xd7, 0xed, 0x1b, 0x67, 0x8d, 0x14, 0x2f, 0xeb, 0x4f, 0x01,
	0x09, 0x13, 0x2b, 0xb6, 0xc5, 0x07, 0x80, 0x6a, 0xd7, 0x51, 0x28, 0x78, 0x31, 0x0a, 0x35, 0x6e,
	0x91, 0x9d, 0x4f, 0xa6, 0xde, 0x1c, 0xe4, 0x0c, 0x1f, 0x61, 0xf1, 0xa7, 0x64, 0xff, 0x23, 0x95,
	0x8b, 0x52, 0xe0, 0x33, 0x24, 0x83, 0x62, 0x78, 0x39, 0xe2, 0x4f, 0xd5, 0xce, 0xa6, 0x5f, 0x55,
	0xce, 0x4a, 0xf1, 0x1a, 0xfb, 0xcf, 0xbd, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x8a, 0x20,
	0xdd, 0xf4, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReviewServiceClient is the client API for ReviewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReviewServiceClient interface {
	SaveReview(ctx context.Context, in *NewReviewRequest, opts ...grpc.CallOption) (*ReviewId, error)
	SaveReviews(ctx context.Context, opts ...grpc.CallOption) (ReviewService_SaveReviewsClient, error)
	GetTopReviews(ctx context.Context, in *GetTopReviewsRequest, opts ...grpc.CallOption) (ReviewService_GetTopReviewsClient, error)
}

type reviewServiceClient struct {
	cc *grpc.ClientConn
}

func NewReviewServiceClient(cc *grpc.ClientConn) ReviewServiceClient {
	return &reviewServiceClient{cc}
}

func (c *reviewServiceClient) SaveReview(ctx context.Context, in *NewReviewRequest, opts ...grpc.CallOption) (*ReviewId, error) {
	out := new(ReviewId)
	err := c.cc.Invoke(ctx, "/reviewservice.ReviewService/SaveReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) SaveReviews(ctx context.Context, opts ...grpc.CallOption) (ReviewService_SaveReviewsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ReviewService_serviceDesc.Streams[0], "/reviewservice.ReviewService/SaveReviews", opts...)
	if err != nil {
		return nil, err
	}
	x := &reviewServiceSaveReviewsClient{stream}
	return x, nil
}

type ReviewService_SaveReviewsClient interface {
	Send(*NewReviewRequest) error
	Recv() (*ReviewId, error)
	grpc.ClientStream
}

type reviewServiceSaveReviewsClient struct {
	grpc.ClientStream
}

func (x *reviewServiceSaveReviewsClient) Send(m *NewReviewRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *reviewServiceSaveReviewsClient) Recv() (*ReviewId, error) {
	m := new(ReviewId)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *reviewServiceClient) GetTopReviews(ctx context.Context, in *GetTopReviewsRequest, opts ...grpc.CallOption) (ReviewService_GetTopReviewsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ReviewService_serviceDesc.Streams[1], "/reviewservice.ReviewService/GetTopReviews", opts...)
	if err != nil {
		return nil, err
	}
	x := &reviewServiceGetTopReviewsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ReviewService_GetTopReviewsClient interface {
	Recv() (*Review, error)
	grpc.ClientStream
}

type reviewServiceGetTopReviewsClient struct {
	grpc.ClientStream
}

func (x *reviewServiceGetTopReviewsClient) Recv() (*Review, error) {
	m := new(Review)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ReviewServiceServer is the server API for ReviewService service.
type ReviewServiceServer interface {
	SaveReview(context.Context, *NewReviewRequest) (*ReviewId, error)
	SaveReviews(ReviewService_SaveReviewsServer) error
	GetTopReviews(*GetTopReviewsRequest, ReviewService_GetTopReviewsServer) error
}

func RegisterReviewServiceServer(s *grpc.Server, srv ReviewServiceServer) {
	s.RegisterService(&_ReviewService_serviceDesc, srv)
}

func _ReviewService_SaveReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).SaveReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reviewservice.ReviewService/SaveReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).SaveReview(ctx, req.(*NewReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_SaveReviews_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ReviewServiceServer).SaveReviews(&reviewServiceSaveReviewsServer{stream})
}

type ReviewService_SaveReviewsServer interface {
	Send(*ReviewId) error
	Recv() (*NewReviewRequest, error)
	grpc.ServerStream
}

type reviewServiceSaveReviewsServer struct {
	grpc.ServerStream
}

func (x *reviewServiceSaveReviewsServer) Send(m *ReviewId) error {
	return x.ServerStream.SendMsg(m)
}

func (x *reviewServiceSaveReviewsServer) Recv() (*NewReviewRequest, error) {
	m := new(NewReviewRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ReviewService_GetTopReviews_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetTopReviewsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReviewServiceServer).GetTopReviews(m, &reviewServiceGetTopReviewsServer{stream})
}

type ReviewService_GetTopReviewsServer interface {
	Send(*Review) error
	grpc.ServerStream
}

type reviewServiceGetTopReviewsServer struct {
	grpc.ServerStream
}

func (x *reviewServiceGetTopReviewsServer) Send(m *Review) error {
	return x.ServerStream.SendMsg(m)
}

var _ReviewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "reviewservice.ReviewService",
	HandlerType: (*ReviewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveReview",
			Handler:    _ReviewService_SaveReview_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SaveReviews",
			Handler:       _ReviewService_SaveReviews_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetTopReviews",
			Handler:       _ReviewService_GetTopReviews_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "review_service.proto",
}
