package models

import "fmt"

type Coords struct {
	Lat string // latitude
	Lon string // longitude
}

func (c Coords) String() string {
	return fmt.Sprintf("lat %s and lon %s", c.Lat, c.Lon)
}
