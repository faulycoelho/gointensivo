package database

import (
	"database/sql"

	"github.com/dev/gointensivo/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	_, err := r.Db.Exec("insert into order (id, price, tax, final_price) Values (?, ?, ?, ?)",
		order.ID, order.Tax, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("select count(*) from order").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil

}
