// Code generated by protoc-gen-gogo.
// source: volume_mount.proto
// DO NOT EDIT!

package models

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strconv "strconv"

import bytes "bytes"

import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DeprecatedBindMountMode int32

const (
	DeprecatedBindMountMode_RO DeprecatedBindMountMode = 0
	DeprecatedBindMountMode_RW DeprecatedBindMountMode = 1
)

var DeprecatedBindMountMode_name = map[int32]string{
	0: "RO",
	1: "RW",
}
var DeprecatedBindMountMode_value = map[string]int32{
	"RO": 0,
	"RW": 1,
}

func (x DeprecatedBindMountMode) Enum() *DeprecatedBindMountMode {
	p := new(DeprecatedBindMountMode)
	*p = x
	return p
}
func (x DeprecatedBindMountMode) MarshalJSON() ([]byte, error) {
	return proto.MarshalJSONEnum(DeprecatedBindMountMode_name, int32(x))
}
func (x *DeprecatedBindMountMode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DeprecatedBindMountMode_value, data, "DeprecatedBindMountMode")
	if err != nil {
		return err
	}
	*x = DeprecatedBindMountMode(value)
	return nil
}
func (DeprecatedBindMountMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorVolumeMount, []int{0}
}

type SharedDevice struct {
	VolumeId    string `protobuf:"bytes,1,opt,name=volume_id,json=volumeId" json:"volume_id"`
	MountConfig string `protobuf:"bytes,2,opt,name=mount_config,json=mountConfig" json:"mount_config"`
}

func (m *SharedDevice) Reset()                    { *m = SharedDevice{} }
func (*SharedDevice) ProtoMessage()               {}
func (*SharedDevice) Descriptor() ([]byte, []int) { return fileDescriptorVolumeMount, []int{0} }

func (m *SharedDevice) GetVolumeId() string {
	if m != nil {
		return m.VolumeId
	}
	return ""
}

func (m *SharedDevice) GetMountConfig() string {
	if m != nil {
		return m.MountConfig
	}
	return ""
}

type VolumeMount struct {
	DeprecatedVolumeId string                  `protobuf:"bytes,2,opt,name=deprecated_volume_id,json=deprecatedVolumeId" json:"deprecated_volume_id"`
	DeprecatedMode     DeprecatedBindMountMode `protobuf:"varint,4,opt,name=deprecated_mode,json=deprecatedMode,enum=models.DeprecatedBindMountMode" json:"deprecated_mode"`
	DeprecatedConfig   []byte                  `protobuf:"bytes,5,opt,name=deprecated_config,json=deprecatedConfig" json:"deprecated_config"`
	Driver             string                  `protobuf:"bytes,1,opt,name=driver" json:"driver"`
	ContainerDir       string                  `protobuf:"bytes,3,opt,name=container_dir,json=containerDir" json:"container_dir"`
	Mode               string                  `protobuf:"bytes,6,opt,name=mode" json:"mode"`
	// oneof device {
	Shared *SharedDevice `protobuf:"bytes,7,opt,name=shared" json:"shared"`
}

func (m *VolumeMount) Reset()                    { *m = VolumeMount{} }
func (*VolumeMount) ProtoMessage()               {}
func (*VolumeMount) Descriptor() ([]byte, []int) { return fileDescriptorVolumeMount, []int{1} }

func (m *VolumeMount) GetDeprecatedVolumeId() string {
	if m != nil {
		return m.DeprecatedVolumeId
	}
	return ""
}

func (m *VolumeMount) GetDeprecatedMode() DeprecatedBindMountMode {
	if m != nil {
		return m.DeprecatedMode
	}
	return DeprecatedBindMountMode_RO
}

func (m *VolumeMount) GetDeprecatedConfig() []byte {
	if m != nil {
		return m.DeprecatedConfig
	}
	return nil
}

func (m *VolumeMount) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *VolumeMount) GetContainerDir() string {
	if m != nil {
		return m.ContainerDir
	}
	return ""
}

func (m *VolumeMount) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

func (m *VolumeMount) GetShared() *SharedDevice {
	if m != nil {
		return m.Shared
	}
	return nil
}

type VolumePlacement struct {
	DriverNames []string `protobuf:"bytes,1,rep,name=driver_names,json=driverNames" json:"driver_names"`
}

