package etudiants

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"juniorconseiltaker-siaje-api/server"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

type EtudiantPayload struct {
	IdEtudiant         string `json:"id_etudiant"`
	ReturnPage         string `json:"returnPage"`
	NumeroEtudiant     string `json:"numero_etudiant"`
	DateInscription    string `json:"date_inscription"`
	Titre              string `json:"titre"`
	Nom                string `json:"nom"`
	Prenom             string `json:"prenom"`
	Promo              string `json:"promo"`
	Mail               string `json:"mail"`
	Portable           string `json:"portable"`
	FileAvatar         string `json:"file_avatar"`
	Avatar             string `json:"avatar"`
	Login              string `json:"login"`
	Group              string `json:"groupe"`
	CreateAccess       string `json:"create_access"`
	Adresse            string `json:"adresse"`
	CodePostal         string `json:"code_postal"`
	Ville              string `json:"ville"`
	Pays               string `json:"pays"`
	CommuneNaissance   string `json:"commune_naissance"`
	DptNaissance       string `json:"dpt_naissance"`
	DateNaissanceJour  string `json:"date_naissance[jour]"`
	DateNaissanceMois  string `json:"date_naissance[mois]"`
	DateNaissanceAnnee string `json:"date_naissance[annee]"`
	Commentaire        string `json:"commentaire"`
	RibBanque          string `json:"rib_banque"`
	RibIndicatif       string `json:"rib_indicatif"`
	RibCompte          string `json:"rib_compte"`
	RibClef            string `json:"rib_clef"`
	RibDomiciliation   string `json:"rib_domiciliation"`
	RibTitulaire       string `json:"rib_titulaire"`
	RibIban            string `json:"rib_iban"`
	RibBic             string `json:"rib_bic"`
	Nationalite        string `json:"nationalite"`
	NumeroSecu         string `json:"n_secu_sociale"`
	CaisseSecu         string `json:"secu"`
	SubAddEtudiant     string `json:"sub_add_etudiant"`
}

func Get(context *gin.Context) {

	requestId := context.Request.URL.Query().Get("id")

	if requestId == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Missing parameter id",
		})
	} else {
		if _, err := strconv.Atoi(requestId); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "Invalid parameter id",
			})
		}
	}

	siajeStudentGetEndpoint := fmt.Sprintf("https://pro.siaje.com/jctaker/etudiants.php?mode=edit&etudiant=%s", requestId)
	method := "GET"

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Jar: server.CookieJar,
	}
	req, err := http.NewRequest(method, siajeStudentGetEndpoint, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
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

	etudiant := make(map[string]string)

	const void string = ""

	htmlParserRegex := regexp.MustCompile("<head>(?:.|\\n|\\r)+?<\\/head>")

	firstBatch := htmlParserRegex.ReplaceAllString(string(body), void)

	htmlParserRegex = regexp.MustCompile("<div(?:.|\\n|\\r)+?<div id=\"content\" mode=\"\">")

	secondBatch := htmlParserRegex.ReplaceAllString(firstBatch, void)

	htmlParserRegex = regexp.MustCompile("<script(?:.|\\n|\\r)+?<\\/script>")

	thirdBatch := htmlParserRegex.ReplaceAllString(secondBatch, void)

	htmlParserRegex = regexp.MustCompile("<!-- #suggestions -->(?:.|\\n|\\r)+?<\\/div><!-- fin new_top -->")

	studentSectionContent := htmlParserRegex.ReplaceAllString(thirdBatch, void)

	htmlTagRegex := regexp.MustCompile(`(?P<tag>id|name|value)="(?P<item>.*?)"`)

	htmlTagMatch := htmlTagRegex.FindAllStringSubmatch(string(studentSectionContent), -1)

	for index := 0; index < len(htmlTagMatch); index += 1 {

		if index+2 >= len(htmlTagMatch) {
			break
		}

		if htmlTagMatch[index][1] == "id" && htmlTagMatch[index+1][1] == "name" && htmlTagMatch[index+2][1] == "value" {
			etudiant[htmlTagMatch[index+1][2]] = htmlTagMatch[index+2][2]
		}
	}

	textareaTagRegex := regexp.MustCompile(`(?:<textarea)(.|\n|\r)+?(<\/textarea>)`)
	textTagRegex := regexp.MustCompile(`id="([^"]+)"`)
	textContentRegex := regexp.MustCompile(`(?:class="">)((?:.|\n|\r)+?)(?:<\/textarea>)`)

	textarea := textareaTagRegex.FindAllStringSubmatch(string(studentSectionContent), -1)

	for item := 0; item < len(textarea); item += 1 {

		text := strings.Join(textarea[item], " ")
		match_id := textTagRegex.FindAllStringSubmatch(text, 1)
		match_text := textContentRegex.FindAllStringSubmatch(text, -1)

		if match_id[0][1] == "adresse" {

			endOfLine := match_text[0][1][len(match_text[0][1])-1:]
			if endOfLine == "\n" {
				match_text[0][1] = match_text[0][1][:len(match_text[0][1])-1]
			}

			etudiant[match_id[0][1]] = match_text[0][1]
		}

	}

	birthDateRegex := regexp.MustCompile(`(?:selected="selected">)((.|\n|\r)+?)(?:<\/option>)`)

	birthDateRegexMatch := birthDateRegex.FindAllStringSubmatch(string(studentSectionContent), -1)

	const day int = 1
	const month int = 2
	const year int = 3

	etudiant["date_naissance[jour]"] = birthDateRegexMatch[day][1]
	etudiant["date_naissance[mois]"] = birthDateRegexMatch[month][1]
	etudiant["date_naissance[annee]"] = birthDateRegexMatch[year][1]

	context.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"message":  "Found student",
		"etudiant": etudiant,
	})
}

