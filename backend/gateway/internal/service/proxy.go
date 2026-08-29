package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"

	"gateway/internal/domain"

	"github.com/labstack/echo/v4"
)

type proxyRoute struct {
	route domain.Route
	proxy *httputil.ReverseProxy
}

type ProxyService struct {
	routes []proxyRoute
}

func NewProxyService(routes []domain.Route) (*ProxyService, error) {
	proxyRoutes := make([]proxyRoute, 0, len(routes))
	for _, route := range routes {
		target, err := url.Parse(route.Target)
		if err != nil || target.Scheme == "" || target.Host == "" {
			return nil, fmt.Errorf("invalid target %q", route.Target)
		}

		proxy := httputil.NewSingleHostReverseProxy(target)
		originalDirector := proxy.Director
		originalDirectorWithPrefix := func(request *http.Request) {
			originalPath := request.URL.Path
			originalDirector(request)
			if route.StripPrefix {
				request.URL.Path = strings.TrimPrefix(originalPath, route.Prefix)
				if request.URL.Path == "" {
					request.URL.Path = "/"
				}
			}
		}
		proxy.Director = originalDirectorWithPrefix
		proxyRoutes = append(proxyRoutes, proxyRoute{route: route, proxy: proxy})
	}

	return &ProxyService{routes: proxyRoutes}, nil
}

func NewMockService() *ProxyService {
	return &ProxyService{}
}

func (s *ProxyService) Handle(ctx echo.Context) error {
	return s.handleMock(ctx.Response(), ctx.Request())
}

func (s *ProxyService) Forward(writer http.ResponseWriter, request *http.Request) error {
	if err := s.handleMock(writer, request); err == nil {
		return nil
	}
	matched := s.match(request.URL.Path)
	if matched == nil {
		return fmt.Errorf("no upstream route for %s", request.URL.Path)
	}
	matched.proxy.ServeHTTP(writer, request)
	return nil
}

func (s *ProxyService) handleMock(writer http.ResponseWriter, request *http.Request) error {
	path := strings.TrimSpace(request.URL.Path)
	if path == "" || path == "/" {
		return fmt.Errorf("empty path")
	}
	path = strings.TrimPrefix(path, "/")

	if request.Method != http.MethodGet {
		return fmt.Errorf("unsupported method %s", request.Method)
	}

	switch {
	case path == "api/spectacles":
		writeJSON(writer, http.StatusOK, mockSpectaclesResponse)
		return nil
	case strings.HasPrefix(path, "api/spectacles/"):
		idStr := strings.TrimPrefix(path, "api/spectacles/")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return writeJSONError(writer, http.StatusBadRequest, "INVALID_ID", "Spectacle id must be an integer")
		}
		if spectacle, ok := mockSpectacleByID(id); ok {
			writeJSON(writer, http.StatusOK, spectacle)
			return nil
		}
		return writeJSONError(writer, http.StatusNotFound, "NOT_FOUND", "Spectacle not found")
	case strings.HasPrefix(path, "api/shows/"):
		id := strings.TrimPrefix(path, "api/shows/")
		if show, ok := mockShowByID(id); ok {
			writeJSON(writer, http.StatusOK, show)
			return nil
		}
		return writeJSONError(writer, http.StatusNotFound, "NOT_FOUND", "Show not found")
	default:
		return fmt.Errorf("no mock response for %s", request.URL.Path)
	}
}

func writeJSON(writer http.ResponseWriter, status int, payload interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	_ = json.NewEncoder(writer).Encode(payload)
}

func writeJSONError(writer http.ResponseWriter, status int, code, message string) error {
	writeJSON(writer, status, map[string]string{"code": code, "message": message})
	return nil
}

func (s *ProxyService) match(path string) *proxyRoute {
	var result *proxyRoute
	for index := range s.routes {
		prefix := s.routes[index].route.Prefix
		if prefix == "/" || path == prefix || strings.HasPrefix(path, prefix+"/") {
			if result == nil || len(prefix) > len(result.route.Prefix) {
				result = &s.routes[index]
			}
		}
	}
	return result
}

type GetSpectaclesResponse struct {
	Spectacles []SpectacleSummary `json:"spectacles"`
}

type SpectacleSummary struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	PreviewURL string `json:"preview_url"`
}

type Spectacle struct {
	ID          int64         `json:"id"`
	Name        string        `json:"name"`
	PreviewURL  string        `json:"preview_url"`
	Description string        `json:"description,omitempty"`
	AgeLimit    int           `json:"age_limit,omitempty"`
	Duration    int           `json:"duration_minutes,omitempty"`
	Genre       string        `json:"genre,omitempty"`
	PushkinCard bool          `json:"pushkin_card,omitempty"`
	Theatre     *Theatre      `json:"theatre,omitempty"`
	Images      []Image       `json:"images,omitempty"`
	Shows       []ShowPreview `json:"shows,omitempty"`
}

