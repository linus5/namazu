// Code generated by protoc-gen-go.
// source: inspector_message.proto
// DO NOT EDIT!

/*
Package pbinspectorhandler is a generated protocol buffer package.

It is generated from these files:
	inspector_message.proto

It has these top-level messages:
	InspectorMsgReq_Event_FuncCall
	InspectorMsgReq_Event_FuncReturn
	InspectorMsgReq_Event_Exit
	InspectorMsgReq_Event
	InspectorMsgReq_Initiation
	InspectorMsgReq_JavaSpecificFields_StackTraceElement
	InspectorMsgReq_JavaSpecificFields_Params
	InspectorMsgReq_JavaSpecificFields
	InspectorMsgReq
	InspectorMsgRsp
*/
package pbinspectorhandler

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type InspectorMsgReq_Event_Type int32

const (
	InspectorMsgReq_Event_FUNC_CALL   InspectorMsgReq_Event_Type = 1
	InspectorMsgReq_Event_FUNC_RETURN InspectorMsgReq_Event_Type = 2
	InspectorMsgReq_Event_EXIT        InspectorMsgReq_Event_Type = 3
)

var InspectorMsgReq_Event_Type_name = map[int32]string{
	1: "FUNC_CALL",
	2: "FUNC_RETURN",
	3: "EXIT",
}
var InspectorMsgReq_Event_Type_value = map[string]int32{
	"FUNC_CALL":   1,
	"FUNC_RETURN": 2,
	"EXIT":        3,
}

func (x InspectorMsgReq_Event_Type) Enum() *InspectorMsgReq_Event_Type {
	p := new(InspectorMsgReq_Event_Type)
	*p = x
	return p
}
func (x InspectorMsgReq_Event_Type) String() string {
	return proto.EnumName(InspectorMsgReq_Event_Type_name, int32(x))
}
func (x *InspectorMsgReq_Event_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(InspectorMsgReq_Event_Type_value, data, "InspectorMsgReq_Event_Type")
	if err != nil {
		return err
	}
	*x = InspectorMsgReq_Event_Type(value)
	return nil
}

type InspectorMsgReq_Type int32

const (
	InspectorMsgReq_EVENT      InspectorMsgReq_Type = 1
	InspectorMsgReq_INITIATION InspectorMsgReq_Type = 2
)

var InspectorMsgReq_Type_name = map[int32]string{
	1: "EVENT",
	2: "INITIATION",
}
var InspectorMsgReq_Type_value = map[string]int32{
	"EVENT":      1,
	"INITIATION": 2,
}

func (x InspectorMsgReq_Type) Enum() *InspectorMsgReq_Type {
	p := new(InspectorMsgReq_Type)
	*p = x
	return p
}
func (x InspectorMsgReq_Type) String() string {
	return proto.EnumName(InspectorMsgReq_Type_name, int32(x))
}
func (x *InspectorMsgReq_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(InspectorMsgReq_Type_value, data, "InspectorMsgReq_Type")
	if err != nil {
		return err
	}
	*x = InspectorMsgReq_Type(value)
	return nil
}

type InspectorMsgRsp_Result int32

const (
	InspectorMsgRsp_ACK   InspectorMsgRsp_Result = 1
	InspectorMsgRsp_ERROR InspectorMsgRsp_Result = 2
	InspectorMsgRsp_END   InspectorMsgRsp_Result = 3
)

var InspectorMsgRsp_Result_name = map[int32]string{
	1: "ACK",
	2: "ERROR",
	3: "END",
}
var InspectorMsgRsp_Result_value = map[string]int32{
	"ACK":   1,
	"ERROR": 2,
	"END":   3,
}

func (x InspectorMsgRsp_Result) Enum() *InspectorMsgRsp_Result {
	p := new(InspectorMsgRsp_Result)
	*p = x
	return p
}
func (x InspectorMsgRsp_Result) String() string {
	return proto.EnumName(InspectorMsgRsp_Result_name, int32(x))
}
func (x *InspectorMsgRsp_Result) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(InspectorMsgRsp_Result_value, data, "InspectorMsgRsp_Result")
	if err != nil {
		return err
	}
	*x = InspectorMsgRsp_Result(value)
	return nil
}

type InspectorMsgReq_Event_FuncCall struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *InspectorMsgReq_Event_FuncCall) Reset()         { *m = InspectorMsgReq_Event_FuncCall{} }
func (m *InspectorMsgReq_Event_FuncCall) String() string { return proto.CompactTextString(m) }
func (*InspectorMsgReq_Event_FuncCall) ProtoMessage()    {}

