// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lxd/migration/migrate.proto

/*
Package migration is a generated protocol buffer package.

It is generated from these files:
	lxd/migration/migrate.proto

It has these top-level messages:
	IDMapType
	Config
	Device
	Snapshot
	MigrationHeader
	MigrationControl
	MigrationSync
	DumpStatsEntry
	RestoreStatsEntry
	StatsEntry
*/
package migration

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MigrationFSType int32

const (
	MigrationFSType_RSYNC MigrationFSType = 0
	MigrationFSType_BTRFS MigrationFSType = 1
	MigrationFSType_ZFS   MigrationFSType = 2
	MigrationFSType_RBD   MigrationFSType = 3
)

var MigrationFSType_name = map[int32]string{
	0: "RSYNC",
	1: "BTRFS",
	2: "ZFS",
	3: "RBD",
}
var MigrationFSType_value = map[string]int32{
	"RSYNC": 0,
	"BTRFS": 1,
	"ZFS":   2,
	"RBD":   3,
}

func (x MigrationFSType) Enum() *MigrationFSType {
	p := new(MigrationFSType)
	*p = x
	return p
}
func (x MigrationFSType) String() string {
	return proto.EnumName(MigrationFSType_name, int32(x))
}
func (x *MigrationFSType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MigrationFSType_value, data, "MigrationFSType")
	if err != nil {
		return err
	}
	*x = MigrationFSType(value)
	return nil
}
func (MigrationFSType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CRIUType int32

const (
	CRIUType_CRIU_RSYNC CRIUType = 0
	CRIUType_PHAUL      CRIUType = 1
	CRIUType_NONE       CRIUType = 2
)

var CRIUType_name = map[int32]string{
	0: "CRIU_RSYNC",
	1: "PHAUL",
	2: "NONE",
}
var CRIUType_value = map[string]int32{
	"CRIU_RSYNC": 0,
	"PHAUL":      1,
	"NONE":       2,
}

func (x CRIUType) Enum() *CRIUType {
	p := new(CRIUType)
	*p = x
	return p
}
func (x CRIUType) String() string {
	return proto.EnumName(CRIUType_name, int32(x))
}
func (x *CRIUType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CRIUType_value, data, "CRIUType")
	if err != nil {
		return err
	}
	*x = CRIUType(value)
	return nil
}
func (CRIUType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type IDMapType struct {
	Isuid            *bool  `protobuf:"varint,1,req,name=isuid" json:"isuid,omitempty"`
	Isgid            *bool  `protobuf:"varint,2,req,name=isgid" json:"isgid,omitempty"`
	Hostid           *int32 `protobuf:"varint,3,req,name=hostid" json:"hostid,omitempty"`
	Nsid             *int32 `protobuf:"varint,4,req,name=nsid" json:"nsid,omitempty"`
	Maprange         *int32 `protobuf:"varint,5,req,name=maprange" json:"maprange,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *IDMapType) Reset()                    { *m = IDMapType{} }
func (m *IDMapType) String() string            { return proto.CompactTextString(m) }
func (*IDMapType) ProtoMessage()               {}
func (*IDMapType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IDMapType) GetIsuid() bool {
	if m != nil && m.Isuid != nil {
		return *m.Isuid
	}
	return false
}

func (m *IDMapType) GetIsgid() bool {
	if m != nil && m.Isgid != nil {
		return *m.Isgid
	}
	return false
}

func (m *IDMapType) GetHostid() int32 {
	if m != nil && m.Hostid != nil {
		return *m.Hostid
	}
	return 0
}

func (m *IDMapType) GetNsid() int32 {
	if m != nil && m.Nsid != nil {
		return *m.Nsid
	}
	return 0
}

func (m *IDMapType) GetMaprange() int32 {
	if m != nil && m.Maprange != nil {
		return *m.Maprange
	}
	return 0
}

type Config struct {
	Key              *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            *string `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Config) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Config) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type Device struct {
	Name             *string   `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Config           []*Config `protobuf:"bytes,2,rep,name=config" json:"config,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Device) Reset()                    { *m = Device{} }
func (m *Device) String() string            { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()               {}
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Device) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Device) GetConfig() []*Config {
	if m != nil {
		return m.Config
	}
	return nil
}

