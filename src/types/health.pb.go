// Code generated by protoc-gen-gogo.
// source: health.proto
// DO NOT EDIT!

package types

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// skipping weak import gogoproto "gogoproto"

import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HealthCheck struct {
	ID                  string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address             string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	TaskID              string   `protobuf:"bytes,3,opt,name=taskid,proto3" json:"taskid,omitempty"`
	AppID               string   `protobuf:"bytes,4,opt,name=appid,proto3" json:"appid,omitempty"`
	Protocol            string   `protobuf:"bytes,5,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Port                int32    `protobuf:"varint,6,opt,name=port,proto3" json:"port,omitempty"`
	PortIndex           int32    `protobuf:"varint,7,opt,name=portIndex,proto3" json:"portIndex,omitempty"`
	Command             *Command `protobuf:"bytes,8,opt,name=command" json:"command,omitempty"`
	Path                string   `protobuf:"bytes,9,opt,name=path,proto3" json:"path,omitempty"`
	DelaySeconds        float64  `protobuf:"fixed64,10,opt,name=delaySeconds,proto3" json:"delaySeconds,omitempty"`
	ConsecutiveFailures uint32   `protobuf:"varint,11,opt,name=consecutiveFailures,proto3" json:"consecutiveFailures,omitempty"`
	GracePeriodSeconds  float64  `protobuf:"fixed64,12,opt,name=gracePeriodSeconds,proto3" json:"gracePeriodSeconds,omitempty"`
	IntervalSeconds     float64  `protobuf:"fixed64,13,opt,name=intervalSeconds,proto3" json:"intervalSeconds,omitempty"`
	TimeoutSeconds      float64  `protobuf:"fixed64,14,opt,name=timeoutSeconds,proto3" json:"timeoutSeconds,omitempty"`
}

func (m *HealthCheck) Reset()                    { *m = HealthCheck{} }
func (m *HealthCheck) String() string            { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()               {}
func (*HealthCheck) Descriptor() ([]byte, []int) { return fileDescriptorHealth, []int{0} }

type Command struct {
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptorHealth, []int{1} }

func init() {
	proto.RegisterType((*HealthCheck)(nil), "types.HealthCheck")
	proto.RegisterType((*Command)(nil), "types.Command")
}
func (this *HealthCheck) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*HealthCheck)
	if !ok {
		that2, ok := that.(HealthCheck)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *HealthCheck")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *HealthCheck but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *HealthCheck but is not nil && this == nil")
	}
	if this.ID != that1.ID {
		return fmt.Errorf("ID this(%v) Not Equal that(%v)", this.ID, that1.ID)
	}
	if this.Address != that1.Address {
		return fmt.Errorf("Address this(%v) Not Equal that(%v)", this.Address, that1.Address)
	}
	if this.TaskID != that1.TaskID {
		return fmt.Errorf("TaskID this(%v) Not Equal that(%v)", this.TaskID, that1.TaskID)
	}
	if this.AppID != that1.AppID {
		return fmt.Errorf("AppID this(%v) Not Equal that(%v)", this.AppID, that1.AppID)
	}
	if this.Protocol != that1.Protocol {
		return fmt.Errorf("Protocol this(%v) Not Equal that(%v)", this.Protocol, that1.Protocol)
	}
	if this.Port != that1.Port {
		return fmt.Errorf("Port this(%v) Not Equal that(%v)", this.Port, that1.Port)
	}
	if this.PortIndex != that1.PortIndex {
		return fmt.Errorf("PortIndex this(%v) Not Equal that(%v)", this.PortIndex, that1.PortIndex)
	}
	if !this.Command.Equal(that1.Command) {
		return fmt.Errorf("Command this(%v) Not Equal that(%v)", this.Command, that1.Command)
	}
	if this.Path != that1.Path {
		return fmt.Errorf("Path this(%v) Not Equal that(%v)", this.Path, that1.Path)
	}
	if this.DelaySeconds != that1.DelaySeconds {
		return fmt.Errorf("DelaySeconds this(%v) Not Equal that(%v)", this.DelaySeconds, that1.DelaySeconds)
	}
	if this.ConsecutiveFailures != that1.ConsecutiveFailures {
		return fmt.Errorf("ConsecutiveFailures this(%v) Not Equal that(%v)", this.ConsecutiveFailures, that1.ConsecutiveFailures)
	}
	if this.GracePeriodSeconds != that1.GracePeriodSeconds {
		return fmt.Errorf("GracePeriodSeconds this(%v) Not Equal that(%v)", this.GracePeriodSeconds, that1.GracePeriodSeconds)
	}
	if this.IntervalSeconds != that1.IntervalSeconds {
		return fmt.Errorf("IntervalSeconds this(%v) Not Equal that(%v)", this.IntervalSeconds, that1.IntervalSeconds)
	}
	if this.TimeoutSeconds != that1.TimeoutSeconds {
		return fmt.Errorf("TimeoutSeconds this(%v) Not Equal that(%v)", this.TimeoutSeconds, that1.TimeoutSeconds)
	}
	return nil
}
func (this *HealthCheck) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*HealthCheck)
	if !ok {
		that2, ok := that.(HealthCheck)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.ID != that1.ID {
		return false
	}
	if this.Address != that1.Address {
		return false
	}
	if this.TaskID != that1.TaskID {
		return false
	}
	if this.AppID != that1.AppID {
		return false
	}
	if this.Protocol != that1.Protocol {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	if this.PortIndex != that1.PortIndex {
		return false
	}
	if !this.Command.Equal(that1.Command) {
		return false
	}
	if this.Path != that1.Path {
		return false
	}
	if this.DelaySeconds != that1.DelaySeconds {
		return false
	}
	if this.ConsecutiveFailures != that1.ConsecutiveFailures {
		return false
	}
	if this.GracePeriodSeconds != that1.GracePeriodSeconds {
		return false
	}
	if this.IntervalSeconds != that1.IntervalSeconds {
		return false
	}
	if this.TimeoutSeconds != that1.TimeoutSeconds {
		return false
	}
	return true
}
func (this *Command) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Command)
	if !ok {
		that2, ok := that.(Command)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Command")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Command but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Command but is not nil && this == nil")
	}
	if this.Value != that1.Value {
		return fmt.Errorf("Value this(%v) Not Equal that(%v)", this.Value, that1.Value)
	}
	return nil
}
func (this *Command) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Command)
	if !ok {
		that2, ok := that.(Command)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	return true
}
func (this *HealthCheck) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 18)
	s = append(s, "&types.HealthCheck{")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	s = append(s, "Address: "+fmt.Sprintf("%#v", this.Address)+",\n")
	s = append(s, "TaskID: "+fmt.Sprintf("%#v", this.TaskID)+",\n")
	s = append(s, "AppID: "+fmt.Sprintf("%#v", this.AppID)+",\n")
	s = append(s, "Protocol: "+fmt.Sprintf("%#v", this.Protocol)+",\n")
	s = append(s, "Port: "+fmt.Sprintf("%#v", this.Port)+",\n")
	s = append(s, "PortIndex: "+fmt.Sprintf("%#v", this.PortIndex)+",\n")
	if this.Command != nil {
		s = append(s, "Command: "+fmt.Sprintf("%#v", this.Command)+",\n")
	}
	s = append(s, "Path: "+fmt.Sprintf("%#v", this.Path)+",\n")
	s = append(s, "DelaySeconds: "+fmt.Sprintf("%#v", this.DelaySeconds)+",\n")
	s = append(s, "ConsecutiveFailures: "+fmt.Sprintf("%#v", this.ConsecutiveFailures)+",\n")
	s = append(s, "GracePeriodSeconds: "+fmt.Sprintf("%#v", this.GracePeriodSeconds)+",\n")
	s = append(s, "IntervalSeconds: "+fmt.Sprintf("%#v", this.IntervalSeconds)+",\n")
	s = append(s, "TimeoutSeconds: "+fmt.Sprintf("%#v", this.TimeoutSeconds)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Command) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&types.Command{")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringHealth(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringHealth(m github_com_gogo_protobuf_proto.Message) string {
	e := github_com_gogo_protobuf_proto.GetUnsafeExtensionsMap(m)
	if e == nil {
		return "nil"
	}
	s := "proto.NewUnsafeXXX_InternalExtensions(map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "})"
	return s
}
func (m *HealthCheck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HealthCheck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHealth(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if len(m.Address) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHealth(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
	}
	if len(m.TaskID) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintHealth(dAtA, i, uint64(len(m.TaskID)))
		i += copy(dAtA[i:], m.TaskID)
	}
	if len(m.AppID) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintHealth(dAtA, i, uint64(len(m.AppID)))
		i += copy(dAtA[i:], m.AppID)
	}
	if len(m.Protocol) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintHealth(dAtA, i, uint64(len(m.Protocol)))
		i += copy(dAtA[i:], m.Protocol)
	}
	if m.Port != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintHealth(dAtA, i, uint64(m.Port))
	}
	if m.PortIndex != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintHealth(dAtA, i, uint64(m.PortIndex))
	}
	if m.Command != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintHealth(dAtA, i, uint64(m.Command.Size()))
		n1, err := m.Command.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Path) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintHealth(dAtA, i, uint64(len(m.Path)))
		i += copy(dAtA[i:], m.Path)
	}
	if m.DelaySeconds != 0 {
		dAtA[i] = 0x51
		i++
		i = encodeFixed64Health(dAtA, i, uint64(math.Float64bits(float64(m.DelaySeconds))))
	}
	if m.ConsecutiveFailures != 0 {
		dAtA[i] = 0x58
		i++
		i = encodeVarintHealth(dAtA, i, uint64(m.ConsecutiveFailures))
	}
	if m.GracePeriodSeconds != 0 {
		dAtA[i] = 0x61
		i++
		i = encodeFixed64Health(dAtA, i, uint64(math.Float64bits(float64(m.GracePeriodSeconds))))
	}
	if m.IntervalSeconds != 0 {
		dAtA[i] = 0x69
		i++
		i = encodeFixed64Health(dAtA, i, uint64(math.Float64bits(float64(m.IntervalSeconds))))
	}
	if m.TimeoutSeconds != 0 {
		dAtA[i] = 0x71
		i++
		i = encodeFixed64Health(dAtA, i, uint64(math.Float64bits(float64(m.TimeoutSeconds))))
	}
	return i, nil
}

func (m *Command) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Command) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHealth(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	return i, nil
}

func encodeFixed64Health(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Health(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintHealth(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedHealthCheck(r randyHealth, easy bool) *HealthCheck {
	this := &HealthCheck{}
	this.ID = string(randStringHealth(r))
	this.Address = string(randStringHealth(r))
	this.TaskID = string(randStringHealth(r))
	this.AppID = string(randStringHealth(r))
	this.Protocol = string(randStringHealth(r))
	this.Port = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Port *= -1
	}
	this.PortIndex = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.PortIndex *= -1
	}
	if r.Intn(10) != 0 {
		this.Command = NewPopulatedCommand(r, easy)
	}
	this.Path = string(randStringHealth(r))
	this.DelaySeconds = float64(r.Float64())
	if r.Intn(2) == 0 {
		this.DelaySeconds *= -1
	}
	this.ConsecutiveFailures = uint32(r.Uint32())
	this.GracePeriodSeconds = float64(r.Float64())
	if r.Intn(2) == 0 {
		this.GracePeriodSeconds *= -1
	}
	this.IntervalSeconds = float64(r.Float64())
	if r.Intn(2) == 0 {
		this.IntervalSeconds *= -1
	}
	this.TimeoutSeconds = float64(r.Float64())
	if r.Intn(2) == 0 {
		this.TimeoutSeconds *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedCommand(r randyHealth, easy bool) *Command {
	this := &Command{}
	this.Value = string(randStringHealth(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyHealth interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneHealth(r randyHealth) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringHealth(r randyHealth) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneHealth(r)
	}
	return string(tmps)
}
func randUnrecognizedHealth(r randyHealth, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldHealth(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldHealth(dAtA []byte, r randyHealth, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateHealth(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateHealth(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateHealth(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateHealth(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateHealth(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateHealth(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateHealth(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *HealthCheck) Size() (n int) {
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovHealth(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovHealth(uint64(l))
	}
	l = len(m.TaskID)
	if l > 0 {
		n += 1 + l + sovHealth(uint64(l))
	}
	l = len(m.AppID)
	if l > 0 {
		n += 1 + l + sovHealth(uint64(l))
	}
	l = len(m.Protocol)
	if l > 0 {
		n += 1 + l + sovHealth(uint64(l))
	}
	if m.Port != 0 {
		n += 1 + sovHealth(uint64(m.Port))
	}
	if m.PortIndex != 0 {
		n += 1 + sovHealth(uint64(m.PortIndex))
	}
	if m.Command != nil {
		l = m.Command.Size()
		n += 1 + l + sovHealth(uint64(l))
	}
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovHealth(uint64(l))
	}
	if m.DelaySeconds != 0 {
		n += 9
	}
	if m.ConsecutiveFailures != 0 {
		n += 1 + sovHealth(uint64(m.ConsecutiveFailures))
	}
	if m.GracePeriodSeconds != 0 {
		n += 9
	}
	if m.IntervalSeconds != 0 {
		n += 9
	}
	if m.TimeoutSeconds != 0 {
		n += 9
	}
	return n
}

func (m *Command) Size() (n int) {
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovHealth(uint64(l))
	}
	return n
}

func sovHealth(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHealth(x uint64) (n int) {
	return sovHealth(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HealthCheck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHealth
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
			return fmt.Errorf("proto: HealthCheck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HealthCheck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealth
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
				return ErrInvalidLengthHealth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealth
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
				return ErrInvalidLengthHealth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaskID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealth
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
				return ErrInvalidLengthHealth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TaskID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealth
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
				return ErrInvalidLengthHealth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Protocol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealth
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
				return ErrInvalidLengthHealth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Protocol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortIndex", wireType)
			}
			m.PortIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PortIndex |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Command", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealth
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
				return ErrInvalidLengthHealth
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Command == nil {
				m.Command = &Command{}
			}
			if err := m.Command.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealth
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
				return ErrInvalidLengthHealth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelaySeconds", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.DelaySeconds = float64(math.Float64frombits(v))
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsecutiveFailures", wireType)
			}
			m.ConsecutiveFailures = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConsecutiveFailures |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field GracePeriodSeconds", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.GracePeriodSeconds = float64(math.Float64frombits(v))
		case 13:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntervalSeconds", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.IntervalSeconds = float64(math.Float64frombits(v))
		case 14:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutSeconds", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.TimeoutSeconds = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipHealth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHealth
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
func (m *Command) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHealth
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
			return fmt.Errorf("proto: Command: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Command: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealth
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
				return ErrInvalidLengthHealth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHealth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHealth
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
func skipHealth(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHealth
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
					return 0, ErrIntOverflowHealth
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
					return 0, ErrIntOverflowHealth
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
				return 0, ErrInvalidLengthHealth
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHealth
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
				next, err := skipHealth(dAtA[start:])
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
	ErrInvalidLengthHealth = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHealth   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("health.proto", fileDescriptorHealth) }

var fileDescriptorHealth = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x51, 0xcd, 0xce, 0xd2, 0x40,
	0x14, 0xa5, 0x40, 0x5b, 0x7a, 0xf9, 0x31, 0x19, 0x89, 0x99, 0x10, 0x03, 0x86, 0x18, 0xc3, 0xaa,
	0x1a, 0x7d, 0x02, 0xc1, 0x18, 0xd9, 0x99, 0xea, 0xc2, 0xed, 0xd8, 0x99, 0xd0, 0x86, 0xd2, 0x69,
	0xda, 0x29, 0x91, 0x37, 0xf2, 0x11, 0x7c, 0x04, 0x96, 0x3e, 0x81, 0x01, 0x56, 0x2e, 0x5d, 0xba,
	0xfc, 0x2e, 0xb7, 0xf4, 0xfb, 0xf2, 0x11, 0x16, 0x27, 0xbd, 0xf7, 0x9c, 0x73, 0x4f, 0xe7, 0xce,
	0x40, 0x2f, 0x52, 0x22, 0x31, 0x91, 0x9f, 0xe5, 0xda, 0x68, 0x66, 0x9b, 0x5d, 0xa6, 0x8a, 0xd1,
	0x70, 0xa5, 0x57, 0x9a, 0x98, 0xd7, 0xe7, 0xaa, 0x12, 0xa7, 0x7f, 0x5b, 0xd0, 0xfd, 0x44, 0xee,
	0x45, 0xa4, 0xc2, 0x35, 0x7b, 0x06, 0xcd, 0x58, 0x72, 0xeb, 0x85, 0x35, 0xf3, 0xe6, 0xce, 0xe9,
	0xcf, 0xa4, 0xb9, 0xfc, 0x10, 0x20, 0xc3, 0x38, 0xb8, 0x42, 0xca, 0x5c, 0x15, 0x05, 0x6f, 0x9e,
	0xc5, 0xa0, 0x6e, 0xd9, 0x14, 0x1c, 0x23, 0x8a, 0x35, 0x4e, 0xb5, 0x68, 0x0a, 0x70, 0xca, 0xf9,
	0x8a, 0x0c, 0x4e, 0x5e, 0x14, 0x36, 0x01, 0x5b, 0x64, 0x19, 0x5a, 0xda, 0x64, 0xf1, 0xd0, 0x62,
	0xbf, 0xcf, 0x32, 0x74, 0x54, 0x3c, 0x1b, 0x41, 0x87, 0xce, 0x13, 0xea, 0x84, 0xdb, 0x94, 0x7f,
	0xdf, 0x33, 0x06, 0xed, 0x4c, 0xe7, 0x86, 0x3b, 0xc8, 0xdb, 0x01, 0xd5, 0xec, 0x39, 0x78, 0xe7,
	0xef, 0x32, 0x95, 0xea, 0x07, 0x77, 0x49, 0x78, 0x20, 0xd8, 0x0c, 0xdc, 0x50, 0x6f, 0x36, 0x22,
	0x95, 0xbc, 0x83, 0x5a, 0xf7, 0xed, 0xc0, 0xa7, 0x3b, 0xf0, 0x17, 0x15, 0x1b, 0xd4, 0x32, 0x65,
	0x0b, 0x13, 0x71, 0x8f, 0xfe, 0x49, 0x35, 0x2e, 0xd4, 0x93, 0x2a, 0x11, 0xbb, 0x2f, 0x2a, 0xd4,
	0xa9, 0x2c, 0x38, 0xa0, 0x66, 0x05, 0x8f, 0x38, 0xf6, 0x06, 0x9e, 0x62, 0x51, 0xa8, 0xb0, 0x34,
	0xf1, 0x56, 0x7d, 0x14, 0x71, 0x52, 0xe2, 0x65, 0xf0, 0x2e, 0x5a, 0xfb, 0xc1, 0x2d, 0x89, 0xf9,
	0xc0, 0x56, 0xb9, 0x08, 0xd5, 0x67, 0x95, 0xc7, 0x5a, 0xd6, 0xd9, 0x3d, 0xca, 0xbe, 0xa1, 0xe0,
	0x0e, 0x4f, 0xe2, 0xd4, 0xa8, 0x7c, 0x2b, 0x92, 0xda, 0xdc, 0x27, 0xf3, 0x35, 0xcd, 0x5e, 0xc1,
	0xc0, 0xc4, 0x1b, 0xa5, 0x4b, 0x53, 0x1b, 0x07, 0x64, 0xbc, 0x62, 0xa7, 0x13, 0x70, 0x2f, 0xfb,
	0xb3, 0x21, 0xd8, 0x18, 0x50, 0xaa, 0xea, 0xa1, 0x83, 0xaa, 0x99, 0xbf, 0xdc, 0x1f, 0xc7, 0x8d,
	0xc3, 0x71, 0x6c, 0xfd, 0x43, 0xfc, 0x47, 0xfc, 0x3c, 0x8d, 0xad, 0x5f, 0x88, 0x3d, 0xe2, 0x37,
	0xe2, 0x80, 0xf8, 0xd6, 0xf8, 0xee, 0xd0, 0xd3, 0xbc, 0xbb, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x0b,
	0xbc, 0xfe, 0x67, 0x67, 0x02, 0x00, 0x00,
}
