package money

import "strings"

type Currency string

const (
	CNY Currency = "CNY"
	USD Currency = "USD"

	BTC  Currency = "BTC"
	ETH  Currency = "ETH"
	USDT Currency = "USDT"
)

func (c Currency) String() string {
	return string(c)
}

func GetCurrency(c string) (Currency, bool) {
	cc := Currency(strings.ToUpper(c))
	switch cc {
	case CNY, USD, BTC, ETH, USDT:
		return cc, true
	default:
		return "", false
	}
}

func (c Currency) IsCrypto() bool {
	switch c {
	case BTC, ETH, USDT:
		return true
	default:
		return false
	}
}

func (c Currency) FractionDigits() int {
	switch c {
	case CNY:
		return 2
	case USD:
		return 2
	case BTC:
		return 8
	case ETH:
		return 18
	case USDT:
		return 6
	default:
		return 0
	}
}
