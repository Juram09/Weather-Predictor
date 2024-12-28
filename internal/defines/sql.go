package defines

const (
	GetWeather  = `SELECT weather.meteorolic_conditions FROM weather WHERE weather.id = ?;`
	SaveWeather = `INSERT INTO weather (id, meteorolic_conditions) VALUES (?, ?);`
)
