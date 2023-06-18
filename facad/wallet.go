package facad

import "errors"

type wallet struct {
	amount int
}

func initAmount() *wallet {
	return &wallet{
		amount: 0,
	}
}

func (w *wallet) addAmount(amount int) {
	w.amount += amount
}

func (w *wallet) deductAmount(amount int) (err error) {
	if w.amount < amount {
		return errors.New("amount insufficient")
	}
	w.amount = w.amount - amount
	return nil
}
