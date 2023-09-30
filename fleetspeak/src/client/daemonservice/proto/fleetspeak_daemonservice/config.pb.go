// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.24.3
// source: fleetspeak/src/client/daemonservice/proto/fleetspeak_daemonservice/config.proto

package fleetspeak_daemonservice

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The configuration information expected by daemonservice.Factory in
// ClientServiceConfig.config.
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Argv []string `protobuf:"bytes,1,rep,name=argv,proto3" json:"argv,omitempty"`
	// If set, process will be killed after this much inactivity. Any message to
	// or from the process, and any stdin/stderr output counts as inactivity.
	InactivityTimeout *durationpb.Duration `protobuf:"bytes,2,opt,name=inactivity_timeout,json=inactivityTimeout,proto3" json:"inactivity_timeout,omitempty"`
	// If set, start the process only when there is a message for it to work on.
	// Forced to true when inactivity timeout is set.
	LazyStart bool `protobuf:"varint,3,opt,name=lazy_start,json=lazyStart,proto3" json:"lazy_start,omitempty"`
	// By default, daemon services report resource usage every 10 minutes. This
	// flag disables this if set.
	DisableResourceMonitoring bool `protobuf:"varint,4,opt,name=disable_resource_monitoring,json=disableResourceMonitoring,proto3" json:"disable_resource_monitoring,omitempty"`
	// How many samples to aggregate into a report when monitoring resource usage.
	// If unset, defaults to 20.
	ResourceMonitoringSampleSize int32 `protobuf:"varint,5,opt,name=resource_monitoring_sample_size,json=resourceMonitoringSampleSize,proto3" json:"resource_monitoring_sample_size,omitempty"`
	// How long to wait between resource monitoring samples. If unset, defaults to
	// 30.
	ResourceMonitoringSamplePeriodSeconds int32 `protobuf:"varint,6,opt,name=resource_monitoring_sample_period_seconds,json=resourceMonitoringSamplePeriodSeconds,proto3" json:"resource_monitoring_sample_period_seconds,omitempty"`
	// If set, Fleetspeak will kill and restart the child if it exceeds this
	// memory limit, in bytes.
	MemoryLimit int64 `protobuf:"varint,7,opt,name=memory_limit,json=memoryLimit,proto3" json:"memory_limit,omitempty"`
	// If set, Fleetspeak will monitor child's heartbeat messages and kill
	// unresponsive processes. The values below should be set to configure the
	// heartbeat monitoring.
	MonitorHeartbeats bool `protobuf:"varint,8,opt,name=monitor_heartbeats,json=monitorHeartbeats,proto3" json:"monitor_heartbeats,omitempty"`
	// How long to wait for initial heartbeat.
	HeartbeatUnresponsiveGracePeriodSeconds int32 `protobuf:"varint,9,opt,name=heartbeat_unresponsive_grace_period_seconds,json=heartbeatUnresponsiveGracePeriodSeconds,proto3" json:"heartbeat_unresponsive_grace_period_seconds,omitempty"`
	// How long to wait for subsequent heartbeats.
	HeartbeatUnresponsiveKillPeriodSeconds int32             `protobuf:"varint,10,opt,name=heartbeat_unresponsive_kill_period_seconds,json=heartbeatUnresponsiveKillPeriodSeconds,proto3" json:"heartbeat_unresponsive_kill_period_seconds,omitempty"`
	StdParams                              *Config_StdParams `protobuf:"bytes,11,opt,name=std_params,json=stdParams,proto3" json:"std_params,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetArgv() []string {
	if x != nil {
		return x.Argv
	}
	return nil
}

func (x *Config) GetInactivityTimeout() *durationpb.Duration {
	if x != nil {
		return x.InactivityTimeout
	}
	return nil
}

func (x *Config) GetLazyStart() bool {
	if x != nil {
		return x.LazyStart
	}
	return false
}

func (x *Config) GetDisableResourceMonitoring() bool {
	if x != nil {
		return x.DisableResourceMonitoring
	}
	return false
}

func (x *Config) GetResourceMonitoringSampleSize() int32 {
	if x != nil {
		return x.ResourceMonitoringSampleSize
	}
	return 0
}

func (x *Config) GetResourceMonitoringSamplePeriodSeconds() int32 {
	if x != nil {
		return x.ResourceMonitoringSamplePeriodSeconds
	}
	return 0
}

func (x *Config) GetMemoryLimit() int64 {
	if x != nil {
		return x.MemoryLimit
	}
	return 0
}

func (x *Config) GetMonitorHeartbeats() bool {
	if x != nil {
		return x.MonitorHeartbeats
	}
	return false
}

func (x *Config) GetHeartbeatUnresponsiveGracePeriodSeconds() int32 {
	if x != nil {
		return x.HeartbeatUnresponsiveGracePeriodSeconds
	}
	return 0
}

func (x *Config) GetHeartbeatUnresponsiveKillPeriodSeconds() int32 {
	if x != nil {
		return x.HeartbeatUnresponsiveKillPeriodSeconds
	}
	return 0
}

func (x *Config) GetStdParams() *Config_StdParams {
	if x != nil {
		return x.StdParams
	}
	return nil
}

// If set, we forward stderr and stdout data to the server as messages with:
//
// message_type="StdOutput"
// data=<fleetspeak.daemonservice.StdOutputData>
type Config_StdParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"` // Service on the server to forward to. Required.
	// A message will be sent when we have flush_bytes queued, or when we
	// have bytes flush_time_seconds old.
	FlushBytes       int32 `protobuf:"varint,2,opt,name=flush_bytes,json=flushBytes,proto3" json:"flush_bytes,omitempty"`                     // Default and maximum value is 1MB.
	FlushTimeSeconds int32 `protobuf:"varint,3,opt,name=flush_time_seconds,json=flushTimeSeconds,proto3" json:"flush_time_seconds,omitempty"` // Default is 60.
}

