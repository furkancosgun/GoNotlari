package persistence

import (
	"context"
	"errors"
	"product_api/domain"
	"product_api/persistence/common"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

// Soyut Bir ProductRepository Yapısı
type IProductRepository interface {
	GetAllProducts() []domain.Product
	GetAllProductsByStore(store string) []domain.Product
	AddProduct(product domain.Product) error
	GetProductById(id int) (domain.Product, error)
	DeleteById(id int) error
	UpdatePriceById(id int, newPrice float32) error
}

// ProductRepository y
type ProductRepository struct {
	dbPool *pgxpool.Pool
}

// Product Repository Nesnesi Oluşturmak İçin
func NewProductRepository(dbpool *pgxpool.Pool) IProductRepository {
	return &ProductRepository{dbPool: dbpool}
}

// IProductRepository arayuzunun GetAllProducts() methodunu implement eder
func (productRepository *ProductRepository) GetAllProducts() []domain.Product {
	//Boş bir context nesnesi
	ctx := context.Background()

	//Postgresql uzerinde sorgumuz çalıştrılır
	productRows, err := productRepository.dbPool.Query(ctx, "SELECT * FROM products")
	if err != nil { //eger sorguda veya baglantıda bir hata yoksa devam edilir
		log.Errorf("Error While Getting All Products:%v", err)
		return []domain.Product{}
	}

	/*
	   //Return Edecegimiz product listesi
	   var productList = []domain.Product{}

	   //Product Listesini doldurmak için kullanacagımız product nesnesi
	   var product = domain.Product{}

	   //Eger sorgu sonucunda biz sonraki satıra gidebiliyorsa dongu içerisinde

	   	for productRows.Next() {
	   		//Scan ifades SELECT ID,NAME,PRICE FROM ... sırasına gore bizim kolon ve degerlerimi mapler
	   		productRows.Scan(&product.Id, &product.Name, &product.Price, &product.Discount, &product.Store)

	   		//mapledigimiz ve veriyi aldgımız product nesnesini listeye atarız
	   		productList = append(productList, product)
	   	}

	   //en sonda donuş ifademiz
	   return productList
	*/
	return extractProductsFromRows(productRows)
}

// IProductRepository arayuzunun GetAllProducts() methodunu implement eder
func (productRepository *ProductRepository) GetAllProductsByStore(store string) []domain.Product {

	ctx := context.Background()

	productRows, err := productRepository.dbPool.Query(ctx, "SELECT * FROM products WHERE store = $1", store)
	if err != nil {
		log.Errorf("Error While Getting All Products:%v", err)
		return []domain.Product{}
	}

	return extractProductsFromRows(productRows)
}

func (productRepository *ProductRepository) AddProduct(product domain.Product) error {
	ctx := context.Background()

	productAddResult, err := productRepository.dbPool.Exec(ctx,
		"INSERT INTO products (name,price,discount,store) VALUES ($1,$2,$3,$4)",
		product.Name, product.Price, product.Discount, product.Store)
	if err != nil {
		log.Errorf("AddProduct Has Error:%v", err)
		return err
	}
	log.Info("Product Added With", productAddResult.String())

	return nil
}

func (productRepository *ProductRepository) GetProductById(id int) (domain.Product, error) {
	ctx := context.Background()
	productRow := productRepository.dbPool.QueryRow(ctx, "SELECT * FROM products WHERE id = $1", id)

	return extractProductFromRow(productRow)
}
func (productRepository *ProductRepository) DeleteById(id int) error {
	_, err := productRepository.GetProductById(id)
	if err != nil {
		return err
	}

	ctx := context.Background()
	_, err = productRepository.dbPool.Exec(ctx, "DELETE FROM products WHERE id = $1", id)

	return err
}

func (productRepository *ProductRepository) UpdatePriceById(id int, newPrice float32) error {
	_, err := productRepository.GetProductById(id)
	if err != nil {
		return err
	}

	ctx := context.Background()
	_, err = productRepository.dbPool.Exec(ctx, "UPDATE products SET price = $2 WHERE id = $1", id, newPrice)

	return err
}

func extractProductsFromRows(productRows pgx.Rows) []domain.Product {
	var productList = []domain.Product{}

	for productRows.Next() {
		prod, err := extractProductFromRow(productRows)
		if err != nil {
			break
		}
		productList = append(productList, prod)
	}
	return productList
}

func extractProductFromRow(productRow pgx.Row) (domain.Product, error) {
	var product = domain.Product{}
	err := productRow.Scan(&product.Id, &product.Name, &product.Price, &product.Discount, &product.Store)
	if err != nil {
		return product, errors.New(common.PRODUCT_NOT_FOUND)
	}

	return product, nil
}
