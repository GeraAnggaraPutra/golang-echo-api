package repository

import (
	"database/sql"
	"go-echo/entity"
	"time"
)

type Repository interface {
	GetAll() ([]*entity.Biodata, error)
	FindByID(ID int) (*entity.Biodata, error)
	Create(biodata *entity.Biodata) error
	Update(biodata *entity.Biodata) (*entity.Biodata, error)
	Delete(biodata *entity.Biodata) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]*entity.Biodata, error) {
	biodatas := []*entity.Biodata{}

	sqlQuery := "SELECT * FROM biodata"
	rows, err := r.db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		biodata := &entity.Biodata{}
		var createdAt, updatedAt sql.NullTime
		err := rows.Scan(&biodata.ID, &biodata.NAME, &biodata.AGE, &biodata.ADDRESS, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}
		if createdAt.Valid {
			biodata.CreatedAt = createdAt.Time
		}
		if updatedAt.Valid {
			biodata.UpdatedAt = updatedAt.Time
		}
		biodatas = append(biodatas, biodata)
	}

	return biodatas, nil
}

func (r *repository) FindByID(ID int) (*entity.Biodata, error) {
	biodata := &entity.Biodata{}
	var createdAt, updatedAt sql.NullTime
	sqlQuery := "SELECT * FROM biodata WHERE id = ?"
	err := r.db.QueryRow(sqlQuery, ID).Scan(&biodata.ID, &biodata.NAME, &biodata.AGE, &biodata.ADDRESS, &createdAt, &updatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
            return nil, nil
        }
		return nil, err
	}
	if createdAt.Valid {
		biodata.CreatedAt = createdAt.Time
	}
	if updatedAt.Valid {
		biodata.UpdatedAt = updatedAt.Time
	}

	return biodata, nil
}

func (r *repository) Create(biodata *entity.Biodata) error {
	now := time.Now()
	biodata.CreatedAt = now
	biodata.UpdatedAt = now

	sqlQuery := "INSERT INTO biodata(name, age, address, created_at, updated_at) VALUES(?, ?, ?, ?, ?)"
	_, err := r.db.Exec(sqlQuery, biodata.NAME, biodata.AGE, biodata.ADDRESS, biodata.CreatedAt, biodata.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Update(biodata *entity.Biodata) (*entity.Biodata, error) {
	biodata.UpdatedAt = time.Now()

	sqlQuery := "UPDATE biodata SET name = ?, age = ?, address = ?, updated_at = ? WHERE id = ?"
	_, err := r.db.Exec(sqlQuery, biodata.NAME, biodata.AGE, biodata.ADDRESS, biodata.UpdatedAt, biodata.ID)
	if err != nil {
		return nil, err
	}

	return biodata, nil
}

func (r *repository) Delete(biodata *entity.Biodata) error {
	sqlQuery := "DELETE FROM biodata WHERE id = ?"
	_, err := r.db.Exec(sqlQuery, biodata.ID)
	if err != nil {
		return err
	}

	return nil
}