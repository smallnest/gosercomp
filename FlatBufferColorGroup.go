// automatically generated, do not modify

package gosercomp

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type FlatBufferColorGroup struct {
	_tab flatbuffers.Table
}

func GetRootAsFlatBufferColorGroup(buf []byte, offset flatbuffers.UOffsetT) *FlatBufferColorGroup {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FlatBufferColorGroup{}
	x.Init(buf, n + offset)
	return x
}

func (rcv *FlatBufferColorGroup) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FlatBufferColorGroup) CgId() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *FlatBufferColorGroup) Name() string {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.String(o + rcv._tab.Pos)
	}
	return ""
}

func (rcv *FlatBufferColorGroup) Colors(j int) string {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.String(a + flatbuffers.UOffsetT(j * 4))
	}
	return ""
}

func (rcv *FlatBufferColorGroup) ColorsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func FlatBufferColorGroupStart(builder *flatbuffers.Builder) { builder.StartObject(3) }
func FlatBufferColorGroupAddCgId(builder *flatbuffers.Builder, cgId int32) { builder.PrependInt32Slot(0, cgId, 0) }
func FlatBufferColorGroupAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(name), 0) }
func FlatBufferColorGroupAddColors(builder *flatbuffers.Builder, colors flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(colors), 0) }
func FlatBufferColorGroupStartColorsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(4, numElems, 4)
}
func FlatBufferColorGroupEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
