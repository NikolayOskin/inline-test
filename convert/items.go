package convert

import "inline-test/model"

// ItemsToPBItems - упрощенная версия конвертера.
func ItemsToPBItems(items []model.Item) []*model.ItemPB {
	var result []*model.ItemPB

	for _, item := range items {
		result = append(result, itemToPBItem(item))
	}

	return result
}

// ItemsToPBItemsComplex - более сложная версия конвертера, отличающаяся наличием if statement'a внутри цикла.
func ItemsToPBItemsComplex(items []model.Item) []*model.ItemPB {
	var result []*model.ItemPB

	for _, item := range items {
		if !item.IsActive || len(item.Title) < 100 {
			continue
		}

		result = append(result, itemToPBItem(item))
	}

	return result
}

func itemToPBItem(item model.Item) *model.ItemPB {
	return &model.ItemPB{
		ID:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		IsActive:    item.IsActive,
		Brands:      brandsToBrandsPB(item.Brands),
	}
}
