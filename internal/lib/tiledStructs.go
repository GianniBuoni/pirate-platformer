package lib

type MapProps struct {
	Bg        string  `json:"bg"`
	Horizon   float32 `json:"horizon"`
	NextLevel int     `json:"next level"`
	TopLimit  int     `json:"top limit"`
}

// Stores layer data: Data (a tile layer)
// or Objects.
type Layer struct {
	Data    []int    `json:"data"`
	Objects []Object `json:"objects,omitempty"`
	Name    string   `json:"name"`
}

// Basic positional data provided by the Tiled Map
// Properties also stored user defined data.
type Object struct {
	Properties Properties `json:"properties,omitempty"`
	Image      string     `json:"name"`
	Id         int        `json:"id"`
	Height     float32    `json:"height"`
	Width      float32    `json:"width"`
	X          float32    `json:"x"`
	Y          float32    `json:"y"`
}

type Properties struct {
	Loader   string  `json:"loader"`
	Value    int     `json:"value"`
	SpeedX   float32 `json:"speedX"`
	SpeedY   float32 `json:"speedY"`
	DirX     float32 `json:"dirX"`
	DirY     float32 `json:"dirY"`
	FlipH    float32 `json:"flipH"`
	FlipV    float32 `json:"flipV"`
	Lifetime float32 `json:"lifetime"`
}

// Level Maps gloabl into about a tileset.
// This is not the tileset itself, but rather
// a collection of data of how a tileset relates to a level.
type TileRefs struct {
	TileRef []TileRef `json:"tilesets"`
}

type TileRef struct {
	Source   string `json:"source"`
	FirstGID int    `json:"firstgid"`
}

// Data about a single tile
type Tile struct {
	Image string
	X     float32
	Y     float32
	ImgX  float32
	ImgY  float32
}

// Stores the first and last GID's
// that an associated tileset contains.
// Useful for figuring out if a GID belongs to
// a specific tileset.
type GIDRange struct {
	FirstGID int
	LastGID  int
}

// Tileset Data
type Tileset struct {
	Count   int      `json:"tilecount"`
	Columns int      `json:"columns"`
	Tiles   []TsTile `json:"tiles,omitempty"`
}

// Data about a single tile from a tileset file.
// Only emmited by tilesets made of multiple images.
type TsTile struct {
	Image       string      `json:"image"`
	ObjectGroup ObjectGroup `json:"objectgroup,omitempty"`
}

type ObjectGroup struct {
	Hitboxes []Object `json:"objects"`
}

type Template struct {
	Objects Object `json:"object"`
}
