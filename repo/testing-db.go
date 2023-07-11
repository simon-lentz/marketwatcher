package repo

type TestRepo struct{}

func NewTestRepo() *TestRepo {
	return &TestRepo{}
}

// repeat func signatures to satisfy db interface
func (repo *TestRepo) Migrate() error {
	return nil
}

func (repo *TestRepo) InsertHolding(h Holding) (*Holding, error) {
	return &h, nil
}

func (repo *TestRepo) AllHoldings() ([]Holding, error) {
	var all []Holding
	return all, nil
}

func (repo *TestRepo) GetHoldingByID(id int64) (*Holding, error) {
	h := Holding{
		Units: 5,
		Value: 2500,
	}
	return &h, nil
}

func (repo *TestRepo) UpdateHolding(id int64, updated Holding) error {
	return nil
}

func (repo *TestRepo) DeleteHolding(id int64) error {
	return nil
}
