package funtify

type BasePage struct {
	Endpoint  string `json:"href"`
	Limit     int    `json:"limit"`
	Nextpoint string `json:"next"`
	Offset    int    `json:"offset"`
	Prevpoint string `json:"previous"`
	Count     int    `json:"total"`
}

type AlbumPage struct {
	BasePage
	Albums []Album `json:"items"`
}
