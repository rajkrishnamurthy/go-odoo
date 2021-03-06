package types

import (
	"time"
)

type AccountMoveLineReconcileWriteoff struct {
	AnalyticId    Many2One  `xmlrpc:"analytic_id"`
	Comment       string    `xmlrpc:"comment"`
	CreateDate    time.Time `xmlrpc:"create_date"`
	CreateUid     Many2One  `xmlrpc:"create_uid"`
	DateP         time.Time `xmlrpc:"date_p"`
	DisplayName   string    `xmlrpc:"display_name"`
	Id            int64     `xmlrpc:"id"`
	JournalId     Many2One  `xmlrpc:"journal_id"`
	LastUpdate    time.Time `xmlrpc:"__last_update"`
	WriteDate     time.Time `xmlrpc:"write_date"`
	WriteoffAccId Many2One  `xmlrpc:"writeoff_acc_id"`
	WriteUid      Many2One  `xmlrpc:"write_uid"`
}

type AccountMoveLineReconcileWriteoffNil struct {
	AnalyticId    interface{} `xmlrpc:"analytic_id"`
	Comment       interface{} `xmlrpc:"comment"`
	CreateDate    interface{} `xmlrpc:"create_date"`
	CreateUid     interface{} `xmlrpc:"create_uid"`
	DateP         interface{} `xmlrpc:"date_p"`
	DisplayName   interface{} `xmlrpc:"display_name"`
	Id            interface{} `xmlrpc:"id"`
	JournalId     interface{} `xmlrpc:"journal_id"`
	LastUpdate    interface{} `xmlrpc:"__last_update"`
	WriteDate     interface{} `xmlrpc:"write_date"`
	WriteoffAccId interface{} `xmlrpc:"writeoff_acc_id"`
	WriteUid      interface{} `xmlrpc:"write_uid"`
}

var AccountMoveLineReconcileWriteoffModel string = "account.move.line.reconcile.writeoff"

type AccountMoveLineReconcileWriteoffs []AccountMoveLineReconcileWriteoff

type AccountMoveLineReconcileWriteoffsNil []AccountMoveLineReconcileWriteoffNil

func (s *AccountMoveLineReconcileWriteoff) NilableType_() interface{} {
	return &AccountMoveLineReconcileWriteoffNil{}
}

func (ns *AccountMoveLineReconcileWriteoffNil) Type_() interface{} {
	s := &AccountMoveLineReconcileWriteoff{}
	return load(ns, s)
}

func (s *AccountMoveLineReconcileWriteoffs) NilableType_() interface{} {
	return &AccountMoveLineReconcileWriteoffsNil{}
}

func (ns *AccountMoveLineReconcileWriteoffsNil) Type_() interface{} {
	s := &AccountMoveLineReconcileWriteoffs{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountMoveLineReconcileWriteoff))
	}
	return s
}
