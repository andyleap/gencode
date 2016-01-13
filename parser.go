package main

import (
	"strconv"

	. "github.com/andyleap/parser"
)

func MakeGrammar() *Grammar {
	Letter := Set("\\p{L}")
	Digit := Set("\\p{Nd}")
	WS := Ignore(Mult(0, 0, Set("\t\n\f\r ")))
	RWS := Ignore(Mult(1, 0, Set("\t\n\f\r ")))
	NL := Ignore(And(Mult(0, 0, Set("\t\f\r ")), Mult(1, 0, And(Lit("\n"), WS))))

	gIdentifier := And(Letter, Mult(0, 0, Or(Letter, Digit)))
	gIdentifier.Node(func(m Match) (Match, error) {
		return String(m), nil
	})

	gIntField := And(Optional(Tag("Var", Lit("v"))), Optional(Tag("Unsigned", Lit("u"))), Lit("int"), Tag("Bits", Or(Lit("8"), Lit("16"), Lit("32"), Lit("64"))))
	gIntField.Node(func(m Match) (Match, error) {
		bits, err := strconv.ParseInt(String(GetTag(m, "Bits")), 10, 64)
		if err != nil {
			return nil, err
		}
		signed := GetTag(m, "Unsigned") == nil
		if GetTag(m, "Var") == nil {
			return &IntField{
				Signed: signed,
				Bits:   int(bits),
			}, nil
		}
		return &VarIntField{
			Signed: signed,
			Bits:   int(bits),
		}, nil
	})

	gFloatField := And(Lit("float"), Tag("Bits", Or(Lit("32"), Lit("64"))))
	gFloatField.Node(func(m Match) (Match, error) {
		bits, err := strconv.ParseInt(String(GetTag(m, "Bits")), 10, 64)
		if err != nil {
			return nil, err
		}
		return &FloatField{
			Bits: int(bits),
		}, nil
	})

	gStringField := And(Lit("string"))
	gStringField.Node(func(m Match) (Match, error) {
		s := &StringField{}
		return s, nil
	})

	gDeferField := And(gIdentifier)
	gDeferField.Node(func(m Match) (Match, error) {
		s := &DeferField{
			Defer: String(m),
		}
		return s, nil
	})

	gUnionDefer := And(gIdentifier)
	gUnionDefer.Node(func(m Match) (Match, error) {
		s := &UnionDefer{
			Defer: String(m),
		}
		return s, nil
	})

	gUnion := And(Lit("union"), Optional(And(RWS, Tag("Interface", gIdentifier))), WS, Lit("{"), WS,
		Mult(0, 0, And(WS, Tag("Defer", gUnionDefer), NL)),
		Lit("}"),
	)
	gUnion.Node(func(m Match) (Match, error) {
		u := &UnionField{}
		for _, v := range GetTags(m, "Defer") {
			u.Structs = append(u.Structs, v.(*UnionDefer))
		}
		return u, nil
	})

	gType := Or(gIntField, gStringField, gFloatField, gUnion, gDeferField)

	gField := And(Tag("Name", gIdentifier), Require(RWS, Tag("Type", gType), NL))
	gField.Node(func(m Match) (Match, error) {
		f := GetTag(m, "Type").(Field)
		f.SetName(GetTag(m, "Name").(string))
		return TagMatch("Field", f), nil
	})

	gStruct := And(Lit("struct"), Require(RWS, Tag("Name", gIdentifier), WS, Lit("{"), WS,
		Mult(0, 0, gField),
		Lit("}"), WS,
	))
	gStruct.Node(func(m Match) (Match, error) {
		s := &Struct{
			Name: GetTag(m, "Name").(string),
		}
		for _, v := range GetTags(m, "Field") {
			s.Fields = append(s.Fields, v.(Field))
		}
		return TagMatch("Struct", s), nil
	})

	gSchema := And(WS, Mult(0, 0, gStruct), WS)
	gSchema.Node(func(m Match) (Match, error) {
		s := &Schema{}
		for _, v := range GetTags(m, "Struct") {
			s.Structs = append(s.Structs, v.(*Struct))
		}
		return s, nil
	})
	return gSchema
}
