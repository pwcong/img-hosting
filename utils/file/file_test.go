package file

import (
	"fmt"
	"testing"
)

func TestGetExtension(t *testing.T) {

	filename := "test.jpg"
	if o := GetExtension(filename); o != "jpg" {

		t.Errorf("correct is %s, output: %s", "jpg", o)
	}

	filename = "test"
	if o := GetExtension(filename); o != "" {
		fmt.Println(o)
		t.Errorf("correct is %s, output: %s", "<empty>", o)
	}

}
