// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: plugin.proto

package plugin

import (
	bytes "bytes"
	fmt "fmt"
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

type Meta struct {
	URL                  string   `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Enforce              bool     `protobuf:"varint,4,opt,name=enforce,proto3" json:"enforce,omitempty"`
	Dependencies         []string `protobuf:"bytes,5,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Meta) Reset()      { *m = Meta{} }
func (*Meta) ProtoMessage() {}
func (*Meta) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{0}
}
func (m *Meta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Meta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Meta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Meta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Meta.Merge(m, src)
}
func (m *Meta) XXX_Size() int {
	return m.Size()
}
func (m *Meta) XXX_DiscardUnknown() {
	xxx_messageInfo_Meta.DiscardUnknown(m)
}

var xxx_messageInfo_Meta proto.InternalMessageInfo

func (m *Meta) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

func (m *Meta) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Meta) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Meta) GetEnforce() bool {
	if m != nil {
		return m.Enforce
	}
	return false
}

func (m *Meta) GetDependencies() []string {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func (*Meta) XXX_MessageName() string {
	return "plugin.Meta"
}
func init() {
	proto.RegisterType((*Meta)(nil), "plugin.Meta")
}

func init() { proto.RegisterFile("plugin.proto", fileDescriptor_22a625af4bc1cc87) }

var fileDescriptor_22a625af4bc1cc87 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x21, 0x4f, 0x03, 0x31,
	0x18, 0x86, 0xf7, 0x71, 0xc7, 0x60, 0xcd, 0x04, 0xa9, 0xaa, 0xfa, 0x72, 0x99, 0x3a, 0xcc, 0x4e,
	0xf0, 0x0f, 0x40, 0xc2, 0x4c, 0x13, 0x0c, 0xa8, 0xad, 0xfb, 0x76, 0xd7, 0x64, 0x6b, 0x9b, 0x5e,
	0x87, 0x40, 0x21, 0xf8, 0x31, 0xfc, 0x94, 0x49, 0x24, 0x92, 0xeb, 0x0c, 0x72, 0x12, 0x49, 0xae,
	0xcb, 0x04, 0xee, 0x7d, 0x9e, 0x57, 0x3d, 0x6c, 0xec, 0xd6, 0xdb, 0x5a, 0x9b, 0xa9, 0xf3, 0x36,
	0x58, 0x3e, 0x3c, 0xd2, 0xe4, 0x1d, 0x58, 0x3e, 0xa3, 0x30, 0xe7, 0x57, 0x2c, 0x7b, 0x94, 0x0f,
	0x02, 0x0a, 0x28, 0x47, 0xb2, 0x9f, 0x9c, 0xb3, 0xdc, 0xcc, 0x37, 0x24, 0xce, 0x92, 0x4a, 0x9b,
	0x0b, 0x76, 0xf1, 0x42, 0xbe, 0xd5, 0xd6, 0x88, 0x2c, 0xe9, 0x13, 0xf6, 0x0f, 0x99, 0x95, 0xf5,
	0x8a, 0x44, 0x5e, 0x40, 0x79, 0x29, 0x4f, 0xc8, 0x27, 0x6c, 0xbc, 0x24, 0x47, 0x66, 0x49, 0x46,
	0x69, 0x6a, 0xc5, 0x79, 0x91, 0x95, 0x23, 0xf9, 0xcf, 0xdd, 0x3e, 0x7f, 0x75, 0x38, 0x38, 0x74,
	0x08, 0xbf, 0x1d, 0xc2, 0x5b, 0x44, 0xf8, 0x88, 0x08, 0xbb, 0x88, 0xf0, 0x19, 0x11, 0xbe, 0x23,
	0xc2, 0x4f, 0xc4, 0xc1, 0xa1, 0xf7, 0x7b, 0x84, 0xdd, 0x1e, 0xe1, 0xe9, 0xba, 0xd6, 0xa1, 0xd9,
	0x2e, 0xa6, 0xca, 0x6e, 0xaa, 0xbb, 0xc6, 0xeb, 0x76, 0xa6, 0xee, 0xc9, 0xbc, 0x6a, 0xaa, 0x9c,
	0xa7, 0xd5, 0x5a, 0xd7, 0x4d, 0xa8, 0x8e, 0x8d, 0x8b, 0x61, 0x4a, 0xbe, 0xf9, 0x0b, 0x00, 0x00,
	0xff, 0xff, 0x07, 0xb2, 0xdc, 0xa9, 0x02, 0x01, 0x00, 0x00,
}

