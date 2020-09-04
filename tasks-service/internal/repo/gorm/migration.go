package gorm

import (
	"time"

	"github.com/pedroquerido/sword-challenge/tasks-service/pkg/task"
	"gorm.io/gorm"
)

// CreateTables ...
func CreateTables(db *gorm.DB) error {

	return db.AutoMigrate(&taskRow{})
}

// DropTables ...
func DropTables(db *gorm.DB) error {

	return db.Migrator().DropTable(&taskRow{})
}

// Populate ...
func Populate(db *gorm.DB) error {

	taskRows := []*taskRow{
		fromTask(&task.Task{
			ID:      "3b48d64d-f484-4fe9-a018-06cd3d7950a0",
			UserID:  "user-1",
			Summary: "summary-1",
			Date:    time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		}),
		fromTask(&task.Task{
			ID:      "3c5fd02a-ff20-4bfb-b943-e6540dbca4fa",
			UserID:  "user-2",
			Summary: "summary-2",
			Date:    time.Date(2020, time.September, 3, 21, 10, 0, 0, time.UTC),
		}),
	}

	for i := range taskRows {
		if err := db.Model(&taskRow{}).Create(taskRows[i]).Error; err != nil {
			return err
		}
	}

	return nil
}
