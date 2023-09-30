// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.24.3
// source: fleetspeak/src/server/proto/fleetspeak_server/resource.proto

package fleetspeak_server

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represents client resource-usage data in the data-store.
// Next id: 15
type ClientResourceUsageRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the client service that resource usage is charged/attributed to
	// e.g 'system' for the system Fleetspeak service, or the name of a daemon
	// service as specified in its config.
	Scope            string                 `protobuf:"bytes,1,opt,name=scope,proto3" json:"scope,omitempty"`
	Pid              int64                  `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	ProcessStartTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=process_start_time,json=processStartTime,proto3" json:"process_start_time,omitempty"`
	// When the resource-usage metrics were measured on the client.
	ClientTimestamp *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=client_timestamp,json=clientTimestamp,proto3" json:"client_timestamp,omitempty"`
	// When the resource usage record was written to the data-store.
	ServerTimestamp *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=server_timestamp,json=serverTimestamp,proto3" json:"server_timestamp,omitempty"`
	// If true, indicates that the process has terminated, and that this is
	// the final resource-usage record for that process.
	ProcessTerminated bool `protobuf:"varint,12,opt,name=process_terminated,json=processTerminated,proto3" json:"process_terminated,omitempty"`
	// CPU-usage is in millis per second.
	MeanUserCpuRate       float32 `protobuf:"fixed32,6,opt,name=mean_user_cpu_rate,json=meanUserCpuRate,proto3" json:"mean_user_cpu_rate,omitempty"`
	MaxUserCpuRate        float32 `protobuf:"fixed32,7,opt,name=max_user_cpu_rate,json=maxUserCpuRate,proto3" json:"max_user_cpu_rate,omitempty"`
	MeanSystemCpuRate     float32 `protobuf:"fixed32,8,opt,name=mean_system_cpu_rate,json=meanSystemCpuRate,proto3" json:"mean_system_cpu_rate,omitempty"`
	MaxSystemCpuRate      float32 `protobuf:"fixed32,9,opt,name=max_system_cpu_rate,json=maxSystemCpuRate,proto3" json:"max_system_cpu_rate,omitempty"`
	MeanResidentMemoryMib int32   `protobuf:"varint,10,opt,name=mean_resident_memory_mib,json=meanResidentMemoryMib,proto3" json:"mean_resident_memory_mib,omitempty"`
	MaxResidentMemoryMib  int32   `protobuf:"varint,11,opt,name=max_resident_memory_mib,json=maxResidentMemoryMib,proto3" json:"max_resident_memory_mib,omitempty"`
	MeanNumFds            int32   `protobuf:"varint,13,opt,name=mean_num_fds,json=meanNumFds,proto3" json:"mean_num_fds,omitempty"`
	MaxNumFds             int32   `protobuf:"varint,14,opt,name=max_num_fds,json=maxNumFds,proto3" json:"max_num_fds,omitempty"`
}

func (x *ClientResourceUsageRecord) Reset() {
	*x = ClientResourceUsageRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_server_proto_fleetspeak_server_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientResourceUsageRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientResourceUsageRecord) ProtoMessage() {}

func (x *ClientResourceUsageRecord) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_server_proto_fleetspeak_server_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientResourceUsageRecord.ProtoReflect.Descriptor instead.
func (*ClientResourceUsageRecord) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_server_proto_fleetspeak_server_resource_proto_rawDescGZIP(), []int{0}
}

func (x *ClientResourceUsageRecord) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *ClientResourceUsageRecord) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *ClientResourceUsageRecord) GetProcessStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ProcessStartTime
	}
	return nil
}

func (x *ClientResourceUsageRecord) GetClientTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.ClientTimestamp
	}
	return nil
}

func (x *ClientResourceUsageRecord) GetServerTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.ServerTimestamp
	}
	return nil
}

func (x *ClientResourceUsageRecord) GetProcessTerminated() bool {
	if x != nil {
		return x.ProcessTerminated
	}
	return false
}

func (x *ClientResourceUsageRecord) GetMeanUserCpuRate() float32 {
	if x != nil {
		return x.MeanUserCpuRate
	}
	return 0
}

func (x *ClientResourceUsageRecord) GetMaxUserCpuRate() float32 {
	if x != nil {
		return x.MaxUserCpuRate
	}
	return 0
}

func (x *ClientResourceUsageRecord) GetMeanSystemCpuRate() float32 {
	if x != nil {
		return x.MeanSystemCpuRate
	}
	return 0
}

func (x *ClientResourceUsageRecord) GetMaxSystemCpuRate() float32 {
	if x != nil {
		return x.MaxSystemCpuRate
	}
	return 0
}

func (x *ClientResourceUsageRecord) GetMeanResidentMemoryMib() int32 {
	if x != nil {
		return x.MeanResidentMemoryMib
	}
	return 0
}

func (x *ClientResourceUsageRecord) GetMaxResidentMemoryMib() int32 {
	if x != nil {
		return x.MaxResidentMemoryMib
	}
	return 0
}

func (x *ClientResourceUsageRecord) GetMeanNumFds() int32 {
	if x != nil {
		return x.MeanNumFds
	}
	return 0
}

func (x *ClientResourceUsageRecord) GetMaxNumFds() int32 {
	if x != nil {
		return x.MaxNumFds
	}
	return 0
}

