package paysuper_i18n


import (
	"fmt"
	"github.com/vube/i18n"
	"go.uber.org/zap"
	"os"
)

type Formatter interface {
	FormatCurrency(locale string, amount float64, currency string) (string, error)
}

type formatterImpl struct {
	factory *i18n.TranslatorFactory
}

func NewFormatter() (Formatter, error) {
	dir, _ := os.Getwd()

	f, errs := i18n.NewTranslatorFactory(
		[]string{fmt.Sprintf("%s/data/rules", dir)},
		[]string{"data/messages"},
		"en",
	)

	if len(errs) > 0 {
		zap.S().Errorw("Could not create new factory factory", "err", errs[0])
		return nil, errs[0]
	}

	return &formatterImpl{
		factory: f,
	}, nil
}

func (t *formatterImpl) FormatCurrency(locale string, amount float64, currency string) (string, error) {
	translator, errs := t.factory.GetTranslator(locale)
	if len(errs) > 0 {
		zap.S().Errorw("Can't get translator", "err", errs[0])
		return "", errs[0]
	}

	result, err := translator.FormatCurrency(amount, currency)
	if err != nil {
		zap.S().Errorw("Can't format currency", "err", err)
		return result, err
	}

	return result, nil
}
