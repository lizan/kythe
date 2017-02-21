// Code generated by protoc-gen-gogo.
// source: kythe/proto/common.proto
// DO NOT EDIT!

/*
	Package common_proto is a generated protocol buffer package.

	It is generated from these files:
		kythe/proto/common.proto

	It has these top-level messages:
		Fact
		Point
		Span
		NodeInfo
*/
package common_proto

import proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Fact represents a single key/value fact from the graph.
type Fact struct {
	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Fact) Reset()                    { *m = Fact{} }
func (m *Fact) String() string            { return proto.CompactTextString(m) }
func (*Fact) ProtoMessage()               {}
func (*Fact) Descriptor() ([]byte, []int) { return fileDescriptorCommon, []int{0} }

func (m *Fact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Fact) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// A Point represents a location within a file or buffer.
//
// If line_number ≤ 0, the line number and column offset are considered
// unknown and will be ignored.
//
// A point with line_number > 0 is said to be _normalized_ if it satisfies
// the constraint 0 ≤ column_offset ≤ bytelen(line_number); that is, if the
// column_offset is within the actual range of the corresponding line.  A
// point can be normalized by adjusting line_number and column_offset so that
// this constraint is satisfied.  This may be impossible if the column offset
// exceeds the bounds of the file.
type Point struct {
	ByteOffset   int32 `protobuf:"varint,1,opt,name=byte_offset,json=byteOffset,proto3" json:"byte_offset,omitempty"`
	LineNumber   int32 `protobuf:"varint,2,opt,name=line_number,json=lineNumber,proto3" json:"line_number,omitempty"`
	ColumnOffset int32 `protobuf:"varint,3,opt,name=column_offset,json=columnOffset,proto3" json:"column_offset,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptorCommon, []int{1} }

func (m *Point) GetByteOffset() int32 {
	if m != nil {
		return m.ByteOffset
	}
	return 0
}

func (m *Point) GetLineNumber() int32 {
	if m != nil {
		return m.LineNumber
	}
	return 0
}

func (m *Point) GetColumnOffset() int32 {
	if m != nil {
		return m.ColumnOffset
	}
	return 0
}

// A Span represents an inclusive-exclusive range inside of a file or buffer.
type Span struct {
	Start *Point `protobuf:"bytes,1,opt,name=start" json:"start,omitempty"`
	End   *Point `protobuf:"bytes,2,opt,name=end" json:"end,omitempty"`
}

func (m *Span) Reset()                    { *m = Span{} }
func (m *Span) String() string            { return proto.CompactTextString(m) }
func (*Span) ProtoMessage()               {}
func (*Span) Descriptor() ([]byte, []int) { return fileDescriptorCommon, []int{2} }

func (m *Span) GetStart() *Point {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *Span) GetEnd() *Point {
	if m != nil {
		return m.End
	}
	return nil
}

type NodeInfo struct {
	// The matching facts known for that node, a map from fact name to value.
	Facts map[string][]byte `protobuf:"bytes,2,rep,name=facts" json:"facts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// If known and unambiguous, an anchor ticket for this node's definition
	// location.
	// Tickets are Kythe URIs (http://www.kythe.io/docs/kythe-uri-spec.html).
	Definition string `protobuf:"bytes,5,opt,name=definition,proto3" json:"definition,omitempty"`
}

func (m *NodeInfo) Reset()                    { *m = NodeInfo{} }
func (m *NodeInfo) String() string            { return proto.CompactTextString(m) }
func (*NodeInfo) ProtoMessage()               {}
func (*NodeInfo) Descriptor() ([]byte, []int) { return fileDescriptorCommon, []int{3} }

func (m *NodeInfo) GetFacts() map[string][]byte {
	if m != nil {
		return m.Facts
	}
	return nil
}

func (m *NodeInfo) GetDefinition() string {
	if m != nil {
		return m.Definition
	}
	return ""
}

func init() {
	proto.RegisterType((*Fact)(nil), "kythe.proto.common.Fact")
	proto.RegisterType((*Point)(nil), "kythe.proto.common.Point")
	proto.RegisterType((*Span)(nil), "kythe.proto.common.Span")
	proto.RegisterType((*NodeInfo)(nil), "kythe.proto.common.NodeInfo")
}
func (m *Fact) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Fact) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	return i, nil
}

