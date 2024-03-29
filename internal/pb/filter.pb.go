// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: filter.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PostConnectionFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PostConnectionFilter) Reset() {
	*x = PostConnectionFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostConnectionFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostConnectionFilter) ProtoMessage() {}

func (x *PostConnectionFilter) ProtoReflect() protoreflect.Message {
	mi := &file_filter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostConnectionFilter.ProtoReflect.Descriptor instead.
func (*PostConnectionFilter) Descriptor() ([]byte, []int) {
	return file_filter_proto_rawDescGZIP(), []int{0}
}

type UserConnectionFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserConnectionFilter) Reset() {
	*x = UserConnectionFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserConnectionFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserConnectionFilter) ProtoMessage() {}

func (x *UserConnectionFilter) ProtoReflect() protoreflect.Message {
	mi := &file_filter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserConnectionFilter.ProtoReflect.Descriptor instead.
func (*UserConnectionFilter) Descriptor() ([]byte, []int) {
	return file_filter_proto_rawDescGZIP(), []int{1}
}

type CommentConnectionFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CommentConnectionFilter) Reset() {
	*x = CommentConnectionFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentConnectionFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentConnectionFilter) ProtoMessage() {}

func (x *CommentConnectionFilter) ProtoReflect() protoreflect.Message {
	mi := &file_filter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentConnectionFilter.ProtoReflect.Descriptor instead.
func (*CommentConnectionFilter) Descriptor() ([]byte, []int) {
	return file_filter_proto_rawDescGZIP(), []int{2}
}

type SearchResultsPeopleFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SearchResultsPeopleFilter) Reset() {
	*x = SearchResultsPeopleFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResultsPeopleFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResultsPeopleFilter) ProtoMessage() {}

func (x *SearchResultsPeopleFilter) ProtoReflect() protoreflect.Message {
	mi := &file_filter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResultsPeopleFilter.ProtoReflect.Descriptor instead.
func (*SearchResultsPeopleFilter) Descriptor() ([]byte, []int) {
	return file_filter_proto_rawDescGZIP(), []int{3}
}

type SearchResultPostFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SearchResultPostFilter) Reset() {
	*x = SearchResultPostFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResultPostFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResultPostFilter) ProtoMessage() {}

func (x *SearchResultPostFilter) ProtoReflect() protoreflect.Message {
	mi := &file_filter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResultPostFilter.ProtoReflect.Descriptor instead.
func (*SearchResultPostFilter) Descriptor() ([]byte, []int) {
	return file_filter_proto_rawDescGZIP(), []int{4}
}

type SearchResultsCommentFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SearchResultsCommentFilter) Reset() {
	*x = SearchResultsCommentFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResultsCommentFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResultsCommentFilter) ProtoMessage() {}

func (x *SearchResultsCommentFilter) ProtoReflect() protoreflect.Message {
	mi := &file_filter_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResultsCommentFilter.ProtoReflect.Descriptor instead.
func (*SearchResultsCommentFilter) Descriptor() ([]byte, []int) {
	return file_filter_proto_rawDescGZIP(), []int{5}
}

var File_filter_proto protoreflect.FileDescriptor

var file_filter_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x16, 0x0a, 0x14, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x16,
	0x0a, 0x14, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x19, 0x0a, 0x17, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x22, 0x1b, 0x0a, 0x19, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x18,
	0x0a, 0x16, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x50, 0x6f,
	0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x1c, 0x0a, 0x1a, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x26, 0x48, 0x01, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x6d, 0x76, 0x37, 0x2f, 0x62, 0x6c, 0x6f,
	0x67, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_filter_proto_rawDescOnce sync.Once
	file_filter_proto_rawDescData = file_filter_proto_rawDesc
)

func file_filter_proto_rawDescGZIP() []byte {
	file_filter_proto_rawDescOnce.Do(func() {
		file_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_filter_proto_rawDescData)
	})
	return file_filter_proto_rawDescData
}

var file_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_filter_proto_goTypes = []interface{}{
	(*PostConnectionFilter)(nil),       // 0: filter.PostConnectionFilter
	(*UserConnectionFilter)(nil),       // 1: filter.UserConnectionFilter
	(*CommentConnectionFilter)(nil),    // 2: filter.CommentConnectionFilter
	(*SearchResultsPeopleFilter)(nil),  // 3: filter.SearchResultsPeopleFilter
	(*SearchResultPostFilter)(nil),     // 4: filter.SearchResultPostFilter
	(*SearchResultsCommentFilter)(nil), // 5: filter.SearchResultsCommentFilter
}
var file_filter_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_filter_proto_init() }
func file_filter_proto_init() {
	if File_filter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_filter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostConnectionFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_filter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserConnectionFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_filter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentConnectionFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_filter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResultsPeopleFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_filter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResultPostFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_filter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResultsCommentFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_filter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_filter_proto_goTypes,
		DependencyIndexes: file_filter_proto_depIdxs,
		MessageInfos:      file_filter_proto_msgTypes,
	}.Build()
	File_filter_proto = out.File
	file_filter_proto_rawDesc = nil
	file_filter_proto_goTypes = nil
	file_filter_proto_depIdxs = nil
}
