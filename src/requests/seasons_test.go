package requests

import (
	"basket-go-count/src/mocks"
	"testing"
)

func TestDoSeasonRecordsCount(t *testing.T) {
	conn := mocks.Connect()
	count, err := DoSeasonRecordsCount(conn)
	if err != nil {
		t.Error(err)
	}
	t.Log(count)
	err = conn.Close()
	if err != nil {
		t.Error(err)
	}
}
