package valueobjects

import (
	"errors"
	"github.com/shopspring/decimal"
)

var (
	MoneyZero = Money{money: decimal.Zero}
)

var (
	ErrorNegativeMoney = errors.New("money cannot be less than zero")
)

type Money struct {
	money decimal.Decimal
}

func NewMoney(money decimal.Decimal) (Money, error) {
	if money.GreaterThan(decimal.Zero) {
		return MoneyZero, ErrorNegativeMoney
	}
	return Money{money: money}, nil
}

func (m Money) Add(money Money) Money {
	return Money{
		money: m.money.Add(money.money),
	}
}

func (m Money) Mul(factor decimal.Decimal) Money {
	return Money{
		money: m.money.Mul(factor),
	}
}

func (m Money) MulQuantity(quantity Quantity) Money {
	return Money{
		money: m.money.Mul(decimal.NewFromInt(int64(quantity.Value()))),
	}
}

func (m Money) Value() decimal.Decimal {
	return m.money
}
