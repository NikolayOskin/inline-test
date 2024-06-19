package convert

import "inline-test/model"

func ItemsToPBItems(items []model.Item) []*model.ItemPB {
	var result []*model.ItemPB

	for _, item := range items {
		if !item.IsActive {
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

func brandsToBrandsPB(brands []model.Brand) []*model.BrandPB {
	result := make([]*model.BrandPB, 0, len(brands))

	for _, brand := range brands {
		result = append(result, brandToBrandPB(&brand))
	}

	return result
}

func brandToBrandPB(brand *model.Brand) *model.BrandPB {
	return &model.BrandPB{
		Title:       brand.Title,
		Description: brand.Description,
	}
}
