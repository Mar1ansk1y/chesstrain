package bill

import "go/token"

type Bill struct {
	BillToken token.Token
	EnableBill bool
}

func (b *Bill) BillProcess() {
	// здесь происходит проверка оплаты
}
