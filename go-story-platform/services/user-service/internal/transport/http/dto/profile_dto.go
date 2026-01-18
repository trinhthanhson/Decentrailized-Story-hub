package dto

type ProfileResponse struct {
	Avatar      string `json:"avatar"`
	Bio         string `json:"bio"`
	Preferences string `json:"preferences"`
}

type ProfileRequest struct {
	Avatar string `json:"avatar"`
	Bio    string `json:"bio"`
}

type UpdateProfileRequest struct {
	Avatar      string `json:"avatar"`
	Bio         string `json:"bio"`
	Preferences string `json:"preferences"`
}
