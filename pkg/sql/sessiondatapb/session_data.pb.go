// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sql/sessiondatapb/session_data.proto

package sessiondatapb

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

// BytesEncodeFormat is the configuration for bytes to string conversions.
type BytesEncodeFormat int32

const (
	// BytesEncodeHex uses the hex format: e'abc\n'::BYTES::STRING -> '\x61626312'.
	// This is the default, for compatibility with PostgreSQL.
	BytesEncodeHex BytesEncodeFormat = 0
	// BytesEncodeEscape uses the escaped format: e'abc\n'::BYTES::STRING -> 'abc\012'.
	BytesEncodeEscape BytesEncodeFormat = 1
	// BytesEncodeBase64 uses base64 encoding.
	BytesEncodeBase64 BytesEncodeFormat = 2
)

var BytesEncodeFormat_name = map[int32]string{
	0: "BytesEncodeHex",
	1: "BytesEncodeEscape",
	2: "BytesEncodeBase64",
}
var BytesEncodeFormat_value = map[string]int32{
	"BytesEncodeHex":    0,
	"BytesEncodeEscape": 1,
	"BytesEncodeBase64": 2,
}

func (BytesEncodeFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_session_data_c24dfdb3837a3a66, []int{0}
}

// VectorizeExecMode controls if an when the Executor executes queries using
// the columnar execution engine.
type VectorizeExecMode int32

const (
	// VectorizeOff means that columnar execution is disabled.
	VectorizeOff VectorizeExecMode = 0
	// Vectorize201Auto means that that any supported queries that use only
	// streaming operators (i.e. those that do not require any buffering) will
	// be run using the columnar execution. If any part of a query is not
	// supported by the vectorized execution engine, the whole query will fall
	// back to row execution.
	// This is the default setting in 20.1.
	Vectorize201Auto VectorizeExecMode = 1
	// VectorizeOn means that any supported queries will be run using the
	// columnar execution.
	VectorizeOn VectorizeExecMode = 2
	// VectorizeExperimentalAlways means that we attempt to vectorize all
	// queries; unsupported queries will fail. Mostly used for testing.
	VectorizeExperimentalAlways VectorizeExecMode = 3
)

var VectorizeExecMode_name = map[int32]string{
	0: "VectorizeOff",
	1: "Vectorize201Auto",
	2: "VectorizeOn",
	3: "VectorizeExperimentalAlways",
}
var VectorizeExecMode_value = map[string]int32{
	"VectorizeOff":                0,
	"Vectorize201Auto":            1,
	"VectorizeOn":                 2,
	"VectorizeExperimentalAlways": 3,
}

func (VectorizeExecMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_session_data_c24dfdb3837a3a66, []int{1}
}

