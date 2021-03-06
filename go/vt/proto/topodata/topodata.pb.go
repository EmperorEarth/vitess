// Code generated by protoc-gen-go.
// source: topodata.proto
// DO NOT EDIT!

/*
Package topodata is a generated protocol buffer package.

It is generated from these files:
	topodata.proto

It has these top-level messages:
	KeyRange
	TabletAlias
	Tablet
	Shard
	Keyspace
	ShardReplication
	ShardReference
	SrvKeyspace
*/
package topodata

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

// KeyspaceIdType describes the type of the sharding key for a
// range-based sharded keyspace.
type KeyspaceIdType int32

const (
	// UNSET is the default value, when range-based sharding is not used.
	KeyspaceIdType_UNSET KeyspaceIdType = 0
	// UINT64 is when uint64 value is used.
	// This is represented as 'unsigned bigint' in mysql
	KeyspaceIdType_UINT64 KeyspaceIdType = 1
	// BYTES is when an array of bytes is used.
	// This is represented as 'varbinary' in mysql
	KeyspaceIdType_BYTES KeyspaceIdType = 2
)

var KeyspaceIdType_name = map[int32]string{
	0: "UNSET",
	1: "UINT64",
	2: "BYTES",
}
var KeyspaceIdType_value = map[string]int32{
	"UNSET":  0,
	"UINT64": 1,
	"BYTES":  2,
}

