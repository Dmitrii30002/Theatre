package service

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestMockServiceRoutes(t *testing.T) {
	service := NewMockService()
	e := echo.New()

	t.Run("spectacles list", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/spectacles", nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		if err := service.Handle(ctx); err != nil {
			t.Fatalf("Handle() error = %v", err)
		}

		if rec.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
		}
		if !strings.Contains(rec.Body.String(), "\"spectacles\"") {
			t.Fatalf("body = %q, want spectacles payload", rec.Body.String())
		}
	})

	t.Run("spectacle by id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/spectacles/1", nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		if err := service.Handle(ctx); err != nil {
			t.Fatalf("Handle() error = %v", err)
		}

		if rec.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
		}
		if !strings.Contains(rec.Body.String(), "\"name\":\"Гамлет\"") {
			t.Fatalf("body = %q, want Hamlet payload", rec.Body.String())
		}
	})

	t.Run("show by id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/shows/show_001", nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		if err := service.Handle(ctx); err != nil {
			t.Fatalf("Handle() error = %v", err)
		}

		if rec.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
		}
		if !strings.Contains(rec.Body.String(), "\"id\":\"show_001\"") {
			t.Fatalf("body = %q, want show_001 payload", rec.Body.String())
		}
	})
}
