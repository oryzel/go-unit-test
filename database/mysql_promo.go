package database

import "database/sql"

type Discount interface {
	FindCurrentPromo() int
}

type DiscountRepository struct {
	db *sql.DB
}

func NewDiscountRepository(db *sql.DB) *DiscountRepository {
	return &DiscountRepository{
		db: db,
	}
}

func (dc *DiscountRepository) FindCurrentPromo() (discount int) {

	sql := "SELECT value FROM discount LIMIT 1"
	row := dc.db.QueryRow(sql)
	row.Scan(&discount)

	return discount
}
