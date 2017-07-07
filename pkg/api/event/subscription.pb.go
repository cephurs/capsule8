// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event/v0/subscription.proto

package event

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ThrottleModifier_IntervalType int32

const (
	ThrottleModifier_MILLISECOND ThrottleModifier_IntervalType = 0
	ThrottleModifier_SECOND      ThrottleModifier_IntervalType = 1
	ThrottleModifier_MINUTE      ThrottleModifier_IntervalType = 2
	ThrottleModifier_HOUR        ThrottleModifier_IntervalType = 3
)

var ThrottleModifier_IntervalType_name = map[int32]string{
	0: "MILLISECOND",
	1: "SECOND",
	2: "MINUTE",
	3: "HOUR",
}
var ThrottleModifier_IntervalType_value = map[string]int32{
	"MILLISECOND": 0,
	"SECOND":      1,
	"MINUTE":      2,
	"HOUR":        3,
}

func (x ThrottleModifier_IntervalType) String() string {
	return proto.EnumName(ThrottleModifier_IntervalType_name, int32(x))
}
func (ThrottleModifier_IntervalType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{12, 0}
}

//
// SignedSubscription wraps a subscription and signature.
//
type SignedSubscription struct {
	Subscription   *Subscription `protobuf:"bytes,1,opt,name=subscription" json:"subscription,omitempty"`
	SubscriptionId string        `protobuf:"bytes,2,opt,name=subscription_id,json=subscriptionId" json:"subscription_id,omitempty"`
	Signature      string        `protobuf:"bytes,3,opt,name=signature" json:"signature,omitempty"`
}

func (m *SignedSubscription) Reset()                    { *m = SignedSubscription{} }
func (m *SignedSubscription) String() string            { return proto.CompactTextString(m) }
func (*SignedSubscription) ProtoMessage()               {}
func (*SignedSubscription) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *SignedSubscription) GetSubscription() *Subscription {
	if m != nil {
		return m.Subscription
	}
	return nil
}

func (m *SignedSubscription) GetSubscriptionId() string {
	if m != nil {
		return m.SubscriptionId
	}
	return ""
}

func (m *SignedSubscription) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

//
// Subscription identifies a subscriber's interest in events. The optional
// Filter field restricts the subscription to the nodes and processes
// specified within. The optional Selector field specifies the types and
// properties of the events desired by the subscriber.
//
type Subscription struct {
	// Return events matching one or more of the specified event
	// filters. If no event filters are specified, then no events
	// will be returned.
	EventFilter *EventFilter `protobuf:"bytes,1,opt,name=event_filter,json=eventFilter" json:"event_filter,omitempty"`
	// If not empty, then only return events from containers matched
	// by one or more of the specified container filters.
	ContainerFilter *ContainerFilter `protobuf:"bytes,2,opt,name=container_filter,json=containerFilter" json:"container_filter,omitempty"`
	// If not empty, then only return events from nodes matched by
	// one or more of the specified node filters.
	NodeFilter *NodeFilter `protobuf:"bytes,3,opt,name=node_filter,json=nodeFilter" json:"node_filter,omitempty"`
	// If not empty, then only return events that occurred after
	// the specified relative duration subtracted from the current
	// time. If the resulting time is in the past, then the
	// subscription will search for historic events before streaming
	// live ones.
	SinceDuration int64 `protobuf:"varint,10,opt,name=since_duration,json=sinceDuration" json:"since_duration,omitempty"`
	// If not empty, then only return events that occurred before
	// the specified relative duration added to the current time. If
	// the resulting time is in the past, then the subscription will
	// search for historic events before streaming live ones. If the
	// resulting time is in the future, then the subscription will be
	// automatically closed at that time.
	UntilDuration int64 `protobuf:"varint,11,opt,name=until_duration,json=untilDuration" json:"until_duration,omitempty"`
	// If not empty, apply the specified modifier to the subscription.
	Modifier *Modifier `protobuf:"bytes,20,opt,name=modifier" json:"modifier,omitempty"`
}

func (m *Subscription) Reset()                    { *m = Subscription{} }
func (m *Subscription) String() string            { return proto.CompactTextString(m) }
func (*Subscription) ProtoMessage()               {}
func (*Subscription) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Subscription) GetEventFilter() *EventFilter {
	if m != nil {
		return m.EventFilter
	}
	return nil
}

