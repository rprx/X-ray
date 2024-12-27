// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.2
// source: proxy/freedom/config.proto

package freedom

import (
	protocol "github.com/GFW-knocker/Xray-core/common/protocol"
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

type Config_DomainStrategy int32

const (
	Config_AS_IS      Config_DomainStrategy = 0
	Config_USE_IP     Config_DomainStrategy = 1
	Config_USE_IP4    Config_DomainStrategy = 2
	Config_USE_IP6    Config_DomainStrategy = 3
	Config_USE_IP46   Config_DomainStrategy = 4
	Config_USE_IP64   Config_DomainStrategy = 5
	Config_FORCE_IP   Config_DomainStrategy = 6
	Config_FORCE_IP4  Config_DomainStrategy = 7
	Config_FORCE_IP6  Config_DomainStrategy = 8
	Config_FORCE_IP46 Config_DomainStrategy = 9
	Config_FORCE_IP64 Config_DomainStrategy = 10
)

// Enum value maps for Config_DomainStrategy.
var (
	Config_DomainStrategy_name = map[int32]string{
		0:  "AS_IS",
		1:  "USE_IP",
		2:  "USE_IP4",
		3:  "USE_IP6",
		4:  "USE_IP46",
		5:  "USE_IP64",
		6:  "FORCE_IP",
		7:  "FORCE_IP4",
		8:  "FORCE_IP6",
		9:  "FORCE_IP46",
		10: "FORCE_IP64",
	}
	Config_DomainStrategy_value = map[string]int32{
		"AS_IS":      0,
		"USE_IP":     1,
		"USE_IP4":    2,
		"USE_IP6":    3,
		"USE_IP46":   4,
		"USE_IP64":   5,
		"FORCE_IP":   6,
		"FORCE_IP4":  7,
		"FORCE_IP6":  8,
		"FORCE_IP46": 9,
		"FORCE_IP64": 10,
	}
)

func (x Config_DomainStrategy) Enum() *Config_DomainStrategy {
	p := new(Config_DomainStrategy)
	*p = x
	return p
}

func (x Config_DomainStrategy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Config_DomainStrategy) Descriptor() protoreflect.EnumDescriptor {
	return file_proxy_freedom_config_proto_enumTypes[0].Descriptor()
}

func (Config_DomainStrategy) Type() protoreflect.EnumType {
	return &file_proxy_freedom_config_proto_enumTypes[0]
}

func (x Config_DomainStrategy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Config_DomainStrategy.Descriptor instead.
func (Config_DomainStrategy) EnumDescriptor() ([]byte, []int) {
	return file_proxy_freedom_config_proto_rawDescGZIP(), []int{3, 0}
}

type DestinationOverride struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server *protocol.ServerEndpoint `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
}

func (x *DestinationOverride) Reset() {
	*x = DestinationOverride{}
	mi := &file_proxy_freedom_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DestinationOverride) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestinationOverride) ProtoMessage() {}

func (x *DestinationOverride) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_freedom_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestinationOverride.ProtoReflect.Descriptor instead.
func (*DestinationOverride) Descriptor() ([]byte, []int) {
	return file_proxy_freedom_config_proto_rawDescGZIP(), []int{0}
}

func (x *DestinationOverride) GetServer() *protocol.ServerEndpoint {
	if x != nil {
		return x.Server
	}
	return nil
}

type Fragment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PacketsFrom uint64 `protobuf:"varint,1,opt,name=packets_from,json=packetsFrom,proto3" json:"packets_from,omitempty"`
	PacketsTo   uint64 `protobuf:"varint,2,opt,name=packets_to,json=packetsTo,proto3" json:"packets_to,omitempty"`
	LengthMin   uint64 `protobuf:"varint,3,opt,name=length_min,json=lengthMin,proto3" json:"length_min,omitempty"`
	LengthMax   uint64 `protobuf:"varint,4,opt,name=length_max,json=lengthMax,proto3" json:"length_max,omitempty"`
	IntervalMin uint64 `protobuf:"varint,5,opt,name=interval_min,json=intervalMin,proto3" json:"interval_min,omitempty"`
	IntervalMax uint64 `protobuf:"varint,6,opt,name=interval_max,json=intervalMax,proto3" json:"interval_max,omitempty"`
	FakeHost    bool   `protobuf:"varint,7,opt,name=fake_host,json=fakeHost,proto3" json:"fake_host,omitempty"`
	Host1Header string `protobuf:"bytes,8,opt,name=host1_header,json=host1Header,proto3" json:"host1_header,omitempty"`
	Host1Domain string `protobuf:"bytes,9,opt,name=host1_domain,json=host1Domain,proto3" json:"host1_domain,omitempty"`
	Host2Header string `protobuf:"bytes,10,opt,name=host2_header,json=host2Header,proto3" json:"host2_header,omitempty"`
	Host2Domain string `protobuf:"bytes,11,opt,name=host2_domain,json=host2Domain,proto3" json:"host2_domain,omitempty"`
}

func (x *Fragment) Reset() {
	*x = Fragment{}
	mi := &file_proxy_freedom_config_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Fragment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fragment) ProtoMessage() {}

func (x *Fragment) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_freedom_config_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fragment.ProtoReflect.Descriptor instead.
func (*Fragment) Descriptor() ([]byte, []int) {
	return file_proxy_freedom_config_proto_rawDescGZIP(), []int{1}
}

func (x *Fragment) GetPacketsFrom() uint64 {
	if x != nil {
		return x.PacketsFrom
	}
	return 0
}

func (x *Fragment) GetPacketsTo() uint64 {
	if x != nil {
		return x.PacketsTo
	}
	return 0
}

func (x *Fragment) GetLengthMin() uint64 {
	if x != nil {
		return x.LengthMin
	}
	return 0
}

func (x *Fragment) GetLengthMax() uint64 {
	if x != nil {
		return x.LengthMax
	}
	return 0
}

func (x *Fragment) GetIntervalMin() uint64 {
	if x != nil {
		return x.IntervalMin
	}
	return 0
}

func (x *Fragment) GetIntervalMax() uint64 {
	if x != nil {
		return x.IntervalMax
	}
	return 0
}

func (x *Fragment) GetFakeHost() bool {
	if x != nil {
		return x.FakeHost
	}
	return false
}

func (x *Fragment) GetHost1Header() string {
	if x != nil {
		return x.Host1Header
	}
	return ""
}

func (x *Fragment) GetHost1Domain() string {
	if x != nil {
		return x.Host1Domain
	}
	return ""
}

func (x *Fragment) GetHost2Header() string {
	if x != nil {
		return x.Host2Header
	}
	return ""
}

func (x *Fragment) GetHost2Domain() string {
	if x != nil {
		return x.Host2Domain
	}
	return ""
}

type Noise struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LengthMin uint64 `protobuf:"varint,1,opt,name=length_min,json=lengthMin,proto3" json:"length_min,omitempty"`
	LengthMax uint64 `protobuf:"varint,2,opt,name=length_max,json=lengthMax,proto3" json:"length_max,omitempty"`
	DelayMin  uint64 `protobuf:"varint,3,opt,name=delay_min,json=delayMin,proto3" json:"delay_min,omitempty"`
	DelayMax  uint64 `protobuf:"varint,4,opt,name=delay_max,json=delayMax,proto3" json:"delay_max,omitempty"`
	StrNoise  []byte `protobuf:"bytes,5,opt,name=str_noise,json=strNoise,proto3" json:"str_noise,omitempty"`
}

func (x *Noise) Reset() {
	*x = Noise{}
	mi := &file_proxy_freedom_config_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Noise) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Noise) ProtoMessage() {}

func (x *Noise) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_freedom_config_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Noise.ProtoReflect.Descriptor instead.
func (*Noise) Descriptor() ([]byte, []int) {
	return file_proxy_freedom_config_proto_rawDescGZIP(), []int{2}
}

func (x *Noise) GetLengthMin() uint64 {
	if x != nil {
		return x.LengthMin
	}
	return 0
}

func (x *Noise) GetLengthMax() uint64 {
	if x != nil {
		return x.LengthMax
	}
	return 0
}

func (x *Noise) GetDelayMin() uint64 {
	if x != nil {
		return x.DelayMin
	}
	return 0
}

func (x *Noise) GetDelayMax() uint64 {
	if x != nil {
		return x.DelayMax
	}
	return 0
}

func (x *Noise) GetStrNoise() []byte {
	if x != nil {
		return x.StrNoise
	}
	return nil
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DomainStrategy      Config_DomainStrategy `protobuf:"varint,1,opt,name=domain_strategy,json=domainStrategy,proto3,enum=xray.proxy.freedom.Config_DomainStrategy" json:"domain_strategy,omitempty"`
	DestinationOverride *DestinationOverride  `protobuf:"bytes,3,opt,name=destination_override,json=destinationOverride,proto3" json:"destination_override,omitempty"`
	UserLevel           uint32                `protobuf:"varint,4,opt,name=user_level,json=userLevel,proto3" json:"user_level,omitempty"`
	Fragment            *Fragment             `protobuf:"bytes,5,opt,name=fragment,proto3" json:"fragment,omitempty"`
	ProxyProtocol       uint32                `protobuf:"varint,6,opt,name=proxy_protocol,json=proxyProtocol,proto3" json:"proxy_protocol,omitempty"`
	Noises              []*Noise              `protobuf:"bytes,7,rep,name=noises,proto3" json:"noises,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_proxy_freedom_config_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_freedom_config_proto_msgTypes[3]
	if x != nil {
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
	return file_proxy_freedom_config_proto_rawDescGZIP(), []int{3}
}

func (x *Config) GetDomainStrategy() Config_DomainStrategy {
	if x != nil {
		return x.DomainStrategy
	}
	return Config_AS_IS
}

func (x *Config) GetDestinationOverride() *DestinationOverride {
	if x != nil {
		return x.DestinationOverride
	}
	return nil
}

func (x *Config) GetUserLevel() uint32 {
	if x != nil {
		return x.UserLevel
	}
	return 0
}

func (x *Config) GetFragment() *Fragment {
	if x != nil {
		return x.Fragment
	}
	return nil
}

func (x *Config) GetProxyProtocol() uint32 {
	if x != nil {
		return x.ProxyProtocol
	}
	return 0
}

func (x *Config) GetNoises() []*Noise {
	if x != nil {
		return x.Noises
	}
	return nil
}

var File_proxy_freedom_config_proto protoreflect.FileDescriptor

var file_proxy_freedom_config_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x66, 0x72, 0x65, 0x65, 0x64, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x78, 0x72,
	0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x66, 0x72, 0x65, 0x65, 0x64, 0x6f, 0x6d,
	0x1a, 0x21, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x78, 0x72, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0xf9, 0x02, 0x0a, 0x08, 0x46, 0x72, 0x61,
	0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x70, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x73, 0x5f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x54, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x4d, 0x69, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x5f, 0x6d, 0x61, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6c, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x4d, 0x61, 0x78, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x4d, 0x69, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x4d, 0x61, 0x78, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x61, 0x6b, 0x65, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x66, 0x61, 0x6b, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x6f, 0x73, 0x74,
	0x31, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x68, 0x6f, 0x73, 0x74, 0x31, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x68,
	0x6f, 0x73, 0x74, 0x31, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x68, 0x6f, 0x73, 0x74, 0x31, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x21,
	0x0a, 0x0c, 0x68, 0x6f, 0x73, 0x74, 0x32, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x6f, 0x73, 0x74, 0x32, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x6f, 0x73, 0x74, 0x32, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x6f, 0x73, 0x74, 0x32, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x22, 0x9c, 0x01, 0x0a, 0x05, 0x4e, 0x6f, 0x69, 0x73, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x09, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x4d, 0x69, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x4d, 0x61, 0x78, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x4d, 0x69, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x6c,
	0x61, 0x79, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x65,
	0x6c, 0x61, 0x79, 0x4d, 0x61, 0x78, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x72, 0x5f, 0x6e, 0x6f,
	0x69, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x73, 0x74, 0x72, 0x4e, 0x6f,
	0x69, 0x73, 0x65, 0x22, 0x97, 0x04, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x52,
	0x0a, 0x0f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x66, 0x72, 0x65, 0x65, 0x64, 0x6f, 0x6d, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x52, 0x0e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x12, 0x5a, 0x0a, 0x14, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x66, 0x72,
	0x65, 0x65, 0x64, 0x6f, 0x6d, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x52, 0x13, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x38, 0x0a,
	0x08, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x66, 0x72, 0x65,
	0x65, 0x64, 0x6f, 0x6d, 0x2e, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x66,
	0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x31,
	0x0a, 0x06, 0x6e, 0x6f, 0x69, 0x73, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x66, 0x72, 0x65, 0x65,
	0x64, 0x6f, 0x6d, 0x2e, 0x4e, 0x6f, 0x69, 0x73, 0x65, 0x52, 0x06, 0x6e, 0x6f, 0x69, 0x73, 0x65,
	0x73, 0x22, 0xa9, 0x01, 0x0a, 0x0e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x53, 0x5f, 0x49, 0x53, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x55, 0x53, 0x45, 0x5f, 0x49, 0x50, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x53, 0x45, 0x5f, 0x49, 0x50, 0x34, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x53, 0x45, 0x5f,
	0x49, 0x50, 0x36, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x53, 0x45, 0x5f, 0x49, 0x50, 0x34,
	0x36, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x53, 0x45, 0x5f, 0x49, 0x50, 0x36, 0x34, 0x10,
	0x05, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x5f, 0x49, 0x50, 0x10, 0x06, 0x12,
	0x0d, 0x0a, 0x09, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x5f, 0x49, 0x50, 0x34, 0x10, 0x07, 0x12, 0x0d,
	0x0a, 0x09, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x5f, 0x49, 0x50, 0x36, 0x10, 0x08, 0x12, 0x0e, 0x0a,
	0x0a, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x5f, 0x49, 0x50, 0x34, 0x36, 0x10, 0x09, 0x12, 0x0e, 0x0a,
	0x0a, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x5f, 0x49, 0x50, 0x36, 0x34, 0x10, 0x0a, 0x42, 0x5f, 0x0a,
	0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x66, 0x72, 0x65, 0x65, 0x64, 0x6f, 0x6d, 0x50, 0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x46, 0x57, 0x2d, 0x6b, 0x6e, 0x6f, 0x63, 0x6b, 0x65,
	0x72, 0x2f, 0x58, 0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x66, 0x72, 0x65, 0x65, 0x64, 0x6f, 0x6d, 0xaa, 0x02, 0x12, 0x58, 0x72, 0x61, 0x79,
	0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x46, 0x72, 0x65, 0x65, 0x64, 0x6f, 0x6d, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proxy_freedom_config_proto_rawDescOnce sync.Once
	file_proxy_freedom_config_proto_rawDescData = file_proxy_freedom_config_proto_rawDesc
)

func file_proxy_freedom_config_proto_rawDescGZIP() []byte {
	file_proxy_freedom_config_proto_rawDescOnce.Do(func() {
		file_proxy_freedom_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_proxy_freedom_config_proto_rawDescData)
	})
	return file_proxy_freedom_config_proto_rawDescData
}

