package math_tool

import "github.com/shopspring/decimal"

/**
 * math_tool.go
 */

func AddDeci(a decimal.Decimal, b *decimal.Decimal) decimal.Decimal {
	if b == nil {
		return a
	}
	return a.Add(*b)
}
func AddDeciPoint(a *decimal.Decimal, b *decimal.Decimal) *decimal.Decimal {
	if a == nil && b == nil {
		return &decimal.Zero
	}
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}

	c := a.Add(*b)
	return &c
}

func DeciNil0(a *decimal.Decimal) decimal.Decimal {
	if a == nil {
		return decimal.Zero
	}
	return *a
}

func DeciPoint4str(str string) *decimal.Decimal {
	d := decimal.RequireFromString(str)
	return &d
}

func DeciPointMin(a, b *decimal.Decimal) *decimal.Decimal {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.LessThan(*b) {
		return a
	} else {
		return b
	}
}

func DeciPointMax(a, b *decimal.Decimal) *decimal.Decimal {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.LessThan(*b) {
		return b
	} else {
		return a
	}
}
