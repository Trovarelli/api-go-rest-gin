package controllers

import (
    "errors"
    "net/http"
    "strconv"

    "api-go-rest-gin/src/models"
    "api-go-rest-gin/src/repository"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type AlunosHandler struct {
    repo repository.AlunosRepository
}

func NewAlunosController(repo repository.AlunosRepository) *AlunosHandler {
    return &AlunosHandler{repo: repo}
}

func (h *AlunosHandler) GetAll(c *gin.Context) {
    alunos, err := h.repo.GetAll(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, alunos)
}

func (h *AlunosHandler) GetById(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.ParseInt(idStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
        return
    }
    aluno, err := h.repo.GetById(c.Request.Context(), id)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusNotFound, gin.H{"error": "aluno não encontrado"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }
    c.JSON(http.StatusOK, aluno)
}

func (h *AlunosHandler) Create(c *gin.Context) {
    var aluno models.Aluno
    if err := c.ShouldBindJSON(&aluno); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    created, err := h.repo.Create(c.Request.Context(), aluno)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, created)
}

func (h *AlunosHandler) Update(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.ParseInt(idStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
        return
    }
    var aluno models.Aluno
    if err := c.ShouldBindJSON(&aluno); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    aluno.Id = id
    if err := h.repo.Update(c.Request.Context(), &aluno); err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusNotFound, gin.H{"error": "aluno não encontrado"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }
    c.JSON(http.StatusOK, aluno)
}

func (h *AlunosHandler) Delete(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.ParseInt(idStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
        return
    }
    if err := h.repo.Delete(c.Request.Context(), id); err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusNotFound, gin.H{"error": "aluno não encontrado"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }
    c.Status(http.StatusNoContent)
}