func (m *VolumePlacement) Reset()                    { *m = VolumePlacement{} }
func (*VolumePlacement) ProtoMessage()               {}
func (*VolumePlacement) Descriptor() ([]byte, []int) { return fileDescriptorVolumeMount, []int{2} }

func (m *VolumePlacement) GetDriverNames() []string {
	if m != nil {
		return m.DriverNames
	}
	return nil
}

func init() {
	proto.RegisterType((*SharedDevice)(nil), "models.SharedDevice")
	proto.RegisterType((*VolumeMount)(nil), "models.VolumeMount")
	proto.RegisterType((*VolumePlacement)(nil), "models.VolumePlacement")
	proto.RegisterEnum("models.DeprecatedBindMountMode", DeprecatedBindMountMode_name, DeprecatedBindMountMode_value)
}
func (x DeprecatedBindMountMode) String() string {
	s, ok := DeprecatedBindMountMode_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *SharedDevice) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*SharedDevice)
	if !ok {
		that2, ok := that.(SharedDevice)
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
	if this.VolumeId != that1.VolumeId {
		return false
	}
	if this.MountConfig != that1.MountConfig {
		return false
	}
	return true
}
func (this *VolumeMount) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*VolumeMount)
	if !ok {
		that2, ok := that.(VolumeMount)
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
	if this.DeprecatedVolumeId != that1.DeprecatedVolumeId {
		return false
	}
	if this.DeprecatedMode != that1.DeprecatedMode {
		return false
	}
	if !bytes.Equal(this.DeprecatedConfig, that1.DeprecatedConfig) {
		return false
	}
	if this.Driver != that1.Driver {
		return false
	}
	if this.ContainerDir != that1.ContainerDir {
		return false
	}
	if this.Mode != that1.Mode {
		return false
	}
	if !this.Shared.Equal(that1.Shared) {
		return false
	}
	return true
}
func (this *VolumePlacement) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*VolumePlacement)
	if !ok {
		that2, ok := that.(VolumePlacement)
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
	if len(this.DriverNames) != len(that1.DriverNames) {
		return false
	}
	for i := range this.DriverNames {
		if this.DriverNames[i] != that1.DriverNames[i] {
			return false
		}
	}
	return true
}
func (this *SharedDevice) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&models.SharedDevice{")
	s = append(s, "VolumeId: "+fmt.Sprintf("%#v", this.VolumeId)+",\n")
	s = append(s, "MountConfig: "+fmt.Sprintf("%#v", this.MountConfig)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *VolumeMount) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&models.VolumeMount{")
	s = append(s, "DeprecatedVolumeId: "+fmt.Sprintf("%#v", this.DeprecatedVolumeId)+",\n")
	s = append(s, "DeprecatedMode: "+fmt.Sprintf("%#v", this.DeprecatedMode)+",\n")
	s = append(s, "DeprecatedConfig: "+fmt.Sprintf("%#v", this.DeprecatedConfig)+",\n")
	s = append(s, "Driver: "+fmt.Sprintf("%#v", this.Driver)+",\n")
	s = append(s, "ContainerDir: "+fmt.Sprintf("%#v", this.ContainerDir)+",\n")
	s = append(s, "Mode: "+fmt.Sprintf("%#v", this.Mode)+",\n")
	if this.Shared != nil {
		s = append(s, "Shared: "+fmt.Sprintf("%#v", this.Shared)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *VolumePlacement) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&models.VolumePlacement{")
	if this.DriverNames != nil {
		s = append(s, "DriverNames: "+fmt.Sprintf("%#v", this.DriverNames)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringVolumeMount(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringVolumeMount(m github_com_gogo_protobuf_proto.Message) string {
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
func (m *SharedDevice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SharedDevice) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintVolumeMount(dAtA, i, uint64(len(m.VolumeId)))
	i += copy(dAtA[i:], m.VolumeId)
	dAtA[i] = 0x12
	i++
	i = encodeVarintVolumeMount(dAtA, i, uint64(len(m.MountConfig)))
	i += copy(dAtA[i:], m.MountConfig)
	return i, nil
}

func (m *VolumeMount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VolumeMount) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintVolumeMount(dAtA, i, uint64(len(m.Driver)))
	i += copy(dAtA[i:], m.Driver)
	dAtA[i] = 0x12
	i++
	i = encodeVarintVolumeMount(dAtA, i, uint64(len(m.DeprecatedVolumeId)))
	i += copy(dAtA[i:], m.DeprecatedVolumeId)
	dAtA[i] = 0x1a
	i++
	i = encodeVarintVolumeMount(dAtA, i, uint64(len(m.ContainerDir)))
	i += copy(dAtA[i:], m.ContainerDir)
	dAtA[i] = 0x20
	i++
	i = encodeVarintVolumeMount(dAtA, i, uint64(m.DeprecatedMode))
	if m.DeprecatedConfig != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintVolumeMount(dAtA, i, uint64(len(m.DeprecatedConfig)))
		i += copy(dAtA[i:], m.DeprecatedConfig)
	}
	dAtA[i] = 0x32
	i++
	i = encodeVarintVolumeMount(dAtA, i, uint64(len(m.Mode)))
	i += copy(dAtA[i:], m.Mode)
	if m.Shared != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintVolumeMount(dAtA, i, uint64(m.Shared.Size()))
		n1, err := m.Shared.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *VolumePlacement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VolumePlacement) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.DriverNames) > 0 {
		for _, s := range m.DriverNames {
			dAtA[i] = 0xa
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
	return i, nil
}

func encodeFixed64VolumeMount(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32VolumeMount(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintVolumeMount(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SharedDevice) Size() (n int) {
	var l int
	_ = l
	l = len(m.VolumeId)
	n += 1 + l + sovVolumeMount(uint64(l))
	l = len(m.MountConfig)
	n += 1 + l + sovVolumeMount(uint64(l))
	return n
}

func (m *VolumeMount) Size() (n int) {
	var l int
	_ = l
	l = len(m.Driver)
	n += 1 + l + sovVolumeMount(uint64(l))
	l = len(m.DeprecatedVolumeId)
	n += 1 + l + sovVolumeMount(uint64(l))
	l = len(m.ContainerDir)
	n += 1 + l + sovVolumeMount(uint64(l))
	n += 1 + sovVolumeMount(uint64(m.DeprecatedMode))
	if m.DeprecatedConfig != nil {
		l = len(m.DeprecatedConfig)
		n += 1 + l + sovVolumeMount(uint64(l))
	}
	l = len(m.Mode)
	n += 1 + l + sovVolumeMount(uint64(l))
	if m.Shared != nil {
		l = m.Shared.Size()
		n += 1 + l + sovVolumeMount(uint64(l))
	}
	return n
}

func (m *VolumePlacement) Size() (n int) {
	var l int
	_ = l
	if len(m.DriverNames) > 0 {
		for _, s := range m.DriverNames {
			l = len(s)
			n += 1 + l + sovVolumeMount(uint64(l))
		}
	}
	return n
}

func sovVolumeMount(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozVolumeMount(x uint64) (n int) {
	return sovVolumeMount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *SharedDevice) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SharedDevice{`,
		`VolumeId:` + fmt.Sprintf("%v", this.VolumeId) + `,`,
		`MountConfig:` + fmt.Sprintf("%v", this.MountConfig) + `,`,
		`}`,
	}, "")
	return s
}
func (this *VolumeMount) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&VolumeMount{`,
		`Driver:` + fmt.Sprintf("%v", this.Driver) + `,`,
		`DeprecatedVolumeId:` + fmt.Sprintf("%v", this.DeprecatedVolumeId) + `,`,
		`ContainerDir:` + fmt.Sprintf("%v", this.ContainerDir) + `,`,
		`DeprecatedMode:` + fmt.Sprintf("%v", this.DeprecatedMode) + `,`,
		`DeprecatedConfig:` + fmt.Sprintf("%v", this.DeprecatedConfig) + `,`,
		`Mode:` + fmt.Sprintf("%v", this.Mode) + `,`,
		`Shared:` + strings.Replace(fmt.Sprintf("%v", this.Shared), "SharedDevice", "SharedDevice", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *VolumePlacement) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&VolumePlacement{`,
		`DriverNames:` + fmt.Sprintf("%v", this.DriverNames) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringVolumeMount(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *SharedDevice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVolumeMount
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
			return fmt.Errorf("proto: SharedDevice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SharedDevice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VolumeId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVolumeMount
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
				return ErrInvalidLengthVolumeMount
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VolumeId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MountConfig", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVolumeMount
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
				return ErrInvalidLengthVolumeMount
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MountConfig = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVolumeMount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVolumeMount
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
func (m *VolumeMount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVolumeMount
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
			return fmt.Errorf("proto: VolumeMount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VolumeMount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Driver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVolumeMount
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
				return ErrInvalidLengthVolumeMount
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Driver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeprecatedVolumeId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVolumeMount
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
				return ErrInvalidLengthVolumeMount
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeprecatedVolumeId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContainerDir", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVolumeMount
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
				return ErrInvalidLengthVolumeMount
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContainerDir = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeprecatedMode", wireType)
			}
			m.DeprecatedMode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVolumeMount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DeprecatedMode |= (DeprecatedBindMountMode(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeprecatedConfig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVolumeMount
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
				return ErrInvalidLengthVolumeMount
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeprecatedConfig = append(m.DeprecatedConfig[:0], dAtA[iNdEx:postIndex]...)
			if m.DeprecatedConfig == nil {
				m.DeprecatedConfig = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVolumeMount
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
				return ErrInvalidLengthVolumeMount
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Mode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shared", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVolumeMount
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
				return ErrInvalidLengthVolumeMount
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Shared == nil {
				m.Shared = &SharedDevice{}
			}
			if err := m.Shared.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVolumeMount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVolumeMount
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
func (m *VolumePlacement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVolumeMount
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
			return fmt.Errorf("proto: VolumePlacement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VolumePlacement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DriverNames", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVolumeMount
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
				return ErrInvalidLengthVolumeMount
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DriverNames = append(m.DriverNames, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVolumeMount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVolumeMount
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
func skipVolumeMount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVolumeMount
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
					return 0, ErrIntOverflowVolumeMount
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
					return 0, ErrIntOverflowVolumeMount
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
				return 0, ErrInvalidLengthVolumeMount
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowVolumeMount
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
				next, err := skipVolumeMount(dAtA[start:])
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
	ErrInvalidLengthVolumeMount = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVolumeMount   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("volume_mount.proto", fileDescriptorVolumeMount) }

var fileDescriptorVolumeMount = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x51, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xf5, 0x36, 0xc5, 0x90, 0x89, 0x49, 0xd3, 0x55, 0x10, 0x16, 0x87, 0x8d, 0x95, 0x03, 0x32,
	0x08, 0x5c, 0xa9, 0x3d, 0x50, 0x71, 0x34, 0x11, 0x12, 0x87, 0x42, 0xb5, 0x48, 0xe5, 0x68, 0xb9,
	0xde, 0x6d, 0xba, 0x52, 0xec, 0xad, 0x1c, 0x27, 0x5c, 0xf9, 0x02, 0xc4, 0x67, 0xf0, 0x29, 0x3d,
	0xf6, 0xc8, 0xa9, 0x22, 0xe6, 0x82, 0x38, 0xf5, 0x13, 0x90, 0x67, 0xb7, 0x71, 0x7a, 0xe8, 0xc5,
	0xb3, 0x6f, 0x66, 0xde, 0xcc, 0xf3, 0x1b, 0xa0, 0x4b, 0x3d, 0x5b, 0xe4, 0x32, 0xc9, 0xf5, 0xa2,
	0xa8, 0xa2, 0x8b, 0x52, 0x57, 0x9a, 0xba, 0xb9, 0x16, 0x72, 0x36, 0x7f, 0xf6, 0x7a, 0xaa, 0xaa,
	0xf3, 0xc5, 0x69, 0x94, 0xe9, 0x7c, 0x6f, 0xaa, 0xa7, 0x7a, 0x0f, 0xcb, 0xa7, 0x8b, 0x33, 0x44,
	0x08, 0xf0, 0x65, 0x68, 0xe3, 0xaf, 0xe0, 0x7d, 0x3e, 0x4f, 0x4b, 0x29, 0x26, 0x72, 0xa9, 0x32,
	0x49, 0x23, 0xe8, 0xda, 0xe1, 0x4a, 0xf8, 0x24, 0x20, 0x61, 0x37, 0xde, 0xbd, 0xbc, 0x1e, 0x39,
	0xff, 0xae, 0x47, 0x6d, 0x81, 0x3f, 0x32, 0xcf, 0x0f, 0x82, 0xbe, 0x01, 0x0f, 0x55, 0x24, 0x99,
	0x2e, 0xce, 0xd4, 0xd4, 0xdf, 0x42, 0xca, 0xd0, 0x52, 0xee, 0xd4, 0x78, 0x0f, 0xd1, 0x3b, 0x04,
	0xe3, 0xef, 0x1d, 0xe8, 0x9d, 0xe0, 0x94, 0xa3, 0x26, 0x4b, 0x9f, 0x83, 0x2b, 0x4a, 0xb5, 0x94,
	0xa5, 0xdd, 0xda, 0xb7, 0x23, 0x6c, 0x96, 0xdb, 0x48, 0x0f, 0x61, 0x28, 0xe4, 0x45, 0x29, 0xb3,
	0xb4, 0x92, 0x22, 0x69, 0xb5, 0x9a, 0xc5, 0x6e, 0xc3, 0xf2, 0x09, 0xa7, 0x6d, 0xcf, 0xc9, 0xad,
	0xd4, 0xb7, 0xf0, 0x38, 0xd3, 0x45, 0x95, 0xaa, 0x42, 0x96, 0x89, 0x50, 0xa5, 0xdf, 0x41, 0xca,
	0x13, 0xbb, 0xe8, 0x6e, 0x91, 0x7b, 0x6b, 0x38, 0x51, 0x25, 0x3d, 0x86, 0x9d, 0x8d, 0xad, 0x8d,
	0xd5, 0xfe, 0x76, 0x40, 0xc2, 0xfe, 0xfe, 0x28, 0x32, 0xbe, 0x47, 0x93, 0x75, 0x39, 0x56, 0x85,
	0xc0, 0x7f, 0x3a, 0xd2, 0x42, 0xae, 0x15, 0xf5, 0x5b, 0x7e, 0x93, 0xa7, 0x07, 0xb0, 0xbb, 0x31,
	0xd1, 0xba, 0xf7, 0x20, 0x20, 0xa1, 0xb7, 0xa6, 0x0c, 0xda, 0x06, 0x63, 0x1a, 0x0d, 0x60, 0x1b,
	0x77, 0xbb, 0xa8, 0xdc, 0xb3, 0xca, 0x31, 0xc7, 0xf1, 0x4b, 0x0f, 0xc1, 0x9d, 0xe3, 0x3d, 0xfd,
	0x87, 0x01, 0x09, 0x7b, 0xfb, 0xc3, 0x5b, 0x7d, 0x9b, 0x57, 0x8e, 0xa1, 0x31, 0xd6, 0xf4, 0x71,
	0x1b, 0xc7, 0xef, 0x61, 0xc7, 0x58, 0x75, 0x3c, 0x4b, 0x33, 0x99, 0xcb, 0xa2, 0xa2, 0x07, 0xe0,
	0x19, 0xd7, 0x93, 0x22, 0xcd, 0xe5, 0xdc, 0x27, 0x41, 0x27, 0xec, 0xc6, 0x83, 0xe6, 0xb0, 0x9b,
	0x79, 0xde, 0x33, 0xe8, 0x63, 0x03, 0x5e, 0xbe, 0x80, 0xa7, 0xf7, 0x78, 0x41, 0x5d, 0xd8, 0xe2,
	0x9f, 0x06, 0x0e, 0xc6, 0x2f, 0x03, 0x12, 0xbf, 0xba, 0x5a, 0x31, 0xf2, 0x6b, 0xc5, 0x9c, 0x9b,
	0x15, 0x23, 0xdf, 0x6a, 0x46, 0x7e, 0xd6, 0x8c, 0x5c, 0xd6, 0x8c, 0x5c, 0xd5, 0x8c, 0xfc, 0xae,
	0x19, 0xf9, 0x5b, 0x33, 0xe7, 0xa6, 0x66, 0xe4, 0xc7, 0x1f, 0xe6, 0xfc, 0x0f, 0x00, 0x00, 0xff,
	0xff, 0xaa, 0xa4, 0x9d, 0x40, 0xf6, 0x02, 0x00, 0x00,
}
