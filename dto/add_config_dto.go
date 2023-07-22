package dto

type EditConfigRequest struct {
	Id        int64     `param:"id"`
	CacheKey  string    `json:"cacheKey"`
	RateLimit rateLimit `json:"rateLimit"`
}

type rateLimit struct {
	Zone          string
	Burst         int
	RatePerSecond int64
	MaxSize       string
	path          string
}
