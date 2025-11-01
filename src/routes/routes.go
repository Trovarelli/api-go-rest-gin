package routes

import (
    "api-go-rest-gin/src/controllers"
    "github.com/gin-gonic/gin"
)

func HandleRequest(alunos *controllers.AlunosHandler) {
    r := gin.Default()
    r.GET("/alunos", alunos.GetAll)
    r.GET("/alunos/:id", alunos.GetById)
    r.POST("/alunos", alunos.Create)
    r.PUT("/alunos/:id", alunos.Update)
    r.DELETE("/alunos/:id", alunos.Delete)
    r.Run()
}
