package authentication
import (
    "fmt"
    "net/http"
    "errors"
	"encoding/json"

)

type Conf struct {
	ClientID     string 
	ClientSecret string 
	RedirectURL  string 
}

var conf = Conf{
	ClientID:     "94143fe4a712d77c2983",  	// fill in with your id before test
	ClientSecret: "0eaa973b77bf15fd7e431ba167195b8e3a1818bc",   // fill in with your secret before test
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
