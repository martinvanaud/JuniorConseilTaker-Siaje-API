package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

func Authenticate(context *gin.Context) {

	url := "https://pro.siaje.com/jctaker/connexion.php?r=index.php"
	authorized_url := "https://pro.siaje.com/jctaker/index.php"

	method := "POST"

	payload := strings.NewReader("login=m-vanaud&mot_de_passe=TBjR9fu5&r=index.php&sub_connexion=Se%20connecter%20%E2%86%92")

	authentication_successful := false

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) > 1 {
				fmt.Println(via[0].Header.Get("Location"))
				if req.Header.Get("Referer") == authorized_url {
					authentication_successful = true
					return http.ErrUseLastResponse
				}
			}
			return nil
		},
	}
	req, err := http.NewRequest(method, url, payload)

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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

	var siajeToken string
	for _, cookie := range res.Cookies() {
		if strings.Compare(cookie.Name, "PHPSESSID") == 0 {
			siajeToken = cookie.Value
		}
	}

	if authentication_successful {
		context.JSON(http.StatusOK, gin.H{
			"code":  http.StatusOK,
			"token": siajeToken,
		})
	} else {
		context.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Invalid credentials",
		})
	}
}
