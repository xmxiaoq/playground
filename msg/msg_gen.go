package msg

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Data) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 []byte
		zb0001, err = dc.ReadBytes([]byte((*z)))
		if err != nil {
			return
		}
		(*z) = Data(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Data) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteBytes([]byte(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Data) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendBytes(o, []byte(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Data) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 []byte
		zb0001, bts, err = msgp.ReadBytesBytes(bts, []byte((*z)))
		if err != nil {
			return
		}
		(*z) = Data(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Data) Msgsize() (s int) {
	s = msgp.BytesPrefixSize + len([]byte(z))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *MyInt) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 int
		zb0001, err = dc.ReadInt()
		if err != nil {
			return
		}
		(*z) = MyInt(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z MyInt) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteInt(int(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z MyInt) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt(o, int(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MyInt) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 int
		zb0001, bts, err = msgp.ReadIntBytes(bts)
		if err != nil {
			return
		}
		(*z) = MyInt(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z MyInt) Msgsize() (s int) {
	s = msgp.IntSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *MyStruct) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "which":
			var zb0002 uint32
			zb0002, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Which == nil && zb0002 > 0 {
				z.Which = make(map[string]*MyInt, zb0002)
			} else if len(z.Which) > 0 {
				for key, _ := range z.Which {
					delete(z.Which, key)
				}
			}
			for zb0002 > 0 {
				zb0002--
				var za0001 string
				var za0002 *MyInt
				za0001, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					za0002 = nil
				} else {
					if za0002 == nil {
						za0002 = new(MyInt)
					}
					{
						var zb0003 int
						zb0003, err = dc.ReadInt()
						if err != nil {
							return
						}
						*za0002 = MyInt(zb0003)
					}
				}
				z.Which[za0001] = za0002
			}
		case "other":
			{
				var zb0004 []byte
				zb0004, err = dc.ReadBytes([]byte(z.Other))
				if err != nil {
					return
				}
				z.Other = Data(zb0004)
			}
		case "nums":
			var zb0005 uint32
			zb0005, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if zb0005 != uint32(Eight) {
				err = msgp.ArrayError{Wanted: uint32(Eight), Got: zb0005}
				return
			}
			for za0003 := range z.Nums {
				z.Nums[za0003], err = dc.ReadFloat64()
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *MyStruct) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "which"
	err = en.Append(0x83, 0xa5, 0x77, 0x68, 0x69, 0x63, 0x68)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.Which)))
	if err != nil {
		return
	}
	for za0001, za0002 := range z.Which {
		err = en.WriteString(za0001)
		if err != nil {
			return
		}
		if za0002 == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = en.WriteInt(int(*za0002))
			if err != nil {
				return
			}
		}
	}
	// write "other"
	err = en.Append(0xa5, 0x6f, 0x74, 0x68, 0x65, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteBytes([]byte(z.Other))
	if err != nil {
		return
	}
	// write "nums"
	err = en.Append(0xa4, 0x6e, 0x75, 0x6d, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(Eight))
	if err != nil {
		return
	}
	for za0003 := range z.Nums {
		err = en.WriteFloat64(z.Nums[za0003])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *MyStruct) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "which"
	o = append(o, 0x83, 0xa5, 0x77, 0x68, 0x69, 0x63, 0x68)
	o = msgp.AppendMapHeader(o, uint32(len(z.Which)))
	for za0001, za0002 := range z.Which {
		o = msgp.AppendString(o, za0001)
		if za0002 == nil {
			o = msgp.AppendNil(o)
		} else {
			o = msgp.AppendInt(o, int(*za0002))
		}
	}
	// string "other"
	o = append(o, 0xa5, 0x6f, 0x74, 0x68, 0x65, 0x72)
	o = msgp.AppendBytes(o, []byte(z.Other))
	// string "nums"
	o = append(o, 0xa4, 0x6e, 0x75, 0x6d, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(Eight))
	for za0003 := range z.Nums {
		o = msgp.AppendFloat64(o, z.Nums[za0003])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MyStruct) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "which":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Which == nil && zb0002 > 0 {
				z.Which = make(map[string]*MyInt, zb0002)
			} else if len(z.Which) > 0 {
				for key, _ := range z.Which {
					delete(z.Which, key)
				}
			}
			for zb0002 > 0 {
				var za0001 string
				var za0002 *MyInt
				zb0002--
				za0001, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					za0002 = nil
				} else {
					if za0002 == nil {
						za0002 = new(MyInt)
					}
					{
						var zb0003 int
						zb0003, bts, err = msgp.ReadIntBytes(bts)
						if err != nil {
							return
						}
						*za0002 = MyInt(zb0003)
					}
				}
				z.Which[za0001] = za0002
			}
		case "other":
			{
				var zb0004 []byte
				zb0004, bts, err = msgp.ReadBytesBytes(bts, []byte(z.Other))
				if err != nil {
					return
				}
				z.Other = Data(zb0004)
			}
		case "nums":
			var zb0005 uint32
			zb0005, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if zb0005 != uint32(Eight) {
				err = msgp.ArrayError{Wanted: uint32(Eight), Got: zb0005}
				return
			}
			for za0003 := range z.Nums {
				z.Nums[za0003], bts, err = msgp.ReadFloat64Bytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *MyStruct) Msgsize() (s int) {
	s = 1 + 6 + msgp.MapHeaderSize
	if z.Which != nil {
		for za0001, za0002 := range z.Which {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001)
			if za0002 == nil {
				s += msgp.NilSize
			} else {
				s += msgp.IntSize
			}
		}
	}
	s += 6 + msgp.BytesPrefixSize + len([]byte(z.Other)) + 5 + msgp.ArrayHeaderSize + (Eight * (msgp.Float64Size))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Sample3) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Foo":
			z.Foo, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "Bar":
			z.Bar, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "Age":
			z.Age, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "FirstName":
			z.FirstName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "LastName":
			z.LastName, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Sample3) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "Foo"
	err = en.Append(0x85, 0xa3, 0x46, 0x6f, 0x6f)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.Foo)
	if err != nil {
		return
	}
	// write "Bar"
	err = en.Append(0xa3, 0x42, 0x61, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.Bar)
	if err != nil {
		return
	}
	// write "Age"
	err = en.Append(0xa3, 0x41, 0x67, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.Age)
	if err != nil {
		return
	}
	// write "FirstName"
	err = en.Append(0xa9, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.FirstName)
	if err != nil {
		return
	}
	// write "LastName"
	err = en.Append(0xa8, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.LastName)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Sample3) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "Foo"
	o = append(o, 0x85, 0xa3, 0x46, 0x6f, 0x6f)
	o = msgp.AppendInt(o, z.Foo)
	// string "Bar"
	o = append(o, 0xa3, 0x42, 0x61, 0x72)
	o = msgp.AppendInt(o, z.Bar)
	// string "Age"
	o = append(o, 0xa3, 0x41, 0x67, 0x65)
	o = msgp.AppendInt(o, z.Age)
	// string "FirstName"
	o = append(o, 0xa9, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.FirstName)
	// string "LastName"
	o = append(o, 0xa8, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.LastName)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Sample3) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Foo":
			z.Foo, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Bar":
			z.Bar, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Age":
			z.Age, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "FirstName":
			z.FirstName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "LastName":
			z.LastName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Sample3) Msgsize() (s int) {
	s = 1 + 4 + msgp.IntSize + 4 + msgp.IntSize + 4 + msgp.IntSize + 10 + msgp.StringPrefixSize + len(z.FirstName) + 9 + msgp.StringPrefixSize + len(z.LastName)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Sample3List) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Sample3List) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 0
	err = en.Append(0x80)
	if err != nil {
		return err
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Sample3List) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 0
	o = append(o, 0x80)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Sample3List) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Sample3List) Msgsize() (s int) {
	s = 1
	return
}