func (this *Meta) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*Meta)
	if !ok {
		that2, ok := that.(Meta)
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
	if this.URL != that1.URL {
		if this.URL < that1.URL {
			return -1
		}
		return 1
	}
	if this.Name != that1.Name {
		if this.Name < that1.Name {
			return -1
		}
		return 1
	}
	if this.Version != that1.Version {
		if this.Version < that1.Version {
			return -1
		}
		return 1
	}
	if this.Enforce != that1.Enforce {
		if !this.Enforce {
			return -1
		}
		return 1
	}
	if len(this.Dependencies) != len(that1.Dependencies) {
		if len(this.Dependencies) < len(that1.Dependencies) {
			return -1
		}
		return 1
	}
	for i := range this.Dependencies {
		if this.Dependencies[i] != that1.Dependencies[i] {
			if this.Dependencies[i] < that1.Dependencies[i] {
				return -1
			}
			return 1
		}
	}
	if c := bytes.Compare(this.XXX_unrecognized, that1.XXX_unrecognized); c != 0 {
		return c
	}
	return 0
}
func (this *Meta) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Meta)
	if !ok {
		that2, ok := that.(Meta)
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
	if this.URL != that1.URL {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	if this.Enforce != that1.Enforce {
		return false
	}
	if len(this.Dependencies) != len(that1.Dependencies) {
		return false
	}
	for i := range this.Dependencies {
		if this.Dependencies[i] != that1.Dependencies[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Meta) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&plugin.Meta{")
	s = append(s, "URL: "+fmt.Sprintf("%#v", this.URL)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Version: "+fmt.Sprintf("%#v", this.Version)+",\n")
	s = append(s, "Enforce: "+fmt.Sprintf("%#v", this.Enforce)+",\n")
	s = append(s, "Dependencies: "+fmt.Sprintf("%#v", this.Dependencies)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPlugin(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Meta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Meta) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Meta) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Dependencies) > 0 {
		for iNdEx := len(m.Dependencies) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Dependencies[iNdEx])
			copy(dAtA[i:], m.Dependencies[iNdEx])
			i = encodeVarintPlugin(dAtA, i, uint64(len(m.Dependencies[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Enforce {
		i--
		if m.Enforce {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintPlugin(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPlugin(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.URL) > 0 {
		i -= len(m.URL)
		copy(dAtA[i:], m.URL)
		i = encodeVarintPlugin(dAtA, i, uint64(len(m.URL)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPlugin(dAtA []byte, offset int, v uint64) int {
	offset -= sovPlugin(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedMeta(r randyPlugin, easy bool) *Meta {
	this := &Meta{}
	this.URL = string(randStringPlugin(r))
	this.Name = string(randStringPlugin(r))
	this.Version = string(randStringPlugin(r))
	this.Enforce = bool(bool(r.Intn(2) == 0))
	v1 := r.Intn(10)
	this.Dependencies = make([]string, v1)
	for i := 0; i < v1; i++ {
		this.Dependencies[i] = string(randStringPlugin(r))
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedPlugin(r, 6)
	}
	return this
}

type randyPlugin interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RunePlugin(r randyPlugin) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringPlugin(r randyPlugin) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RunePlugin(r)
	}
	return string(tmps)
}
func randUnrecognizedPlugin(r randyPlugin, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldPlugin(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldPlugin(dAtA []byte, r randyPlugin, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulatePlugin(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulatePlugin(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulatePlugin(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulatePlugin(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulatePlugin(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulatePlugin(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulatePlugin(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Meta) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.URL)
	if l > 0 {
		n += 1 + l + sovPlugin(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPlugin(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovPlugin(uint64(l))
	}
	if m.Enforce {
		n += 2
	}
	if len(m.Dependencies) > 0 {
		for _, s := range m.Dependencies {
			l = len(s)
			n += 1 + l + sovPlugin(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPlugin(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPlugin(x uint64) (n int) {
	return sovPlugin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Meta) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Meta{`,
		`URL:` + fmt.Sprintf("%v", this.URL) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Version:` + fmt.Sprintf("%v", this.Version) + `,`,
		`Enforce:` + fmt.Sprintf("%v", this.Enforce) + `,`,
		`Dependencies:` + fmt.Sprintf("%v", this.Dependencies) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringPlugin(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Meta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlugin
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
			return fmt.Errorf("proto: Meta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Meta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field URL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
				return ErrInvalidLengthPlugin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlugin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.URL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
				return ErrInvalidLengthPlugin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlugin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
				return ErrInvalidLengthPlugin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlugin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enforce", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enforce = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dependencies", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
				return ErrInvalidLengthPlugin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlugin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Dependencies = append(m.Dependencies, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPlugin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlugin
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPlugin
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
func skipPlugin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPlugin
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
					return 0, ErrIntOverflowPlugin
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
					return 0, ErrIntOverflowPlugin
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
				return 0, ErrInvalidLengthPlugin
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPlugin
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPlugin
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPlugin        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlugin          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPlugin = fmt.Errorf("proto: unexpected end of group")
)