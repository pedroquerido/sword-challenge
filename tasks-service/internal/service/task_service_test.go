package service_test

import (
	"os"
	"github.com/pedroquerido/sword-challenge/tasks-service/internal/service"
	"testing"

	"gopkg.in/go-playground/validator.v9"
)

func TestMain(m *testing.M) {

	testTaskValidator = service.NewTaskValidator(validator.New())

	os.Exit(m.Run())
}
