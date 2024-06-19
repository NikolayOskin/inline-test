package convert

import "inline-test/model"

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
