package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"richard_adekponya_fasttrack_quizapp.com/app/constants"
	"richard_adekponya_fasttrack_quizapp.com/app/models"
	"richard_adekponya_fasttrack_quizapp.com/app/services"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	// Public routes
	router.HandleFunc("/user/get-questions", getQuestions).Methods(http.MethodGet)
	router.HandleFunc("/user/submit-answers", submitAnswers).Methods(http.MethodPost)
	return router
}

func getQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	getAllQuestions(w)
}

func submitAnswers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	submitUserAnswersToQuizQuestions(w,r)
}

func getAllQuestions(w http.ResponseWriter) {
	quizQuestion, err := services.GetAllQuestions()
	if err != nil {
		log.Println("there is an error getting quiz questions ", err)
		errorResponse := models.ApiErrorResponse{ResponseCode: constants.FAILED_MESSAGE, ResponseMessage: constants.FAILED_MESSAGE, ErrorMessage: err.Error()}
		jsonResponse(w, errorResponse, http.StatusInternalServerError)
	} else {
		quizQuestionResponse := models.QuizQuestionResponse{ResponseCode: constants.SUCCESS_CODE, ResponseMessage: constants.SUCCESS_MESSAGE, QuizQuestion: quizQuestion}

		jsonResponse(w, quizQuestionResponse, http.StatusOK)
	}
}

func submitUserAnswersToQuizQuestions(w http.ResponseWriter, r*http.Request) {
	var userSubmissions models.UserSubmissions
	if err := json.NewDecoder(r.Body).Decode(&userSubmissions); err != nil {
        errorReponse:=models.ApiErrorResponse{ResponseCode: constants.FAILED_CODE,ResponseMessage: constants.FAILED_MESSAGE,ErrorMessage: err.Error()}
		jsonResponse(w,errorReponse,http.StatusBadRequest)
		return
    }
	res,err:= services.SubmitUserAnswers(userSubmissions)
	if(err!=nil){
		log.Println("there is an exception submiting user answers ",err)
		errorReponse:=models.ApiErrorResponse{ResponseCode: constants.FAILED_CODE,ResponseMessage: constants.FAILED_MESSAGE,ErrorMessage: err.Error()}
		jsonResponse(w,errorReponse,http.StatusInternalServerError)
	}else{
	jsonResponse(w,res,http.StatusAccepted)
	}

}

func jsonResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
