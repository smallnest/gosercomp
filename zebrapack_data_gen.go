package gosercomp

// NOTE: THIS FILE WAS PRODUCED BY THE
// ZEBRAPACK CODE GENERATION TOOL (github.com/glycerine/zebrapack)
// DO NOT EDIT

import "github.com/glycerine/zebrapack/msgp"

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *ZColorGroup) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields0zckf = 3

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields0zckf uint32
	totalEncodedFields0zckf, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zckf := totalEncodedFields0zckf
	missingFieldsLeft0zckf := maxFields0zckf - totalEncodedFields0zckf

	var nextMiss0zckf int = -1
	var found0zckf [maxFields0zckf]bool
	var curField0zckf int

doneWithStruct0zckf:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zckf > 0 || missingFieldsLeft0zckf > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zckf, missingFieldsLeft0zckf, msgp.ShowFound(found0zckf[:]), decodeMsgFieldOrder0zckf)
		if encodedFieldsLeft0zckf > 0 {
			encodedFieldsLeft0zckf--
			curField0zckf, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss0zckf < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zckf = 0
			}
			for nextMiss0zckf < maxFields0zckf && (found0zckf[nextMiss0zckf] || decodeMsgFieldSkip0zckf[nextMiss0zckf]) {
				nextMiss0zckf++
			}
			if nextMiss0zckf == maxFields0zckf {
				// filled all the empty fields!
				break doneWithStruct0zckf
			}
			missingFieldsLeft0zckf--
			curField0zckf = nextMiss0zckf
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zckf)
		switch curField0zckf {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "id"
			found0zckf[0] = true
			z.Id, err = dc.ReadInt()
			if err != nil {
				panic(err)
			}
		case 1:
			// zid 1 for "name"
			found0zckf[1] = true
			z.Name, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case 2:
			// zid 2 for "colors"
			found0zckf[2] = true
			var zxzo uint32
			zxzo, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Colors) >= int(zxzo) {
				z.Colors = (z.Colors)[:zxzo]
			} else {
				z.Colors = make([]string, zxzo)
			}
			for zxup := range z.Colors {
				z.Colors[zxup], err = dc.ReadString()
				if err != nil {
					panic(err)
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss0zckf != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of ZColorGroup
var decodeMsgFieldOrder0zckf = []string{"id", "name", "colors"}

var decodeMsgFieldSkip0zckf = []bool{false, false, false}

// fieldsNotEmpty supports omitempty tags
func (z *ZColorGroup) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 3
	}
	var fieldsInUse uint32 = 3
	isempty[0] = (z.Id == 0) // number, omitempty
	if isempty[0] {
		fieldsInUse--
	}
	isempty[1] = (len(z.Name) == 0) // string, omitempty
	if isempty[1] {
		fieldsInUse--
	}
	isempty[2] = (len(z.Colors) == 0) // string, omitempty
	if isempty[2] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *ZColorGroup) EncodeMsg(en *msgp.Writer) (err error) {

	// honor the omitempty tags
	var empty_zrkb [3]bool
	fieldsInUse_ztvc := z.fieldsNotEmpty(empty_zrkb[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_ztvc + 1)
	if err != nil {
		return err
	}

	if !empty_zrkb[0] {
		// zid 0 for "id"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteInt(z.Id)
		if err != nil {
			panic(err)
		}
	}

	if !empty_zrkb[1] {
		// zid 1 for "name"
		err = en.Append(0x1)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Name)
		if err != nil {
			panic(err)
		}
	}

	if !empty_zrkb[2] {
		// zid 2 for "colors"
		err = en.Append(0x2)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Colors)))
		if err != nil {
			panic(err)
		}
		for zxup := range z.Colors {
			err = en.WriteString(z.Colors[zxup])
			if err != nil {
				panic(err)
			}
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ZColorGroup) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [3]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[0] {
		// zid 0 for "id"
		o = append(o, 0x0)
		o = msgp.AppendInt(o, z.Id)
	}

	if !empty[1] {
		// zid 1 for "name"
		o = append(o, 0x1)
		o = msgp.AppendString(o, z.Name)
	}

	if !empty[2] {
		// zid 2 for "colors"
		o = append(o, 0x2)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Colors)))
		for zxup := range z.Colors {
			o = msgp.AppendString(o, z.Colors[zxup])
		}
	}

	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ZColorGroup) UnmarshalMsg(bts []byte) (o []byte, err error) {
	cfg := &msgp.RuntimeConfig{UnsafeZeroCopy: true}
	return z.UnmarshalMsgWithCfg(bts, cfg)
}
func (z *ZColorGroup) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields1zhyh = 3

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields1zhyh uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zhyh, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zhyh := totalEncodedFields1zhyh
	missingFieldsLeft1zhyh := maxFields1zhyh - totalEncodedFields1zhyh

	var nextMiss1zhyh int = -1
	var found1zhyh [maxFields1zhyh]bool
	var curField1zhyh int

