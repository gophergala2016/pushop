package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	defaultTargetSegment   = "_site"
	imagesSegment          = "i"
	originalImagesSegment  = "o"
	thumbnailImagesSegment = "t"
	pagesSegment           = "image"

	maxWidth  = 480
	maxHeight = 480
)

type Build struct {
	projectPath string
	config      *Config
}

type TemplateImage struct {
	OriginalPath  string
	ThumbnailPath string
	ThumbnailURL  string
	ImagePath     string
	ImageURL      string
	HtmlPath      string
	HtmlURL       string
	Name          string
	Description   string
	Next          string
	Previous      string
}

type TemplateImages []*TemplateImage

func NewBuild(projectPath string, config *Config) *Build {
	return &Build{projectPath, config}
}

func (b *Build) Generate(targetPath string) error {
	var err error

	// Clean the project
	if err = cleanProject(targetPath); err != nil {
		return err
	}

	// Collect Files
	var templateImages TemplateImages
	templateImages, err = b.collectFiles(b.projectPath, targetPath)
	if err != nil {
		return err
	}

	// Create directories
	for _, path := range []string{
		targetPath,
		path.Join(targetPath, imagesSegment, originalImagesSegment),
		path.Join(targetPath, imagesSegment, thumbnailImagesSegment),
	} {
		if err = b.createDirectory(path); err != nil {
			return err
		}
	}

	// Copy images and generate thumbnails
	if err = b.generateImages(templateImages); err != nil {
		return err
	}

	// Generate the index page
	if err = b.generateIndex(targetPath, templateImages); err != nil {
		return err
	}

	// Generate the image pages
	pagesPath := path.Join(targetPath, pagesSegment)
	if err = b.createDirectory(pagesPath); err != nil {
		return err
	}
	return b.generateImagePages(pagesPath, templateImages)
}

func cleanProject(targetPath string) error {
	return os.RemoveAll(targetPath)
}

func (b *Build) createDirectory(path string) error {
	return os.MkdirAll(path, 0766)
}

func (b *Build) generateImages(templateImages TemplateImages) error {
	for _, ti := range templateImages {
		if err := copyFile(ti.OriginalPath, ti.ImagePath); err != nil {
			return err
		}
		if err := generateThumbnail(ti.OriginalPath, ti.ThumbnailPath); err != nil {
			return err
		}
	}
	return nil
}

func (b *Build) collectFiles(projectPath, targetPath string) (TemplateImages, error) {
	var templateImages TemplateImages
	files, err := ioutil.ReadDir(projectPath)
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
		ti := &TemplateImage{
			Name:          fileName,
			Description:   "",
			OriginalPath:  path.Join(projectPath, fileName),
			ThumbnailPath: path.Join(targetPath, imagesSegment, thumbnailImagesSegment, fileName),
			ThumbnailURL:  "/" + imagesSegment + "/" + thumbnailImagesSegment + "/" + fileName,
			ImagePath:     path.Join(targetPath, imagesSegment, originalImagesSegment, fileName),
			ImageURL:      "/" + imagesSegment + "/" + originalImagesSegment + "/" + fileName,
			HtmlPath:      path.Join(targetPath, pagesSegment, getImageFileName(fileName)),
			HtmlURL:       "/" + pagesSegment + "/" + getImageFileName(fileName),
		}
		if _, ok := b.config.Content[fileInfo.Name()]; ok {
			file := b.config.Content[fileInfo.Name()]
			ti.Name = file.Name
			ti.Description = file.Description
		}
		templateImages = append(templateImages, ti)
	}
	last := len(templateImages) - 1
	for i, _ := range templateImages {
		templateImages[i].Previous = templateImages[last].HtmlURL
		if i > 0 {
			templateImages[i].Previous = templateImages[i-1].HtmlURL
		}
		templateImages[i].Next = templateImages[0].HtmlURL
		if i < last {
			templateImages[i].Next = templateImages[i+1].HtmlURL
		}
	}

	return templateImages, nil
}

func (b *Build) generateIndex(targetPath string, templateImages TemplateImages) error {
	indexFile, err := os.Create(path.Join(targetPath, "index.html"))
	if err != nil {
		return err
	}
	defer indexFile.Close()
	return b.frontPage(indexFile, templateImages)
}

func (b *Build) frontPage(w io.Writer, templateImages TemplateImages) error {
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
		"Images":      templateImages,
	}
	return t.Execute(w, templateData)
}

func (b *Build) generateImagePages(pagesPath string, templateImages TemplateImages) error {
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

	for _, ti := range templateImages {
		if err = b.generateaImagePage(ti, t); err != nil {
			return err
		}
	}
	return nil
}

func getImageFileName(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + ".html"
}

func (b *Build) generateaImagePage(ti *TemplateImage, t *template.Template) error {
	imageFile, err := os.Create(ti.HtmlPath)
	if err != nil {
		return err
	}
	defer imageFile.Close()

	templateData := map[string]interface{}{
		"Title":       b.config.Title,
		"Description": b.config.Description,
		"Analytics":   b.config.Analytics,
		"Image":       ti,
	}
	return t.Execute(imageFile, templateData)
}

func getTumbnailImagesPath(targetPath string) string {
	return path.Join(targetPath, imagesSegment, thumbnailImagesSegment)
}
