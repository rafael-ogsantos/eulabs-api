package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rafael-ogsantos/eulabs-api/application/repositories"
	"github.com/rafael-ogsantos/eulabs-api/application/services"
	"github.com/rafael-ogsantos/eulabs-api/domain"
	"gorm.io/gorm"
)

// New returns a new echo instance
func New(conn *gorm.DB, e *echo.Echo) *echo.Echo {
	g := e.Group("/api")

	// FindById
	g.GET("/product/:id", func(c echo.Context) error {
		id := c.Param("id")
		ctx := c.Request().Context()

		productRepository := repositories.NewProductRepositoryDb(conn)
		productService := services.NewProductService(productRepository)

		product, err := productService.FindById(ctx, id)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, product)
	})

	// Create
	g.POST("/product", func(c echo.Context) (err error) {
		p := new(domain.Product)
		if err = c.Bind(p); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		productRepository := repositories.NewProductRepositoryDb(conn)
		productService := services.NewProductService(productRepository)

		product := &domain.Product{
			Name:        p.Name,
			Description: p.Description,
		}

		productCreated, err := productService.Insert(c.Request().Context(), product)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, productCreated)
	})

	// Update
	g.PUT("/product/:id", func(c echo.Context) error {
		p := new(domain.Product)
		id := c.Param("id")
		ctx := c.Request().Context()

		if err := c.Bind(p); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		productRepository := repositories.NewProductRepositoryDb(conn)
		productService := services.NewProductService(productRepository)

		productUpdated, err := productService.Update(ctx, id, p)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, productUpdated)
	})

	// Delete
	g.DELETE("/product/:id", func(c echo.Context) error {
		id := c.Param("id")
		ctx := c.Request().Context()
		productRepository := repositories.NewProductRepositoryDb(conn)
		productService := services.NewProductService(productRepository)

		err := productService.Delete(ctx, id)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, "Product deleted")
	})

	return e
}
