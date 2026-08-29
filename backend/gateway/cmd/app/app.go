package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"gateway/internal/config"
	"gateway/internal/controller"
	"gateway/internal/domain"
	"gateway/internal/service"
)

func Run() error {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "config/config.yaml"
	}

	cfg, err := config.Load(configPath)
	if err != nil {
		return err
	}

	routes := make([]domain.Route, 0, len(cfg.Routes))
	for _, route := range cfg.Routes {
		routes = append(routes, domain.Route{
			Prefix:      route.Prefix,
			Target:      route.Target,
			StripPrefix: route.StripPrefix,
		})
	}
	proxyService, err := service.NewProxyService(routes)
	if err != nil {
		return err
	}

	e := echo.New()
	e.HideBanner = true
	e.Logger.SetOutput(os.Stdout)
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.HTTPErrorHandler = func(err error, ctx echo.Context) {
		code := echo.ErrInternalServerError.Code
		message := interface{}("internal server error")
		if httpError, ok := err.(*echo.HTTPError); ok {
			code = httpError.Code
			message = httpError.Message
			if httpError.Internal != nil {
				logger.Error("request failed", "error", httpError.Internal, "path", ctx.Path())
			}
		}
		_ = ctx.JSON(code, map[string]interface{}{"error": message})
	}

	controller.New(proxyService).Register(e)

	server := &httpServer{
		echo:         e,
		address:      cfg.Server.Address,
		readTimeout:  config.ParseDuration(cfg.Server.ReadTimeout, 10*time.Second),
		writeTimeout: config.ParseDuration(cfg.Server.WriteTimeout, 15*time.Second),
		idleTimeout:  config.ParseDuration(cfg.Server.IdleTimeout, 60*time.Second),
	}

	go func() {
		logger.Info("gateway started", "address", server.address, "config", configPath)
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			logger.Error("gateway stopped unexpectedly", "error", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	shutdownTimeout := config.ParseDuration(cfg.Server.ShutdownTimeout, 10*time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	logger.Info("shutting down gateway")
	return server.Shutdown(ctx)
}

type httpServer struct {
	echo         *echo.Echo
	address      string
	readTimeout  time.Duration
	writeTimeout time.Duration
	idleTimeout  time.Duration
}

func (s *httpServer) Start() error {
	s.echo.Server.ReadTimeout = s.readTimeout
	s.echo.Server.WriteTimeout = s.writeTimeout
	s.echo.Server.IdleTimeout = s.idleTimeout
	return s.echo.Start(s.address)
}

func (s *httpServer) Shutdown(ctx context.Context) error {
	if err := s.echo.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdown gateway: %w", err)
	}
	return nil
}
