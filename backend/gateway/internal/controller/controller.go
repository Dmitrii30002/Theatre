package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gateway/internal/service"
)

type Controller struct {
	theatre service.Theatre
}

func New(theatre service.Theatre) *Controller {
	return &Controller{theatre: theatre}
}

func (c *Controller) Register(e *echo.Echo) {
	e.GET("/health", c.Health)
	e.GET("/ready", c.Ready)
	e.GET("/api/spectacles", c.GetSpectacles)
	e.GET("/api/spectacles/:id", c.GetSpectacleByID)
	e.GET("/api/shows/:id", c.GetShowByID)
}

func (c *Controller) Health(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func (c *Controller) Ready(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"status": "ready"})
}

func (c *Controller) GetSpectacles(ctx echo.Context) error {
	page, err := queryInt32(ctx, "page", 0)
	if err != nil {
		return err
	}
	size, err := queryInt32(ctx, "size", 20)
	if err != nil {
		return err
	}
	if page < 0 || size <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "page must be non-negative and size must be positive")
	}

	response, err := c.theatre.GetSpectacles(ctx.Request().Context(), page, size)
	if err != nil {
		return grpcError(err)
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c *Controller) GetSpectacleByID(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "spectacle id must be a positive integer")
	}

	response, err := c.theatre.GetSpectacleByID(ctx.Request().Context(), id)
	if err != nil {
		return grpcError(err)
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c *Controller) GetShowByID(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "show id is required")
	}

	response, err := c.theatre.GetShowByID(ctx.Request().Context(), id)
	if err != nil {
		return grpcError(err)
	}
	return ctx.JSON(http.StatusOK, response)
}

func queryInt32(ctx echo.Context, name string, defaultValue int32) (int32, error) {
	value := ctx.QueryParam(name)
	if value == "" {
		return defaultValue, nil
	}
	parsed, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0, echo.NewHTTPError(http.StatusBadRequest, name+" must be an integer")
	}
	return int32(parsed), nil
}

func grpcError(err error) error {
	code := http.StatusBadGateway
	switch status.Code(err) {
	case codes.InvalidArgument:
		code = http.StatusBadRequest
	case codes.NotFound:
		code = http.StatusNotFound
	case codes.Unauthenticated:
		code = http.StatusUnauthorized
	case codes.PermissionDenied:
		code = http.StatusForbidden
	}
	return echo.NewHTTPError(code, status.Convert(err).Message()).SetInternal(err)
}
