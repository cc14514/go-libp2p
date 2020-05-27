// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: identify.proto

package identify_pb

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

type Delta struct {
	// new protocols now serviced by the peer.
	AddedProtocols []string `protobuf:"bytes,1,rep,name=added_protocols,json=addedProtocols" json:"added_protocols,omitempty"`
	// protocols dropped by the peer.
	RmProtocols          []string `protobuf:"bytes,2,rep,name=rm_protocols,json=rmProtocols" json:"rm_protocols,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Delta) Reset()         { *m = Delta{} }
func (m *Delta) String() string { return proto.CompactTextString(m) }
func (*Delta) ProtoMessage()    {}
func (*Delta) Descriptor() ([]byte, []int) {
	return fileDescriptor_83f1e7e6b485409f, []int{0}
}
func (m *Delta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Delta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Delta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Delta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Delta.Merge(m, src)
}
func (m *Delta) XXX_Size() int {
	return m.Size()
}
func (m *Delta) XXX_DiscardUnknown() {
	xxx_messageInfo_Delta.DiscardUnknown(m)
}

var xxx_messageInfo_Delta proto.InternalMessageInfo

func (m *Delta) GetAddedProtocols() []string {
	if m != nil {
		return m.AddedProtocols
	}
	return nil
}

func (m *Delta) GetRmProtocols() []string {
	if m != nil {
		return m.RmProtocols
	}
	return nil
}

type Identify struct {
	// protocolVersion determines compatibility between peers
	ProtocolVersion *string `protobuf:"bytes,5,opt,name=protocolVersion" json:"protocolVersion,omitempty"`
	// agentVersion is like a UserAgent string in browsers, or client version in bittorrent
	// includes the client name and client.
	AgentVersion *string `protobuf:"bytes,6,opt,name=agentVersion" json:"agentVersion,omitempty"`
	// publicKey is this node's public key (which also gives its node.ID)
	// - may not need to be sent, as secure channel implies it has been sent.
	// - then again, if we change / disable secure channel, may still want it.
	PublicKey []byte `protobuf:"bytes,1,opt,name=publicKey" json:"publicKey,omitempty"`
	// listenAddrs are the multiaddrs the sender node listens for open connections on
	ListenAddrs [][]byte `protobuf:"bytes,2,rep,name=listenAddrs" json:"listenAddrs,omitempty"`
	// oservedAddr is the multiaddr of the remote endpoint that the sender node perceives
	// this is useful information to convey to the other side, as it helps the remote endpoint
	// determine whether its connection to the local peer goes through NAT.
	ObservedAddr []byte `protobuf:"bytes,4,opt,name=observedAddr" json:"observedAddr,omitempty"`
	// protocols are the services this node is running
	Protocols []string `protobuf:"bytes,3,rep,name=protocols" json:"protocols,omitempty"`
	// a delta update is incompatible with everything else. If this field is included, none of the others can appear.
	Delta *Delta `protobuf:"bytes,7,opt,name=delta" json:"delta,omitempty"`
	// add by liangc : take groupid to join private routing table.
	Groupid []byte `protobuf:"bytes,9,opt,name=groupid" json:"groupid,omitempty"`
	// signedPeerRecord contains a serialized SignedEnvelope containing a PeerRecord,
	// signed by the sending node. It contains the same addresses as the listenAddrs field, but
	// in a form that lets us share authenticated addrs with other peers.
	// see github.com/libp2p/go-libp2p-core/record/pb/envelope.proto and
	// github.com/libp2p/go-libp2p-core/peer/pb/peer_record.proto for message definitions.
	SignedPeerRecord     []byte   `protobuf:"bytes,8,opt,name=signedPeerRecord" json:"signedPeerRecord,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Identify) Reset()         { *m = Identify{} }
func (m *Identify) String() string { return proto.CompactTextString(m) }
func (*Identify) ProtoMessage()    {}
func (*Identify) Descriptor() ([]byte, []int) {
	return fileDescriptor_83f1e7e6b485409f, []int{1}
}
func (m *Identify) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Identify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Identify.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Identify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identify.Merge(m, src)
}
func (m *Identify) XXX_Size() int {
	return m.Size()
}
func (m *Identify) XXX_DiscardUnknown() {
	xxx_messageInfo_Identify.DiscardUnknown(m)
}