// SessionData contains session parameters that are easily serializable and are
// required to be propagated to the remote nodes for the correct execution of
// DistSQL flows.
type SessionData struct {
	// Database indicates the "current" database for the purpose of resolving
	// names.
	Database string `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	// ApplicationName is the name of the application running the current
	// session. This can be used for logging and per-application statistics.
	ApplicationName string `protobuf:"bytes,2,opt,name=application_name,json=applicationName,proto3" json:"application_name,omitempty"`
	// User is the name of the user logged into the session.
	User string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	// DataConversion gives access to the data conversion configuration.
	DataConversionConfig DataConversionConfig `protobuf:"bytes,4,opt,name=data_conversion_config,json=dataConversionConfig,proto3" json:"data_conversion_config"`
	// VectorizeMode indicates which kinds of queries to use vectorized execution
	// engine for.
	VectorizeMode VectorizeExecMode `protobuf:"varint,5,opt,name=vectorize_mode,json=vectorizeMode,proto3,enum=cockroach.sql.sessiondatapb.VectorizeExecMode" json:"vectorize_mode,omitempty"`
	// TestingVectorizeInjectPanics indicates whether random panics are injected
	// into the vectorized flow execution. The goal of such behavior is making
	// sure that errors that are propagated as panics in the vectorized engine
	// are caught in all scenarios.
	TestingVectorizeInjectPanics bool `protobuf:"varint,6,opt,name=testing_vectorize_inject_panics,json=testingVectorizeInjectPanics,proto3" json:"testing_vectorize_inject_panics,omitempty"`
	// DefaultIntSize specifies the size in bits or bytes (preferred) of how a
	// "naked" INT type should be parsed.
	DefaultIntSize int32 `protobuf:"varint,7,opt,name=default_int_size,json=defaultIntSize,proto3" json:"default_int_size,omitempty"`
	// The name of the location according to whose current timezone we're going to
	// parse timestamps. Used to initialize sessiondata.SessionData.Location on
	// the remote nodes.
	//
	// Note that the current serialization of the time.Location objects as
	// strings has many drawbacks which could lead to unstable computation on the
	// remote nodes. See #36864 and
	// https://github.com/cockroachdb/cockroach/pull/55377#issuecomment-707794695
	// for more details.
	Location string `protobuf:"bytes,8,opt,name=location,proto3" json:"location,omitempty"`
	// SearchPath is a list of namespaces to search builtins in. Used to
	// initialize sessiondata.SessionData.SearchPath on the remote nodes.
	SearchPath          []string `protobuf:"bytes,9,rep,name=search_path,json=searchPath,proto3" json:"search_path,omitempty"`
	TemporarySchemaName string   `protobuf:"bytes,10,opt,name=temporary_schema_name,json=temporarySchemaName,proto3" json:"temporary_schema_name,omitempty"`
	// SeqState gives access to the SQL sequences that have been manipulated by
	// the session.
	SeqState SequenceState `protobuf:"bytes,11,opt,name=seq_state,json=seqState,proto3" json:"seq_state"`
}

func (m *SessionData) Reset()         { *m = SessionData{} }
func (m *SessionData) String() string { return proto.CompactTextString(m) }
func (*SessionData) ProtoMessage()    {}
func (*SessionData) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_data_c24dfdb3837a3a66, []int{0}
}
func (m *SessionData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SessionData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *SessionData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionData.Merge(dst, src)
}
func (m *SessionData) XXX_Size() int {
	return m.Size()
}
func (m *SessionData) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionData.DiscardUnknown(m)
}

var xxx_messageInfo_SessionData proto.InternalMessageInfo

// DataConversionConfig contains the parameters that influence the conversion
// between SQL data types and strings/byte arrays.
type DataConversionConfig struct {
	// BytesEncodeFormat indicates how to encode byte arrays when converting to
	// string.
	BytesEncodeFormat BytesEncodeFormat `protobuf:"varint,1,opt,name=bytes_encode_format,json=bytesEncodeFormat,proto3,enum=cockroach.sql.sessiondatapb.BytesEncodeFormat" json:"bytes_encode_format,omitempty"`
	// ExtraFloatDigits indicates the number of digits beyond the standard number
	// to use for float conversions.This must be set to a value between -15 and
	// 3, inclusive.
	ExtraFloatDigits int32 `protobuf:"varint,2,opt,name=extra_float_digits,json=extraFloatDigits,proto3" json:"extra_float_digits,omitempty"`
}

func (m *DataConversionConfig) Reset()         { *m = DataConversionConfig{} }
func (m *DataConversionConfig) String() string { return proto.CompactTextString(m) }
func (*DataConversionConfig) ProtoMessage()    {}
func (*DataConversionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_data_c24dfdb3837a3a66, []int{1}
}
func (m *DataConversionConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataConversionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *DataConversionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataConversionConfig.Merge(dst, src)
}
func (m *DataConversionConfig) XXX_Size() int {
	return m.Size()
}
func (m *DataConversionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DataConversionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DataConversionConfig proto.InternalMessageInfo

// SequenceState is used to marshall the sessiondata.SequenceState struct.
type SequenceState struct {
	Seqs []*SequenceState_Seq `protobuf:"bytes,1,rep,name=seqs,proto3" json:"seqs,omitempty"`
	// last_seq_incremented is the id of the last sequence incremented by the
	// session. This field is filled in iff seqs is not empty.
	LastSeqIncremented uint32 `protobuf:"varint,2,opt,name=last_seq_incremented,json=lastSeqIncremented,proto3" json:"last_seq_incremented,omitempty"`
}

func (m *SequenceState) Reset()         { *m = SequenceState{} }
func (m *SequenceState) String() string { return proto.CompactTextString(m) }
func (*SequenceState) ProtoMessage()    {}
func (*SequenceState) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_data_c24dfdb3837a3a66, []int{2}
}
func (m *SequenceState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SequenceState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *SequenceState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SequenceState.Merge(dst, src)
}
func (m *SequenceState) XXX_Size() int {
	return m.Size()
}
func (m *SequenceState) XXX_DiscardUnknown() {
	xxx_messageInfo_SequenceState.DiscardUnknown(m)
}

var xxx_messageInfo_SequenceState proto.InternalMessageInfo

// Seq represents the last value of one sequence modified by the session.
type SequenceState_Seq struct {
	SeqID     uint32 `protobuf:"varint,1,opt,name=seq_id,json=seqId,proto3" json:"seq_id,omitempty"`
	LatestVal int64  `protobuf:"varint,2,opt,name=latest_val,json=latestVal,proto3" json:"latest_val,omitempty"`
}

func (m *SequenceState_Seq) Reset()         { *m = SequenceState_Seq{} }
func (m *SequenceState_Seq) String() string { return proto.CompactTextString(m) }
func (*SequenceState_Seq) ProtoMessage()    {}
func (*SequenceState_Seq) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_data_c24dfdb3837a3a66, []int{2, 0}
}
func (m *SequenceState_Seq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SequenceState_Seq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *SequenceState_Seq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SequenceState_Seq.Merge(dst, src)
}
func (m *SequenceState_Seq) XXX_Size() int {
	return m.Size()
}
func (m *SequenceState_Seq) XXX_DiscardUnknown() {
	xxx_messageInfo_SequenceState_Seq.DiscardUnknown(m)
}

var xxx_messageInfo_SequenceState_Seq proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SessionData)(nil), "cockroach.sql.sessiondatapb.SessionData")
	proto.RegisterType((*DataConversionConfig)(nil), "cockroach.sql.sessiondatapb.DataConversionConfig")
	proto.RegisterType((*SequenceState)(nil), "cockroach.sql.sessiondatapb.SequenceState")
	proto.RegisterType((*SequenceState_Seq)(nil), "cockroach.sql.sessiondatapb.SequenceState.Seq")
	proto.RegisterEnum("cockroach.sql.sessiondatapb.BytesEncodeFormat", BytesEncodeFormat_name, BytesEncodeFormat_value)
	proto.RegisterEnum("cockroach.sql.sessiondatapb.VectorizeExecMode", VectorizeExecMode_name, VectorizeExecMode_value)
}
func (m *SessionData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SessionData) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Database) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSessionData(dAtA, i, uint64(len(m.Database)))
		i += copy(dAtA[i:], m.Database)
	}
	if len(m.ApplicationName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSessionData(dAtA, i, uint64(len(m.ApplicationName)))
		i += copy(dAtA[i:], m.ApplicationName)
	}
	if len(m.User) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSessionData(dAtA, i, uint64(len(m.User)))
		i += copy(dAtA[i:], m.User)
	}
	dAtA[i] = 0x22
	i++
	i = encodeVarintSessionData(dAtA, i, uint64(m.DataConversionConfig.Size()))
	n1, err := m.DataConversionConfig.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.VectorizeMode != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintSessionData(dAtA, i, uint64(m.VectorizeMode))
	}
	if m.TestingVectorizeInjectPanics {
		dAtA[i] = 0x30
		i++
		if m.TestingVectorizeInjectPanics {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.DefaultIntSize != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintSessionData(dAtA, i, uint64(m.DefaultIntSize))
	}
	if len(m.Location) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintSessionData(dAtA, i, uint64(len(m.Location)))
		i += copy(dAtA[i:], m.Location)
	}
	if len(m.SearchPath) > 0 {
		for _, s := range m.SearchPath {
			dAtA[i] = 0x4a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.TemporarySchemaName) > 0 {
		dAtA[i] = 0x52
		i++
		i = encodeVarintSessionData(dAtA, i, uint64(len(m.TemporarySchemaName)))
		i += copy(dAtA[i:], m.TemporarySchemaName)
	}
	dAtA[i] = 0x5a
	i++
	i = encodeVarintSessionData(dAtA, i, uint64(m.SeqState.Size()))
	n2, err := m.SeqState.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	return i, nil
}

func (m *DataConversionConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataConversionConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.BytesEncodeFormat != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSessionData(dAtA, i, uint64(m.BytesEncodeFormat))
	}
	if m.ExtraFloatDigits != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSessionData(dAtA, i, uint64(m.ExtraFloatDigits))
	}
	return i, nil
}

func (m *SequenceState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SequenceState) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Seqs) > 0 {
		for _, msg := range m.Seqs {
			dAtA[i] = 0xa
			i++
			i = encodeVarintSessionData(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.LastSeqIncremented != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSessionData(dAtA, i, uint64(m.LastSeqIncremented))
	}
	return i, nil
}

func (m *SequenceState_Seq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SequenceState_Seq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.SeqID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSessionData(dAtA, i, uint64(m.SeqID))
	}
	if m.LatestVal != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSessionData(dAtA, i, uint64(m.LatestVal))
	}
	return i, nil
}

func encodeVarintSessionData(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SessionData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Database)
	if l > 0 {
		n += 1 + l + sovSessionData(uint64(l))
	}
	l = len(m.ApplicationName)
	if l > 0 {
		n += 1 + l + sovSessionData(uint64(l))
	}
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovSessionData(uint64(l))
	}
	l = m.DataConversionConfig.Size()
	n += 1 + l + sovSessionData(uint64(l))
	if m.VectorizeMode != 0 {
		n += 1 + sovSessionData(uint64(m.VectorizeMode))
	}
	if m.TestingVectorizeInjectPanics {
		n += 2
	}
	if m.DefaultIntSize != 0 {
		n += 1 + sovSessionData(uint64(m.DefaultIntSize))
	}
	l = len(m.Location)
	if l > 0 {
		n += 1 + l + sovSessionData(uint64(l))
	}
	if len(m.SearchPath) > 0 {
		for _, s := range m.SearchPath {
			l = len(s)
			n += 1 + l + sovSessionData(uint64(l))
		}
	}
	l = len(m.TemporarySchemaName)
	if l > 0 {
		n += 1 + l + sovSessionData(uint64(l))
	}
	l = m.SeqState.Size()
	n += 1 + l + sovSessionData(uint64(l))
	return n
}

func (m *DataConversionConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BytesEncodeFormat != 0 {
		n += 1 + sovSessionData(uint64(m.BytesEncodeFormat))
	}
	if m.ExtraFloatDigits != 0 {
		n += 1 + sovSessionData(uint64(m.ExtraFloatDigits))
	}
	return n
}

func (m *SequenceState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Seqs) > 0 {
		for _, e := range m.Seqs {
			l = e.Size()
			n += 1 + l + sovSessionData(uint64(l))
		}
	}
	if m.LastSeqIncremented != 0 {
		n += 1 + sovSessionData(uint64(m.LastSeqIncremented))
	}
	return n
}

func (m *SequenceState_Seq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SeqID != 0 {
		n += 1 + sovSessionData(uint64(m.SeqID))
	}
	if m.LatestVal != 0 {
		n += 1 + sovSessionData(uint64(m.LatestVal))
	}
	return n
}

func sovSessionData(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSessionData(x uint64) (n int) {
	return sovSessionData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SessionData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSessionData
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
			return fmt.Errorf("proto: SessionData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SessionData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Database", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessionData
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
				return ErrInvalidLengthSessionData
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Database = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessionData
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
				return ErrInvalidLengthSessionData
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApplicationName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessionData
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
				return ErrInvalidLengthSessionData
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataConversionConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessionData
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
				return ErrInvalidLengthSessionData
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DataConversionConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VectorizeMode", wireType)
			}
			m.VectorizeMode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessionData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VectorizeMode |= (VectorizeExecMode(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TestingVectorizeInjectPanics", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessionData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.TestingVectorizeInjectPanics = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultIntSize", wireType)
			}
			m.DefaultIntSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessionData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DefaultIntSize |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Location", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessionData
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
				return ErrInvalidLengthSessionData
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Location = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SearchPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessionData
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
				return ErrInvalidLengthSessionData
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SearchPath = append(m.SearchPath, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TemporarySchemaName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessionData
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
				return ErrInvalidLengthSessionData
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TemporarySchemaName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SeqState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessionData
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
				return ErrInvalidLengthSessionData
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SeqState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSessionData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSessionData
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
func (m *DataConversionConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSessionData
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
			return fmt.Errorf("proto: DataConversionConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataConversionConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BytesEncodeFormat", wireType)
			}
			m.BytesEncodeFormat = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessionData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BytesEncodeFormat |= (BytesEncodeFormat(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtraFloatDigits", wireType)
			}
			m.ExtraFloatDigits = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessionData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExtraFloatDigits |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSessionData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSessionData
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
func (m *SequenceState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSessionData
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
			return fmt.Errorf("proto: SequenceState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SequenceState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seqs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessionData
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
				return ErrInvalidLengthSessionData
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Seqs = append(m.Seqs, &SequenceState_Seq{})
			if err := m.Seqs[len(m.Seqs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastSeqIncremented", wireType)
			}
			m.LastSeqIncremented = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessionData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastSeqIncremented |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSessionData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSessionData
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
func (m *SequenceState_Seq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSessionData
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
			return fmt.Errorf("proto: Seq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Seq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SeqID", wireType)
			}
			m.SeqID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessionData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SeqID |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestVal", wireType)
			}
			m.LatestVal = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSessionData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LatestVal |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSessionData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSessionData
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
func skipSessionData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSessionData
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
					return 0, ErrIntOverflowSessionData
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
					return 0, ErrIntOverflowSessionData
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
				return 0, ErrInvalidLengthSessionData
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSessionData
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
				next, err := skipSessionData(dAtA[start:])
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
	ErrInvalidLengthSessionData = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSessionData   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("sql/sessiondatapb/session_data.proto", fileDescriptor_session_data_c24dfdb3837a3a66)
}

var fileDescriptor_session_data_c24dfdb3837a3a66 = []byte{
	// 738 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xcd, 0x6e, 0xea, 0x46,
	0x14, 0xb6, 0x03, 0xa4, 0x30, 0x14, 0xae, 0x33, 0x97, 0x5b, 0x59, 0xdc, 0xd6, 0x58, 0x51, 0x17,
	0x6e, 0x54, 0x91, 0x7b, 0x69, 0xd5, 0x7d, 0x48, 0x88, 0xca, 0x22, 0x6d, 0x64, 0xd4, 0x2c, 0xba,
	0xe8, 0x68, 0x18, 0x1f, 0xc0, 0xad, 0xed, 0xb1, 0x3d, 0x43, 0xfe, 0x9e, 0xa0, 0xcb, 0xbe, 0x43,
	0xb3, 0xe8, 0xa3, 0x64, 0x53, 0x29, 0xcb, 0xac, 0xd2, 0x96, 0xbc, 0x48, 0x35, 0x43, 0x02, 0x21,
	0x89, 0x92, 0xdd, 0x9c, 0xef, 0xe7, 0xcc, 0xe1, 0x7c, 0x83, 0xd1, 0x97, 0x22, 0x8b, 0xb6, 0x05,
	0x08, 0x11, 0xf2, 0x24, 0xa0, 0x92, 0xa6, 0xc3, 0xfb, 0x8a, 0xa8, 0xb2, 0x9d, 0xe6, 0x5c, 0x72,
	0xfc, 0x9e, 0x71, 0xf6, 0x5b, 0xce, 0x29, 0x9b, 0xb4, 0x45, 0x16, 0xb5, 0x57, 0xf4, 0xcd, 0xc6,
	0x98, 0x8f, 0xb9, 0xd6, 0x6d, 0xab, 0xd3, 0xdc, 0xb2, 0xf9, 0x4f, 0x11, 0x55, 0x07, 0x73, 0xdd,
	0x1e, 0x95, 0x14, 0x37, 0x51, 0x59, 0xe9, 0x87, 0x54, 0x80, 0x6d, 0xba, 0xa6, 0x57, 0xf1, 0x17,
	0x35, 0xfe, 0x0a, 0x59, 0x34, 0x4d, 0xa3, 0x90, 0x51, 0xa9, 0x2e, 0x4e, 0x68, 0x0c, 0xf6, 0x9a,
	0xd6, 0xbc, 0x79, 0x80, 0xff, 0x40, 0x63, 0xc0, 0x18, 0x15, 0xa7, 0x02, 0x72, 0xbb, 0xa0, 0x69,
	0x7d, 0xc6, 0x31, 0xfa, 0x4c, 0xb5, 0x22, 0x8c, 0x27, 0xc7, 0x90, 0xeb, 0xd9, 0x19, 0x4f, 0x46,
	0xe1, 0xd8, 0x2e, 0xba, 0xa6, 0x57, 0xed, 0x7c, 0x6c, 0xbf, 0x30, 0x7e, 0x5b, 0x4d, 0xb7, 0xbb,
	0x70, 0xee, 0x6a, 0x63, 0xb7, 0x78, 0x79, 0xd3, 0x32, 0xfc, 0x46, 0xf0, 0x0c, 0x87, 0x7f, 0x42,
	0xf5, 0x63, 0x60, 0x92, 0xe7, 0xe1, 0x39, 0x90, 0x98, 0x07, 0x60, 0x97, 0x5c, 0xd3, 0xab, 0x77,
	0xda, 0x2f, 0x5e, 0x73, 0x74, 0x6f, 0xe9, 0x9d, 0x02, 0x3b, 0xe0, 0x01, 0xf8, 0xb5, 0x45, 0x17,
	0x55, 0xe2, 0x1e, 0x6a, 0x49, 0x10, 0x32, 0x4c, 0xc6, 0x64, 0xd9, 0x3e, 0x4c, 0x7e, 0x05, 0x26,
	0x49, 0x4a, 0x93, 0x90, 0x09, 0x7b, 0xdd, 0x35, 0xbd, 0xb2, 0xff, 0xf9, 0x9d, 0x6c, 0xd1, 0xb1,
	0xaf, 0x45, 0x87, 0x5a, 0x83, 0x3d, 0x64, 0x05, 0x30, 0xa2, 0xd3, 0x48, 0x92, 0x30, 0x91, 0x44,
	0x84, 0xe7, 0x60, 0x7f, 0xe2, 0x9a, 0x5e, 0xc9, 0xaf, 0xdf, 0xe1, 0xfd, 0x44, 0x0e, 0xc2, 0x73,
	0x50, 0x89, 0x44, 0x7c, 0xbe, 0x5a, 0xbb, 0x3c, 0x4f, 0xe4, 0xbe, 0xc6, 0x2d, 0x54, 0x15, 0x40,
	0x73, 0x36, 0x21, 0x29, 0x95, 0x13, 0xbb, 0xe2, 0x16, 0xbc, 0x8a, 0x8f, 0xe6, 0xd0, 0x21, 0x95,
	0x13, 0xdc, 0x41, 0xef, 0x24, 0xc4, 0x29, 0xcf, 0x69, 0x7e, 0x46, 0x04, 0x9b, 0x40, 0x4c, 0xe7,
	0xb9, 0x21, 0xdd, 0xe9, 0xed, 0x82, 0x1c, 0x68, 0x4e, 0x67, 0x77, 0x80, 0x2a, 0x02, 0x32, 0x22,
	0x24, 0x95, 0x60, 0x57, 0x75, 0x34, 0x5b, 0x2f, 0xee, 0x6c, 0x00, 0xd9, 0x14, 0x12, 0x06, 0x03,
	0xe5, 0xb8, 0xcb, 0xa4, 0x2c, 0x20, 0xd3, 0xf5, 0xe6, 0x85, 0x89, 0x1a, 0xcf, 0x85, 0x87, 0x7f,
	0x41, 0x6f, 0x87, 0x67, 0x12, 0x04, 0x81, 0x84, 0xf1, 0x00, 0xc8, 0x88, 0xe7, 0x31, 0x95, 0xfa,
	0xd5, 0xbd, 0x96, 0x52, 0x57, 0xf9, 0x7a, 0xda, 0xb6, 0xaf, 0x5d, 0xfe, 0xc6, 0xf0, 0x31, 0x84,
	0xbf, 0x46, 0x18, 0x4e, 0x65, 0x4e, 0xc9, 0x28, 0xe2, 0x54, 0x92, 0x20, 0x1c, 0x87, 0x52, 0xe8,
	0x07, 0x5b, 0xf2, 0x2d, 0xcd, 0xec, 0x2b, 0x62, 0x4f, 0xe3, 0x9b, 0x7f, 0x9b, 0xa8, 0xb6, 0xf2,
	0x43, 0x70, 0x17, 0x15, 0x05, 0x64, 0xc2, 0x36, 0xdd, 0x82, 0x57, 0x7d, 0x65, 0xa0, 0x15, 0xa7,
	0xaa, 0x7c, 0xed, 0xc5, 0x1f, 0x50, 0x23, 0xa2, 0x42, 0x12, 0xb5, 0xd0, 0x30, 0x61, 0x39, 0xc4,
	0x90, 0x48, 0x08, 0xf4, 0x14, 0x35, 0x1f, 0x2b, 0x6e, 0x00, 0x59, 0x7f, 0xc9, 0x34, 0xf7, 0x51,
	0x61, 0x00, 0x19, 0x76, 0xd1, 0xba, 0xf6, 0x04, 0x7a, 0x1f, 0xb5, 0x6e, 0x65, 0x76, 0xd3, 0x2a,
	0x29, 0xe9, 0x9e, 0x5f, 0x12, 0x90, 0xf5, 0x03, 0xfc, 0x05, 0x42, 0x11, 0x55, 0x6f, 0x8c, 0x1c,
	0xd3, 0x48, 0x37, 0x2c, 0xf8, 0x95, 0x39, 0x72, 0x44, 0xa3, 0x2d, 0x82, 0x36, 0x9e, 0x6c, 0x09,
	0x63, 0x54, 0x7f, 0x00, 0x7e, 0x0f, 0xa7, 0x96, 0x81, 0xdf, 0xad, 0x08, 0x7b, 0x82, 0xd1, 0x14,
	0x2c, 0xf3, 0x11, 0xdc, 0xa5, 0x02, 0xbe, 0xfb, 0xd6, 0x5a, 0x6b, 0x96, 0x7f, 0xff, 0xd3, 0x31,
	0xfe, 0xba, 0x70, 0x8c, 0xad, 0x13, 0xb4, 0xf1, 0xe4, 0xcf, 0x82, 0x2d, 0xf4, 0xe9, 0x02, 0xfc,
	0x71, 0x34, 0xb2, 0x0c, 0xdc, 0x40, 0xd6, 0x02, 0xe9, 0x7c, 0xf8, 0xb8, 0x33, 0x95, 0xdc, 0x32,
	0xf1, 0x1b, 0x54, 0x5d, 0xea, 0x12, 0x6b, 0x0d, 0xb7, 0xd0, 0xfb, 0x07, 0xdd, 0x52, 0xc8, 0x43,
	0xb5, 0x0f, 0x1a, 0xed, 0x44, 0x27, 0xf4, 0x4c, 0x58, 0x85, 0xe5, 0xc5, 0xdd, 0xed, 0xcb, 0xff,
	0x1c, 0xe3, 0x72, 0xe6, 0x98, 0x57, 0x33, 0xc7, 0xbc, 0x9e, 0x39, 0xe6, 0xbf, 0x33, 0xc7, 0xfc,
	0xe3, 0xd6, 0x31, 0xae, 0x6e, 0x1d, 0xe3, 0xfa, 0xd6, 0x31, 0x7e, 0xae, 0xad, 0x84, 0x33, 0x5c,
	0xd7, 0x9f, 0xba, 0x6f, 0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x35, 0xd6, 0x95, 0x08, 0x45, 0x05,
	0x00, 0x00,
}