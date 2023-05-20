package usecases

import (
	"backend/internal/models"
	"backend/internal/repos"
	"log"
)

// Sample usecase definition.
// Contains its repo to perform db queries.
type SampleUsecase struct {
	repo *repos.SampleRepo
}

func NewSampleUsecase() (uc *SampleUsecase) {
	r, err := repos.NewSampleRepo()
	if err != nil {
		log.Fatal("Could not create the SampleRepo")
	}

	uc = &SampleUsecase{repo: r}
	return
}

// Usecase (inner logic) for getting the samples list.
func (u *SampleUsecase) GetAllSamples() (samples []*models.Sample, err error) {

	samples, err = u.repo.GetList()
	return
}

// Usecase (inner logic) for getting a sample by id.
func (u *SampleUsecase) GetSampleById(id int) (sample *models.Sample, err error) {

	sample, err = u.repo.GetInstanceById(id)
	return
}

// Usecase (inner logic) for creating a new sample.
func (u *SampleUsecase) CreateSample(sampleReceived *models.Sample) (sampleReturned *models.Sample, err error) {

	sampleReturned, err = u.repo.CreateInstance(sampleReceived)
	return
}