func (x KeyspaceIdType) String() string {
	return proto.EnumName(KeyspaceIdType_name, int32(x))
}
func (KeyspaceIdType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// TabletType represents the type of a given tablet.
type TabletType int32

const (
	// UNKNOWN is not a valid value.
	TabletType_UNKNOWN TabletType = 0
	// MASTER is the master server for the shard. Only MASTER allows DMLs.
	TabletType_MASTER TabletType = 1
	// REPLICA is a slave type. It is used to serve live traffic.
	// A REPLICA can be promoted to MASTER. A demoted MASTER will go to REPLICA.
	TabletType_REPLICA TabletType = 2
	// RDONLY (old name) / BATCH (new name) is used to serve traffic for
	// long-running jobs. It is a separate type from REPLICA so
	// long-running queries don't affect web-like traffic.
	TabletType_RDONLY TabletType = 3
	TabletType_BATCH  TabletType = 3
	// SPARE is a type of servers that cannot serve queries, but is available
	// in case an extra server is needed.
	TabletType_SPARE TabletType = 4
	// EXPERIMENTAL is like SPARE, except it can serve queries. This
	// type can be used for usages not planned by Vitess, like online
	// export to another storage engine.
	TabletType_EXPERIMENTAL TabletType = 5
	// BACKUP is the type a server goes to when taking a backup. No queries
	// can be served in BACKUP mode.
	TabletType_BACKUP TabletType = 6
	// RESTORE is the type a server uses when restoring a backup, at
	// startup time.  No queries can be served in RESTORE mode.
	TabletType_RESTORE TabletType = 7
	// WORKER is the type a server goes into when used by a vtworker
	// process to perform an offline action. It is a serving type (as
	// the vtworker processes may need queries to run). In this state,
	// this tablet is dedicated to the vtworker process that uses it.
	TabletType_WORKER TabletType = 8
)

var TabletType_name = map[int32]string{
	0: "UNKNOWN",
	1: "MASTER",
	2: "REPLICA",
	3: "RDONLY",
	// Duplicate value: 3: "BATCH",
	4: "SPARE",
	5: "EXPERIMENTAL",
	6: "BACKUP",
	7: "RESTORE",
	8: "WORKER",
}
var TabletType_value = map[string]int32{
	"UNKNOWN":      0,
	"MASTER":       1,
	"REPLICA":      2,
	"RDONLY":       3,
	"BATCH":        3,
	"SPARE":        4,
	"EXPERIMENTAL": 5,
	"BACKUP":       6,
	"RESTORE":      7,
	"WORKER":       8,
}

func (x TabletType) String() string {
	return proto.EnumName(TabletType_name, int32(x))
}
func (TabletType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// KeyRange describes a range of sharding keys, when range-based
// sharding is used.
type KeyRange struct {
	Start []byte `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End   []byte `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (m *KeyRange) Reset()                    { *m = KeyRange{} }
func (m *KeyRange) String() string            { return proto.CompactTextString(m) }
func (*KeyRange) ProtoMessage()               {}
func (*KeyRange) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// TabletAlias is a globally unique tablet identifier.
type TabletAlias struct {
	// cell is the cell (or datacenter) the tablet is in
	Cell string `protobuf:"bytes,1,opt,name=cell" json:"cell,omitempty"`
	// uid is a unique id for this tablet within the shard
	// (this is the MySQL server id as well).
	Uid uint32 `protobuf:"varint,2,opt,name=uid" json:"uid,omitempty"`
}

func (m *TabletAlias) Reset()                    { *m = TabletAlias{} }
func (m *TabletAlias) String() string            { return proto.CompactTextString(m) }
func (*TabletAlias) ProtoMessage()               {}
func (*TabletAlias) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Tablet represents information about a running instance of vttablet.
type Tablet struct {
	// alias is the unique name of the tablet.
	Alias *TabletAlias `protobuf:"bytes,1,opt,name=alias" json:"alias,omitempty"`
	// Fully qualified domain name of the host.
	Hostname string `protobuf:"bytes,2,opt,name=hostname" json:"hostname,omitempty"`
	// IP address, stored as a string.
	Ip string `protobuf:"bytes,3,opt,name=ip" json:"ip,omitempty"`
	// Map of named ports. Normally this should include vt, grpc, and mysql.
	PortMap map[string]int32 `protobuf:"bytes,4,rep,name=port_map,json=portMap" json:"port_map,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	// Keyspace name.
	Keyspace string `protobuf:"bytes,5,opt,name=keyspace" json:"keyspace,omitempty"`
	// Shard name. If range based sharding is used, it should match
	// key_range.
	Shard string `protobuf:"bytes,6,opt,name=shard" json:"shard,omitempty"`
	// If range based sharding is used, range for the tablet's shard.
	KeyRange *KeyRange `protobuf:"bytes,7,opt,name=key_range,json=keyRange" json:"key_range,omitempty"`
	// type is the current type of the tablet.
	Type TabletType `protobuf:"varint,8,opt,name=type,enum=topodata.TabletType" json:"type,omitempty"`
	// It this is set, it is used as the database name instead of the
	// normal "vt_" + keyspace.
	DbNameOverride string `protobuf:"bytes,9,opt,name=db_name_override,json=dbNameOverride" json:"db_name_override,omitempty"`
	// tablet tags
	Tags map[string]string `protobuf:"bytes,10,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Tablet) Reset()                    { *m = Tablet{} }
func (m *Tablet) String() string            { return proto.CompactTextString(m) }
func (*Tablet) ProtoMessage()               {}
func (*Tablet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Tablet) GetAlias() *TabletAlias {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *Tablet) GetPortMap() map[string]int32 {
	if m != nil {
		return m.PortMap
	}
	return nil
}

func (m *Tablet) GetKeyRange() *KeyRange {
	if m != nil {
		return m.KeyRange
	}
	return nil
}

func (m *Tablet) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

// A Shard contains data about a subset of the data whithin a keyspace.
type Shard struct {
	// No lock is necessary to update this field, when for instance
	// TabletExternallyReparented updates this. However, we lock the
	// shard for reparenting operations (InitShardMaster,
	// PlannedReparentShard,EmergencyReparentShard), to guarantee
	// exclusive operation.
	MasterAlias *TabletAlias `protobuf:"bytes,1,opt,name=master_alias,json=masterAlias" json:"master_alias,omitempty"`
	// key_range is the KeyRange for this shard. It can be unset if:
	// - we are not using range-based sharding in this shard.
	// - the shard covers the entire keyrange.
	// This must match the shard name based on our other conventions, but
	// helpful to have it decomposed here.
	// Once set at creation time, it is never changed.
	KeyRange *KeyRange `protobuf:"bytes,2,opt,name=key_range,json=keyRange" json:"key_range,omitempty"`
	// served_types has at most one entry per TabletType
	// The keyspace lock is always taken when changing this.
	ServedTypes []*Shard_ServedType `protobuf:"bytes,3,rep,name=served_types,json=servedTypes" json:"served_types,omitempty"`
	// SourceShards is the list of shards we're replicating from,
	// using filtered replication.
	// The keyspace lock is always taken when changing this.
	SourceShards []*Shard_SourceShard `protobuf:"bytes,4,rep,name=source_shards,json=sourceShards" json:"source_shards,omitempty"`
	// Cells is the list of cells that contain tablets for this shard.
	// No lock is necessary to update this field.
	Cells []string `protobuf:"bytes,5,rep,name=cells" json:"cells,omitempty"`
	// tablet_controls has at most one entry per TabletType.
	// The keyspace lock is always taken when changing this.
	TabletControls []*Shard_TabletControl `protobuf:"bytes,6,rep,name=tablet_controls,json=tabletControls" json:"tablet_controls,omitempty"`
}

func (m *Shard) Reset()                    { *m = Shard{} }
func (m *Shard) String() string            { return proto.CompactTextString(m) }
func (*Shard) ProtoMessage()               {}
func (*Shard) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Shard) GetMasterAlias() *TabletAlias {
	if m != nil {
		return m.MasterAlias
	}
	return nil
}

func (m *Shard) GetKeyRange() *KeyRange {
	if m != nil {
		return m.KeyRange
	}
	return nil
}

func (m *Shard) GetServedTypes() []*Shard_ServedType {
	if m != nil {
		return m.ServedTypes
	}
	return nil
}

func (m *Shard) GetSourceShards() []*Shard_SourceShard {
	if m != nil {
		return m.SourceShards
	}
	return nil
}

func (m *Shard) GetTabletControls() []*Shard_TabletControl {
	if m != nil {
		return m.TabletControls
	}
	return nil
}

// ServedType is an entry in the served_types
type Shard_ServedType struct {
	TabletType TabletType `protobuf:"varint,1,opt,name=tablet_type,json=tabletType,enum=topodata.TabletType" json:"tablet_type,omitempty"`
	Cells      []string   `protobuf:"bytes,2,rep,name=cells" json:"cells,omitempty"`
}

func (m *Shard_ServedType) Reset()                    { *m = Shard_ServedType{} }
func (m *Shard_ServedType) String() string            { return proto.CompactTextString(m) }
func (*Shard_ServedType) ProtoMessage()               {}
func (*Shard_ServedType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

// SourceShard represents a data source for filtered replication
// accross shards. When this is used in a destination shard, the master
// of that shard will run filtered replication.
type Shard_SourceShard struct {
	// Uid is the unique ID for this SourceShard object.
	Uid uint32 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	// the source keyspace
	Keyspace string `protobuf:"bytes,2,opt,name=keyspace" json:"keyspace,omitempty"`
	// the source shard
	Shard string `protobuf:"bytes,3,opt,name=shard" json:"shard,omitempty"`
	// the source shard keyrange
	KeyRange *KeyRange `protobuf:"bytes,4,opt,name=key_range,json=keyRange" json:"key_range,omitempty"`
	// the source table list to replicate
	Tables []string `protobuf:"bytes,5,rep,name=tables" json:"tables,omitempty"`
}

func (m *Shard_SourceShard) Reset()                    { *m = Shard_SourceShard{} }
func (m *Shard_SourceShard) String() string            { return proto.CompactTextString(m) }
func (*Shard_SourceShard) ProtoMessage()               {}
func (*Shard_SourceShard) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 1} }

func (m *Shard_SourceShard) GetKeyRange() *KeyRange {
	if m != nil {
		return m.KeyRange
	}
	return nil
}

// TabletControl controls tablet's behavior
type Shard_TabletControl struct {
	// which tablet type is affected
	TabletType TabletType `protobuf:"varint,1,opt,name=tablet_type,json=tabletType,enum=topodata.TabletType" json:"tablet_type,omitempty"`
	Cells      []string   `protobuf:"bytes,2,rep,name=cells" json:"cells,omitempty"`
	// what to do
	DisableQueryService bool     `protobuf:"varint,3,opt,name=disable_query_service,json=disableQueryService" json:"disable_query_service,omitempty"`
	BlacklistedTables   []string `protobuf:"bytes,4,rep,name=blacklisted_tables,json=blacklistedTables" json:"blacklisted_tables,omitempty"`
}

func (m *Shard_TabletControl) Reset()                    { *m = Shard_TabletControl{} }
func (m *Shard_TabletControl) String() string            { return proto.CompactTextString(m) }
func (*Shard_TabletControl) ProtoMessage()               {}
func (*Shard_TabletControl) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 2} }