var File_fleetspeak_src_server_proto_fleetspeak_server_resource_proto protoreflect.FileDescriptor

var file_fleetspeak_src_server_proto_fleetspeak_server_resource_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2f, 0x73, 0x72, 0x63,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6c,
	0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb4, 0x05, 0x0a, 0x19, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x48, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x10, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x45, 0x0a, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x45, 0x0a, 0x10, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x2d, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x2b, 0x0a, 0x12, 0x6d, 0x65, 0x61, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x70, 0x75,
	0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x6d, 0x65, 0x61,
	0x6e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x70, 0x75, 0x52, 0x61, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x11,
	0x6d, 0x61, 0x78, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x70, 0x75, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x6d, 0x61, 0x78, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x70, 0x75, 0x52, 0x61, 0x74, 0x65, 0x12, 0x2f, 0x0a, 0x14, 0x6d, 0x65, 0x61, 0x6e, 0x5f,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x70, 0x75, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x11, 0x6d, 0x65, 0x61, 0x6e, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x43, 0x70, 0x75, 0x52, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x70, 0x75, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x10, 0x6d, 0x61, 0x78, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x43, 0x70, 0x75, 0x52, 0x61, 0x74, 0x65, 0x12, 0x37, 0x0a, 0x18, 0x6d, 0x65, 0x61, 0x6e, 0x5f,
	0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f,
	0x6d, 0x69, 0x62, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x6d, 0x65, 0x61, 0x6e, 0x52,
	0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4d, 0x69, 0x62,
	0x12, 0x35, 0x0a, 0x17, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6d, 0x69, 0x62, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x14, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x4d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x4d, 0x69, 0x62, 0x12, 0x20, 0x0a, 0x0c, 0x6d, 0x65, 0x61, 0x6e, 0x5f,
	0x6e, 0x75, 0x6d, 0x5f, 0x66, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d,
	0x65, 0x61, 0x6e, 0x4e, 0x75, 0x6d, 0x46, 0x64, 0x73, 0x12, 0x1e, 0x0a, 0x0b, 0x6d, 0x61, 0x78,
	0x5f, 0x6e, 0x75, 0x6d, 0x5f, 0x66, 0x64, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x6d, 0x61, 0x78, 0x4e, 0x75, 0x6d, 0x46, 0x64, 0x73, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73,
	0x70, 0x65, 0x61, 0x6b, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fleetspeak_src_server_proto_fleetspeak_server_resource_proto_rawDescOnce sync.Once
	file_fleetspeak_src_server_proto_fleetspeak_server_resource_proto_rawDescData = file_fleetspeak_src_server_proto_fleetspeak_server_resource_proto_rawDesc
)

func file_fleetspeak_src_server_proto_fleetspeak_server_resource_proto_rawDescGZIP() []byte {
	file_fleetspeak_src_server_proto_fleetspeak_server_resource_proto_rawDescOnce.Do(func() {
		file_fleetspeak_src_server_proto_fleetspeak_server_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_fleetspeak_src_server_proto_fleetspeak_server_resource_proto_rawDescData)
	})
	return file_fleetspeak_src_server_proto_fleetspeak_server_resource_proto_rawDescData
}

var file_fleetspeak_src_server_proto_fleetspeak_server_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_fleetspeak_src_server_proto_fleetspeak_server_resource_proto_goTypes = []interface{}{
	(*ClientResourceUsageRecord)(nil), // 0: fleetspeak.server.ClientResourceUsageRecord
	(*timestamppb.Timestamp)(nil),     // 1: google.protobuf.Timestamp
}
var file_fleetspeak_src_server_proto_fleetspeak_server_resource_proto_depIdxs = []int32{
	1, // 0: fleetspeak.server.ClientResourceUsageRecord.process_start_time:type_name -> google.protobuf.Timestamp
	1, // 1: fleetspeak.server.ClientResourceUsageRecord.client_timestamp:type_name -> google.protobuf.Timestamp
	1, // 2: fleetspeak.server.ClientResourceUsageRecord.server_timestamp:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_fleetspeak_src_server_proto_fleetspeak_server_resource_proto_init() }
func file_fleetspeak_src_server_proto_fleetspeak_server_resource_proto_init() {
	if File_fleetspeak_src_server_proto_fleetspeak_server_resource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fleetspeak_src_server_proto_fleetspeak_server_resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientResourceUsageRecord); i {
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
			RawDescriptor: file_fleetspeak_src_server_proto_fleetspeak_server_resource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fleetspeak_src_server_proto_fleetspeak_server_resource_proto_goTypes,
		DependencyIndexes: file_fleetspeak_src_server_proto_fleetspeak_server_resource_proto_depIdxs,
		MessageInfos:      file_fleetspeak_src_server_proto_fleetspeak_server_resource_proto_msgTypes,
	}.Build()
	File_fleetspeak_src_server_proto_fleetspeak_server_resource_proto = out.File
	file_fleetspeak_src_server_proto_fleetspeak_server_resource_proto_rawDesc = nil
	file_fleetspeak_src_server_proto_fleetspeak_server_resource_proto_goTypes = nil
	file_fleetspeak_src_server_proto_fleetspeak_server_resource_proto_depIdxs = nil
}
