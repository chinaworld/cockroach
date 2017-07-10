// Code generated by protoc-gen-gogo.
// source: cockroach/pkg/server/diagnostics.proto
// DO NOT EDIT!

/*
	Package server is a generated protocol buffer package.

	It is generated from these files:
		cockroach/pkg/server/diagnostics.proto

	It has these top-level messages:
		DiagnosticReport
		NodeInfo
		StoreInfo
*/
package server

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import cockroach_sql "github.com/cockroachdb/cockroach/pkg/roachpb"
import cockroach_sql_sqlbase1 "github.com/cockroachdb/cockroach/pkg/sql/sqlbase"

import github_com_cockroachdb_cockroach_pkg_roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"

import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

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

type DiagnosticReport struct {
	Node                NodeInfo                                     `protobuf:"bytes,1,opt,name=node" json:"node"`
	Stores              []StoreInfo                                  `protobuf:"bytes,2,rep,name=stores" json:"stores"`
	Schema              []cockroach_sql_sqlbase1.TableDescriptor     `protobuf:"bytes,3,rep,name=schema" json:"schema"`
	SqlStats            []cockroach_sql.CollectedStatementStatistics `protobuf:"bytes,4,rep,name=sql_stats,json=sqlStats" json:"sql_stats"`
	UnimplementedErrors map[string]int64                             `protobuf:"bytes,5,rep,name=unimplemented_errors,json=unimplementedErrors" json:"unimplemented_errors,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (m *DiagnosticReport) Reset()                    { *m = DiagnosticReport{} }
func (m *DiagnosticReport) String() string            { return proto.CompactTextString(m) }
func (*DiagnosticReport) ProtoMessage()               {}
func (*DiagnosticReport) Descriptor() ([]byte, []int) { return fileDescriptorDiagnostics, []int{0} }

type NodeInfo struct {
	NodeID     github_com_cockroachdb_cockroach_pkg_roachpb.NodeID `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3,casttype=github.com/cockroachdb/cockroach/pkg/roachpb.NodeID" json:"node_id,omitempty"`
	Bytes      int64                                               `protobuf:"varint,2,opt,name=bytes,proto3" json:"bytes,omitempty"`
	KeyCount   int64                                               `protobuf:"varint,3,opt,name=key_count,json=keyCount,proto3" json:"key_count,omitempty"`
	RangeCount int64                                               `protobuf:"varint,4,opt,name=range_count,json=rangeCount,proto3" json:"range_count,omitempty"`
}

func (m *NodeInfo) Reset()                    { *m = NodeInfo{} }
func (m *NodeInfo) String() string            { return proto.CompactTextString(m) }
func (*NodeInfo) ProtoMessage()               {}
func (*NodeInfo) Descriptor() ([]byte, []int) { return fileDescriptorDiagnostics, []int{1} }

type StoreInfo struct {
	NodeID     github_com_cockroachdb_cockroach_pkg_roachpb.NodeID  `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3,casttype=github.com/cockroachdb/cockroach/pkg/roachpb.NodeID" json:"node_id,omitempty"`
	StoreID    github_com_cockroachdb_cockroach_pkg_roachpb.StoreID `protobuf:"varint,2,opt,name=store_id,json=storeId,proto3,casttype=github.com/cockroachdb/cockroach/pkg/roachpb.StoreID" json:"store_id,omitempty"`
	Bytes      int64                                                `protobuf:"varint,3,opt,name=bytes,proto3" json:"bytes,omitempty"`
	KeyCount   int64                                                `protobuf:"varint,4,opt,name=key_count,json=keyCount,proto3" json:"key_count,omitempty"`
	RangeCount int64                                                `protobuf:"varint,5,opt,name=range_count,json=rangeCount,proto3" json:"range_count,omitempty"`
}

func (m *StoreInfo) Reset()                    { *m = StoreInfo{} }
func (m *StoreInfo) String() string            { return proto.CompactTextString(m) }
func (*StoreInfo) ProtoMessage()               {}
func (*StoreInfo) Descriptor() ([]byte, []int) { return fileDescriptorDiagnostics, []int{2} }

