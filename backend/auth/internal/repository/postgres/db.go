package postgres

import (
    "database/sql"
    _ "github.com/lib/pq"
    "github.com/rajivgeraev/flippy-toys/backend/auth/internal/config"
)

func NewDB(cfg *config.PostgresConfig) (*sql.DB, error) {
    db, err := sql.Open("postgres", cfg.GetDSN())
    if err != nil {
        return nil, err
    }

    if err = db.Ping(); err != nil {
        return nil, err
    }

    return db, nil
}