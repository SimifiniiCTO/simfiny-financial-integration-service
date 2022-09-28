package plaidhandler

import (
	"log"
	"os"
	"testing"
)

// setup sets up a database connection to the test db node
func setup() {
	var err error
	plaidTestClient, err = GetPlaidWrapperForTest()
	if err != nil {
		log.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}