var file_proxy_freedom_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proxy_freedom_config_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proxy_freedom_config_proto_goTypes = []any{
	(Config_DomainStrategy)(0),      // 0: xray.proxy.freedom.Config.DomainStrategy
	(*DestinationOverride)(nil),     // 1: xray.proxy.freedom.DestinationOverride
	(*Fragment)(nil),                // 2: xray.proxy.freedom.Fragment
	(*Noise)(nil),                   // 3: xray.proxy.freedom.Noise
	(*Config)(nil),                  // 4: xray.proxy.freedom.Config
	(*protocol.ServerEndpoint)(nil), // 5: xray.common.protocol.ServerEndpoint
}
var file_proxy_freedom_config_proto_depIdxs = []int32{
	5, // 0: xray.proxy.freedom.DestinationOverride.server:type_name -> xray.common.protocol.ServerEndpoint
	0, // 1: xray.proxy.freedom.Config.domain_strategy:type_name -> xray.proxy.freedom.Config.DomainStrategy
	1, // 2: xray.proxy.freedom.Config.destination_override:type_name -> xray.proxy.freedom.DestinationOverride
	2, // 3: xray.proxy.freedom.Config.fragment:type_name -> xray.proxy.freedom.Fragment
	3, // 4: xray.proxy.freedom.Config.noises:type_name -> xray.proxy.freedom.Noise
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proxy_freedom_config_proto_init() }
func file_proxy_freedom_config_proto_init() {
	if File_proxy_freedom_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proxy_freedom_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proxy_freedom_config_proto_goTypes,
		DependencyIndexes: file_proxy_freedom_config_proto_depIdxs,
		EnumInfos:         file_proxy_freedom_config_proto_enumTypes,
		MessageInfos:      file_proxy_freedom_config_proto_msgTypes,
	}.Build()
	File_proxy_freedom_config_proto = out.File
	file_proxy_freedom_config_proto_rawDesc = nil
	file_proxy_freedom_config_proto_goTypes = nil
	file_proxy_freedom_config_proto_depIdxs = nil
}
