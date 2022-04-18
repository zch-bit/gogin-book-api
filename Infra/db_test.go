package Infra

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const TestDb = "./test.db"

func TestConnectDatabase(t *testing.T) {
	asserts := assert.New(t)

	// Test create and close DB
	db := ConnectDatabase()
	_, err := os.Stat(TestDb)
	asserts.NoError(err, "DB should exist")
	asserts.NoError(db.DB().Ping(), "DB should be able to ping")

	// Test get a connecting from connection pools
	connection := GetDB()
	asserts.NoError(connection.DB().Ping(), "Db should be able to ping")
	db.Close()
	clean()
}

func clean() {
	_ = os.Remove(TestDb)
}
