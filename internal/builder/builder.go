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
	"go.uber.org/zap"
)

type Builder struct {
	Logger *zap.Logger
}

func (b *Builder) Start(settings map[string]string) error {
	// TODO: make this a relative path
	files, err := filepath.Glob("/Users/brianseitel/Code/go/src/github.com/brianseitel/gleeman/tales/entries/*.md")
	if err != nil {
		b.Logger.Sugar().Error(err)
		return err
	}

	layoutFile, _ := ioutil.ReadFile("/Users/brianseitel/Code/go/src/github.com/brianseitel/gleeman/tales/layout/_layout.html")

	layout := string(layoutFile)

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

		f, _ := os.Create(fmt.Sprintf("./public/%s.html", slugify.Slugify(entryDetails["title"])))
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
