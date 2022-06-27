package exchanges

import (
	"fmt"

	. "github.com/matthew188/crex"
	"github.com/matthew188/crex/exchanges/binancefutures"
	"github.com/matthew188/crex/exchanges/bitmex"
	"github.com/matthew188/crex/exchanges/bybit"
	"github.com/matthew188/crex/exchanges/deribit"
	"github.com/matthew188/crex/exchanges/hbdm"
	"github.com/matthew188/crex/exchanges/hbdmswap"
	"github.com/matthew188/crex/exchanges/okexfutures"
	"github.com/matthew188/crex/exchanges/okexswap"
)

func NewExchange(name string, opts ...ApiOption) Exchange {
	params := &Parameters{}

	for _, opt := range opts {
		opt(params)
	}

	return NewExchangeFromParameters(name, params)
}

func NewExchangeFromParameters(name string, params *Parameters) Exchange {
	switch name {
	case BinanceFutures:
		return binancefutures.NewBinanceFutures(params)
	case BitMEX:
		return bitmex.NewBitMEX(params)
	case Deribit:
		return deribit.NewDeribit(params)
	case Bybit:
		return bybit.NewBybit(params)
	case Hbdm:
		return hbdm.NewHbdm(params)
	case HbdmSwap:
		return hbdmswap.NewHbdmSwap(params)
	case OkexFutures:
		return okexfutures.NewOkexFutures(params)
	case OkexSwap:
		return okexswap.NewOkexSwap(params)
	default:
		panic(fmt.Sprintf("new exchange error [%v]", name))
	}
}
