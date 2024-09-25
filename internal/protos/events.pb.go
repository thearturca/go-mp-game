// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: internal/protos/events.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Direction int32

const (
	Direction_left  Direction = 0
	Direction_right Direction = 1
	Direction_up    Direction = 2
	Direction_down  Direction = 3
)

// Enum value maps for Direction.
var (
	Direction_name = map[int32]string{
		0: "left",
		1: "right",
		2: "up",
		3: "down",
	}
	Direction_value = map[string]int32{
		"left":  0,
		"right": 1,
		"up":    2,
		"down":  3,
	}
)

func (x Direction) Enum() *Direction {
	p := new(Direction)
	*p = x
	return p
}

func (x Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_protos_events_proto_enumTypes[0].Descriptor()
}

func (Direction) Type() protoreflect.EnumType {
	return &file_internal_protos_events_proto_enumTypes[0]
}

func (x Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Direction.Descriptor instead.
func (Direction) EnumDescriptor() ([]byte, []int) {
	return file_internal_protos_events_proto_rawDescGZIP(), []int{0}
}

type Action int32

const (
	Action_idle Action = 0
	Action_run  Action = 1
)

// Enum value maps for Action.
var (
	Action_name = map[int32]string{
		0: "idle",
		1: "run",
	}
	Action_value = map[string]int32{
		"idle": 0,
		"run":  1,
	}
)

func (x Action) Enum() *Action {
	p := new(Action)
	*p = x
	return p
}

func (x Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_protos_events_proto_enumTypes[1].Descriptor()
}

func (Action) Type() protoreflect.EnumType {
	return &file_internal_protos_events_proto_enumTypes[1]
}

func (x Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action.Descriptor instead.
func (Action) EnumDescriptor() ([]byte, []int) {
	return file_internal_protos_events_proto_rawDescGZIP(), []int{1}
}

type Skin int32

const (
	Skin_big_demon  Skin = 0
	Skin_big_zombie Skin = 1
	Skin_elf_f      Skin = 2
)

// Enum value maps for Skin.
var (
	Skin_name = map[int32]string{
		0: "big_demon",
		1: "big_zombie",
		2: "elf_f",
	}
	Skin_value = map[string]int32{
		"big_demon":  0,
		"big_zombie": 1,
		"elf_f":      2,
	}
)

func (x Skin) Enum() *Skin {
	p := new(Skin)
	*p = x
	return p
}

func (x Skin) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Skin) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_protos_events_proto_enumTypes[2].Descriptor()
}

func (Skin) Type() protoreflect.EnumType {
	return &file_internal_protos_events_proto_enumTypes[2]
}

func (x Skin) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Skin.Descriptor instead.
func (Skin) EnumDescriptor() ([]byte, []int) {
	return file_internal_protos_events_proto_rawDescGZIP(), []int{2}
}

type EventType int32

const (
	EventType_init        EventType = 0
	EventType_connect     EventType = 1
	EventType_exit        EventType = 2
	EventType_stop        EventType = 3
	EventType_move        EventType = 4
	EventType_empty       EventType = 5
	EventType_state       EventType = 6
	EventType_state_patch EventType = 7
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "init",
		1: "connect",
		2: "exit",
		3: "stop",
		4: "move",
		5: "empty",
		6: "state",
		7: "state_patch",
	}
	EventType_value = map[string]int32{
		"init":        0,
		"connect":     1,
		"exit":        2,
		"stop":        3,
		"move":        4,
		"empty":       5,
		"state":       6,
		"state_patch": 7,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_protos_events_proto_enumTypes[3].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_internal_protos_events_proto_enumTypes[3]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_internal_protos_events_proto_rawDescGZIP(), []int{3}
}

type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protos_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protos_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_internal_protos_events_proto_rawDescGZIP(), []int{0}
}

func (x *Position) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Position) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

type Velocity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Speed     float64   `protobuf:"fixed64,1,opt,name=speed,proto3" json:"speed,omitempty"`
	Direction Direction `protobuf:"varint,2,opt,name=direction,proto3,enum=protos.Direction" json:"direction,omitempty"`
}

func (x *Velocity) Reset() {
	*x = Velocity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protos_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Velocity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Velocity) ProtoMessage() {}

