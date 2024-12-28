package defines

import "github.com/Juram09/Weather-Predictor/internal/domain/entities"

var (
	Ferengi = entities.Planet{
		Name:      "ferengi",
		Distance:  500,
		Position:  0,
		Velocity:  1,
		Clockwise: true,
	}
	Vulcano = entities.Planet{
		Name:      "vulcano",
		Distance:  1000,
		Position:  0,
		Velocity:  5,
		Clockwise: false,
	}
	Betazoide = entities.Planet{
		Name:      "betazoide",
		Distance:  2000,
		Position:  0,
		Velocity:  3,
		Clockwise: true,
	}
)
