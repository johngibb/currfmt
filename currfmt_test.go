package currfmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/johngibb/currfmt"
)

func TestFormatPrice(t *testing.T) {
	tests := []struct {
		val          int64
		currencyCode string
		want         string
		err          string
	}{
		{100, "USD", "$1.00", ""},
		{100000, "USD", "$1,000.00", ""},
		{-100000, "USD", "-$1,000.00", ""},
		{-100000, "GBP", "-£1.000,00", ""},
		{100, "BLAH", "", `unknown currency: "BLAH"`},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v %v", tt.val, tt.currencyCode), func(t *testing.T) {
			got, err := currfmt.FormatPrice(tt.val, tt.currencyCode)
			assertEqual(t, got, tt.want, "formatted")
			assertEqual(t, errMsg(err), tt.err, "error")
		})
	}
}

func errMsg(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

func assertEqual(t *testing.T, got, want interface{}, msg string, args ...interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("%s: got %v, want %v", fmt.Sprintf(msg, args...), got, want)
	}
}
