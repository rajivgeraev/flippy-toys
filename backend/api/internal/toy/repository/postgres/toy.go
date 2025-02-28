// internal/toy/repository/postgres/toy.go
package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/rajivgeraev/flippy-toys/backend/api/internal/toy/model"
)

type ToyRepository struct {
	db *sql.DB
}

// AddPhoto implements repository.ToyRepository.
func (r *ToyRepository) AddPhoto(photo *model.Photo) error {
	query := `
        INSERT INTO toy_photos (
            toy_id, url, cloudinary_id, asset_id, is_main, created_at
        ) VALUES ($1, $2, $3, $4, $5, NOW())
        RETURNING id`

	err := r.db.QueryRow(
		query,
		photo.ToyID,
		photo.URL,
		photo.CloudinaryID,
		photo.AssetID,
		photo.IsMain,
	).Scan(&photo.ID)

	if err != nil {
		return fmt.Errorf("failed to add photo: %w", err)
	}

	return nil
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
	query := `
        UPDATE toys 
        SET title = $1, 
            description = $2, 
            condition = $3, 
            category = $4, 
            updated_at = NOW()
        WHERE id = $5 AND is_deleted IS NULL
        RETURNING updated_at`

	err := r.db.QueryRow(
		query,
		toy.Title,
		toy.Description,
		toy.Condition,
		toy.Category,
		toy.ID,
	).Scan(&toy.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to update toy: %w", err)
	}

	return nil
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

	fmt.Printf("Creating toy: %+v", toy)

	query := `
		INSERT INTO toys (
			user_id, title, description, condition, category, status, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
		RETURNING id, created_at, updated_at`

	err = tx.QueryRow(
		query,
		toy.UserID,
		toy.Title,
		toy.Description,
		toy.Condition, // Добавлено поле condition
		toy.Category,  // Добавлено поле category
		toy.Status,
	).Scan(&toy.ID, &toy.CreatedAt, &toy.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to insert toy: %w", err)
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

	var (
		ageRangeJSON []byte
		condition    *string // Используем указатели
		category     *string
	)

	err := r.db.QueryRow(query, id).Scan(
		&toy.ID, &toy.UserID, &toy.Title, &toy.Description,
		&ageRangeJSON, &condition, &category, &toy.Status,
		&toy.IsDeleted, &toy.CreatedAt, &toy.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if ageRangeJSON != nil {
		var ageRange model.AgeRange
		if err := json.Unmarshal(ageRangeJSON, &ageRange); err != nil {
			return nil, err
		}
		toy.AgeRange = &ageRange
	}

	if condition != nil {
		c := model.ToyCondition(*condition)
		toy.Condition = &c
	}

	if category != nil {
		c := model.ToyCategory(*category)
		toy.Category = &c
	}

	// Загружаем фотографии
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
		var (
			ageRangeJSON []byte
			condition    *string
			category     *string
		)

		err := rows.Scan(
			&toy.ID, &toy.UserID, &toy.Title, &toy.Description,
			&ageRangeJSON, &condition, &category, &toy.Status,
			&toy.IsDeleted, &toy.CreatedAt, &toy.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		// Декодируем age_range
		if ageRangeJSON != nil {
			var ageRange model.AgeRange
			if err := json.Unmarshal(ageRangeJSON, &ageRange); err != nil {
				return nil, err
			}
			toy.AgeRange = &ageRange
		}

		if condition != nil {
			c := model.ToyCondition(*condition)
			toy.Condition = &c
		}

		if category != nil {
			c := model.ToyCategory(*category)
			toy.Category = &c
		}

		// Загружаем фотографии
		photos, err := r.getPhotos(toy.ID)
		if err != nil {
			return nil, err
		}
		toy.Photos = photos

		toys = append(toys, toy)
	}

	return toys, rows.Err()
}

func (r *ToyRepository) ListWithFilters(filters *model.ToyFilters) ([]model.Toy, error) {
	query := `
        SELECT t.id, t.user_id, t.title, t.description, 
               t.condition, t.category, t.status,
               t.is_deleted, t.created_at, t.updated_at
        FROM toys t
        WHERE t.status = 'active' AND t.is_deleted IS NULL`

	args := make([]interface{}, 0)
	if len(filters.Categories) > 0 {
		query += ` AND t.category = ANY($1)`
		args = append(args, pq.Array(filters.Categories))
	}

	query += ` ORDER BY t.created_at DESC`

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to list toys: %w", err)
	}
	defer rows.Close()

	var toys []model.Toy
	for rows.Next() {
		var toy model.Toy
		err := rows.Scan(
			&toy.ID,
			&toy.UserID,
			&toy.Title,
			&toy.Description,
			&toy.Condition,
			&toy.Category,
			&toy.Status,
			&toy.IsDeleted,
			&toy.CreatedAt,
			&toy.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan toy: %w", err)
		}

		// Загружаем фотографии для каждой игрушки
		photos, err := r.getPhotos(toy.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get photos: %w", err)
		}
		toy.Photos = photos

		toys = append(toys, toy)
	}

	return toys, rows.Err()
}
