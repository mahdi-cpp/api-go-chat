package model

import "time"

type Temperature struct {
	ID        uint      `gorm:"primarykey"`
	NodeId    string    `json:"nodeId"`
	Value     int32     `json:"value"`
	CreatedAt time.Time `json:"createdAt"`
}

type Lamp struct {
	ID    uint   `gorm:"primarykey"`
	Name  string `json:"name"`
	Icon  string `json:"icon"`
	Value int32  `json:"value"`
}

type Object struct {
	Type       string `json:"type"`
	JsonString string `json:"jsonString"`
}
