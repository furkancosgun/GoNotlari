package service

import (
	"errors"
	"product_api/domain"
	"product_api/persistence"
	"product_api/persistence/common"
	"product_api/service/model"
)

type IProductService interface {
	CreateProduct(product model.ProductCreate) error
	GetProductById(id int) (domain.Product, error)
	UpdateProductPriceById(id int, newPrice float32) error
	DeleteProductById(id int) error
	GetAllProducts() []domain.Product
	GetAllProductsByStore(store string) []domain.Product
}

type ProductService struct {
	productRepository persistence.IProductRepository
}

func NewProductService(repository persistence.IProductRepository) IProductService {
	return &ProductService{productRepository: repository}
}

// CreateProduct implements IProductService.
func (service *ProductService) CreateProduct(ProductCreate model.ProductCreate) error {
	valitationErr := validateProductCreate(ProductCreate)
	if valitationErr != nil {
		return valitationErr
	}
	return service.productRepository.AddProduct(domain.Product{
		Name:     ProductCreate.Name,
		Price:    ProductCreate.Price,
		Discount: ProductCreate.Discount,
		Store:    ProductCreate.Store,
	})
}

// DeleteProductById implements IProductService.
func (service *ProductService) DeleteProductById(id int) error {
	if id == 0 {
		return errors.New("Id Not Be Empty")
	}
	return service.productRepository.DeleteById(id)
}

// GetAllProducts implements IProductService.
func (service *ProductService) GetAllProducts() []domain.Product {
	return service.productRepository.GetAllProducts()
}

// GetAllProductsByStore implements IProductService.
func (service *ProductService) GetAllProductsByStore(store string) []domain.Product {
	return service.productRepository.GetAllProductsByStore(store)
}

// GetProductById implements IProductService.
func (service *ProductService) GetProductById(id int) (domain.Product, error) {
	if id == 0 {
		return domain.Product{}, errors.New(common.PRODUCT_ID_NOT_BE_EMPTY)
	}
	return service.productRepository.GetProductById(id)
}

// UpdateProductPriceById implements IProductService.
func (service *ProductService) UpdateProductPriceById(id int, newPrice float32) error {
	if id == 0 {
		return errors.New(common.PRODUCT_ID_NOT_BE_EMPTY)
	}
	return service.productRepository.UpdatePriceById(id, newPrice)
}

func validateProductCreate(product model.ProductCreate) error {
	if product.Discount > 70 {
		return errors.New(common.PRODUCT_DISCOUNT_NOT_BE_GREATHER_THAN_70)
	}
	return nil
}
