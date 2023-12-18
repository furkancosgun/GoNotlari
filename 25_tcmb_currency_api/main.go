package main

import (
	"net/http"
	"tcmb_currency_api/service"
	"tcmb_currency_api/service/response"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	currencyService := service.New()

	e.GET("api/v1/currencies", func(c echo.Context) error {
		currencies, err := currencyService.GetAllCurrencies()
		if err != nil {
			return c.JSON(http.StatusBadGateway, response.NewErrorResponseWithError(err))
		}

		return c.JSON(http.StatusOK, currencies)
	})

	e.GET("api/v1/currencies/:currCode", func(c echo.Context) error {
		paramCurrCode := c.Param("currCode")
		if paramCurrCode == "" {
			return c.JSON(http.StatusBadRequest, response.NewErrorResponseWithString(
				"Para birimi parametresi bulunamadi - api/v1/currencies/{PARA_BIRIMI}",
			))
		}

		curr, err := currencyService.GetCurrencyByCurrCode(paramCurrCode)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.NewErrorResponseWithError(err))
		}

		return c.JSON(http.StatusOK, curr)
	})

	e.Start("localhost:8080")
}
