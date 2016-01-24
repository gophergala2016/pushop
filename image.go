package main

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"

	"github.com/nfnt/resize"
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

func copyFile(source, destination string) error {
	fr, err := os.Open(source)
	if err != nil {
		return err
	}
	defer fr.Close()
	fw, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer fw.Close()
	_, err = io.Copy(fw, fr)
	return err
}

func generateThumbnail(source, destination string) error {
	file, err := os.Open(source)
	if err != nil {
		return err
	}
	var img image.Image
	var imageType string
	img, imageType, err = image.Decode(file)
	if err != nil {
		return err
	}
	file.Close()

	m := resize.Thumbnail(maxWidth, maxHeight, img, resize.Lanczos3)

	out, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer out.Close()
	switch imageType {
	case "gif":
		return gif.Encode(out, m, nil)
	case "jpeg":
		return jpeg.Encode(out, m, nil)
	case "png":
		return png.Encode(out, m)
	}
	return nil
}
