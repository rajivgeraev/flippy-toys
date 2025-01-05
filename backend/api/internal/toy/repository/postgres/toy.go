// internal/toy/repository/postgres/toy.go
package postgres

import (
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/toy/model"
)

type ToyRepository struct {
	db *sql.DB
}

// AddPhoto implements repository.ToyRepository.
func (r *ToyRepository) AddPhoto(photo *model.Photo) error {
	panic("unimplemented")
}

// DeletePhoto implements repository.ToyRepository.
func (r *ToyRepository) DeletePhoto(photoID uuid.UUID) error {
	panic("unimplemented")
}

// SetMainPhoto implements repository.ToyRepository.
func (r *ToyRepository) SetMainPhoto(photoID uuid.UUID) error {
	panic("unimplemented")
}

// SoftDelete implements repository.ToyRepository.
func (r *ToyRepository) SoftDelete(toyID uuid.UUID) error {
	panic("unimplemented")
}

// Update implements repository.ToyRepository.
func (r *ToyRepository) Update(toy *model.Toy) error {
	panic("unimplemented")
}

func NewToyRepository(db *sql.DB) *ToyRepository {
	return &ToyRepository{db: db}
}

func (r *ToyRepository) Create(toy *model.Toy) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	ageRange, err := json.Marshal(toy.AgeRange)
	if err != nil {
		return err
	}

	query := `
        INSERT INTO toys (
            user_id, title, description, age_range,
            condition, category, status
        ) VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id, created_at, updated_at`

	err = tx.QueryRow(
		query,
		toy.UserID,
		toy.Title,
		toy.Description,
		ageRange,
		toy.Condition,
		toy.Category,
		toy.Status,
	).Scan(&toy.ID, &toy.CreatedAt, &toy.UpdatedAt)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *ToyRepository) GetByID(id uuid.UUID) (*model.Toy, error) {
	toy := &model.Toy{}

	query := `
        SELECT t.id, t.user_id, t.title, t.description, 
               t.age_range, t.condition, t.category, t.status,
               t.is_deleted, t.created_at, t.updated_at
        FROM toys t
        WHERE t.id = $1 AND t.is_deleted IS NULL`

	var ageRangeJSON []byte
	err := r.db.QueryRow(query, id).Scan(
		&toy.ID, &toy.UserID, &toy.Title, &toy.Description,
		&ageRangeJSON, &toy.Condition, &toy.Category, &toy.Status,
		&toy.IsDeleted, &toy.CreatedAt, &toy.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(ageRangeJSON, &toy.AgeRange)
	if err != nil {
		return nil, err
	}

	// Получаем фотографии
	photos, err := r.getPhotos(toy.ID)
	if err != nil {
		return nil, err
	}
	toy.Photos = photos

	return toy, nil
}

func (r *ToyRepository) getPhotos(toyID uuid.UUID) ([]model.Photo, error) {
	query := `
        SELECT id, toy_id, url, cloudinary_id, is_main, created_at
        FROM toy_photos
        WHERE toy_id = $1
        ORDER BY is_main DESC, created_at DESC`

	rows, err := r.db.Query(query, toyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var photos []model.Photo
	for rows.Next() {
		var photo model.Photo
		err := rows.Scan(
			&photo.ID, &photo.ToyID, &photo.URL,
			&photo.CloudinaryID, &photo.IsMain, &photo.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		photos = append(photos, photo)
	}

	return photos, rows.Err()
}

func (r *ToyRepository) ListActive(limit, offset int) ([]model.Toy, error) {
	query := `
        SELECT t.id, t.user_id, t.title, t.description, 
               t.age_range, t.condition, t.category, t.status,
               t.is_deleted, t.created_at, t.updated_at
        FROM toys t
        WHERE t.status = 'active' AND t.is_deleted IS NULL
        ORDER BY t.created_at DESC
        LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var toys []model.Toy
	for rows.Next() {
		var toy model.Toy
		var ageRangeJSON []byte

		err := rows.Scan(
			&toy.ID, &toy.UserID, &toy.Title, &toy.Description,
			&ageRangeJSON, &toy.Condition, &toy.Category, &toy.Status,
			&toy.IsDeleted, &toy.CreatedAt, &toy.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(ageRangeJSON, &toy.AgeRange)
		if err != nil {
			return nil, err
		}

		photos, err := r.getPhotos(toy.ID)
		if err != nil {
			return nil, err
		}
		toy.Photos = photos

		toys = append(toys, toy)
	}

	return toys, rows.Err()
}

func (r *ToyRepository) GetByUserID(userID uuid.UUID) ([]model.Toy, error) {
	query := `
        SELECT t.id, t.user_id, t.title, t.description, 
               t.age_range, t.condition, t.category, t.status,
               t.is_deleted, t.created_at, t.updated_at
        FROM toys t
        WHERE t.user_id = $1 AND t.is_deleted IS NULL
        ORDER BY t.created_at DESC`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var toys []model.Toy
	for rows.Next() {
		var toy model.Toy
		var ageRangeJSON []byte

		err := rows.Scan(
			&toy.ID, &toy.UserID, &toy.Title, &toy.Description,
			&ageRangeJSON, &toy.Condition, &toy.Category, &toy.Status,
			&toy.IsDeleted, &toy.CreatedAt, &toy.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(ageRangeJSON, &toy.AgeRange)
		if err != nil {
			return nil, err
		}

		photos, err := r.getPhotos(toy.ID)
		if err != nil {
			return nil, err
		}
		toy.Photos = photos

		toys = append(toys, toy)
	}

	return toys, rows.Err()
}
