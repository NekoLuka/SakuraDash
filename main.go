package main

import (
	"embed"
	"gopkg.in/yaml.v3"
	"html/template"
	"log"
	"net/http"
	"os"
	"slices"
	"strings"
)

var DEVELOPMENT string

//go:embed static
var staticFiles embed.FS

//go:embed index.html
var indexTemplateFile string
var indexTemplate *template.Template

var config Config

func init() {
	DEVELOPMENT = os.Getenv("SAKURADASH_DEVELOPMENT")

	configFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal("Could not import config.yaml")
	}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatal("Could not unmarshall config.yaml")
	}

	// Parse the index template when not in development mode
	if DEVELOPMENT != "1" {
		tmpl, err := template.New("index").Funcs(template.FuncMap{
			"isVisible": isVisible,
			"add":       add,
		}).Parse(indexTemplateFile)
		if err != nil {
			log.Fatal("Could not parse index.html")
		}
		indexTemplate = tmpl
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templateContext := map[string]any{
		"Config":     config,
		"UserGroups": strings.Split(strings.Join(r.Header["Remote-Groups"], ","), ","),
	}

	// When in development mode reload the index file on every request
	if DEVELOPMENT == "1" {
		content, err := os.ReadFile("index.html")
		if err != nil {
			log.Fatal("Could not import index.html")
		}
		indexTemplateFile = string(content)

		tmpl, err := template.New("index").Funcs(template.FuncMap{
			"isVisible": isVisible,
			"add":       add,
		}).Parse(indexTemplateFile)
		if err != nil {
			log.Fatal("Could not parse index.html", err)
		}
		tmpl.Execute(w, templateContext)
	} else {
		indexTemplate.Execute(w, templateContext)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/static/", http.FileServerFS(staticFiles))
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func isVisible(slice []string, item string) bool {
	if item == "" {
		return true
	}
	return slices.Contains(slice, item)
}

func add(a int, b int) int {
	return a + b
}
