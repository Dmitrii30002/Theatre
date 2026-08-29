package service

import (
	"testing"

	"gateway/internal/domain"
)

func TestProxyServiceMatchesLongestPrefix(t *testing.T) {
	service, err := NewProxyService([]domain.Route{
		{Prefix: "/api", Target: "http://localhost:8080"},
		{Prefix: "/api/users", Target: "http://localhost:8081"},
	})
	if err != nil {
		t.Fatalf("NewProxyService() error = %v", err)
	}

	tests := []struct {
		name   string
		path   string
		prefix string
	}{
		{name: "specific route", path: "/api/users/42", prefix: "/api/users"},
		{name: "generic route", path: "/api/orders", prefix: "/api"},
		{name: "no partial match", path: "/apix", prefix: ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			matched := service.match(test.path)
			if test.prefix == "" {
				if matched != nil {
					t.Fatalf("match(%q) = %q, want no match", test.path, matched.route.Prefix)
				}
				return
			}
			if matched == nil || matched.route.Prefix != test.prefix {
				t.Fatalf("match(%q) = %v, want prefix %q", test.path, matched, test.prefix)
			}
		})
	}
}
