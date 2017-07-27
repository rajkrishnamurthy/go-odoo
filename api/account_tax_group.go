package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountTaxGroupService struct {
	client *Client
}

func NewAccountTaxGroupService(c *Client) *AccountTaxGroupService {
	return &AccountTaxGroupService{client: c}
}

func (svc *AccountTaxGroupService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountTaxGroupModel, name)
}

func (svc *AccountTaxGroupService) GetByIds(ids []int) (*types.AccountTaxGroups, error) {
	a := &types.AccountTaxGroups{}
	return a, svc.client.getByIds(types.AccountTaxGroupModel, ids, a)
}

func (svc *AccountTaxGroupService) GetByName(name string) (*types.AccountTaxGroups, error) {
	a := &types.AccountTaxGroups{}
	return a, svc.client.getByName(types.AccountTaxGroupModel, name, a)
}

func (svc *AccountTaxGroupService) GetAll() (*types.AccountTaxGroups, error) {
	a := &types.AccountTaxGroups{}
	return a, svc.client.getAll(types.AccountTaxGroupModel, a)
}

func (svc *AccountTaxGroupService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountTaxGroupModel, fields, relations)
}

func (svc *AccountTaxGroupService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountTaxGroupModel, ids, fields, relations)
}

func (svc *AccountTaxGroupService) Delete(ids []int) error {
	return svc.client.delete(types.AccountTaxGroupModel, ids)
}