func (m *Subscription) GetContainerFilter() *ContainerFilter {
	if m != nil {
		return m.ContainerFilter
	}
	return nil
}

func (m *Subscription) GetNodeFilter() *NodeFilter {
	if m != nil {
		return m.NodeFilter
	}
	return nil
}

func (m *Subscription) GetSinceDuration() int64 {
	if m != nil {
		return m.SinceDuration
	}
	return 0
}

func (m *Subscription) GetUntilDuration() int64 {
	if m != nil {
		return m.UntilDuration
	}
	return 0
}

func (m *Subscription) GetModifier() *Modifier {
	if m != nil {
		return m.Modifier
	}
	return nil
}

type NodeFilter struct {
}

func (m *NodeFilter) Reset()                    { *m = NodeFilter{} }
func (m *NodeFilter) String() string            { return proto.CompactTextString(m) }
func (*NodeFilter) ProtoMessage()               {}
func (*NodeFilter) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type ContainerFilter struct {
	Ids      []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
	Names    []string `protobuf:"bytes,2,rep,name=names" json:"names,omitempty"`
	ImageIds []string `protobuf:"bytes,3,rep,name=image_ids,json=imageIds" json:"image_ids,omitempty"`
	// Container image name (shell-style globs are supported)
	ImageNames []string `protobuf:"bytes,4,rep,name=image_names,json=imageNames" json:"image_names,omitempty"`
}

func (m *ContainerFilter) Reset()                    { *m = ContainerFilter{} }
func (m *ContainerFilter) String() string            { return proto.CompactTextString(m) }
func (*ContainerFilter) ProtoMessage()               {}
func (*ContainerFilter) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *ContainerFilter) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *ContainerFilter) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *ContainerFilter) GetImageIds() []string {
	if m != nil {
		return m.ImageIds
	}
	return nil
}

func (m *ContainerFilter) GetImageNames() []string {
	if m != nil {
		return m.ImageNames
	}
	return nil
}

type EventFilter struct {
	//
	// Kernel-level events
	//
	SyscallEvents []*SyscallEventFilter `protobuf:"bytes,1,rep,name=syscall_events,json=syscallEvents" json:"syscall_events,omitempty"`
	ProcessEvents []*ProcessEventFilter `protobuf:"bytes,2,rep,name=process_events,json=processEvents" json:"process_events,omitempty"`
	FileEvents    []*FileEventFilter    `protobuf:"bytes,3,rep,name=file_events,json=fileEvents" json:"file_events,omitempty"`
	//
	// Operating System-level events (containers, etc)
	//
	ContainerEvents []*ContainerEventFilter `protobuf:"bytes,10,rep,name=container_events,json=containerEvents" json:"container_events,omitempty"`
	//
	// Debugging events (>= 100)
	//
	ChargenEvents []*ChargenEventFilter `protobuf:"bytes,100,rep,name=chargen_events,json=chargenEvents" json:"chargen_events,omitempty"`
	TickerEvents  []*TickerEventFilter  `protobuf:"bytes,101,rep,name=ticker_events,json=tickerEvents" json:"ticker_events,omitempty"`
}

func (m *EventFilter) Reset()                    { *m = EventFilter{} }
func (m *EventFilter) String() string            { return proto.CompactTextString(m) }
func (*EventFilter) ProtoMessage()               {}
func (*EventFilter) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *EventFilter) GetSyscallEvents() []*SyscallEventFilter {
	if m != nil {
		return m.SyscallEvents
	}
	return nil
}

func (m *EventFilter) GetProcessEvents() []*ProcessEventFilter {
	if m != nil {
		return m.ProcessEvents
	}
	return nil
}

func (m *EventFilter) GetFileEvents() []*FileEventFilter {
	if m != nil {
		return m.FileEvents
	}
	return nil
}

func (m *EventFilter) GetContainerEvents() []*ContainerEventFilter {
	if m != nil {
		return m.ContainerEvents
	}
	return nil
}

func (m *EventFilter) GetChargenEvents() []*ChargenEventFilter {
	if m != nil {
		return m.ChargenEvents
	}
	return nil
}

func (m *EventFilter) GetTickerEvents() []*TickerEventFilter {
	if m != nil {
		return m.TickerEvents
	}
	return nil
}

