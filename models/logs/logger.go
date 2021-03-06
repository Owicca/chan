package logs

import (
	"gorm.io/gorm"
	"errors"
)


type Entry struct {
	ID int64 `gorm:"primaryKey,column:id"`
	Action string
	Subject int64
	Object int64
	ObjectType string
	Data string
}

func (e Entry) TableName() string {
	return "log_actions"
}

// A generic interface that writes an "Entry" to a "Store".
// Similar to io.Writer.
type Logger interface {
	Write(data Entry) error
}

type ActionLog struct {
	Store *gorm.DB
}

// Create an ActionLog.
func NewActionLog(store *gorm.DB) *ActionLog {
	if store == nil {
		return nil
	}

	lg := &ActionLog{
		Store: store,
	}

	return lg
}

// Write "Entry" to "Store",
// return error if fields are invalid or could not write to "Store".
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

	if entry.ObjectType == "" {
		return errors.New("Invalid action object_type")
	}

	if RequiresData(entry.Action) && entry.Data == "" {
		return errors.New("No data provided on insert/update")
	}

	l.Store.Create(&entry)

	return nil
}