package regulation

import (
	"github.com/pechenini/foodapi/model"
	"math"
	"math/rand"
	"time"
)

//Regulator is an interface that changes something in the db on the fly
type Regulator interface {
	Regulate(db *model.DB)
}

type PriceRegulator struct {
	timeout time.Duration
}

func NewPriceRegulator(timeout time.Duration) *PriceRegulator {
	return &PriceRegulator{timeout: timeout}
}

func (regulator *PriceRegulator) Regulate(db *model.DB) {
	rand.Seed(time.Now().UnixNano())
	ticker := time.NewTicker(regulator.timeout)
	defer ticker.Stop()
	for range ticker.C {
		for i := 0; i < len(db.Restaurants); i++ {
			for j := 0; j < len(db.Restaurants[i].Menu); j++ {
				newPrice := 2 + rand.Float64()*(8-2)
				db.Restaurants[i].Menu[j].Price = math.Round(newPrice*100) / 100
			}
		}
	}
}
