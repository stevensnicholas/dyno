package authentication

import (
	"dyno/internal/logger"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GitHubUserInfo struct {
	Name  string
	Photo string
	Email string
	ID    float64
}

type Conf struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

var conf = Conf{
	ClientID:     "", // fill in with your id before test
	ClientSecret: "", // fill in with your secret before test
	RedirectURL:  "",
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

func GetTokenAuthURL(code string) (string, error) {
	ClientID, err := getSSMParameterValue("client_id")
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	ClientSecret, err := getSSMParameterValue("client_secret")
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		ClientID, ClientSecret, code,
	), nil
}

func GetToken(url string) (*Token, error) {
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodPost, url, nil); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")

	var httpClient = http.Client{}
	var res *http.Response
	if res, err = httpClient.Do(req); err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve token")
	}

	var token Token
	if err = json.NewDecoder(res.Body).Decode(&token); err != nil {
		return nil, err
	}
	return &token, nil
}

func GetUserInfo(token *Token) (*GitHubUserInfo, error) {
	var userInfoURL = "https://api.github.com/user"
	var req *http.Request
	var err error

	if req, err = http.NewRequest(http.MethodGet, userInfoURL, nil); err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	var client = http.Client{}
	var res *http.Response

	if res, err = client.Do(req); err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve user")
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var Info map[string]interface{}

	if err := json.Unmarshal(resBody, &Info); err != nil {
		return nil, err
	}

	userinfo := &GitHubUserInfo{
		Email: Info["email"].(string),
		Name:  Info["login"].(string),
		Photo: Info["avatar_url"].(string),
		ID:    Info["id"].(float64),
	}

	return userinfo, nil
}
