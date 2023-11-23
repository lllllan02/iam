package data

import (
	context "context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lllllan02/iam/internal/data"
	model "github.com/lllllan02/iam/internal/model"
	"github.com/stretchr/testify/assert"
)

func setupUserData(t *testing.T) (data.UserData, sqlmock.Sqlmock) {
	d, mock := setupData(t)

	// mock userData
	userData := data.NewUserData(d)

	return userData, mock
}

func TestUserData_Create(t *testing.T) {
	userData, mock := setupUserData(t)

	ctx := context.Background()
	user := &model.User{
		Model: model.Model{
			InstanceID: model.NewInstanceID("uid-", 1),
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		Username: "test",
		Password: "password",
		Email:    "test@example.com",
	}

	mock.ExpectBegin()
	// 针对 beforeCreate 对 username 的重复检查进行的 select
	mock.ExpectQuery("SELECT").WithArgs(user.Username).WillReturnRows(sqlmock.NewRows([]string{"count(*)"}))
	mock.ExpectExec("INSERT INTO").WithArgs(user.InstanceID, user.CreatedAt, user.UpdatedAt, user.DeletedAt, user.Username, user.Password, user.Email).WillReturnResult(sqlmock.NewResult(1, 1))
	// 针对 afterCreate 对 instance_id 的更新，此处注意需要忽略 updated_at 字段（因为具体时间匹配不上，可以直接忽略）
	mock.ExpectExec("UPDATE").WithArgs(user.InstanceID, sqlmock.AnyArg(), 1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := userData.Create(ctx, user)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
