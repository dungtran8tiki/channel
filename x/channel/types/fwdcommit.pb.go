// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: channel/fwdcommit.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type Fwdcommit struct {
	Index            string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Channelid        string `protobuf:"bytes,2,opt,name=channelid,proto3" json:"channelid,omitempty"`
	Hashcodedest     string `protobuf:"bytes,3,opt,name=hashcodedest,proto3" json:"hashcodedest,omitempty"`
	Timelockreceiver string `protobuf:"bytes,4,opt,name=timelockreceiver,proto3" json:"timelockreceiver,omitempty"`
	Timelocksender   string `protobuf:"bytes,5,opt,name=timelocksender,proto3" json:"timelocksender,omitempty"`
	Hashcodehtlc     string `protobuf:"bytes,6,opt,name=hashcodehtlc,proto3" json:"hashcodehtlc,omitempty"`
}

func (m *Fwdcommit) Reset()         { *m = Fwdcommit{} }
func (m *Fwdcommit) String() string { return proto.CompactTextString(m) }
func (*Fwdcommit) ProtoMessage()    {}
func (*Fwdcommit) Descriptor() ([]byte, []int) {
	return fileDescriptor_54a5573fa7ed7002, []int{0}
}
func (m *Fwdcommit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Fwdcommit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Fwdcommit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Fwdcommit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fwdcommit.Merge(m, src)
}
func (m *Fwdcommit) XXX_Size() int {
	return m.Size()
}
func (m *Fwdcommit) XXX_DiscardUnknown() {
	xxx_messageInfo_Fwdcommit.DiscardUnknown(m)
}

var xxx_messageInfo_Fwdcommit proto.InternalMessageInfo

func (m *Fwdcommit) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Fwdcommit) GetChannelid() string {
	if m != nil {
		return m.Channelid
	}
	return ""
}

func (m *Fwdcommit) GetHashcodedest() string {
	if m != nil {
		return m.Hashcodedest
	}
	return ""
}

func (m *Fwdcommit) GetTimelockreceiver() string {
	if m != nil {
		return m.Timelockreceiver
	}
	return ""
}

func (m *Fwdcommit) GetTimelocksender() string {
	if m != nil {
		return m.Timelocksender
	}
	return ""
}

func (m *Fwdcommit) GetHashcodehtlc() string {
	if m != nil {
		return m.Hashcodehtlc
	}
	return ""
}

func init() {
	proto.RegisterType((*Fwdcommit)(nil), "channel.channel.Fwdcommit")
}

func init() { proto.RegisterFile("channel/fwdcommit.proto", fileDescriptor_54a5573fa7ed7002) }

var fileDescriptor_54a5573fa7ed7002 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0xce, 0x48, 0xcc,
	0xcb, 0x4b, 0xcd, 0xd1, 0x4f, 0x2b, 0x4f, 0x49, 0xce, 0xcf, 0xcd, 0xcd, 0x2c, 0xd1, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x87, 0x4a, 0xe8, 0x41, 0x69, 0xa5, 0xdb, 0x8c, 0x5c, 0x9c, 0x6e,
	0x30, 0x45, 0x42, 0x22, 0x5c, 0xac, 0x99, 0x79, 0x29, 0xa9, 0x15, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0x9c, 0x41, 0x10, 0x8e, 0x90, 0x0c, 0x17, 0x27, 0x54, 0x79, 0x66, 0x8a, 0x04, 0x13, 0x58, 0x06,
	0x21, 0x20, 0xa4, 0xc4, 0xc5, 0x93, 0x91, 0x58, 0x9c, 0x91, 0x9c, 0x9f, 0x92, 0x9a, 0x92, 0x5a,
	0x5c, 0x22, 0xc1, 0x0c, 0x56, 0x80, 0x22, 0x26, 0xa4, 0xc5, 0x25, 0x50, 0x92, 0x99, 0x9b, 0x9a,
	0x93, 0x9f, 0x9c, 0x5d, 0x94, 0x9a, 0x9c, 0x9a, 0x59, 0x96, 0x5a, 0x24, 0xc1, 0x02, 0x56, 0x87,
	0x21, 0x2e, 0xa4, 0xc6, 0xc5, 0x07, 0x13, 0x2b, 0x4e, 0xcd, 0x4b, 0x49, 0x2d, 0x92, 0x60, 0x05,
	0xab, 0x44, 0x13, 0x45, 0xb6, 0x37, 0xa3, 0x24, 0x27, 0x59, 0x82, 0x0d, 0xd5, 0x5e, 0x90, 0x98,
	0x93, 0xe1, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1,
	0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xc1, 0x43, 0xa8, 0x42,
	0x1f, 0xc6, 0x2a, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x07, 0x94, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x14, 0xc8, 0x55, 0xd9, 0x43, 0x01, 0x00, 0x00,
}

