package service

import (
	"context"

	"gateway/internal/domain"
)

type Theatre interface {
	GetSpectacles(context.Context, int32, int32) (domain.Spectacles, error)
	GetSpectacleByID(context.Context, int64) (domain.Spectacle, error)
	GetShowByID(context.Context, string) (domain.Show, error)
}
