package etudiants

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

//type EtudiantPayload struct {
//	IdEtudiant         string `json:"id_etudiant"`
//	ReturnPage         string `json:"returnPage"`
//	NumeroEtudiant     string `json:"numero_etudiant"`
//	DateInscription    string `json:"date_inscription"`
//	Titre              string `json:"titre"`
//	Nom                string `json:"nom"`
//	Prenom             string `json:"prenom"`
//	Promo              string `json:"promo"`
//	Mail               string `json:"mail"`
//	Portable           string `json:"portable"`
//	FileAvatar         string `json:"file_avatar"`
//	Avatar             string `json:"avatar"`
//	Login              string `json:"login"`
//	Group              string `json:"groupe"`
//	CreateAccess       string `json:"create_access"`
//	Adresse            string `json:"adresse"`
//	CodePostal         string `json:"code_postal"`
//	Ville              string `json:"ville"`
//	Pays               string `json:"pays"`
//	CommuneNaissance   string `json:"commune_naissance"`
//	DptNaissance       string `json:"dpt_naissance"`
//	DateNaissanceJour  string `json:"date_naissance[jour]"`
//	DateNaissanceMois  string `json:"date_naissance[mois]"`
//	DateNaissanceAnnee string `json:"date_naissance[annee]"`
//	Commentaire        string `json:"commentaire"`
//	RibBanque          string `json:"rib_banque"`
//	RibIndicatif       string `json:"rib_indicatif"`
//	RibCompte          string `json:"rib_compte"`
//	RibClef            string `json:"rib_clef"`
//	RibDomiciliation   string `json:"rib_domiciliation"`
//	RibTitulaire       string `json:"rib_titulaire"`
//	RibIban            string `json:"rib_iban"`
//	RibBic             string `json:"rib_bic"`
//	Nationalite        string `json:"nationalite"`
//	NumeroSecu         string `json:"n_secu_sociale"`
//	CaisseSecu         string `json:"secu"`
//	SubAddEtudiant     string `json:"sub_add_etudiant"`
//}

func Get(context *gin.Context) {

	//fmt.Println(context.Request.URL.Query().Get("id"))
	session := sessions.Default(context)

	id := context.Request.URL.Query().Get("id")

	if _, err := strconv.Atoi(id); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid parameters",
		})
	}

	url := fmt.Sprintf("https://pro.siaje.com/jctaker/etudiants.php?etudiant=%s", id)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Cookie", fmt.Sprintf("PHPSESSID=%s", session.Get("PHPSESSID")))

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

	context.JSON(http.StatusBadRequest, gin.H{
		"code":    http.StatusBadRequest,
		"message": string(body),
	})
}

func Create(context *gin.Context) {

	session := sessions.Default(context)

	siajeEtudiantCreateEndpoint := "https://pro.siaje.com/jctaker/etudiants.php?mode=add"
	//authorized_url := "etudiants.php?etudiant=879"

	method := "POST"

	//var etudiantCreateRequestPayload EtudiantPayload
	//err := json.NewDecoder(context.Request.Body).Decode(&etudiantCreateRequestPayload)
	//if err != nil {
	//	http.Error(context.Writer, err.Error(), http.StatusBadRequest)
	//	return
	//}
	//
	//apiRequestPayload := url.Values{
	//	"id_etudiant":           {etudiantCreateRequestPayload.IdEtudiant},
	//	"returnPage":            {etudiantCreateRequestPayload.ReturnPage},
	//	"numero_etudiant":       {etudiantCreateRequestPayload.NumeroEtudiant},
	//	"date_inscription":      {etudiantCreateRequestPayload.DateInscription},
	//	"titre":                 {etudiantCreateRequestPayload.Titre},
	//	"nom":                   {etudiantCreateRequestPayload.Nom},
	//	"prenom":                {etudiantCreateRequestPayload.Prenom},
	//	"promo":                 {etudiantCreateRequestPayload.Promo},
	//	"mail":                  {etudiantCreateRequestPayload.Mail},
	//	"portable":              {etudiantCreateRequestPayload.Portable},
	//	"file_avatar":           {etudiantCreateRequestPayload.FileAvatar},
	//	"avatar":                {etudiantCreateRequestPayload.Avatar},
	//	"login":                 {etudiantCreateRequestPayload.Login},
	//	"groupe":                {etudiantCreateRequestPayload.Group},
	//	"create_access":         {etudiantCreateRequestPayload.CreateAccess},
	//	"adresse":               {etudiantCreateRequestPayload.Adresse},
	//	"code_postal":           {etudiantCreateRequestPayload.CodePostal},
	//	"ville":                 {etudiantCreateRequestPayload.Ville},
	//	"pays":                  {etudiantCreateRequestPayload.Pays},
	//	"commune_naissance":     {etudiantCreateRequestPayload.CommuneNaissance},
	//	"dpt_naissance":         {etudiantCreateRequestPayload.DptNaissance},
	//	"date_naissance[jour]":  {etudiantCreateRequestPayload.DateNaissanceJour},
	//	"date_naissance[mois]":  {etudiantCreateRequestPayload.DateNaissanceMois},
	//	"date_naissance[annee]": {etudiantCreateRequestPayload.DateNaissanceAnnee},
	//	"commentaire":           {etudiantCreateRequestPayload.Commentaire},
	//	"rib_banque":            {etudiantCreateRequestPayload.RibBanque},
	//	"rib_indicatif":         {etudiantCreateRequestPayload.RibIndicatif},
	//	"rib_compte":            {etudiantCreateRequestPayload.RibCompte},
	//	"rib_clef":              {etudiantCreateRequestPayload.RibClef},
	//	"rib_domiciliation":     {etudiantCreateRequestPayload.RibDomiciliation},
	//	"rib_titulaire":         {etudiantCreateRequestPayload.RibTitulaire},
	//	"rib_iban":              {etudiantCreateRequestPayload.RibIban},
	//	"rib_bic":               {etudiantCreateRequestPayload.RibBic},
	//	"nationalite":           {etudiantCreateRequestPayload.Nationalite},
	//	"n_secu_sociale":        {etudiantCreateRequestPayload.NumeroSecu},
	//	"secu":                  {etudiantCreateRequestPayload.CaisseSecu},
	//	"sub_add_etudiant":      {etudiantCreateRequestPayload.SubAddEtudiant},
	//}

	//payload := context.Request.Body

	body, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	requestBody := string(body)

	payload := strings.NewReader(requestBody)

	//fmt.Println("payload:", payload)

	student_creation_successful := false

	client := &http.Client{
		//	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		//		if len(via) > 1 {
		//			student_creation_successful = true
		//			return http.ErrUseLastResponse
		//		}
		//		return nil
		//	},
	}

	req, err := http.NewRequest(method, siajeEtudiantCreateEndpoint, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cookie", fmt.Sprintf("PHPSESSID=%s", session.Get("PHPSESSID")))

	fmt.Println("\nsiaje-token:", session.Get("PHPSESSID"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	if student_creation_successful {
		context.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "Student created",
		})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Not a valid student",
		})
	}
}

func Update(context *gin.Context) {

}
