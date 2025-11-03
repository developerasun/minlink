package pkg

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/minlink/internal/constant"
	"github.com/shopspring/decimal"
)

func validateAddress(address string) error {
	_, found := strings.CutPrefix(address, "0x")

	if !found || len(address) != 42 {
		error := errors.New("validateAddress.go: invalid ethereum address")
		return error
	}

	return nil
}

/*
@return e.g `1000000000000000000`
*/
func toWei(_amount string, _tokenType string) string {
	amount, err := decimal.NewFromString(_amount)

	if err != nil {
		log.Fatalln(err.Error())
	}

	targetDecimal := constant.ETH_DECIMAL
	if _tokenType == "usdt" {
		targetDecimal = constant.USDT_DECIMAL
	}

	plain := decimal.NewFromFloat(targetDecimal)
	calculated := amount.Mul(plain)

	return calculated.String()
}

/*
@return e.g `N*1e18`
*/
func toWeiAsExponent(_amount string, _tokenType string) string {
	targetDecimal := constant.ETH_DECIMAL
	if _tokenType == "usdt" {
		targetDecimal = constant.USDT_DECIMAL
	}

	toFloat, _ := strconv.ParseFloat(_amount, 64)
	target := fmt.Sprintf("%.e", toFloat*targetDecimal)

	converted := strings.Replace(target, "+0", "", 1)
	return converted
}
