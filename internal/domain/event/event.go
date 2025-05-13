package event

type OrderCreated struct {
	OrderId   int32
	ClientId  string
	CreatedAt int64
}

type OrderPaid struct {
	OrderId int32
	Amount  int64
	PaidAt  int64
}

type OrderCancelled struct {
	OrderId     int32
	CancelledAt int64
}

type ShipmentCreated struct {
	ShipmentId int32
	OrderId    int32
	CreatedAt  int64
}

type OrderShipped struct {
	OrderId   int32
	ShippedAt int64
}
