package logs

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"context"
	"errors"
)


type Entry struct {
	Action string
	Subject int64
	Object int64
	Object_type string
	Data string
}

// A generic interface that writes an "Entry" to a "Store".
// Similar to io.Writer
type Logger interface {
	Write(data Entry) error
}

type ActionLog struct {
	Store *pgxpool.Pool
}

// Create an ActionLog
func NewActionLog(store *pgxpool.Pool) *ActionLog {
	if store == nil {
		return nil
	}

	lg := &ActionLog{
		Store: store,
	}

	return lg
}

// Write "Entry" to "Store",
// return error if fields are invalid or could not write to "Store"
func (l *ActionLog) Write(entry Entry) error {
	if !AllowedAction(entry.Action).IsValid() {
		return errors.New("Invalid action provided")
	}

	if entry.Subject <= 0 {
		return errors.New("Invalid action subject")
	}

	if entry.Object <= 0 {
		return errors.New("Invalid action object")
	}

	if entry.Object_type == "" {
		return errors.New("Invalid action object_type")
	}

	if RequiresData(entry.Action) && entry.Data == "" {
		return errors.New("No data provided on insert/update")
	}

	sql := `
	INSERT INTO
	log_actions(action, subject, object, object_type, data)
	VALUES($1, $2, $3, $4, $5)
	`
	_, err := l.Store.Exec(context.Background(), sql, entry.Action, entry.Subject, entry.Object, entry.Object_type, entry.Data)
	if err != nil {
		return err
	}

	return nil
}