package domain

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	id           uint64       `json:"id"`
	name         string       `json:"name"`
	familyName   string       `json:"familyName"`
	age          uint8        `json:"age"`
	birthday     time.Time    `json:"birthday"`
	sharedBy     SharedBy     `json:"sharedBy"`
	relationship Relationship `json:"relationship"`
	status       Status       `json:"status"`
	createdAt    time.Time    `json:"createdAt"`
	updatedAt    time.Time    `json:"updatedAt"`
}

type Status string

const (
	ACTIVE   Status = "ACTIVE"
	INACTIVE        = "INACTIVE"
	BANISH          = "BANISH"
)

type SharedBy string

const (
	ME      SharedBy = "ME"
	FRIENDS          = "FRIENDS"
	ALL              = "ALL"
)

type Relationship string

const (
	MARIED Relationship = "MARIED"
	SINGLE              = "SINGLE"
)
