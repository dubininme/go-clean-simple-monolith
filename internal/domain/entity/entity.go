package entity

const (
	OrderStatusPending   = "pending"
	OrderStatusPaid      = "paid"
	OrderStatusCancelled = "cancelled"

	ShipmentStatusPending   = "pending"
	ShipmentStatusShipped   = "shipped"
	ShipmentStatusCancelled = "cancelled"
)

type Product struct {
	Id          int32
	Name        string
	Description string
	Price       int64
	SKU         string
	CreatedAt   int64
	UpdatedAt   int64
}

type Order struct {
	Id          int32
	Status      string
	Amount      int64
	Currency    string
	CreatedAt   int64
	UpdatedAt   int64
	CancelledAt *int64
}

type OrderItem struct {
	Id        int32
	OrderId   int32
	ProductId int32
	Quantity  int32
	Price     int64
	CreatedAt int64
	UpdatedAt int64
}

type Shipment struct {
	Id        int32
	OrderId   int32
	Address   string
	Phone     string
	Status    string
	ShippedAt *int64
	CreatedAt int64
	UpdatedAt int64
}
