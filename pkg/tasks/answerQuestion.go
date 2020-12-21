package tasks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

type AnswerFront struct {
	Question_id string `json:"questionId"`
	Text        string `json:"response"`
}

type AnswerMeli struct {
	Question_id int64  `json:"question_id"`
	Text        string `json:"text"`
}

var AnswerToPost AnswerFront

func AnswerQuestion(c *gin.Context) {

	bodyFront, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		fmt.Errorf("Error", err.Error())
		return
	}

	answerToPost := string(bodyFront)

	fmt.Println("AnswerFront: ", answerToPost)

	json.Unmarshal(bodyFront, &AnswerToPost)

	//AnswerToPost.id = answerToPost

	fmt.Printf("%+v\n", AnswerToPost)
	question_id, err := strconv.ParseInt(AnswerToPost.Question_id, 10, 64)
	answerQuestion := AnswerMeli{
		Question_id: question_id,
		Text:        AnswerToPost.Text,
	}

	jsonAnswerQuestion, _ := json.Marshal(answerQuestion)

	fmt.Println("JSON para MELI: ", string(jsonAnswerQuestion))

	responseAnswerQuestion, err := http.Post("https://api.mercadolibre.com/answers?access_token="+TokenR.Access_token, "application/json; application/x-www-form-urlencoded", bytes.NewBuffer(jsonAnswerQuestion))

	if err != nil {
		fmt.Errorf("Error", err.Error())
		return
	}

	defer responseAnswerQuestion.Body.Close()

	response, err := ioutil.ReadAll(responseAnswerQuestion.Body)

	if err != nil {
		fmt.Errorf("Error", err.Error())
		return
	}

	bodyString := string(response)

	fmt.Println(bodyString)

	json.Unmarshal(response, &AnswerToPost)

	//c.JSON(200, AnswerToPost)
}
