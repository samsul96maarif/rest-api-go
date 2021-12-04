package entities

type Product struct {
	BaseEntities
	Name string `json:"name" validate:"required"`
}

func (Product) TableName() string {
	return "products"
}
