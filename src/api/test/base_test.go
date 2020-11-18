package test

import (
	"os"
	"testing"

	"../app"
	"github.com/mercadolibre/golang-restclient/rest"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	go app.Start()
	os.Exit(m.Run())
}
