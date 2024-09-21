package services

import (
	"errors"
	"log"
	"math"
	"richard_adekponya_fasttrack_quizapp.com/app/constants"
	"richard_adekponya_fasttrack_quizapp.com/app/models"
	"math/rand"
)

func generateQuizQuestions() models.QuizQuestion {
	return models.QuizQuestion{
		UserId:generateRandomUserIdFromList(getRecentActiveQuizzers()),
		Questions: []models.Question{
			{Id: 1, Question: "What is the capital of France?", PossibleAnswers: map[string]string{"A": "Berlin", "B": "Madrid", "C": "Paris", "D": "Rome"}},
			{Id: 2, Question: "What is 2 + 2?", PossibleAnswers: map[string]string{"A": "3", "B": "4", "C": "5", "D": "6"}},
			{Id: 3, Question: "What is the capital Of Ghana?", PossibleAnswers: map[string]string{"A": "Accra", "B": "Ahafo", "C": "Sunyani", "D": "Moli"}},
			{Id: 4, Question: "What is the continent of Ghana?", PossibleAnswers: map[string]string{"A": "Africa", "B": "Asia", "C": "Europe", "D": "America"}},
			{Id: 5, Question: "Which of these languages is not spoken in Malta?", PossibleAnswers: map[string]string{"A": "Maltese", "B": "English", "C": "Italia", "D": "Akan"}},
		},
	}
}

func generateQuizAnswers() models.QuizAnswer {
	return models.QuizAnswer{
		CorrectAnswer: []models.CorrectAnswer{
			{QuestionId: 1, CorrectAnswer: "C", AnswerDetails: "Paris"},
			{QuestionId: 2, CorrectAnswer: "B", AnswerDetails: "4"},
			{QuestionId: 3, CorrectAnswer: "A", AnswerDetails: "Accra"},
			{QuestionId: 4, CorrectAnswer: "A", AnswerDetails: "Africa"},
			{QuestionId: 5, CorrectAnswer: "D", AnswerDetails: "Akan"},
		},
	}
}

func generateRandomUserIdFromList(recentUsers [] models.RecentUserScoreBoard)int{
	sizeOfQuizzers:=len(recentUsers)
	return rand.Intn(sizeOfQuizzers- 1 - 0) + 0 + 1

}

func getRecentActiveQuizzers() ( []models.RecentUserScoreBoard) {
	var rBoards=[]models.RecentUserScoreBoard{
		{UserId: 1, UserName: "kfolly@c2", PercentageScore: 32.65},
		{UserId: 2, UserName: "fgdgsy@a2", PercentageScore: 12.76},
		{UserId: 3, UserName: "turysa@3h", PercentageScore: 86.54},
		{UserId: 4, UserName: "tyrh4y@gg", PercentageScore: 62.31},
		{UserId: 5, UserName: "kassls@wd", PercentageScore: 70.25},
		{UserId: 6, UserName: "kdgsf3@5h", PercentageScore: 18.25},
		{UserId: 7, UserName: "cash56@sd", PercentageScore: 58.25},
		{UserId: 8, UserName: "ashra5@qs", PercentageScore: 48.25},
	}
	return rBoards
}

func generateRecentUsersScoreBoard(userId int,userRecentPercentageScore float64) ( []models.RecentUserScoreBoard) {
	var rBoards=[]models.RecentUserScoreBoard{
		{UserId: 1, UserName: "kfolly@c2", PercentageScore: 32.65},
		{UserId: 2, UserName: "fgdgsy@a2", PercentageScore: 12.76},
		{UserId: 3, UserName: "turysa@3h", PercentageScore: 86.54},
		{UserId: 4, UserName: "tyrh4y@gg", PercentageScore: 62.31},
		{UserId: 5, UserName: "kassls@wd", PercentageScore: 70.25},
		{UserId: 6, UserName: "kdgsf3@5h", PercentageScore: 18.25},
		{UserId: 7, UserName: "cash56@sd", PercentageScore: 58.25},
		{UserId: 8, UserName: "ashra5@qs", PercentageScore: 48.25},
	}
	for i,recentUserBoard:=range rBoards{
		if userId==recentUserBoard.UserId {
			log.Println("found recent Id",recentUserBoard.UserId)
			rBoards[i].PercentageScore=userRecentPercentageScore
			log.Println("found recent ",recentUserBoard)
		}
	}
	return rBoards
}



func GetAllQuestions() (models.QuizQuestion, error) {
	return generateQuizQuestions(), nil
}

func SubmitUserAnswers(userSubmissions models.UserSubmissions) (models.FinalQuizResponse, error) {
	userSubmissions,err:=validateUserAnswersAndQuizAnswerLength(userSubmissions)
	log.Println("user selections ",userSubmissions)
	if(err!=nil){
		return models.FinalQuizResponse{},err
	}
	// var userId = userSubmissions.UserId
	var userAnswers = userSubmissions.UserAnswers
	var quizAnswers = generateQuizAnswers()
	var totalQuestions=len(generateQuizQuestions().Questions)
	
	totalCorrectAnswer := 0

	for i:=0;i<len(userAnswers);i++{
		if userAnswers[i].QuestionId==quizAnswers.CorrectAnswer[i].QuestionId && userAnswers[i].UserSelection==quizAnswers.CorrectAnswer[i].CorrectAnswer{
			totalCorrectAnswer+=1;
		}
	}
	var totalPercentageScore=calculatePercentageScoreInQuiz(totalCorrectAnswer,totalQuestions)
	var recentUserScoreBoard=generateRecentUsersScoreBoard(userSubmissions.UserId,totalPercentageScore);
	var overAllPercentageUserScore=calculateComparison(float64(totalPercentageScore),len(recentUserScoreBoard),recentUserScoreBoard)
	
	return models.FinalQuizResponse{ResponseCode: constants.SUCCESS_CODE,ResponseMessage: constants.SUCCESS_MESSAGE,TotalQuestions:totalQuestions,TotalCorrectAnswers: totalCorrectAnswer,PercentageScoreInQuiz: totalPercentageScore,RecentUsersScoreBoard: recentUserScoreBoard,PercentageOfUsersPerformedBetterThan: overAllPercentageUserScore },nil

}

func calculatePercentageScoreInQuiz(totalCorrectScore int,totalQuestions int)(float64){
	return math.Ceil((float64(totalCorrectScore)/float64(totalQuestions))*100)

}

func validateUserAnswersAndQuizAnswerLength(userSubmissions models.UserSubmissions)(models.UserSubmissions,error){
	if(len(generateQuizAnswers().CorrectAnswer)!=len(userSubmissions.UserAnswers)){
		return models.UserSubmissions{}, errors.New("user answers not matching quiz answers in length. possible non answered selection to questions")
	}
	return userSubmissions,nil
}

func calculateComparison(currentScore float64,totalUsers int,recentUsersScoreBoard []models.RecentUserScoreBoard) float64 {
	totalPlusActiveUsers := totalUsers
	performedBetterThanUsers := 0

	for _, result := range recentUsersScoreBoard {
		if currentScore > result.PercentageScore {
			performedBetterThanUsers++
		}
	}

	return math.Ceil(float64(performedBetterThanUsers) / float64(totalPlusActiveUsers-1) * 100)
}
