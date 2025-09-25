package rest

import "github.com/Sakrafux/stack-experiment-monorepo/internal/domain/profile"

type ProfileResponse struct {
	Profile *Profile `json:"profile"`
}

type Profile struct {
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Image     string `json:"image"`
	Following bool   `json:"following"`
}

func toProfile(profile *profile.Profile) *Profile {
	return &Profile{
		Username:  profile.Username,
		Bio:       profile.Bio,
		Image:     profile.Image,
		Following: profile.Following,
	}
}
