// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/containerd/containerd/api/types/descriptor.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
	github_com_opencontainers_go_digest "github.com/opencontainers/go-digest"
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

// Descriptor describes a blob in a content store.
//
// This descriptor can be used to reference content from an
// oci descriptor found in a manifest.
// See https://godoc.org/github.com/opencontainers/image-spec/specs-go/v1#Descriptor
type Descriptor struct {
	MediaType            string                                     `protobuf:"bytes,1,opt,name=media_type,json=mediaType,proto3" json:"media_type,omitempty"`
	Digest               github_com_opencontainers_go_digest.Digest `protobuf:"bytes,2,opt,name=digest,proto3,customtype=github.com/opencontainers/go-digest.Digest" json:"digest"`
	Size_                int64                                      `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Annotations          map[string]string                          `protobuf:"bytes,5,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *Descriptor) Reset()      { *m = Descriptor{} }
func (*Descriptor) ProtoMessage() {}
func (*Descriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_37f958df3707db9e, []int{0}
}
func (m *Descriptor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Descriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Descriptor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Descriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Descriptor.Merge(m, src)
}
func (m *Descriptor) XXX_Size() int {
	return m.Size()
}
func (m *Descriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_Descriptor.DiscardUnknown(m)
}

var xxx_messageInfo_Descriptor proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Descriptor)(nil), "containerd.types.Descriptor")
	proto.RegisterMapType((map[string]string)(nil), "containerd.types.Descriptor.AnnotationsEntry")
}

func init() {
	proto.RegisterFile("github.com/containerd/containerd/api/types/descriptor.proto", fileDescriptor_37f958df3707db9e)
}

var fileDescriptor_37f958df3707db9e = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4e, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0xcf, 0x2b, 0x49, 0xcc, 0xcc, 0x4b, 0x2d,
	0x4a, 0x41, 0x66, 0x26, 0x16, 0x64, 0xea, 0x97, 0x54, 0x16, 0xa4, 0x16, 0xeb, 0xa7, 0xa4, 0x16,
	0x27, 0x17, 0x65, 0x16, 0x94, 0xe4, 0x17, 0xe9, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x09, 0x20,
	0x94, 0xe9, 0x81, 0x95, 0x48, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x25, 0xf5, 0x41, 0x2c, 0x88,
	0x3a, 0xa5, 0x39, 0x4c, 0x5c, 0x5c, 0x2e, 0x70, 0xcd, 0x42, 0xb2, 0x5c, 0x5c, 0xb9, 0xa9, 0x29,
	0x99, 0x89, 0xf1, 0x20, 0x3d, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x9c, 0x60, 0x91, 0x90,
	0xca, 0x82, 0x54, 0x21, 0x2f, 0x2e, 0xb6, 0x94, 0xcc, 0xf4, 0xd4, 0xe2, 0x12, 0x09, 0x26, 0x90,
	0x94, 0x93, 0xd1, 0x89, 0x7b, 0xf2, 0x0c, 0xb7, 0xee, 0xc9, 0x6b, 0x21, 0x39, 0x35, 0xbf, 0x20,
	0x35, 0x0f, 0x6e, 0x79, 0xb1, 0x7e, 0x7a, 0xbe, 0x2e, 0x44, 0x8b, 0x9e, 0x0b, 0x98, 0x0a, 0x82,
	0x9a, 0x20, 0x24, 0xc4, 0xc5, 0x52, 0x9c, 0x59, 0x95, 0x2a, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x1c,
	0x04, 0x66, 0x0b, 0xf9, 0x73, 0x71, 0x27, 0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x96, 0x64, 0xe6, 0xe7,
	0x15, 0x4b, 0xb0, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0xe9, 0xea, 0xa1, 0xfb, 0x45, 0x0f, 0xe1, 0x62,
	0x3d, 0x47, 0x84, 0x7a, 0xd7, 0xbc, 0x92, 0xa2, 0xca, 0x20, 0x64, 0x13, 0xa4, 0xec, 0xb8, 0x04,
	0xd0, 0x15, 0x08, 0x09, 0x70, 0x31, 0x67, 0xa7, 0x56, 0x42, 0x3d, 0x07, 0x62, 0x0a, 0x89, 0x70,
	0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6, 0x42, 0x7c, 0x15, 0x04, 0xe1, 0x58, 0x31, 0x59, 0x30, 0x3a,
	0x79, 0x9d, 0x78, 0x28, 0xc7, 0x70, 0xe3, 0xa1, 0x1c, 0x43, 0xc3, 0x23, 0x39, 0xc6, 0x13, 0x8f,
	0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x31, 0xca, 0x80, 0xf8, 0xd8, 0xb1,
	0x06, 0x93, 0x49, 0x6c, 0xe0, 0x10, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x81, 0x35, 0x34,
	0x94, 0xd8, 0x01, 0x00, 0x00,
}