// A Keyspace contains data about a keyspace.
type Keyspace struct {
	// name of the column used for sharding
	// empty if the keyspace is not sharded
	ShardingColumnName string `protobuf:"bytes,1,opt,name=sharding_column_name,json=shardingColumnName" json:"sharding_column_name,omitempty"`
	// type of the column used for sharding
	// UNSET if the keyspace is not sharded
	ShardingColumnType KeyspaceIdType `protobuf:"varint,2,opt,name=sharding_column_type,json=shardingColumnType,enum=topodata.KeyspaceIdType" json:"sharding_column_type,omitempty"`
	// ServedFrom will redirect the appropriate traffic to
	// another keyspace.
	ServedFroms []*Keyspace_ServedFrom `protobuf:"bytes,4,rep,name=served_froms,json=servedFroms" json:"served_froms,omitempty"`
}

func (m *Keyspace) Reset()                    { *m = Keyspace{} }
func (m *Keyspace) String() string            { return proto.CompactTextString(m) }
func (*Keyspace) ProtoMessage()               {}
func (*Keyspace) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Keyspace) GetServedFroms() []*Keyspace_ServedFrom {
	if m != nil {
		return m.ServedFroms
	}
	return nil
}

// ServedFrom indicates a relationship between a TabletType and the
// keyspace name that's serving it.
type Keyspace_ServedFrom struct {
	// the tablet type (key for the map)
	TabletType TabletType `protobuf:"varint,1,opt,name=tablet_type,json=tabletType,enum=topodata.TabletType" json:"tablet_type,omitempty"`
	// the cells to limit this to
	Cells []string `protobuf:"bytes,2,rep,name=cells" json:"cells,omitempty"`
	// the keyspace name that's serving it
	Keyspace string `protobuf:"bytes,3,opt,name=keyspace" json:"keyspace,omitempty"`
}