func (m *Fwdcommit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Fwdcommit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Fwdcommit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hashcodehtlc) > 0 {
		i -= len(m.Hashcodehtlc)
		copy(dAtA[i:], m.Hashcodehtlc)
		i = encodeVarintFwdcommit(dAtA, i, uint64(len(m.Hashcodehtlc)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Timelocksender) > 0 {
		i -= len(m.Timelocksender)
		copy(dAtA[i:], m.Timelocksender)
		i = encodeVarintFwdcommit(dAtA, i, uint64(len(m.Timelocksender)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Timelockreceiver) > 0 {
		i -= len(m.Timelockreceiver)
		copy(dAtA[i:], m.Timelockreceiver)
		i = encodeVarintFwdcommit(dAtA, i, uint64(len(m.Timelockreceiver)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Hashcodedest) > 0 {
		i -= len(m.Hashcodedest)
		copy(dAtA[i:], m.Hashcodedest)
		i = encodeVarintFwdcommit(dAtA, i, uint64(len(m.Hashcodedest)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Channelid) > 0 {
		i -= len(m.Channelid)
		copy(dAtA[i:], m.Channelid)
		i = encodeVarintFwdcommit(dAtA, i, uint64(len(m.Channelid)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintFwdcommit(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFwdcommit(dAtA []byte, offset int, v uint64) int {
	offset -= sovFwdcommit(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Fwdcommit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovFwdcommit(uint64(l))
	}
	l = len(m.Channelid)
	if l > 0 {
		n += 1 + l + sovFwdcommit(uint64(l))
	}
	l = len(m.Hashcodedest)
	if l > 0 {
		n += 1 + l + sovFwdcommit(uint64(l))
	}
	l = len(m.Timelockreceiver)
	if l > 0 {
		n += 1 + l + sovFwdcommit(uint64(l))
	}
	l = len(m.Timelocksender)
	if l > 0 {
		n += 1 + l + sovFwdcommit(uint64(l))
	}
	l = len(m.Hashcodehtlc)
	if l > 0 {
		n += 1 + l + sovFwdcommit(uint64(l))
	}
	return n
}

func sovFwdcommit(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFwdcommit(x uint64) (n int) {
	return sovFwdcommit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Fwdcommit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFwdcommit
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
			return fmt.Errorf("proto: Fwdcommit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fwdcommit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFwdcommit
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
				return ErrInvalidLengthFwdcommit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFwdcommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channelid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFwdcommit
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
				return ErrInvalidLengthFwdcommit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFwdcommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channelid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hashcodedest", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFwdcommit
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
				return ErrInvalidLengthFwdcommit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFwdcommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hashcodedest = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timelockreceiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFwdcommit
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
				return ErrInvalidLengthFwdcommit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFwdcommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Timelockreceiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timelocksender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFwdcommit
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
				return ErrInvalidLengthFwdcommit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFwdcommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Timelocksender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hashcodehtlc", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFwdcommit
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
				return ErrInvalidLengthFwdcommit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFwdcommit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hashcodehtlc = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFwdcommit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFwdcommit
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
func skipFwdcommit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFwdcommit
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
					return 0, ErrIntOverflowFwdcommit
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
					return 0, ErrIntOverflowFwdcommit
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
				return 0, ErrInvalidLengthFwdcommit
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFwdcommit
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFwdcommit
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFwdcommit        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFwdcommit          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFwdcommit = fmt.Errorf("proto: unexpected end of group")
)