func init() {
	proto.RegisterType((*DiagnosticReport)(nil), "cockroach.server.DiagnosticReport")
	proto.RegisterType((*NodeInfo)(nil), "cockroach.server.NodeInfo")
	proto.RegisterType((*StoreInfo)(nil), "cockroach.server.StoreInfo")
}
func (m *DiagnosticReport) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DiagnosticReport) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintDiagnostics(dAtA, i, uint64(m.Node.Size()))
	n1, err := m.Node.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.Stores) > 0 {
		for _, msg := range m.Stores {
			dAtA[i] = 0x12
			i++
			i = encodeVarintDiagnostics(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Schema) > 0 {
		for _, msg := range m.Schema {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintDiagnostics(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.SqlStats) > 0 {
		for _, msg := range m.SqlStats {
			dAtA[i] = 0x22
			i++
			i = encodeVarintDiagnostics(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.UnimplementedErrors) > 0 {
		keysForUnimplementedErrors := make([]string, 0, len(m.UnimplementedErrors))
		for k := range m.UnimplementedErrors {
			keysForUnimplementedErrors = append(keysForUnimplementedErrors, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForUnimplementedErrors)
		for _, k := range keysForUnimplementedErrors {
			dAtA[i] = 0x2a
			i++
			v := m.UnimplementedErrors[string(k)]
			mapSize := 1 + len(k) + sovDiagnostics(uint64(len(k))) + 1 + sovDiagnostics(uint64(v))
			i = encodeVarintDiagnostics(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintDiagnostics(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x10
			i++
			i = encodeVarintDiagnostics(dAtA, i, uint64(v))
		}
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
	if m.NodeID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintDiagnostics(dAtA, i, uint64(m.NodeID))
	}
	if m.Bytes != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintDiagnostics(dAtA, i, uint64(m.Bytes))
	}
	if m.KeyCount != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintDiagnostics(dAtA, i, uint64(m.KeyCount))
	}
	if m.RangeCount != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintDiagnostics(dAtA, i, uint64(m.RangeCount))
	}
	return i, nil
}

func (m *StoreInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoreInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.NodeID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintDiagnostics(dAtA, i, uint64(m.NodeID))
	}
	if m.StoreID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintDiagnostics(dAtA, i, uint64(m.StoreID))
	}
	if m.Bytes != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintDiagnostics(dAtA, i, uint64(m.Bytes))
	}
	if m.KeyCount != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintDiagnostics(dAtA, i, uint64(m.KeyCount))
	}
	if m.RangeCount != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintDiagnostics(dAtA, i, uint64(m.RangeCount))
	}
	return i, nil
}

