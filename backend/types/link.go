package types

type LinkPostRequest struct {
	FullLink string `json:"full_link"`
}

type LinkPostResponse struct {
	Code string `json:"code"`
}

type LinkGetResponse struct {
	FullLink string `json:"full_link"`
}
