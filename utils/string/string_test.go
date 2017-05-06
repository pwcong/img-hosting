package string

import "testing"

func TestPickFilenameFromContentDisposition(t *testing.T) {

	str := "form-data; name=\"img\"; filename=\"59432931_p0.jpg\""
	if o := PickFilenameFromContentDisposition(str); o != "59432931_p0.jpg" {
		t.Errorf("correct is %s, output: %s", "59432931_p0.jpg", o)
	}

	str = "sssssssssss"
	if o := PickFilenameFromContentDisposition(str); o != "" {
		t.Errorf("correct is %s, output: %s", "", o)
	}

}