func encodeFixed64Diagnostics(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Diagnostics(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintDiagnostics(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *DiagnosticReport) Size() (n int) {
	var l int
	_ = l
	l = m.Node.Size()
	n += 1 + l + sovDiagnostics(uint64(l))
	if len(m.Stores) > 0 {
		for _, e := range m.Stores {
			l = e.Size()
			n += 1 + l + sovDiagnostics(uint64(l))
		}
	}
	if len(m.Schema) > 0 {
		for _, e := range m.Schema {
			l = e.Size()
			n += 1 + l + sovDiagnostics(uint64(l))
		}
	}
	if len(m.SqlStats) > 0 {
		for _, e := range m.SqlStats {
			l = e.Size()
			n += 1 + l + sovDiagnostics(uint64(l))
		}
	}
	if len(m.UnimplementedErrors) > 0 {
		for k, v := range m.UnimplementedErrors {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovDiagnostics(uint64(len(k))) + 1 + sovDiagnostics(uint64(v))
			n += mapEntrySize + 1 + sovDiagnostics(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *NodeInfo) Size() (n int) {
	var l int
	_ = l
	if m.NodeID != 0 {
		n += 1 + sovDiagnostics(uint64(m.NodeID))
	}
	if m.Bytes != 0 {
		n += 1 + sovDiagnostics(uint64(m.Bytes))
	}
	if m.KeyCount != 0 {
		n += 1 + sovDiagnostics(uint64(m.KeyCount))
	}
	if m.RangeCount != 0 {
		n += 1 + sovDiagnostics(uint64(m.RangeCount))
	}
	return n
}

func (m *StoreInfo) Size() (n int) {
	var l int
	_ = l
	if m.NodeID != 0 {
		n += 1 + sovDiagnostics(uint64(m.NodeID))
	}
	if m.StoreID != 0 {
		n += 1 + sovDiagnostics(uint64(m.StoreID))
	}
	if m.Bytes != 0 {
		n += 1 + sovDiagnostics(uint64(m.Bytes))
	}
	if m.KeyCount != 0 {
		n += 1 + sovDiagnostics(uint64(m.KeyCount))
	}
	if m.RangeCount != 0 {
		n += 1 + sovDiagnostics(uint64(m.RangeCount))
	}
	return n
}

func sovDiagnostics(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDiagnostics(x uint64) (n int) {
	return sovDiagnostics(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DiagnosticReport) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDiagnostics
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
			return fmt.Errorf("proto: DiagnosticReport: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DiagnosticReport: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiagnostics
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
				return ErrInvalidLengthDiagnostics
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Node.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stores", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiagnostics
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
				return ErrInvalidLengthDiagnostics
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stores = append(m.Stores, StoreInfo{})
			if err := m.Stores[len(m.Stores)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Schema", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiagnostics
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
				return ErrInvalidLengthDiagnostics
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Schema = append(m.Schema, cockroach_sql_sqlbase1.TableDescriptor{})
			if err := m.Schema[len(m.Schema)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SqlStats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiagnostics
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
				return ErrInvalidLengthDiagnostics
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SqlStats = append(m.SqlStats, cockroach_sql.CollectedStatementStatistics{})
			if err := m.SqlStats[len(m.SqlStats)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnimplementedErrors", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiagnostics
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
				return ErrInvalidLengthDiagnostics
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiagnostics
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
					return ErrIntOverflowDiagnostics
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
				return ErrInvalidLengthDiagnostics
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.UnimplementedErrors == nil {
				m.UnimplementedErrors = make(map[string]int64)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowDiagnostics
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
				var mapvalue int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowDiagnostics
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					mapvalue |= (int64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.UnimplementedErrors[mapkey] = mapvalue
			} else {
				var mapvalue int64
				m.UnimplementedErrors[mapkey] = mapvalue
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDiagnostics(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDiagnostics
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
				return ErrIntOverflowDiagnostics
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
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
			}
			m.NodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiagnostics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NodeID |= (github_com_cockroachdb_cockroach_pkg_roachpb.NodeID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bytes", wireType)
			}
			m.Bytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiagnostics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Bytes |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyCount", wireType)
			}
			m.KeyCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiagnostics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyCount |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeCount", wireType)
			}
			m.RangeCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiagnostics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RangeCount |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDiagnostics(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDiagnostics
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
func (m *StoreInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDiagnostics
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
			return fmt.Errorf("proto: StoreInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoreInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
			}
			m.NodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiagnostics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NodeID |= (github_com_cockroachdb_cockroach_pkg_roachpb.NodeID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoreID", wireType)
			}
			m.StoreID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiagnostics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StoreID |= (github_com_cockroachdb_cockroach_pkg_roachpb.StoreID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bytes", wireType)
			}
			m.Bytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiagnostics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Bytes |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyCount", wireType)
			}
			m.KeyCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiagnostics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyCount |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeCount", wireType)
			}
			m.RangeCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiagnostics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RangeCount |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDiagnostics(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDiagnostics
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
func skipDiagnostics(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDiagnostics
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
					return 0, ErrIntOverflowDiagnostics
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
					return 0, ErrIntOverflowDiagnostics
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
				return 0, ErrInvalidLengthDiagnostics
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDiagnostics
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
				next, err := skipDiagnostics(dAtA[start:])
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
	ErrInvalidLengthDiagnostics = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDiagnostics   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("cockroach/pkg/server/diagnostics.proto", fileDescriptorDiagnostics) }

var fileDescriptorDiagnostics = []byte{
	// 555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xbd, 0x8e, 0xda, 0x40,
	0x10, 0xc6, 0x18, 0x0c, 0x2c, 0x0d, 0x72, 0x28, 0x2c, 0x4e, 0x32, 0x08, 0x29, 0x27, 0x94, 0x48,
	0xb6, 0x74, 0x77, 0x45, 0x7e, 0xaa, 0x00, 0x17, 0xe9, 0x9a, 0x2b, 0x7c, 0x49, 0x73, 0x45, 0x90,
	0x7f, 0x26, 0xc6, 0xc2, 0x78, 0xcd, 0xee, 0xfa, 0x24, 0xde, 0x22, 0x7d, 0xde, 0x24, 0x4f, 0x40,
	0x99, 0x74, 0xa9, 0x50, 0xe2, 0xbc, 0x45, 0xaa, 0x68, 0xc7, 0x0e, 0x22, 0x90, 0x8b, 0x94, 0xe2,
	0x0a, 0x4b, 0x33, 0xfb, 0xfd, 0xec, 0xec, 0xcc, 0x98, 0x9c, 0xfa, 0xd4, 0x5f, 0x30, 0xea, 0xfa,
	0x73, 0x3b, 0x5d, 0x84, 0x36, 0x07, 0x76, 0x07, 0xcc, 0x0e, 0x22, 0x37, 0x4c, 0x28, 0x17, 0x91,
	0xcf, 0xad, 0x94, 0x51, 0x41, 0xf5, 0xce, 0x8e, 0x67, 0x15, 0x9c, 0x5e, 0x37, 0xa4, 0x21, 0x45,
	0xd0, 0x96, 0x51, 0xc1, 0xeb, 0x3d, 0xfe, 0xd3, 0x0f, 0xa3, 0xd4, 0xb3, 0xdd, 0x34, 0x9d, 0x71,
	0xe1, 0x8a, 0xd2, 0xae, 0xf7, 0xe4, 0xe0, 0xda, 0x55, 0x2c, 0x3f, 0xcf, 0xe5, 0x60, 0x73, 0xc1,
	0x32, 0x5f, 0x64, 0x0c, 0x82, 0x82, 0x3b, 0xfc, 0xa2, 0x92, 0xce, 0x74, 0x57, 0x90, 0x03, 0x29,
	0x65, 0x42, 0xbf, 0x20, 0xb5, 0x84, 0x06, 0x60, 0x28, 0x03, 0x65, 0xd4, 0x3e, 0xeb, 0x59, 0x87,
	0xe5, 0x59, 0xd7, 0x34, 0x80, 0xab, 0xe4, 0x3d, 0x1d, 0xd7, 0x36, 0xdb, 0x7e, 0xc5, 0x41, 0xb6,
	0xfe, 0x9c, 0x68, 0x5c, 0x50, 0x06, 0xdc, 0xa8, 0x0e, 0xd4, 0x51, 0xfb, 0xec, 0xe4, 0x58, 0x77,
	0x23, 0xf1, 0x3d, 0x61, 0x29, 0xd0, 0xa7, 0x44, 0xe3, 0xfe, 0x1c, 0x96, 0xae, 0xa1, 0xa2, 0xf4,
	0x74, 0x5f, 0xba, 0x8a, 0xad, 0xb2, 0x7c, 0xeb, 0x8d, 0xeb, 0xc5, 0x30, 0x05, 0xee, 0xb3, 0x28,
	0x15, 0x94, 0xed, 0x5c, 0x50, 0xab, 0x5f, 0x93, 0x16, 0x5f, 0xc5, 0x45, 0x2b, 0x8c, 0x1a, 0x1a,
	0x3d, 0x3d, 0x30, 0x9a, 0xd0, 0x38, 0x06, 0x5f, 0x40, 0x70, 0x23, 0x5c, 0x01, 0x4b, 0x48, 0x84,
	0x0c, 0x22, 0x1c, 0x46, 0xe9, 0xd6, 0xe4, 0xab, 0x58, 0x1e, 0x72, 0x3d, 0x21, 0xdd, 0x2c, 0x89,
	0x96, 0x69, 0x8c, 0x44, 0x08, 0x66, 0xc0, 0x18, 0x65, 0xdc, 0xa8, 0xa3, 0xf5, 0xcb, 0xe3, 0xe7,
	0x1d, 0x36, 0xd2, 0x7a, 0xbb, 0x2f, 0xbf, 0x44, 0xf5, 0x65, 0x22, 0xd8, 0xda, 0x79, 0x94, 0x1d,
	0x23, 0xbd, 0xd7, 0xc4, 0xb8, 0x4f, 0xa0, 0x77, 0x88, 0xba, 0x80, 0x35, 0x4e, 0xa4, 0xe5, 0xc8,
	0x50, 0xef, 0x92, 0xfa, 0x9d, 0x1b, 0x67, 0x60, 0x54, 0x07, 0xca, 0x48, 0x75, 0x8a, 0xe4, 0x45,
	0xf5, 0x99, 0x32, 0xfc, 0xa4, 0x90, 0xe6, 0xef, 0x09, 0xe9, 0xb7, 0xa4, 0x21, 0xa7, 0x33, 0x8b,
	0x02, 0x14, 0xd7, 0xc7, 0xaf, 0xf2, 0x6d, 0x5f, 0x43, 0x78, 0xfa, 0x73, 0xdb, 0x3f, 0x0f, 0x23,
	0x31, 0xcf, 0x3c, 0xcb, 0xa7, 0x4b, 0x7b, 0xf7, 0x9e, 0xc0, 0xb3, 0xff, 0xba, 0x69, 0xc5, 0xdc,
	0xa7, 0x8e, 0x26, 0x1d, 0xaf, 0x02, 0x59, 0x82, 0xb7, 0x16, 0x38, 0x70, 0x2c, 0x01, 0x13, 0xfd,
	0x84, 0xb4, 0x16, 0xb0, 0x9e, 0xf9, 0x34, 0x4b, 0x84, 0xa1, 0x22, 0xd2, 0x5c, 0xc0, 0x7a, 0x22,
	0x73, 0xbd, 0x4f, 0xda, 0xcc, 0x4d, 0x42, 0x28, 0xe1, 0x1a, 0xc2, 0x04, 0x8f, 0x90, 0x30, 0xfc,
	0x58, 0x25, 0xad, 0xdd, 0x9a, 0x3c, 0x68, 0xf5, 0xef, 0x48, 0x13, 0xd7, 0x4f, 0x9a, 0x57, 0xd1,
	0x7c, 0x92, 0x6f, 0xfb, 0x8d, 0xe2, 0x72, 0xe9, 0x7e, 0xf1, 0x5f, 0xee, 0xa5, 0xce, 0x69, 0xa0,
	0xe9, 0x7e, 0x77, 0xd4, 0x7b, 0xbb, 0x53, 0xfb, 0x77, 0x77, 0xea, 0x87, 0xdd, 0x19, 0x0f, 0x36,
	0xdf, 0xcd, 0xca, 0x26, 0x37, 0x95, 0xcf, 0xb9, 0xa9, 0x7c, 0xcd, 0x4d, 0xe5, 0x5b, 0x6e, 0x2a,
	0x1f, 0x7e, 0x98, 0x95, 0x5b, 0xad, 0xd8, 0x41, 0x4f, 0xc3, 0xff, 0xfa, 0xfc, 0x57, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x60, 0x42, 0x13, 0x53, 0x7c, 0x04, 0x00, 0x00,
}
