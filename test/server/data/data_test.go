package data

import (
	"os"
	"testing"

	"github.com/lllllan02/iam/internal/data"
	"github.com/lllllan02/iam/pkg/config"
	"github.com/lllllan02/iam/pkg/log"
)

var (
	d        *data.Data
	userData data.UserData
)

func setup() {
	conf := config.NewConfig("../../../config/test.yml")

	logger := log.NewLog(conf)
	db := data.NewDB(conf, logger)
	rdb := data.NewRedis(conf)

	d = data.NewData(db, rdb, logger)
	userData = data.NewUserData(d)

	// TODO: Insert data for the test
}

func TestMain(m *testing.M) {
	setup()

	os.Exit(m.Run())
}