func (m *Point) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Point) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ByteOffset != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintCommon(dAtA, i, uint64(m.ByteOffset))
	}
	if m.LineNumber != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintCommon(dAtA, i, uint64(m.LineNumber))
	}
	if m.ColumnOffset != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintCommon(dAtA, i, uint64(m.ColumnOffset))
	}
	return i, nil
}

func (m *Span) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Span) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Start != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCommon(dAtA, i, uint64(m.Start.Size()))
		n1, err := m.Start.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.End != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCommon(dAtA, i, uint64(m.End.Size()))
		n2, err := m.End.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *NodeInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodeInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Facts) > 0 {
		for k, _ := range m.Facts {
			dAtA[i] = 0x12
			i++
			v := m.Facts[k]
			byteSize := 0
			if len(v) > 0 {
				byteSize = 1 + len(v) + sovCommon(uint64(len(v)))
			}
			mapSize := 1 + len(k) + sovCommon(uint64(len(k))) + byteSize
			i = encodeVarintCommon(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintCommon(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			if len(v) > 0 {
				dAtA[i] = 0x12
				i++
				i = encodeVarintCommon(dAtA, i, uint64(len(v)))
				i += copy(dAtA[i:], v)
			}
		}
	}
	if len(m.Definition) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Definition)))
		i += copy(dAtA[i:], m.Definition)
	}
	return i, nil
}

