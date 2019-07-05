package currfmt

import (
	"fmt"
	"math"
	"strings"
)

type currencyInfo struct {
	// Symbol is the short symbol used to represent the price value (e.g. $, £).
	Symbol string

	// MinorPerMajor is the number of minor currency units per major units (i.e.
	// the number of cents per dollar).
	MinorPerMajor int64

	// DecimalSeparator is the symbol between the integer and fractional part of
	// the number (e.g. the period in "1.23").
	DecimalSeparator string

	// PlacesSeparator is the symbol between multiple integer digits (e.g. ","
	// in "1,000,000").
	PlacesSeparator string

	// PlacesMagnitude is the magnitude of each group of digits separated by the
	// places separator (e.g. 1000 for "1,000,000").
	PlacesMagnitude int64
}

// currencies is a map of a currency's ISO 4217 code to its formatting
// information.
var currencies = map[string]currencyInfo{
	"USD": {
		Symbol:           "$",
		MinorPerMajor:    100,
		DecimalSeparator: ".",
		PlacesSeparator:  ",",
		PlacesMagnitude:  1000,
	},
	"GBP": {
		Symbol:           "£",
		MinorPerMajor:    100,
		DecimalSeparator: ",",
		PlacesSeparator:  ".",
		PlacesMagnitude:  1000,
	},
	"JPY": {
		Symbol:        "¥",
		MinorPerMajor: 2,
	},
}

func FormatPrice(val int64, currencyCode string) (string, error) {
	info, ok := currencies[currencyCode]
	if !ok {
		return "", fmt.Errorf("unknown currency: %q", currencyCode)
	}

	neg := val < 0
	prefix := ""
	if neg {
		val *= -1
		prefix = "-"
	}

	major, minor := val/info.MinorPerMajor, val%info.MinorPerMajor
	return prefix + info.Symbol + formatNumWithGroupSeparator(
		major,
		info.PlacesMagnitude,
		info.PlacesSeparator,
	) + info.DecimalSeparator + padLeftZeros(fmt.Sprint(minor), info.MinorPerMajor), nil
}

func formatNumWithGroupSeparator(val, mag int64, separator string) string {
	var result string
	for val != 0 {
		group := fmt.Sprint(val % mag)
		if val/mag != 0 {
			group = padLeftZeros(group, mag)
		}
		if result != "" {
			group = group + separator
		}
		result = group + result
		val /= mag
	}

	return result
}

func padLeftZeros(s string, mag int64) string {
	digits := int(math.Log10(float64(mag)))
	return strings.Repeat("0", digits-len(s)) + s
}
