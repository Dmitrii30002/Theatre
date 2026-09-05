package theatregrpc

import (
	"context"

	"gateway/internal/domain"
	theatre "gateway/internal/gen/theatre"

	"google.golang.org/grpc"
)

type Client struct {
	spectacles theatre.SpectacleServiceClient
	shows      theatre.ShowServiceClient
}

func NewClient(conn grpc.ClientConnInterface) *Client {
	return &Client{
		spectacles: theatre.NewSpectacleServiceClient(conn),
		shows:      theatre.NewShowServiceClient(conn),
	}
}

func (c *Client) GetSpectacles(ctx context.Context, page, size int32) (domain.Spectacles, error) {
	response, err := c.spectacles.GetSpectacles(ctx, &theatre.GetSpectaclesRequest{Page: page, Size: size})
	if err != nil {
		return domain.Spectacles{}, err
	}

	items := make([]domain.SpectaclePreview, 0, len(response.GetSpectacles()))
	for _, spectacle := range response.GetSpectacles() {
		items = append(items, domain.SpectaclePreview{
			ID:         spectacle.GetId(),
			Name:       spectacle.GetName(),
			PreviewURL: spectacle.GetPreviewUrl(),
		})
	}
	return domain.Spectacles{Spectacles: items}, nil
}

func (c *Client) GetSpectacleByID(ctx context.Context, id int64) (domain.Spectacle, error) {
	response, err := c.spectacles.GetSpectacleById(ctx, &theatre.GetSpectacleByIdRequest{Id: id})
	if err != nil {
		return domain.Spectacle{}, err
	}

	var theatreInfo *domain.Theatre
	if response.GetTheatre() != nil {
		theatreInfo = &domain.Theatre{ID: response.GetTheatre().GetId(), Name: response.GetTheatre().GetName()}
	}
	images := make([]domain.Image, 0, len(response.GetImages()))
	for _, image := range response.GetImages() {
		images = append(images, domain.Image{URL: image.GetUrl()})
	}
	shows := make([]domain.ShowPreview, 0, len(response.GetShows()))
	for _, show := range response.GetShows() {
		shows = append(shows, domain.ShowPreview{ID: show.GetId(), PlatformName: show.GetPlatformName(), Date: show.GetDate()})
	}

	return domain.Spectacle{
		ID: response.GetId(), Name: response.GetName(), PreviewURL: response.GetPreviewUrl(),
		Description: response.GetDescription(), AgeLimit: response.GetAgeLimit(), Duration: response.GetDurationMinutes(),
		Genre: response.GetGenre(), PushkinCard: response.GetPushkinCard(), Theatre: theatreInfo, Images: images, Shows: shows,
	}, nil
}

func (c *Client) GetShowByID(ctx context.Context, id string) (domain.Show, error) {
	response, err := c.shows.GetShowById(ctx, &theatre.GetShowByIdRequest{Id: id})
	if err != nil {
		return domain.Show{}, err
	}

	seats := make([]domain.Seat, 0, len(response.GetSeats()))
	for _, seat := range response.GetSeats() {
		seats = append(seats, domain.Seat{
			ID: seat.GetId(), Row: seat.GetRow(), Number: seat.GetNumber(), Color: seat.GetColor(),
			Price: seat.GetPrice(), Status: seat.GetStatus(),
		})
	}
	return domain.Show{ID: response.GetId(), Scheme: response.GetScheme(), PlatformName: response.GetPlatformName(), Date: response.GetDate(), Seats: seats}, nil
}
