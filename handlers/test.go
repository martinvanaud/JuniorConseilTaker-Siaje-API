package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"juniorconseiltaker-siaje-api/server"
)

//type AuthenticationPayload struct {
//	Login        string `json:"login"`
//	Password     string `json:"mot_de_passe"`
//	Redirect     string `json:"r"`
//	SubConnexion string `json:"sub_connexion"`
//}

func Test(context *gin.Context) {

	//var authenticationRequestPayload AuthenticationPayload
	//err := json.NewDecoder(context.Request.Body).Decode(&authenticationRequestPayload)
	//if err != nil {
	//	http.Error(context.Writer, err.Error(), http.StatusBadRequest)
	//	return
	//}
	//
	//payload := url.Values{
	//	"login":         {authenticationRequestPayload.Login},
	//	"mot_de_passe":  {authenticationRequestPayload.Password},
	//	"r":             {authenticationRequestPayload.Redirect},
	//	"sub_connexion": {authenticationRequestPayload.SubConnexion},
	//}

	//requestBody := strings.NewReader(payload.Encode())
	// &{login=m-vanaud&mot_de_passe=Mv01Mars2002%24&r=index.php&sub_connexion=Se+connecter+%E2%86%92 0 -1}

	session, _ := server.Store.Get(context.Request, "je-name")

	fmt.Println(session.Values["PHPSESSID"])

	for _, cookie := range context.Request.Cookies() {
		fmt.Println("Found a cookie (request) named:", cookie.Name, cookie.Value)
	}
}
