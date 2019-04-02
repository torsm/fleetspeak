// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fleetspeak/src/server/components/proto/fleetspeak_components/config.proto

package fleetspeak_components

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Config struct {
	// Mysql connection string. Required.
	//
	// https://github.com/go-sql-driver/mysql#dsn-data-source-name
	MysqlDataSourceName string `protobuf:"bytes,1,opt,name=mysql_data_source_name,json=mysqlDataSourceName,proto3" json:"mysql_data_source_name,omitempty"`
	// The parameters required to stand up an https server. Required.
	HttpsConfig *HttpsConfig `protobuf:"bytes,2,opt,name=https_config,json=httpsConfig,proto3" json:"https_config,omitempty"`
	// If set, expects connections to arrive through a load balance implementing
	// the PROXY protocol.
	//
	// https://www.haproxy.org/download/1.8/doc/proxy-protocol.txt
	ProxyProtocol bool `protobuf:"varint,3,opt,name=proxy_protocol,json=proxyProtocol,proto3" json:"proxy_protocol,omitempty"`
	// If set, only clients reporting this label will be allowed to
	// connect. Meant as a sanity check that the client and server are for the
	// same Fleetspeak installation.
	RequiredLabel        string   `protobuf:"bytes,4,opt,name=required_label,json=requiredLabel,proto3" json:"required_label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_48a3be44b76012e1, []int{0}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (dst *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(dst, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetMysqlDataSourceName() string {
	if m != nil {
		return m.MysqlDataSourceName
	}
	return ""
}

func (m *Config) GetHttpsConfig() *HttpsConfig {
	if m != nil {
		return m.HttpsConfig
	}
	return nil
}

func (m *Config) GetProxyProtocol() bool {
	if m != nil {
		return m.ProxyProtocol
	}
	return false
}

func (m *Config) GetRequiredLabel() string {
	if m != nil {
		return m.RequiredLabel
	}
	return ""
}

type HttpsConfig struct {
	// The bind address to listen on for connections on, e.g. ":443" or
	// "localhost:1234". Required.
	ListenAddress string `protobuf:"bytes,1,opt,name=listen_address,json=listenAddress,proto3" json:"listen_address,omitempty"`
	// A certificate chain which identifies the server to clients. Must lead to a
	// certificate known to the clients. x509 format. Required.
	Certificates string `protobuf:"bytes,2,opt,name=certificates,proto3" json:"certificates,omitempty"`
	// The private key used to identify the server. Must match the first entry in
	// certificates. x509 format. Required.
	Key string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	// If set, disables long running (streaming) connections.
	DisableStreaming     bool     `protobuf:"varint,4,opt,name=disable_streaming,json=disableStreaming,proto3" json:"disable_streaming,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HttpsConfig) Reset()         { *m = HttpsConfig{} }
func (m *HttpsConfig) String() string { return proto.CompactTextString(m) }
func (*HttpsConfig) ProtoMessage()    {}
func (*HttpsConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_48a3be44b76012e1, []int{1}
}
func (m *HttpsConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpsConfig.Unmarshal(m, b)
}
func (m *HttpsConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpsConfig.Marshal(b, m, deterministic)
}
func (dst *HttpsConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpsConfig.Merge(dst, src)
}
func (m *HttpsConfig) XXX_Size() int {
	return xxx_messageInfo_HttpsConfig.Size(m)
}
func (m *HttpsConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpsConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HttpsConfig proto.InternalMessageInfo

func (m *HttpsConfig) GetListenAddress() string {
	if m != nil {
		return m.ListenAddress
	}
	return ""
}

func (m *HttpsConfig) GetCertificates() string {
	if m != nil {
		return m.Certificates
	}
	return ""
}

func (m *HttpsConfig) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *HttpsConfig) GetDisableStreaming() bool {
	if m != nil {
		return m.DisableStreaming
	}
	return false
}

func init() {
	proto.RegisterType((*Config)(nil), "fleetspeak.components.Config")
	proto.RegisterType((*HttpsConfig)(nil), "fleetspeak.components.HttpsConfig")
}

func init() {
	proto.RegisterFile("fleetspeak/src/server/components/proto/fleetspeak_components/config.proto", fileDescriptor_config_48a3be44b76012e1)
}

var fileDescriptor_config_48a3be44b76012e1 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xe1, 0x6a, 0xf2, 0x30,
	0x14, 0x86, 0xe9, 0xe7, 0x87, 0x68, 0xd4, 0xe1, 0x32, 0x36, 0xfa, 0x53, 0x0a, 0x03, 0x61, 0xd0,
	0xc2, 0xbc, 0x82, 0xb1, 0x0d, 0x1c, 0x8c, 0x31, 0xea, 0x05, 0x84, 0x63, 0x7a, 0xd4, 0x60, 0x9a,
	0xd4, 0x9c, 0x38, 0xe6, 0x95, 0xec, 0xbe, 0x76, 0x45, 0xa3, 0xa9, 0xd2, 0x09, 0xfb, 0xd7, 0x3e,
	0xef, 0x43, 0xce, 0xfb, 0xb2, 0x97, 0x95, 0x46, 0xf4, 0x54, 0x21, 0x6c, 0x33, 0x72, 0x32, 0x23,
	0x74, 0x1f, 0xe8, 0x32, 0x69, 0xcb, 0xca, 0x1a, 0x34, 0x9e, 0xb2, 0xca, 0x59, 0x6f, 0xb3, 0x56,
	0x13, 0xbf, 0x32, 0x69, 0xcd, 0x4a, 0xad, 0xd3, 0xa0, 0xf0, 0xeb, 0xd6, 0x49, 0x5b, 0x27, 0xf9,
	0x8e, 0x58, 0xf7, 0x31, 0x78, 0x7c, 0xc6, 0x6e, 0xca, 0x03, 0xed, 0xb4, 0x28, 0xc0, 0x83, 0x20,
	0xbb, 0x77, 0x12, 0x85, 0x81, 0x12, 0xe3, 0x68, 0x12, 0x4d, 0xfb, 0xf9, 0x55, 0x48, 0x9f, 0xc0,
	0xc3, 0x22, 0x64, 0x6f, 0x50, 0x22, 0x7f, 0x66, 0xc3, 0x8d, 0xf7, 0x15, 0x89, 0xe6, 0x58, 0xfc,
	0x6f, 0x12, 0x4d, 0x07, 0xf7, 0x49, 0xfa, 0xe7, 0xb5, 0x74, 0x5e, 0xab, 0xcd, 0xb9, 0x7c, 0xb0,
	0x69, 0x7f, 0xf8, 0x2d, 0xbb, 0xa8, 0x9c, 0xfd, 0x3c, 0x88, 0x50, 0x56, 0x5a, 0x1d, 0x77, 0x26,
	0xd1, 0xb4, 0x97, 0x8f, 0x02, 0x7d, 0x3f, 0xc2, 0x5a, 0x73, 0xb8, 0xdb, 0x2b, 0x87, 0x85, 0xd0,
	0xb0, 0x44, 0x1d, 0xff, 0x0f, 0xd5, 0x46, 0x27, 0xfa, 0x5a, 0xc3, 0xe4, 0x2b, 0x62, 0x83, 0xf9,
	0xf9, 0xeb, 0x5a, 0x91, 0x47, 0x23, 0xa0, 0x28, 0x1c, 0x12, 0x1d, 0x17, 0x8d, 0x1a, 0xfa, 0xd0,
	0x40, 0x9e, 0xb0, 0xa1, 0x44, 0xe7, 0xd5, 0x4a, 0x49, 0xf0, 0x48, 0x61, 0x4b, 0x3f, 0x3f, 0x63,
	0x7c, 0xcc, 0x3a, 0x5b, 0x3c, 0x84, 0x76, 0xfd, 0xbc, 0xfe, 0xe4, 0x77, 0xec, 0xb2, 0x50, 0x04,
	0x4b, 0x8d, 0x82, 0xbc, 0x43, 0x28, 0x95, 0x59, 0x87, 0x5a, 0xbd, 0x7c, 0x7c, 0x0c, 0x16, 0x27,
	0xbe, 0xec, 0x86, 0x7d, 0xb3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x02, 0x59, 0xbf, 0xd9,
	0x01, 0x00, 0x00,
}
