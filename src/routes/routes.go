package routes

import (
    "api-go-rest-gin/src/controllers"

    "github.com/gin-gonic/gin"
)

func SetupRouter(alunos *controllers.AlunosHandler) *gin.Engine {
    r := gin.Default()
    r.GET("/alunos", alunos.GetAll)
    r.GET("/alunos/:id", alunos.GetById)
    r.GET("/alunos/cpf/:cpf", alunos.GetByCPF)
    r.POST("/alunos", alunos.Create)
    r.PUT("/alunos/:id", alunos.Update)
    r.DELETE("/alunos/:id", alunos.Delete)

    return r
}

func HandleRequest(alunos *controllers.AlunosHandler) {
    r := SetupRouter(alunos)
    r.Run()
}