type Snapshot struct {
	Name             *string   `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	LocalConfig      []*Config `protobuf:"bytes,2,rep,name=localConfig" json:"localConfig,omitempty"`
	Profiles         []string  `protobuf:"bytes,3,rep,name=profiles" json:"profiles,omitempty"`
	Ephemeral        *bool     `protobuf:"varint,4,req,name=ephemeral" json:"ephemeral,omitempty"`
	LocalDevices     []*Device `protobuf:"bytes,5,rep,name=localDevices" json:"localDevices,omitempty"`
	Architecture     *int32    `protobuf:"varint,6,req,name=architecture" json:"architecture,omitempty"`
	Stateful         *bool     `protobuf:"varint,7,req,name=stateful" json:"stateful,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Snapshot) Reset()                    { *m = Snapshot{} }
func (m *Snapshot) String() string            { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()               {}
func (*Snapshot) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Snapshot) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Snapshot) GetLocalConfig() []*Config {
	if m != nil {
		return m.LocalConfig
	}
	return nil
}

func (m *Snapshot) GetProfiles() []string {
	if m != nil {
		return m.Profiles
	}
	return nil
}

func (m *Snapshot) GetEphemeral() bool {
	if m != nil && m.Ephemeral != nil {
		return *m.Ephemeral
	}
	return false
}

func (m *Snapshot) GetLocalDevices() []*Device {
	if m != nil {
		return m.LocalDevices
	}
	return nil
}

func (m *Snapshot) GetArchitecture() int32 {
	if m != nil && m.Architecture != nil {
		return *m.Architecture
	}
	return 0
}

func (m *Snapshot) GetStateful() bool {
	if m != nil && m.Stateful != nil {
		return *m.Stateful
	}
	return false
}

type MigrationHeader struct {
	Fs               *MigrationFSType `protobuf:"varint,1,req,name=fs,enum=migration.MigrationFSType" json:"fs,omitempty"`
	Criu             *CRIUType        `protobuf:"varint,2,opt,name=criu,enum=migration.CRIUType" json:"criu,omitempty"`
	Idmap            []*IDMapType     `protobuf:"bytes,3,rep,name=idmap" json:"idmap,omitempty"`
	SnapshotNames    []string         `protobuf:"bytes,4,rep,name=snapshotNames" json:"snapshotNames,omitempty"`
	Snapshots        []*Snapshot      `protobuf:"bytes,5,rep,name=snapshots" json:"snapshots,omitempty"`
	Predump          *bool            `protobuf:"varint,7,opt,name=predump" json:"predump,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *MigrationHeader) Reset()                    { *m = MigrationHeader{} }
func (m *MigrationHeader) String() string            { return proto.CompactTextString(m) }
func (*MigrationHeader) ProtoMessage()               {}
func (*MigrationHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MigrationHeader) GetFs() MigrationFSType {
	if m != nil && m.Fs != nil {
		return *m.Fs
	}
	return MigrationFSType_RSYNC
}

func (m *MigrationHeader) GetCriu() CRIUType {
	if m != nil && m.Criu != nil {
		return *m.Criu
	}
	return CRIUType_CRIU_RSYNC
}

func (m *MigrationHeader) GetIdmap() []*IDMapType {
	if m != nil {
		return m.Idmap
	}
	return nil
}

func (m *MigrationHeader) GetSnapshotNames() []string {
	if m != nil {
		return m.SnapshotNames
	}
	return nil
}

func (m *MigrationHeader) GetSnapshots() []*Snapshot {
	if m != nil {
		return m.Snapshots
	}
	return nil
}

func (m *MigrationHeader) GetPredump() bool {
	if m != nil && m.Predump != nil {
		return *m.Predump
	}
	return false
}

type MigrationControl struct {
	Success *bool `protobuf:"varint,1,req,name=success" json:"success,omitempty"`
	// optional failure message if sending a failure
	Message          *string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MigrationControl) Reset()                    { *m = MigrationControl{} }
func (m *MigrationControl) String() string            { return proto.CompactTextString(m) }
func (*MigrationControl) ProtoMessage()               {}
func (*MigrationControl) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *MigrationControl) GetSuccess() bool {
	if m != nil && m.Success != nil {
		return *m.Success
	}
	return false
}

func (m *MigrationControl) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type MigrationSync struct {
	FinalPreDump     *bool  `protobuf:"varint,1,req,name=finalPreDump" json:"finalPreDump,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MigrationSync) Reset()                    { *m = MigrationSync{} }
func (m *MigrationSync) String() string            { return proto.CompactTextString(m) }
func (*MigrationSync) ProtoMessage()               {}
func (*MigrationSync) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *MigrationSync) GetFinalPreDump() bool {
	if m != nil && m.FinalPreDump != nil {
		return *m.FinalPreDump
	}
	return false
}

