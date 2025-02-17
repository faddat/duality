// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dualitylabs/duality/incentives/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// GenesisState defines the incentives module's various parameters when first
// initialized
type GenesisState struct {
	// params are all the parameters of the module
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// gauges are all gauges that should exist at genesis
	Gauges []*Gauge `protobuf:"bytes,2,rep,name=gauges,proto3" json:"gauges,omitempty"`
	// last_gauge_id is what the gauge number will increment from when creating
	// the next gauge after genesis
	LastGaugeId      uint64            `protobuf:"varint,3,opt,name=last_gauge_id,json=lastGaugeId,proto3" json:"last_gauge_id,omitempty"`
	LastStakeId      uint64            `protobuf:"varint,4,opt,name=last_stake_id,json=lastStakeId,proto3" json:"last_stake_id,omitempty"`
	Stakes           []*Stake          `protobuf:"bytes,5,rep,name=stakes,proto3" json:"stakes,omitempty"`
	AccountHistories []*AccountHistory `protobuf:"bytes,6,rep,name=accountHistories,proto3" json:"accountHistories,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_440efef4d375f023, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetGauges() []*Gauge {
	if m != nil {
		return m.Gauges
	}
	return nil
}

func (m *GenesisState) GetLastGaugeId() uint64 {
	if m != nil {
		return m.LastGaugeId
	}
	return 0
}

func (m *GenesisState) GetLastStakeId() uint64 {
	if m != nil {
		return m.LastStakeId
	}
	return 0
}

func (m *GenesisState) GetStakes() []*Stake {
	if m != nil {
		return m.Stakes
	}
	return nil
}

func (m *GenesisState) GetAccountHistories() []*AccountHistory {
	if m != nil {
		return m.AccountHistories
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "dualitylabs.duality.incentives.GenesisState")
}

func init() {
	proto.RegisterFile("dualitylabs/duality/incentives/genesis.proto", fileDescriptor_440efef4d375f023)
}

var fileDescriptor_440efef4d375f023 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x4b, 0xfb, 0x30,
	0x1c, 0xc6, 0x9b, 0x6d, 0xbf, 0x1e, 0xb2, 0x9f, 0x20, 0xc5, 0x43, 0xd9, 0x21, 0x8e, 0x81, 0x32,
	0xfc, 0x93, 0xc2, 0xf4, 0xea, 0xc1, 0x21, 0xcc, 0x81, 0x07, 0xe9, 0x6e, 0xbb, 0x8c, 0xac, 0x0d,
	0x5d, 0x70, 0x6b, 0xc6, 0x92, 0x8a, 0x7d, 0x17, 0xbe, 0xac, 0x1d, 0x77, 0xf4, 0x24, 0xd2, 0xbe,
	0x07, 0xcf, 0x92, 0x34, 0x73, 0x15, 0xc1, 0xee, 0xf6, 0x6d, 0xfa, 0x79, 0xbe, 0x79, 0xf2, 0x3c,
	0xf0, 0x22, 0x4c, 0xc8, 0x9c, 0xc9, 0x74, 0x4e, 0xa6, 0xc2, 0x33, 0xb3, 0xc7, 0xe2, 0x80, 0xc6,
	0x92, 0x3d, 0x53, 0xe1, 0x45, 0x34, 0xa6, 0x82, 0x09, 0xbc, 0x5c, 0x71, 0xc9, 0x1d, 0x54, 0xa2,
	0xb1, 0x99, 0xf1, 0x8e, 0x6e, 0x1d, 0x45, 0x3c, 0xe2, 0x1a, 0xf5, 0xd4, 0x54, 0xa8, 0x5a, 0xe7,
	0x15, 0x77, 0x2c, 0xc9, 0x8a, 0x2c, 0xcc, 0x15, 0xad, 0xb3, 0x2a, 0x43, 0x24, 0x89, 0xe8, 0x9e,
	0xac, 0x90, 0xe4, 0x69, 0xcb, 0x5e, 0x57, 0xb0, 0x24, 0x08, 0x78, 0x12, 0xcb, 0xc9, 0x8c, 0x09,
	0xc9, 0x57, 0x69, 0xa1, 0xea, 0x7c, 0xd6, 0xe0, 0xff, 0x41, 0x11, 0xc1, 0x48, 0x12, 0x49, 0x9d,
	0x3b, 0x68, 0x17, 0x76, 0x5d, 0xd0, 0x06, 0xdd, 0x66, 0xef, 0x14, 0xff, 0x1d, 0x09, 0x7e, 0xd4,
	0x74, 0xbf, 0xb1, 0x7e, 0x3f, 0xb6, 0x7c, 0xa3, 0x75, 0x6e, 0xa0, 0xad, 0xdf, 0x21, 0xdc, 0x5a,
	0xbb, 0xde, 0x6d, 0xf6, 0x4e, 0xaa, 0xb6, 0x0c, 0x14, 0xed, 0x1b, 0x91, 0xd3, 0x81, 0x07, 0x73,
	0x22, 0xe4, 0x44, 0x7f, 0x4e, 0x58, 0xe8, 0xd6, 0xdb, 0xa0, 0xdb, 0xf0, 0x9b, 0xea, 0x50, 0x93,
	0xc3, 0xf0, 0x9b, 0xd1, 0x19, 0x28, 0xa6, 0xb1, 0x63, 0x46, 0xea, 0x6c, 0x18, 0x2a, 0x1b, 0xfa,
	0xb7, 0x70, 0xff, 0xed, 0x67, 0x43, 0x0b, 0x7d, 0x23, 0x72, 0xc6, 0xf0, 0xd0, 0xa4, 0x76, 0xaf,
	0x43, 0x63, 0x54, 0xb8, 0xb6, 0x5e, 0x84, 0xab, 0x16, 0xdd, 0x96, 0x75, 0xa9, 0xff, 0x6b, 0x4f,
	0xff, 0x61, 0x9d, 0x21, 0xb0, 0xc9, 0x10, 0xf8, 0xc8, 0x10, 0x78, 0xcd, 0x91, 0xb5, 0xc9, 0x91,
	0xf5, 0x96, 0x23, 0x6b, 0xdc, 0x8b, 0x98, 0x9c, 0x25, 0x53, 0x1c, 0xf0, 0xc5, 0xb6, 0xc7, 0xcb,
	0x1f, 0xa5, 0xbe, 0x94, 0x6b, 0x95, 0xe9, 0x92, 0x8a, 0xa9, 0xad, 0xdb, 0xbc, 0xfa, 0x0a, 0x00,
	0x00, 0xff, 0xff, 0x1b, 0x7f, 0x45, 0x03, 0xee, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AccountHistories) > 0 {
		for iNdEx := len(m.AccountHistories) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AccountHistories[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Stakes) > 0 {
		for iNdEx := len(m.Stakes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Stakes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.LastStakeId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastStakeId))
		i--
		dAtA[i] = 0x20
	}
	if m.LastGaugeId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastGaugeId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Gauges) > 0 {
		for iNdEx := len(m.Gauges) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Gauges[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Gauges) > 0 {
		for _, e := range m.Gauges {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.LastGaugeId != 0 {
		n += 1 + sovGenesis(uint64(m.LastGaugeId))
	}
	if m.LastStakeId != 0 {
		n += 1 + sovGenesis(uint64(m.LastStakeId))
	}
	if len(m.Stakes) > 0 {
		for _, e := range m.Stakes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.AccountHistories) > 0 {
		for _, e := range m.AccountHistories {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gauges", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Gauges = append(m.Gauges, &Gauge{})
			if err := m.Gauges[len(m.Gauges)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastGaugeId", wireType)
			}
			m.LastGaugeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastGaugeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastStakeId", wireType)
			}
			m.LastStakeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastStakeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stakes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stakes = append(m.Stakes, &Stake{})
			if err := m.Stakes[len(m.Stakes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountHistories", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountHistories = append(m.AccountHistories, &AccountHistory{})
			if err := m.AccountHistories[len(m.AccountHistories)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
