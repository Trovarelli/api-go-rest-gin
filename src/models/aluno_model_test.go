package models

import "testing"

func TestAlunoValidator_InvalidRGLength(t *testing.T) {
    a := &Aluno{Nome: "Ana", CPF: "123456789", RG: "123"}
    if err := AlunoValidator(a); err == nil {
        t.Fatalf("esperava erro de validação para RG inválido")
    }
}
