package entity

type OrderRepositoryInterface interface {
	Save(Order *Order) error
	GetTotal() (int, error)
}
