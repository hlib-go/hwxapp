package hwxapp

import (
	"testing"
)

func TestGetPhoneNumber(t *testing.T) {

	r, err := GetPhoneNumber(nil, "", "", "")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(r.JSON())

}
