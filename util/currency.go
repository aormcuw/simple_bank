package util

const (
	USD = "USD"
	EUR = "EUR"
	YEN = "YEN"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, YEN:
		return true
	}
	return false
}
