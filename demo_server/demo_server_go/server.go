

package main
import (
    "fmt"
    "net/http"
    "errors"
	"html/template"
	"encoding/json"

)

type Conf struct {
	ClientId     string // Client ID
	ClientSecret string // Client Secret
	RedirectUrl  string // Authorization callback URL
}

var conf = Conf{
	ClientId:     "",  	// fill in with your id before test
	ClientSecret: "",   // fill in with your secret before test
	RedirectUrl:  "http://localhost:3000/login",
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"` 
	Scope       string `json:"scope"`     
}

func Hello(w http.ResponseWriter, r *http.Request) {

	var temp *template.Template
	var err error
	if temp, err = template.ParseFiles("test-frontend.html"); err != nil {
		fmt.Println("read frontend failed, error:", err)
		return
	}


	if err = temp.Execute(w, conf); err != nil {
		fmt.Println("read html page failed, error:", err)
		return
	}
}


func GetTokenAuthUrl(code string) string {
	return fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		conf.ClientId, conf.ClientSecret, code,
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


func Oauth(w http.ResponseWriter, r *http.Request) {

	var err error
	
	// get code
	var code = r.URL.Query().Get("code")

	// get token
	var tokenAuthUrl = GetTokenAuthUrl(code)
	var token *Token	
	if token, err = GetToken(tokenAuthUrl); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v",token)
	
}


func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/login", Oauth)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("listening error:", err)
		return
	}
}