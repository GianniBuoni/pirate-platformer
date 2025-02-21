package lib

type Layer struct {
	Objects []Object `json:"objects,omitempty"`
	Name    string   `json:"name"`
}

type Object struct {
	Image  string  `json:"name"`
	Height float32 `json:"height"`
	Width  float32 `json:"width"`
	X      float32 `json:"x"`
	Y      float32 `json:"y"`
}

type Property struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
}
