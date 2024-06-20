package model

type Item struct {
	ID          int
	Title       string
	Description string
	IsActive    bool
	Brands      []Brand
}

type ItemPB struct {
	ID          int
	Title       string
	Description string
	IsActive    bool
	Brands      []*BrandPB
}
