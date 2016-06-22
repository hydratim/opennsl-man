// Code generated by protoc-gen-go.
// source: stat.proto
// DO NOT EDIT!

/*
Package stat is a generated protocol buffer package.

It is generated from these files:
	stat.proto

It has these top-level messages:
	InitRequest
	InitResponse
	ClearRequest
	ClearResponse
	SyncRequest
	SyncResponse
	GetRequest
	GetResponse
	MultiGetRequest
	MultiGetResponse
*/
package stat

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

type StatType int32

const (
	StatType_STAT_TYPE_IF_IN_OCTETS                            StatType = 0
	StatType_STAT_TYPE_IF_IN_UCAST_PKTS                        StatType = 1
	StatType_STAT_TYPE_IF_IN_NUCAST_PKTS                       StatType = 2
	StatType_STAT_TYPE_IF_IN_DISCARDS                          StatType = 3
	StatType_STAT_TYPE_IF_IN_ERRORS                            StatType = 4
	StatType_STAT_TYPE_IF_IN_UNKNOWN_PROTOS                    StatType = 5
	StatType_STAT_TYPE_IF_OUT_OCTETS                           StatType = 6
	StatType_STAT_TYPE_IF_OUT_UCAST_PKTS                       StatType = 7
	StatType_STAT_TYPE_IF_OUT_NUCAST_PKTS                      StatType = 8
	StatType_STAT_TYPE_IF_OUT_DISCARDS                         StatType = 9
	StatType_STAT_TYPE_IF_OUT_ERRORS                           StatType = 10
	StatType_STAT_TYPE_IF_OUT_QLEN                             StatType = 11
	StatType_STAT_TYPE_IP_IN_RECEIVES                          StatType = 12
	StatType_STAT_TYPE_IP_IN_HDR_ERRORS                        StatType = 13
	StatType_STAT_TYPE_IP_FORW_DATAGRAMS                       StatType = 14
	StatType_STAT_TYPE_IP_IN_DISCARDS                          StatType = 15
	StatType_STAT_TYPE_DOT1D_BASE_PORT_DELAY_EXCEEDED_DISCARDS StatType = 16
	StatType_STAT_TYPE_DOT1D_BASE_PORT_MTU_EXCEEDED_DISCARDS   StatType = 17
	StatType_STAT_TYPE_DOT1D_TP_PORT_IN_FRAMES                 StatType = 18
	StatType_STAT_TYPE_DOT1D_TP_PORT_OUT_FRAMES                StatType = 19
	StatType_STAT_TYPE_DOT1D_PORT_IN_DISCARDS                  StatType = 20
	StatType_STAT_TYPE_ETHER_STATS_DROP_EVENTS                 StatType = 21
	StatType_STAT_TYPE_ETHER_STATS_MULTICAST_PKTS              StatType = 22
	StatType_STAT_TYPE_ETHER_STATS_BROADCAST_PKTS              StatType = 23
	StatType_STAT_TYPE_ETHER_STATS_UNDERSIZE_PKTS              StatType = 24
	StatType_STAT_TYPE_ETHER_STATS_FRAGMENTS                   StatType = 25
	StatType_STAT_TYPE_ETHER_STATS_PKTS_64_OCTETS              StatType = 26
	StatType_STAT_TYPE_ETHER_STATS_PKTS_65TO127_OCTETS         StatType = 27
	StatType_STAT_TYPE_ETHER_STATS_PKTS_128TO255_OCTETS        StatType = 28
	StatType_STAT_TYPE_ETHER_STATS_PKTS_256TO511_OCTETS        StatType = 29
	StatType_STAT_TYPE_ETHER_STATS_PKTS_512TO1023_OCTETS       StatType = 30
	StatType_STAT_TYPE_ETHER_STATS_PKTS_1024TO1518_OCTETS      StatType = 31
	StatType_STAT_TYPE_ETHER_STATS_OVERSIZE_PKTS               StatType = 32
	StatType_STAT_TYPE_ETHER_RX_OVER_SIZE_PKTS                 StatType = 33
	StatType_STAT_TYPE_ETHER_TX_OVER_SIZE_PKTS                 StatType = 34
	StatType_STAT_TYPE_ETHER_STATS_JABBERS                     StatType = 35
	StatType_STAT_TYPE_ETHER_STATS_OCTETS                      StatType = 36
	StatType_STAT_TYPE_ETHER_STATS_PKTS                        StatType = 37
	StatType_STAT_TYPE_ETHER_STATS_COLLISIONS                  StatType = 38
	StatType_STAT_TYPE_ETHER_STATS_CRC_ALIGN_ERRORS            StatType = 39
	StatType_STAT_TYPE_ETHER_STATS_TX_NO_ERRORS                StatType = 40
	StatType_STAT_TYPE_ETHER_STATS_RX_NO_ERRORS                StatType = 41
	StatType_STAT_TYPE_DOT3_STATS_ALIGNMENT_ERRORS             StatType = 42
	StatType_STAT_TYPE_DOT3_STATS_FCS_ERRORS                   StatType = 43
	StatType_STAT_TYPE_DOT3_STATS_SINGLE_COLLISION_FRAMES      StatType = 44
	StatType_STAT_TYPE_DOT3_STATS_MULTIPLE_COLLISION_FRAMES    StatType = 45
	StatType_STAT_TYPE_DOT3_STATS_SQET_TEST_ERRORS             StatType = 46
	StatType_STAT_TYPE_DOT3_STATS_DEFERRED_TRANSMISSIONS       StatType = 47
	StatType_STAT_TYPE_DOT3_STATS_LATE_COLLISIONS              StatType = 48
	StatType_STAT_TYPE_DOT3_STATS_EXCESSIVE_COLLISIONS         StatType = 49
	StatType_STAT_TYPE_DOT3_STATS_INTERNAL_MAC_TRANSMIT_ERRORS StatType = 50
	StatType_STAT_TYPE_DOT3_STATS_CARRIER_SENSE_ERRORS         StatType = 51
	StatType_STAT_TYPE_DOT3_STATS_FRAME_TOO_LONGS              StatType = 52
	StatType_STAT_TYPE_DOT3_STATS_INTERNAL_MAC_RECEIVE_ERRORS  StatType = 53
	StatType_STAT_TYPE_DOT3_STATS_SYMBOL_ERRORS                StatType = 54
	StatType_STAT_TYPE_DOT3_CONTROL_IN_UNKNOWN_OP_CODES        StatType = 55
	StatType_STAT_TYPE_DOT3_IN_PAUSE_FRAMES                    StatType = 56
	StatType_STAT_TYPE_DOT3_OUT_PAUSE_FRAMES                   StatType = 57
	StatType_STAT_TYPE_IF_HC_IN_OCTETS                         StatType = 58
	StatType_STAT_TYPE_IF_HC_IN_UCAST_PKTS                     StatType = 59
	StatType_STAT_TYPE_IF_HC_IN_MULTICAST_PKTS                 StatType = 60
	StatType_STAT_TYPE_IF_HC_IN_BROADCAST_PKTS                 StatType = 61
	StatType_STAT_TYPE_IF_HC_OUT_OCTETS                        StatType = 62
	StatType_STAT_TYPE_IF_HC_OUT_UCAST_PKTS                    StatType = 63
	StatType_STAT_TYPE_IF_HC_OUT_MULTICAST_PKTS                StatType = 64
	StatType_STAT_TYPE_IF_HC_OUT_BROADCAST_PKTS                StatType = 65
	StatType_STAT_TYPE_IPV6_IF_STATS_IN_RECEIVES               StatType = 66
	StatType_STAT_TYPE_IPV6_IF_STATS_IN_HDR_ERRORS             StatType = 67
	StatType_STAT_TYPE_IPV6_IF_STATS_IN_ADDR_ERRORS            StatType = 68
	StatType_STAT_TYPE_IPV6_IF_STATS_IN_DISCARDS               StatType = 69
	StatType_STAT_TYPE_IPV6_IF_STATS_OUT_FORW_DATAGRAMS        StatType = 70
	StatType_STAT_TYPE_IPV6_IF_STATS_OUT_DISCARDS              StatType = 71
	StatType_STAT_TYPE_IPV6_IF_STATS_IN_MCAST_PKTS             StatType = 72
	StatType_STAT_TYPE_IPV6_IF_STATS_OUT_MCAST_PKTS            StatType = 73
	StatType_STAT_TYPE_IF_IN_BROADCAST_PKTS                    StatType = 74
	StatType_STAT_TYPE_IF_IN_MULTICAST_PKTS                    StatType = 75
	StatType_STAT_TYPE_IF_OUT_BROADCAST_PKTS                   StatType = 76
	StatType_STAT_TYPE_IF_OUT_MULTICAST_PKTS                   StatType = 77
	StatType_STAT_TYPE_IEEE8021_PFC_REQUESTS                   StatType = 78
	StatType_STAT_TYPE_IEEE8021_PFC_INDICATIONS                StatType = 79
)

