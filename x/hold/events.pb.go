// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: provenance/hold/v1/events.proto

package hold

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

// EventHoldAdded is an event indicating that some funds were placed on hold in an account.
type EventHoldAdded struct {
	// address is the bech32 address string of the account with the funds.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// amount is a Coins string of the funds placed on hold.
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	// reason is a human-readable indicator of why this hold was added.
	Reason string `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (m *EventHoldAdded) Reset()         { *m = EventHoldAdded{} }
func (m *EventHoldAdded) String() string { return proto.CompactTextString(m) }
func (*EventHoldAdded) ProtoMessage()    {}
func (*EventHoldAdded) Descriptor() ([]byte, []int) {
	return fileDescriptor_3be3cec6aa38cf10, []int{0}
}
func (m *EventHoldAdded) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventHoldAdded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventHoldAdded.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventHoldAdded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventHoldAdded.Merge(m, src)
}
func (m *EventHoldAdded) XXX_Size() int {
	return m.Size()
}
func (m *EventHoldAdded) XXX_DiscardUnknown() {
	xxx_messageInfo_EventHoldAdded.DiscardUnknown(m)
}

var xxx_messageInfo_EventHoldAdded proto.InternalMessageInfo

func (m *EventHoldAdded) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *EventHoldAdded) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *EventHoldAdded) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

// EventHoldReleased is an event indicating that some funds were released from hold for an account.
type EventHoldReleased struct {
	// address is the bech32 address string of the account with the funds.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// amount is a Coins string of the funds released from hold.
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *EventHoldReleased) Reset()         { *m = EventHoldReleased{} }
func (m *EventHoldReleased) String() string { return proto.CompactTextString(m) }
func (*EventHoldReleased) ProtoMessage()    {}
func (*EventHoldReleased) Descriptor() ([]byte, []int) {
	return fileDescriptor_3be3cec6aa38cf10, []int{1}
}
func (m *EventHoldReleased) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventHoldReleased) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventHoldReleased.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventHoldReleased) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventHoldReleased.Merge(m, src)
}
func (m *EventHoldReleased) XXX_Size() int {
	return m.Size()
}
func (m *EventHoldReleased) XXX_DiscardUnknown() {
	xxx_messageInfo_EventHoldReleased.DiscardUnknown(m)
}

var xxx_messageInfo_EventHoldReleased proto.InternalMessageInfo

func (m *EventHoldReleased) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *EventHoldReleased) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

// EventUnlockVestingAccounts is an event indicating that a vesting account has been unlocked.
type EventVestingAccountUnlocked struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *EventVestingAccountUnlocked) Reset()         { *m = EventVestingAccountUnlocked{} }
func (m *EventVestingAccountUnlocked) String() string { return proto.CompactTextString(m) }
func (*EventVestingAccountUnlocked) ProtoMessage()    {}
func (*EventVestingAccountUnlocked) Descriptor() ([]byte, []int) {
	return fileDescriptor_3be3cec6aa38cf10, []int{2}
}
func (m *EventVestingAccountUnlocked) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventVestingAccountUnlocked) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventVestingAccountUnlocked.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventVestingAccountUnlocked) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventVestingAccountUnlocked.Merge(m, src)
}
func (m *EventVestingAccountUnlocked) XXX_Size() int {
	return m.Size()
}
func (m *EventVestingAccountUnlocked) XXX_DiscardUnknown() {
	xxx_messageInfo_EventVestingAccountUnlocked.DiscardUnknown(m)
}

var xxx_messageInfo_EventVestingAccountUnlocked proto.InternalMessageInfo

func (m *EventVestingAccountUnlocked) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*EventHoldAdded)(nil), "provenance.hold.v1.EventHoldAdded")
	proto.RegisterType((*EventHoldReleased)(nil), "provenance.hold.v1.EventHoldReleased")
	proto.RegisterType((*EventVestingAccountUnlocked)(nil), "provenance.hold.v1.EventVestingAccountUnlocked")
}

func init() { proto.RegisterFile("provenance/hold/v1/events.proto", fileDescriptor_3be3cec6aa38cf10) }

var fileDescriptor_3be3cec6aa38cf10 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0x28, 0xca, 0x2f,
	0x4b, 0xcd, 0x4b, 0xcc, 0x4b, 0x4e, 0xd5, 0xcf, 0xc8, 0xcf, 0x49, 0xd1, 0x2f, 0x33, 0xd4, 0x4f,
	0x2d, 0x4b, 0xcd, 0x2b, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x42, 0x28, 0xd0,
	0x03, 0x29, 0xd0, 0x2b, 0x33, 0x94, 0x92, 0x4c, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0x8e, 0x07, 0xab,
	0xd0, 0x87, 0x70, 0x20, 0xca, 0x95, 0x4a, 0xb8, 0xf8, 0x5c, 0x41, 0xda, 0x3d, 0xf2, 0x73, 0x52,
	0x1c, 0x53, 0x52, 0x52, 0x53, 0x84, 0x8c, 0xb8, 0xd8, 0x13, 0x53, 0x52, 0x8a, 0x52, 0x8b, 0x8b,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x9d, 0x24, 0x2e, 0x6d, 0xd1, 0x15, 0x81, 0x6a, 0x72, 0x84,
	0xc8, 0x04, 0x97, 0x14, 0x65, 0xe6, 0xa5, 0x07, 0xc1, 0x14, 0x0a, 0x89, 0x71, 0xb1, 0x25, 0xe6,
	0xe6, 0x97, 0xe6, 0x95, 0x48, 0x30, 0x81, 0xb4, 0x04, 0x41, 0x79, 0x20, 0xf1, 0xa2, 0xd4, 0xc4,
	0xe2, 0xfc, 0x3c, 0x09, 0x66, 0x88, 0x38, 0x84, 0xa7, 0x14, 0xcf, 0x25, 0x08, 0xb7, 0x35, 0x28,
	0x35, 0x27, 0x35, 0xb1, 0x98, 0xba, 0x16, 0x2b, 0x05, 0x72, 0x49, 0x83, 0x2d, 0x08, 0x4b, 0x2d,
	0x2e, 0xc9, 0xcc, 0x4b, 0x77, 0x4c, 0x4e, 0x06, 0x09, 0x87, 0xe6, 0xe5, 0xe4, 0x27, 0x67, 0x93,
	0x67, 0x95, 0x53, 0xec, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7,
	0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x70, 0x89, 0x66,
	0xe6, 0xeb, 0x61, 0x86, 0x7a, 0x00, 0x63, 0x94, 0x56, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e,
	0x72, 0x7e, 0xae, 0x3e, 0x42, 0x81, 0x6e, 0x66, 0x3e, 0x12, 0x4f, 0xbf, 0x02, 0x1c, 0x8f, 0x49,
	0x6c, 0xe0, 0xf8, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x03, 0xa5, 0x4b, 0x1c, 0xe1, 0x01,
	0x00, 0x00,
}

func (m *EventHoldAdded) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventHoldAdded) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventHoldAdded) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Reason) > 0 {
		i -= len(m.Reason)
		copy(dAtA[i:], m.Reason)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Reason)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventHoldReleased) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventHoldReleased) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventHoldReleased) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventVestingAccountUnlocked) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventVestingAccountUnlocked) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventVestingAccountUnlocked) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventHoldAdded) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Reason)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventHoldReleased) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventVestingAccountUnlocked) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventHoldAdded) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventHoldAdded: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventHoldAdded: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventHoldReleased) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventHoldReleased: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventHoldReleased: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventVestingAccountUnlocked) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventVestingAccountUnlocked: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventVestingAccountUnlocked: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
