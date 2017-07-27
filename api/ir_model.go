package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrModelService struct {
	client *Client
}

func NewIrModelService(c *Client) *IrModelService {
	return &IrModelService{client: c}
}

func (svc *IrModelService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrModelModel, name)
}

func (svc *IrModelService) GetByIds(ids []int) (*types.IrModels, error) {
	i := &types.IrModels{}
	return i, svc.client.getByIds(types.IrModelModel, ids, i)
}

func (svc *IrModelService) GetByName(name string) (*types.IrModels, error) {
	i := &types.IrModels{}
	return i, svc.client.getByName(types.IrModelModel, name, i)
}

func (svc *IrModelService) GetAll() (*types.IrModels, error) {
	i := &types.IrModels{}
	return i, svc.client.getAll(types.IrModelModel, i)
}

func (svc *IrModelService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrModelModel, fields, relations)
}

func (svc *IrModelService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrModelModel, ids, fields, relations)
}

func (svc *IrModelService) Delete(ids []int) error {
	return svc.client.delete(types.IrModelModel, ids)
}
