package mapper_mysql

import (
	"salt-final-transaction/domain/entity"
	model "salt-final-transaction/internal/repository/mysql/models"
)

func ItemModelToEntity(model_item *model.ModelItem) *entity.Item {
	entity_item := &entity.Item{}
	entity_item.SetId(model_item.Id)
	entity_item.SetItemsTypeId(model_item.Items_type_id)
	entity_item.SetName(model_item.Name)
	entity_item.SetDescription(model_item.Description)
	entity_item.SetPrice(model_item.Price)
	entity_item.SetStock(model_item.Stock)
	entity_item.SetIsService(model_item.Is_service)
	entity_item.SetStatus(int16(model_item.Status))
	return entity_item
}