func (x *Velocity) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protos_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Velocity.ProtoReflect.Descriptor instead.
func (*Velocity) Descriptor() ([]byte, []int) {
	return file_internal_protos_events_proto_rawDescGZIP(), []int{1}
}

func (x *Velocity) GetSpeed() float64 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *Velocity) GetDirection() Direction {
	if x != nil {
		return x.Direction
	}
	return Direction_left
}

type Unit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Position *Position `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	Frame    int32     `protobuf:"varint,3,opt,name=frame,proto3" json:"frame,omitempty"`
	Skin     Skin      `protobuf:"varint,4,opt,name=skin,proto3,enum=protos.Skin" json:"skin,omitempty"`
	Action   Action    `protobuf:"varint,5,opt,name=action,proto3,enum=protos.Action" json:"action,omitempty"`
	Velocity *Velocity `protobuf:"bytes,6,opt,name=velocity,proto3" json:"velocity,omitempty"`
	Side     Direction `protobuf:"varint,7,opt,name=side,proto3,enum=protos.Direction" json:"side,omitempty"`
}

func (x *Unit) Reset() {
	*x = Unit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protos_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unit) ProtoMessage() {}

func (x *Unit) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protos_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unit.ProtoReflect.Descriptor instead.
func (*Unit) Descriptor() ([]byte, []int) {
	return file_internal_protos_events_proto_rawDescGZIP(), []int{2}
}

func (x *Unit) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Unit) GetPosition() *Position {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *Unit) GetFrame() int32 {
	if x != nil {
		return x.Frame
	}
	return 0
}

func (x *Unit) GetSkin() Skin {
	if x != nil {
		return x.Skin
	}
	return Skin_big_demon
}

func (x *Unit) GetAction() Action {
	if x != nil {
		return x.Action
	}
	return Action_idle
}

func (x *Unit) GetVelocity() *Velocity {
	if x != nil {
		return x.Velocity
	}
	return nil
}

func (x *Unit) GetSide() Direction {
	if x != nil {
		return x.Side
	}
	return Direction_left
}

type PatchUnit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Position *Position  `protobuf:"bytes,2,opt,name=position,proto3,oneof" json:"position,omitempty"`
	Action   *Action    `protobuf:"varint,5,opt,name=action,proto3,enum=protos.Action,oneof" json:"action,omitempty"`
	Velocity *Velocity  `protobuf:"bytes,6,opt,name=velocity,proto3,oneof" json:"velocity,omitempty"`
	Side     *Direction `protobuf:"varint,7,opt,name=side,proto3,enum=protos.Direction,oneof" json:"side,omitempty"`
}

func (x *PatchUnit) Reset() {
	*x = PatchUnit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protos_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchUnit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchUnit) ProtoMessage() {}

func (x *PatchUnit) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protos_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchUnit.ProtoReflect.Descriptor instead.
func (*PatchUnit) Descriptor() ([]byte, []int) {
	return file_internal_protos_events_proto_rawDescGZIP(), []int{3}
}

func (x *PatchUnit) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PatchUnit) GetPosition() *Position {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *PatchUnit) GetAction() Action {
	if x != nil && x.Action != nil {
		return *x.Action
	}
	return Action_idle
}

func (x *PatchUnit) GetVelocity() *Velocity {
	if x != nil {
		return x.Velocity
	}
	return nil
}

func (x *PatchUnit) GetSide() Direction {
	if x != nil && x.Side != nil {
		return *x.Side
	}
	return Direction_left
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type EventType `protobuf:"varint,1,opt,name=type,proto3,enum=protos.EventType" json:"type,omitempty"`
	// Types that are assignable to Data:
	//
	//	*Event_Connect
	//	*Event_Move
	//	*Event_State
	//	*Event_StatePatch
	Data     isEvent_Data `protobuf_oneof:"data"`
	PlayerId uint32       `protobuf:"varint,9,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protos_events_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protos_events_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_internal_protos_events_proto_rawDescGZIP(), []int{4}
}

func (x *Event) GetType() EventType {
	if x != nil {
		return x.Type
	}
	return EventType_init
}

func (m *Event) GetData() isEvent_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *Event) GetConnect() *EventConnect {
	if x, ok := x.GetData().(*Event_Connect); ok {
		return x.Connect
	}
	return nil
}