func (m *Keyspace_ServedFrom) Reset()                    { *m = Keyspace_ServedFrom{} }
func (m *Keyspace_ServedFrom) String() string            { return proto.CompactTextString(m) }
func (*Keyspace_ServedFrom) ProtoMessage()               {}
func (*Keyspace_ServedFrom) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

// ShardReplication describes the MySQL replication relationships
// whithin a cell.
type ShardReplication struct {
	// Note there can be only one Node in this array
	// for a given tablet.
	Nodes []*ShardReplication_Node `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *ShardReplication) Reset()                    { *m = ShardReplication{} }
func (m *ShardReplication) String() string            { return proto.CompactTextString(m) }
func (*ShardReplication) ProtoMessage()               {}
func (*ShardReplication) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ShardReplication) GetNodes() []*ShardReplication_Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

// Node describes a tablet instance within the cell
type ShardReplication_Node struct {
	TabletAlias *TabletAlias `protobuf:"bytes,1,opt,name=tablet_alias,json=tabletAlias" json:"tablet_alias,omitempty"`
}

func (m *ShardReplication_Node) Reset()                    { *m = ShardReplication_Node{} }
func (m *ShardReplication_Node) String() string            { return proto.CompactTextString(m) }
func (*ShardReplication_Node) ProtoMessage()               {}
func (*ShardReplication_Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

func (m *ShardReplication_Node) GetTabletAlias() *TabletAlias {
	if m != nil {
		return m.TabletAlias
	}
	return nil
}

// ShardReference is used as a pointer from a SrvKeyspace to a Shard
type ShardReference struct {
	// Copied from Shard.
	Name     string    `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	KeyRange *KeyRange `protobuf:"bytes,2,opt,name=key_range,json=keyRange" json:"key_range,omitempty"`
}

func (m *ShardReference) Reset()                    { *m = ShardReference{} }
func (m *ShardReference) String() string            { return proto.CompactTextString(m) }
func (*ShardReference) ProtoMessage()               {}
func (*ShardReference) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ShardReference) GetKeyRange() *KeyRange {
	if m != nil {
		return m.KeyRange
	}
	return nil
}

// SrvKeyspace is a rollup node for the keyspace itself.
type SrvKeyspace struct {
	// The partitions this keyspace is serving, per tablet type.
	Partitions []*SrvKeyspace_KeyspacePartition `protobuf:"bytes,1,rep,name=partitions" json:"partitions,omitempty"`
	// copied from Keyspace
	ShardingColumnName string                    `protobuf:"bytes,2,opt,name=sharding_column_name,json=shardingColumnName" json:"sharding_column_name,omitempty"`
	ShardingColumnType KeyspaceIdType            `protobuf:"varint,3,opt,name=sharding_column_type,json=shardingColumnType,enum=topodata.KeyspaceIdType" json:"sharding_column_type,omitempty"`
	ServedFrom         []*SrvKeyspace_ServedFrom `protobuf:"bytes,4,rep,name=served_from,json=servedFrom" json:"served_from,omitempty"`
}

func (m *SrvKeyspace) Reset()                    { *m = SrvKeyspace{} }
func (m *SrvKeyspace) String() string            { return proto.CompactTextString(m) }
func (*SrvKeyspace) ProtoMessage()               {}
func (*SrvKeyspace) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *SrvKeyspace) GetPartitions() []*SrvKeyspace_KeyspacePartition {
	if m != nil {
		return m.Partitions
	}
	return nil
}

