// internal/user/repository/postgres/user.go
package postgres

import (
	"database/sql"

	"github.com/google/uuid"

	"github.com/rajivgeraev/flippy-toys/backend/api/internal/user/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUserWithTelegram создает пользователя и его телеграм профиль в транзакции
func (r *UserRepository) CreateUserWithTelegram(user *model.User, telegramProfile *model.TelegramProfile) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Создаем основного пользователя
	query := `
       INSERT INTO users (
           display_name, access_level
       ) VALUES ($1, $2)
       RETURNING id, created_at, updated_at`

	err = tx.QueryRow(
		query,
		telegramProfile.FirstName, // Используем имя из телеграма как display_name
		"basic",
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return err
	}

	// Создаем телеграм профиль
	query = `
       INSERT INTO user_telegram (
           user_id, telegram_id, username, first_name, last_name,
           photo_url, language_code, is_premium, phone
       ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
       RETURNING id, created_at, updated_at`

	err = tx.QueryRow(
		query,
		user.ID,
		telegramProfile.TelegramID,
		telegramProfile.Username,
		telegramProfile.FirstName,
		telegramProfile.LastName,
		telegramProfile.PhotoURL,
		telegramProfile.LanguageCode,
		telegramProfile.IsPremium,
		telegramProfile.Phone,
	).Scan(&telegramProfile.ID, &telegramProfile.CreatedAt, &telegramProfile.UpdatedAt)
	if err != nil {
		return err
	}

	user.TelegramProfile = telegramProfile

	return tx.Commit()
}

// GetByTelegramID получает пользователя по telegram_id
func (r *UserRepository) GetByTelegramID(telegramID int64) (*model.User, error) {
	query := `
       SELECT u.id, u.display_name, u.email, u.phone, 
              u.real_first_name, u.real_last_name, u.access_level,
              u.is_deleted, u.created_at, u.updated_at,
              t.id, t.telegram_id, t.username, t.first_name, t.last_name,
              t.photo_url, t.language_code, t.is_premium, t.phone,
              t.created_at, t.updated_at
       FROM users u
       JOIN user_telegram t ON t.user_id = u.id
       WHERE t.telegram_id = $1 AND u.is_deleted IS NULL`

	var user model.User
	var telegram model.TelegramProfile

	err := r.db.QueryRow(query, telegramID).Scan(
		&user.ID, &user.DisplayName, &user.Email, &user.Phone,
		&user.RealFirstName, &user.RealLastName, &user.AccessLevel,
		&user.IsDeleted, &user.CreatedAt, &user.UpdatedAt,
		&telegram.ID, &telegram.TelegramID, &telegram.Username,
		&telegram.FirstName, &telegram.LastName, &telegram.PhotoURL,
		&telegram.LanguageCode, &telegram.IsPremium, &telegram.Phone,
		&telegram.CreatedAt, &telegram.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	telegram.UserID = user.ID
	user.TelegramProfile = &telegram

	return &user, nil
}

// UpdateTelegramProfile обновляет данные телеграм профиля
func (r *UserRepository) UpdateTelegramProfile(profile *model.TelegramProfile) error {
	query := `
       UPDATE user_telegram
       SET username = $1,
           first_name = $2,
           last_name = $3,
           photo_url = $4,
           language_code = $5,
           is_premium = $6,
           phone = $7,
           updated_at = CURRENT_TIMESTAMP
       WHERE telegram_id = $8
       RETURNING updated_at`

	return r.db.QueryRow(
		query,
		profile.Username,
		profile.FirstName,
		profile.LastName,
		profile.PhotoURL,
		profile.LanguageCode,
		profile.IsPremium,
		profile.Phone,
		profile.TelegramID,
	).Scan(&profile.UpdatedAt)
}

// UpdatePhone обновляет номер телефона в обеих таблицах
func (r *UserRepository) UpdatePhone(telegramID int64, phone string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Обновляем телефон в telegram профиле
	query := `
       UPDATE user_telegram 
       SET phone = $1, 
           updated_at = CURRENT_TIMESTAMP
       WHERE telegram_id = $2
       RETURNING user_id`

	var userID uuid.UUID
	err = tx.QueryRow(query, phone, telegramID).Scan(&userID)
	if err != nil {
		return err
	}

	// Обновляем телефон в основной таблице
	query = `
       UPDATE users 
       SET phone = $1,
           access_level = 'advanced',
           updated_at = CURRENT_TIMESTAMP
       WHERE id = $2`

	_, err = tx.Exec(query, phone, userID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *UserRepository) GetByID(id uuid.UUID) (*model.User, error) {
	user := &model.User{}

	query := `
        SELECT u.id, u.display_name, u.phone, u.access_level, u.created_at, u.updated_at,
               t.telegram_id, t.username, t.first_name, t.last_name, 
               t.photo_url, t.language_code, t.is_premium
        FROM users u
        LEFT JOIN user_telegram t ON t.user_id = u.id
        WHERE u.id = $1 AND u.is_deleted IS NULL`

	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.DisplayName,
		&user.Phone,
		&user.AccessLevel,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.TelegramProfile.TelegramID,
		&user.TelegramProfile.Username,
		&user.TelegramProfile.FirstName,
		&user.TelegramProfile.LastName,
		&user.TelegramProfile.PhotoURL,
		&user.TelegramProfile.LanguageCode,
		&user.TelegramProfile.IsPremium,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}
