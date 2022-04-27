package solution

import (
	"testing"

	"github.com/kyokomi/emoji"
)

var testCase = emoji.Sprint("Hello :world_map:!")

func TestGetMessage(t *testing.T) {
	result := GetMessage()

	if result != testCase {
		t.Error("Strings not match!")
	}
}
