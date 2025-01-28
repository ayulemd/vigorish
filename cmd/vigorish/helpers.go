package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/shopspring/decimal"
)

func (app *application) readJSON(r io.Reader, dst any) error {
	err := json.NewDecoder(r).Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		default:
			return err
		}
	}
	return nil
}

func (app *application) impliedProbability(price int64) decimal.Decimal {
	var impliedProbability decimal.Decimal
	decimalPrice := decimal.NewFromInt(price)

	if decimalPrice.LessThan(decimal.NewFromInt(0)) {
		decimalPrice = decimalPrice.Abs()
		impliedProbability = decimalPrice.Div(decimalPrice.Add(decimal.NewFromInt(100))).Round(4)
	} else {
		impliedProbability = decimal.NewFromInt(100).Div(decimalPrice.Add(decimal.NewFromInt(100))).Round(4)
	}

	impliedProbability = impliedProbability.Mul(decimal.NewFromInt(100))

	return impliedProbability
}
