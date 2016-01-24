package main

import (
	"html/template"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

const (
	defaultTargetDirectory = "_site"
	maxWidth               = 480
	maxHeight              = 480
)

type Build struct {
	config *Config
	target string
}

type TemplateFile struct {
	OriginalPath  string
	ThumbnailPath string
	ThumbnailURL  string
	ImagePath     string
	ImageURL      string
	HtmlPath      string
	HtmlURL       string
	Name          string
	Description   string
}

type TemplateFiles []*TemplateFile

func NewBuild(target string, config *Config) *Build {
	return &Build{config, target}
}

func (b *Build) Generate(source string) error {

	fileList, err := b.collectFiles(source, b.target)
	if err != nil {
		return err
	}

	// Clean the project
	b.Clean()
	b.createDirectory(b.target)
	originalImagesPath := path.Join(b.target, "i", "o")
	b.createDirectory(originalImagesPath)
	b.copyImages(source, originalImagesPath, fileList)
	thumbnailImagesPath := path.Join(b.target, "i", "t")
	b.createDirectory(thumbnailImagesPath)
	if err := b.generateThumbnails(source, thumbnailImagesPath, fileList); err != nil {
		panic(err)
	}

	indexFile, err := os.Create(path.Join(b.target, "index.html"))
	if err != nil {
		return err
	}
	defer indexFile.Close()
	err = b.frontPage(fileList, indexFile)
	if err != nil {
		return err
	}

	pagesPath := path.Join(b.target, "image")
	b.createDirectory(pagesPath)
	return b.generateImagePages(pagesPath, fileList)
}

func (b *Build) Clean() error {
	return os.RemoveAll(b.target)
}

func (b *Build) createDirectory(path string) error {
	return os.MkdirAll(path, 0766)
}

func (b *Build) copyImages(source, imagesPath string, fileList TemplateFiles) error {
	for _, templateFile := range fileList {
		if err := copyFile(templateFile.OriginalPath, templateFile.ImagePath); err != nil {
			return err
		}
	}
	return nil
}

func (b *Build) generateThumbnails(source, imagesPath string, fileList TemplateFiles) error {
	for _, templateFile := range fileList {
		if err := generateThumbnail(templateFile.OriginalPath, templateFile.ThumbnailPath); err != nil {
			return err
		}
	}
	return nil
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

func (b *Build) collectFiles(source, destination string) (TemplateFiles, error) {
	var content TemplateFiles
	files, err := ioutil.ReadDir(source)
	if err != nil {
		return nil, err
	}
	for _, fileInfo := range files {
		if strings.HasPrefix(fileInfo.Name(), skipPrefix) || strings.HasPrefix(fileInfo.Name(), hiddenFilePrefix) {
			continue
		}
		if fileInfo.IsDir() {
			// Skip directories
			continue
		}
		file, err := os.Open(fileInfo.Name())
		if err != nil {
			continue
		}
		defer file.Close()
		if !isSupported(getContentType(file)) {
			continue
		}
		fileName := fileInfo.Name()
		templateFile := &TemplateFile{
			Name:          fileName,
			Description:   "",
			OriginalPath:  path.Join(source, fileName),
			ThumbnailPath: path.Join(destination, "i", "t", fileName),
			ThumbnailURL:  "/i/t/" + fileName,
			ImagePath:     path.Join(destination, "i", "o", fileName),
			ImageURL:      "/i/o/" + fileName,
			HtmlPath:      path.Join(destination, "image", getImageFileName(fileName)),
			HtmlURL:       "/image/" + getImageFileName(fileName),
		}
		if _, ok := b.config.Content[fileInfo.Name()]; ok {
			file := b.config.Content[fileInfo.Name()]
			templateFile.Name = file.Name
			templateFile.Description = file.Description
		}
		content = append(content, templateFile)
	}
	return content, nil
}

func (b *Build) frontPage(fileList TemplateFiles, w io.Writer) error {
	layoutData, err := Asset("assets/templates/layout.html")
	if err != nil {
		return err
	}
	indexData, err := Asset("assets/templates/index.html")
	if err != nil {
		return err
	}
	t, err := template.New("layout").Parse(string(layoutData))
	t.Parse(string(indexData))
	if err != nil {
		return err
	}
	templateData := map[string]interface{}{
		"Title":       b.config.Title,
		"Description": b.config.Description,
		"Analytics":   b.config.Analytics,
		"Images":      fileList,
	}
	return t.Execute(w, templateData)
}

func (b *Build) generateImagePages(pagesPath string, tf TemplateFiles) error {
	layoutData, err := Asset("assets/templates/layout.html")
	if err != nil {
		return err
	}
	imageData, err := Asset("assets/templates/image.html")
	if err != nil {
		return err
	}
	t, err := template.New("layout").Parse(string(layoutData))
	if err != nil {
		return err
	}
	t.Parse(string(imageData))

	for _, templateFile := range tf {
		if err = b.generateaImagePage(templateFile, t); err != nil {
			return err
		}
	}
	return nil
}

func getImageFileName(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + ".html"
}

func (b *Build) generateaImagePage(tf *TemplateFile, t *template.Template) error {
	imageFile, err := os.Create(tf.HtmlPath)
	if err != nil {
		return err
	}
	defer imageFile.Close()

	templateData := map[string]interface{}{
		"Title":       b.config.Title,
		"Description": b.config.Description,
		"Analytics":   b.config.Analytics,
		"Image":       tf,
	}
	return t.Execute(imageFile, templateData)
}
