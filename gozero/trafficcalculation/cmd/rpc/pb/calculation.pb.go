// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: calculation.proto

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

type Box struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LatMin float32 `protobuf:"fixed32,1,opt,name=lat_min,json=latMin,proto3" json:"lat_min,omitempty"` // 区域位置边界
	LngMin float32 `protobuf:"fixed32,2,opt,name=lng_min,json=lngMin,proto3" json:"lng_min,omitempty"` // 区域位置边界
	LatMax float32 `protobuf:"fixed32,3,opt,name=lat_max,json=latMax,proto3" json:"lat_max,omitempty"` // 区域位置边界
	LngMax float32 `protobuf:"fixed32,4,opt,name=lng_max,json=lngMax,proto3" json:"lng_max,omitempty"` // 区域位置边界
}

func (x *Box) Reset() {
	*x = Box{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Box) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Box) ProtoMessage() {}

func (x *Box) ProtoReflect() protoreflect.Message {
	mi := &file_calculation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Box.ProtoReflect.Descriptor instead.
func (*Box) Descriptor() ([]byte, []int) {
	return file_calculation_proto_rawDescGZIP(), []int{0}
}

func (x *Box) GetLatMin() float32 {
	if x != nil {
		return x.LatMin
	}
	return 0
}

func (x *Box) GetLngMin() float32 {
	if x != nil {
		return x.LngMin
	}
	return 0
}

func (x *Box) GetLatMax() float32 {
	if x != nil {
		return x.LatMax
	}
	return 0
}

func (x *Box) GetLngMax() float32 {
	if x != nil {
		return x.LngMax
	}
	return 0
}

type QueryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SegId  []int64 `protobuf:"varint,1,rep,packed,name=seg_id,json=segId,proto3" json:"seg_id,omitempty"`
	LaneId []int64 `protobuf:"varint,2,rep,packed,name=lane_id,json=laneId,proto3" json:"lane_id,omitempty"`
	Area   *Box    `protobuf:"bytes,3,opt,name=area,proto3" json:"area,omitempty"`
}

func (x *QueryReq) Reset() {
	*x = QueryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryReq) ProtoMessage() {}

func (x *QueryReq) ProtoReflect() protoreflect.Message {
	mi := &file_calculation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryReq.ProtoReflect.Descriptor instead.
func (*QueryReq) Descriptor() ([]byte, []int) {
	return file_calculation_proto_rawDescGZIP(), []int{1}
}

func (x *QueryReq) GetSegId() []int64 {
	if x != nil {
		return x.SegId
	}
	return nil
}

func (x *QueryReq) GetLaneId() []int64 {
	if x != nil {
		return x.LaneId
	}
	return nil
}

func (x *QueryReq) GetArea() *Box {
	if x != nil {
		return x.Area
	}
	return nil
}

type LaneStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LaneId int64  `protobuf:"varint,1,opt,name=lane_id,json=laneId,proto3" json:"lane_id,omitempty"`
	Level  uint32 `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *LaneStatus) Reset() {
	*x = LaneStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LaneStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LaneStatus) ProtoMessage() {}

func (x *LaneStatus) ProtoReflect() protoreflect.Message {
	mi := &file_calculation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LaneStatus.ProtoReflect.Descriptor instead.
func (*LaneStatus) Descriptor() ([]byte, []int) {
	return file_calculation_proto_rawDescGZIP(), []int{2}
}

func (x *LaneStatus) GetLaneId() int64 {
	if x != nil {
		return x.LaneId
	}
	return 0
}

func (x *LaneStatus) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

type TrafficStatusResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      int64         `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message   string        `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp int64         `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Traffics  []*LaneStatus `protobuf:"bytes,4,rep,name=traffics,proto3" json:"traffics,omitempty"`
}

func (x *TrafficStatusResp) Reset() {
	*x = TrafficStatusResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrafficStatusResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrafficStatusResp) ProtoMessage() {}

