package types

import (
	"time"
)

type IrQwebFieldDatetime struct {
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
}

type IrQwebFieldDatetimeNil struct {
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
}

var IrQwebFieldDatetimeModel string = "ir.qweb.field.datetime"

type IrQwebFieldDatetimes []IrQwebFieldDatetime

type IrQwebFieldDatetimesNil []IrQwebFieldDatetimeNil

func (s *IrQwebFieldDatetime) NilableType_() interface{} {
	return &IrQwebFieldDatetimeNil{}
}

func (ns *IrQwebFieldDatetimeNil) Type_() interface{} {
	s := &IrQwebFieldDatetime{}
	return load(ns, s)
}

func (s *IrQwebFieldDatetimes) NilableType_() interface{} {
	return &IrQwebFieldDatetimesNil{}
}

func (ns *IrQwebFieldDatetimesNil) Type_() interface{} {
	s := &IrQwebFieldDatetimes{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*IrQwebFieldDatetime))
	}
	return s
}
