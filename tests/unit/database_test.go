package unit

import (
	"rest_api/pkg/client/postgres"
	"testing"
)

func TestDBConnection(t *testing.T) {
	postgres.GetPostgresEngine()
}
