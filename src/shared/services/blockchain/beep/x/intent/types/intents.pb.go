// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: beep/intent/intents.proto

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

type Intents struct {
	Id           uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator      string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	ActionType   string `protobuf:"bytes,3,opt,name=actionType,proto3" json:"actionType,omitempty"`
	Memo         string `protobuf:"bytes,4,opt,name=memo,proto3" json:"memo,omitempty"`
	TargetChain  string `protobuf:"bytes,5,opt,name=targetChain,proto3" json:"targetChain,omitempty"`
	MinOutput    uint64 `protobuf:"varint,6,opt,name=minOutput,proto3" json:"minOutput,omitempty"`
	ExpiryHeight uint64 `protobuf:"varint,7,opt,name=expiryHeight,proto3" json:"expiryHeight,omitempty"`
	Status       string `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	Executor     string `protobuf:"bytes,9,opt,name=executor,proto3" json:"executor,omitempty"`
}

func (m *Intents) Reset()         { *m = Intents{} }
func (m *Intents) String() string { return proto.CompactTextString(m) }
func (*Intents) ProtoMessage()    {}
func (*Intents) Descriptor() ([]byte, []int) {
	return fileDescriptor_d27fe3ec0d909c43, []int{0}
}
func (m *Intents) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Intents) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Intents.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Intents) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Intents.Merge(m, src)
}
func (m *Intents) XXX_Size() int {
	return m.Size()
}
func (m *Intents) XXX_DiscardUnknown() {
	xxx_messageInfo_Intents.DiscardUnknown(m)
}

var xxx_messageInfo_Intents proto.InternalMessageInfo

func (m *Intents) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Intents) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Intents) GetActionType() string {
	if m != nil {
		return m.ActionType
	}
	return ""
}

func (m *Intents) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *Intents) GetTargetChain() string {
	if m != nil {
		return m.TargetChain
	}
	return ""
}

func (m *Intents) GetMinOutput() uint64 {
	if m != nil {
		return m.MinOutput
	}
	return 0
}

func (m *Intents) GetExpiryHeight() uint64 {
	if m != nil {
		return m.ExpiryHeight
	}
	return 0
}

func (m *Intents) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Intents) GetExecutor() string {
	if m != nil {
		return m.Executor
	}
	return ""
}

func init() {
	proto.RegisterType((*Intents)(nil), "beep.intent.Intents")
}

func init() { proto.RegisterFile("beep/intent/intents.proto", fileDescriptor_d27fe3ec0d909c43) }

var fileDescriptor_d27fe3ec0d909c43 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xe3, 0x10, 0x92, 0xe6, 0x8a, 0x18, 0x0e, 0x09, 0x19, 0x84, 0xac, 0xa8, 0x53, 0x17,
	0xca, 0xc0, 0x1b, 0xc0, 0x02, 0x13, 0x52, 0xc5, 0xc4, 0x96, 0xa6, 0xa7, 0xd6, 0x43, 0x12, 0x2b,
	0xb9, 0x48, 0xc9, 0x5b, 0xf0, 0x58, 0x8c, 0x1d, 0x19, 0x51, 0xf2, 0x20, 0xa0, 0x5e, 0x5b, 0x28,
	0x93, 0xfd, 0x7f, 0x9f, 0x65, 0xdd, 0xfd, 0x70, 0xb5, 0x20, 0x72, 0x77, 0xb6, 0x60, 0x2a, 0x78,
	0x7f, 0xd4, 0x33, 0x57, 0x95, 0x5c, 0xe2, 0x78, 0xab, 0x66, 0x3b, 0x36, 0xf9, 0x56, 0x10, 0x3d,
	0xef, 0x34, 0x9e, 0x83, 0x6f, 0x97, 0x5a, 0x25, 0x6a, 0x1a, 0xcc, 0x7d, 0xbb, 0x44, 0x0d, 0x51,
	0x56, 0x51, 0xca, 0x65, 0xa5, 0xfd, 0x44, 0x4d, 0xe3, 0xf9, 0x21, 0xa2, 0x01, 0x48, 0x33, 0xb6,
	0x65, 0xf1, 0xda, 0x39, 0xd2, 0x27, 0x22, 0x8f, 0x08, 0x22, 0x04, 0x39, 0xe5, 0xa5, 0x0e, 0xc4,
	0xc8, 0x1d, 0x13, 0x18, 0x73, 0x5a, 0xad, 0x88, 0x1f, 0xd7, 0xa9, 0x2d, 0xf4, 0xa9, 0xa8, 0x63,
	0x84, 0x37, 0x10, 0xe7, 0xb6, 0x78, 0x69, 0xd8, 0x35, 0xac, 0x43, 0x19, 0xe3, 0x0f, 0xe0, 0x04,
	0xce, 0xa8, 0x75, 0xb6, 0xea, 0x9e, 0xc8, 0xae, 0xd6, 0xac, 0x23, 0x79, 0xf0, 0x8f, 0xe1, 0x25,
	0x84, 0x35, 0xa7, 0xdc, 0xd4, 0x7a, 0x24, 0xdf, 0xef, 0x13, 0x5e, 0xc3, 0x88, 0x5a, 0xca, 0x9a,
	0xed, 0x2a, 0xb1, 0x98, 0xdf, 0xfc, 0x70, 0xfb, 0xd1, 0x1b, 0xb5, 0xe9, 0x8d, 0xfa, 0xea, 0x8d,
	0x7a, 0x1f, 0x8c, 0xb7, 0x19, 0x8c, 0xf7, 0x39, 0x18, 0xef, 0xed, 0x42, 0x3a, 0x6c, 0x0f, 0x2d,
	0x72, 0xe7, 0xa8, 0x5e, 0x84, 0x52, 0xe2, 0xfd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x50, 0x5f,
	0x7e, 0xe6, 0x61, 0x01, 0x00, 0x00,
}

func (m *Intents) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Intents) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Intents) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Executor) > 0 {
		i -= len(m.Executor)
		copy(dAtA[i:], m.Executor)
		i = encodeVarintIntents(dAtA, i, uint64(len(m.Executor)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintIntents(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x42
	}
	if m.ExpiryHeight != 0 {
		i = encodeVarintIntents(dAtA, i, uint64(m.ExpiryHeight))
		i--
		dAtA[i] = 0x38
	}
	if m.MinOutput != 0 {
		i = encodeVarintIntents(dAtA, i, uint64(m.MinOutput))
		i--
		dAtA[i] = 0x30
	}
	if len(m.TargetChain) > 0 {
		i -= len(m.TargetChain)
		copy(dAtA[i:], m.TargetChain)
		i = encodeVarintIntents(dAtA, i, uint64(len(m.TargetChain)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Memo) > 0 {
		i -= len(m.Memo)
		copy(dAtA[i:], m.Memo)
		i = encodeVarintIntents(dAtA, i, uint64(len(m.Memo)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ActionType) > 0 {
		i -= len(m.ActionType)
		copy(dAtA[i:], m.ActionType)
		i = encodeVarintIntents(dAtA, i, uint64(len(m.ActionType)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintIntents(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintIntents(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintIntents(dAtA []byte, offset int, v uint64) int {
	offset -= sovIntents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Intents) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovIntents(uint64(m.Id))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovIntents(uint64(l))
	}
	l = len(m.ActionType)
	if l > 0 {
		n += 1 + l + sovIntents(uint64(l))
	}
	l = len(m.Memo)
	if l > 0 {
		n += 1 + l + sovIntents(uint64(l))
	}
	l = len(m.TargetChain)
	if l > 0 {
		n += 1 + l + sovIntents(uint64(l))
	}
	if m.MinOutput != 0 {
		n += 1 + sovIntents(uint64(m.MinOutput))
	}
	if m.ExpiryHeight != 0 {
		n += 1 + sovIntents(uint64(m.ExpiryHeight))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovIntents(uint64(l))
	}
	l = len(m.Executor)
	if l > 0 {
		n += 1 + l + sovIntents(uint64(l))
	}
	return n
}

func sovIntents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIntents(x uint64) (n int) {
	return sovIntents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Intents) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIntents
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
			return fmt.Errorf("proto: Intents: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Intents: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIntents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIntents
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
				return ErrInvalidLengthIntents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIntents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIntents
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
				return ErrInvalidLengthIntents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIntents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ActionType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIntents
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
				return ErrInvalidLengthIntents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIntents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Memo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIntents
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
				return ErrInvalidLengthIntents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIntents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TargetChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinOutput", wireType)
			}
			m.MinOutput = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIntents
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
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiryHeight", wireType)
			}
			m.ExpiryHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIntents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpiryHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIntents
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
				return ErrInvalidLengthIntents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIntents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Executor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIntents
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
				return ErrInvalidLengthIntents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIntents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Executor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIntents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIntents
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
func skipIntents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIntents
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
					return 0, ErrIntOverflowIntents
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
					return 0, ErrIntOverflowIntents
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
				return 0, ErrInvalidLengthIntents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIntents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIntents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIntents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIntents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIntents = fmt.Errorf("proto: unexpected end of group")
)
