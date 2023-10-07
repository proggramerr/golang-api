package unit

import (
	"rest_api/app"
	"testing"
)

func TestDBConnection(t *testing.T) {
	app.GetDBEngine()
}
