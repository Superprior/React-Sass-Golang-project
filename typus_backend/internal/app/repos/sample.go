package repos

import (
	"backend/internal/app/models"
	"backend/internal/app/store"
	"database/sql"
	"fmt"
)

type SampleRepo struct {
	store *store.Store
}

func NewSampleRepo() (*SampleRepo, error) {
	sConf := store.NewConfig()
	s := store.New(sConf)

	if err := s.Open(); err != nil {
		return nil, err
	}

	return &SampleRepo{
		store: s,
	}, nil
}

func (r *SampleRepo) GetList() ([]*models.Sample, error) {

	// Use GetInstanceById function for every ID to avoid nesting

	query := fmt.Sprintf("SELECT id FROM code_samples;")

	rows, err := r.store.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var samples []*models.Sample

	for rows.Next() {
		var id int

		if err := rows.Scan(&id); err != nil {
			return nil, err
		}

		sample, err := r.GetInstanceById(id)
		if err != nil {
			return nil, err
		}
		samples = append(samples, sample)
	}

	return samples, nil
}

func (r *SampleRepo) GetInstanceById(id int) (*models.Sample, error) {

	// Perform select query to check if the instance with that id exists

	var _id int
	query := fmt.Sprintf("SELECT id FROM code_samples WHERE id=%d;", id)

	err := r.store.DB.QueryRow(query).Scan(&_id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Sample with id %d does not exist", id)
		}
		return nil, err
	}

	// Use unnest function from postgres to separate the code lines
	query = fmt.Sprintf("SELECT id, title, unnest(content), language FROM code_samples WHERE id=%d;", id)

	// Perform the query
	rows, err := r.store.DB.Query(query)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Sample with id %d does not exist", id)
		}
		return nil, err
	}
	defer rows.Close()

	// Loop through the list of rows and create an array of lines
	var sample models.Sample
	var lines []string

	for rows.Next() {
		var line string

		if err := rows.Scan(&sample.ID, &sample.Title, &line, &sample.Language); err != nil {
			return nil, err
		}

		lines = append(lines, line)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	sample.Content = lines
	return &sample, nil
}

func (r *SampleRepo) CreateInstance(sample *models.Sample) (*models.Sample, error) {

	// Create postgres representation of the array to insert

	linesString := "["

	for i, line := range sample.Content {
		linesString += "'" + line + "'"
		if i != len(sample.Content)-1 {
			linesString += ", "
		}
	}
	linesString += "]"

	// Perform the insert

	query := fmt.Sprintf(
		"INSERT INTO code_samples (title, content, language) VALUES ('%s', ARRAY%s, '%s') RETURNING id;",
		sample.Title, linesString, sample.Language,
	)

	// Get the created sample's id and check for errors
	err := r.store.DB.QueryRow(query).Scan(&sample.ID)

	if err != nil {
		return nil, err
	}

	return sample, nil
}
