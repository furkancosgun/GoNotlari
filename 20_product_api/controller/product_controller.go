package controller

import (
	"net/http"
	"product_api/controller/request"
	"product_api/controller/response"
	"product_api/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productService service.IProductService
}

func NewProductController(productService service.IProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (productController *ProductController) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/v1/products/:id", productController.GetProductById)
	e.GET("/api/v1/products", productController.GetAllProducts)
	e.POST("/api/v1/products", productController.AddProduct)
	e.PUT("/api/v1/products/:id", productController.UpdatePrice)
	e.DELETE("/api/v1/products/:id", productController.DeleteProductById)
}

func (productController *ProductController) GetProductById(e echo.Context) error {
	idParam, atoiErr := strconv.Atoi(e.Param("id"))
	if atoiErr != nil {
		return e.JSON(http.StatusBadRequest, response.ErrorResponse{Error: atoiErr.Error()})
	}
	product, err := productController.productService.GetProductById(idParam)
	if err != nil {
		return e.JSON(http.StatusNotFound, response.ErrorResponse{Error: err.Error()})
	}
	return e.JSON(http.StatusOK, product)
}
func (productController *ProductController) GetAllProducts(e echo.Context) error {
	storeParam := e.QueryParam("store")
	if storeParam != "" {
		return e.JSON(http.StatusOK, productController.productService.GetAllProductsByStore(storeParam))
	}
	return e.JSON(http.StatusOK, productController.productService.GetAllProducts())
}
func (productController *ProductController) AddProduct(e echo.Context) error {
	var addProductRequest request.AddProductRequest
	err := e.Bind(&addProductRequest)
	if err != nil {
		return e.JSON(http.StatusBadRequest, response.ErrorResponse{Error: err.Error()})
	}
	err = productController.productService.CreateProduct(addProductRequest.ToModel())
	if err != nil {
		return e.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{Error: err.Error()})
	}
	return e.NoContent(http.StatusCreated)

}
func (productController *ProductController) UpdatePrice(e echo.Context) error {
	var newPriceRequest request.NewPriceRequest
	idParam, atoiErr := strconv.Atoi(e.Param("id"))
	err := e.Bind(&newPriceRequest)
	if atoiErr != nil {
		return e.JSON(http.StatusBadRequest, response.ErrorResponse{Error: atoiErr.Error()})
	}
	if err != nil {
		return e.JSON(http.StatusBadRequest, response.ErrorResponse{Error: err.Error()})
	}

	err = productController.productService.UpdateProductPriceById(idParam, newPriceRequest.NewPrice)
	if err != nil {
		return e.JSON(http.StatusBadRequest, response.ErrorResponse{Error: err.Error()})
	}
	return e.NoContent(http.StatusOK)
}
func (productController *ProductController) DeleteProductById(e echo.Context) error {
	idParam, atoiErr := strconv.Atoi(e.Param("id"))
	if atoiErr != nil {
		return e.JSON(http.StatusBadRequest, response.ErrorResponse{Error: atoiErr.Error()})
	}
	err := productController.productService.DeleteProductById(idParam)
	if err != nil {
		return e.JSON(http.StatusBadRequest, response.ErrorResponse{Error: err.Error()})
	}
	return e.NoContent(http.StatusOK)
}
