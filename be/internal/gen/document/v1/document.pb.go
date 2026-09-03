package documentv1

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Document struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content       string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	OwnerId       string                 `protobuf:"bytes,4,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Document) Reset()         { *x = Document{} }
func (x *Document) String() string { return protoimpl.X.MessageStringOf(x) }
func (*Document) ProtoMessage()    {}
func (x *Document) ProtoReflect() protoreflect.Message {
	mi := &file_document_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*Document) Descriptor() ([]byte, []int) { return file_document_proto_rawDescGZIP(), []int{0} }
func (x *Document) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}
func (x *Document) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}
func (x *Document) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}
func (x *Document) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}
func (x *Document) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}
func (x *Document) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreateDocumentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateDocumentRequest) Reset()         { *x = CreateDocumentRequest{} }
func (x *CreateDocumentRequest) String() string { return protoimpl.X.MessageStringOf(x) }
func (*CreateDocumentRequest) ProtoMessage()    {}
func (x *CreateDocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_document_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*CreateDocumentRequest) Descriptor() ([]byte, []int) {
	return file_document_proto_rawDescGZIP(), []int{1}
}
func (x *CreateDocumentRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type CreateDocumentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Document      *Document              `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateDocumentResponse) Reset()         { *x = CreateDocumentResponse{} }
func (x *CreateDocumentResponse) String() string { return protoimpl.X.MessageStringOf(x) }
func (*CreateDocumentResponse) ProtoMessage()    {}
func (x *CreateDocumentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_document_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*CreateDocumentResponse) Descriptor() ([]byte, []int) {
	return file_document_proto_rawDescGZIP(), []int{2}
}
func (x *CreateDocumentResponse) GetDocument() *Document {
	if x != nil {
		return x.Document
	}
	return nil
}

type GetDocumentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDocumentRequest) Reset()         { *x = GetDocumentRequest{} }
func (x *GetDocumentRequest) String() string { return protoimpl.X.MessageStringOf(x) }
func (*GetDocumentRequest) ProtoMessage()    {}
func (x *GetDocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_document_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*GetDocumentRequest) Descriptor() ([]byte, []int) {
	return file_document_proto_rawDescGZIP(), []int{3}
}
func (x *GetDocumentRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetDocumentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Document      *Document              `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDocumentResponse) Reset()         { *x = GetDocumentResponse{} }
func (x *GetDocumentResponse) String() string { return protoimpl.X.MessageStringOf(x) }
func (*GetDocumentResponse) ProtoMessage()    {}
func (x *GetDocumentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_document_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*GetDocumentResponse) Descriptor() ([]byte, []int) {
	return file_document_proto_rawDescGZIP(), []int{4}
}
func (x *GetDocumentResponse) GetDocument() *Document {
	if x != nil {
		return x.Document
	}
	return nil
}

type ListDocumentsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListDocumentsRequest) Reset()         { *x = ListDocumentsRequest{} }
func (x *ListDocumentsRequest) String() string { return protoimpl.X.MessageStringOf(x) }
func (*ListDocumentsRequest) ProtoMessage()    {}
func (x *ListDocumentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_document_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*ListDocumentsRequest) Descriptor() ([]byte, []int) {
	return file_document_proto_rawDescGZIP(), []int{5}
}

type ListDocumentsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Documents     []*Document            `protobuf:"bytes,1,rep,name=documents,proto3" json:"documents,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListDocumentsResponse) Reset()         { *x = ListDocumentsResponse{} }
func (x *ListDocumentsResponse) String() string { return protoimpl.X.MessageStringOf(x) }
func (*ListDocumentsResponse) ProtoMessage()    {}
func (x *ListDocumentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_document_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*ListDocumentsResponse) Descriptor() ([]byte, []int) {
	return file_document_proto_rawDescGZIP(), []int{6}
}
func (x *ListDocumentsResponse) GetDocuments() []*Document {
	if x != nil {
		return x.Documents
	}
	return nil
}

type UpdateDocumentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content       string                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateDocumentRequest) Reset()         { *x = UpdateDocumentRequest{} }
func (x *UpdateDocumentRequest) String() string { return protoimpl.X.MessageStringOf(x) }
func (*UpdateDocumentRequest) ProtoMessage()    {}
func (x *UpdateDocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_document_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*UpdateDocumentRequest) Descriptor() ([]byte, []int) {
	return file_document_proto_rawDescGZIP(), []int{7}
}
func (x *UpdateDocumentRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}
func (x *UpdateDocumentRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type UpdateDocumentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Document      *Document              `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateDocumentResponse) Reset()         { *x = UpdateDocumentResponse{} }
func (x *UpdateDocumentResponse) String() string { return protoimpl.X.MessageStringOf(x) }
func (*UpdateDocumentResponse) ProtoMessage()    {}
func (x *UpdateDocumentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_document_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*UpdateDocumentResponse) Descriptor() ([]byte, []int) {
	return file_document_proto_rawDescGZIP(), []int{8}
}
func (x *UpdateDocumentResponse) GetDocument() *Document {
	if x != nil {
		return x.Document
	}
	return nil
}

type DeleteDocumentRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteDocumentRequest) Reset()         {}
func (x *DeleteDocumentRequest) String() string { return x.Id }
func (*DeleteDocumentRequest) ProtoMessage()    {}
func (x *DeleteDocumentRequest) ProtoReflect() protoreflect.Message {
	return protoimpl.X.MessageStateOf(protoimpl.Pointer(x)).LoadMessageInfo().MessageOf(x)
}
func (x *DeleteDocumentRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteDocumentResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteDocumentResponse) Reset()         {}
func (x *DeleteDocumentResponse) String() string { return x.Message }
func (*DeleteDocumentResponse) ProtoMessage()    {}
func (x *DeleteDocumentResponse) ProtoReflect() protoreflect.Message {
	return protoimpl.X.MessageStateOf(protoimpl.Pointer(x)).LoadMessageInfo().MessageOf(x)
}
func (x *DeleteDocumentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_document_proto protoreflect.FileDescriptor

const file_document_proto_rawDesc = "" +
	"\n" +
	"\x0edocument.proto\x12\vdocument.v1\x1a\x1fgoogle/protobuf/timestamp.proto\"\xdb\x01\n" +
	"\bDocument\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12\x18\n" +
	"\acontent\x18\x03 \x01(\tR\acontent\x12\x19\n" +
	"\bowner_id\x18\x04 \x01(\tR\aownerId\x129\n" +
	"\n" +
	"created_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\"-\n" +
	"\x15CreateDocumentRequest\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\"K\n" +
	"\x16CreateDocumentResponse\x121\n" +
	"\bdocument\x18\x01 \x01(\v2\x15.document.v1.DocumentR\bdocument\"\x24\n" +
	"\x12GetDocumentRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"H\n" +
	"\x13GetDocumentResponse\x121\n" +
	"\bdocument\x18\x01 \x01(\v2\x15.document.v1.DocumentR\bdocument\"\x16\n" +
	"\x14ListDocumentsRequest\"L\n" +
	"\x15ListDocumentsResponse\x123\n" +
	"\tdocuments\x18\x01 \x03(\v2\x15.document.v1.DocumentR\tdocuments\"A\n" +
	"\x15UpdateDocumentRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x18\n" +
	"\acontent\x18\x02 \x01(\tR\acontent\"K\n" +
	"\x16UpdateDocumentResponse\x121\n" +
	"\bdocument\x18\x01 \x01(\v2\x15.document.v1.DocumentR\bdocument2\xf1\x02\n" +
	"\x0fDocumentService\x12Y\n" +
	"\x0eCreateDocument\x12\".document.v1.CreateDocumentRequest\x1a#.document.v1.CreateDocumentResponse\x12P\n" +
	"\vGetDocument\x12\x1f.document.v1.GetDocumentRequest\x1a .document.v1.GetDocumentResponse\x12V\n" +
	"\rListDocuments\x12!.document.v1.ListDocumentsRequest\x1a\".document.v1.ListDocumentsResponse\x12Y\n" +
	"\x0eUpdateDocument\x12\".document.v1.UpdateDocumentRequest\x1a#.document.v1.UpdateDocumentResponseBNZLgithub.com/Khansa01/collaboration-app/be/internal/gen/document/v1;documentv1b\x06proto3"

var (
	file_document_proto_rawDescOnce sync.Once
	file_document_proto_rawDescData []byte
)

func file_document_proto_rawDescGZIP() []byte {
	file_document_proto_rawDescOnce.Do(func() {
		file_document_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_document_proto_rawDesc), len(file_document_proto_rawDesc)))
	})
	return file_document_proto_rawDescData
}

var file_document_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_document_proto_goTypes = []any{
	(*Document)(nil),               // 0: document.v1.Document
	(*CreateDocumentRequest)(nil),  // 1: document.v1.CreateDocumentRequest
	(*CreateDocumentResponse)(nil), // 2: document.v1.CreateDocumentResponse
	(*GetDocumentRequest)(nil),     // 3: document.v1.GetDocumentRequest
	(*GetDocumentResponse)(nil),    // 4: document.v1.GetDocumentResponse
	(*ListDocumentsRequest)(nil),   // 5: document.v1.ListDocumentsRequest
	(*ListDocumentsResponse)(nil),  // 6: document.v1.ListDocumentsResponse
	(*UpdateDocumentRequest)(nil),  // 7: document.v1.UpdateDocumentRequest
	(*UpdateDocumentResponse)(nil), // 8: document.v1.UpdateDocumentResponse
	(*timestamppb.Timestamp)(nil),  // 9: google.protobuf.Timestamp
}
var file_document_proto_depIdxs = []int32{
	9, 9, 0, 0, 0, 0, 1, 3, 5, 7, 2, 4, 6, 8,
	10, 6, 6, 6, 0,
}

func init() { file_document_proto_init() }
func file_document_proto_init() {
	if File_document_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_document_proto_rawDesc), len(file_document_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_document_proto_goTypes,
		DependencyIndexes: file_document_proto_depIdxs,
		MessageInfos:      file_document_proto_msgTypes,
	}.Build()
	File_document_proto = out.File
	file_document_proto_goTypes = nil
	file_document_proto_depIdxs = nil
}