func (x *Event) GetMove() *EventMove {
	if x, ok := x.GetData().(*Event_Move); ok {
		return x.Move
	}
	return nil
}

func (x *Event) GetState() *GameState {
	if x, ok := x.GetData().(*Event_State); ok {
		return x.State
	}
	return nil
}

func (x *Event) GetStatePatch() *GameStatePatche {
	if x, ok := x.GetData().(*Event_StatePatch); ok {
		return x.StatePatch
	}
	return nil
}

func (x *Event) GetPlayerId() uint32 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

type isEvent_Data interface {
	isEvent_Data()
}

type Event_Connect struct {
	Connect *EventConnect `protobuf:"bytes,3,opt,name=connect,proto3,oneof"`
}

type Event_Move struct {
	Move *EventMove `protobuf:"bytes,6,opt,name=move,proto3,oneof"`
}

type Event_State struct {
	State *GameState `protobuf:"bytes,7,opt,name=state,proto3,oneof"`
}

type Event_StatePatch struct {
	StatePatch *GameStatePatche `protobuf:"bytes,8,opt,name=state_patch,json=statePatch,proto3,oneof"`
}

func (*Event_Connect) isEvent_Data() {}

func (*Event_Move) isEvent_Data() {}

func (*Event_State) isEvent_Data() {}

func (*Event_StatePatch) isEvent_Data() {}

type EventConnect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unit *Unit `protobuf:"bytes,1,opt,name=unit,proto3" json:"unit,omitempty"`
}

func (x *EventConnect) Reset() {
	*x = EventConnect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protos_events_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventConnect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventConnect) ProtoMessage() {}

func (x *EventConnect) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protos_events_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventConnect.ProtoReflect.Descriptor instead.
func (*EventConnect) Descriptor() ([]byte, []int) {
	return file_internal_protos_events_proto_rawDescGZIP(), []int{5}
}

func (x *EventConnect) GetUnit() *Unit {
	if x != nil {
		return x.Unit
	}
	return nil
}

type EventMove struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Direction Direction `protobuf:"varint,2,opt,name=direction,proto3,enum=protos.Direction" json:"direction,omitempty"`
}

func (x *EventMove) Reset() {
	*x = EventMove{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protos_events_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventMove) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventMove) ProtoMessage() {}

func (x *EventMove) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protos_events_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventMove.ProtoReflect.Descriptor instead.
func (*EventMove) Descriptor() ([]byte, []int) {
	return file_internal_protos_events_proto_rawDescGZIP(), []int{6}
}

func (x *EventMove) GetDirection() Direction {
	if x != nil {
		return x.Direction
	}
	return Direction_left
}

type GameState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Units map[uint32]*Unit `protobuf:"bytes,1,rep,name=units,proto3" json:"units,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GameState) Reset() {
	*x = GameState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protos_events_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameState) ProtoMessage() {}

func (x *GameState) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protos_events_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameState.ProtoReflect.Descriptor instead.
func (*GameState) Descriptor() ([]byte, []int) {
	return file_internal_protos_events_proto_rawDescGZIP(), []int{7}
}

func (x *GameState) GetUnits() map[uint32]*Unit {
	if x != nil {
		return x.Units
	}
	return nil
}

type GameStatePatche struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Units map[uint32]*PatchUnit `protobuf:"bytes,1,rep,name=units,proto3" json:"units,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GameStatePatche) Reset() {
	*x = GameStatePatche{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protos_events_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameStatePatche) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameStatePatche) ProtoMessage() {}

func (x *GameStatePatche) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protos_events_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameStatePatche.ProtoReflect.Descriptor instead.
func (*GameStatePatche) Descriptor() ([]byte, []int) {
	return file_internal_protos_events_proto_rawDescGZIP(), []int{8}
}

func (x *GameStatePatche) GetUnits() map[uint32]*PatchUnit {
	if x != nil {
		return x.Units
	}
	return nil
}

var File_internal_protos_events_proto protoreflect.FileDescriptor