// This one contains statistics about dump/restore process
type DumpStatsEntry struct {
	FreezingTime       *uint32 `protobuf:"varint,1,req,name=freezing_time,json=freezingTime" json:"freezing_time,omitempty"`
	FrozenTime         *uint32 `protobuf:"varint,2,req,name=frozen_time,json=frozenTime" json:"frozen_time,omitempty"`
	MemdumpTime        *uint32 `protobuf:"varint,3,req,name=memdump_time,json=memdumpTime" json:"memdump_time,omitempty"`
	MemwriteTime       *uint32 `protobuf:"varint,4,req,name=memwrite_time,json=memwriteTime" json:"memwrite_time,omitempty"`
	PagesScanned       *uint64 `protobuf:"varint,5,req,name=pages_scanned,json=pagesScanned" json:"pages_scanned,omitempty"`
	PagesSkippedParent *uint64 `protobuf:"varint,6,req,name=pages_skipped_parent,json=pagesSkippedParent" json:"pages_skipped_parent,omitempty"`
	PagesWritten       *uint64 `protobuf:"varint,7,req,name=pages_written,json=pagesWritten" json:"pages_written,omitempty"`
	IrmapResolve       *uint32 `protobuf:"varint,8,opt,name=irmap_resolve,json=irmapResolve" json:"irmap_resolve,omitempty"`
	PagesLazy          *uint64 `protobuf:"varint,9,req,name=pages_lazy,json=pagesLazy" json:"pages_lazy,omitempty"`
	PagePipes          *uint64 `protobuf:"varint,10,opt,name=page_pipes,json=pagePipes" json:"page_pipes,omitempty"`
	PagePipeBufs       *uint64 `protobuf:"varint,11,opt,name=page_pipe_bufs,json=pagePipeBufs" json:"page_pipe_bufs,omitempty"`
	XXX_unrecognized   []byte  `json:"-"`
}

func (m *DumpStatsEntry) Reset()                    { *m = DumpStatsEntry{} }
func (m *DumpStatsEntry) String() string            { return proto.CompactTextString(m) }
func (*DumpStatsEntry) ProtoMessage()               {}
func (*DumpStatsEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DumpStatsEntry) GetFreezingTime() uint32 {
	if m != nil && m.FreezingTime != nil {
		return *m.FreezingTime
	}
	return 0
}

func (m *DumpStatsEntry) GetFrozenTime() uint32 {
	if m != nil && m.FrozenTime != nil {
		return *m.FrozenTime
	}
	return 0
}

func (m *DumpStatsEntry) GetMemdumpTime() uint32 {
	if m != nil && m.MemdumpTime != nil {
		return *m.MemdumpTime
	}
	return 0
}

func (m *DumpStatsEntry) GetMemwriteTime() uint32 {
	if m != nil && m.MemwriteTime != nil {
		return *m.MemwriteTime
	}
	return 0
}

func (m *DumpStatsEntry) GetPagesScanned() uint64 {
	if m != nil && m.PagesScanned != nil {
		return *m.PagesScanned
	}
	return 0
}

func (m *DumpStatsEntry) GetPagesSkippedParent() uint64 {
	if m != nil && m.PagesSkippedParent != nil {
		return *m.PagesSkippedParent
	}
	return 0
}

func (m *DumpStatsEntry) GetPagesWritten() uint64 {
	if m != nil && m.PagesWritten != nil {
		return *m.PagesWritten
	}
	return 0
}

func (m *DumpStatsEntry) GetIrmapResolve() uint32 {
	if m != nil && m.IrmapResolve != nil {
		return *m.IrmapResolve
	}
	return 0
}

