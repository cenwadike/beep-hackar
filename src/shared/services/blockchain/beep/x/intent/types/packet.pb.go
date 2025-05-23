// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: beep/intent/packet.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type IntentPacketData struct {
	// Types that are valid to be assigned to Packet:
	//	*IntentPacketData_NoData
	//	*IntentPacketData_IntentPacketPacket
	Packet isIntentPacketData_Packet `protobuf_oneof:"packet"`
}

func (m *IntentPacketData) Reset()         { *m = IntentPacketData{} }
func (m *IntentPacketData) String() string { return proto.CompactTextString(m) }
func (*IntentPacketData) ProtoMessage()    {}
func (*IntentPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c832f5c48d30ca2, []int{0}
}
func (m *IntentPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IntentPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IntentPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IntentPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntentPacketData.Merge(m, src)
}
func (m *IntentPacketData) XXX_Size() int {
	return m.Size()
}
func (m *IntentPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_IntentPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_IntentPacketData proto.InternalMessageInfo

type isIntentPacketData_Packet interface {
	isIntentPacketData_Packet()
	MarshalTo([]byte) (int, error)
	Size() int
}

type IntentPacketData_NoData struct {
	NoData *NoData `protobuf:"bytes,1,opt,name=noData,proto3,oneof" json:"noData,omitempty"`
}
type IntentPacketData_IntentPacketPacket struct {
	IntentPacketPacket *IntentPacketPacketData `protobuf:"bytes,2,opt,name=intentPacketPacket,proto3,oneof" json:"intentPacketPacket,omitempty"`
}

func (*IntentPacketData_NoData) isIntentPacketData_Packet()             {}
func (*IntentPacketData_IntentPacketPacket) isIntentPacketData_Packet() {}

func (m *IntentPacketData) GetPacket() isIntentPacketData_Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

func (m *IntentPacketData) GetNoData() *NoData {
	if x, ok := m.GetPacket().(*IntentPacketData_NoData); ok {
		return x.NoData
	}
	return nil
}

func (m *IntentPacketData) GetIntentPacketPacket() *IntentPacketPacketData {
	if x, ok := m.GetPacket().(*IntentPacketData_IntentPacketPacket); ok {
		return x.IntentPacketPacket
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*IntentPacketData) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*IntentPacketData_NoData)(nil),
		(*IntentPacketData_IntentPacketPacket)(nil),
	}
}

type NoData struct {
}

func (m *NoData) Reset()         { *m = NoData{} }
func (m *NoData) String() string { return proto.CompactTextString(m) }
func (*NoData) ProtoMessage()    {}
func (*NoData) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c832f5c48d30ca2, []int{1}
}
func (m *NoData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NoData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NoData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NoData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoData.Merge(m, src)
}
func (m *NoData) XXX_Size() int {
	return m.Size()
}
func (m *NoData) XXX_DiscardUnknown() {
	xxx_messageInfo_NoData.DiscardUnknown(m)
}

var xxx_messageInfo_NoData proto.InternalMessageInfo

