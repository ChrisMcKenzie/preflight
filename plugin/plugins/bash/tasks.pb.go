// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: plugins/bash/tasks.proto

package bash

import (
	bytes "bytes"
	fmt "fmt"
	plugin "github.com/ChrisMcKenzie/preflight/plugin"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

type Script struct {
	Meta                 *plugin.Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	EvaluateScript       string       `protobuf:"bytes,2,opt,name=evaluateScript,proto3" json:"evaluateScript,omitempty"`
	ApplyScript          string       `protobuf:"bytes,3,opt,name=applyScript,proto3" json:"applyScript,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Script) Reset()      { *m = Script{} }
func (*Script) ProtoMessage() {}
func (*Script) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef8f88c59fcb432f, []int{0}
}
func (m *Script) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Script) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Script.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Script) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Script.Merge(m, src)
}
func (m *Script) XXX_Size() int {
	return m.Size()
}
func (m *Script) XXX_DiscardUnknown() {
	xxx_messageInfo_Script.DiscardUnknown(m)
}

var xxx_messageInfo_Script proto.InternalMessageInfo

func (m *Script) GetMeta() *plugin.Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *Script) GetEvaluateScript() string {
	if m != nil {
		return m.EvaluateScript
	}
	return ""
}

func (m *Script) GetApplyScript() string {
	if m != nil {
		return m.ApplyScript
	}
	return ""
}

func (*Script) XXX_MessageName() string {
	return "bash.script"
}
func init() {
	proto.RegisterType((*Script)(nil), "bash.script")
}

func init() { proto.RegisterFile("plugins/bash/tasks.proto", fileDescriptor_ef8f88c59fcb432f) }

var fileDescriptor_ef8f88c59fcb432f = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0xc8, 0x29, 0x4d,
	0xcf, 0xcc, 0x2b, 0xd6, 0x4f, 0x4a, 0x2c, 0xce, 0xd0, 0x2f, 0x49, 0x2c, 0xce, 0x2e, 0xd6, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0x89, 0x48, 0x09, 0x43, 0xe4, 0xf5, 0x21, 0x14, 0x44,
	0x4a, 0xa9, 0x84, 0x8b, 0xad, 0x38, 0xb9, 0x28, 0xb3, 0xa0, 0x44, 0x48, 0x81, 0x8b, 0x25, 0x37,
	0xb5, 0x24, 0x51, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0x88, 0x47, 0x0f, 0xaa, 0xcc, 0x37, 0xb5,
	0x24, 0x31, 0x08, 0x2c, 0x23, 0xa4, 0xc6, 0xc5, 0x97, 0x5a, 0x96, 0x98, 0x53, 0x9a, 0x58, 0x92,
	0x1a, 0x0c, 0xd6, 0x23, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x84, 0x26, 0x2a, 0xa4, 0xc0, 0xc5,
	0x9d, 0x58, 0x50, 0x90, 0x53, 0x09, 0x55, 0xc4, 0x0c, 0x56, 0x84, 0x2c, 0xe4, 0x94, 0x78, 0xe3,
	0xa1, 0x1c, 0xc3, 0x87, 0x87, 0x72, 0x8c, 0x3f, 0x1e, 0xca, 0x31, 0x36, 0x3c, 0x92, 0x63, 0x5c,
	0xf1, 0x48, 0x8e, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63,
	0x7c, 0xf1, 0x48, 0x8e, 0xe1, 0x03, 0x48, 0xfc, 0xb1, 0x1c, 0xe3, 0x89, 0xc7, 0x72, 0x8c, 0x51,
	0xfa, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xce, 0x19, 0x45, 0x99,
	0xc5, 0xbe, 0xc9, 0xde, 0xa9, 0x79, 0x55, 0x99, 0xa9, 0xfa, 0x05, 0x45, 0xa9, 0x69, 0x39, 0x99,
	0xe9, 0x19, 0x25, 0xfa, 0xc8, 0xfe, 0x4f, 0x62, 0x03, 0xfb, 0xcf, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0xff, 0x1c, 0x58, 0x61, 0x16, 0x01, 0x00, 0x00,
}

func (this *Script) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*Script)
	if !ok {
		that2, ok := that.(Script)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if c := this.Meta.Compare(that1.Meta); c != 0 {
		return c
	}
	if this.EvaluateScript != that1.EvaluateScript {
		if this.EvaluateScript < that1.EvaluateScript {
			return -1
		}
		return 1
	}
	if this.ApplyScript != that1.ApplyScript {
		if this.ApplyScript < that1.ApplyScript {
			return -1
		}
		return 1
	}
	if c := bytes.Compare(this.XXX_unrecognized, that1.XXX_unrecognized); c != 0 {
		return c
	}
	return 0
}
func (this *Script) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Script)
	if !ok {
		that2, ok := that.(Script)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Meta.Equal(that1.Meta) {
		return false
	}
	if this.EvaluateScript != that1.EvaluateScript {
		return false
	}
	if this.ApplyScript != that1.ApplyScript {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Script) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&bash.Script{")
	if this.Meta != nil {
		s = append(s, "Meta: "+fmt.Sprintf("%#v", this.Meta)+",\n")
	}
	s = append(s, "EvaluateScript: "+fmt.Sprintf("%#v", this.EvaluateScript)+",\n")
	s = append(s, "ApplyScript: "+fmt.Sprintf("%#v", this.ApplyScript)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTasks(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Script) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Script) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Script) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ApplyScript) > 0 {
		i -= len(m.ApplyScript)
		copy(dAtA[i:], m.ApplyScript)
		i = encodeVarintTasks(dAtA, i, uint64(len(m.ApplyScript)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.EvaluateScript) > 0 {
		i -= len(m.EvaluateScript)
		copy(dAtA[i:], m.EvaluateScript)
		i = encodeVarintTasks(dAtA, i, uint64(len(m.EvaluateScript)))
		i--
		dAtA[i] = 0x12
	}
	if m.Meta != nil {
		{
			size, err := m.Meta.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTasks(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTasks(dAtA []byte, offset int, v uint64) int {
	offset -= sovTasks(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedScript(r randyTasks, easy bool) *Script {
	this := &Script{}
	if r.Intn(5) != 0 {
		this.Meta = plugin.NewPopulatedMeta(r, easy)
	}
	this.EvaluateScript = string(randStringTasks(r))
	this.ApplyScript = string(randStringTasks(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTasks(r, 4)
	}
	return this
}

type randyTasks interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneTasks(r randyTasks) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringTasks(r randyTasks) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneTasks(r)
	}
	return string(tmps)
}
func randUnrecognizedTasks(r randyTasks, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldTasks(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldTasks(dAtA []byte, r randyTasks, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateTasks(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateTasks(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateTasks(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateTasks(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateTasks(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateTasks(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateTasks(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Script) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Meta != nil {
		l = m.Meta.Size()
		n += 1 + l + sovTasks(uint64(l))
	}
	l = len(m.EvaluateScript)
	if l > 0 {
		n += 1 + l + sovTasks(uint64(l))
	}
	l = len(m.ApplyScript)
	if l > 0 {
		n += 1 + l + sovTasks(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTasks(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTasks(x uint64) (n int) {
	return sovTasks(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Script) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Script{`,
		`Meta:` + strings.Replace(fmt.Sprintf("%v", this.Meta), "Meta", "plugin.Meta", 1) + `,`,
		`EvaluateScript:` + fmt.Sprintf("%v", this.EvaluateScript) + `,`,
		`ApplyScript:` + fmt.Sprintf("%v", this.ApplyScript) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTasks(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Script) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTasks
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
			return fmt.Errorf("proto: script: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: script: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
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
				return ErrInvalidLengthTasks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTasks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Meta == nil {
				m.Meta = &plugin.Meta{}
			}
			if err := m.Meta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EvaluateScript", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
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
				return ErrInvalidLengthTasks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTasks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EvaluateScript = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplyScript", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
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
				return ErrInvalidLengthTasks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTasks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApplyScript = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTasks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTasks
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTasks
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTasks(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTasks
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
					return 0, ErrIntOverflowTasks
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
					return 0, ErrIntOverflowTasks
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
				return 0, ErrInvalidLengthTasks
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTasks
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTasks
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTasks        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTasks          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTasks = fmt.Errorf("proto: unexpected end of group")
)