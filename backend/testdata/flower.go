package testdata

import (
	"time"

	"github.com/Slowers-team/Slowers-App/database"
)

func GetTestFlowers() []database.Flower {
    return []database.Flower{
		{
			Name: "sunflower",
			LatinName: "Helianthus annuus",
			AddedTime: time.Date(2024, 9, 30, 21, 11, 54, 0, time.UTC),
		},
		{
			Name: "daisy",
			LatinName: "Bellis perennis",
			AddedTime: time.Date(2024, 8, 27, 7, 4, 32, 0, time.UTC),
		},
		{
			Name: "spring crocus",
			LatinName: "Crocus vernus",
			AddedTime: time.Date(2024, 9, 29, 19, 58, 1, 0, time.UTC),
		},
	}
}

func GetTestID() string {
	return "842af389e234e768923475bc"
}
