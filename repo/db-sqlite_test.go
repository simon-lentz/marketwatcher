package repo

import (
	"testing"
)

func TestSQLiteRepo_Migrate(t *testing.T) {
	err := testRepo.Migrate()
	if err != nil {
		t.Error("migrate failed:", err)
	}
}

func TestSQLiteRepo_InsertHolding(t *testing.T) {
	h := Holding{
		Units: 5,
		Value: 2500,
	}

	result, err := testRepo.InsertHolding(h)
	if err != nil {
		t.Error("insert failed:", err)
	}

	if result.ID <= 0 {
		t.Error("invalid ID returned:", result.ID)
	}

	if result.Units != 5 || result.Value != 2500 {
		t.Error("invalid units and/or value returned:", result.Units, result.Value)
	}
}

func TestSQLiteRepo_AllHoldings(t *testing.T) {
	h, err := testRepo.AllHoldings()
	if err != nil {
		t.Error("get all failed:", err)
	}

	if len(h) != 1 {
		t.Error("unexpected number of rows created:", err)
	}
}

func TestSQLiteRepo_GetHoldingsByID(t *testing.T) {
	h, err := testRepo.GetHoldingByID(1)
	if err != nil {
		t.Error("failed to retrieve by ID:", err)
	}

	if h.Units != 5 || h.Value != 2500 {
		t.Error("invalid units and/or value retrieved:", h.Units, h.Value)
	}

	_, err = testRepo.GetHoldingByID(2)
	if err == nil {
		t.Error("nonexistent ID accessed:", err)
	}
}

func TestSQLiteRepo_UpdateHolding(t *testing.T) {
	h, err := testRepo.GetHoldingByID(1)
	if err != nil {
		t.Error(err)
	}

	h.Value = 500

	if err = testRepo.UpdateHolding(1, *h); err != nil {
		t.Error("failed to update holding:", err)
		if err != errUpdateFailed {
			t.Error("unexpected error returned:", err)
		}
	}
}

func TestSQLiteRepo_DeleteHolding(t *testing.T) {
	if err := testRepo.DeleteHolding(1); err != nil {
		t.Error("failed to delete holding:", err)
		if err != errDeleteFailed {
			t.Error("unexpected error returned", err)
		}
	}

	if err := testRepo.DeleteHolding(2); err == nil {
		t.Error("nonexistent ID accessed:", err)
	}
}
