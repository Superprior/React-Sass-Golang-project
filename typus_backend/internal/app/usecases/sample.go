package usecases

import (
	"backend/internal/app/models"
	"backend/internal/app/repos"
	"log"
)

type SampleUsecase struct {
	repo *repos.SampleRepo
}

func NewSampleUsecase() *SampleUsecase {
	r, err := repos.NewSampleRepo()
	if err != nil {
		log.Fatal("Could not create the SampleRepo")
	}

	return &SampleUsecase{
		repo: r,
	}
}

func (u *SampleUsecase) GetAllSamples() ([]*models.Sample, error) {
	samples, err := u.repo.GetList()

	if err != nil {
		return nil, err
	}
	return samples, nil
}

func (u *SampleUsecase) GetSampleById(id int) (*models.Sample, error) {

	sample, err := u.repo.GetInstanceById(id)

	if err != nil {
		return nil, err
	}

	return sample, nil
}

func (u *SampleUsecase) CreateSample(sample *models.Sample) (*models.Sample, error) {

	sample, err := u.repo.CreateInstance(sample)

	if err != nil {
		return nil, err
	}

	return sample, nil
}
