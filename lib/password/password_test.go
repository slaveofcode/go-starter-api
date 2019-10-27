package password

import "testing"

func TestPassword(t *testing.T) {
	hashed, _ := Hash("simplePasswordH4SH")
	if !Compare("simplePasswordH4SH", hashed) {
		t.Error("Password hash doesnt match")
	}
}
