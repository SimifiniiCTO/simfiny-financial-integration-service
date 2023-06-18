package plaidhandler

import (
	"log"
	"os"
	"testing"
)

var (
	testAccessToken string
	testClientId    string
	testItemId      string
)

// setup sets up a database connection to the test db node
func setup() {
	var err error
	plaidTestClient, err = GetPlaidWrapperForTest()
	if err != nil {
		log.Fatal(err)
	}

	res, err := plaidTestClient.GetAccessTokenForSandboxAcct()
	if err != nil {
		log.Fatal(err)
	}

	testAccessToken = res.AccessToken
	testItemId = res.ItemId
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}
