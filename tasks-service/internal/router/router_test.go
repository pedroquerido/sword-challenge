package router_test

import (
	"os"
	"github.com/pedroquerido/sword-challenge/tasks-service/internal/router"
	"testing"

	"gopkg.in/go-playground/validator.v9"
)

func TestMain(m *testing.M) {

	testRequestValidator = router.NewRequestValidator(validator.New())

	os.Exit(m.Run())
}
