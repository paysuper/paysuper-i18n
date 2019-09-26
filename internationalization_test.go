package paysuper_i18n

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_formatterImpl_FormatCurrency(t *testing.T) {
	type args struct {
		locale   string
		amount   float64
		currency string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"USD", args{locale: "en", amount: 32.45, currency: "USD"}, "$32.45", false},
		{"RUB", args{locale: "en", amount: 32.45, currency: "RUB"}, "₽32.45", false},
		{"EUR", args{locale: "en", amount: 32.45, currency: "EUR"}, "€32.45", false},
		{"RUB in ru locale", args{locale: "ru", amount: 32.45, currency: "RUB"}, "32,45 руб.", false},
	}

	loc, err := NewFormatter([]string{"internal/data/rules"}, []string{"internal/data/messages"})
	if err != nil {
		t.Errorf("Error during ctor. Error %s", err.Error())
		return
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := loc.FormatCurrency(tt.args.locale, tt.args.amount, tt.args.currency)
			if (err != nil) != tt.wantErr {
				t.Errorf("formatterImpl.FormatCurrency() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.EqualValues(t, tt.want, got)
		})
	}
}
