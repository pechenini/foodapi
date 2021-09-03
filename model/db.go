package model

type DB struct {
	Restaurants []Restaurants `yaml:"restaurants" json:"restaurants"`
}

type Restaurants struct {
	Name  string `yaml:"name" json:"name"`
	ID    int    `yaml:"id" json:"id"`
	Image string `yaml:"image" json:"image"`
	Menu  []Menu `yaml:"menu" json:"menu"`
}

type Menu struct {
	ID          int      `yaml:"id" json:"id"`
	Name        string   `yaml:"name" json:"name"`
	Price       float64  `yaml:"price" json:"price"`
	Image       string   `yaml:"image" json:"image"`
	Type        string   `yaml:"type" json:"type"`
	Ingredients []string `yaml:"ingredients" json:"ingredients"`
}