func encodeFixed64Common(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Common(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintCommon(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Fact) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func (m *Point) Size() (n int) {
	var l int
	_ = l
	if m.ByteOffset != 0 {
		n += 1 + sovCommon(uint64(m.ByteOffset))
	}
	if m.LineNumber != 0 {
		n += 1 + sovCommon(uint64(m.LineNumber))
	}
	if m.ColumnOffset != 0 {
		n += 1 + sovCommon(uint64(m.ColumnOffset))
	}
	return n
}

func (m *Span) Size() (n int) {
	var l int
	_ = l
	if m.Start != nil {
		l = m.Start.Size()
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.End != nil {
		l = m.End.Size()
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func (m *NodeInfo) Size() (n int) {
	var l int
	_ = l
	if len(m.Facts) > 0 {
		for k, v := range m.Facts {
			_ = k
			_ = v
			l = 0
			if len(v) > 0 {
				l = 1 + len(v) + sovCommon(uint64(len(v)))
			}
			mapEntrySize := 1 + len(k) + sovCommon(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovCommon(uint64(mapEntrySize))
		}
	}
	l = len(m.Definition)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func sovCommon(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCommon(x uint64) (n int) {
	return sovCommon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Fact) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: Fact: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fact: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
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
func (m *Point) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: Point: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Point: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ByteOffset", wireType)
			}
			m.ByteOffset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ByteOffset |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LineNumber", wireType)
			}
			m.LineNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LineNumber |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ColumnOffset", wireType)
			}
			m.ColumnOffset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ColumnOffset |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
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
func (m *Span) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: Span: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Span: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Start == nil {
				m.Start = &Point{}
			}
			if err := m.Start.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.End == nil {
				m.End = &Point{}
			}
			if err := m.End.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
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
func (m *NodeInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: NodeInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Facts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthCommon
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.Facts == nil {
				m.Facts = make(map[string][]byte)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCommon
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var mapbyteLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCommon
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					mapbyteLen |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intMapbyteLen := int(mapbyteLen)
				if intMapbyteLen < 0 {
					return ErrInvalidLengthCommon
				}
				postbytesIndex := iNdEx + intMapbyteLen
				if postbytesIndex > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := make([]byte, mapbyteLen)
				copy(mapvalue, dAtA[iNdEx:postbytesIndex])
				iNdEx = postbytesIndex
				m.Facts[mapkey] = mapvalue
			} else {
				var mapvalue []byte
				m.Facts[mapkey] = mapvalue
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Definition", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Definition = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
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
func skipCommon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
				return 0, ErrInvalidLengthCommon
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCommon
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
				next, err := skipCommon(dAtA[start:])
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
	ErrInvalidLengthCommon = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommon   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("kythe/proto/common.proto", fileDescriptorCommon) }

var fileDescriptorCommon = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x4a, 0xfb, 0x40,
	0x10, 0xc7, 0x7f, 0x9b, 0x3f, 0xa5, 0xbf, 0x69, 0x85, 0xb2, 0x78, 0x88, 0x1e, 0xd2, 0x12, 0x0f,
	0x16, 0x84, 0x54, 0xda, 0x4b, 0x11, 0xbc, 0x08, 0x0a, 0x7a, 0xa8, 0x25, 0x3e, 0x40, 0x49, 0x93,
	0x49, 0x0d, 0x4d, 0x76, 0x4a, 0xb2, 0x2d, 0xe4, 0x4d, 0x7c, 0x0d, 0xdf, 0xc2, 0xa3, 0x8f, 0x20,
	0xf5, 0x45, 0x24, 0xbb, 0x16, 0x0b, 0x16, 0x6f, 0x33, 0xdf, 0xfd, 0x7c, 0x67, 0x66, 0xbf, 0xe0,
	0x2c, 0x2b, 0xf9, 0x8c, 0x83, 0x55, 0x41, 0x92, 0x06, 0x11, 0xe5, 0x39, 0x09, 0x5f, 0x35, 0x9c,
	0xab, 0x17, 0xdd, 0xf8, 0xfa, 0xc5, 0xbb, 0x04, 0xeb, 0x2e, 0x8c, 0x24, 0xe7, 0x60, 0x89, 0x30,
	0x47, 0x87, 0xf5, 0x58, 0xff, 0x7f, 0xa0, 0x6a, 0x7e, 0x0c, 0xf6, 0x26, 0xcc, 0xd6, 0xe8, 0x18,
	0x3d, 0xd6, 0x6f, 0x07, 0xba, 0xf1, 0x04, 0xd8, 0x53, 0x4a, 0x85, 0xe4, 0x5d, 0x68, 0xcd, 0x2b,
	0x89, 0x33, 0x4a, 0x92, 0x12, 0xa5, 0x72, 0xda, 0x01, 0xd4, 0xd2, 0xa3, 0x52, 0x6a, 0x20, 0x4b,
	0x05, 0xce, 0xc4, 0x3a, 0x9f, 0x63, 0xa1, 0xa6, 0xd8, 0x01, 0xd4, 0xd2, 0x44, 0x29, 0xfc, 0x0c,
	0x8e, 0x22, 0xca, 0xd6, 0xb9, 0xd8, 0xcd, 0x30, 0x15, 0xd2, 0xd6, 0xa2, 0x9e, 0xe2, 0xc5, 0x60,
	0x3d, 0xad, 0x42, 0xc1, 0x07, 0x60, 0x97, 0x32, 0x2c, 0xf4, 0xa2, 0xd6, 0xf0, 0xc4, 0xff, 0xfd,
	0x1b, 0x5f, 0x1d, 0x16, 0x68, 0x8e, 0x5f, 0x80, 0x89, 0x22, 0x56, 0x6b, 0xff, 0xc4, 0x6b, 0xca,
	0x7b, 0x65, 0xd0, 0x9c, 0x50, 0x8c, 0xf7, 0x22, 0x21, 0x7e, 0x0d, 0x76, 0x12, 0x46, 0xb2, 0x74,
	0x8c, 0x9e, 0xd9, 0x6f, 0x0d, 0xcf, 0x0f, 0x79, 0x77, 0xb0, 0x5f, 0xc7, 0x57, 0xde, 0x0a, 0x59,
	0x54, 0x81, 0x76, 0x71, 0x17, 0x20, 0xc6, 0x24, 0x15, 0xa9, 0x4c, 0x49, 0x38, 0xb6, 0x4a, 0x74,
	0x4f, 0x39, 0x1d, 0x03, 0xfc, 0x98, 0x78, 0x07, 0xcc, 0x25, 0x56, 0xdf, 0xc1, 0xd7, 0xe5, 0xe1,
	0xdc, 0xaf, 0x8c, 0x31, 0x7b, 0xb0, 0x9a, 0xac, 0x63, 0x04, 0x0d, 0x99, 0x46, 0x4b, 0x94, 0x37,
	0xa3, 0xb7, 0xad, 0xcb, 0xde, 0xb7, 0x2e, 0xfb, 0xd8, 0xba, 0xec, 0xe5, 0xd3, 0xfd, 0x07, 0xdd,
	0x88, 0x72, 0x7f, 0x41, 0xb4, 0xc8, 0xd0, 0x8f, 0x71, 0x23, 0x89, 0xb2, 0x72, 0xff, 0xf8, 0x29,
	0x9b, 0x37, 0x54, 0x31, 0xfa, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x73, 0xe4, 0x0f, 0x27, 0x02,
	0x00, 0x00,
}