type SyscallEventFilter struct {
	// Type of system call event (entry w/ args or exit w/ ret)
	Type SyscallEventType `protobuf:"varint,1,opt,name=type,enum=capsule8.event.v0.SyscallEventType" json:"type,omitempty"`
	// System call number (arch/x86/entry/syscalls/syscall_64.tbl)
	Id   *google_protobuf.Int64Value  `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Arg0 *google_protobuf.UInt64Value `protobuf:"bytes,10,opt,name=arg0" json:"arg0,omitempty"`
	Arg1 *google_protobuf.UInt64Value `protobuf:"bytes,11,opt,name=arg1" json:"arg1,omitempty"`
	Arg2 *google_protobuf.UInt64Value `protobuf:"bytes,12,opt,name=arg2" json:"arg2,omitempty"`
	Arg3 *google_protobuf.UInt64Value `protobuf:"bytes,13,opt,name=arg3" json:"arg3,omitempty"`
	Arg4 *google_protobuf.UInt64Value `protobuf:"bytes,14,opt,name=arg4" json:"arg4,omitempty"`
	Arg5 *google_protobuf.UInt64Value `protobuf:"bytes,15,opt,name=arg5" json:"arg5,omitempty"`
	Ret  *google_protobuf.Int64Value  `protobuf:"bytes,20,opt,name=ret" json:"ret,omitempty"`
}

func (m *SyscallEventFilter) Reset()                    { *m = SyscallEventFilter{} }
func (m *SyscallEventFilter) String() string            { return proto.CompactTextString(m) }
func (*SyscallEventFilter) ProtoMessage()               {}
func (*SyscallEventFilter) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *SyscallEventFilter) GetType() SyscallEventType {
	if m != nil {
		return m.Type
	}
	return SyscallEventType_SYSCALL_EVENT_TYPE_UNKNOWN
}

func (m *SyscallEventFilter) GetId() *google_protobuf.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SyscallEventFilter) GetArg0() *google_protobuf.UInt64Value {
	if m != nil {
		return m.Arg0
	}
	return nil
}

func (m *SyscallEventFilter) GetArg1() *google_protobuf.UInt64Value {
	if m != nil {
		return m.Arg1
	}
	return nil
}

func (m *SyscallEventFilter) GetArg2() *google_protobuf.UInt64Value {
	if m != nil {
		return m.Arg2
	}
	return nil
}

func (m *SyscallEventFilter) GetArg3() *google_protobuf.UInt64Value {
	if m != nil {
		return m.Arg3
	}
	return nil
}

func (m *SyscallEventFilter) GetArg4() *google_protobuf.UInt64Value {
	if m != nil {
		return m.Arg4
	}
	return nil
}

func (m *SyscallEventFilter) GetArg5() *google_protobuf.UInt64Value {
	if m != nil {
		return m.Arg5
	}
	return nil
}

func (m *SyscallEventFilter) GetRet() *google_protobuf.Int64Value {
	if m != nil {
		return m.Ret
	}
	return nil
}

type ProcessEventFilter struct {
	// Required; the process event type to match
	Type ProcessEventType `protobuf:"varint,1,opt,name=type,enum=capsule8.event.v0.ProcessEventType" json:"type,omitempty"`
}

func (m *ProcessEventFilter) Reset()                    { *m = ProcessEventFilter{} }
func (m *ProcessEventFilter) String() string            { return proto.CompactTextString(m) }
func (*ProcessEventFilter) ProtoMessage()               {}
func (*ProcessEventFilter) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *ProcessEventFilter) GetType() ProcessEventType {
	if m != nil {
		return m.Type
	}
	return ProcessEventType_PROCESS_EVENT_TYPE_UNKNOWN
}

type FileEventFilter struct {
	// Required
	Type FileEventType `protobuf:"varint,1,opt,name=type,enum=capsule8.event.v0.FileEventType" json:"type,omitempty"`
	// Optional filename exact match
	Filename *google_protobuf.StringValue `protobuf:"bytes,10,opt,name=filename" json:"filename,omitempty"`
	// Optional filename pattern
	FilenamePattern *google_protobuf.StringValue `protobuf:"bytes,11,opt,name=filename_pattern,json=filenamePattern" json:"filename_pattern,omitempty"`
	// Optional open(2) flags mask value
	OpenFlagsMask *google_protobuf.Int32Value `protobuf:"bytes,12,opt,name=open_flags_mask,json=openFlagsMask" json:"open_flags_mask,omitempty"`
	// Optional open(2)/creat(2) mode mask value
	CreateModeMask *google_protobuf.Int32Value `protobuf:"bytes,13,opt,name=create_mode_mask,json=createModeMask" json:"create_mode_mask,omitempty"`
}

func (m *FileEventFilter) Reset()                    { *m = FileEventFilter{} }
func (m *FileEventFilter) String() string            { return proto.CompactTextString(m) }
func (*FileEventFilter) ProtoMessage()               {}
func (*FileEventFilter) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *FileEventFilter) GetType() FileEventType {
	if m != nil {
		return m.Type
	}
	return FileEventType_FILE_EVENT_TYPE_UNKNOWN
}

func (m *FileEventFilter) GetFilename() *google_protobuf.StringValue {
	if m != nil {
		return m.Filename
	}
	return nil
}

func (m *FileEventFilter) GetFilenamePattern() *google_protobuf.StringValue {
	if m != nil {
		return m.FilenamePattern
	}
	return nil
}

func (m *FileEventFilter) GetOpenFlagsMask() *google_protobuf.Int32Value {
	if m != nil {
		return m.OpenFlagsMask
	}
	return nil
}

func (m *FileEventFilter) GetCreateModeMask() *google_protobuf.Int32Value {
	if m != nil {
		return m.CreateModeMask
	}
	return nil
}

type ContainerEventFilter struct {
	// Required.
	Type ContainerEventType `protobuf:"varint,1,opt,name=type,enum=capsule8.event.v0.ContainerEventType" json:"type,omitempty"`
}

func (m *ContainerEventFilter) Reset()                    { *m = ContainerEventFilter{} }
func (m *ContainerEventFilter) String() string            { return proto.CompactTextString(m) }
func (*ContainerEventFilter) ProtoMessage()               {}
func (*ContainerEventFilter) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *ContainerEventFilter) GetType() ContainerEventType {
	if m != nil {
		return m.Type
	}
	return ContainerEventType_CONTAINER_EVENT_TYPE_UNKNOWN
}

type ChargenEventFilter struct {
	// Length of character sequence strings to generate
	Length uint64 `protobuf:"varint,1,opt,name=length" json:"length,omitempty"`
}

func (m *ChargenEventFilter) Reset()                    { *m = ChargenEventFilter{} }
func (m *ChargenEventFilter) String() string            { return proto.CompactTextString(m) }
func (*ChargenEventFilter) ProtoMessage()               {}
func (*ChargenEventFilter) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *ChargenEventFilter) GetLength() uint64 {
	if m != nil {
		return m.Length
	}
	return 0
}

type TickerEventFilter struct {
	// The interval at which ticker events are generated
	Interval int64 `protobuf:"varint,1,opt,name=interval" json:"interval,omitempty"`
}

func (m *TickerEventFilter) Reset()                    { *m = TickerEventFilter{} }
func (m *TickerEventFilter) String() string            { return proto.CompactTextString(m) }
func (*TickerEventFilter) ProtoMessage()               {}
func (*TickerEventFilter) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *TickerEventFilter) GetInterval() int64 {
	if m != nil {
		return m.Interval
	}
	return 0
}

//
// Modifier specifies which stream modifiers to apply if any. For a given
// stream, a modifier can apply a throttle or limit etc. Modifiers can be
// used together.
//
type Modifier struct {
	Throttle *ThrottleModifier `protobuf:"bytes,1,opt,name=throttle" json:"throttle,omitempty"`
	Limit    *LimitModifier    `protobuf:"bytes,2,opt,name=limit" json:"limit,omitempty"`
}

func (m *Modifier) Reset()                    { *m = Modifier{} }
func (m *Modifier) String() string            { return proto.CompactTextString(m) }
func (*Modifier) ProtoMessage()               {}
func (*Modifier) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

func (m *Modifier) GetThrottle() *ThrottleModifier {
	if m != nil {
		return m.Throttle
	}
	return nil
}

func (m *Modifier) GetLimit() *LimitModifier {
	if m != nil {
		return m.Limit
	}
	return nil
}

type ThrottleModifier struct {
	Interval     int64                         `protobuf:"varint,1,opt,name=interval" json:"interval,omitempty"`
	IntervalType ThrottleModifier_IntervalType `protobuf:"varint,2,opt,name=interval_type,json=intervalType,enum=capsule8.event.v0.ThrottleModifier_IntervalType" json:"interval_type,omitempty"`
}

func (m *ThrottleModifier) Reset()                    { *m = ThrottleModifier{} }
func (m *ThrottleModifier) String() string            { return proto.CompactTextString(m) }
func (*ThrottleModifier) ProtoMessage()               {}
func (*ThrottleModifier) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

func (m *ThrottleModifier) GetInterval() int64 {
	if m != nil {
		return m.Interval
	}
	return 0
}

func (m *ThrottleModifier) GetIntervalType() ThrottleModifier_IntervalType {
	if m != nil {
		return m.IntervalType
	}
	return ThrottleModifier_MILLISECOND
}

type LimitModifier struct {
	// Limit the number of events
	Limit int64 `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
}

