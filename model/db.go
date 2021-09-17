package model

type DB struct {
	Suppliers []Suppliers `yaml:"suppliers" json:"suppliers"`
}

type Suppliers struct {
	Name         string       `yaml:"name" json:"name"`
	Type         string       `yaml:"type" json:"type"`
	WorkingHours WorkingHours `yaml:"working_hours" json:"workingHours"`
	ID           int          `yaml:"id" json:"id"`
	Image        string       `yaml:"image" json:"image"`
	Menu         []Menu       `yaml:"menu" json:"menu"`
}

type Menu struct {
	ID          int      `yaml:"id" json:"id"`
	Name        string   `yaml:"name" json:"name"`
	Price       float64  `yaml:"price" json:"price"`
	Image       string   `yaml:"image" json:"image"`
	Type        string   `yaml:"type" json:"type"`
	Ingredients []string `yaml:"ingredients" json:"ingredients"`
}

type WorkingHours struct {
	Opening string `yaml:"opening" json:"opening"`
	Closing string `yaml:"closing" json:"closing"`
}
