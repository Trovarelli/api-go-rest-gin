package main

import (
    "api-go-rest-gin/src/controllers"
    "api-go-rest-gin/src/database"
    "api-go-rest-gin/src/repository"
    "api-go-rest-gin/src/routes"
)

func main() {
    db, err := database.DatabaseConnect()
    if err != nil {
        panic(err)
    }

    repo := repository.NewAlunosRepository(db)
    alunosHandler := controllers.NewAlunosController(repo)

    routes.HandleRequest(alunosHandler)
}
 
