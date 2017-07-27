package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BaseImportTestsModelsPreviewService struct {
	client *Client
}

func NewBaseImportTestsModelsPreviewService(c *Client) *BaseImportTestsModelsPreviewService {
	return &BaseImportTestsModelsPreviewService{client: c}
}

func (svc *BaseImportTestsModelsPreviewService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.BaseImportTestsModelsPreviewModel, name)
}

func (svc *BaseImportTestsModelsPreviewService) GetByIds(ids []int) (*types.BaseImportTestsModelsPreviews, error) {
	b := &types.BaseImportTestsModelsPreviews{}
	return b, svc.client.getByIds(types.BaseImportTestsModelsPreviewModel, ids, b)
}

func (svc *BaseImportTestsModelsPreviewService) GetByName(name string) (*types.BaseImportTestsModelsPreviews, error) {
	b := &types.BaseImportTestsModelsPreviews{}
	return b, svc.client.getByName(types.BaseImportTestsModelsPreviewModel, name, b)
}

func (svc *BaseImportTestsModelsPreviewService) GetAll() (*types.BaseImportTestsModelsPreviews, error) {
	b := &types.BaseImportTestsModelsPreviews{}
	return b, svc.client.getAll(types.BaseImportTestsModelsPreviewModel, b)
}

func (svc *BaseImportTestsModelsPreviewService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.BaseImportTestsModelsPreviewModel, fields, relations)
}

func (svc *BaseImportTestsModelsPreviewService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseImportTestsModelsPreviewModel, ids, fields, relations)
}

func (svc *BaseImportTestsModelsPreviewService) Delete(ids []int) error {
	return svc.client.delete(types.BaseImportTestsModelsPreviewModel, ids)
}