func (x *Config_StdParams) Reset() {
	*x = Config_StdParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config_StdParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config_StdParams) ProtoMessage() {}

func (x *Config_StdParams) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config_StdParams.ProtoReflect.Descriptor instead.
func (*Config_StdParams) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Config_StdParams) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *Config_StdParams) GetFlushBytes() int32 {
	if x != nil {
		return x.FlushBytes
	}
	return 0
}

func (x *Config_StdParams) GetFlushTimeSeconds() int32 {
	if x != nil {
		return x.FlushTimeSeconds
	}
	return 0
}

var File_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto protoreflect.FileDescriptor

var file_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto_rawDesc = []byte{
	0x0a, 0x4f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2f, 0x73, 0x72, 0x63,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6c, 0x65, 0x65,
	0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x5f, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x18, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2e, 0x64, 0x61,
	0x65, 0x6d, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x06, 0x0a, 0x06,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x76, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x76, 0x12, 0x48, 0x0a, 0x12, 0x69, 0x6e,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x11, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x61, 0x7a, 0x79, 0x5f, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6c, 0x61, 0x7a, 0x79, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x3e, 0x0a, 0x1b, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x19, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x12, 0x45, 0x0a, 0x1f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x1c, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x58, 0x0a, 0x29, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x25, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x2d, 0x0a, 0x12, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x5f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x11, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x73, 0x12, 0x5c, 0x0a, 0x2b, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x5f, 0x75, 0x6e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x69, 0x76, 0x65,
	0x5f, 0x67, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x27, 0x68, 0x65, 0x61,
	0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x55, 0x6e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x69,
	0x76, 0x65, 0x47, 0x72, 0x61, 0x63, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x12, 0x5a, 0x0a, 0x2a, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61,
	0x74, 0x5f, 0x75, 0x6e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x69, 0x76, 0x65, 0x5f, 0x6b,
	0x69, 0x6c, 0x6c, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x26, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x55, 0x6e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x69, 0x76, 0x65, 0x4b,
	0x69, 0x6c, 0x6c, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73,
	0x12, 0x49, 0x0a, 0x0a, 0x73, 0x74, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61,
	0x6b, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x74, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x52, 0x09, 0x73, 0x74, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x7d, 0x0a, 0x09, 0x53,
	0x74, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66,
	0x6c, 0x75, 0x73, 0x68, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x66, 0x6c, 0x75, 0x73, 0x68, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x12,
	0x66, 0x6c, 0x75, 0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x66, 0x6c, 0x75, 0x73, 0x68, 0x54,
	0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x42, 0x61, 0x5a, 0x5f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74,
	0x73, 0x70, 0x65, 0x61, 0x6b, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x5f,
	0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto_rawDescOnce sync.Once
	file_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto_rawDescData = file_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto_rawDesc
)

func file_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto_rawDescGZIP() []byte {
	file_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto_rawDescOnce.Do(func() {
		file_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto_rawDescData)
	})
	return file_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto_rawDescData
}

var file_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto_goTypes = []interface{}{
	(*Config)(nil),              // 0: fleetspeak.daemonservice.Config
	(*Config_StdParams)(nil),    // 1: fleetspeak.daemonservice.Config.StdParams
	(*durationpb.Duration)(nil), // 2: google.protobuf.Duration
}
var file_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto_depIdxs = []int32{
	2, // 0: fleetspeak.daemonservice.Config.inactivity_timeout:type_name -> google.protobuf.Duration
	1, // 1: fleetspeak.daemonservice.Config.std_params:type_name -> fleetspeak.daemonservice.Config.StdParams
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() {
	file_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto_init()
}
func file_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto_init() {
	if File_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config_StdParams); i {
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
			RawDescriptor: file_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto_goTypes,
		DependencyIndexes: file_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto_depIdxs,
		MessageInfos:      file_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto_msgTypes,
	}.Build()
	File_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto = out.File
	file_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto_rawDesc = nil
	file_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto_goTypes = nil
	file_fleetspeak_src_client_daemonservice_proto_fleetspeak_daemonservice_config_proto_depIdxs = nil
}