// IntentPacketPacketData defines a struct for the packet payload
type IntentPacketPacketData struct {
	IntentType  string `protobuf:"bytes,1,opt,name=intentType,proto3" json:"intentType,omitempty"`
	Memo        string `protobuf:"bytes,2,opt,name=memo,proto3" json:"memo,omitempty"`
	TargetChain string `protobuf:"bytes,3,opt,name=targetChain,proto3" json:"targetChain,omitempty"`
	MinOutput   uint64 `protobuf:"varint,4,opt,name=minOutput,proto3" json:"minOutput,omitempty"`
	Creator     string `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	Amount      uint64 `protobuf:"varint,6,opt,name=amount,proto3" json:"amount,omitempty"`
	InputToken  string `protobuf:"bytes,7,opt,name=inputToken,proto3" json:"inputToken,omitempty"`
	OutputToken string `protobuf:"bytes,8,opt,name=outputToken,proto3" json:"outputToken,omitempty"`
}

func (m *IntentPacketPacketData) Reset()         { *m = IntentPacketPacketData{} }
func (m *IntentPacketPacketData) String() string { return proto.CompactTextString(m) }
func (*IntentPacketPacketData) ProtoMessage()    {}
func (*IntentPacketPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c832f5c48d30ca2, []int{2}
}
func (m *IntentPacketPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IntentPacketPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IntentPacketPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IntentPacketPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntentPacketPacketData.Merge(m, src)
}
func (m *IntentPacketPacketData) XXX_Size() int {
	return m.Size()
}
func (m *IntentPacketPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_IntentPacketPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_IntentPacketPacketData proto.InternalMessageInfo

func (m *IntentPacketPacketData) GetIntentType() string {
	if m != nil {
		return m.IntentType
	}
	return ""
}

func (m *IntentPacketPacketData) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *IntentPacketPacketData) GetTargetChain() string {
	if m != nil {
		return m.TargetChain
	}
	return ""
}

func (m *IntentPacketPacketData) GetMinOutput() uint64 {
	if m != nil {
		return m.MinOutput
	}
	return 0
}

func (m *IntentPacketPacketData) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *IntentPacketPacketData) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *IntentPacketPacketData) GetInputToken() string {
	if m != nil {
		return m.InputToken
	}
	return ""
}

func (m *IntentPacketPacketData) GetOutputToken() string {
	if m != nil {
		return m.OutputToken
	}
	return ""
}

// IntentPacketPacketAck defines a struct for the packet acknowledgment
type IntentPacketPacketAck struct {
	IntentId int32 `protobuf:"varint,1,opt,name=intentId,proto3" json:"intentId,omitempty"`
}

func (m *IntentPacketPacketAck) Reset()         { *m = IntentPacketPacketAck{} }
func (m *IntentPacketPacketAck) String() string { return proto.CompactTextString(m) }
func (*IntentPacketPacketAck) ProtoMessage()    {}
func (*IntentPacketPacketAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c832f5c48d30ca2, []int{3}
}
func (m *IntentPacketPacketAck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IntentPacketPacketAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IntentPacketPacketAck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IntentPacketPacketAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntentPacketPacketAck.Merge(m, src)
}
func (m *IntentPacketPacketAck) XXX_Size() int {
	return m.Size()
}
func (m *IntentPacketPacketAck) XXX_DiscardUnknown() {
	xxx_messageInfo_IntentPacketPacketAck.DiscardUnknown(m)
}

var xxx_messageInfo_IntentPacketPacketAck proto.InternalMessageInfo

func (m *IntentPacketPacketAck) GetIntentId() int32 {
	if m != nil {
		return m.IntentId
	}
	return 0
}

func init() {
	proto.RegisterType((*IntentPacketData)(nil), "beep.intent.IntentPacketData")
	proto.RegisterType((*NoData)(nil), "beep.intent.NoData")
	proto.RegisterType((*IntentPacketPacketData)(nil), "beep.intent.IntentPacketPacketData")
	proto.RegisterType((*IntentPacketPacketAck)(nil), "beep.intent.IntentPacketPacketAck")
}

func init() { proto.RegisterFile("beep/intent/packet.proto", fileDescriptor_5c832f5c48d30ca2) }

var fileDescriptor_5c832f5c48d30ca2 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xc1, 0x4e, 0xf2, 0x40,
	0x10, 0x6e, 0xf9, 0xa1, 0x94, 0xe1, 0xf2, 0x67, 0x89, 0x64, 0x63, 0xcc, 0x86, 0xd4, 0x0b, 0x17,
	0x4a, 0x22, 0x4f, 0x20, 0x7a, 0x90, 0x8b, 0x9a, 0x06, 0x2f, 0xde, 0x16, 0xdc, 0x68, 0x43, 0xba,
	0xbb, 0xa9, 0xd3, 0x44, 0x6e, 0x3e, 0x82, 0xcf, 0xe0, 0xd3, 0x78, 0xe4, 0xe8, 0xd1, 0xc0, 0x8b,
	0x18, 0x66, 0xa9, 0xd6, 0xc0, 0xa5, 0x99, 0xf9, 0xe6, 0xfb, 0xbe, 0x7e, 0xb3, 0x19, 0xe0, 0x33,
	0xa5, 0xec, 0x30, 0xd5, 0xa8, 0x34, 0x0e, 0xad, 0x9c, 0x2f, 0x14, 0xc6, 0x36, 0x37, 0x68, 0x58,
	0x7b, 0x3b, 0x89, 0xdd, 0x24, 0x7a, 0xf7, 0xe1, 0xff, 0x84, 0xca, 0x5b, 0xe2, 0x5c, 0x4a, 0x94,
	0x6c, 0x00, 0x81, 0x36, 0xdb, 0x8a, 0xfb, 0x3d, 0xbf, 0xdf, 0x3e, 0xeb, 0xc4, 0x15, 0x49, 0x7c,
	0x4d, 0xa3, 0x2b, 0x2f, 0xd9, 0x91, 0xd8, 0x1d, 0xb0, 0xb4, 0x62, 0xe1, 0xbe, 0xbc, 0x46, 0xd2,
	0xd3, 0x3f, 0xd2, 0xc9, 0x1e, 0x6d, 0x67, 0x75, 0xc0, 0x60, 0x1c, 0x42, 0xe0, 0x72, 0x47, 0x21,
	0x04, 0xee, 0xa7, 0xd1, 0x6b, 0x0d, 0xba, 0x87, 0x4d, 0x98, 0x00, 0x70, 0x26, 0xd3, 0xa5, 0x55,
	0x14, 0xbc, 0x95, 0x54, 0x10, 0xc6, 0xa0, 0x9e, 0xa9, 0xcc, 0x50, 0xae, 0x56, 0x42, 0x35, 0xeb,
	0x41, 0x1b, 0x65, 0xfe, 0xa8, 0xf0, 0xe2, 0x49, 0xa6, 0x9a, 0xff, 0xa3, 0x51, 0x15, 0x62, 0x27,
	0xd0, 0xca, 0x52, 0x7d, 0x53, 0xa0, 0x2d, 0x90, 0xd7, 0x7b, 0x7e, 0xbf, 0x9e, 0xfc, 0x02, 0x8c,
	0x43, 0x73, 0x9e, 0x2b, 0x89, 0x26, 0xe7, 0x0d, 0xd2, 0x96, 0x2d, 0xeb, 0x42, 0x20, 0x33, 0x53,
	0x68, 0xe4, 0x01, 0x89, 0x76, 0x9d, 0x4b, 0x69, 0x0b, 0x9c, 0x9a, 0x85, 0xd2, 0xbc, 0x59, 0xa6,
	0x2c, 0x91, 0x6d, 0x22, 0x43, 0xde, 0x8e, 0x10, 0xba, 0x44, 0x15, 0x28, 0x1a, 0xc1, 0xd1, 0xfe,
	0x0b, 0x9c, 0xcf, 0x17, 0xec, 0x18, 0x42, 0xb7, 0xee, 0xe4, 0x81, 0xd6, 0x6f, 0x24, 0x3f, 0xfd,
	0x78, 0xf0, 0xb1, 0x16, 0xfe, 0x6a, 0x2d, 0xfc, 0xaf, 0xb5, 0xf0, 0xdf, 0x36, 0xc2, 0x5b, 0x6d,
	0x84, 0xf7, 0xb9, 0x11, 0xde, 0x7d, 0x87, 0xee, 0xe4, 0xa5, 0xbc, 0x14, 0x5c, 0x5a, 0xf5, 0x3c,
	0x0b, 0xe8, 0x52, 0x46, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x42, 0xe7, 0xe8, 0x70, 0x45, 0x02,
	0x00, 0x00,
}

func (m *IntentPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IntentPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IntentPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Packet != nil {
		{
			size := m.Packet.Size()
			i -= size
			if _, err := m.Packet.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *IntentPacketData_NoData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IntentPacketData_NoData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NoData != nil {
		{
			size, err := m.NoData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPacket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *IntentPacketData_IntentPacketPacket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IntentPacketData_IntentPacketPacket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.IntentPacketPacket != nil {
		{
			size, err := m.IntentPacketPacket.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPacket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *NoData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NoData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NoData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *IntentPacketPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IntentPacketPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IntentPacketPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OutputToken) > 0 {
		i -= len(m.OutputToken)
		copy(dAtA[i:], m.OutputToken)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.OutputToken)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.InputToken) > 0 {
		i -= len(m.InputToken)
		copy(dAtA[i:], m.InputToken)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.InputToken)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Amount != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x2a
	}
	if m.MinOutput != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.MinOutput))
		i--
		dAtA[i] = 0x20
	}
	if len(m.TargetChain) > 0 {
		i -= len(m.TargetChain)
		copy(dAtA[i:], m.TargetChain)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.TargetChain)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Memo) > 0 {
		i -= len(m.Memo)
		copy(dAtA[i:], m.Memo)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Memo)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.IntentType) > 0 {
		i -= len(m.IntentType)
		copy(dAtA[i:], m.IntentType)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.IntentType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IntentPacketPacketAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IntentPacketPacketAck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IntentPacketPacketAck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IntentId != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.IntentId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPacket(dAtA []byte, offset int, v uint64) int {
	offset -= sovPacket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IntentPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Packet != nil {
		n += m.Packet.Size()
	}
	return n
}

func (m *IntentPacketData_NoData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NoData != nil {
		l = m.NoData.Size()
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}
func (m *IntentPacketData_IntentPacketPacket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IntentPacketPacket != nil {
		l = m.IntentPacketPacket.Size()
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}
func (m *NoData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *IntentPacketPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IntentType)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.Memo)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.TargetChain)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	if m.MinOutput != 0 {
		n += 1 + sovPacket(uint64(m.MinOutput))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovPacket(uint64(m.Amount))
	}
	l = len(m.InputToken)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.OutputToken)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}

func (m *IntentPacketPacketAck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IntentId != 0 {
		n += 1 + sovPacket(uint64(m.IntentId))
	}
	return n
}

func sovPacket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPacket(x uint64) (n int) {
	return sovPacket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IntentPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IntentPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IntentPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &NoData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Packet = &IntentPacketData_NoData{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntentPacketPacket", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &IntentPacketPacketData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Packet = &IntentPacketData_IntentPacketPacket{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NoData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NoData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NoData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IntentPacketPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IntentPacketPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IntentPacketPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntentType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IntentType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Memo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TargetChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinOutput", wireType)
			}
			m.MinOutput = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinOutput |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InputToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InputToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutputToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OutputToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IntentPacketPacketAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IntentPacketPacketAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IntentPacketPacketAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntentId", wireType)
			}
			m.IntentId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IntentId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPacket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPacket
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPacket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPacket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPacket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPacket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPacket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPacket = fmt.Errorf("proto: unexpected end of group")
)
