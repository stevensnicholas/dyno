package main
import (
    "fmt"
    "net/http"
    "net/url"
    "io/ioutil"
    "bytes"
    "errors"

)

func GetGitHubOauthToken(c_id string, c_secret string) (string, error) {
	const rootURl = "https://github.com/login/oauth/access_token"

	values := url.Values{}
	// values.Add("code", code)
	values.Add("client_id", c_id)
	values.Add("client_secret", c_secret)

	query := values.Encode()

	queryString := fmt.Sprintf("%s?%s", rootURl, bytes.NewBufferString(query))
	req, err := http.NewRequest("POST", queryString, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", errors.New("could not retrieve token")
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	parsedQuery, err := url.ParseQuery(string(resBody))
	if err != nil {
		return "", err
	}

	tokenBody := parsedQuery["access_token"][0]

	return tokenBody, nil
}

func main() {

    var GITHUB_OAUTH_CLIENT_ID = 
    var GITHUB_OAUTH_CLIENT_SECRET = 

    token, erro := GetGitHubOauthToken(GITHUB_OAUTH_CLIENT_ID, GITHUB_OAUTH_CLIENT_SECRET)
    fmt.Println(token, erro)
}