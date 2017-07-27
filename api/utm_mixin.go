package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type UtmMixinService struct {
	client *Client
}

func NewUtmMixinService(c *Client) *UtmMixinService {
	return &UtmMixinService{client: c}
}

func (svc *UtmMixinService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.UtmMixinModel, name)
}

func (svc *UtmMixinService) GetByIds(ids []int) (*types.UtmMixins, error) {
	u := &types.UtmMixins{}
	return u, svc.client.getByIds(types.UtmMixinModel, ids, u)
}

func (svc *UtmMixinService) GetByName(name string) (*types.UtmMixins, error) {
	u := &types.UtmMixins{}
	return u, svc.client.getByName(types.UtmMixinModel, name, u)
}

func (svc *UtmMixinService) GetAll() (*types.UtmMixins, error) {
	u := &types.UtmMixins{}
	return u, svc.client.getAll(types.UtmMixinModel, u)
}

func (svc *UtmMixinService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.UtmMixinModel, fields, relations)
}

func (svc *UtmMixinService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.UtmMixinModel, ids, fields, relations)
}

func (svc *UtmMixinService) Delete(ids []int) error {
	return svc.client.delete(types.UtmMixinModel, ids)
}
