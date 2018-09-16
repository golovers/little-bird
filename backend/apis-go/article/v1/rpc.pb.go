// Code generated by protoc-gen-go. DO NOT EDIT.
// source: article/v1/rpc.proto

package littlebird_article_v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

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

type CreateRequest struct {
	Article              *Article `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7eb0a1d4c43f90a9, []int{0}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetArticle() *Article {
	if m != nil {
		return m.Article
	}
	return nil
}

type CreateResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7eb0a1d4c43f90a9, []int{1}
}
func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type TrendingRequest struct {
	Offset               int64    `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int64    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrendingRequest) Reset()         { *m = TrendingRequest{} }
func (m *TrendingRequest) String() string { return proto.CompactTextString(m) }
func (*TrendingRequest) ProtoMessage()    {}
func (*TrendingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7eb0a1d4c43f90a9, []int{2}
}
func (m *TrendingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrendingRequest.Unmarshal(m, b)
}
func (m *TrendingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrendingRequest.Marshal(b, m, deterministic)
}
func (m *TrendingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrendingRequest.Merge(m, src)
}
func (m *TrendingRequest) XXX_Size() int {
	return xxx_messageInfo_TrendingRequest.Size(m)
}
func (m *TrendingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TrendingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TrendingRequest proto.InternalMessageInfo

func (m *TrendingRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *TrendingRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type UpdateRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Article              *Article `protobuf:"bytes,2,opt,name=article,proto3" json:"article,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7eb0a1d4c43f90a9, []int{3}
}
func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateRequest) GetArticle() *Article {
	if m != nil {
		return m.Article
	}
	return nil
}

type ListCreatedByRequest struct {
	OwnerId              string   `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCreatedByRequest) Reset()         { *m = ListCreatedByRequest{} }
func (m *ListCreatedByRequest) String() string { return proto.CompactTextString(m) }
func (*ListCreatedByRequest) ProtoMessage()    {}
func (*ListCreatedByRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7eb0a1d4c43f90a9, []int{4}
}
func (m *ListCreatedByRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCreatedByRequest.Unmarshal(m, b)
}
func (m *ListCreatedByRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCreatedByRequest.Marshal(b, m, deterministic)
}
func (m *ListCreatedByRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCreatedByRequest.Merge(m, src)
}
func (m *ListCreatedByRequest) XXX_Size() int {
	return xxx_messageInfo_ListCreatedByRequest.Size(m)
}
func (m *ListCreatedByRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCreatedByRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCreatedByRequest proto.InternalMessageInfo

func (m *ListCreatedByRequest) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

type GetRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7eb0a1d4c43f90a9, []int{5}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7eb0a1d4c43f90a9, []int{6}
}
func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListRequest struct {
	Offset               int64    `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int64    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7eb0a1d4c43f90a9, []int{7}
}
func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ListResponse struct {
	Articles             []*Article `protobuf:"bytes,1,rep,name=articles,proto3" json:"articles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7eb0a1d4c43f90a9, []int{8}
}
func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetArticles() []*Article {
	if m != nil {
		return m.Articles
	}
	return nil
}

type Article struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	LastUpdate           string   `protobuf:"bytes,4,opt,name=last_update,json=lastUpdate,proto3" json:"last_update,omitempty"`
	CreatedBy            string   `protobuf:"bytes,5,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	CreatedById          string   `protobuf:"bytes,6,opt,name=created_by_id,json=createdById,proto3" json:"created_by_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Article) Reset()         { *m = Article{} }
func (m *Article) String() string { return proto.CompactTextString(m) }
func (*Article) ProtoMessage()    {}
func (*Article) Descriptor() ([]byte, []int) {
	return fileDescriptor_7eb0a1d4c43f90a9, []int{9}
}
func (m *Article) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Article.Unmarshal(m, b)
}
func (m *Article) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Article.Marshal(b, m, deterministic)
}
func (m *Article) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Article.Merge(m, src)
}
func (m *Article) XXX_Size() int {
	return xxx_messageInfo_Article.Size(m)
}
func (m *Article) XXX_DiscardUnknown() {
	xxx_messageInfo_Article.DiscardUnknown(m)
}

var xxx_messageInfo_Article proto.InternalMessageInfo

func (m *Article) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Article) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Article) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Article) GetLastUpdate() string {
	if m != nil {
		return m.LastUpdate
	}
	return ""
}

