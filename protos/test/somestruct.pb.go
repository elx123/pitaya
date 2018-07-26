// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: somestruct.proto

package test

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SomeStruct struct {
	A                    int32    `protobuf:"varint,1,opt,name=A,proto3" json:"A,omitempty"`
	B                    string   `protobuf:"bytes,2,opt,name=B,proto3" json:"B,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SomeStruct) Reset()         { *m = SomeStruct{} }
func (m *SomeStruct) String() string { return proto.CompactTextString(m) }
func (*SomeStruct) ProtoMessage()    {}
func (*SomeStruct) Descriptor() ([]byte, []int) {
	return fileDescriptor_somestruct_a546fe052fde1e0d, []int{0}
}
func (m *SomeStruct) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SomeStruct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SomeStruct.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SomeStruct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SomeStruct.Merge(dst, src)
}
func (m *SomeStruct) XXX_Size() int {
	return m.Size()
}
func (m *SomeStruct) XXX_DiscardUnknown() {
	xxx_messageInfo_SomeStruct.DiscardUnknown(m)
}

var xxx_messageInfo_SomeStruct proto.InternalMessageInfo

func (m *SomeStruct) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *SomeStruct) GetB() string {
	if m != nil {
		return m.B
	}
	return ""
}

func init() {
	proto.RegisterType((*SomeStruct)(nil), "test.SomeStruct")
}
func (m *SomeStruct) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SomeStruct) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.A != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSomestruct(dAtA, i, uint64(m.A))
	}
	if len(m.B) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSomestruct(dAtA, i, uint64(len(m.B)))
		i += copy(dAtA[i:], m.B)
	}
	return i, nil
}

func encodeVarintSomestruct(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SomeStruct) Size() (n int) {
	var l int
	_ = l
	if m.A != 0 {
		n += 1 + sovSomestruct(uint64(m.A))
	}
	l = len(m.B)
	if l > 0 {
		n += 1 + l + sovSomestruct(uint64(l))
	}
	return n
}

func sovSomestruct(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSomestruct(x uint64) (n int) {
	return sovSomestruct(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SomeStruct) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSomestruct
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SomeStruct: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SomeStruct: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			m.A = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSomestruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.A |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field B", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSomestruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSomestruct
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.B = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSomestruct(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSomestruct
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
func skipSomestruct(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSomestruct
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
					return 0, ErrIntOverflowSomestruct
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSomestruct
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSomestruct
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSomestruct
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSomestruct(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSomestruct = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSomestruct   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("somestruct.proto", fileDescriptor_somestruct_a546fe052fde1e0d) }

var fileDescriptor_somestruct_a546fe052fde1e0d = []byte{
	// 109 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0xce, 0xcf, 0x4d,
	0x2d, 0x2e, 0x29, 0x2a, 0x4d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x49,
	0x2d, 0x2e, 0x51, 0xd2, 0xe0, 0xe2, 0x0a, 0xce, 0xcf, 0x4d, 0x0d, 0x06, 0xcb, 0x08, 0xf1, 0x70,
	0x31, 0x3a, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x31, 0x3a, 0x82, 0x78, 0x4e, 0x12, 0x4c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x8c, 0x4e, 0x4e, 0x02, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24,
	0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x49, 0x6c, 0x60, 0x83, 0x8c, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x03, 0xc0, 0xd4, 0xcb, 0x5c, 0x00, 0x00, 0x00,
}
