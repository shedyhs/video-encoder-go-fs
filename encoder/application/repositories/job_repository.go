package repositories

import (
	"encoder/domain"
	"fmt"

	"github.com/jinzhu/gorm"
)

type JobRepository interface {
	Insert(Job *domain.Job) (*domain.Job, error)
	Find(id string) (*domain.Job, error)
	Update(Job *domain.Job) (*domain.Job, error)
}

type JobRepositoryDb struct {
	Db *gorm.DB
}

func NewJobRepository(db *gorm.DB) *JobRepositoryDb {
	return &JobRepositoryDb{
		Db: db,
	}
}

func (repo JobRepositoryDb) Insert(Job *domain.Job) (*domain.Job, error) {
	err := repo.Db.Create(Job).Error
	if err != nil {
		return nil, err
	}
	return Job, nil
}

func (repo JobRepositoryDb) Find(id string) (*domain.Job, error) {
	var Job domain.Job
	repo.Db.Preload("Video").First(&Job, "id = ?", id)

	if Job.ID == "" {
		return nil, fmt.Errorf("job does not exists")
	}
	return &Job, nil
}

func (repo JobRepositoryDb) Update(job *domain.Job) (*domain.Job, error) {
	err := repo.Db.Save(&job).Error
	if err != nil {
		return nil, err
	}
	return job, nil
}
