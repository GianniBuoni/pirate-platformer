package lib

// Stores layer data: Data (a tile layer)
// or Objects.
type Layer struct {
	Data    []int    `json:"data"`
	Objects []Object `json:"objects,omitempty"`
	Name    string   `json:"name"`
}

type Object struct {
	Properties Properties `json:"properties,omitempty"`
	Image      string     `json:"name"`
	Height     float32    `json:"height"`
	Width      float32    `json:"width"`
	X          float32    `json:"x"`
	Y          float32    `json:"y"`
}

type Properties struct {
	Loader string  `json:"loader"`
	FlipH  float32 `json:"flipH"`
	FlipV  float32 `json:"flipV"`
}

type TileRefs struct {
	TileRef []TileRef `json:"tilesets"`
}

type TileRef struct {
	Source   string `json:"source"`
	FirstGID int    `json:"firstgid"`
}

type Tile struct {
	Image string
	X     float32
	Y     float32
	ImgX  float32
	ImgY  float32
}

type GIDRange struct {
	FirstGID int
	LastGID  int
}

type Tileset struct {
	Count   int `json:"tilecount"`
	Columns int `json:"columns"`
}