func (x *TrafficStatusResp) ProtoReflect() protoreflect.Message {
	mi := &file_calculation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrafficStatusResp.ProtoReflect.Descriptor instead.
func (*TrafficStatusResp) Descriptor() ([]byte, []int) {
	return file_calculation_proto_rawDescGZIP(), []int{3}
}

func (x *TrafficStatusResp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *TrafficStatusResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *TrafficStatusResp) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *TrafficStatusResp) GetTraffics() []*LaneStatus {
	if x != nil {
		return x.Traffics
	}
	return nil
}

var File_calculation_proto protoreflect.FileDescriptor

var file_calculation_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x69, 0x0a, 0x03, 0x42, 0x6f, 0x78, 0x12, 0x17,
	0x0a, 0x07, 0x6c, 0x61, 0x74, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x06, 0x6c, 0x61, 0x74, 0x4d, 0x69, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x6e, 0x67, 0x5f, 0x6d,
	0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x6c, 0x6e, 0x67, 0x4d, 0x69, 0x6e,
	0x12, 0x17, 0x0a, 0x07, 0x6c, 0x61, 0x74, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x06, 0x6c, 0x61, 0x74, 0x4d, 0x61, 0x78, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x6e, 0x67,
	0x5f, 0x6d, 0x61, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x6c, 0x6e, 0x67, 0x4d,
	0x61, 0x78, 0x22, 0x57, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x12, 0x15,
	0x0a, 0x06, 0x73, 0x65, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05,
	0x73, 0x65, 0x67, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x61, 0x6e, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x61, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70,
	0x62, 0x2e, 0x42, 0x6f, 0x78, 0x52, 0x04, 0x61, 0x72, 0x65, 0x61, 0x22, 0x3b, 0x0a, 0x0a, 0x4c,
	0x61, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x61, 0x6e,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x61, 0x6e, 0x65,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x8b, 0x01, 0x0a, 0x11, 0x54, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2a, 0x0a, 0x08, 0x74, 0x72,
	0x61, 0x66, 0x66, 0x69, 0x63, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70,
	0x62, 0x2e, 0x4c, 0x61, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x08, 0x74, 0x72,
	0x61, 0x66, 0x66, 0x69, 0x63, 0x73, 0x32, 0x50, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x13,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x12, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calculation_proto_rawDescOnce sync.Once
	file_calculation_proto_rawDescData = file_calculation_proto_rawDesc
)

func file_calculation_proto_rawDescGZIP() []byte {
	file_calculation_proto_rawDescOnce.Do(func() {
		file_calculation_proto_rawDescData = protoimpl.X.CompressGZIP(file_calculation_proto_rawDescData)
	})
	return file_calculation_proto_rawDescData
}

var file_calculation_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_calculation_proto_goTypes = []interface{}{
	(*Box)(nil),               // 0: pb.Box
	(*QueryReq)(nil),          // 1: pb.QueryReq
	(*LaneStatus)(nil),        // 2: pb.LaneStatus
	(*TrafficStatusResp)(nil), // 3: pb.TrafficStatusResp
}
var file_calculation_proto_depIdxs = []int32{
	0, // 0: pb.QueryReq.area:type_name -> pb.Box
	2, // 1: pb.TrafficStatusResp.traffics:type_name -> pb.LaneStatus
	1, // 2: pb.TrafficCalculation.QueryCurrentTraffic:input_type -> pb.QueryReq
	3, // 3: pb.TrafficCalculation.QueryCurrentTraffic:output_type -> pb.TrafficStatusResp
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_calculation_proto_init() }
func file_calculation_proto_init() {
	if File_calculation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calculation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Box); i {
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
		file_calculation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryReq); i {
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
		file_calculation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LaneStatus); i {
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
		file_calculation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrafficStatusResp); i {
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
			RawDescriptor: file_calculation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculation_proto_goTypes,
		DependencyIndexes: file_calculation_proto_depIdxs,
		MessageInfos:      file_calculation_proto_msgTypes,
	}.Build()
	File_calculation_proto = out.File
	file_calculation_proto_rawDesc = nil
	file_calculation_proto_goTypes = nil
	file_calculation_proto_depIdxs = nil
}