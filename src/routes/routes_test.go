package routes_test

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "api-go-rest-gin/src/models"
    "api-go-rest-gin/src/testutil"
)

func TestGetAllAlunos_Returns200AndList(t *testing.T) {
    seed := []models.Aluno{
        {Id: 1, Nome: "Ana", CPF: "123456789", RG: "12345678901"},
        {Id: 2, Nome: "Bia", CPF: "987654321", RG: "10987654321"},
    }
    r := testutil.NewTestRouter(seed)

    req := httptest.NewRequest(http.MethodGet, "/alunos", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Fatalf("status inesperado: got=%d want=%d body=%s", w.Code, http.StatusOK, w.Body.String())
    }
    var got []models.Aluno
    if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
        t.Fatalf("falha ao decodificar resposta: %v", err)
    }
    if len(got) != len(seed) {
        t.Fatalf("quantidade inesperada: got=%d want=%d", len(got), len(seed))
    }
}

func TestGetAlunoByID_NotFound(t *testing.T) {
    r := testutil.NewTestRouter(nil)

    req := httptest.NewRequest(http.MethodGet, "/alunos/999", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    if w.Code != http.StatusNotFound {
        t.Fatalf("status inesperado: got=%d want=%d body=%s", w.Code, http.StatusNotFound, w.Body.String())
    }
}

func TestGetAlunoByCPF_Returns200(t *testing.T) {
    seed := []models.Aluno{
        {Id: 10, Nome: "Carlos", CPF: "111222333", RG: "12345678901"},
    }
    r := testutil.NewTestRouter(seed)

    req := httptest.NewRequest(http.MethodGet, "/alunos/cpf/111222333", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Fatalf("status inesperado: got=%d want=%d body=%s", w.Code, http.StatusOK, w.Body.String())
    }
}
