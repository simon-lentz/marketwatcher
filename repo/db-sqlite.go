package repo

import (
	"database/sql"
	"errors"
)

type SQLiteRepo struct {
	Conn *sql.DB
}

func NewSQLiteRepo(db *sql.DB) *SQLiteRepo {
	return &SQLiteRepo{
		Conn: db,
	}
}

func (repo *SQLiteRepo) Migrate() error {
	query := `
	create table if not exists holdings(
		id integer primary key autoincrement,
		units integer not null,
		value integer not null);
	`
	_, err := repo.Conn.Exec(query)
	return err
}

func (repo *SQLiteRepo) InsertHolding(h Holding) (*Holding, error) {
	insert := "insert into holdings (units, value) values (?, ?)"
	result, err := repo.Conn.Exec(insert, h.Units, h.Value)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	h.ID = id

	return &h, nil
}

func (repo *SQLiteRepo) AllHoldings() ([]Holding, error) {
	query := "select id, units, value from holdings order by value"
	rows, err := repo.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []Holding
	for rows.Next() {
		var h Holding
		if err := rows.Scan(
			&h.ID,
			&h.Units,
			&h.Value,
		); err != nil {
			return nil, err
		}

		all = append(all, h)
	}

	return all, nil
}

func (repo *SQLiteRepo) GetHoldingByID(id int64) (*Holding, error) {
	row := repo.Conn.QueryRow("select id, units, value from holdings where id = ?", id)

	var h Holding

	if err := row.Scan(
		&h.ID,
		&h.Units,
		&h.Value,
	); err != nil {
		return nil, err
	}

	return &h, nil
}

func (repo *SQLiteRepo) UpdateHolding(id int64, updated Holding) error {
	if id == 0 {
		return errors.New("Invalid ID")
	}

	update := "update holdings set units = ?, value = ? where id = ?"

	result, err := repo.Conn.Exec(update, updated.Units, updated.Value, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errUpdateFailed
	}

	return nil
}

func (repo *SQLiteRepo) DeleteHolding(id int64) error {
	result, err := repo.Conn.Exec("delete from holdings where id = ?", id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errDeleteFailed
	}

	return nil
}