func (m *InspectorMsgReq_Event_FuncCall) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type InspectorMsgReq_Event_FuncReturn struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *InspectorMsgReq_Event_FuncReturn) Reset()         { *m = InspectorMsgReq_Event_FuncReturn{} }
func (m *InspectorMsgReq_Event_FuncReturn) String() string { return proto.CompactTextString(m) }
func (*InspectorMsgReq_Event_FuncReturn) ProtoMessage()    {}

func (m *InspectorMsgReq_Event_FuncReturn) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type InspectorMsgReq_Event_Exit struct {
	ExitCode         *int32 `protobuf:"varint,1,req,name=exitCode" json:"exitCode,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *InspectorMsgReq_Event_Exit) Reset()         { *m = InspectorMsgReq_Event_Exit{} }
func (m *InspectorMsgReq_Event_Exit) String() string { return proto.CompactTextString(m) }
func (*InspectorMsgReq_Event_Exit) ProtoMessage()    {}

func (m *InspectorMsgReq_Event_Exit) GetExitCode() int32 {
	if m != nil && m.ExitCode != nil {
		return *m.ExitCode
	}
	return 0
}

type InspectorMsgReq_Event struct {
	Type             *InspectorMsgReq_Event_Type       `protobuf:"varint,1,req,name=type,enum=pbinspectorhandler.InspectorMsgReq_Event_Type" json:"type,omitempty"`
	FuncCall         *InspectorMsgReq_Event_FuncCall   `protobuf:"bytes,2,opt" json:"FuncCall,omitempty"`
	FuncReturn       *InspectorMsgReq_Event_FuncReturn `protobuf:"bytes,3,opt" json:"FuncReturn,omitempty"`
	Exit             *InspectorMsgReq_Event_Exit       `protobuf:"bytes,4,opt" json:"Exit,omitempty"`
	XXX_unrecognized []byte                            `json:"-"`
}

func (m *InspectorMsgReq_Event) Reset()         { *m = InspectorMsgReq_Event{} }
func (m *InspectorMsgReq_Event) String() string { return proto.CompactTextString(m) }
func (*InspectorMsgReq_Event) ProtoMessage()    {}

func (m *InspectorMsgReq_Event) GetType() InspectorMsgReq_Event_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return InspectorMsgReq_Event_FUNC_CALL
}

func (m *InspectorMsgReq_Event) GetFuncCall() *InspectorMsgReq_Event_FuncCall {
	if m != nil {
		return m.FuncCall
	}
	return nil
}

func (m *InspectorMsgReq_Event) GetFuncReturn() *InspectorMsgReq_Event_FuncReturn {
	if m != nil {
		return m.FuncReturn
	}
	return nil
}

func (m *InspectorMsgReq_Event) GetExit() *InspectorMsgReq_Event_Exit {
	if m != nil {
		return m.Exit
	}
	return nil
}

type InspectorMsgReq_Initiation struct {
	ProcessId        *string `protobuf:"bytes,1,req,name=processId" json:"processId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *InspectorMsgReq_Initiation) Reset()         { *m = InspectorMsgReq_Initiation{} }
func (m *InspectorMsgReq_Initiation) String() string { return proto.CompactTextString(m) }
func (*InspectorMsgReq_Initiation) ProtoMessage()    {}

func (m *InspectorMsgReq_Initiation) GetProcessId() string {
	if m != nil && m.ProcessId != nil {
		return *m.ProcessId
	}
	return ""
}

type InspectorMsgReq_JavaSpecificFields_StackTraceElement struct {
	FileName         *string `protobuf:"bytes,1,req,name=fileName" json:"fileName,omitempty"`
	ClassName        *string `protobuf:"bytes,2,req,name=className" json:"className,omitempty"`
	MethodName       *string `protobuf:"bytes,3,req,name=methodName" json:"methodName,omitempty"`
	LineNumber       *int32  `protobuf:"varint,4,req,name=lineNumber" json:"lineNumber,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *InspectorMsgReq_JavaSpecificFields_StackTraceElement) Reset() {
	*m = InspectorMsgReq_JavaSpecificFields_StackTraceElement{}
}
func (m *InspectorMsgReq_JavaSpecificFields_StackTraceElement) String() string {
	return proto.CompactTextString(m)
}
func (*InspectorMsgReq_JavaSpecificFields_StackTraceElement) ProtoMessage() {}

func (m *InspectorMsgReq_JavaSpecificFields_StackTraceElement) GetFileName() string {
	if m != nil && m.FileName != nil {
		return *m.FileName
	}
	return ""
}

func (m *InspectorMsgReq_JavaSpecificFields_StackTraceElement) GetClassName() string {
	if m != nil && m.ClassName != nil {
		return *m.ClassName
	}
	return ""
}

func (m *InspectorMsgReq_JavaSpecificFields_StackTraceElement) GetMethodName() string {
	if m != nil && m.MethodName != nil {
		return *m.MethodName
	}
	return ""
}

func (m *InspectorMsgReq_JavaSpecificFields_StackTraceElement) GetLineNumber() int32 {
	if m != nil && m.LineNumber != nil {
		return *m.LineNumber
	}
	return 0
}

type InspectorMsgReq_JavaSpecificFields_Params struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Value            *string `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *InspectorMsgReq_JavaSpecificFields_Params) Reset() {
	*m = InspectorMsgReq_JavaSpecificFields_Params{}
}
func (m *InspectorMsgReq_JavaSpecificFields_Params) String() string { return proto.CompactTextString(m) }
func (*InspectorMsgReq_JavaSpecificFields_Params) ProtoMessage()    {}

func (m *InspectorMsgReq_JavaSpecificFields_Params) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *InspectorMsgReq_JavaSpecificFields_Params) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type InspectorMsgReq_JavaSpecificFields struct {
	ThreadName           *string                                                 `protobuf:"bytes,1,req,name=threadName" json:"threadName,omitempty"`
	NrStackTraceElements *int32                                                  `protobuf:"varint,2,req,name=nrStackTraceElements" json:"nrStackTraceElements,omitempty"`
	StackTraceElements   []*InspectorMsgReq_JavaSpecificFields_StackTraceElement `protobuf:"bytes,3,rep,name=stackTraceElements" json:"stackTraceElements,omitempty"`
	NrParams             *int32                                                  `protobuf:"varint,4,req,name=nrParams" json:"nrParams,omitempty"`
	Params               []*InspectorMsgReq_JavaSpecificFields_Params            `protobuf:"bytes,5,rep,name=params" json:"params,omitempty"`
	XXX_unrecognized     []byte                                                  `json:"-"`
}

