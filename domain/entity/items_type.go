package entity

import (
	"database/sql"
	"errors"
	"time"
)

type ItemsType struct {
	id         int64
	name       string
	status     int16
	created_at time.Time
	updated_at sql.NullTime
	deleted_at sql.NullTime
}

type DTOItemsType struct {
	Id         int64
	Name       string
	Status     int16
	Created_at time.Time
	Updated_at sql.NullTime
	Deleted_at sql.NullTime
}

func NewItemsType(dto DTOItemsType) (*ItemsType, error) {
	if dto.Name == "" {
		return nil, errors.New("Name is required")
	}

	items_type := &ItemsType{
		id:         dto.Id,
		name:       dto.Name,
		created_at: dto.Created_at,
		updated_at: dto.Updated_at,
		deleted_at: dto.Deleted_at,
	}

	return items_type, nil
}

// Item's Type Getter

func (it *ItemsType) GetId() int64 {
	return it.id
}

func (it *ItemsType) GetName() string {
	return it.name
}

func (it *ItemsType) GetCreatedAt() time.Time {
	return it.created_at
}
func (it *ItemsType) GetUpdatedAt() sql.NullTime {
	return it.updated_at
}

func (it *ItemsType) GetDeletedAt() sql.NullTime {
	return it.deleted_at
}

// Item's Type Setter

func (it *ItemsType) SetId(value int64) {
	it.id = value
}

func (it *ItemsType) SetName(value string) {
	it.name = value
}

func (it *ItemsType) SetCreatedAt(value time.Time) {
	it.created_at = value
}
func (it *ItemsType) SetUpdatedAt(value sql.NullTime) {
	it.updated_at = value
}

func (it *ItemsType) SetDeletedAt(value sql.NullTime) {
	it.deleted_at = value
}
