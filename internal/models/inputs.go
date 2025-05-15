package models

type UserInput struct {
	Name     string `json:"name" example:"John Doe"`
	Email    string `json:"email" example:"john@example.com"`
	Age      int    `json:"age" example:"30"`
	Password string `json:"password" example:"secret123"`
}

type UpdateUserInput struct {
	Name  string `json:"name" example:"Updated John"`
	Email string `json:"email" example:"updated@example.com"`
	Age   int    `json:"age" example:"35"`
}

type LoginInput struct {
	Email    string `json:"email" example:"john@example.com"`
	Password string `json:"password" example:"secret123"`
}

type OrderInput struct {
	Product  string  `json:"product" example:"Laptop"`
	Quantity int     `json:"quantity" example:"2"`
	Price    float64 `json:"price" example:"1499.99"`
}