func (m *InspectorMsgReq_JavaSpecificFields) Reset()         { *m = InspectorMsgReq_JavaSpecificFields{} }
func (m *InspectorMsgReq_JavaSpecificFields) String() string { return proto.CompactTextString(m) }
func (*InspectorMsgReq_JavaSpecificFields) ProtoMessage()    {}

func (m *InspectorMsgReq_JavaSpecificFields) GetThreadName() string {
	if m != nil && m.ThreadName != nil {
		return *m.ThreadName
	}
	return ""
}

func (m *InspectorMsgReq_JavaSpecificFields) GetNrStackTraceElements() int32 {
	if m != nil && m.NrStackTraceElements != nil {
		return *m.NrStackTraceElements
	}
	return 0
}

func (m *InspectorMsgReq_JavaSpecificFields) GetStackTraceElements() []*InspectorMsgReq_JavaSpecificFields_StackTraceElement {
	if m != nil {
		return m.StackTraceElements
	}
	return nil
}

func (m *InspectorMsgReq_JavaSpecificFields) GetNrParams() int32 {
	if m != nil && m.NrParams != nil {
		return *m.NrParams
	}
	return 0
}

func (m *InspectorMsgReq_JavaSpecificFields) GetParams() []*InspectorMsgReq_JavaSpecificFields_Params {
	if m != nil {
		return m.Params
	}
	return nil
}

