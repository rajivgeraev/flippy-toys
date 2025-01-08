package model

import (
	"time"

	"github.com/google/uuid"
)

type ToyCondition string
type ToyStatus string
type ToyCategory string

const (
	ConditionNew        ToyCondition = "new"
	ConditionLikeNew    ToyCondition = "like_new"
	ConditionGood       ToyCondition = "good"
	ConditionAcceptable ToyCondition = "acceptable"

	StatusActive    ToyStatus = "active"
	StatusReserved  ToyStatus = "reserved"
	StatusExchanged ToyStatus = "exchanged"
	StatusDeleted   ToyStatus = "deleted"

	CategoryConstruction ToyCategory = "construction_toys"
	CategoryDolls        ToyCategory = "dolls"
	CategoryVehicles     ToyCategory = "vehicles"
	CategoryEducational  ToyCategory = "educational"
	CategoryOutdoor      ToyCategory = "outdoor"
	CategoryBoardGames   ToyCategory = "board_games"
	CategoryElectronic   ToyCategory = "electronic"
	CategoryStuffed      ToyCategory = "stuffed_animals"
	CategoryAction       ToyCategory = "action_figures"
	CategoryArtsCrafts   ToyCategory = "arts_crafts"
	CategoryMusical      ToyCategory = "musical"
	CategoryOther        ToyCategory = "other"
)

type AgeRange struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type Toy struct {
	ID          uuid.UUID     `json:"id"`
	UserID      uuid.UUID     `json:"user_id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	AgeRange    *AgeRange     `json:"age_range,omitempty"`
	Condition   *ToyCondition `json:"condition,omitempty"`
	Category    *ToyCategory  `json:"category,omitempty"`
	Status      ToyStatus     `json:"status"`
	Photos      []Photo       `json:"photos,omitempty"`
	IsDeleted   *time.Time    `json:"is_deleted,omitempty"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

type Photo struct {
	ID           uuid.UUID `json:"id"`
	ToyID        uuid.UUID `json:"toy_id"`
	URL          string    `json:"url"`           // secure_url из Cloudinary
	CloudinaryID string    `json:"cloudinary_id"` // public_id из Cloudinary
	AssetID      string    `json:"asset_id"`      // asset_id из Cloudinary
	IsMain       bool      `json:"is_main"`
	CreatedAt    time.Time `json:"created_at"`
}
