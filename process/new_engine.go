package process

import (
	"github.com/shopspring/decimal"
	"matching-engine/engine"
	"matching-engine/errcode"
	"matching-engine/middleware/cache"
)

func NewEngine(symbol string, price decimal.Decimal) *errcode.Errcode {
	if engine.ChanMap[symbol] != nil {
		return errcode.EngineExist
	}

	engine.ChanMap[symbol] = make(chan engine.Order, 100)
	go engine.Run(symbol, price)

	cache.SaveSymbol(symbol)
	cache.SavePrice(symbol, price)

	return errcode.OK
}
