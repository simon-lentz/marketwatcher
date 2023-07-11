package repo

import "errors"

var (
	errUpdateFailed = errors.New("Update Failed...")
	errDeleteFailed = errors.New("Delete Failed...")
)

type Repo interface {
	Migrate() error
	InsertHolding(h Holding) (*Holding, error)
	AllHoldings() ([]Holding, error)
	GetHoldingByID(id int64) (*Holding, error)
	UpdateHolding(id int64, updated Holding) error
	DeleteHolding(id int64) error
}

type Holding struct {
	ID    int64   `json:"id"`
	Units int64   `json:"units"`
	Value float64 `json:"value"`
}
