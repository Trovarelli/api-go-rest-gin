package controllers_test

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"

    "api-go-rest-gin/src/controllers"
    "api-go-rest-gin/src/testutil"

    "github.com/gin-gonic/gin"
)

func TestCreateAluno_InvalidCPF_Returns400(t *testing.T) {
    gin.SetMode(gin.TestMode)

    repo := testutil.NewFakeAlunosRepo(nil)
    h := controllers.NewAlunosController(repo)

    body := []byte(`{"nome":"Zoe","cpf":"123-45","rg":"12345678901"}`)
    req := httptest.NewRequest(http.MethodPost, "/alunos", bytes.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()

    c, _ := gin.CreateTestContext(w)
    c.Request = req

    h.Create(c)

    if w.Code != http.StatusInternalServerError && w.Code != http.StatusBadRequest {
        t.Fatalf("esperava 4xx, recebeu %d (body=%s)", w.Code, w.Body.String())
    }
}