func (m *SrvKeyspace) GetServedFrom() []*SrvKeyspace_ServedFrom {
	if m != nil {
		return m.ServedFrom
	}
	return nil
}

type SrvKeyspace_KeyspacePartition struct {
	// The type this partition applies to.
	ServedType TabletType `protobuf:"varint,1,opt,name=served_type,json=servedType,enum=topodata.TabletType" json:"served_type,omitempty"`
	// List of non-overlapping continuous shards sorted by range.
	ShardReferences []*ShardReference `protobuf:"bytes,2,rep,name=shard_references,json=shardReferences" json:"shard_references,omitempty"`
}

func (m *SrvKeyspace_KeyspacePartition) Reset()         { *m = SrvKeyspace_KeyspacePartition{} }
func (m *SrvKeyspace_KeyspacePartition) String() string { return proto.CompactTextString(m) }
func (*SrvKeyspace_KeyspacePartition) ProtoMessage()    {}
func (*SrvKeyspace_KeyspacePartition) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{7, 0}
}

func (m *SrvKeyspace_KeyspacePartition) GetShardReferences() []*ShardReference {
	if m != nil {
		return m.ShardReferences
	}
	return nil
}

// ServedFrom indicates a relationship between a TabletType and the
// keyspace name that's serving it.
type SrvKeyspace_ServedFrom struct {
	// the tablet type
	TabletType TabletType `protobuf:"varint,1,opt,name=tablet_type,json=tabletType,enum=topodata.TabletType" json:"tablet_type,omitempty"`
	// the keyspace name that's serving it
	Keyspace string `protobuf:"bytes,2,opt,name=keyspace" json:"keyspace,omitempty"`
}

func (m *SrvKeyspace_ServedFrom) Reset()                    { *m = SrvKeyspace_ServedFrom{} }
func (m *SrvKeyspace_ServedFrom) String() string            { return proto.CompactTextString(m) }
func (*SrvKeyspace_ServedFrom) ProtoMessage()               {}
func (*SrvKeyspace_ServedFrom) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 1} }

func init() {
	proto.RegisterType((*KeyRange)(nil), "topodata.KeyRange")
	proto.RegisterType((*TabletAlias)(nil), "topodata.TabletAlias")
	proto.RegisterType((*Tablet)(nil), "topodata.Tablet")
	proto.RegisterType((*Shard)(nil), "topodata.Shard")
	proto.RegisterType((*Shard_ServedType)(nil), "topodata.Shard.ServedType")
	proto.RegisterType((*Shard_SourceShard)(nil), "topodata.Shard.SourceShard")
	proto.RegisterType((*Shard_TabletControl)(nil), "topodata.Shard.TabletControl")
	proto.RegisterType((*Keyspace)(nil), "topodata.Keyspace")
	proto.RegisterType((*Keyspace_ServedFrom)(nil), "topodata.Keyspace.ServedFrom")
	proto.RegisterType((*ShardReplication)(nil), "topodata.ShardReplication")
	proto.RegisterType((*ShardReplication_Node)(nil), "topodata.ShardReplication.Node")
	proto.RegisterType((*ShardReference)(nil), "topodata.ShardReference")
	proto.RegisterType((*SrvKeyspace)(nil), "topodata.SrvKeyspace")
	proto.RegisterType((*SrvKeyspace_KeyspacePartition)(nil), "topodata.SrvKeyspace.KeyspacePartition")
	proto.RegisterType((*SrvKeyspace_ServedFrom)(nil), "topodata.SrvKeyspace.ServedFrom")
	proto.RegisterEnum("topodata.KeyspaceIdType", KeyspaceIdType_name, KeyspaceIdType_value)
	proto.RegisterEnum("topodata.TabletType", TabletType_name, TabletType_value)
}

