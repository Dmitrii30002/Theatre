package controller

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"gateway/internal/domain"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type fakeTheatre struct {
	id   int64
	show string
}

func (f *fakeTheatre) GetSpectacles(context.Context, int32, int32) (domain.Spectacles, error) {
	return domain.Spectacles{}, errors.New("not implemented")
}

func (f *fakeTheatre) GetSpectacleByID(_ context.Context, id int64) (domain.Spectacle, error) {
	f.id = id
	return domain.Spectacle{ID: id, Name: "Hamlet"}, nil
}

func (f *fakeTheatre) GetShowByID(_ context.Context, id string) (domain.Show, error) {
	f.show = id
	return domain.Show{ID: id}, status.Error(codes.NotFound, "show not found")
}

func TestGetSpectacleByIDForwardsID(t *testing.T) {
	theatre := &fakeTheatre{}
	e := echo.New()
	New(theatre).Register(e)

	record := httptest.NewRecorder()
	e.ServeHTTP(record, httptest.NewRequest(http.MethodGet, "/api/spectacles/42", nil))

	if record.Code != http.StatusOK || theatre.id != 42 {
		t.Fatalf("status = %d, forwarded id = %d", record.Code, theatre.id)
	}
}

func TestGetShowByIDMapsNotFound(t *testing.T) {
	theatre := &fakeTheatre{}
	e := echo.New()
	New(theatre).Register(e)

	record := httptest.NewRecorder()
	e.ServeHTTP(record, httptest.NewRequest(http.MethodGet, "/api/shows/show-42", nil))

	if record.Code != http.StatusNotFound || theatre.show != "show-42" {
		t.Fatalf("status = %d, forwarded id = %q", record.Code, theatre.show)
	}
}