func (m *Article) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *Article) GetCreatedById() string {
	if m != nil {
		return m.CreatedById
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "littlebird.article.v1.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "littlebird.article.v1.CreateResponse")
	proto.RegisterType((*TrendingRequest)(nil), "littlebird.article.v1.TrendingRequest")
	proto.RegisterType((*UpdateRequest)(nil), "littlebird.article.v1.UpdateRequest")
	proto.RegisterType((*ListCreatedByRequest)(nil), "littlebird.article.v1.ListCreatedByRequest")
	proto.RegisterType((*GetRequest)(nil), "littlebird.article.v1.GetRequest")
	proto.RegisterType((*DeleteRequest)(nil), "littlebird.article.v1.DeleteRequest")
	proto.RegisterType((*ListRequest)(nil), "littlebird.article.v1.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "littlebird.article.v1.ListResponse")
	proto.RegisterType((*Article)(nil), "littlebird.article.v1.Article")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ArticleServiceClient is the client API for ArticleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArticleServiceClient interface {
	// List articles
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// ListCreatedBy list articles by a specific user
	ListCreatedBy(ctx context.Context, in *ListCreatedByRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// Delete delete an article
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Get get a specific article
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Article, error)
	// Update update a specific article
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Update update a specific article
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
}

type articleServiceClient struct {
	cc *grpc.ClientConn
}

func NewArticleServiceClient(cc *grpc.ClientConn) ArticleServiceClient {
	return &articleServiceClient{cc}
}

func (c *articleServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/littlebird.article.v1.ArticleService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) ListCreatedBy(ctx context.Context, in *ListCreatedByRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/littlebird.article.v1.ArticleService/ListCreatedBy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/littlebird.article.v1.ArticleService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Article, error) {
	out := new(Article)
	err := c.cc.Invoke(ctx, "/littlebird.article.v1.ArticleService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/littlebird.article.v1.ArticleService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/littlebird.article.v1.ArticleService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticleServiceServer is the server API for ArticleService service.
type ArticleServiceServer interface {
	// List articles
	List(context.Context, *ListRequest) (*ListResponse, error)
	// ListCreatedBy list articles by a specific user
	ListCreatedBy(context.Context, *ListCreatedByRequest) (*ListResponse, error)
	// Delete delete an article
	Delete(context.Context, *DeleteRequest) (*empty.Empty, error)
	// Get get a specific article
	Get(context.Context, *GetRequest) (*Article, error)
	// Update update a specific article
	Update(context.Context, *UpdateRequest) (*empty.Empty, error)
	// Update update a specific article
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
}

func RegisterArticleServiceServer(s *grpc.Server, srv ArticleServiceServer) {
	s.RegisterService(&_ArticleService_serviceDesc, srv)
}

func _ArticleService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/littlebird.article.v1.ArticleService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_ListCreatedBy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCreatedByRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).ListCreatedBy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/littlebird.article.v1.ArticleService/ListCreatedBy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).ListCreatedBy(ctx, req.(*ListCreatedByRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/littlebird.article.v1.ArticleService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/littlebird.article.v1.ArticleService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/littlebird.article.v1.ArticleService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/littlebird.article.v1.ArticleService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ArticleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "littlebird.article.v1.ArticleService",
	HandlerType: (*ArticleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _ArticleService_List_Handler,
		},
		{
			MethodName: "ListCreatedBy",
			Handler:    _ArticleService_ListCreatedBy_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ArticleService_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ArticleService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ArticleService_Update_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ArticleService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "article/v1/rpc.proto",
}

func init() { proto.RegisterFile("article/v1/rpc.proto", fileDescriptor_7eb0a1d4c43f90a9) }

var fileDescriptor_7eb0a1d4c43f90a9 = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x6d, 0x92, 0xc6, 0x69, 0x26, 0x38, 0x48, 0xab, 0x50, 0x19, 0x53, 0x68, 0x58, 0x40, 0x42,
	0x42, 0xb2, 0x95, 0x72, 0x41, 0x70, 0x40, 0x7c, 0xa9, 0xa4, 0xe2, 0x82, 0x01, 0x21, 0x4e, 0x91,
	0x63, 0x4f, 0xa2, 0x95, 0x5c, 0xaf, 0xb1, 0x27, 0x41, 0x39, 0xf2, 0x83, 0xf8, 0x8f, 0xc8, 0xbb,
	0xeb, 0x34, 0x29, 0x38, 0x2d, 0xdc, 0x3c, 0x3b, 0x6f, 0xe7, 0xbd, 0x79, 0x3b, 0x63, 0x18, 0x84,
	0x39, 0x89, 0x28, 0x41, 0x7f, 0x39, 0xf2, 0xf3, 0x2c, 0xf2, 0xb2, 0x5c, 0x92, 0x64, 0xb7, 0x12,
	0x41, 0x94, 0xe0, 0x54, 0xe4, 0xb1, 0x67, 0x00, 0xde, 0x72, 0xe4, 0x1e, 0xcd, 0xa5, 0x9c, 0x27,
	0xe8, 0x87, 0x99, 0xf0, 0xc3, 0x34, 0x95, 0x14, 0x92, 0x90, 0x69, 0xa1, 0x2f, 0xb9, 0x77, 0x4c,
	0x56, 0x45, 0xd3, 0xc5, 0xcc, 0xc7, 0xf3, 0x8c, 0x56, 0x3a, 0xc9, 0xc7, 0x60, 0xbf, 0xc9, 0x31,
	0x24, 0x0c, 0xf0, 0xfb, 0x02, 0x0b, 0x62, 0xcf, 0xa0, 0x63, 0x2a, 0x3b, 0x8d, 0x61, 0xe3, 0x71,
	0xef, 0xe4, 0x9e, 0xf7, 0x57, 0x52, 0xef, 0x95, 0xfe, 0x0c, 0x2a, 0x38, 0x1f, 0x42, 0xbf, 0x2a,
	0x55, 0x64, 0x32, 0x2d, 0x90, 0xf5, 0xa1, 0x29, 0x62, 0x55, 0xa6, 0x1b, 0x34, 0x45, 0xcc, 0x5f,
	0xc2, 0xcd, 0xcf, 0x39, 0xa6, 0xb1, 0x48, 0xe7, 0x15, 0xdd, 0x21, 0x58, 0x72, 0x36, 0x2b, 0x90,
	0x14, 0xac, 0x15, 0x98, 0x88, 0x0d, 0xa0, 0x9d, 0x88, 0x73, 0x41, 0x4e, 0x53, 0x1d, 0xeb, 0x80,
	0x7f, 0x03, 0xfb, 0x4b, 0x16, 0x6f, 0xa8, 0xbd, 0xc4, 0xb0, 0xa9, 0xbe, 0xf9, 0x6f, 0xea, 0x47,
	0x30, 0xf8, 0x20, 0x0a, 0xd2, 0x1d, 0xc4, 0xaf, 0x57, 0x15, 0xc3, 0x6d, 0x38, 0x90, 0x3f, 0x52,
	0xcc, 0x27, 0x6b, 0x9e, 0x8e, 0x8a, 0xc7, 0x31, 0x3f, 0x02, 0x38, 0x45, 0xaa, 0x91, 0xc2, 0x8f,
	0xc1, 0x7e, 0x8b, 0x09, 0xd6, 0x6a, 0xe5, 0x2f, 0xa0, 0x57, 0x32, 0xfe, 0x9f, 0x13, 0x67, 0x70,
	0x43, 0x5f, 0x36, 0x56, 0x3f, 0x87, 0x03, 0xd3, 0x49, 0xe1, 0x34, 0x86, 0xad, 0x6b, 0x74, 0xbe,
	0xc6, 0xf3, 0x5f, 0x0d, 0xe8, 0x98, 0xd3, 0x3f, 0x0c, 0x1d, 0x40, 0x9b, 0x04, 0x19, 0x3b, 0xbb,
	0x81, 0x0e, 0x98, 0x03, 0x9d, 0x48, 0xa6, 0x84, 0x29, 0x39, 0x2d, 0xed, 0x89, 0x09, 0xd9, 0x31,
	0xf4, 0x92, 0xb0, 0xa0, 0xc9, 0x42, 0x3d, 0x93, 0xb3, 0xaf, 0xb2, 0x50, 0x1e, 0xe9, 0x87, 0x63,
	0x77, 0x01, 0x22, 0xed, 0xf1, 0x64, 0xba, 0x72, 0xda, 0x2a, 0xdf, 0x8d, 0x2a, 0xd7, 0x19, 0x07,
	0xfb, 0x22, 0x5d, 0x7a, 0x6e, 0x29, 0x44, 0x6f, 0x8d, 0x18, 0xc7, 0x27, 0x3f, 0xf7, 0xa1, 0x6f,
	0xf4, 0x7e, 0xc2, 0x7c, 0x29, 0x22, 0x64, 0x1f, 0x61, 0xbf, 0xb4, 0x83, 0xf1, 0x9a, 0xa6, 0x37,
	0x8c, 0x76, 0x1f, 0xec, 0xc4, 0x68, 0x3f, 0xf9, 0x1e, 0x8b, 0xc0, 0xde, 0x1a, 0x08, 0xf6, 0x64,
	0xc7, 0xbd, 0xcb, 0x63, 0x73, 0x5d, 0x92, 0xf7, 0x60, 0xe9, 0x21, 0x61, 0x0f, 0x6b, 0x2e, 0x6c,
	0xcd, 0x90, 0x7b, 0xe8, 0xe9, 0x65, 0xf6, 0xaa, 0x65, 0xf6, 0xde, 0x95, 0xcb, 0xcc, 0xf7, 0xd8,
	0x19, 0xb4, 0x4e, 0x91, 0xd8, 0xfd, 0x9a, 0x32, 0x17, 0x83, 0xea, 0x5e, 0x31, 0x18, 0x5a, 0x95,
	0x79, 0xad, 0x3a, 0x55, 0x5b, 0x5b, 0xb8, 0x43, 0xd5, 0x57, 0xb0, 0xb4, 0x35, 0xb5, 0x95, 0xb6,
	0xfe, 0x3e, 0xee, 0xa3, 0x2b, 0x50, 0x95, 0x71, 0x53, 0x4b, 0x51, 0x3d, 0xfd, 0x1d, 0x00, 0x00,
	0xff, 0xff, 0x89, 0x4b, 0x17, 0xe9, 0x28, 0x05, 0x00, 0x00,
}
