package model

type TransformRequest struct {
	ID        string `json:"id"`
	Width     int    `json:"width,omitempty"`
	Height    int    `json:"height,omitempty"`
	X         int    `json:"x,omitempty"` // crop offset
	Y         int    `json:"y,omitempty"`
	Angle     int    `json:"angle,omitempty"`     // rotation
	FlipAxis  string `json:"flip_axis,omitempty"` // "horizontal" / "vertical"
	Watermark string `json:"watermark,omitempty"` // overlay image public_id
	Quality   string `json:"quality,omitempty"`   // "auto", "60", ...
	Filter    string `json:"filter,omitempty"`
	Format    string `json:"format,omitempty"`
}
