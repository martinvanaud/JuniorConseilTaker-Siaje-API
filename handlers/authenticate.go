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

	session, _ := server.Store.Get(context.Request, "je-session")

	siajeAuthenticationEndpoint := "https://pro.siaje.com/jctaker/connexion.php?r=index.php"
	authorized_url := "https://pro.siaje.com/jctaker/index.php"

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

	authentication_successful := false

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) > 1 {
				if req.Header.Get("Referer") == authorized_url {
					authentication_successful = true
					return http.ErrUseLastResponse
				}
			}
			return nil
		},
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

	for _, cookie := range res.Cookies() {
		if strings.Compare(cookie.Name, "PHPSESSID") == 0 {
			session.Values["PHPSESSID"] = cookie.Value
			session.Save(context.Request, context.Writer)
		}
	}

	if authentication_successful {
		context.JSON(http.StatusOK, gin.H{
			"code":  http.StatusOK,
			"token": session.Values["PHPSESSID"],
		})
	} else {
		context.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Invalid credentials",
		})
	}
}
