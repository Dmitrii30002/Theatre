package domain

type Route struct {
	Prefix      string
	Target      string
	StripPrefix bool
}
