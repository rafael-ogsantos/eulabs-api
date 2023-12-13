package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rafael-ogsantos/eulabs-api/application/repositories"
	"github.com/rafael-ogsantos/eulabs-api/application/services"
	"github.com/rafael-ogsantos/eulabs-api/domain"
)

type RouterInterface interface {
	Router() *echo.Echo
}

type Router struct {
	e                 *echo.Echo
	ProductService    services.ProductServiceInterface
	ProductRepository repositories.ProductRepository
}

func NewRouter(
	e *echo.Echo,
	service services.ProductServiceInterface,
	repository repositories.ProductRepository,
) RouterInterface {
	return &Router{
		ProductService:    service,
		ProductRepository: repository,
	}
}

// New returns a new echo instance
func (r *Router) Router() *echo.Echo {
	if r.e == nil {
		r.e = echo.New()
	}

	g := r.e.Group("/api")

	// FindById
	g.GET("/product/:id", func(c echo.Context) error {
		id := c.Param("id")
		ctx := c.Request().Context()

		product, err := r.ProductService.FindById(ctx, id)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, product)
	})

	// Create
	g.POST("/product", func(c echo.Context) (err error) {
		p := new(domain.Product)
		ctx := c.Request().Context()

		if err = c.Bind(p); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		product, err := r.ProductService.Insert(ctx, p)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, product)
	})

	// Update
	g.PUT("/product/:id", func(c echo.Context) error {
		p := new(domain.Product)
		id := c.Param("id")
		ctx := c.Request().Context()

		if err := c.Bind(p); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		product, err := r.ProductService.Update(ctx, id, p)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, product)
	})

	// Delete
	g.DELETE("/product/:id", func(c echo.Context) error {
		id := c.Param("id")
		ctx := c.Request().Context()

		err := r.ProductService.Delete(ctx, id)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	})

	return r.e
}
