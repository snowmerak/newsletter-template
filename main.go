package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/snowmerak/gemail/article"
	"github.com/snowmerak/gemail/newsletter"
	"gopkg.in/alecthomas/kingpin.v2"
)

const HEAD = `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
</head>
<body>
    <div style="width: 600px; padding: 0px 0px 5% 5%;">
        <table>
		`

const TAIL = `        </table>
</div>
</body>
</html>`

func main() {
	newsletters := kingpin.New("newletter generator", "Generate newsletters from articles")
	newsletterNew := newsletters.Command("new", "new a newsletter")
	newsletterNewName := newsletterNew.Arg("name", "Name of the newsletter").Required().String()

	newsletterGenerate := newsletters.Command("generate", "generate a newsletter")
	newsletterGenerateName := newsletterGenerate.Arg("name", "Name of the newsletter").Required().String()

	articles := newsletters.Command("article", "Generate an article")
	articleNew := articles.Command("new", "Create a new article")
	articleNewName := articleNew.Arg("name", "Name of the article").Required().String()

	lettersPath := filepath.Join(".", "newsletter", "letters")
	if err := os.MkdirAll(lettersPath, 0755); err != nil {
		panic(err)
	}
	contentsPath := filepath.Join(".", "article", "contents")
	if err := os.MkdirAll(contentsPath, 0755); err != nil {
		panic(err)
	}
	distPath := filepath.Join(".", "dist")
	if err := os.MkdirAll(distPath, 0755); err != nil {
		panic(err)
	}

	switch kingpin.MustParse(newsletters.Parse(os.Args[1:])) {
	case newsletterGenerate.FullCommand():
		name := *newsletterGenerateName
		if name == "" {
			panic(errors.New("name of the newsletter is required"))
		}
		if !strings.HasSuffix(name, ".yaml") {
			name += ".yaml"
		}
		path := filepath.Join(lettersPath, name)
		newsletter, err := newsletter.Load(path)
		if err != nil {
			panic(err)
		}
		distFile := filepath.Join(distPath, strings.TrimSuffix(name, ".yaml")+".html")
		f, err := os.Create(distFile)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		fmt.Fprint(f, HEAD)
		parsed := newsletter.Parse()
		fmt.Fprint(f, parsed)
		for _, a := range newsletter.Articles {
			article, err := article.Load(filepath.Join(contentsPath, a) + ".yaml")
			if err != nil {
				panic(err)
			}
			fmt.Fprint(f, article.Parse())
		}
		fmt.Fprint(f, TAIL)
	case newsletterNew.FullCommand():
		name := *newsletterNewName
		if name == "" {
			panic(errors.New("Name of the newsletter is required"))
		}
		if !strings.HasSuffix(name, ".yaml") {
			name += ".yaml"
		}
		fmt.Println("Generating newsletter: ", name)
		f, err := os.Create(filepath.Join(lettersPath, name))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if err := newsletter.New(f); err != nil {
			panic(err)
		}
		fmt.Println("Generated newsletter: ", name)
	case articleNew.FullCommand():
		name := *articleNewName
		if name == "" {
			panic(errors.New("Name of the article is required"))
		}
		if !strings.HasSuffix(name, ".yaml") {
			name += ".yaml"
		}
		fmt.Println("Generating article: ", name)
		f, err := os.Create(filepath.Join(contentsPath, name))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if err := article.New(f); err != nil {
			panic(err)
		}
		fmt.Println("Generated article: ", name)
	default:
		fmt.Println(newsletters.Help)
	}
}
