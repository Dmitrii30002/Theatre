package service

import (
	"context"
	"testing"

	"gateway/internal/domain"
)

type fakeTheatre struct{}

func (fakeTheatre) GetSpectacles(context.Context, int32, int32) (domain.Spectacles, error) {
	return domain.Spectacles{Spectacles: []domain.SpectaclePreview{{ID: 1, Name: "Hamlet"}}}, nil
}

func (fakeTheatre) GetSpectacleByID(context.Context, int64) (domain.Spectacle, error) {
	return domain.Spectacle{ID: 1, Name: "Hamlet"}, nil
}

func (fakeTheatre) GetShowByID(context.Context, string) (domain.Show, error) {
	return domain.Show{ID: "show-1"}, nil
}

func TestFakeTheatreImplementsPort(t *testing.T) {
	var _ Theatre = fakeTheatre{}
}
