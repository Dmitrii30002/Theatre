package domain

type Spectacles struct {
	Spectacles []SpectaclePreview `json:"spectacles"`
}

type SpectaclePreview struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	PreviewURL string `json:"preview_url"`
}

type Spectacle struct {
	ID          int64         `json:"id"`
	Name        string        `json:"name"`
	PreviewURL  string        `json:"preview_url"`
	Description string        `json:"description,omitempty"`
	AgeLimit    int32         `json:"age_limit,omitempty"`
	Duration    int32         `json:"duration_minutes,omitempty"`
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
	ID     int64   `json:"id"`
	Row    int32   `json:"row"`
	Number int32   `json:"number"`
	Color  string  `json:"color"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}
