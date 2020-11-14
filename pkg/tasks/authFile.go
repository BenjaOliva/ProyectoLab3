package tasks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

var code string
var TokenR TokenResp

// TOKEN
type Token struct {
	Grant_type string 		`json:"grant_type"`
	Client_id int 			`json:"client_id"`
	Client_secret string 	`json:"client_secret"`
	Code string				`json:"code"`
	Redirect_uri string 	`json:"redirect_uri"`
}

type TokenResp struct {
	Access_token string 	`json:"access_token"`
	Token_type string 		`json:"token_type"`
	Expires_in int 			`json:"expires_in"`
	Scope string 			`json:"scope"`
	User_id int 			`json:"user_id"`
	Refresh_token string 	`json:"refresh_token"`
}

// FUNCIONES PARA INTERCAMBIAR EL CODE POR UN ACCESS TOKEN
func GetToken(c *gin.Context) {
	code = c.Query("code")
	fmt.Println("code: " + code)
	TokenRequest(code, c)
}

func TokenRequest(code string, c *gin.Context) {
	u := Token{
		Grant_type:    "authorization_code",
		Client_id:     4026193299288302,
		Client_secret: "XEiRyQ3sLtnVkoX33YmAA42uV022wcEw",
		Code:          code,
		Redirect_uri:  "http://localhost:8080/auth/code/",
	}

	b, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	// Intercambiamos code por token
	resp, err := http.Post("https://api.mercadolibre.com/oauth/token", "application/json; application/x-www-form-urlencoded", bytes.NewBuffer(b))

	if err != nil {
		fmt.Errorf("Error", err.Error())
		return
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	bodyString := string(data)
	fmt.Println(bodyString)

	json.Unmarshal(data, &TokenR)
	fmt.Printf("%+v\n", TokenR)

	c.JSON(200, TokenR)
}