func init() { proto.RegisterFile("topodata.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1056 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0x66, 0x7f, 0xec, 0xd8, 0x67, 0x5d, 0x77, 0x3b, 0xa4, 0x68, 0xb5, 0x08, 0x11, 0xf9, 0x86,
	0xa8, 0x08, 0x83, 0x5c, 0x0a, 0x51, 0x25, 0xa4, 0x38, 0x66, 0x0b, 0xf9, 0x73, 0xcc, 0xd8, 0x51,
	0xc9, 0xd5, 0x6a, 0x6d, 0x4f, 0xd3, 0x55, 0x6c, 0xef, 0x32, 0x33, 0x8e, 0xe4, 0x67, 0xe8, 0x45,
	0xb9, 0xe6, 0x65, 0xb8, 0xe4, 0xa9, 0x90, 0xd0, 0x9c, 0xd9, 0xb5, 0xd7, 0x0e, 0x09, 0x29, 0xca,
	0x55, 0xe6, 0xec, 0xf9, 0x99, 0xf3, 0x7d, 0xe7, 0x3b, 0x13, 0x43, 0x5d, 0x26, 0x69, 0x32, 0x8e,
	0x64, 0xd4, 0x4c, 0x79, 0x22, 0x13, 0x52, 0xc9, 0xed, 0x46, 0x0b, 0x2a, 0xc7, 0x6c, 0x41, 0xa3,
	0xd9, 0x25, 0x23, 0xdb, 0x50, 0x12, 0x32, 0xe2, 0xd2, 0x33, 0x76, 0x8c, 0xdd, 0x1a, 0xd5, 0x06,
	0x71, 0xc1, 0x62, 0xb3, 0xb1, 0x67, 0xe2, 0x37, 0x75, 0x6c, 0x3c, 0x07, 0x67, 0x10, 0x0d, 0x27,
	0x4c, 0xb6, 0x27, 0x71, 0x24, 0x08, 0x01, 0x7b, 0xc4, 0x26, 0x13, 0xcc, 0xaa, 0x52, 0x3c, 0xab,
	0xa4, 0x79, 0xac, 0x93, 0x1e, 0x51, 0x75, 0x6c, 0xfc, 0x6d, 0x41, 0x59, 0x67, 0x91, 0x2f, 0xa1,
	0x14, 0xa9, 0x4c, 0xcc, 0x70, 0x5a, 0x4f, 0x9b, 0xcb, 0xee, 0x0a, 0x65, 0xa9, 0x8e, 0x21, 0x3e,
	0x54, 0xde, 0x26, 0x42, 0xce, 0xa2, 0x29, 0xc3, 0x72, 0x55, 0xba, 0xb4, 0x49, 0x1d, 0xcc, 0x38,
	0xf5, 0x2c, 0xfc, 0x6a, 0xc6, 0x29, 0xd9, 0x83, 0x4a, 0x9a, 0x70, 0x19, 0x4e, 0xa3, 0xd4, 0xb3,
	0x77, 0xac, 0x5d, 0xa7, 0xf5, 0xd9, 0x66, 0xed, 0x66, 0x2f, 0xe1, 0xf2, 0x34, 0x4a, 0x83, 0x99,
	0xe4, 0x0b, 0xba, 0x95, 0x6a, 0x4b, 0xdd, 0x72, 0xc5, 0x16, 0x22, 0x8d, 0x46, 0xcc, 0x2b, 0xe9,
	0x5b, 0x72, 0x1b, 0x69, 0x79, 0x1b, 0xf1, 0xb1, 0x57, 0x46, 0x87, 0x36, 0xc8, 0xd7, 0x50, 0xbd,
	0x62, 0x8b, 0x90, 0x2b, 0xe6, 0xbc, 0x2d, 0x04, 0x42, 0x56, 0x97, 0xe5, 0x9c, 0x62, 0x19, 0xcd,
	0xee, 0x2e, 0xd8, 0x72, 0x91, 0x32, 0xaf, 0xb2, 0x63, 0xec, 0xd6, 0x5b, 0xdb, 0x9b, 0x8d, 0x0d,
	0x16, 0x29, 0xa3, 0x18, 0x41, 0x76, 0xc1, 0x1d, 0x0f, 0x43, 0x85, 0x30, 0x4c, 0xae, 0x19, 0xe7,
	0xf1, 0x98, 0x79, 0x55, 0xbc, 0xbb, 0x3e, 0x1e, 0x76, 0xa3, 0x29, 0x3b, 0xcb, 0xbe, 0x92, 0x26,
	0xd8, 0x32, 0xba, 0x14, 0x1e, 0x20, 0x58, 0xff, 0x06, 0xd8, 0x41, 0x74, 0x29, 0x34, 0x52, 0x8c,
	0xf3, 0x5f, 0x42, 0xad, 0x88, 0x5f, 0x8d, 0xe9, 0x8a, 0x2d, 0xb2, 0xc9, 0xa9, 0xa3, 0x02, 0x7b,
	0x1d, 0x4d, 0xe6, 0x9a, 0xeb, 0x12, 0xd5, 0xc6, 0x4b, 0x73, 0xcf, 0xf0, 0xbf, 0x87, 0xea, 0xb2,
	0xdc, 0x7f, 0x25, 0x56, 0x0b, 0x89, 0x47, 0x76, 0xc5, 0x71, 0x6b, 0x8d, 0x77, 0x65, 0x28, 0xf5,
	0x91, 0xb9, 0x3d, 0xa8, 0x4d, 0x23, 0x21, 0x19, 0x0f, 0xef, 0xa1, 0x02, 0x47, 0x87, 0x6a, 0xa5,
	0xad, 0x71, 0x6e, 0xde, 0x83, 0xf3, 0x1f, 0xa0, 0x26, 0x18, 0xbf, 0x66, 0xe3, 0x50, 0x11, 0x2b,
	0x3c, 0x6b, 0x93, 0x27, 0xec, 0xa8, 0xd9, 0xc7, 0x18, 0x9c, 0x80, 0x23, 0x96, 0x67, 0x41, 0xf6,
	0xe1, 0x91, 0x48, 0xe6, 0x7c, 0xc4, 0x42, 0x9c, 0xb9, 0xc8, 0x44, 0xf5, 0xe9, 0x8d, 0x7c, 0x0c,
	0xc2, 0x33, 0xad, 0x89, 0x95, 0x21, 0x14, 0x2b, 0x6a, 0x1f, 0x84, 0x57, 0xda, 0xb1, 0x14, 0x2b,
	0x68, 0x90, 0x57, 0xf0, 0x58, 0x22, 0xc6, 0x70, 0x94, 0xcc, 0x24, 0x4f, 0x26, 0xc2, 0x2b, 0x6f,
	0xca, 0x55, 0x57, 0xd6, 0x54, 0x74, 0x74, 0x14, 0xad, 0xcb, 0xa2, 0x29, 0xfc, 0x0b, 0x80, 0x55,
	0xeb, 0xe4, 0x05, 0x38, 0x59, 0x55, 0xd4, 0x99, 0x71, 0x87, 0xce, 0x40, 0x2e, 0xcf, 0xab, 0x16,
	0xcd, 0x42, 0x8b, 0xfe, 0x1f, 0x06, 0x38, 0x05, 0x58, 0xf9, 0x42, 0x1b, 0xcb, 0x85, 0x5e, 0x5b,
	0x19, 0xf3, 0xb6, 0x95, 0xb1, 0x6e, 0x5d, 0x19, 0xfb, 0x1e, 0xe3, 0xfb, 0x04, 0xca, 0xd8, 0x68,
	0x4e, 0x5f, 0x66, 0xf9, 0x7f, 0x1a, 0xf0, 0x68, 0x8d, 0x99, 0x07, 0xc5, 0x4e, 0x5a, 0xf0, 0x74,
	0x1c, 0x0b, 0x15, 0x15, 0xfe, 0x36, 0x67, 0x7c, 0x11, 0x2a, 0x4d, 0xc4, 0x23, 0x86, 0x68, 0x2a,
	0xf4, 0xe3, 0xcc, 0xf9, 0x8b, 0xf2, 0xf5, 0xb5, 0x8b, 0x7c, 0x05, 0x64, 0x38, 0x89, 0x46, 0x57,
	0x93, 0x58, 0x48, 0x25, 0x37, 0xdd, 0xb6, 0x8d, 0x65, 0x9f, 0x14, 0x3c, 0xd8, 0x88, 0x68, 0xfc,
	0x65, 0xe2, 0xbb, 0xab, 0xd9, 0xfa, 0x06, 0xb6, 0x91, 0xa0, 0x78, 0x76, 0x19, 0x8e, 0x92, 0xc9,
	0x7c, 0x3a, 0xc3, 0xe5, 0xcf, 0xb6, 0x8b, 0xe4, 0xbe, 0x0e, 0xba, 0xd4, 0xfe, 0x93, 0xa3, 0x9b,
	0x19, 0x88, 0xdb, 0x44, 0xdc, 0xde, 0x1a, 0xa9, 0x78, 0xc7, 0xa1, 0x56, 0xf7, 0x46, 0x2d, 0xe4,
	0x60, 0x7f, 0xb9, 0x23, 0x6f, 0x78, 0x32, 0x15, 0x37, 0x1f, 0xce, 0xbc, 0x46, 0xb6, 0x26, 0xaf,
	0x78, 0x32, 0xcd, 0xd7, 0x44, 0x9d, 0x85, 0x3f, 0xcf, 0x65, 0xa8, 0xcc, 0x87, 0x1d, 0x45, 0x51,
	0x64, 0xd6, 0xba, 0xc8, 0x8e, 0xec, 0x8a, 0xe5, 0xda, 0x8d, 0x77, 0x06, 0xb8, 0x7a, 0xf3, 0x58,
	0x3a, 0x89, 0x47, 0x91, 0x8c, 0x93, 0x19, 0x79, 0x01, 0xa5, 0x59, 0x32, 0x66, 0xea, 0x6d, 0x51,
	0x60, 0x3e, 0xdf, 0x58, 0xab, 0x42, 0x68, 0xb3, 0x9b, 0x8c, 0x19, 0xd5, 0xd1, 0xfe, 0x3e, 0xd8,
	0xca, 0x54, 0x2f, 0x54, 0x06, 0xe1, 0x3e, 0x2f, 0x94, 0x5c, 0x19, 0x8d, 0x73, 0xa8, 0x67, 0x37,
	0xbc, 0x61, 0x9c, 0xcd, 0x46, 0x4c, 0xfd, 0x77, 0x2c, 0x0c, 0x13, 0xcf, 0x1f, 0xfc, 0x8e, 0x35,
	0x7e, 0xb7, 0xc1, 0xe9, 0xf3, 0xeb, 0xa5, 0x62, 0x7e, 0x02, 0x48, 0x23, 0x2e, 0x63, 0x85, 0x20,
	0x07, 0xf9, 0x45, 0x01, 0xe4, 0x2a, 0x74, 0x39, 0xbd, 0x5e, 0x1e, 0x4f, 0x0b, 0xa9, 0xb7, 0x4a,
	0xcf, 0xfc, 0x60, 0xe9, 0x59, 0xff, 0x43, 0x7a, 0x6d, 0x70, 0x0a, 0xd2, 0xcb, 0x94, 0xb7, 0xf3,
	0xef, 0x38, 0x0a, 0xe2, 0x83, 0x95, 0xf8, 0xfc, 0xf7, 0x06, 0x3c, 0xb9, 0x01, 0x51, 0x69, 0xb0,
	0xf0, 0xee, 0xdf, 0xad, 0xc1, 0xd5, 0x83, 0x4f, 0x3a, 0xe0, 0x62, 0x97, 0x21, 0xcf, 0xc7, 0xa7,
	0xe5, 0xe8, 0x14, 0x71, 0xad, 0xcf, 0x97, 0x3e, 0x16, 0x6b, 0xb6, 0xf0, 0xc3, 0x87, 0xd8, 0x86,
	0x3b, 0x1e, 0xd7, 0x23, 0xbb, 0x52, 0x72, 0xcb, 0xcf, 0x5a, 0x50, 0x5f, 0x67, 0x98, 0x54, 0xa1,
	0x74, 0xde, 0xed, 0x07, 0x03, 0xf7, 0x23, 0x02, 0x50, 0x3e, 0x3f, 0xec, 0x0e, 0xbe, 0xfb, 0xd6,
	0x35, 0xd4, 0xe7, 0x83, 0x8b, 0x41, 0xd0, 0x77, 0xcd, 0x67, 0xef, 0x0d, 0x80, 0xd5, 0x85, 0xc4,
	0x81, 0xad, 0xf3, 0xee, 0x71, 0xf7, 0xec, 0x75, 0x57, 0xa7, 0x9c, 0xb6, 0xfb, 0x83, 0x80, 0xba,
	0x86, 0x72, 0xd0, 0xa0, 0x77, 0x72, 0xd8, 0x69, 0xbb, 0xa6, 0x72, 0xd0, 0x1f, 0xcf, 0xba, 0x27,
	0x17, 0xae, 0x85, 0xb5, 0xda, 0x83, 0xce, 0xcf, 0xfa, 0xd8, 0xef, 0xb5, 0x69, 0xe0, 0xda, 0xc4,
	0x85, 0x5a, 0xf0, 0x6b, 0x2f, 0xa0, 0x87, 0xa7, 0x41, 0x77, 0xd0, 0x3e, 0x71, 0x4b, 0x2a, 0xe7,
	0xa0, 0xdd, 0x39, 0x3e, 0xef, 0xb9, 0x65, 0x5d, 0xac, 0x3f, 0x38, 0xa3, 0x81, 0xbb, 0xa5, 0x1c,
	0xaf, 0xcf, 0xe8, 0x71, 0x40, 0xdd, 0x8a, 0x6f, 0xba, 0xc6, 0x81, 0x0f, 0xde, 0x28, 0x99, 0x36,
	0x17, 0xc9, 0x5c, 0xce, 0x87, 0xac, 0x79, 0x1d, 0x4b, 0x26, 0x84, 0xfe, 0x91, 0x3a, 0x2c, 0xe3,
	0x9f, 0xe7, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x8c, 0x39, 0x13, 0xbd, 0x0a, 0x00, 0x00,
}
