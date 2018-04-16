package main

import (
	"fmt"
	"os"
)

const (
	ENV_MATTERMOST_BASE_URL   = "BASE_URL"
	ENV_PERSONAL_ACCESS_TOKEN = "TOKEN"
	ENV_USERNAME              = "USERNAME"
	ENV_PASSWORD              = "PASSWORD"
)

type ClientConfig struct {
	MattermostBaseURL   string
	PersonalAccessToken string
	Username            string
	Password            string
}

func LoadFromEnv() (*ClientConfig, error) {
	config := &ClientConfig{
		MattermostBaseURL:   os.Getenv(ENV_MATTERMOST_BASE_URL),
		PersonalAccessToken: os.Getenv(ENV_PERSONAL_ACCESS_TOKEN),
		Username:            os.Getenv(ENV_USERNAME),
		Password:            os.Getenv(ENV_PASSWORD),
	}

	hasBaseURL := len(config.MattermostBaseURL) > 0
	if !hasBaseURL {
		return nil, fmt.Errorf("Env %v is required", ENV_MATTERMOST_BASE_URL)
	}

	hasCredential := len(config.Username) > 0 && len(config.Password) > 0
	hasPAT := len(config.PersonalAccessToken) > 0
	if !hasCredential && !hasPAT {
		return nil, fmt.Errorf("Env %v & %v or, %v is required", ENV_USERNAME, ENV_PASSWORD, ENV_PERSONAL_ACCESS_TOKEN)
	}

	return config, nil
}
