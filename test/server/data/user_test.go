package data

import (
	"context"
	"testing"

	"github.com/lllllan02/iam/internal/model"
	"github.com/lllllan02/iam/pkg/code"
	"github.com/lllllan02/iam/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestUserData_Create(t *testing.T) {

	tests := []struct {
		name        string
		user        model.User
		wantErr     bool
		wantErrCode int
	}{
		{"1", model.User{Username: "success", Password: "password", Email: "success@test.com"}, false, 0},
		{"2", model.User{Username: "err_name", Password: "password", Email: "fail@test.com"}, true, code.CInvalidUsername},
		{"3", model.User{Username: "success", Password: "password", Email: "fail@test.com"}, true, code.CDuplicateUsername},
		{"4", model.User{Username: "fail", Password: "password", Email: "err@email"}, true, code.CInvalidEmail},
		{"5", model.User{Username: "fail", Password: "password", Email: "success@test.com"}, true, code.CDuplicaEmail},
		{"6", model.User{Username: "sec", Password: "password", Email: "sec@test.com"}, false, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := userData.Create(context.Background(), &tt.user); (err != nil) != tt.wantErr {
				assert.Equal(t, tt.wantErrCode, errors.Code(err))
			}
		})
	}
}
