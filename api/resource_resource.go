package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ResourceResourceService struct {
	client *Client
}

func NewResourceResourceService(c *Client) *ResourceResourceService {
	return &ResourceResourceService{client: c}
}

func (svc *ResourceResourceService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ResourceResourceModel, name)
}

func (svc *ResourceResourceService) GetByIds(ids []int) (*types.ResourceResources, error) {
	r := &types.ResourceResources{}
	return r, svc.client.getByIds(types.ResourceResourceModel, ids, r)
}

func (svc *ResourceResourceService) GetByName(name string) (*types.ResourceResources, error) {
	r := &types.ResourceResources{}
	return r, svc.client.getByName(types.ResourceResourceModel, name, r)
}

func (svc *ResourceResourceService) GetAll() (*types.ResourceResources, error) {
	r := &types.ResourceResources{}
	return r, svc.client.getAll(types.ResourceResourceModel, r)
}

func (svc *ResourceResourceService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ResourceResourceModel, fields, relations)
}

func (svc *ResourceResourceService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ResourceResourceModel, ids, fields, relations)
}

func (svc *ResourceResourceService) Delete(ids []int) error {
	return svc.client.delete(types.ResourceResourceModel, ids)
}
