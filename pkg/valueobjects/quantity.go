package valueobjects

import "errors"

var (
	QuantityZero = Quantity{quantity: 0}
)

var (
	ErrorNegativeQuatity = errors.New("quantity cannot be less than zero")
)

type Quantity struct {
	quantity int
}

func NewQuantity(quantity int) (Quantity, error) {
	if quantity < 0 {
		return QuantityZero, ErrorNegativeQuatity
	}
	return Quantity{quantity: quantity}, nil
}

func (q Quantity) Value() int {
	return q.quantity
}

func (q Quantity) Add(addend Quantity) Quantity {
	return Quantity{
		quantity: q.quantity + addend.quantity,
	}
}

func (q Quantity) Sub(subtrahend Quantity) Quantity {
	return Quantity{
		quantity: q.quantity - subtrahend.quantity,
	}
}
