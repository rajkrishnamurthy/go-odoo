package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type CalendarContactsService struct {
	client *Client
}

func NewCalendarContactsService(c *Client) *CalendarContactsService {
	return &CalendarContactsService{client: c}
}

func (svc *CalendarContactsService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.CalendarContactsModel, name)
}

func (svc *CalendarContactsService) GetByIds(ids []int) (*types.CalendarContactss, error) {
	c := &types.CalendarContactss{}
	return c, svc.client.getByIds(types.CalendarContactsModel, ids, c)
}

func (svc *CalendarContactsService) GetByName(name string) (*types.CalendarContactss, error) {
	c := &types.CalendarContactss{}
	return c, svc.client.getByName(types.CalendarContactsModel, name, c)
}

func (svc *CalendarContactsService) GetAll() (*types.CalendarContactss, error) {
	c := &types.CalendarContactss{}
	return c, svc.client.getAll(types.CalendarContactsModel, c)
}

func (svc *CalendarContactsService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.CalendarContactsModel, fields, relations)
}

func (svc *CalendarContactsService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.CalendarContactsModel, ids, fields, relations)
}

func (svc *CalendarContactsService) Delete(ids []int) error {
	return svc.client.delete(types.CalendarContactsModel, ids)
}