type InspectorMsgReq struct {
	ProcessId             *string                             `protobuf:"bytes,1,req,name=process_id" json:"process_id,omitempty"`
	Type                  *InspectorMsgReq_Type               `protobuf:"varint,2,req,name=type,enum=pbinspectorhandler.InspectorMsgReq_Type" json:"type,omitempty"`
	Pid                   *int32                              `protobuf:"varint,3,req,name=pid" json:"pid,omitempty"`
	Tid                   *int32                              `protobuf:"varint,4,req,name=tid" json:"tid,omitempty"`
	MsgId                 *int32                              `protobuf:"varint,5,req,name=msg_id" json:"msg_id,omitempty"`
	GaMsgId               *int32                              `protobuf:"varint,6,opt,name=ga_msg_id" json:"ga_msg_id,omitempty"`
	Event                 *InspectorMsgReq_Event              `protobuf:"bytes,7,opt" json:"Event,omitempty"`
	Initiation            *InspectorMsgReq_Initiation         `protobuf:"bytes,8,opt" json:"Initiation,omitempty"`
	HasJavaSpecificFields *int32                              `protobuf:"varint,9,req" json:"HasJavaSpecificFields,omitempty"`
	JavaSpecificFields    *InspectorMsgReq_JavaSpecificFields `protobuf:"bytes,10,opt" json:"JavaSpecificFields,omitempty"`
	XXX_unrecognized      []byte                              `json:"-"`
}

func (m *InspectorMsgReq) Reset()         { *m = InspectorMsgReq{} }
func (m *InspectorMsgReq) String() string { return proto.CompactTextString(m) }
func (*InspectorMsgReq) ProtoMessage()    {}

func (m *InspectorMsgReq) GetProcessId() string {
	if m != nil && m.ProcessId != nil {
		return *m.ProcessId
	}
	return ""
}

func (m *InspectorMsgReq) GetType() InspectorMsgReq_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return InspectorMsgReq_EVENT
}

func (m *InspectorMsgReq) GetPid() int32 {
	if m != nil && m.Pid != nil {
		return *m.Pid
	}
	return 0
}

func (m *InspectorMsgReq) GetTid() int32 {
	if m != nil && m.Tid != nil {
		return *m.Tid
	}
	return 0
}

func (m *InspectorMsgReq) GetMsgId() int32 {
	if m != nil && m.MsgId != nil {
		return *m.MsgId
	}
	return 0
}

func (m *InspectorMsgReq) GetGaMsgId() int32 {
	if m != nil && m.GaMsgId != nil {
		return *m.GaMsgId
	}
	return 0
}

func (m *InspectorMsgReq) GetEvent() *InspectorMsgReq_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *InspectorMsgReq) GetInitiation() *InspectorMsgReq_Initiation {
	if m != nil {
		return m.Initiation
	}
	return nil
}

func (m *InspectorMsgReq) GetHasJavaSpecificFields() int32 {
	if m != nil && m.HasJavaSpecificFields != nil {
		return *m.HasJavaSpecificFields
	}
	return 0
}

func (m *InspectorMsgReq) GetJavaSpecificFields() *InspectorMsgReq_JavaSpecificFields {
	if m != nil {
		return m.JavaSpecificFields
	}
	return nil
}

type InspectorMsgRsp struct {
	Res              *InspectorMsgRsp_Result `protobuf:"varint,1,req,name=res,enum=pbinspectorhandler.InspectorMsgRsp_Result" json:"res,omitempty"`
	MsgId            *int32                  `protobuf:"varint,2,opt,name=msg_id" json:"msg_id,omitempty"`
	GaMsgId          *int32                  `protobuf:"varint,3,opt,name=ga_msg_id" json:"ga_msg_id,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *InspectorMsgRsp) Reset()         { *m = InspectorMsgRsp{} }
func (m *InspectorMsgRsp) String() string { return proto.CompactTextString(m) }
func (*InspectorMsgRsp) ProtoMessage()    {}

func (m *InspectorMsgRsp) GetRes() InspectorMsgRsp_Result {
	if m != nil && m.Res != nil {
		return *m.Res
	}
	return InspectorMsgRsp_ACK
}

func (m *InspectorMsgRsp) GetMsgId() int32 {
	if m != nil && m.MsgId != nil {
		return *m.MsgId
	}
	return 0
}

func (m *InspectorMsgRsp) GetGaMsgId() int32 {
	if m != nil && m.GaMsgId != nil {
		return *m.GaMsgId
	}
	return 0
}

func init() {
	proto.RegisterEnum("pbinspectorhandler.InspectorMsgReq_Event_Type", InspectorMsgReq_Event_Type_name, InspectorMsgReq_Event_Type_value)
	proto.RegisterEnum("pbinspectorhandler.InspectorMsgReq_Type", InspectorMsgReq_Type_name, InspectorMsgReq_Type_value)
	proto.RegisterEnum("pbinspectorhandler.InspectorMsgRsp_Result", InspectorMsgRsp_Result_name, InspectorMsgRsp_Result_value)
}
