package models

// import (
//     "time"

//     _ "github.com/go-sql-driver/mysql"
// )

type User struct{
	UserId int `json:"userId"`
	FullName string `json:"fullName"`
	DateOfBirth string `json:"dateOfBirth"`
	Address string `json:"address"`
	MobileNumber string `json:"mobileNumber"`
	Email string `json:"email"`
	Password string `json:"password"`
	DateCreated string `json:"dateCreated"`
	LatestQuizPercentage float64 `json:"latestQuizScore"`
}