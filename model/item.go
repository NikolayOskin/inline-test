package model

type Item struct {
	ID          int
	Title       string
	Description string
	IsActive    bool
	Brands      []Brand
}

type Brand struct {
	Title       string
	Description string
}

type ItemPB struct {
	ID          int
	Title       string
	Description string
	IsActive    bool
	Brands      []*BrandPB
}

type BrandPB struct {
	Title       string
	Description string
}
