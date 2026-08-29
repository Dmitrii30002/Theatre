package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"gateway/internal/service"
)

type Controller struct {
	proxy *service.ProxyService
}

func New(proxy *service.ProxyService) *Controller {
	return &Controller{proxy: proxy}
}

func (c *Controller) Register(e *echo.Echo) {
	e.GET("/health", c.Health)
	e.GET("/ready", c.Ready)
	e.Any("/*", c.Proxy)
}

func (c *Controller) Health(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func (c *Controller) Ready(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"status": "ready"})
}

func (c *Controller) Proxy(ctx echo.Context) error {
	if err := c.proxy.Forward(ctx.Response().Writer, ctx.Request()); err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, "upstream unavailable").SetInternal(err)
	}
	return nil
}
