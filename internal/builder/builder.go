package builder

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/avelino/slugify"
	"github.com/olebedev/config"
	"github.com/russross/blackfriday"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Builder struct {
	Logger *zap.Logger
}

type PageData struct {
	Name      string
	Post      string
	Entries   []EntryData
	Now       string
	Copyright string
}

type EntryData struct {
	Name      string
	Title     string
	Author    string
	Date      string
	Post      template.HTML
	Now       string
	Copyright string
}

func (b *Builder) Start(settings map[string]string) error {
	path := getCurrentPath()
	entryPath := path + "/tales/entries/*.md"
	b.Logger.Sugar().Infof("Globbing entries from %s", entryPath)
	files, err := filepath.Glob(entryPath)
	if err != nil {
		b.Logger.Sugar().Error(err)
		return err
	}

	indexTemplatePath := path + "/tales/layout/_main.html"
	entryTemplatePath := path + "/tales/layout/_entry.html"
	b.Logger.Sugar().Infof("Loading layouts from %s", entryTemplatePath)

	indexData := PageData{
		Name:      settings["name"],
		Now:       time.Now().Format(time.RFC3339),
		Copyright: time.Now().Format("2006"),
	}

	for _, file := range files {
		entryTemplate := template.Must(template.ParseFiles(entryTemplatePath))

		input, _ := ioutil.ReadFile(file)
		entryDetails, post := parseDetails(string(input))
		output := blackfriday.MarkdownBasic([]byte(post))

		entryData := EntryData{
			Name:      settings["name"],
			Title:     entryDetails["title"],
			Author:    entryDetails["author"],
			Date:      entryDetails["date"],
			Post:      template.HTML(trim(string(output), 200)),
			Now:       time.Now().Format(time.RFC3339),
			Copyright: time.Now().Format("2006"),
		}

		indexData.Entries = append(indexData.Entries, entryData)

		destPath := fmt.Sprintf(path+"/public/%s.html", slugify.Slugify(entryDetails["title"]))
		b.Logger.Sugar().Infof("Saving entry to %s", destPath)
		f, _ := os.Create(destPath)
		err = entryTemplate.Execute(f, entryData)
		if err != nil {
			panic(err)
		}
		f.Close()
	}

	indexTemplate := template.Must(template.ParseFiles(indexTemplatePath))

	indexPath := fmt.Sprintf(path + "/public/index.html")
	b.Logger.Sugar().Infof("Creating index file... %s", indexPath)
	f, _ := os.Create(indexPath)
	err = indexTemplate.Execute(f, indexData)
	if err != nil {
		panic(err)
	}
	f.Close()

	return nil
}

func trim(str string, length int) string {
	if len(str) < length {
		return str
	}

	idx := strings.LastIndex(str, " ")
	str = str[:idx] + "..."

	return str
}

func parseDetails(entry string) (map[string]string, string) {
	var yaml string
	start := strings.Index(entry, "---")
	if start > -1 {
		entry = entry[start+3:]
		end := strings.Index(entry, "---")
		if end > -1 {
			yaml = entry[start:end]
			entry = entry[end:]
		}
	}

	cfg, err := config.ParseYaml(yaml)
	if err != nil {
		panic(err)
	}

	results, _ := cfg.Map("")

	details := make(map[string]string)
	for key, val := range results {
		details[key] = val.(string)
	}

	return details, entry
}

func getCurrentPath() string {
	if path := viper.GetString("base_path"); path != "" {
		return path
	}
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	return exPath
}