func Create(context *gin.Context) {

	siajeStudentCreateEndpoint := "https://pro.siaje.com/jctaker/etudiants.php?mode=add"

	method := "POST"

	var etudiantCreateRequestPayload EtudiantPayload

	err := json.NewDecoder(context.Request.Body).Decode(&etudiantCreateRequestPayload)
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	apiRequestPayload := url.Values{
		"id_etudiant":           {etudiantCreateRequestPayload.IdEtudiant},
		"returnPage":            {etudiantCreateRequestPayload.ReturnPage},
		"numero_etudiant":       {etudiantCreateRequestPayload.NumeroEtudiant},
		"date_inscription":      {etudiantCreateRequestPayload.DateInscription},
		"titre":                 {etudiantCreateRequestPayload.Titre},
		"nom":                   {etudiantCreateRequestPayload.Nom},
		"prenom":                {etudiantCreateRequestPayload.Prenom},
		"promo":                 {etudiantCreateRequestPayload.Promo},
		"mail":                  {etudiantCreateRequestPayload.Mail},
		"portable":              {etudiantCreateRequestPayload.Portable},
		"file_avatar":           {etudiantCreateRequestPayload.FileAvatar},
		"avatar":                {etudiantCreateRequestPayload.Avatar},
		"login":                 {etudiantCreateRequestPayload.Login},
		"groupe":                {etudiantCreateRequestPayload.Group},
		"create_access":         {etudiantCreateRequestPayload.CreateAccess},
		"adresse":               {etudiantCreateRequestPayload.Adresse},
		"code_postal":           {etudiantCreateRequestPayload.CodePostal},
		"ville":                 {etudiantCreateRequestPayload.Ville},
		"pays":                  {etudiantCreateRequestPayload.Pays},
		"commune_naissance":     {etudiantCreateRequestPayload.CommuneNaissance},
		"dpt_naissance":         {etudiantCreateRequestPayload.DptNaissance},
		"date_naissance[jour]":  {etudiantCreateRequestPayload.DateNaissanceJour},
		"date_naissance[mois]":  {etudiantCreateRequestPayload.DateNaissanceMois},
		"date_naissance[annee]": {etudiantCreateRequestPayload.DateNaissanceAnnee},
		"commentaire":           {etudiantCreateRequestPayload.Commentaire},
		"rib_banque":            {etudiantCreateRequestPayload.RibBanque},
		"rib_indicatif":         {etudiantCreateRequestPayload.RibIndicatif},
		"rib_compte":            {etudiantCreateRequestPayload.RibCompte},
		"rib_clef":              {etudiantCreateRequestPayload.RibClef},
		"rib_domiciliation":     {etudiantCreateRequestPayload.RibDomiciliation},
		"rib_titulaire":         {etudiantCreateRequestPayload.RibTitulaire},
		"rib_iban":              {etudiantCreateRequestPayload.RibIban},
		"rib_bic":               {etudiantCreateRequestPayload.RibBic},
		"nationalite":           {etudiantCreateRequestPayload.Nationalite},
		"n_secu_sociale":        {etudiantCreateRequestPayload.NumeroSecu},
		"secu":                  {etudiantCreateRequestPayload.CaisseSecu},
		"sub_add_etudiant":      {etudiantCreateRequestPayload.SubAddEtudiant},
	}

	payload := strings.NewReader(apiRequestPayload.Encode())

	requestSuccessful := false

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			forwardUrl := req.URL.String()

			fmt.Println("forwarding to: ", forwardUrl)

			match, _ := regexp.MatchString("(https:\\/\\/pro.siaje.com\\/jctaker\\/etudiants.php\\?etudiant=)([0-9]+)", forwardUrl)

			if match {

				split := strings.SplitAfter(forwardUrl, "=")
				fmt.Println("id:", split[1])

				requestSuccessful = true
				return http.ErrUseLastResponse
			}

			return nil
		},
		Jar: server.CookieJar,
	}

	req, err := http.NewRequest(method, siajeStudentCreateEndpoint, payload)

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
		fmt.Println("cookie:", cookie)
	}

	if requestSuccessful {
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

func Delete(context *gin.Context) {

}