var xxx_messageInfo_Identify proto.InternalMessageInfo

func (m *Identify) GetProtocolVersion() string {
	if m != nil && m.ProtocolVersion != nil {
		return *m.ProtocolVersion
	}
	return ""
}

func (m *Identify) GetAgentVersion() string {
	if m != nil && m.AgentVersion != nil {
		return *m.AgentVersion
	}
	return ""
}

func (m *Identify) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *Identify) GetListenAddrs() [][]byte {
	if m != nil {
		return m.ListenAddrs
	}
	return nil
}

func (m *Identify) GetObservedAddr() []byte {
	if m != nil {
		return m.ObservedAddr
	}
	return nil
}

func (m *Identify) GetProtocols() []string {
	if m != nil {
		return m.Protocols
	}
	return nil
}

func (m *Identify) GetDelta() *Delta {
	if m != nil {
		return m.Delta
	}
	return nil
}

func (m *Identify) GetGroupid() []byte {
	if m != nil {
		return m.Groupid
	}
	return nil
}

func (m *Identify) GetSignedPeerRecord() []byte {
	if m != nil {
		return m.SignedPeerRecord
	}
	return nil
}

func init() {
	proto.RegisterType((*Delta)(nil), "identify.pb.Delta")
	proto.RegisterType((*Identify)(nil), "identify.pb.Identify")
}

func init() { proto.RegisterFile("identify.proto", fileDescriptor_83f1e7e6b485409f) }

var fileDescriptor_83f1e7e6b485409f = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x14, 0x45, 0x99, 0xd6, 0xda, 0xe6, 0x25, 0xb4, 0x32, 0xab, 0x59, 0x48, 0x89, 0xd9, 0x38, 0xb8,
	0xc8, 0xc2, 0x3f, 0x50, 0xdc, 0x88, 0x9b, 0x32, 0x82, 0x5b, 0x49, 0xf2, 0x9e, 0x61, 0x20, 0xcd,
	0x84, 0x49, 0x2a, 0xf4, 0x0f, 0xc5, 0x95, 0x9f, 0x20, 0xf9, 0x12, 0xc9, 0xd4, 0x34, 0xa9, 0x2e,
	0xdf, 0xe1, 0x30, 0x77, 0xee, 0x85, 0xa5, 0x46, 0x2a, 0x1b, 0xfd, 0xb6, 0x8f, 0x2b, 0x6b, 0x1a,
	0xc3, 0xfd, 0xe1, 0x4e, 0xa3, 0x67, 0x98, 0x3d, 0x50, 0xd1, 0x24, 0xfc, 0x1a, 0x56, 0x09, 0x22,
	0xe1, 0xab, 0x93, 0x32, 0x53, 0xd4, 0x82, 0x85, 0x53, 0xe9, 0xa9, 0xa5, 0xc3, 0x9b, 0x9e, 0xf2,
	0x2b, 0x08, 0xec, 0x76, 0x64, 0x4d, 0x9c, 0xe5, 0xdb, 0xed, 0x51, 0x89, 0x3e, 0x27, 0xb0, 0x78,
	0xfc, 0x0d, 0xe1, 0x12, 0x56, 0xbd, 0xfc, 0x42, 0xb6, 0xd6, 0xa6, 0x14, 0xb3, 0x90, 0x49, 0x4f,
	0xfd, 0xc5, 0x3c, 0x82, 0x20, 0xc9, 0xa9, 0x6c, 0x7a, 0xed, 0xdc, 0x69, 0x27, 0x8c, 0x5f, 0x82,
	0x57, 0xed, 0xd2, 0x42, 0x67, 0x4f, 0xb4, 0x17, 0x2c, 0x64, 0x32, 0x50, 0x03, 0xe0, 0x21, 0xf8,
	0x85, 0xae, 0x1b, 0x2a, 0xef, 0x10, 0xed, 0xe1, 0x6b, 0x81, 0x1a, 0xa3, 0x2e, 0xc3, 0xa4, 0x35,
	0xd9, 0x77, 0xc2, 0x0e, 0x88, 0x33, 0xf7, 0xc4, 0x09, 0x73, 0x19, 0xc7, 0x7a, 0x53, 0x57, 0x6f,
	0x00, 0x5c, 0xc2, 0x0c, 0xbb, 0xc5, 0xc4, 0x3c, 0x64, 0xd2, 0xbf, 0xe5, 0xf1, 0x68, 0xce, 0xd8,
	0x6d, 0xa9, 0x0e, 0x02, 0x17, 0x30, 0xcf, 0xad, 0xd9, 0x55, 0x1a, 0x85, 0xe7, 0x62, 0xfa, 0x93,
	0xdf, 0xc0, 0x45, 0xad, 0xf3, 0x92, 0x70, 0x43, 0x64, 0x15, 0x65, 0xc6, 0xa2, 0x58, 0x38, 0xe5,
	0x1f, 0xbf, 0x0f, 0x3e, 0xda, 0x35, 0xfb, 0x6a, 0xd7, 0xec, 0xbb, 0x5d, 0xb3, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xb4, 0xe0, 0x7c, 0x15, 0xcd, 0x01, 0x00, 0x00,
}