doneWithStruct1zhyh:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zhyh > 0 || missingFieldsLeft1zhyh > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zhyh, missingFieldsLeft1zhyh, msgp.ShowFound(found1zhyh[:]), unmarshalMsgFieldOrder1zhyh)
		if encodedFieldsLeft1zhyh > 0 {
			encodedFieldsLeft1zhyh--
			curField1zhyh, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				panic(err)
				return
			}
		} else {
			//missing fields need handling
			if nextMiss1zhyh < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zhyh = 0
			}
			for nextMiss1zhyh < maxFields1zhyh && (found1zhyh[nextMiss1zhyh] || unmarshalMsgFieldSkip1zhyh[nextMiss1zhyh]) {
				nextMiss1zhyh++
			}
			if nextMiss1zhyh == maxFields1zhyh {
				// filled all the empty fields!
				break doneWithStruct1zhyh
			}
			missingFieldsLeft1zhyh--
			curField1zhyh = nextMiss1zhyh
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zhyh)
		switch curField1zhyh {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "id"
			found1zhyh[0] = true
			z.Id, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				panic(err)
			}
		case 1:
			// zid 1 for "name"
			found1zhyh[1] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case 2:
			// zid 2 for "colors"
			found1zhyh[2] = true
			if nbs.AlwaysNil {
				(z.Colors) = (z.Colors)[:0]
			} else {

				var zudl uint32
				zudl, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Colors) >= int(zudl) {
					z.Colors = (z.Colors)[:zudl]
				} else {
					z.Colors = make([]string, zudl)
				}
				for zxup := range z.Colors {
					z.Colors[zxup], bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss1zhyh != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of ZColorGroup
var unmarshalMsgFieldOrder1zhyh = []string{"id", "name", "colors"}

var unmarshalMsgFieldSkip1zhyh = []bool{false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ZColorGroup) Msgsize() (s int) {
	s = 1 + 3 + msgp.IntSize + 5 + msgp.StringPrefixSize + len(z.Name) + 7 + msgp.ArrayHeaderSize
	for zxup := range z.Colors {
		s += msgp.StringPrefixSize + len(z.Colors[zxup])
	}
	return
}

// ZebraSchemaInMsgpack2Format provides the ZebraPack Schema in msgpack2 format, length 507 bytes
func ZebraSchemaInMsgpack2Format() []byte { return zebraSchemaInMsgpack2Format }

var zebraSchemaInMsgpack2Format = []byte{
	0x85, 0xaa, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68, 0xb1, 0x7a, 0x65, 0x62,
	0x72, 0x61, 0x70, 0x61, 0x63, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x67, 0x6f, 0xad, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0xa9, 0x67, 0x6f, 0x73,
	0x65, 0x72, 0x63, 0x6f, 0x6d, 0x70, 0xad, 0x5a, 0x65, 0x62, 0x72, 0x61, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x49, 0x64, 0x0, 0xa7, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73, 0x81, 0xab, 0x5a,
	0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0xab, 0x5a, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x93, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0x0,
	0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa2, 0x49, 0x64, 0xac,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa2, 0x69, 0x64, 0xac,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa3, 0x69, 0x6e, 0x74,
	0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0xd, 0xad,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b,
	0x69, 0x6e, 0x64, 0xd, 0xa3, 0x53, 0x74, 0x72, 0xa3, 0x69, 0x6e, 0x74, 0x87, 0xa3, 0x5a, 0x69,
	0x64, 0x1, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa4, 0x4e,
	0x61, 0x6d, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65,
	0xa4, 0x6e, 0x61, 0x6d, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53,
	0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72,
	0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75,
	0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x2, 0xa3, 0x53, 0x74,
	0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x86, 0xa3, 0x5a, 0x69, 0x64, 0x2, 0xab, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa6, 0x43, 0x6f, 0x6c, 0x6f, 0x72,
	0x73, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa6, 0x63,
	0x6f, 0x6c, 0x6f, 0x72, 0x73, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53,
	0x74, 0x72, 0xa8, 0x5b, 0x5d, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xad, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x83, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x1a, 0xa3,
	0x53, 0x74, 0x72, 0xa5, 0x53, 0x6c, 0x69, 0x63, 0x65, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x2, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0xa7, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x90,
}

// ZebraSchemaInJsonCompact provides the ZebraPack Schema in compact JSON format, length 644 bytes
func ZebraSchemaInJsonCompact() []byte { return zebraSchemaInJsonCompact }

var zebraSchemaInJsonCompact = []byte(`{"SourcePath":"zebrapack_data.go","SourcePackage":"gosercomp","ZebraSchemaId":0,"Structs":{"ZColorGroup":{"StructName":"ZColorGroup","Fields":[{"Zid":0,"FieldGoName":"Id","FieldTagName":"id","FieldTypeStr":"int","FieldCategory":23,"FieldPrimitive":13,"FieldFullType":{"Kind":13,"Str":"int"}},{"Zid":1,"FieldGoName":"Name","FieldTagName":"name","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}},{"Zid":2,"FieldGoName":"Colors","FieldTagName":"colors","FieldTypeStr":"[]string","FieldCategory":26,"FieldFullType":{"Kind":26,"Str":"Slice","Domain":{"Kind":2,"Str":"string"}}}]}},"Imports":[]}`)

// ZebraSchemaInJsonPretty provides the ZebraPack Schema in pretty JSON format, length 1612 bytes
func ZebraSchemaInJsonPretty() []byte { return zebraSchemaInJsonPretty }

var zebraSchemaInJsonPretty = []byte(`{
    "SourcePath": "zebrapack_data.go",
    "SourcePackage": "gosercomp",
    "ZebraSchemaId": 0,
    "Structs": {
        "ZColorGroup": {
            "StructName": "ZColorGroup",
            "Fields": [
                {
                    "Zid": 0,
                    "FieldGoName": "Id",
                    "FieldTagName": "id",
                    "FieldTypeStr": "int",
                    "FieldCategory": 23,
                    "FieldPrimitive": 13,
                    "FieldFullType": {
                        "Kind": 13,
                        "Str": "int"
                    }
                },
                {
                    "Zid": 1,
                    "FieldGoName": "Name",
                    "FieldTagName": "name",
                    "FieldTypeStr": "string",
                    "FieldCategory": 23,
                    "FieldPrimitive": 2,
                    "FieldFullType": {
                        "Kind": 2,
                        "Str": "string"
                    }
                },
                {
                    "Zid": 2,
                    "FieldGoName": "Colors",
                    "FieldTagName": "colors",
                    "FieldTypeStr": "[]string",
                    "FieldCategory": 26,
                    "FieldFullType": {
                        "Kind": 26,
                        "Str": "Slice",
                        "Domain": {
                            "Kind": 2,
                            "Str": "string"
                        }
                    }
                }
            ]
        }
    },
    "Imports": []
}`)
