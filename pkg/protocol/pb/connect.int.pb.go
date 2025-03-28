// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v5.28.3
// source: connect.int.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DeliverMessageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId int64    `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"` // 设备id
	Message  *Message `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`                    // 数据
}

func (x *DeliverMessageReq) Reset() {
	*x = DeliverMessageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_int_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliverMessageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliverMessageReq) ProtoMessage() {}

func (x *DeliverMessageReq) ProtoReflect() protoreflect.Message {
	mi := &file_connect_int_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliverMessageReq.ProtoReflect.Descriptor instead.
func (*DeliverMessageReq) Descriptor() ([]byte, []int) {
	return file_connect_int_proto_rawDescGZIP(), []int{0}
}

func (x *DeliverMessageReq) GetDeviceId() int64 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *DeliverMessageReq) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

// 房间推送
type PushRoomMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId  int64    `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"` // 设备id
	Message *Message `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`              // 数据
}

func (x *PushRoomMsg) Reset() {
	*x = PushRoomMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_int_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushRoomMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushRoomMsg) ProtoMessage() {}

func (x *PushRoomMsg) ProtoReflect() protoreflect.Message {
	mi := &file_connect_int_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushRoomMsg.ProtoReflect.Descriptor instead.
func (*PushRoomMsg) Descriptor() ([]byte, []int) {
	return file_connect_int_proto_rawDescGZIP(), []int{1}
}

func (x *PushRoomMsg) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *PushRoomMsg) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

// 房间推送
type PushAllMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *Message `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // 数据
}

func (x *PushAllMsg) Reset() {
	*x = PushAllMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_int_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushAllMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushAllMsg) ProtoMessage() {}

func (x *PushAllMsg) ProtoReflect() protoreflect.Message {
	mi := &file_connect_int_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushAllMsg.ProtoReflect.Descriptor instead.
func (*PushAllMsg) Descriptor() ([]byte, []int) {
	return file_connect_int_proto_rawDescGZIP(), []int{2}
}

func (x *PushAllMsg) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

var File_connect_int_proto protoreflect.FileDescriptor

var file_connect_int_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x65, 0x78, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x4d,
	0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x52, 0x6f, 0x6f, 0x6d, 0x4d, 0x73, 0x67, 0x12, 0x17, 0x0a,
	0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x33, 0x0a,
	0x0a, 0x50, 0x75, 0x73, 0x68, 0x41, 0x6c, 0x6c, 0x4d, 0x73, 0x67, 0x12, 0x25, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70,
	0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0x4d, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x74,
	0x12, 0x3f, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x6f, 0x2d, 0x69, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_connect_int_proto_rawDescOnce sync.Once
	file_connect_int_proto_rawDescData = file_connect_int_proto_rawDesc
)

func file_connect_int_proto_rawDescGZIP() []byte {
	file_connect_int_proto_rawDescOnce.Do(func() {
		file_connect_int_proto_rawDescData = protoimpl.X.CompressGZIP(file_connect_int_proto_rawDescData)
	})
	return file_connect_int_proto_rawDescData
}

var file_connect_int_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_connect_int_proto_goTypes = []interface{}{
	(*DeliverMessageReq)(nil), // 0: pb.DeliverMessageReq
	(*PushRoomMsg)(nil),       // 1: pb.PushRoomMsg
	(*PushAllMsg)(nil),        // 2: pb.PushAllMsg
	(*Message)(nil),           // 3: pb.Message
	(*emptypb.Empty)(nil),     // 4: google.protobuf.Empty
}
var file_connect_int_proto_depIdxs = []int32{
	3, // 0: pb.DeliverMessageReq.message:type_name -> pb.Message
	3, // 1: pb.PushRoomMsg.message:type_name -> pb.Message
	3, // 2: pb.PushAllMsg.message:type_name -> pb.Message
	0, // 3: pb.ConnectInt.DeliverMessage:input_type -> pb.DeliverMessageReq
	4, // 4: pb.ConnectInt.DeliverMessage:output_type -> google.protobuf.Empty
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_connect_int_proto_init() }
func file_connect_int_proto_init() {
	if File_connect_int_proto != nil {
		return
	}
	file_basic_ext_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_connect_int_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliverMessageReq); i {
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
		file_connect_int_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushRoomMsg); i {
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
		file_connect_int_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushAllMsg); i {
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
			RawDescriptor: file_connect_int_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_connect_int_proto_goTypes,
		DependencyIndexes: file_connect_int_proto_depIdxs,
		MessageInfos:      file_connect_int_proto_msgTypes,
	}.Build()
	File_connect_int_proto = out.File
	file_connect_int_proto_rawDesc = nil
	file_connect_int_proto_goTypes = nil
	file_connect_int_proto_depIdxs = nil
}