var StatType_name = map[int32]string{
	0:  "STAT_TYPE_IF_IN_OCTETS",
	1:  "STAT_TYPE_IF_IN_UCAST_PKTS",
	2:  "STAT_TYPE_IF_IN_NUCAST_PKTS",
	3:  "STAT_TYPE_IF_IN_DISCARDS",
	4:  "STAT_TYPE_IF_IN_ERRORS",
	5:  "STAT_TYPE_IF_IN_UNKNOWN_PROTOS",
	6:  "STAT_TYPE_IF_OUT_OCTETS",
	7:  "STAT_TYPE_IF_OUT_UCAST_PKTS",
	8:  "STAT_TYPE_IF_OUT_NUCAST_PKTS",
	9:  "STAT_TYPE_IF_OUT_DISCARDS",
	10: "STAT_TYPE_IF_OUT_ERRORS",
	11: "STAT_TYPE_IF_OUT_QLEN",
	12: "STAT_TYPE_IP_IN_RECEIVES",
	13: "STAT_TYPE_IP_IN_HDR_ERRORS",
	14: "STAT_TYPE_IP_FORW_DATAGRAMS",
	15: "STAT_TYPE_IP_IN_DISCARDS",
	16: "STAT_TYPE_DOT1D_BASE_PORT_DELAY_EXCEEDED_DISCARDS",
	17: "STAT_TYPE_DOT1D_BASE_PORT_MTU_EXCEEDED_DISCARDS",
	18: "STAT_TYPE_DOT1D_TP_PORT_IN_FRAMES",
	19: "STAT_TYPE_DOT1D_TP_PORT_OUT_FRAMES",
	20: "STAT_TYPE_DOT1D_PORT_IN_DISCARDS",
	21: "STAT_TYPE_ETHER_STATS_DROP_EVENTS",
	22: "STAT_TYPE_ETHER_STATS_MULTICAST_PKTS",
	23: "STAT_TYPE_ETHER_STATS_BROADCAST_PKTS",
	24: "STAT_TYPE_ETHER_STATS_UNDERSIZE_PKTS",
	25: "STAT_TYPE_ETHER_STATS_FRAGMENTS",
	26: "STAT_TYPE_ETHER_STATS_PKTS_64_OCTETS",
	27: "STAT_TYPE_ETHER_STATS_PKTS_65TO127_OCTETS",
	28: "STAT_TYPE_ETHER_STATS_PKTS_128TO255_OCTETS",
	29: "STAT_TYPE_ETHER_STATS_PKTS_256TO511_OCTETS",
	30: "STAT_TYPE_ETHER_STATS_PKTS_512TO1023_OCTETS",
	31: "STAT_TYPE_ETHER_STATS_PKTS_1024TO1518_OCTETS",
	32: "STAT_TYPE_ETHER_STATS_OVERSIZE_PKTS",
	33: "STAT_TYPE_ETHER_RX_OVER_SIZE_PKTS",
	34: "STAT_TYPE_ETHER_TX_OVER_SIZE_PKTS",
	35: "STAT_TYPE_ETHER_STATS_JABBERS",
	36: "STAT_TYPE_ETHER_STATS_OCTETS",
	37: "STAT_TYPE_ETHER_STATS_PKTS",
	38: "STAT_TYPE_ETHER_STATS_COLLISIONS",
	39: "STAT_TYPE_ETHER_STATS_CRC_ALIGN_ERRORS",
	40: "STAT_TYPE_ETHER_STATS_TX_NO_ERRORS",
	41: "STAT_TYPE_ETHER_STATS_RX_NO_ERRORS",
	42: "STAT_TYPE_DOT3_STATS_ALIGNMENT_ERRORS",
	43: "STAT_TYPE_DOT3_STATS_FCS_ERRORS",
	44: "STAT_TYPE_DOT3_STATS_SINGLE_COLLISION_FRAMES",
	45: "STAT_TYPE_DOT3_STATS_MULTIPLE_COLLISION_FRAMES",
	46: "STAT_TYPE_DOT3_STATS_SQET_TEST_ERRORS",
	47: "STAT_TYPE_DOT3_STATS_DEFERRED_TRANSMISSIONS",
	48: "STAT_TYPE_DOT3_STATS_LATE_COLLISIONS",
	49: "STAT_TYPE_DOT3_STATS_EXCESSIVE_COLLISIONS",
	50: "STAT_TYPE_DOT3_STATS_INTERNAL_MAC_TRANSMIT_ERRORS",
	51: "STAT_TYPE_DOT3_STATS_CARRIER_SENSE_ERRORS",
	52: "STAT_TYPE_DOT3_STATS_FRAME_TOO_LONGS",
	53: "STAT_TYPE_DOT3_STATS_INTERNAL_MAC_RECEIVE_ERRORS",
	54: "STAT_TYPE_DOT3_STATS_SYMBOL_ERRORS",
	55: "STAT_TYPE_DOT3_CONTROL_IN_UNKNOWN_OP_CODES",
	56: "STAT_TYPE_DOT3_IN_PAUSE_FRAMES",
	57: "STAT_TYPE_DOT3_OUT_PAUSE_FRAMES",
	58: "STAT_TYPE_IF_HC_IN_OCTETS",
	59: "STAT_TYPE_IF_HC_IN_UCAST_PKTS",
	60: "STAT_TYPE_IF_HC_IN_MULTICAST_PKTS",
	61: "STAT_TYPE_IF_HC_IN_BROADCAST_PKTS",
	62: "STAT_TYPE_IF_HC_OUT_OCTETS",
	63: "STAT_TYPE_IF_HC_OUT_UCAST_PKTS",
	64: "STAT_TYPE_IF_HC_OUT_MULTICAST_PKTS",
	65: "STAT_TYPE_IF_HC_OUT_BROADCAST_PKTS",
	66: "STAT_TYPE_IPV6_IF_STATS_IN_RECEIVES",
	67: "STAT_TYPE_IPV6_IF_STATS_IN_HDR_ERRORS",
	68: "STAT_TYPE_IPV6_IF_STATS_IN_ADDR_ERRORS",
	69: "STAT_TYPE_IPV6_IF_STATS_IN_DISCARDS",
	70: "STAT_TYPE_IPV6_IF_STATS_OUT_FORW_DATAGRAMS",
	71: "STAT_TYPE_IPV6_IF_STATS_OUT_DISCARDS",
	72: "STAT_TYPE_IPV6_IF_STATS_IN_MCAST_PKTS",
	73: "STAT_TYPE_IPV6_IF_STATS_OUT_MCAST_PKTS",
	74: "STAT_TYPE_IF_IN_BROADCAST_PKTS",
	75: "STAT_TYPE_IF_IN_MULTICAST_PKTS",
	76: "STAT_TYPE_IF_OUT_BROADCAST_PKTS",
	77: "STAT_TYPE_IF_OUT_MULTICAST_PKTS",
	78: "STAT_TYPE_IEEE8021_PFC_REQUESTS",
	79: "STAT_TYPE_IEEE8021_PFC_INDICATIONS",
}
var StatType_value = map[string]int32{
	"STAT_TYPE_IF_IN_OCTETS":                            0,
	"STAT_TYPE_IF_IN_UCAST_PKTS":                        1,
	"STAT_TYPE_IF_IN_NUCAST_PKTS":                       2,
	"STAT_TYPE_IF_IN_DISCARDS":                          3,
	"STAT_TYPE_IF_IN_ERRORS":                            4,
	"STAT_TYPE_IF_IN_UNKNOWN_PROTOS":                    5,
	"STAT_TYPE_IF_OUT_OCTETS":                           6,
	"STAT_TYPE_IF_OUT_UCAST_PKTS":                       7,
	"STAT_TYPE_IF_OUT_NUCAST_PKTS":                      8,
	"STAT_TYPE_IF_OUT_DISCARDS":                         9,
	"STAT_TYPE_IF_OUT_ERRORS":                           10,
	"STAT_TYPE_IF_OUT_QLEN":                             11,
	"STAT_TYPE_IP_IN_RECEIVES":                          12,
	"STAT_TYPE_IP_IN_HDR_ERRORS":                        13,
	"STAT_TYPE_IP_FORW_DATAGRAMS":                       14,
	"STAT_TYPE_IP_IN_DISCARDS":                          15,
	"STAT_TYPE_DOT1D_BASE_PORT_DELAY_EXCEEDED_DISCARDS": 16,
	"STAT_TYPE_DOT1D_BASE_PORT_MTU_EXCEEDED_DISCARDS":   17,
	"STAT_TYPE_DOT1D_TP_PORT_IN_FRAMES":                 18,
	"STAT_TYPE_DOT1D_TP_PORT_OUT_FRAMES":                19,
	"STAT_TYPE_DOT1D_PORT_IN_DISCARDS":                  20,
	"STAT_TYPE_ETHER_STATS_DROP_EVENTS":                 21,
	"STAT_TYPE_ETHER_STATS_MULTICAST_PKTS":              22,
	"STAT_TYPE_ETHER_STATS_BROADCAST_PKTS":              23,
	"STAT_TYPE_ETHER_STATS_UNDERSIZE_PKTS":              24,
	"STAT_TYPE_ETHER_STATS_FRAGMENTS":                   25,
	"STAT_TYPE_ETHER_STATS_PKTS_64_OCTETS":              26,
	"STAT_TYPE_ETHER_STATS_PKTS_65TO127_OCTETS":         27,
	"STAT_TYPE_ETHER_STATS_PKTS_128TO255_OCTETS":        28,
	"STAT_TYPE_ETHER_STATS_PKTS_256TO511_OCTETS":        29,
	"STAT_TYPE_ETHER_STATS_PKTS_512TO1023_OCTETS":       30,
	"STAT_TYPE_ETHER_STATS_PKTS_1024TO1518_OCTETS":      31,
	"STAT_TYPE_ETHER_STATS_OVERSIZE_PKTS":               32,
	"STAT_TYPE_ETHER_RX_OVER_SIZE_PKTS":                 33,
	"STAT_TYPE_ETHER_TX_OVER_SIZE_PKTS":                 34,
	"STAT_TYPE_ETHER_STATS_JABBERS":                     35,
	"STAT_TYPE_ETHER_STATS_OCTETS":                      36,
	"STAT_TYPE_ETHER_STATS_PKTS":                        37,
	"STAT_TYPE_ETHER_STATS_COLLISIONS":                  38,
	"STAT_TYPE_ETHER_STATS_CRC_ALIGN_ERRORS":            39,
	"STAT_TYPE_ETHER_STATS_TX_NO_ERRORS":                40,
	"STAT_TYPE_ETHER_STATS_RX_NO_ERRORS":                41,
	"STAT_TYPE_DOT3_STATS_ALIGNMENT_ERRORS":             42,
	"STAT_TYPE_DOT3_STATS_FCS_ERRORS":                   43,
	"STAT_TYPE_DOT3_STATS_SINGLE_COLLISION_FRAMES":      44,
	"STAT_TYPE_DOT3_STATS_MULTIPLE_COLLISION_FRAMES":    45,
	"STAT_TYPE_DOT3_STATS_SQET_TEST_ERRORS":             46,
	"STAT_TYPE_DOT3_STATS_DEFERRED_TRANSMISSIONS":       47,
	"STAT_TYPE_DOT3_STATS_LATE_COLLISIONS":              48,
	"STAT_TYPE_DOT3_STATS_EXCESSIVE_COLLISIONS":         49,
	"STAT_TYPE_DOT3_STATS_INTERNAL_MAC_TRANSMIT_ERRORS": 50,
	"STAT_TYPE_DOT3_STATS_CARRIER_SENSE_ERRORS":         51,
	"STAT_TYPE_DOT3_STATS_FRAME_TOO_LONGS":              52,
	"STAT_TYPE_DOT3_STATS_INTERNAL_MAC_RECEIVE_ERRORS":  53,
	"STAT_TYPE_DOT3_STATS_SYMBOL_ERRORS":                54,
	"STAT_TYPE_DOT3_CONTROL_IN_UNKNOWN_OP_CODES":        55,
	"STAT_TYPE_DOT3_IN_PAUSE_FRAMES":                    56,
	"STAT_TYPE_DOT3_OUT_PAUSE_FRAMES":                   57,
	"STAT_TYPE_IF_HC_IN_OCTETS":                         58,
	"STAT_TYPE_IF_HC_IN_UCAST_PKTS":                     59,
	"STAT_TYPE_IF_HC_IN_MULTICAST_PKTS":                 60,
	"STAT_TYPE_IF_HC_IN_BROADCAST_PKTS":                 61,
	"STAT_TYPE_IF_HC_OUT_OCTETS":                        62,
	"STAT_TYPE_IF_HC_OUT_UCAST_PKTS":                    63,
	"STAT_TYPE_IF_HC_OUT_MULTICAST_PKTS":                64,
	"STAT_TYPE_IF_HC_OUT_BROADCAST_PKTS":                65,
	"STAT_TYPE_IPV6_IF_STATS_IN_RECEIVES":               66,
	"STAT_TYPE_IPV6_IF_STATS_IN_HDR_ERRORS":             67,
	"STAT_TYPE_IPV6_IF_STATS_IN_ADDR_ERRORS":            68,
	"STAT_TYPE_IPV6_IF_STATS_IN_DISCARDS":               69,
	"STAT_TYPE_IPV6_IF_STATS_OUT_FORW_DATAGRAMS":        70,
	"STAT_TYPE_IPV6_IF_STATS_OUT_DISCARDS":              71,
	"STAT_TYPE_IPV6_IF_STATS_IN_MCAST_PKTS":             72,
	"STAT_TYPE_IPV6_IF_STATS_OUT_MCAST_PKTS":            73,
	"STAT_TYPE_IF_IN_BROADCAST_PKTS":                    74,
	"STAT_TYPE_IF_IN_MULTICAST_PKTS":                    75,
	"STAT_TYPE_IF_OUT_BROADCAST_PKTS":                   76,
	"STAT_TYPE_IF_OUT_MULTICAST_PKTS":                   77,
	"STAT_TYPE_IEEE8021_PFC_REQUESTS":                   78,
	"STAT_TYPE_IEEE8021_PFC_INDICATIONS":                79,
}

