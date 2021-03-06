// Code generated by protoc-gen-go.
// source: EpocFrame.proto
// DO NOT EDIT!

package insensate

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type EpocSensor struct {
	Label            *string `protobuf:"bytes,1,req,name=label" json:"label,omitempty"`
	Value            *int32  `protobuf:"varint,2,req,name=value" json:"value,omitempty"`
	Quality          *int32  `protobuf:"varint,3,opt,name=quality" json:"quality,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EpocSensor) Reset()         { *m = EpocSensor{} }
func (m *EpocSensor) String() string { return proto.CompactTextString(m) }
func (*EpocSensor) ProtoMessage()    {}

func (m *EpocSensor) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *EpocSensor) GetValue() int32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *EpocSensor) GetQuality() int32 {
	if m != nil && m.Quality != nil {
		return *m.Quality
	}
	return 0
}

type EpocFrame struct {
	Counter *uint32 `protobuf:"varint,1,req,name=counter" json:"counter,omitempty"`
	GyroX   *int32  `protobuf:"varint,2,req,name=gyro_x" json:"gyro_x,omitempty"`
	GyroY   *int32  `protobuf:"varint,3,req,name=gyro_y" json:"gyro_y,omitempty"`
	// In order:
	//
	//   F3, FC6, P7, T8, F7, F8, T7, P8, AF4, F4, AF3, O2, O1, FC5
	Sensors          []*EpocSensor `protobuf:"bytes,4,rep,name=sensors" json:"sensors,omitempty"`
	Battery          *uint32       `protobuf:"varint,5,opt,name=battery" json:"battery,omitempty"`
	Timestamp        *uint64       `protobuf:"fixed64,6,opt,name=timestamp" json:"timestamp,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *EpocFrame) Reset()         { *m = EpocFrame{} }
func (m *EpocFrame) String() string { return proto.CompactTextString(m) }
func (*EpocFrame) ProtoMessage()    {}

func (m *EpocFrame) GetCounter() uint32 {
	if m != nil && m.Counter != nil {
		return *m.Counter
	}
	return 0
}

func (m *EpocFrame) GetGyroX() int32 {
	if m != nil && m.GyroX != nil {
		return *m.GyroX
	}
	return 0
}

func (m *EpocFrame) GetGyroY() int32 {
	if m != nil && m.GyroY != nil {
		return *m.GyroY
	}
	return 0
}

func (m *EpocFrame) GetSensors() []*EpocSensor {
	if m != nil {
		return m.Sensors
	}
	return nil
}

func (m *EpocFrame) GetBattery() uint32 {
	if m != nil && m.Battery != nil {
		return *m.Battery
	}
	return 0
}

func (m *EpocFrame) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func init() {
}
