package main

import (
	"bytes"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConfig(t *testing.T) {
	Convey("Given config", t, func() {
		config := NewConfig()
		Convey("Config is not nil", func() {
			So(config, ShouldNotBeNil)
		})
		Convey("Config is initialized", func() {
			So(config.Title, ShouldEqual, defaultTitle)
			So(config.Description, ShouldEqual, defaultDescription)
		})
		Convey("Loading config works", func() {
			testString := `title: Test Title
description: Test Description`
			buffer := bytes.NewBufferString(testString)
			err := config.Load(buffer)
			So(err, ShouldBeNil)
			So(config.Title, ShouldEqual, "Test Title")
			So(config.Description, ShouldEqual, "Test Description")
		})
		Convey("Saving config works", func() {
			expected := `title: Save Name
description: Save Description
analytics: ""
content:
  test:
    name: test
    description: descr
`
			buffer := bytes.NewBufferString("")
			config.Title = "Save Name"
			config.Description = "Save Description"
			config.Content["test"] = &File{
				Name:        "test",
				Description: "descr",
			}
			err := config.Save(buffer)
			So(err, ShouldBeNil)
			So(buffer.String(), ShouldEqual, expected)
		})
	})
}
