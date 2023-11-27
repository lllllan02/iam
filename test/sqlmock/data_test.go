package sqlmock

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-redis/redismock/v9"
	"github.com/lllllan02/iam/internal/data"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupData(t *testing.T) (*data.Data, sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}

	// mock db
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm connection: %v", err)
	}

	// mock rdb
	rdb, _ := redismock.NewClientMock()

	// mock data
	d := data.NewData(db, rdb, nil)

	return d, mock
}