func (x StatType) String() string {
	return proto.EnumName(StatType_name, int32(x))
}
func (StatType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type InitRequest struct {
	Unit int64 `protobuf:"varint,1,opt,name=unit" json:"unit,omitempty"`
}

func (m *InitRequest) Reset()                    { *m = InitRequest{} }
func (m *InitRequest) String() string            { return proto.CompactTextString(m) }
func (*InitRequest) ProtoMessage()               {}
func (*InitRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type InitResponse struct {
}

func (m *InitResponse) Reset()                    { *m = InitResponse{} }
func (m *InitResponse) String() string            { return proto.CompactTextString(m) }
func (*InitResponse) ProtoMessage()               {}
func (*InitResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ClearRequest struct {
	Unit int64 `protobuf:"varint,1,opt,name=unit" json:"unit,omitempty"`
}

func (m *ClearRequest) Reset()                    { *m = ClearRequest{} }
func (m *ClearRequest) String() string            { return proto.CompactTextString(m) }
func (*ClearRequest) ProtoMessage()               {}
func (*ClearRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ClearResponse struct {
}

func (m *ClearResponse) Reset()                    { *m = ClearResponse{} }
func (m *ClearResponse) String() string            { return proto.CompactTextString(m) }
func (*ClearResponse) ProtoMessage()               {}
func (*ClearResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type SyncRequest struct {
	Unit int64 `protobuf:"varint,1,opt,name=unit" json:"unit,omitempty"`
}

func (m *SyncRequest) Reset()                    { *m = SyncRequest{} }
func (m *SyncRequest) String() string            { return proto.CompactTextString(m) }
func (*SyncRequest) ProtoMessage()               {}
func (*SyncRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type SyncResponse struct {
}

func (m *SyncResponse) Reset()                    { *m = SyncResponse{} }
func (m *SyncResponse) String() string            { return proto.CompactTextString(m) }
func (*SyncResponse) ProtoMessage()               {}
func (*SyncResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type GetRequest struct {
	Unit int64    `protobuf:"varint,1,opt,name=unit" json:"unit,omitempty"`
	Port int64    `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	Type StatType `protobuf:"varint,3,opt,name=type,enum=stat.StatType" json:"type,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type GetResponse struct {
	Value uint64 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type MultiGetRequest struct {
	Unit int64      `protobuf:"varint,1,opt,name=unit" json:"unit,omitempty"`
	Port int64      `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	Type []StatType `protobuf:"varint,3,rep,name=type,enum=stat.StatType" json:"type,omitempty"`
}

func (m *MultiGetRequest) Reset()                    { *m = MultiGetRequest{} }
func (m *MultiGetRequest) String() string            { return proto.CompactTextString(m) }
func (*MultiGetRequest) ProtoMessage()               {}
func (*MultiGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type MultiGetResponse struct {
	Value []uint64 `protobuf:"varint,1,rep,name=value" json:"value,omitempty"`
}

func (m *MultiGetResponse) Reset()                    { *m = MultiGetResponse{} }
func (m *MultiGetResponse) String() string            { return proto.CompactTextString(m) }
func (*MultiGetResponse) ProtoMessage()               {}
func (*MultiGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func init() {
	proto.RegisterType((*InitRequest)(nil), "stat.InitRequest")
	proto.RegisterType((*InitResponse)(nil), "stat.InitResponse")
	proto.RegisterType((*ClearRequest)(nil), "stat.ClearRequest")
	proto.RegisterType((*ClearResponse)(nil), "stat.ClearResponse")
	proto.RegisterType((*SyncRequest)(nil), "stat.SyncRequest")
	proto.RegisterType((*SyncResponse)(nil), "stat.SyncResponse")
	proto.RegisterType((*GetRequest)(nil), "stat.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "stat.GetResponse")
	proto.RegisterType((*MultiGetRequest)(nil), "stat.MultiGetRequest")
	proto.RegisterType((*MultiGetResponse)(nil), "stat.MultiGetResponse")
	proto.RegisterEnum("stat.StatType", StatType_name, StatType_value)
}

func init() { proto.RegisterFile("stat.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1122 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x97, 0xfb, 0x53, 0xdb, 0x46,
	0x10, 0xc7, 0x9b, 0x40, 0xd2, 0x74, 0x21, 0xa0, 0x5e, 0xf3, 0x30, 0x6f, 0x30, 0x8f, 0x00, 0x49,
	0xc0, 0x16, 0x98, 0xd0, 0x77, 0x65, 0xe9, 0x6c, 0x14, 0x24, 0x9d, 0xb8, 0x3b, 0x13, 0xe8, 0x4c,
	0xe7, 0x86, 0x76, 0xfc, 0x43, 0x66, 0x18, 0xa0, 0x60, 0x3a, 0x93, 0x3f, 0xba, 0xff, 0x43, 0x4f,
	0xb2, 0x5e, 0xa7, 0x87, 0xf9, 0xa1, 0xbf, 0x30, 0x96, 0xee, 0xb3, 0xbb, 0xdf, 0xbd, 0xdb, 0x5b,
	0x2d, 0x00, 0x77, 0x83, 0x8b, 0xc1, 0xce, 0xcd, 0xed, 0xf5, 0xe0, 0x1a, 0x8d, 0x07, 0xbf, 0xeb,
	0x2b, 0x30, 0x61, 0x5f, 0x7d, 0x1e, 0xd0, 0xfe, 0xdf, 0xf7, 0xfd, 0xbb, 0x01, 0x42, 0x30, 0x7e,
	0x2f, 0x1f, 0x6b, 0x8f, 0x96, 0x1f, 0x6d, 0x8e, 0xd1, 0xf0, 0x77, 0x7d, 0x0a, 0x26, 0x87, 0xc8,
	0xdd, 0xcd, 0xf5, 0xd5, 0x5d, 0xbf, 0x5e, 0x87, 0x49, 0xf3, 0xb2, 0x7f, 0x71, 0x3b, 0xca, 0x66,
	0x1a, 0x9e, 0x47, 0x4c, 0x64, 0x24, 0xe3, 0xb0, 0x2f, 0x57, 0x7f, 0x3d, 0x10, 0x67, 0x88, 0x44,
	0x26, 0x67, 0x00, 0xdd, 0xfe, 0x28, 0x65, 0xc1, 0xbb, 0x9b, 0xeb, 0xdb, 0x41, 0xed, 0xf1, 0xf0,
	0x5d, 0xf0, 0x1b, 0xd5, 0x61, 0x7c, 0xf0, 0xe5, 0xa6, 0x5f, 0x1b, 0x93, 0xef, 0xa6, 0xf4, 0xa9,
	0x9d, 0x30, 0x63, 0x26, 0xff, 0x70, 0xf9, 0x96, 0x86, 0x6b, 0xf5, 0x55, 0x98, 0x08, 0x3d, 0x0f,
	0x03, 0xa1, 0x17, 0xf0, 0xe4, 0x9f, 0x8b, 0xcb, 0xfb, 0x7e, 0xe8, 0x7b, 0x9c, 0x0e, 0x1f, 0xea,
	0x7f, 0xc0, 0xb4, 0x7b, 0x7f, 0x39, 0xf8, 0xfc, 0xbf, 0x34, 0x8c, 0x55, 0x6a, 0xd8, 0x04, 0x2d,
	0x75, 0x5f, 0x14, 0x32, 0x96, 0x08, 0xd9, 0xfe, 0xb7, 0x06, 0xcf, 0x62, 0x63, 0x34, 0x0b, 0xaf,
	0x18, 0x37, 0xb8, 0xe0, 0xe7, 0x3e, 0x16, 0x76, 0x47, 0xd8, 0x9e, 0x20, 0x26, 0xc7, 0x9c, 0x69,
	0x5f, 0xa1, 0x45, 0x98, 0xcd, 0xaf, 0xf5, 0x4c, 0x83, 0x71, 0xe1, 0x1f, 0xcb, 0xf5, 0x47, 0x68,
	0x09, 0xe6, 0xf2, 0xeb, 0x5e, 0x06, 0x78, 0x8c, 0xe6, 0xa1, 0x96, 0x07, 0x2c, 0x9b, 0x99, 0x06,
	0xb5, 0x98, 0x36, 0x56, 0x16, 0x1a, 0x53, 0x4a, 0x28, 0xd3, 0xc6, 0x65, 0xc6, 0x8b, 0x85, 0xd0,
	0xde, 0xb1, 0x47, 0x3e, 0x79, 0xc2, 0xa7, 0x84, 0x13, 0xa6, 0x3d, 0x41, 0x73, 0xf0, 0x5a, 0x61,
	0x48, 0x8f, 0xc7, 0xda, 0x9f, 0x16, 0xb4, 0x05, 0x8b, 0x19, 0x6d, 0x5f, 0xa3, 0x65, 0x98, 0x2f,
	0x00, 0x59, 0xf5, 0xcf, 0xd0, 0x02, 0xcc, 0x14, 0x88, 0x44, 0xfe, 0x37, 0xa5, 0xe1, 0x23, 0xfd,
	0x80, 0x66, 0xe0, 0x65, 0x61, 0xf1, 0xc4, 0xc1, 0x9e, 0x36, 0x91, 0xdb, 0x14, 0x3f, 0x48, 0x8d,
	0x62, 0x13, 0xdb, 0xa7, 0x98, 0x69, 0x93, 0xb9, 0x3d, 0x0f, 0x57, 0x8f, 0x2c, 0x1a, 0x3b, 0x7e,
	0x9e, 0xcb, 0xcb, 0x17, 0x1d, 0x42, 0x3f, 0x09, 0xcb, 0xe0, 0x46, 0x97, 0x1a, 0x2e, 0xd3, 0xa6,
	0xca, 0xdc, 0x27, 0xa2, 0xa7, 0x51, 0x0b, 0x9a, 0xe9, 0xaa, 0x45, 0x78, 0xd3, 0x12, 0x6d, 0x83,
	0x61, 0xe1, 0x13, 0x2a, 0x93, 0xc3, 0x8e, 0x71, 0x2e, 0xf0, 0x99, 0x89, 0xb1, 0x85, 0xad, 0xd4,
	0x4c, 0x43, 0x7b, 0xb0, 0x5b, 0x6d, 0xe6, 0xf2, 0x5e, 0x89, 0xd1, 0xb7, 0x68, 0x1d, 0x56, 0xf2,
	0x46, 0xdc, 0x1f, 0x9a, 0x48, 0x51, 0x1d, 0xa9, 0x57, 0x66, 0x8c, 0xd0, 0x06, 0xd4, 0xab, 0xb0,
	0x60, 0xd7, 0x22, 0xee, 0x3b, 0xb4, 0x06, 0xcb, 0x79, 0x2e, 0xf6, 0x95, 0x04, 0x7d, 0xa1, 0x06,
	0xc5, 0xfc, 0x08, 0x53, 0x11, 0x3c, 0x33, 0x61, 0x51, 0xe2, 0x0b, 0x7c, 0x8a, 0x3d, 0x79, 0xb6,
	0x2f, 0xd1, 0x26, 0xac, 0x95, 0x63, 0x6e, 0xcf, 0xe1, 0x76, 0x5a, 0x05, 0xaf, 0xaa, 0xc9, 0x36,
	0x25, 0x86, 0x95, 0x92, 0xaf, 0xab, 0xc9, 0x9e, 0x67, 0x61, 0xca, 0xec, 0xdf, 0xf1, 0x90, 0xac,
	0xa1, 0x55, 0x58, 0x2a, 0x27, 0x65, 0xb2, 0x5d, 0x37, 0x94, 0x38, 0x53, 0xed, 0x2e, 0x70, 0x22,
	0x0e, 0xf6, 0xe3, 0x5a, 0x9f, 0x45, 0xef, 0x61, 0x6b, 0x14, 0xd9, 0xe2, 0xa4, 0xa9, 0x7f, 0x88,
	0xf1, 0x39, 0xb4, 0x03, 0xdb, 0x23, 0xf0, 0xa6, 0x7e, 0xc8, 0x89, 0xde, 0x6a, 0xc5, 0xfc, 0xfc,
	0x03, 0xbc, 0xde, 0x3a, 0xe0, 0xa4, 0xd5, 0x6c, 0xc6, 0xfc, 0x02, 0xda, 0x85, 0xb7, 0x23, 0xf8,
	0x56, 0x53, 0x97, 0x7a, 0x1a, 0xfa, 0x5e, 0x6c, 0xb0, 0x88, 0x1a, 0xf0, 0x6e, 0x94, 0xa0, 0x86,
	0xbe, 0x2f, 0x2d, 0x5a, 0xcd, 0xc3, 0xd8, 0x62, 0x09, 0xbd, 0x81, 0xd5, 0x72, 0x0b, 0x72, 0x9a,
	0xdd, 0xe9, 0xe5, 0xb2, 0x72, 0xa0, 0x67, 0x21, 0x25, 0x52, 0x6c, 0xa5, 0x0c, 0xe3, 0x05, 0xac,
	0x8e, 0x56, 0x60, 0xa1, 0x3c, 0xec, 0x47, 0xa3, 0xdd, 0x96, 0x81, 0xb5, 0x55, 0xb5, 0xad, 0x28,
	0xca, 0x86, 0xda, 0xd7, 0xd4, 0x1b, 0x9e, 0xcf, 0x56, 0x5b, 0x57, 0xeb, 0x3c, 0xbb, 0x6e, 0x12,
	0xc7, 0xb1, 0x99, 0x4d, 0x3c, 0xa6, 0x6d, 0xa0, 0x6d, 0xd8, 0xa8, 0xa0, 0xa8, 0x29, 0x0c, 0xc7,
	0xee, 0x26, 0xcd, 0xf4, 0x8d, 0x7a, 0xc3, 0xb2, 0xac, 0xcc, 0xd1, 0x23, 0x31, 0xb7, 0x59, 0xcd,
	0xd1, 0x2c, 0xb7, 0x85, 0xb6, 0x60, 0x5d, 0xb9, 0x89, 0x7b, 0x11, 0x16, 0x86, 0x0d, 0xca, 0x37,
	0x46, 0xb7, 0xd5, 0x4a, 0xcf, 0xa0, 0x1d, 0x93, 0xc5, 0xd0, 0x5b, 0xf5, 0xfc, 0x33, 0x10, 0xb3,
	0xbd, 0xae, 0x83, 0xd3, 0xbc, 0xe3, 0x5e, 0xf0, 0x0e, 0xe9, 0xb0, 0x53, 0x6a, 0x11, 0xde, 0x5e,
	0xbf, 0xcc, 0xe6, 0x7d, 0xa5, 0x6a, 0x76, 0x82, 0xe5, 0x4b, 0xcc, 0x12, 0xd5, 0x3b, 0x6a, 0x05,
	0x67, 0x50, 0x0b, 0x77, 0x24, 0x23, 0xbb, 0x1c, 0xa7, 0x86, 0xc7, 0x5c, 0x9b, 0x0d, 0x4f, 0x63,
	0x57, 0xbd, 0xab, 0x19, 0x03, 0xc7, 0xe0, 0x38, 0x7b, 0x6e, 0x0d, 0xf5, 0xae, 0x66, 0xc8, 0xa0,
	0x81, 0x4a, 0x77, 0xa7, 0x0a, 0xde, 0x2c, 0xf4, 0xeb, 0x18, 0xb7, 0x3d, 0x8e, 0xa9, 0x67, 0x38,
	0xc2, 0x35, 0xcc, 0x58, 0x4d, 0x92, 0x80, 0x5e, 0x19, 0x45, 0x76, 0x49, 0x6a, 0x07, 0xc7, 0x8a,
	0x3d, 0xd9, 0xbd, 0x23, 0x7c, 0xaf, 0x52, 0x7e, 0xb8, 0x77, 0x82, 0x13, 0x22, 0x1c, 0xe2, 0x75,
	0x99, 0xb6, 0x8f, 0xf6, 0xa1, 0xf1, 0xb0, 0x9e, 0xe8, 0x73, 0x16, 0xfb, 0x6f, 0x15, 0x5a, 0x7c,
	0xb2, 0xf5, 0xe7, 0x6e, 0x9b, 0x38, 0x31, 0x77, 0xa0, 0x76, 0x9a, 0x90, 0x33, 0x89, 0xc7, 0xa9,
	0x44, 0x32, 0xdf, 0x7f, 0xd9, 0xc6, 0x4d, 0x62, 0xc9, 0x23, 0xfd, 0xa0, 0x4e, 0x09, 0x21, 0x2f,
	0x39, 0xdf, 0xe8, 0xc9, 0xdc, 0xa2, 0x63, 0x3f, 0x2c, 0xa9, 0xc0, 0xe0, 0xab, 0xa2, 0x40, 0xdf,
	0x17, 0x3e, 0xf5, 0x47, 0x66, 0x66, 0x10, 0xfa, 0x41, 0xbd, 0xf7, 0xc9, 0x72, 0x66, 0x58, 0xf8,
	0x51, 0xed, 0x20, 0x09, 0x92, 0xfb, 0x9a, 0xfc, 0x54, 0x81, 0xe5, 0x3e, 0x25, 0x3f, 0x17, 0x26,
	0x2f, 0x89, 0x65, 0xa6, 0x9b, 0x5f, 0x0a, 0xe3, 0x51, 0xb4, 0x9e, 0x51, 0xf4, 0xab, 0xba, 0xe9,
	0x29, 0x93, 0x93, 0xf4, 0x5b, 0x15, 0x97, 0xd3, 0x64, 0xa8, 0x3d, 0xd7, 0xf6, 0x4f, 0x0f, 0x02,
	0x38, 0x3e, 0xfd, 0x74, 0x84, 0x69, 0xab, 0x17, 0xad, 0x00, 0x66, 0xa6, 0x19, 0x53, 0xed, 0x62,
	0x05, 0xd4, 0xb0, 0x52, 0xd6, 0x7a, 0x20, 0x7e, 0x32, 0x02, 0x60, 0xb5, 0x8a, 0x54, 0x30, 0x1c,
	0x28, 0xd4, 0x89, 0xa9, 0xa3, 0x56, 0x7f, 0x91, 0x4f, 0x3c, 0x77, 0x1f, 0xc8, 0xcc, 0x4d, 0x77,
	0xeb, 0x68, 0x54, 0x66, 0xe1, 0x29, 0xa4, 0xac, 0x5d, 0x36, 0xec, 0xe6, 0x76, 0xff, 0x63, 0x19,
	0x93, 0x3b, 0xc9, 0x63, 0xb5, 0xd4, 0xa3, 0xa1, 0x33, 0xe7, 0xc8, 0x29, 0x85, 0x72, 0x9e, 0xdc,
	0x1c, 0x84, 0x31, 0x3e, 0x6c, 0xe8, 0x4d, 0xe1, 0x77, 0x82, 0xbb, 0x7d, 0xd2, 0x93, 0xbd, 0x92,
	0x69, 0x5e, 0xae, 0x70, 0xb2, 0x90, 0xed, 0x59, 0xd2, 0x1d, 0x0f, 0x7b, 0x18, 0xf9, 0xf3, 0x69,
	0xf8, 0xff, 0xe1, 0xde, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x10, 0x09, 0xc7, 0x2d, 0x0e,
	0x00, 0x00,
}
