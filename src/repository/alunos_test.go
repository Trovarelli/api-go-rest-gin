package repository_test

import (
    "context"
    "testing"

    "api-go-rest-gin/src/models"
    "api-go-rest-gin/src/repository"

    "github.com/glebarez/sqlite"
    "gorm.io/gorm"
)

func newTestDB(t *testing.T) *gorm.DB {
    t.Helper()
    dsn := "file:" + t.Name() + "?mode=memory&cache=shared"
    db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
    if err != nil {
        t.Fatalf("falha ao abrir sqlite em memoria: %v", err)
    }
    if err := db.AutoMigrate(&models.Aluno{}); err != nil {
        t.Fatalf("falha ao migrar schema: %v", err)
    }
    return db
}

func newRepo(t *testing.T) repository.AlunosRepository {
    t.Helper()
    db := newTestDB(t)
    return repository.NewAlunosRepository(db)
}

func TestCreateAndGetById(t *testing.T) {
    repo := newRepo(t)
    ctx := context.TODO()

    in := models.Aluno{Nome: "Ana", CPF: "123456789", RG: "12345678901"}

    created, err := repo.Create(ctx, in)
    if err != nil {
        t.Fatalf("Create falhou: %v", err)
    }
    if created.Id == 0 {
        t.Fatalf("esperava ID gerado > 0")
    }

    got, err := repo.GetById(ctx, created.Id)
    if err != nil {
        t.Fatalf("GetById falhou: %v", err)
    }
    if got.Nome != in.Nome || got.CPF != in.CPF || got.RG != in.RG {
        t.Fatalf("registro divergente: got=%+v want=%+v", got, in)
    }
}

func TestGetAll_ReturnsAll(t *testing.T) {
    repo := newRepo(t)
    ctx := context.TODO()

    _, _ = repo.Create(ctx, models.Aluno{Nome: "Ana", CPF: "123456789", RG: "12345678901"})
    _, _ = repo.Create(ctx, models.Aluno{Nome: "Bia", CPF: "987654321", RG: "10987654321"})

    got, err := repo.GetAll(ctx)
    if err != nil {
        t.Fatalf("GetAll falhou: %v", err)
    }
    if len(got) != 2 {
        t.Fatalf("quantidade inesperada: got=%d want=%d", len(got), 2)
    }
}

func TestGetByCPF_FoundAndNotFound(t *testing.T) {
    repo := newRepo(t)
    ctx := context.TODO()

    _, _ = repo.Create(ctx, models.Aluno{Nome: "Ana", CPF: "111222333", RG: "12345678901"})

    a, err := repo.GetByCPF(ctx, "111222333")
    if err != nil {
        t.Fatalf("GetByCPF falhou (found): %v", err)
    }
    if a.Nome != "Ana" {
        t.Fatalf("esperava Ana, got=%s", a.Nome)
    }

    if _, err := repo.GetByCPF(ctx, "000000000"); err == nil {
        t.Fatalf("esperava erro para CPF inexistente")
    }
}

func TestUpdate_UpdatesFields_And_NotFound(t *testing.T) {
    repo := newRepo(t)
    ctx := context.TODO()

    created, _ := repo.Create(ctx, models.Aluno{Nome: "Ana", CPF: "123456789", RG: "12345678901"})

    created.Nome = "Ana Maria"
    created.CPF = "987654321"
    if err := repo.Update(ctx, &created); err != nil {
        t.Fatalf("Update falhou: %v", err)
    }

    got, _ := repo.GetById(ctx, created.Id)
    if got.Nome != "Ana Maria" || got.CPF != "987654321" {
        t.Fatalf("update nao persistiu: got=%+v", got)
    }

    missing := models.Aluno{Id: 999, Nome: "X", CPF: "111222333", RG: "12345678901"}
    if err := repo.Update(ctx, &missing); err == nil {
        t.Fatalf("esperava erro para Update de ID inexistente")
    }
}

func TestDelete_Deletes_And_NotFound(t *testing.T) {
    repo := newRepo(t)
    ctx := context.TODO()

    created, _ := repo.Create(ctx, models.Aluno{Nome: "Ana", CPF: "123456789", RG: "12345678901"})

    if err := repo.Delete(ctx, created.Id); err != nil {
        t.Fatalf("Delete falhou: %v", err)
    }
    if _, err := repo.GetById(ctx, created.Id); err == nil {
        t.Fatalf("esperava erro para registro deletado")
    }

    if err := repo.Delete(ctx, created.Id); err == nil {
        t.Fatalf("esperava erro ao deletar ID inexistente")
    }
}
