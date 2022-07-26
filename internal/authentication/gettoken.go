package authentication

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"io/ioutil"
)
type GitHubUserInfo struct {
	Name  string
	Photo string
	// Email string
	ID    float64
}

type Conf struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

var conf = Conf{
	ClientID:     "94143fe4a712d77c2983", // fill in with your id before test
	ClientSecret: "ce8c95475b2874e2204b3f878f3ebedb2a2320dd", // fill in with your secret before test
	RedirectURL:  "http://localhost:8080/login",
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

func GetTokenAuthURL(code string) string {
	return fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		conf.ClientID, conf.ClientSecret, code,
	)
}

func GetToken(url string) (*Token, error) {
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
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

	var userInfoUrl = "https://api.github.com/user"	
	var req *http.Request
	var err error

	if req, err = http.NewRequest(http.MethodGet, userInfoUrl, nil); err != nil {
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

	// var userInfo = make(map[string]interface{})
	// if err = json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
	// 	return nil, err
	// }

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var Info map[string]interface{}
	
	if err := json.Unmarshal(resBody, &Info); err != nil {
		return nil, err
	}

	userinfo := &GitHubUserInfo{
		// Email: Info["email"].(string),
		Name:  Info["login"].(string),
		Photo: Info["avatar_url"].(string),
		ID:    Info["id"].(float64),
	}

	return userinfo, nil
}