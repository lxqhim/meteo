package weatherType

type Response struct {
	Coord    Coord
	Weather  []Weather
	Base     string
	Main     Main
	Wind     Wind
	Clouds   Clouds
	Dt       int
	Sys      Sys
	Timezone int
	Id       int
	Name     string
	Cod      int
}

type Coord struct {
	Log float64
	Lat float64
}

type Weather struct {
	Id          int
	Main        string
	Description string
	Icon        string
}

type Main struct {
	Temp       float32
	Feels_like float32
	Temp_min   float32
	Temp_max   float32
	Pressure   int
	Humidity   int
}

type Wind struct {
	Speed float32
	Deg   float64
}

type Clouds struct {
	All int
}

type Sys struct {
	TypeSys int
	Id      int
	Message float64
	Country string
	Sunrise int
	Sunset  int
}