func (m *Descriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Descriptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Descriptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Annotations) > 0 {
		for k := range m.Annotations {
			v := m.Annotations[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintDescriptor(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintDescriptor(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintDescriptor(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Size_ != 0 {
		i = encodeVarintDescriptor(dAtA, i, uint64(m.Size_))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Digest) > 0 {
		i -= len(m.Digest)
		copy(dAtA[i:], m.Digest)
		i = encodeVarintDescriptor(dAtA, i, uint64(len(m.Digest)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.MediaType) > 0 {
		i -= len(m.MediaType)
		copy(dAtA[i:], m.MediaType)
		i = encodeVarintDescriptor(dAtA, i, uint64(len(m.MediaType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDescriptor(dAtA []byte, offset int, v uint64) int {
	offset -= sovDescriptor(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Descriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MediaType)
	if l > 0 {
		n += 1 + l + sovDescriptor(uint64(l))
	}
	l = len(m.Digest)
	if l > 0 {
		n += 1 + l + sovDescriptor(uint64(l))
	}
	if m.Size_ != 0 {
		n += 1 + sovDescriptor(uint64(m.Size_))
	}
	if len(m.Annotations) > 0 {
		for k, v := range m.Annotations {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovDescriptor(uint64(len(k))) + 1 + len(v) + sovDescriptor(uint64(len(v)))
			n += mapEntrySize + 1 + sovDescriptor(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDescriptor(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDescriptor(x uint64) (n int) {
	return sovDescriptor(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Descriptor) String() string {
	if this == nil {
		return "nil"
	}
	keysForAnnotations := make([]string, 0, len(this.Annotations))
	for k, _ := range this.Annotations {
		keysForAnnotations = append(keysForAnnotations, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForAnnotations)
	mapStringForAnnotations := "map[string]string{"
	for _, k := range keysForAnnotations {
		mapStringForAnnotations += fmt.Sprintf("%v: %v,", k, this.Annotations[k])
	}
	mapStringForAnnotations += "}"
	s := strings.Join([]string{`&Descriptor{`,
		`MediaType:` + fmt.Sprintf("%v", this.MediaType) + `,`,
		`Digest:` + fmt.Sprintf("%v", this.Digest) + `,`,
		`Size_:` + fmt.Sprintf("%v", this.Size_) + `,`,
		`Annotations:` + mapStringForAnnotations + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringDescriptor(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Descriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDescriptor
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
			return fmt.Errorf("proto: Descriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Descriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MediaType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDescriptor
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
				return ErrInvalidLengthDescriptor
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDescriptor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MediaType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Digest", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDescriptor
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
				return ErrInvalidLengthDescriptor
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDescriptor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Digest = github_com_opencontainers_go_digest.Digest(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDescriptor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Annotations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDescriptor
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
				return ErrInvalidLengthDescriptor
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDescriptor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Annotations == nil {
				m.Annotations = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowDescriptor
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowDescriptor
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthDescriptor
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthDescriptor
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowDescriptor
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthDescriptor
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthDescriptor
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipDescriptor(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthDescriptor
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Annotations[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDescriptor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDescriptor
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
func skipDescriptor(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDescriptor
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
					return 0, ErrIntOverflowDescriptor
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
					return 0, ErrIntOverflowDescriptor
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
				return 0, ErrInvalidLengthDescriptor
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDescriptor
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDescriptor
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDescriptor        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDescriptor          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDescriptor = fmt.Errorf("proto: unexpected end of group")
)
