package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"juniorconseiltaker-siaje-api/server"
	"net/http"
	"net/url"
	"strings"
)

type AuthenticationPayload struct {
	Login        string `json:"login"`
	Password     string `json:"mot_de_passe"`
	Redirect     string `json:"r"`
	SubConnexion string `json:"sub_connexion"`
}

func Authenticate(context *gin.Context) {

	siajeAuthenticationEndpoint := "https://pro.siaje.com/jctaker/connexion.php?r=index.php"

	const authorizedUrl string = "https://pro.siaje.com/jctaker/index.php"
	const unauthorizedUrl string = "https://pro.siaje.com/jctaker/connexion.php?m=error"

	method := "POST"

	var authenticationRequestPayload AuthenticationPayload

	err := json.NewDecoder(context.Request.Body).Decode(&authenticationRequestPayload)
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	apiRequestPayload := url.Values{
		"login":         {authenticationRequestPayload.Login},
		"mot_de_passe":  {authenticationRequestPayload.Password},
		"r":             {authenticationRequestPayload.Redirect},
		"sub_connexion": {authenticationRequestPayload.SubConnexion},
	}

	payload := strings.NewReader(apiRequestPayload.Encode())

	authenticationSuccessful := false

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {

			forwardUrl := req.URL.String()

			fmt.Println("forwarding to: ", forwardUrl)

			if forwardUrl == authorizedUrl {
				authenticationSuccessful = true
				return http.ErrUseLastResponse
			} else if forwardUrl == unauthorizedUrl {
				return http.ErrUseLastResponse
			}
			return nil
		},
		Jar: server.CookieJar,
	}

	req, err := http.NewRequest(method, siajeAuthenticationEndpoint, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	var token string

	for _, cookie := range res.Cookies() {
		fmt.Println("cookie:", cookie)
		if strings.Compare(cookie.Name, "PHPSESSID") == 0 {
			token = cookie.Value

			urlObj, _ := url.Parse("http://localhost:8018/")
			client.Jar.SetCookies(urlObj, []*http.Cookie{cookie})

		}
	}

	if authenticationSuccessful {
		context.JSON(http.StatusOK, gin.H{
			"code":  http.StatusOK,
			"token": token,
		})
	} else {
		context.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Invalid credentials",
		})
	}
}
