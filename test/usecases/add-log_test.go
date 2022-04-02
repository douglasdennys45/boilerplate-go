package usecases_test

import (
	"clean-architeture/lib/usecases"
	mock_repositories "clean-architeture/test/ports/repositories"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddLogSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_repositories.NewMockAddLogRepository(ctrl)
	repo.EXPECT().Add(gomock.Any()).Return(gomock.Any()).AnyTimes()

	service := usecases.NewAddLog(repo)
	result := service.Add([]byte("type"))
	assert.NotNil(t, result.Data)
	assert.Nil(t, result.Error)
}
