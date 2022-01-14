package builder

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/avelino/slugify"
	"github.com/olebedev/config"
	"github.com/russross/blackfriday"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Builder struct {
	Logger *zap.Logger
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

	layoutPath := path + "/tales/layout/_layout.html"
	b.Logger.Sugar().Infof("Loading layouts from %s", layoutPath)
	layoutFile, _ := ioutil.ReadFile(layoutPath)
	layout := string(layoutFile)

	b.Logger.Sugar().Infof("Populating layout with settings...")
	for setting, value := range settings {
		layout = strings.ReplaceAll(layout, fmt.Sprintf(`{{ .%s }}`, setting), value)
	}

	for _, file := range files {
		input, _ := ioutil.ReadFile(file)
		entryDetails, post := parseDetails(string(input))
		output := blackfriday.MarkdownBasic([]byte(post))

		entry := layout

		for key, value := range entryDetails {
			entry = strings.Replace(entry, fmt.Sprintf(`{{ .%s }}`, key), value, -1)
			post = strings.Replace(post, fmt.Sprintf(`{{ .%s }}`, key), value, 1)
		}

		entry = strings.Replace(entry, `{{ .post }}`, string(output), 1)

		destPath := fmt.Sprintf(path+"/public/%s.html", slugify.Slugify(entryDetails["title"]))
		b.Logger.Sugar().Infof("Saving entry to %s", destPath)
		f, _ := os.Create(destPath)
		f.Write([]byte(entry))
		f.Close()
	}
	return nil
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