var file_internal_protos_events_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x26, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x78,
	0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x79, 0x22, 0x51,
	0x0a, 0x08, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70,
	0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64,
	0x12, 0x2f, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0xf9, 0x01, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x04, 0x73, 0x6b, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x6b, 0x69, 0x6e, 0x52, 0x04, 0x73, 0x6b, 0x69, 0x6e,
	0x12, 0x26, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x08, 0x76, 0x65, 0x6c, 0x6f,
	0x63, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x52, 0x08, 0x76, 0x65,
	0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x12, 0x25, 0x0a, 0x04, 0x73, 0x69, 0x64, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x73, 0x69, 0x64, 0x65, 0x22, 0x88, 0x02,
	0x0a, 0x09, 0x50, 0x61, 0x74, 0x63, 0x68, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x00, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x2b,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x01,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x08, 0x76,
	0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x48,
	0x02, 0x52, 0x08, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x88, 0x01, 0x01, 0x12, 0x2a,
	0x0a, 0x04, 0x73, 0x69, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x03, 0x52, 0x04, 0x73, 0x69, 0x64, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x73, 0x69, 0x64, 0x65, 0x22, 0x95, 0x02, 0x0a, 0x05, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x6d,
	0x6f, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x76, 0x65, 0x48, 0x00, 0x52, 0x04,
	0x6d, 0x6f, 0x76, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x61, 0x6d,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x3a, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x61,
	0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x63, 0x68, 0x65, 0x48, 0x00, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x30, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x12, 0x20, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x04, 0x75, 0x6e,
	0x69, 0x74, 0x22, 0x3c, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x76, 0x65, 0x12,
	0x2f, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x87, 0x01, 0x0a, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x32,
	0x0a, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x75, 0x6e, 0x69,
	0x74, 0x73, 0x1a, 0x46, 0x0a, 0x0a, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x22, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x98, 0x01, 0x0a, 0x0f, 0x47,
	0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x63, 0x68, 0x65, 0x12, 0x38,
	0x0a, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x74, 0x63, 0x68, 0x65, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x1a, 0x4b, 0x0a, 0x0a, 0x55, 0x6e, 0x69, 0x74,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x32, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x08, 0x0a, 0x04, 0x6c, 0x65, 0x66, 0x74, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x72, 0x69, 0x67, 0x68, 0x74, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x75, 0x70, 0x10, 0x02, 0x12,
	0x08, 0x0a, 0x04, 0x64, 0x6f, 0x77, 0x6e, 0x10, 0x03, 0x2a, 0x1b, 0x0a, 0x06, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x08, 0x0a, 0x04, 0x69, 0x64, 0x6c, 0x65, 0x10, 0x00, 0x12, 0x07, 0x0a,
	0x03, 0x72, 0x75, 0x6e, 0x10, 0x01, 0x2a, 0x30, 0x0a, 0x04, 0x53, 0x6b, 0x69, 0x6e, 0x12, 0x0d,
	0x0a, 0x09, 0x62, 0x69, 0x67, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x6e, 0x10, 0x00, 0x12, 0x0e, 0x0a,
	0x0a, 0x62, 0x69, 0x67, 0x5f, 0x7a, 0x6f, 0x6d, 0x62, 0x69, 0x65, 0x10, 0x01, 0x12, 0x09, 0x0a,
	0x05, 0x65, 0x6c, 0x66, 0x5f, 0x66, 0x10, 0x02, 0x2a, 0x67, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x69, 0x6e, 0x69, 0x74, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04,
	0x65, 0x78, 0x69, 0x74, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x73, 0x74, 0x6f, 0x70, 0x10, 0x03,
	0x12, 0x08, 0x0a, 0x04, 0x6d, 0x6f, 0x76, 0x65, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x10, 0x06,
	0x12, 0x0f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x10,
	0x07, 0x42, 0x11, 0x5a, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_protos_events_proto_rawDescOnce sync.Once
	file_internal_protos_events_proto_rawDescData = file_internal_protos_events_proto_rawDesc
)

func file_internal_protos_events_proto_rawDescGZIP() []byte {
	file_internal_protos_events_proto_rawDescOnce.Do(func() {
		file_internal_protos_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_protos_events_proto_rawDescData)
	})
	return file_internal_protos_events_proto_rawDescData
}

