// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.2
// source: service.proto

package dbsdk

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DBTuple struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts   int64  `protobuf:"varint,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DBTuple) Reset() {
	*x = DBTuple{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBTuple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBTuple) ProtoMessage() {}

func (x *DBTuple) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBTuple.ProtoReflect.Descriptor instead.
func (*DBTuple) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *DBTuple) GetTs() int64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

func (x *DBTuple) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type TableStatTuple struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RowCount int64 `protobuf:"varint,1,opt,name=rowCount,proto3" json:"rowCount,omitempty"`
	OldestTS int64 `protobuf:"varint,2,opt,name=oldestTS,proto3" json:"oldestTS,omitempty"`
	NewestTS int64 `protobuf:"varint,3,opt,name=newestTS,proto3" json:"newestTS,omitempty"`
}

func (x *TableStatTuple) Reset() {
	*x = TableStatTuple{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableStatTuple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableStatTuple) ProtoMessage() {}

func (x *TableStatTuple) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableStatTuple.ProtoReflect.Descriptor instead.
func (*TableStatTuple) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *TableStatTuple) GetRowCount() int64 {
	if x != nil {
		return x.RowCount
	}
	return 0
}

func (x *TableStatTuple) GetOldestTS() int64 {
	if x != nil {
		return x.OldestTS
	}
	return 0
}

func (x *TableStatTuple) GetNewestTS() int64 {
	if x != nil {
		return x.NewestTS
	}
	return 0
}

type AppendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table string   `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Data  *DBTuple `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AppendRequest) Reset() {
	*x = AppendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendRequest) ProtoMessage() {}

func (x *AppendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendRequest.ProtoReflect.Descriptor instead.
func (*AppendRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *AppendRequest) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *AppendRequest) GetData() *DBTuple {
	if x != nil {
		return x.Data
	}
	return nil
}

type AppendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AppendResponse) Reset() {
	*x = AppendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendResponse) ProtoMessage() {}

func (x *AppendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendResponse.ProtoReflect.Descriptor instead.
func (*AppendResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

type QueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table string `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Start int64  `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	Stop  int64  `protobuf:"varint,3,opt,name=stop,proto3" json:"stop,omitempty"`
}

func (x *QueryRequest) Reset() {
	*x = QueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRequest) ProtoMessage() {}

func (x *QueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRequest.ProtoReflect.Descriptor instead.
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

func (x *QueryRequest) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *QueryRequest) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *QueryRequest) GetStop() int64 {
	if x != nil {
		return x.Stop
	}
	return 0
}

type QueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*DBTuple `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *QueryResponse) Reset() {
	*x = QueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResponse) ProtoMessage() {}

func (x *QueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResponse.ProtoReflect.Descriptor instead.
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{5}
}

func (x *QueryResponse) GetData() []*DBTuple {
	if x != nil {
		return x.Data
	}
	return nil
}

type StatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table string `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
}

func (x *StatsRequest) Reset() {
	*x = StatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsRequest) ProtoMessage() {}

func (x *StatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsRequest.ProtoReflect.Descriptor instead.
func (*StatsRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{6}
}

func (x *StatsRequest) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x2d, 0x0a, 0x07, 0x44, 0x42, 0x54, 0x75, 0x70, 0x6c, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x64, 0x0a, 0x0e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f, 0x77, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x6f, 0x77, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x65, 0x73, 0x74, 0x54, 0x53, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6f, 0x6c, 0x64, 0x65, 0x73, 0x74, 0x54, 0x53, 0x12, 0x1a,
	0x0a, 0x08, 0x6e, 0x65, 0x77, 0x65, 0x73, 0x74, 0x54, 0x53, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x6e, 0x65, 0x77, 0x65, 0x73, 0x74, 0x54, 0x53, 0x22, 0x48, 0x0a, 0x0d, 0x41, 0x70,
	0x70, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x21, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x44, 0x42, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x10, 0x0a, 0x0e, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4e, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x6f, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x73, 0x74, 0x6f, 0x70, 0x22, 0x32, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x44, 0x42, 0x54,
	0x75, 0x70, 0x6c, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x24, 0x0a, 0x0c, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x32, 0xab, 0x01, 0x0a, 0x09, 0x44, 0x42, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35,
	0x0a, 0x06, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x12, 0x13, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12,
	0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x05, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x12, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x22, 0x00, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x3b, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_service_proto_goTypes = []interface{}{
	(*DBTuple)(nil),        // 0: main.DBTuple
	(*TableStatTuple)(nil), // 1: main.TableStatTuple
	(*AppendRequest)(nil),  // 2: main.AppendRequest
	(*AppendResponse)(nil), // 3: main.AppendResponse
	(*QueryRequest)(nil),   // 4: main.QueryRequest
	(*QueryResponse)(nil),  // 5: main.QueryResponse
	(*StatsRequest)(nil),   // 6: main.StatsRequest
}
var file_service_proto_depIdxs = []int32{
	0, // 0: main.AppendRequest.data:type_name -> main.DBTuple
	0, // 1: main.QueryResponse.data:type_name -> main.DBTuple
	2, // 2: main.DBService.Append:input_type -> main.AppendRequest
	4, // 3: main.DBService.Query:input_type -> main.QueryRequest
	6, // 4: main.DBService.Stats:input_type -> main.StatsRequest
	3, // 5: main.DBService.Append:output_type -> main.AppendResponse
	5, // 6: main.DBService.Query:output_type -> main.QueryResponse
	1, // 7: main.DBService.Stats:output_type -> main.TableStatTuple
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBTuple); i {
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
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableStatTuple); i {
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
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendRequest); i {
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
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendResponse); i {
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
		file_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRequest); i {
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
		file_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryResponse); i {
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
		file_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsRequest); i {
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
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
