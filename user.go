package main

type User struct {
	handle     string `json: handle`
	imageURL   string `json: image_url`
	profileURL string `json: profile_url`
	name       string `json: screen_name`
}
