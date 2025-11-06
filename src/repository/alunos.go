package repository

import (
	"context"

	"api-go-rest-gin/src/models"

	"gorm.io/gorm"
)

type AlunosRepository interface {
	GetAll(ctx context.Context) ([]models.Aluno, error)
	GetById(ctx context.Context, id int64) (models.Aluno, error)
	GetByCPF(ctx context.Context, cpf string) (models.Aluno, error)
	Create(ctx context.Context, aluno models.Aluno) (models.Aluno, error)
	Update(ctx context.Context, aluno *models.Aluno) error
	Delete(ctx context.Context, id int64) error
}

type AlunosDB struct {
	db *gorm.DB
}

func NewAlunosRepository(db *gorm.DB) AlunosRepository {
	return &AlunosDB{db: db}
}

func (r *AlunosDB) GetAll(ctx context.Context) ([]models.Aluno, error) {
	var alunos []models.Aluno
	if err := r.db.WithContext(ctx).Find(&alunos).Error; err != nil {
		return nil, err
	}
	return alunos, nil
}

func (r *AlunosDB) GetById(ctx context.Context, id int64) (models.Aluno, error) {
	var aluno models.Aluno

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&aluno).Error; err != nil {
		return models.Aluno{}, err
	}

	return aluno, nil
}

func (r *AlunosDB) GetByCPF(ctx context.Context, cpf string) (models.Aluno, error) {
	var aluno models.Aluno

	if err := r.db.WithContext(ctx).Where("cpf = ?", cpf).First(&aluno).Error; err != nil {
		return models.Aluno{}, err
	}

	return aluno, nil
}

func (r *AlunosDB) Create(ctx context.Context, aluno models.Aluno) (models.Aluno, error) {
	if err := r.db.WithContext(ctx).Create(&aluno).Error; err != nil {
		return models.Aluno{}, err
	}
	return aluno, nil
}

func (r *AlunosDB) Update(ctx context.Context, aluno *models.Aluno) error {
	var id = aluno.Id

	tx := r.db.WithContext(ctx).
		Where("id = ?", id).
		Select("Nome", "CPF", "RG").
		Updates(aluno)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		var count int64
		if err := r.db.WithContext(ctx).Model(&models.Aluno{}).Where("id = ?", id).Count(&count).Error; err != nil {
			return err
		}
		if count == 0 {
			return gorm.ErrRecordNotFound
		}
	}
	return nil
}

func (r *AlunosDB) Delete(ctx context.Context, id int64) error {
	tx := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Aluno{})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
