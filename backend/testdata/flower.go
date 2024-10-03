package testdata

import (
	"time"

	"github.com/Slowers-team/Slowers-App/database"
)

func GetTestFlowers() []database.Flower {
    return []database.Flower{
		{
			ID: database.NewID("842af389e234e768923475bc"),
			Name: "sunflower",
			LatinName: "Helianthus annuus",
			AddedTime: time.Date(2024, 9, 30, 21, 11, 54, 0, time.UTC),
		},
		{
			ID: database.NewID("485a28e70545c378ff29b438"),
			Name: "daisy",
			LatinName: "Bellis perennis",
			AddedTime: time.Date(2024, 8, 27, 7, 4, 32, 0, time.UTC),
		},
		{
			ID: database.NewID("ac83264ff67837e87eb82322"),
			Name: "spring crocus",
			LatinName: "Crocus vernus",
			AddedTime: time.Date(2024, 9, 29, 19, 58, 1, 0, time.UTC),
		},
	}
}

func GetTestFlowersConcise() []database.Flower {
	flowers := []database.Flower{}
	for _, flower := range GetTestFlowers() {
		flowers = append(flowers, database.Flower{
			Name: flower.Name,
			LatinName: flower.LatinName,
		})
	}
	return flowers
}
