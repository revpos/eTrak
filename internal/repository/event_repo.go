package repository

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/revpos/eTrak/internal/models"
)

type EventRepository struct {
	DB *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{DB: db}
}

func (repo *EventRepository) InsertEvent(ctx context.Context, e *models.Event) error {
	props, err := json.Marshal(e.Properties)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO events (user_id, event_type, timestamp, properties)
		VALUES ($1, $2, $3, $4)
	`

	_, err = repo.DB.ExecContext(ctx, query, e.UserID, e.EventType, e.Timestamp, props)

	return err
}
