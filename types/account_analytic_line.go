package types

import (
	"time"
)

type AccountAnalyticLine struct {
	AccountId              Many2One    `xmlrpc:"account_id"`
	Amount                 float64     `xmlrpc:"amount"`
	AmountCurrency         float64     `xmlrpc:"amount_currency"`
	AnalyticAmountCurrency float64     `xmlrpc:"analytic_amount_currency"`
	Code                   string      `xmlrpc:"code"`
	CompanyCurrencyId      Many2One    `xmlrpc:"company_currency_id"`
	CompanyId              Many2One    `xmlrpc:"company_id"`
	CreateDate             time.Time   `xmlrpc:"create_date"`
	CreateUid              Many2One    `xmlrpc:"create_uid"`
	CurrencyId             Many2One    `xmlrpc:"currency_id"`
	Date                   time.Time   `xmlrpc:"date"`
	DepartmentId           Many2One    `xmlrpc:"department_id"`
	DisplayName            string      `xmlrpc:"display_name"`
	EmployeeId             Many2One    `xmlrpc:"employee_id"`
	GeneralAccountId       Many2One    `xmlrpc:"general_account_id"`
	HolidayId              Many2One    `xmlrpc:"holiday_id"`
	Id                     int64       `xmlrpc:"id"`
	LastUpdate             time.Time   `xmlrpc:"__last_update"`
	MoveId                 Many2One    `xmlrpc:"move_id"`
	Name                   string      `xmlrpc:"name"`
	PartnerId              Many2One    `xmlrpc:"partner_id"`
	ProductId              Many2One    `xmlrpc:"product_id"`
	ProductUomId           Many2One    `xmlrpc:"product_uom_id"`
	ProjectId              Many2One    `xmlrpc:"project_id"`
	Ref                    string      `xmlrpc:"ref"`
	SoLine                 Many2One    `xmlrpc:"so_line"`
	TagIds                 []int64     `xmlrpc:"tag_ids"`
	TaskId                 Many2One    `xmlrpc:"task_id"`
	TimesheetInvoiceId     Many2One    `xmlrpc:"timesheet_invoice_id"`
	TimesheetInvoiceType   interface{} `xmlrpc:"timesheet_invoice_type"`
	TimesheetRevenue       float64     `xmlrpc:"timesheet_revenue"`
	UnitAmount             float64     `xmlrpc:"unit_amount"`
	UserId                 Many2One    `xmlrpc:"user_id"`
	WriteDate              time.Time   `xmlrpc:"write_date"`
	WriteUid               Many2One    `xmlrpc:"write_uid"`
}

type AccountAnalyticLineNil struct {
	AccountId              interface{} `xmlrpc:"account_id"`
	Amount                 interface{} `xmlrpc:"amount"`
	AmountCurrency         interface{} `xmlrpc:"amount_currency"`
	AnalyticAmountCurrency interface{} `xmlrpc:"analytic_amount_currency"`
	Code                   interface{} `xmlrpc:"code"`
	CompanyCurrencyId      interface{} `xmlrpc:"company_currency_id"`
	CompanyId              interface{} `xmlrpc:"company_id"`
	CreateDate             interface{} `xmlrpc:"create_date"`
	CreateUid              interface{} `xmlrpc:"create_uid"`
	CurrencyId             interface{} `xmlrpc:"currency_id"`
	Date                   interface{} `xmlrpc:"date"`
	DepartmentId           interface{} `xmlrpc:"department_id"`
	DisplayName            interface{} `xmlrpc:"display_name"`
	EmployeeId             interface{} `xmlrpc:"employee_id"`
	GeneralAccountId       interface{} `xmlrpc:"general_account_id"`
	HolidayId              interface{} `xmlrpc:"holiday_id"`
	Id                     interface{} `xmlrpc:"id"`
	LastUpdate             interface{} `xmlrpc:"__last_update"`
	MoveId                 interface{} `xmlrpc:"move_id"`
	Name                   interface{} `xmlrpc:"name"`
	PartnerId              interface{} `xmlrpc:"partner_id"`
	ProductId              interface{} `xmlrpc:"product_id"`
	ProductUomId           interface{} `xmlrpc:"product_uom_id"`
	ProjectId              interface{} `xmlrpc:"project_id"`
	Ref                    interface{} `xmlrpc:"ref"`
	SoLine                 interface{} `xmlrpc:"so_line"`
	TagIds                 interface{} `xmlrpc:"tag_ids"`
	TaskId                 interface{} `xmlrpc:"task_id"`
	TimesheetInvoiceId     interface{} `xmlrpc:"timesheet_invoice_id"`
	TimesheetInvoiceType   interface{} `xmlrpc:"timesheet_invoice_type"`
	TimesheetRevenue       interface{} `xmlrpc:"timesheet_revenue"`
	UnitAmount             interface{} `xmlrpc:"unit_amount"`
	UserId                 interface{} `xmlrpc:"user_id"`
	WriteDate              interface{} `xmlrpc:"write_date"`
	WriteUid               interface{} `xmlrpc:"write_uid"`
}

var AccountAnalyticLineModel string = "account.analytic.line"

type AccountAnalyticLines []AccountAnalyticLine

type AccountAnalyticLinesNil []AccountAnalyticLineNil

func (s *AccountAnalyticLine) NilableType_() interface{} {
	return &AccountAnalyticLineNil{}
}

func (ns *AccountAnalyticLineNil) Type_() interface{} {
	s := &AccountAnalyticLine{}
	return load(ns, s)
}

func (s *AccountAnalyticLines) NilableType_() interface{} {
	return &AccountAnalyticLinesNil{}
}

func (ns *AccountAnalyticLinesNil) Type_() interface{} {
	s := &AccountAnalyticLines{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountAnalyticLine))
	}
	return s
}
