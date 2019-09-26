package paysuper_i18n

import (
	"github.com/vube/i18n"
	"go.uber.org/zap"
	"time"
)

type Formatter interface {
	FormatCurrency(locale string, amount float64, currency string) (string, error)
	FormatDateTime(locale string, datetime time.Time) (string, error)
}

type formatterImpl struct {
	factory *i18n.TranslatorFactory
}

func NewFormatter(rulesPath, messagesPath []string) (Formatter, error) {
	f, errs := i18n.NewTranslatorFactory(
		rulesPath,
		messagesPath,
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

func (t *formatterImpl) FormatDateTime(locale string, datetime time.Time) (string, error) {
	translator, errs := t.factory.GetTranslator(locale)
	if len(errs) > 0 {
		zap.S().Errorw("Can't get translator", "err", errs[0])
		return "", errs[0]
	}

	result, err := translator.FormatDateTime(i18n.DateFormatMedium, datetime)
	if err != nil {
		zap.S().Errorw("Can't format datetime", "err", err)
		return result, err
	}

	return result, nil
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