func (m *LimitModifier) Reset()                    { *m = LimitModifier{} }
func (m *LimitModifier) String() string            { return proto.CompactTextString(m) }
func (*LimitModifier) ProtoMessage()               {}
func (*LimitModifier) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{13} }

func (m *LimitModifier) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func init() {
	proto.RegisterType((*SignedSubscription)(nil), "capsule8.event.v0.SignedSubscription")
	proto.RegisterType((*Subscription)(nil), "capsule8.event.v0.Subscription")
	proto.RegisterType((*NodeFilter)(nil), "capsule8.event.v0.NodeFilter")
	proto.RegisterType((*ContainerFilter)(nil), "capsule8.event.v0.ContainerFilter")
	proto.RegisterType((*EventFilter)(nil), "capsule8.event.v0.EventFilter")
	proto.RegisterType((*SyscallEventFilter)(nil), "capsule8.event.v0.SyscallEventFilter")
	proto.RegisterType((*ProcessEventFilter)(nil), "capsule8.event.v0.ProcessEventFilter")
	proto.RegisterType((*FileEventFilter)(nil), "capsule8.event.v0.FileEventFilter")
	proto.RegisterType((*ContainerEventFilter)(nil), "capsule8.event.v0.ContainerEventFilter")
	proto.RegisterType((*ChargenEventFilter)(nil), "capsule8.event.v0.ChargenEventFilter")
	proto.RegisterType((*TickerEventFilter)(nil), "capsule8.event.v0.TickerEventFilter")
	proto.RegisterType((*Modifier)(nil), "capsule8.event.v0.Modifier")
	proto.RegisterType((*ThrottleModifier)(nil), "capsule8.event.v0.ThrottleModifier")
	proto.RegisterType((*LimitModifier)(nil), "capsule8.event.v0.LimitModifier")
	proto.RegisterEnum("capsule8.event.v0.ThrottleModifier_IntervalType", ThrottleModifier_IntervalType_name, ThrottleModifier_IntervalType_value)
}

