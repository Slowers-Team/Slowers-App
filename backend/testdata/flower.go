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
	}
}

func GetTestID() string {
	return "842af389e234e768923475bc"
}