func (m *Delta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Delta) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Delta) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.RmProtocols) > 0 {
		for iNdEx := len(m.RmProtocols) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.RmProtocols[iNdEx])
			copy(dAtA[i:], m.RmProtocols[iNdEx])
			i = encodeVarintIdentify(dAtA, i, uint64(len(m.RmProtocols[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.AddedProtocols) > 0 {
		for iNdEx := len(m.AddedProtocols) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AddedProtocols[iNdEx])
			copy(dAtA[i:], m.AddedProtocols[iNdEx])
			i = encodeVarintIdentify(dAtA, i, uint64(len(m.AddedProtocols[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Identify) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Identify) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Identify) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Groupid != nil {
		i -= len(m.Groupid)
		copy(dAtA[i:], m.Groupid)
		i = encodeVarintIdentify(dAtA, i, uint64(len(m.Groupid)))
		i--
		dAtA[i] = 0x4a
	}
	if m.SignedPeerRecord != nil {
		i -= len(m.SignedPeerRecord)
		copy(dAtA[i:], m.SignedPeerRecord)
		i = encodeVarintIdentify(dAtA, i, uint64(len(m.SignedPeerRecord)))
		i--
		dAtA[i] = 0x42
	}
	if m.Delta != nil {
		{
			size, err := m.Delta.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIdentify(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.AgentVersion != nil {
		i -= len(*m.AgentVersion)
		copy(dAtA[i:], *m.AgentVersion)
		i = encodeVarintIdentify(dAtA, i, uint64(len(*m.AgentVersion)))
		i--
		dAtA[i] = 0x32
	}
	if m.ProtocolVersion != nil {
		i -= len(*m.ProtocolVersion)
		copy(dAtA[i:], *m.ProtocolVersion)
		i = encodeVarintIdentify(dAtA, i, uint64(len(*m.ProtocolVersion)))
		i--
		dAtA[i] = 0x2a
	}
	if m.ObservedAddr != nil {
		i -= len(m.ObservedAddr)
		copy(dAtA[i:], m.ObservedAddr)
		i = encodeVarintIdentify(dAtA, i, uint64(len(m.ObservedAddr)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Protocols) > 0 {
		for iNdEx := len(m.Protocols) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Protocols[iNdEx])
			copy(dAtA[i:], m.Protocols[iNdEx])
			i = encodeVarintIdentify(dAtA, i, uint64(len(m.Protocols[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ListenAddrs) > 0 {
		for iNdEx := len(m.ListenAddrs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ListenAddrs[iNdEx])
			copy(dAtA[i:], m.ListenAddrs[iNdEx])
			i = encodeVarintIdentify(dAtA, i, uint64(len(m.ListenAddrs[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.PublicKey != nil {
		i -= len(m.PublicKey)
		copy(dAtA[i:], m.PublicKey)
		i = encodeVarintIdentify(dAtA, i, uint64(len(m.PublicKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIdentify(dAtA []byte, offset int, v uint64) int {
	offset -= sovIdentify(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Delta) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AddedProtocols) > 0 {
		for _, s := range m.AddedProtocols {
			l = len(s)
			n += 1 + l + sovIdentify(uint64(l))
		}
	}
	if len(m.RmProtocols) > 0 {
		for _, s := range m.RmProtocols {
			l = len(s)
			n += 1 + l + sovIdentify(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Identify) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PublicKey != nil {
		l = len(m.PublicKey)
		n += 1 + l + sovIdentify(uint64(l))
	}
	if len(m.ListenAddrs) > 0 {
		for _, b := range m.ListenAddrs {
			l = len(b)
			n += 1 + l + sovIdentify(uint64(l))
		}
	}
	if len(m.Protocols) > 0 {
		for _, s := range m.Protocols {
			l = len(s)
			n += 1 + l + sovIdentify(uint64(l))
		}
	}
	if m.ObservedAddr != nil {
		l = len(m.ObservedAddr)
		n += 1 + l + sovIdentify(uint64(l))
	}
	if m.ProtocolVersion != nil {
		l = len(*m.ProtocolVersion)
		n += 1 + l + sovIdentify(uint64(l))
	}
	if m.AgentVersion != nil {
		l = len(*m.AgentVersion)
		n += 1 + l + sovIdentify(uint64(l))
	}
	if m.Delta != nil {
		l = m.Delta.Size()
		n += 1 + l + sovIdentify(uint64(l))
	}
	if m.SignedPeerRecord != nil {
		l = len(m.SignedPeerRecord)
		n += 1 + l + sovIdentify(uint64(l))
	}
	if m.Groupid != nil {
		l = len(m.Groupid)
		n += 1 + l + sovIdentify(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovIdentify(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIdentify(x uint64) (n int) {
	return sovIdentify(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Delta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIdentify
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
			return fmt.Errorf("proto: Delta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Delta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddedProtocols", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentify
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
				return ErrInvalidLengthIdentify
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AddedProtocols = append(m.AddedProtocols, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RmProtocols", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentify
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
				return ErrInvalidLengthIdentify
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RmProtocols = append(m.RmProtocols, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIdentify(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIdentify
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthIdentify
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
func (m *Identify) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIdentify
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
			return fmt.Errorf("proto: Identify: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Identify: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentify
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthIdentify
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKey = append(m.PublicKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PublicKey == nil {
				m.PublicKey = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListenAddrs", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentify
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthIdentify
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ListenAddrs = append(m.ListenAddrs, make([]byte, postIndex-iNdEx))
			copy(m.ListenAddrs[len(m.ListenAddrs)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Protocols", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentify
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
				return ErrInvalidLengthIdentify
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Protocols = append(m.Protocols, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObservedAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentify
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthIdentify
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ObservedAddr = append(m.ObservedAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.ObservedAddr == nil {
				m.ObservedAddr = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtocolVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentify
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
				return ErrInvalidLengthIdentify
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.ProtocolVersion = &s
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AgentVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentify
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
				return ErrInvalidLengthIdentify
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.AgentVersion = &s
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentify
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
				return ErrInvalidLengthIdentify
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIdentify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Delta == nil {
				m.Delta = &Delta{}
			}
			if err := m.Delta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignedPeerRecord", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentify
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthIdentify
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignedPeerRecord = append(m.SignedPeerRecord[:0], dAtA[iNdEx:postIndex]...)
			if m.SignedPeerRecord == nil {
				m.SignedPeerRecord = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Groupid", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentify
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthIdentify
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Groupid = append(m.Groupid[:0], dAtA[iNdEx:postIndex]...)
			if m.Groupid == nil {
				m.Groupid = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIdentify(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIdentify
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthIdentify
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
func skipIdentify(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIdentify
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
					return 0, ErrIntOverflowIdentify
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
					return 0, ErrIntOverflowIdentify
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
				return 0, ErrInvalidLengthIdentify
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIdentify
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIdentify
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIdentify        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIdentify          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIdentify = fmt.Errorf("proto: unexpected end of group")
)
