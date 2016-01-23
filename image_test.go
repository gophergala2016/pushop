package main

import (
	"bytes"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUtils(t *testing.T) {
	Convey("Given text file", t, func() {
		buffer := bytes.NewBuffer([]byte("GIF87a"))
		Convey("ContentType is correctly identified", func() {
			So(getContentType(buffer), ShouldEqual, "image/gif")
		})
	})
}