var file_internal_protos_events_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_internal_protos_events_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_internal_protos_events_proto_goTypes = []any{
	(Direction)(0),          // 0: protos.Direction
	(Action)(0),             // 1: protos.Action
	(Skin)(0),               // 2: protos.Skin
	(EventType)(0),          // 3: protos.EventType
	(*Position)(nil),        // 4: protos.Position
	(*Velocity)(nil),        // 5: protos.Velocity
	(*Unit)(nil),            // 6: protos.Unit
	(*PatchUnit)(nil),       // 7: protos.PatchUnit
	(*Event)(nil),           // 8: protos.Event
	(*EventConnect)(nil),    // 9: protos.EventConnect
	(*EventMove)(nil),       // 10: protos.EventMove
	(*GameState)(nil),       // 11: protos.GameState
	(*GameStatePatche)(nil), // 12: protos.GameStatePatche
	nil,                     // 13: protos.GameState.UnitsEntry
	nil,                     // 14: protos.GameStatePatche.UnitsEntry
}
var file_internal_protos_events_proto_depIdxs = []int32{
	0,  // 0: protos.Velocity.direction:type_name -> protos.Direction
	4,  // 1: protos.Unit.position:type_name -> protos.Position
	2,  // 2: protos.Unit.skin:type_name -> protos.Skin
	1,  // 3: protos.Unit.action:type_name -> protos.Action
	5,  // 4: protos.Unit.velocity:type_name -> protos.Velocity
	0,  // 5: protos.Unit.side:type_name -> protos.Direction
	4,  // 6: protos.PatchUnit.position:type_name -> protos.Position
	1,  // 7: protos.PatchUnit.action:type_name -> protos.Action
	5,  // 8: protos.PatchUnit.velocity:type_name -> protos.Velocity
	0,  // 9: protos.PatchUnit.side:type_name -> protos.Direction
	3,  // 10: protos.Event.type:type_name -> protos.EventType
	9,  // 11: protos.Event.connect:type_name -> protos.EventConnect
	10, // 12: protos.Event.move:type_name -> protos.EventMove
	11, // 13: protos.Event.state:type_name -> protos.GameState
	12, // 14: protos.Event.state_patch:type_name -> protos.GameStatePatche
	6,  // 15: protos.EventConnect.unit:type_name -> protos.Unit
	0,  // 16: protos.EventMove.direction:type_name -> protos.Direction
	13, // 17: protos.GameState.units:type_name -> protos.GameState.UnitsEntry
	14, // 18: protos.GameStatePatche.units:type_name -> protos.GameStatePatche.UnitsEntry
	6,  // 19: protos.GameState.UnitsEntry.value:type_name -> protos.Unit
	7,  // 20: protos.GameStatePatche.UnitsEntry.value:type_name -> protos.PatchUnit
	21, // [21:21] is the sub-list for method output_type
	21, // [21:21] is the sub-list for method input_type
	21, // [21:21] is the sub-list for extension type_name
	21, // [21:21] is the sub-list for extension extendee
	0,  // [0:21] is the sub-list for field type_name
}

func init() { file_internal_protos_events_proto_init() }
func file_internal_protos_events_proto_init() {
	if File_internal_protos_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_protos_events_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Position); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_protos_events_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Velocity); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_protos_events_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Unit); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_protos_events_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*PatchUnit); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_protos_events_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_protos_events_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*EventConnect); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_protos_events_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*EventMove); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_protos_events_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GameState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_protos_events_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GameStatePatche); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_internal_protos_events_proto_msgTypes[3].OneofWrappers = []any{}
	file_internal_protos_events_proto_msgTypes[4].OneofWrappers = []any{
		(*Event_Connect)(nil),
		(*Event_Move)(nil),
		(*Event_State)(nil),
		(*Event_StatePatch)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_protos_events_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_protos_events_proto_goTypes,
		DependencyIndexes: file_internal_protos_events_proto_depIdxs,
		EnumInfos:         file_internal_protos_events_proto_enumTypes,
		MessageInfos:      file_internal_protos_events_proto_msgTypes,
	}.Build()
	File_internal_protos_events_proto = out.File
	file_internal_protos_events_proto_rawDesc = nil
	file_internal_protos_events_proto_goTypes = nil
	file_internal_protos_events_proto_depIdxs = nil
}
