package types

import (
	"time"
)

type UtmMedium struct {
	Active      bool      `xmlrpc:"active"`
	CreateDate  time.Time `xmlrpc:"create_date"`
	CreateUid   Many2One  `xmlrpc:"create_uid"`
	DisplayName string    `xmlrpc:"display_name"`
	Id          int64     `xmlrpc:"id"`
	LastUpdate  time.Time `xmlrpc:"__last_update"`
	Name        string    `xmlrpc:"name"`
	WriteDate   time.Time `xmlrpc:"write_date"`
	WriteUid    Many2One  `xmlrpc:"write_uid"`
}

type UtmMediumNil struct {
	Active      bool        `xmlrpc:"active"`
	CreateDate  interface{} `xmlrpc:"create_date"`
	CreateUid   interface{} `xmlrpc:"create_uid"`
	DisplayName interface{} `xmlrpc:"display_name"`
	Id          interface{} `xmlrpc:"id"`
	LastUpdate  interface{} `xmlrpc:"__last_update"`
	Name        interface{} `xmlrpc:"name"`
	WriteDate   interface{} `xmlrpc:"write_date"`
	WriteUid    interface{} `xmlrpc:"write_uid"`
}

var UtmMediumModel string = "utm.medium"

type UtmMediums []UtmMedium

type UtmMediumsNil []UtmMediumNil

func (s *UtmMedium) NilableType_() interface{} {
	return &UtmMediumNil{}
}

func (ns *UtmMediumNil) Type_() interface{} {
	s := &UtmMedium{}
	return load(ns, s)
}

func (s *UtmMediums) NilableType_() interface{} {
	return &UtmMediumsNil{}
}

func (ns *UtmMediumsNil) Type_() interface{} {
	s := &UtmMediums{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*UtmMedium))
	}
	return s
}
