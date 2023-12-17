package service

import (
	"context"
	"os"
	"product_api/common/postgresql"
	"product_api/persistence"
	"product_api/service"
	"product_api/service/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

var productService service.IProductService

func TestMain(m *testing.M) {

	//bir db pool nesnesi alınır
	dbPool := postgresql.GetConnectionPool(context.Background(), postgresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		DbName:                "productdb",
		UserName:              "postgres",
		Password:              "postgres",
		MaxConnection:         "10",
		MaxConnectionIdleTime: "100s",
	})

	productService = service.NewProductService(persistence.NewProductRepository(dbPool))

	code := m.Run()

	os.Exit(code)
}

func TestCreateProduct(t *testing.T) {
	t.Run("Create Product With No Error", func(t *testing.T) {
		err := productService.CreateProduct(model.ProductCreate{
			Name:     "Test1",
			Price:    10,
			Discount: 1,
			Store:    "test.com",
		})

		assert.Equal(t, nil, err)
	})
	t.Run("Create Product With Discount Error Greather 70", func(t *testing.T) {
		err := productService.CreateProduct(model.ProductCreate{
			Name:     "Test1",
			Price:    1,
			Discount: 71,
			Store:    "test.com",
		})

		assert.NotEqual(t, nil, err)
	})
}