type Theatre struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Image struct {
	URL string `json:"url"`
}

type ShowPreview struct {
	ID           string `json:"id"`
	PlatformName string `json:"platform_name"`
	Date         string `json:"date"`
}

type Show struct {
	ID           string `json:"id"`
	Scheme       string `json:"scheme"`
	PlatformName string `json:"platform_name"`
	Date         string `json:"date"`
	Seats        []Seat `json:"seats"`
}

type Seat struct {
	ID     int64  `json:"id"`
	Row    int    `json:"row"`
	Number int    `json:"number"`
	Color  string `json:"color"`
	Price  int    `json:"price"`
	Status string `json:"status"`
}

var mockSpectaclesResponse = GetSpectaclesResponse{
	Spectacles: []SpectacleSummary{
		{ID: 1, Name: "Гамлет", PreviewURL: "https://api.sourcesplash.com/i/random?q=architecture&w=800&h=600"},
		{ID: 2, Name: "Щелкунчик", PreviewURL: "https://api.sourcesplash.com/i/random?q=architecture&w=800&h=600"},
		{ID: 3, Name: "Мастер и Маргарита", PreviewURL: "https://api.sourcesplash.com/i/random?q=architecture&w=800&h=600"},
		{ID: 4, Name: "Гамлет2", PreviewURL: "https://api.sourcesplash.com/i/random?q=architecture&w=800&h=600"},
		{ID: 5, Name: "Щелкунчик2", PreviewURL: "https://api.sourcesplash.com/i/random?q=architecture&w=800&h=600"},
		{ID: 6, Name: "Мастер и Маргарита2", PreviewURL: "https://api.sourcesplash.com/i/random?q=architecture&w=800&h=600"},
	},
}

var mockSpectacleDetail = Spectacle{
	ID:          1,
	Name:        "Гамлет",
	PreviewURL:  "https://api.sourcesplash.com/i/random?q=architecture&w=800&h=600",
	Description: "Трагедия Уильяма Шекспира о принце Датском, стоящем перед выбором между местью, долгом и собственной совестью.",
	AgeLimit:    16,
	Duration:    180,
	Genre:       "Драма",
	PushkinCard: true,
	Theatre: &Theatre{
		ID:   1,
		Name: "Большой театр",
	},
	Images: []Image{
		{URL: "https://api.sourcesplash.com/i/random?q=architecture&w=800&h=600"},
		{URL: "https://api.sourcesplash.com/i/random?q=architecture&w=800&h=600"},
	},
	Shows: []ShowPreview{
		{ID: "show_001", PlatformName: "Большой зал", Date: "2026-08-25T19:00:00+03:00"},
		{ID: "show_002", PlatformName: "Большой зал", Date: "2026-08-27T19:00:00+03:00"},
		{ID: "show_003", PlatformName: "Большой зал", Date: "2026-08-30T14:00:00+03:00"},
	},
}

var mockShowDetail = Show{
	ID:           "show_001",
	Scheme:       "СХЕМА ПОСАДКИ",
	PlatformName: "Большой зал",
	Date:         "2026-08-25T19:00:00+03:00",
	Seats: []Seat{
		{ID: 10001, Row: 1, Number: 1, Color: "#4CAF50", Price: 5000, Status: "AVAILABLE"},
		{ID: 10002, Row: 1, Number: 2, Color: "#4CAF50", Price: 5000, Status: "SOLD"},
		{ID: 10003, Row: 1, Number: 3, Color: "#4CAF50", Price: 5000, Status: "AVAILABLE"},
		{ID: 10004, Row: 1, Number: 4, Color: "#4CAF50", Price: 5000, Status: "HELD"},
		{ID: 10005, Row: 1, Number: 5, Color: "#4CAF50", Price: 5000, Status: "AVAILABLE"},
		{ID: 10006, Row: 2, Number: 1, Color: "#2196F3", Price: 3500, Status: "AVAILABLE"},
		{ID: 10007, Row: 2, Number: 2, Color: "#2196F3", Price: 3500, Status: "AVAILABLE"},
		{ID: 10008, Row: 2, Number: 3, Color: "#2196F3", Price: 3500, Status: "SOLD"},
		{ID: 10009, Row: 2, Number: 4, Color: "#2196F3", Price: 3500, Status: "AVAILABLE"},
		{ID: 10010, Row: 2, Number: 5, Color: "#2196F3", Price: 3500, Status: "BLOCKED"},
	},
}

func mockSpectacleByID(id int64) (Spectacle, bool) {
	if id == mockSpectacleDetail.ID {
		return mockSpectacleDetail, true
	}
	return Spectacle{}, false
}

func mockShowByID(id string) (Show, bool) {
	if id == mockShowDetail.ID {
		return mockShowDetail, true
	}
	return Show{}, false
}
