package infrastructure

import (
	"context"
	"os"
	"product_api/common/postgresql"
	"product_api/domain"
	"product_api/persistence"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/random"
	"github.com/stretchr/testify/assert"
)

// Product Repository Arayuzu
var productRepository persistence.IProductRepository

// Connection Pool
var dbPool *pgxpool.Pool

// Boş bir context nesnesi
var ctx = context.Background()

// Ilk burası çalışır testleri çalıştırıdıgımızda
func TestMain(m *testing.M) {

	//bir db pool nesnesi alınır
	dbPool = postgresql.GetConnectionPool(ctx, postgresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		DbName:                "productdb",
		UserName:              "postgres",
		Password:              "postgres",
		MaxConnection:         "10",
		MaxConnectionIdleTime: "100s",
	})

	//Product repository nesnesi uretilir
	productRepository = persistence.NewProductRepository(dbPool)

	// Inıtalize Test Data
	setup(ctx, dbPool)

	//butun testler calistirilir
	code := m.Run()

	// Clear All Test Data
	clear(ctx, dbPool)

	//test sonucundan donen code os ile sisteme verilir
	os.Exit(code)
}

// Inıtalize Test Data
func setup(ctx context.Context, dbPool *pgxpool.Pool) {
	TestDataInitialize(ctx, dbPool)
}

// Clear All Test Data
func clear(ctx context.Context, dbPool *pgxpool.Pool) {
	TruncateTestData(ctx, dbPool)
}

func TestGetAllProducts(t *testing.T) {

	t.Run("Test-GetAllProductsLengthGreatherThan1", func(t *testing.T) {
		products := productRepository.GetAllProducts()

		assert.Greater(t, len(products), 1)
	})

	t.Run("Test-GetAllProductsLenghtEqual8", func(t *testing.T) {
		products := productRepository.GetAllProducts()

		assert.Equal(t, len(products), 8)
	})

	t.Run("Test-GetAllProductsIndex1EqualtTo", func(t *testing.T) {
		products := productRepository.GetAllProducts()

		actualProduct := products[0]
		expectedProduct := domain.Product{Id: 1, Name: "KALEM", Price: 10, Discount: 1, Store: "www.trendyol.com"}

		assert.Equal(t, actualProduct, expectedProduct)
	})

}

func TestGetAllProductsByStore(t *testing.T) {

	t.Run("Test-GetAllProductsByStoreLengthGreatherThan1", func(t *testing.T) {
		products := productRepository.GetAllProductsByStore("www.sahibinden.com")

		//Gelen Liste Uzunlugu 1 e eşit
		assert.Equal(t, len(products), 1)
	})

	t.Run("Test-GetAllProductsByStoreEqualtTo", func(t *testing.T) {
		products := productRepository.GetAllProductsByStore("www.sahibinden.com")

		actualProduct := products[0]
		expectedProduct := domain.Product{Id: 7, Name: "FORD CONNECT 1.8 TDCİ 110HP GLX", Price: 800000, Discount: 0, Store: "www.sahibinden.com"}

		assert.Equal(t, actualProduct, expectedProduct)
	})

	t.Run("Test-GetAllProductsByStorePriceEqualTo", func(t *testing.T) {
		products := productRepository.GetAllProductsByStore("www.sahibinden.com")

		actualProduct := products[0]
		expectedProduct := domain.Product{Name: "FORD CONNECT 1.8 TDCİ 110HP GLX", Price: 800000, Discount: 0, Store: "www.sahibinden.com"}

		assert.Equal(t, actualProduct.Name, expectedProduct.Name)
		assert.Equal(t, actualProduct.Price, expectedProduct.Price)
		assert.Equal(t, actualProduct.Store, expectedProduct.Store)
		assert.Equal(t, actualProduct.Discount, expectedProduct.Discount)
	})

}

func TestProductAdd(t *testing.T) {
	t.Run("Test-AddNewProduct With No Error", func(t *testing.T) {
		err := productRepository.AddProduct(domain.Product{
			Name:     "Test Product",
			Price:    10,
			Discount: 0,
			Store:    "www.test.com",
		})
		assert.Equal(t, nil, err)
	})
	t.Run("Test-ProductAddWith Long String", func(t *testing.T) {

		testProduct := domain.Product{
			Name:     random.String(255),
			Price:    1,
			Discount: 1,
			Store:    random.String(255),
		}
		err := productRepository.AddProduct(testProduct)
		assert.Equal(t, nil, err)
	})

	t.Run("Test-ProductAddAndCompare", func(t *testing.T) {
		testProduct := domain.Product{
			Name:     random.String(255),
			Price:    1,
			Discount: 1,
			Store:    random.String(255),
		}
		err := productRepository.AddProduct(testProduct)
		assert.Equal(t, nil, err)

		prodList := productRepository.GetAllProductsByStore(testProduct.Store)
		expectedProduct := prodList[0]
		assert.Equal(t, expectedProduct.Name, testProduct.Name)
		assert.Equal(t, expectedProduct.Price, testProduct.Price)
		assert.Equal(t, expectedProduct.Discount, testProduct.Discount)
		assert.Equal(t, expectedProduct.Store, testProduct.Store)
	})
}

func TestGetProductById(t *testing.T) {
	clear(ctx, dbPool)
	setup(ctx, dbPool)
	t.Run("Test-Get Product Is Not Null", func(t *testing.T) {
		prod, err := productRepository.GetProductById(1)
		assert.NotEqual(t, prod, nil)
		assert.Equal(t, err, nil)
	})
	t.Run("Test-Get Product Null", func(t *testing.T) {
		emptyProduct := domain.Product{}
		prod, err := productRepository.GetProductById(12)
		assert.Equal(t, prod, emptyProduct)
		assert.NotEqual(t, err, nil)
	})
}

func TestDeleteProductById(t *testing.T) {

	t.Run("DELETE BY ID WITH SUCCESS", func(t *testing.T) {
		err := productRepository.DeleteById(1)
		assert.Equal(t, err, nil)
	})
	clear(ctx, dbPool)
	t.Run("DELETE BY ID WITH ERROR", func(t *testing.T) {
		err := productRepository.DeleteById(1)
		assert.NotEqual(t, err, nil)
	})
	setup(ctx, dbPool)
	t.Run("DELETE PRODUCT BEFORA AFTER COL SIZE -1", func(t *testing.T) {
		//Get all Prods
		products := productRepository.GetAllProducts()

		//Delete Any Prod
		productRepository.DeleteById(int(products[0].Id))

		//Get All Prods
		products2 := productRepository.GetAllProducts()
		assert.NotEqual(t, products, products2)
		assert.NotEqual(t, len(products), len(products2))
		assert.Equal(t, len(products), len(products2)+1)
		assert.Equal(t, len(products)-1, len(products2))
	})
}

func TestUpdatePriceById(t *testing.T) {
	clear(ctx, dbPool)
	setup(ctx, dbPool)
	t.Run("Update Price With Success", func(t *testing.T) {
		err := productRepository.UpdatePriceById(3, 50)
		assert.Equal(t, nil, err)

		newProduct, err := productRepository.GetProductById(3)
		assert.Equal(t, nil, err)

		assert.Equal(t, float32(50), newProduct.Price)
	})
	t.Run("Not Update Price Becase Not Found Product", func(t *testing.T) {
		err := productRepository.UpdatePriceById(50, 50)
		assert.NotEmpty(t, err)
	})
}
