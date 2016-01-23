package main

import (
	"io"
	"net/http"
)

const (
	imageGIF = "image/gif"
	imagePNG = "image/png"
	imageJPG = "image/jpeg"
)

func getContentType(r io.Reader) string {
	buff := make([]byte, 512)
	r.Read(buff)
	return http.DetectContentType(buff)
}

func isSupported(contentType string) bool {
	return contentType == imageGIF || contentType == imageJPG || contentType == imagePNG
}
