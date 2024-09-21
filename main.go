package main

import (
    "log"
    "net/http"
    "github.com/rs/cors"
	"richard_adekponya_fasttrack_quizapp.com/app/controllers"
     
)

func main() {
    corsOpts := cors.New(cors.Options{
        AllowedOrigins: []string{"*"}, //you service is available and allowed for this base url 
        AllowedMethods: []string{
            http.MethodGet,
            http.MethodPost,
            http.MethodPut,
            http.MethodPatch,
            http.MethodDelete,
            http.MethodOptions,
            http.MethodHead,
        },
    
        AllowedHeaders: []string{
            "*",//allow requests from cross functional platforms
    
        },
    })
    router := controllers.SetupRoutes()
    log.Println("Server running on port 9090 for quiz app backend engine")
    log.Fatal(http.ListenAndServe(":9090", corsOpts.Handler(router)))
}
