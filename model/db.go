package model

type DB struct {
	Restaurants []Restaurants `yaml:"restaurants" json:"restaurants"`
}

type Restaurants struct {
	Name string `yaml:"name" json:"name"`
	ID   int    `yaml:"id" json:"id"`
	Menu []Menu `yaml:"menu" json:"menu"`
}

type Menu struct {
	Name        string   `yaml:"name" json:"name"`
	Price       float64  `yaml:"price" json:"price"`
	Type        string   `yaml:"type" json:"type"`
	Ingredients []string `yaml:"ingredients" json:"ingredients"`
}
