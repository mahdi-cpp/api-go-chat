package model

type IconBase struct {
	Key    int     `json:"key"`
	Name   string  `json:"name"`
	Width  int     `json:"width,omitempty"`
	Height int     `json:"height,omitempty"`
	Dx     float32 `json:"dx,omitempty"`
	Dy     float32 `json:"dy,omitempty"`
	Color  int     `json:"color,omitempty"`
	Alpha  int     `json:"alpha,omitempty"`
}
