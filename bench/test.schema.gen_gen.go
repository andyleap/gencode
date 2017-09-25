package bench

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *A) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "BirthDay":
			z.BirthDay, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "Phone":
			z.Phone, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Siblings":
			z.Siblings, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "Spouse":
			z.Spouse, err = dc.ReadUint8()
			if err != nil {
				return
			}
		case "Money":
			z.Money, err = dc.ReadFloat64()
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
func (z *A) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 6
	// write "Name"
	err = en.Append(0x86, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "BirthDay"
	err = en.Append(0xa8, 0x42, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.BirthDay)
	if err != nil {
		return
	}
	// write "Phone"
	err = en.Append(0xa5, 0x50, 0x68, 0x6f, 0x6e, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Phone)
	if err != nil {
		return
	}
	// write "Siblings"
	err = en.Append(0xa8, 0x53, 0x69, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Siblings)
	if err != nil {
		return
	}
	// write "Spouse"
	err = en.Append(0xa6, 0x53, 0x70, 0x6f, 0x75, 0x73, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteUint8(z.Spouse)
	if err != nil {
		return
	}
	// write "Money"
	err = en.Append(0xa5, 0x4d, 0x6f, 0x6e, 0x65, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteFloat64(z.Money)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *A) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "Name"
	o = append(o, 0x86, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "BirthDay"
	o = append(o, 0xa8, 0x42, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x79)
	o = msgp.AppendInt64(o, z.BirthDay)
	// string "Phone"
	o = append(o, 0xa5, 0x50, 0x68, 0x6f, 0x6e, 0x65)
	o = msgp.AppendString(o, z.Phone)
	// string "Siblings"
	o = append(o, 0xa8, 0x53, 0x69, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x73)
	o = msgp.AppendInt64(o, z.Siblings)
	// string "Spouse"
	o = append(o, 0xa6, 0x53, 0x70, 0x6f, 0x75, 0x73, 0x65)
	o = msgp.AppendUint8(o, z.Spouse)
	// string "Money"
	o = append(o, 0xa5, 0x4d, 0x6f, 0x6e, 0x65, 0x79)
	o = msgp.AppendFloat64(o, z.Money)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *A) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "BirthDay":
			z.BirthDay, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "Phone":
			z.Phone, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Siblings":
			z.Siblings, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "Spouse":
			z.Spouse, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "Money":
			z.Money, bts, err = msgp.ReadFloat64Bytes(bts)
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

func (z *A) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 9 + msgp.Int64Size + 6 + msgp.StringPrefixSize + len(z.Phone) + 9 + msgp.Int64Size + 7 + msgp.Uint8Size + 6 + msgp.Float64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Group) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Members":
			var xsz uint32
			xsz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Members) >= int(xsz) {
				z.Members = z.Members[:xsz]
			} else {
				z.Members = make([]Person, xsz)
			}
			for xvk := range z.Members {
				var isz uint32
				isz, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				for isz > 0 {
					isz--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Name":
						z.Members[xvk].Name, err = dc.ReadString()
						if err != nil {
							return
						}
					case "Age":
						z.Members[xvk].Age, err = dc.ReadUint8()
						if err != nil {
							return
						}
					case "Height":
						z.Members[xvk].Height, err = dc.ReadFloat64()
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
func (z *Group) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Name"
	err = en.Append(0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "Members"
	err = en.Append(0xa7, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Members)))
	if err != nil {
		return
	}
	for xvk := range z.Members {
		// map header, size 3
		// write "Name"
		err = en.Append(0x83, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Members[xvk].Name)
		if err != nil {
			return
		}
		// write "Age"
		err = en.Append(0xa3, 0x41, 0x67, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteUint8(z.Members[xvk].Age)
		if err != nil {
			return
		}
		// write "Height"
		err = en.Append(0xa6, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74)
		if err != nil {
			return err
		}
		err = en.WriteFloat64(z.Members[xvk].Height)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Group) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Name"
	o = append(o, 0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Members"
	o = append(o, 0xa7, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Members)))
	for xvk := range z.Members {
		// map header, size 3
		// string "Name"
		o = append(o, 0x83, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Members[xvk].Name)
		// string "Age"
		o = append(o, 0xa3, 0x41, 0x67, 0x65)
		o = msgp.AppendUint8(o, z.Members[xvk].Age)
		// string "Height"
		o = append(o, 0xa6, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74)
		o = msgp.AppendFloat64(o, z.Members[xvk].Height)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Group) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Members":
			var xsz uint32
			xsz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Members) >= int(xsz) {
				z.Members = z.Members[:xsz]
			} else {
				z.Members = make([]Person, xsz)
			}
			for xvk := range z.Members {
				var isz uint32
				isz, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for isz > 0 {
					isz--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Name":
						z.Members[xvk].Name, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "Age":
						z.Members[xvk].Age, bts, err = msgp.ReadUint8Bytes(bts)
						if err != nil {
							return
						}
					case "Height":
						z.Members[xvk].Height, bts, err = msgp.ReadFloat64Bytes(bts)
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

func (z *Group) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 8 + msgp.ArrayHeaderSize
	for xvk := range z.Members {
		s += 1 + 5 + msgp.StringPrefixSize + len(z.Members[xvk].Name) + 4 + msgp.Uint8Size + 7 + msgp.Float64Size
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Person) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Age":
			z.Age, err = dc.ReadUint8()
			if err != nil {
				return
			}
		case "Height":
			z.Height, err = dc.ReadFloat64()
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
func (z Person) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Name"
	err = en.Append(0x83, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "Age"
	err = en.Append(0xa3, 0x41, 0x67, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteUint8(z.Age)
	if err != nil {
		return
	}
	// write "Height"
	err = en.Append(0xa6, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteFloat64(z.Height)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Person) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Name"
	o = append(o, 0x83, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Age"
	o = append(o, 0xa3, 0x41, 0x67, 0x65)
	o = msgp.AppendUint8(o, z.Age)
	// string "Height"
	o = append(o, 0xa6, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74)
	o = msgp.AppendFloat64(o, z.Height)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Person) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Age":
			z.Age, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "Height":
			z.Height, bts, err = msgp.ReadFloat64Bytes(bts)
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

func (z Person) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 4 + msgp.Uint8Size + 7 + msgp.Float64Size
	return
}
