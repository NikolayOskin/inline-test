package convert

import "inline-test/model"

func ItemsToPBItems(items []model.Item) []*model.ItemPB {
	var result []*model.ItemPB

	for _, item := range items {
		result = append(result, ItemToPBItem(item))
	}

	return result
}

func ItemsToPBItems2(items []model.Item) []*model.ItemPB {
	var result []*model.ItemPB

	for _, item := range items {
		if !item.IsActive || len(item.Title) < 100 {
			continue
		}

		result = append(result, ItemToPBItem(item))
	}

	return result
}

func ItemToPBItem(item model.Item) *model.ItemPB {
	return &model.ItemPB{
		ID:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		IsActive:    item.IsActive,
		Brands:      brandsToBrandsPB(item.Brands),
	}
}
