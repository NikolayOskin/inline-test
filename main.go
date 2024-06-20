package main

import (
	"fmt"
	"math/rand"
	"time"

	"inline-test/convert"
	"inline-test/model"
)

func main() {
	simpleCase()
	moreComplexCase()
}

func simpleCase() {
	var items []model.Item

	for i := 0; i < randomNumber(1000); i++ {
		items = append(items, model.Item{
			ID:          i,
			Title:       fmt.Sprintf("Item %d", i),
			Description: fmt.Sprintf("Description %d", i),
			IsActive:    rand.Intn(2) == 1,
			Brands: []model.Brand{
				{Title: "Some Title", Description: "Some Description"},
				{Title: "Some Title 2", Description: "Some Description 2"},
			},
		})
	}

	result := convert.ItemsToPBItems(items)
	_ = len(result)
}

func moreComplexCase() {
	var items []model.Item

	for i := 0; i < randomNumber(1000); i++ {
		items = append(items, model.Item{
			ID:          i,
			Title:       fmt.Sprintf("Item %d", i),
			Description: fmt.Sprintf("Description %d", i),
			IsActive:    rand.Intn(2) == 1,
			Brands: []model.Brand{
				{Title: "Some Title", Description: "Some Description"},
				{Title: "Some Title 2", Description: "Some Description 2"},
			},
		})
	}

	result := convert.ItemsToPBItemsComplex(items)
	_ = len(result)
}

func randomNumber(i int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(i)
}
