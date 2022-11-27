package entity

import (
	"database/sql"
	"errors"
	"time"
)

type Item struct {
	id            int64
	items_type_id int64
	name          string
	description   string
	price         float64
	stock         int32
	status        int16
	created_at    time.Time
	updated_at    sql.NullTime
	deleted_at    sql.NullTime
}

type DTOItem struct {
	Id            int64
	Items_type_id int64
	Name          string
	Description   string
	Price         float64
	Stock         int32
	Status        int16
	Created_at    time.Time
	Updated_at    sql.NullTime
	Deleted_at    sql.NullTime
}

func NewItem(dto DTOItem) (*Item, error) {

	if dto.Items_type_id == 0 {
		return nil, errors.New("items_type_id is required")
	}

	if dto.Name == "" {
		return nil, errors.New("name is required")
	}

	item := &Item{
		id:            dto.Id,
		items_type_id: dto.Items_type_id,
		name:          dto.Name,
		description:   dto.Description,
		price:         dto.Price,
		stock:         dto.Stock,
		status:        dto.Status,
		created_at:    dto.Created_at,
		updated_at:    dto.Updated_at,
		deleted_at:    dto.Deleted_at,
	}

	return item, nil
}

// Item Getter

func (i *Item) GetId() int64 {
	return i.id
}

func (i *Item) GetItemsTypeId() int64 {
	return i.items_type_id
}

func (i *Item) GetName() string {
	return i.name
}

func (i *Item) GetDescription() string {
	return i.description
}

func (i *Item) GetPrice() float64 {
	return i.price
}

func (i *Item) GetStock() int32 {
	return i.stock
}

func (i *Item) GetStatus() int16 {
	return i.status
}

func (i *Item) GetCreatedAt() time.Time {
	return i.created_at
}

func (i *Item) GetUpdatedAt() sql.NullTime {
	return i.updated_at
}

func (i *Item) GetDeletedAt() sql.NullTime {
	return i.deleted_at
}

// Item Setter

func (i *Item) SetId(value int64) {
	i.id = value
}

func (i *Item) SetItemsTypeId(value int64) {
	i.items_type_id = value
}

func (i *Item) SetName(value string) {
	i.name = value
}

func (i *Item) SetDescription(value string) {
	i.description = value
}

func (i *Item) SetPrice(value float64) {
	i.price = value
}

func (i *Item) SetStock(value int32) {
	i.stock = value
}

func (i *Item) SetStatus(value int16) {
	i.status = value
}

func (i *Item) SetCreatedAt(value time.Time) {
	i.created_at = value
}

func (i *Item) SetUpdatedAt(value sql.NullTime) {
	i.updated_at = value
}

func (i *Item) SetDeletedAt(value sql.NullTime) {
	i.deleted_at = value
}