func (m *DumpStatsEntry) GetPagesLazy() uint64 {
	if m != nil && m.PagesLazy != nil {
		return *m.PagesLazy
	}
	return 0
}

func (m *DumpStatsEntry) GetPagePipes() uint64 {
	if m != nil && m.PagePipes != nil {
		return *m.PagePipes
	}
	return 0
}

func (m *DumpStatsEntry) GetPagePipeBufs() uint64 {
	if m != nil && m.PagePipeBufs != nil {
		return *m.PagePipeBufs
	}
	return 0
}

type RestoreStatsEntry struct {
	PagesCompared    *uint64 `protobuf:"varint,1,req,name=pages_compared,json=pagesCompared" json:"pages_compared,omitempty"`
	PagesSkippedCow  *uint64 `protobuf:"varint,2,req,name=pages_skipped_cow,json=pagesSkippedCow" json:"pages_skipped_cow,omitempty"`
	ForkingTime      *uint32 `protobuf:"varint,3,req,name=forking_time,json=forkingTime" json:"forking_time,omitempty"`
	RestoreTime      *uint32 `protobuf:"varint,4,req,name=restore_time,json=restoreTime" json:"restore_time,omitempty"`
	PagesRestored    *uint64 `protobuf:"varint,5,opt,name=pages_restored,json=pagesRestored" json:"pages_restored,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RestoreStatsEntry) Reset()                    { *m = RestoreStatsEntry{} }
func (m *RestoreStatsEntry) String() string            { return proto.CompactTextString(m) }
func (*RestoreStatsEntry) ProtoMessage()               {}
func (*RestoreStatsEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RestoreStatsEntry) GetPagesCompared() uint64 {
	if m != nil && m.PagesCompared != nil {
		return *m.PagesCompared
	}
	return 0
}

func (m *RestoreStatsEntry) GetPagesSkippedCow() uint64 {
	if m != nil && m.PagesSkippedCow != nil {
		return *m.PagesSkippedCow
	}
	return 0
}

func (m *RestoreStatsEntry) GetForkingTime() uint32 {
	if m != nil && m.ForkingTime != nil {
		return *m.ForkingTime
	}
	return 0
}

func (m *RestoreStatsEntry) GetRestoreTime() uint32 {
	if m != nil && m.RestoreTime != nil {
		return *m.RestoreTime
	}
	return 0
}

func (m *RestoreStatsEntry) GetPagesRestored() uint64 {
	if m != nil && m.PagesRestored != nil {
		return *m.PagesRestored
	}
	return 0
}

type StatsEntry struct {
	Dump             *DumpStatsEntry    `protobuf:"bytes,1,opt,name=dump" json:"dump,omitempty"`
	Restore          *RestoreStatsEntry `protobuf:"bytes,2,opt,name=restore" json:"restore,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *StatsEntry) Reset()                    { *m = StatsEntry{} }
func (m *StatsEntry) String() string            { return proto.CompactTextString(m) }
func (*StatsEntry) ProtoMessage()               {}
func (*StatsEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *StatsEntry) GetDump() *DumpStatsEntry {
	if m != nil {
		return m.Dump
	}
	return nil
}

func (m *StatsEntry) GetRestore() *RestoreStatsEntry {
	if m != nil {
		return m.Restore
	}
	return nil
}

func init() {
	proto.RegisterType((*IDMapType)(nil), "migration.IDMapType")
	proto.RegisterType((*Config)(nil), "migration.Config")
	proto.RegisterType((*Device)(nil), "migration.Device")
	proto.RegisterType((*Snapshot)(nil), "migration.Snapshot")
	proto.RegisterType((*MigrationHeader)(nil), "migration.MigrationHeader")
	proto.RegisterType((*MigrationControl)(nil), "migration.MigrationControl")
	proto.RegisterType((*MigrationSync)(nil), "migration.MigrationSync")
	proto.RegisterType((*DumpStatsEntry)(nil), "migration.dump_stats_entry")
	proto.RegisterType((*RestoreStatsEntry)(nil), "migration.restore_stats_entry")
	proto.RegisterType((*StatsEntry)(nil), "migration.stats_entry")
	proto.RegisterEnum("migration.MigrationFSType", MigrationFSType_name, MigrationFSType_value)
	proto.RegisterEnum("migration.CRIUType", CRIUType_name, CRIUType_value)
}

func init() { proto.RegisterFile("lxd/migration/migrate.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 894 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x5b, 0x6f, 0xdb, 0x36,
	0x14, 0x9e, 0x64, 0xf9, 0xa2, 0x23, 0x3b, 0x75, 0xd9, 0x60, 0x10, 0xda, 0x5d, 0x3c, 0xb5, 0xc3,
	0x3c, 0x3f, 0x24, 0x9d, 0x8b, 0x01, 0xdb, 0xe3, 0xe2, 0x2c, 0x6b, 0x81, 0x36, 0x0b, 0xe8, 0x14,
	0xc3, 0xf6, 0x22, 0x70, 0xd2, 0x91, 0x43, 0x44, 0x37, 0x90, 0x72, 0x52, 0xe7, 0x65, 0x2f, 0xdb,
	0x4f, 0xd9, 0x4f, 0xda, 0xff, 0x19, 0x48, 0x4a, 0x8a, 0xbc, 0x15, 0xe8, 0x1b, 0xcf, 0x77, 0x3e,
	0x9e, 0xef, 0xdc, 0x48, 0x78, 0x92, 0xbe, 0x8b, 0x8f, 0x33, 0xbe, 0x11, 0xac, 0xe2, 0x45, 0x5e,
	0x9f, 0xf0, 0xa8, 0x14, 0x45, 0x55, 0x10, 0xb7, 0x75, 0x04, 0x7f, 0x80, 0xfb, 0xea, 0xf4, 0x0d,
	0x2b, 0x2f, 0x77, 0x25, 0x92, 0x43, 0xe8, 0x73, 0xb9, 0xe5, 0xb1, 0x6f, 0xcd, 0xec, 0xf9, 0x88,
	0x1a, 0xc3, 0xa0, 0x1b, 0x1e, 0xfb, 0x76, 0x83, 0x6e, 0x78, 0x4c, 0x3e, 0x86, 0xc1, 0x55, 0x21,
	0x2b, 0x1e, 0xfb, 0xbd, 0x99, 0x3d, 0xef, 0xd3, 0xda, 0x22, 0x04, 0x9c, 0x5c, 0xf2, 0xd8, 0x77,
	0x34, 0xaa, 0xcf, 0xe4, 0x31, 0x8c, 0x32, 0x56, 0x0a, 0x96, 0x6f, 0xd0, 0xef, 0x6b, 0xbc, 0xb5,
	0x83, 0xe7, 0x30, 0x58, 0x15, 0x79, 0xc2, 0x37, 0x64, 0x0a, 0xbd, 0x6b, 0xdc, 0x69, 0x6d, 0x97,
	0xaa, 0xa3, 0x52, 0xbe, 0x61, 0xe9, 0x16, 0xb5, 0xb2, 0x4b, 0x8d, 0x11, 0xfc, 0x04, 0x83, 0x53,
	0xbc, 0xe1, 0x11, 0x6a, 0x2d, 0x96, 0x61, 0x7d, 0x45, 0x9f, 0xc9, 0xd7, 0x30, 0x88, 0x74, 0x3c,
	0xdf, 0x9e, 0xf5, 0xe6, 0xde, 0xf2, 0xe1, 0x51, 0x5b, 0xec, 0x91, 0x11, 0xa2, 0x35, 0x21, 0xf8,
	0xd3, 0x86, 0xd1, 0x3a, 0x67, 0xa5, 0xbc, 0x2a, 0xaa, 0xf7, 0xc6, 0x7a, 0x01, 0x5e, 0x5a, 0x44,
	0x2c, 0x5d, 0x7d, 0x20, 0x60, 0x97, 0xa5, 0x8a, 0x2d, 0x45, 0x91, 0xf0, 0x14, 0xa5, 0xdf, 0x9b,
	0xf5, 0xe6, 0x2e, 0x6d, 0x6d, 0xf2, 0x09, 0xb8, 0x58, 0x5e, 0x61, 0x86, 0x82, 0xa5, 0xba, 0x43,
	0x23, 0x7a, 0x0f, 0x90, 0x6f, 0x61, 0xac, 0x03, 0x99, 0xea, 0xa4, 0xdf, 0xff, 0x9f, 0x9e, 0xf1,
	0xd0, 0x3d, 0x1a, 0x09, 0x60, 0xcc, 0x44, 0x74, 0xc5, 0x2b, 0x8c, 0xaa, 0xad, 0x40, 0x7f, 0xa0,
	0x3b, 0xbc, 0x87, 0xa9, 0xa4, 0x64, 0xc5, 0x2a, 0x4c, 0xb6, 0xa9, 0x3f, 0xd4, 0xba, 0xad, 0x1d,
	0xfc, 0x65, 0xc3, 0x83, 0x37, 0x8d, 0xc4, 0x4b, 0x64, 0x31, 0x0a, 0xb2, 0x00, 0x3b, 0x91, 0xba,
	0x17, 0x07, 0xcb, 0xc7, 0x9d, 0x04, 0x5a, 0xde, 0xd9, 0x5a, 0x6d, 0x0c, 0xb5, 0x13, 0x49, 0xbe,
	0x02, 0x27, 0x12, 0x7c, 0xeb, 0xdb, 0x33, 0x6b, 0x7e, 0xb0, 0x7c, 0xd4, 0x6d, 0x0f, 0x7d, 0xf5,
	0x56, 0xd3, 0x34, 0x81, 0x2c, 0xa0, 0xcf, 0xe3, 0x8c, 0x95, 0xba, 0x2d, 0xde, 0xf2, 0xb0, 0xc3,
	0x6c, 0x77, 0x90, 0x1a, 0x0a, 0x79, 0x06, 0x13, 0x59, 0x8f, 0xe6, 0x9c, 0x65, 0x28, 0x7d, 0x47,
	0xb7, 0x72, 0x1f, 0x24, 0xdf, 0x80, 0xdb, 0x00, 0x4d, 0xbb, 0xba, 0xfa, 0xcd, 0x70, 0xe9, 0x3d,
	0x8b, 0xf8, 0x30, 0x2c, 0x05, 0xc6, 0xdb, 0xac, 0xf4, 0x87, 0x33, 0x6b, 0x3e, 0xa2, 0x8d, 0x19,
	0x9c, 0xc1, 0xb4, 0x2d, 0x6f, 0x55, 0xe4, 0x95, 0x28, 0x52, 0xc5, 0x96, 0xdb, 0x28, 0x42, 0x29,
	0xeb, 0x37, 0xd1, 0x98, 0xca, 0x93, 0xa1, 0x94, 0x6c, 0x83, 0xba, 0x70, 0x97, 0x36, 0x66, 0xf0,
	0x02, 0x26, 0x6d, 0x9c, 0xf5, 0x2e, 0x8f, 0xd4, 0x80, 0x12, 0x9e, 0xb3, 0xf4, 0x42, 0xe0, 0xa9,
	0xd2, 0x35, 0x91, 0xf6, 0xb0, 0xe0, 0xef, 0x1e, 0x4c, 0x55, 0x16, 0xa1, 0x1a, 0x8b, 0x0c, 0x31,
	0xaf, 0xc4, 0x8e, 0x3c, 0x85, 0x49, 0x22, 0x10, 0xef, 0x78, 0xbe, 0x09, 0x2b, 0x5e, 0x2f, 0xe7,
	0x84, 0x8e, 0x1b, 0xf0, 0x92, 0x67, 0x48, 0x3e, 0x07, 0x2f, 0x11, 0xc5, 0x1d, 0xe6, 0x86, 0x62,
	0x6b, 0x0a, 0x18, 0x48, 0x13, 0xbe, 0x80, 0x71, 0x86, 0x99, 0x0e, 0xae, 0x19, 0x3d, 0xcd, 0xf0,
	0x6a, 0x4c, 0x53, 0x9e, 0xc2, 0x24, 0xc3, 0xec, 0x56, 0xf0, 0x0a, 0x0d, 0xc7, 0x31, 0x42, 0x0d,
	0xd8, 0x90, 0x4a, 0xb6, 0x41, 0x19, 0xca, 0x88, 0xe5, 0x39, 0xc6, 0xfa, 0x29, 0x3b, 0x74, 0xac,
	0xc1, 0xb5, 0xc1, 0xc8, 0x73, 0x38, 0xac, 0x49, 0xd7, 0xbc, 0x2c, 0x31, 0x0e, 0x4b, 0x26, 0x30,
	0xaf, 0xf4, 0x52, 0x3a, 0x94, 0x18, 0xae, 0x71, 0x5d, 0x68, 0xcf, 0x7d, 0x58, 0xa5, 0x54, 0x61,
	0xae, 0xf7, 0xb3, 0x09, 0xfb, 0x8b, 0xc1, 0x14, 0x89, 0x8b, 0x8c, 0x95, 0xa1, 0x40, 0x59, 0xa4,
	0x37, 0xe8, 0x8f, 0x66, 0x96, 0x4a, 0x50, 0x83, 0xd4, 0x60, 0xe4, 0x53, 0x00, 0x13, 0x29, 0x65,
	0x77, 0x3b, 0xdf, 0xd5, 0x61, 0x5c, 0x8d, 0xbc, 0x66, 0x77, 0xbb, 0xc6, 0x1d, 0x96, 0xbc, 0x44,
	0xe9, 0xc3, 0xcc, 0x6a, 0xdc, 0x17, 0x0a, 0x20, 0xcf, 0xe0, 0xa0, 0x75, 0x87, 0xbf, 0x6f, 0x13,
	0xe9, 0x7b, 0x9a, 0x32, 0x6e, 0x28, 0x27, 0xdb, 0x44, 0x06, 0xff, 0x58, 0xf0, 0x48, 0xa0, 0xac,
	0x0a, 0x81, 0x7b, 0xa3, 0xfa, 0xd2, 0xdc, 0x96, 0x61, 0x54, 0x64, 0xaa, 0x64, 0xf3, 0x87, 0x3a,
	0xd4, 0xd4, 0xb6, 0xaa, 0x41, 0xb2, 0x80, 0x87, 0xfb, 0xed, 0x89, 0x8a, 0x5b, 0x3d, 0x32, 0x87,
	0x3e, 0xe8, 0xf6, 0x66, 0x55, 0xdc, 0xaa, 0xb9, 0x25, 0x85, 0xb8, 0x6e, 0x87, 0x5f, 0xcf, 0xad,
	0xc6, 0x9a, 0xd1, 0x36, 0xc9, 0x74, 0xc6, 0xe6, 0xd5, 0x98, 0xa6, 0xb4, 0x89, 0xd5, 0xa0, 0x1a,
	0x9b, 0xd5, 0x26, 0x46, 0x6b, 0x30, 0x78, 0x07, 0x5e, 0xb7, 0x9c, 0x63, 0x70, 0x62, 0xb3, 0xaa,
	0xd6, 0xdc, 0x5b, 0x3e, 0xe9, 0xbc, 0xa9, 0xff, 0x2e, 0x29, 0xd5, 0x44, 0xf2, 0x1d, 0x0c, 0x6b,
	0x01, 0xfd, 0x1c, 0xbc, 0xe5, 0x67, 0x9d, 0x3b, 0xef, 0x69, 0x18, 0x6d, 0xe8, 0x8b, 0xef, 0x3b,
	0xbf, 0x8f, 0xf9, 0x55, 0x88, 0x0b, 0x7d, 0xba, 0xfe, 0xf5, 0x7c, 0x35, 0xfd, 0x48, 0x1d, 0x4f,
	0x2e, 0xe9, 0xd9, 0x7a, 0x6a, 0x91, 0x21, 0xf4, 0x7e, 0x3b, 0x5b, 0x4f, 0x6d, 0x75, 0xa0, 0x27,
	0xa7, 0xd3, 0xde, 0xe2, 0x18, 0x46, 0xcd, 0x17, 0x43, 0x0e, 0x00, 0xd4, 0x39, 0xec, 0x5c, 0xbc,
	0x78, 0xf9, 0xc3, 0xdb, 0xd7, 0x53, 0x8b, 0x8c, 0xc0, 0x39, 0xff, 0xf9, 0xfc, 0xc7, 0xa9, 0xfd,
	0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x86, 0x78, 0x5b, 0x16, 0x07, 0x00, 0x00,
}