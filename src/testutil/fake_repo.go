package testutil

import (
    "context"

    "api-go-rest-gin/src/controllers"
    "api-go-rest-gin/src/models"
    "api-go-rest-gin/src/repository"
    "api-go-rest-gin/src/routes"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type FakeAlunosRepo struct {
    Data map[int64]models.Aluno
}

func NewFakeAlunosRepo(seed []models.Aluno) repository.AlunosRepository {
    m := make(map[int64]models.Aluno, len(seed))
    for _, a := range seed {
        m[a.Id] = a
    }
    return &FakeAlunosRepo{Data: m}
}

func (f *FakeAlunosRepo) GetAll(_ context.Context) ([]models.Aluno, error) {
    out := make([]models.Aluno, 0, len(f.Data))
    for _, a := range f.Data {
        out = append(out, a)
    }
    return out, nil
}

func (f *FakeAlunosRepo) GetById(_ context.Context, id int64) (models.Aluno, error) {
    if a, ok := f.Data[id]; ok {
        return a, nil
    }
    return models.Aluno{}, gorm.ErrRecordNotFound
}

func (f *FakeAlunosRepo) GetByCPF(_ context.Context, cpf string) (models.Aluno, error) {
    for _, a := range f.Data {
        if a.CPF == cpf {
            return a, nil
        }
    }
    return models.Aluno{}, gorm.ErrRecordNotFound
}

func (f *FakeAlunosRepo) Create(_ context.Context, aluno models.Aluno) (models.Aluno, error) {
    nextID := int64(len(f.Data) + 1)
    aluno.Id = nextID
    f.Data[nextID] = aluno
    return aluno, nil
}

func (f *FakeAlunosRepo) Update(_ context.Context, aluno *models.Aluno) error {
    if _, ok := f.Data[aluno.Id]; !ok {
        return gorm.ErrRecordNotFound
    }
    f.Data[aluno.Id] = *aluno
    return nil
}

func (f *FakeAlunosRepo) Delete(_ context.Context, id int64) error {
    if _, ok := f.Data[id]; !ok {
        return gorm.ErrRecordNotFound
    }
    delete(f.Data, id)
    return nil
}

func NewTestRouter(seed []models.Aluno) *gin.Engine {
    gin.SetMode(gin.TestMode)
    repo := NewFakeAlunosRepo(seed)
    handler := controllers.NewAlunosController(repo)
    return routes.SetupRouter(handler)
}

