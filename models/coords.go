package models

import "fmt"

type Coords struct {
	Lat string
	Lon string
}

func (c Coords) String() string {
	return fmt.Sprintf("lat %s, lon %s", c.Lat, c.Lon)
}