func init() { proto.RegisterFile("event/v0/subscription.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 1010 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0x9e, 0x25, 0x37, 0x70, 0x8e, 0xfc, 0x57, 0x22, 0x18, 0x84, 0xa4, 0x6b, 0x0d, 0x6d, 0x41,
	0x03, 0x74, 0xb5, 0x5d, 0x27, 0x6d, 0xb3, 0x9b, 0x0d, 0x9d, 0x9b, 0x6c, 0x06, 0xe2, 0xb4, 0x53,
	0x92, 0x5d, 0xec, 0x46, 0x50, 0x24, 0x5a, 0x21, 0x22, 0x4b, 0x02, 0x49, 0x3b, 0xc8, 0xf5, 0x5e,
	0x64, 0x8f, 0xb3, 0x17, 0xd8, 0x6e, 0xf6, 0x00, 0x7b, 0x8d, 0x42, 0xa4, 0x64, 0xd1, 0xb1, 0x92,
	0xf8, 0x8e, 0xfc, 0xf8, 0x7d, 0x9f, 0x0e, 0x79, 0x0e, 0x0f, 0x05, 0x3b, 0x78, 0x8e, 0x23, 0xde,
	0x9b, 0xf7, 0x7b, 0x6c, 0x76, 0xc9, 0x3c, 0x4a, 0x12, 0x4e, 0xe2, 0xa8, 0x9b, 0xd0, 0x98, 0xc7,
	0xe8, 0xa9, 0xe7, 0x26, 0x6c, 0x16, 0xe2, 0xc3, 0xae, 0x60, 0x75, 0xe7, 0xfd, 0xed, 0xe7, 0x41,
	0x1c, 0x07, 0x21, 0xee, 0x09, 0xc2, 0xe5, 0x6c, 0xd2, 0xbb, 0xa1, 0x6e, 0x92, 0x60, 0xca, 0xa4,
	0x64, 0x7b, 0x6b, 0xe1, 0x27, 0x25, 0x02, 0xb5, 0xfe, 0xaa, 0x00, 0x3a, 0x23, 0x41, 0x84, 0xfd,
	0x33, 0xe5, 0x2b, 0x68, 0x08, 0x75, 0xf5, 0xab, 0x66, 0xa5, 0x53, 0xd9, 0x33, 0x06, 0x2f, 0xba,
	0x2b, 0x9f, 0xed, 0xaa, 0x32, 0x7b, 0x49, 0x84, 0x5e, 0x42, 0x4b, 0x9d, 0x3b, 0xc4, 0x37, 0xb5,
	0x4e, 0x65, 0x6f, 0xd3, 0x6e, 0xaa, 0xf0, 0xc8, 0x47, 0xcf, 0x60, 0x93, 0x91, 0x20, 0x72, 0xf9,
	0x8c, 0x62, 0x53, 0x17, 0x94, 0x02, 0xb0, 0xfe, 0xd7, 0xa0, 0xbe, 0x14, 0xdc, 0x07, 0xa8, 0x8b,
	0xcf, 0x3b, 0x13, 0x12, 0x72, 0x4c, 0xb3, 0xe0, 0x9e, 0x97, 0x04, 0x77, 0x94, 0x0e, 0x8e, 0x05,
	0xcb, 0x36, 0x70, 0x31, 0x41, 0x63, 0x68, 0x7b, 0x71, 0xc4, 0x5d, 0x12, 0x61, 0x9a, 0xdb, 0x68,
	0xc2, 0xc6, 0x2a, 0xb1, 0x19, 0xe6, 0xd4, 0xcc, 0xaa, 0xe5, 0x2d, 0x03, 0xe8, 0x47, 0x30, 0xa2,
	0xd8, 0xc7, 0xb9, 0x93, 0x2e, 0x9c, 0xbe, 0x29, 0x71, 0x3a, 0x8d, 0x7d, 0x9c, 0x99, 0x40, 0xb4,
	0x18, 0xa3, 0x5d, 0x68, 0x32, 0x12, 0x79, 0xd8, 0xf1, 0x67, 0xd4, 0x15, 0x07, 0x0e, 0x9d, 0xca,
	0x9e, 0x6e, 0x37, 0x04, 0xfa, 0x31, 0x03, 0x53, 0xda, 0x2c, 0xe2, 0x24, 0x2c, 0x68, 0x86, 0xa4,
	0x09, 0x74, 0x41, 0x7b, 0x0f, 0xb5, 0x69, 0xec, 0x93, 0x09, 0xc1, 0xd4, 0xdc, 0x12, 0xa1, 0xec,
	0x94, 0x84, 0x32, 0xce, 0x28, 0xf6, 0x82, 0x6c, 0xd5, 0x01, 0x8a, 0x00, 0xad, 0x1b, 0x68, 0xdd,
	0xd9, 0x38, 0x6a, 0x83, 0x4e, 0x7c, 0x66, 0x56, 0x3a, 0xfa, 0xde, 0xa6, 0x9d, 0x0e, 0xd1, 0x16,
	0x3c, 0x89, 0xdc, 0x29, 0x66, 0xa6, 0x26, 0x30, 0x39, 0x41, 0x3b, 0xb0, 0x49, 0xa6, 0x6e, 0x80,
	0x9d, 0x94, 0xad, 0x8b, 0x95, 0x9a, 0x00, 0x46, 0x3e, 0x43, 0x2f, 0xc0, 0x90, 0x8b, 0x52, 0x58,
	0x15, 0xcb, 0x20, 0xa0, 0xd3, 0x14, 0xb1, 0xfe, 0xd3, 0xc1, 0x50, 0x32, 0x87, 0x4e, 0xa0, 0xc9,
	0x6e, 0x99, 0xe7, 0x86, 0xa1, 0x23, 0xa2, 0x97, 0x01, 0x18, 0x83, 0xdd, 0xb2, 0x72, 0x94, 0x44,
	0x35, 0xf1, 0x0d, 0xa6, 0x60, 0x2c, 0x75, 0x4b, 0x68, 0xec, 0x61, 0xc6, 0x72, 0x37, 0xed, 0x5e,
	0xb7, 0xcf, 0x92, 0xb8, 0xe4, 0x96, 0x28, 0x18, 0x43, 0x43, 0x30, 0x26, 0x24, 0xc4, 0xb9, 0x95,
	0x2e, 0xac, 0xca, 0x6a, 0xe8, 0x98, 0x84, 0x58, 0xf5, 0x81, 0x49, 0x0e, 0x30, 0x64, 0xab, 0xd5,
	0x98, 0x39, 0x81, 0x70, 0x7a, 0xf9, 0x50, 0x35, 0xaa, 0x76, 0x45, 0x49, 0x16, 0xdb, 0xf4, 0xae,
	0x5c, 0x1a, 0xe0, 0x28, 0x77, 0xf4, 0xef, 0xdd, 0xe6, 0x50, 0x12, 0x97, 0xb6, 0xe9, 0x29, 0x18,
	0x43, 0x23, 0x68, 0x70, 0xe2, 0x5d, 0x17, 0xe1, 0x61, 0x61, 0xf6, 0x5d, 0x89, 0xd9, 0xb9, 0xe0,
	0xa9, 0x5e, 0x75, 0x5e, 0x40, 0xcc, 0xfa, 0x47, 0x07, 0xb4, 0x9a, 0x25, 0xf4, 0x1e, 0xaa, 0xfc,
	0x36, 0xc1, 0xe2, 0x32, 0x37, 0x07, 0xdf, 0x3e, 0x92, 0xda, 0xf3, 0xdb, 0x04, 0xdb, 0x42, 0x80,
	0x5e, 0x81, 0x96, 0x35, 0x96, 0xb4, 0xce, 0x65, 0x13, 0xec, 0xe6, 0x4d, 0xb0, 0x3b, 0x8a, 0xf8,
	0xbb, 0x83, 0xdf, 0xdd, 0x70, 0x86, 0x6d, 0x8d, 0xf8, 0xa8, 0x0f, 0x55, 0x97, 0x06, 0x7d, 0x71,
	0xbd, 0x8c, 0xc1, 0xb3, 0x15, 0xfa, 0x85, 0xc2, 0x17, 0xcc, 0x4c, 0xf1, 0x46, 0xdc, 0xb4, 0x75,
	0x14, 0x6f, 0x32, 0xc5, 0xc0, 0xac, 0xaf, 0xa9, 0x18, 0x64, 0x8a, 0x7d, 0xb3, 0xb1, 0xa6, 0x62,
	0x3f, 0x53, 0x1c, 0x98, 0xcd, 0x35, 0x15, 0x07, 0x99, 0xe2, 0xad, 0xd9, 0x5a, 0x53, 0xf1, 0x16,
	0xbd, 0x06, 0x9d, 0x62, 0xbe, 0xe8, 0x20, 0x0f, 0x9c, 0x6c, 0xca, 0xb3, 0xc6, 0x80, 0x56, 0xaf,
	0xcb, 0x1a, 0x69, 0x55, 0x45, 0x45, 0x5a, 0xad, 0x7f, 0x35, 0x68, 0xdd, 0xb9, 0x33, 0xe8, 0x60,
	0xc9, 0xac, 0xf3, 0xd0, 0x2d, 0x53, 0x0a, 0xe4, 0x10, 0x6a, 0xe9, 0x5d, 0x4b, 0xbb, 0xcd, 0xbd,
	0x79, 0x3f, 0xe3, 0x94, 0x44, 0x81, 0xdc, 0xcd, 0x82, 0x8d, 0x7e, 0x81, 0x76, 0x3e, 0x76, 0x12,
	0x97, 0x73, 0x4c, 0xa3, 0x7b, 0xeb, 0x40, 0x75, 0x68, 0xe5, 0xaa, 0xcf, 0x52, 0x84, 0x86, 0xd0,
	0x8a, 0x13, 0x1c, 0x39, 0x93, 0xd0, 0x0d, 0x98, 0x33, 0x75, 0xd9, 0x75, 0x56, 0x1d, 0xa5, 0xc7,
	0xba, 0x3f, 0x90, 0x36, 0x8d, 0x54, 0x73, 0x9c, 0x4a, 0xc6, 0x2e, 0xbb, 0x46, 0x47, 0xd0, 0xf6,
	0x28, 0x76, 0x39, 0x76, 0xa6, 0xe9, 0x5b, 0x23, 0x5c, 0x1a, 0x8f, 0xbb, 0x34, 0xa5, 0x68, 0x1c,
	0xfb, 0x38, 0xb5, 0xb1, 0x7e, 0x83, 0xad, 0xb2, 0x0e, 0x82, 0x7e, 0x58, 0x3a, 0xdc, 0xdd, 0x47,
	0x1b, 0x8f, 0x92, 0xab, 0xef, 0x01, 0xad, 0xb6, 0x10, 0xf4, 0x35, 0x6c, 0x84, 0x38, 0x0a, 0xf8,
	0x95, 0xb0, 0xac, 0xda, 0xd9, 0xcc, 0xea, 0xc1, 0xd3, 0x95, 0x1e, 0x81, 0xb6, 0xa1, 0x46, 0x22,
	0x8e, 0xe9, 0xdc, 0x0d, 0x05, 0x5d, 0xb7, 0x17, 0x73, 0xeb, 0xcf, 0x0a, 0xd4, 0xf2, 0xd7, 0x0a,
	0xfd, 0x04, 0x35, 0x7e, 0x45, 0x63, 0xce, 0x43, 0x9c, 0x3d, 0xfc, 0x65, 0x45, 0x75, 0x9e, 0x51,
	0x8a, 0x47, 0x2e, 0x17, 0xa1, 0x77, 0xf0, 0x24, 0x24, 0x53, 0xc2, 0xb3, 0x96, 0x51, 0x56, 0x45,
	0x27, 0xe9, 0xfa, 0x42, 0x2a, 0xe9, 0xd6, 0xdf, 0x15, 0x68, 0xdf, 0xb5, 0x7d, 0x28, 0x6c, 0x74,
	0x01, 0x8d, 0x7c, 0xec, 0x88, 0x93, 0xd5, 0xc4, 0xc9, 0xf6, 0xd7, 0x08, 0x37, 0xcd, 0x9f, 0x10,
	0x8a, 0x43, 0xae, 0x13, 0x65, 0x66, 0x7d, 0x80, 0xba, 0xba, 0x8a, 0x5a, 0x60, 0x8c, 0x47, 0x27,
	0x27, 0xa3, 0xb3, 0xa3, 0xe1, 0xa7, 0xd3, 0x8f, 0xed, 0xaf, 0x10, 0xc0, 0x46, 0x36, 0xae, 0xa4,
	0xe3, 0xf1, 0xe8, 0xf4, 0xe2, 0xfc, 0xa8, 0xad, 0xa1, 0x1a, 0x54, 0x7f, 0xfd, 0x74, 0x61, 0xb7,
	0x75, 0x6b, 0x17, 0x1a, 0x4b, 0x5b, 0x4c, 0x5f, 0x71, 0x79, 0x26, 0x72, 0x0f, 0x72, 0xf2, 0xf3,
	0xeb, 0x3f, 0x5e, 0x05, 0x84, 0x5f, 0xcd, 0x2e, 0xbb, 0x5e, 0x3c, 0xed, 0xe5, 0x51, 0xf7, 0x28,
	0x76, 0x3d, 0x4e, 0xe6, 0xf8, 0xb0, 0x97, 0x5c, 0x07, 0x3d, 0x37, 0x21, 0xf2, 0x87, 0xf2, 0x72,
	0x43, 0x54, 0xdf, 0xfe, 0x97, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x2e, 0xf0, 0xe1, 0xb9, 0x0a,
	0x00, 0x00,
}