package article

import (
	"bytes"
	"io"
	"os"
	"text/template"

	"github.com/AlecAivazis/survey/v2"
	"gopkg.in/yaml.v3"
)

type Article struct {
	Title   string   `yaml:"title"`
	Author  string   `yaml:"author"`
	Tags    []string `yaml:"tags"`
	Image   string   `yaml:"image"`
	Link    string   `yaml:"link"`
	Content string   `yaml:"Content"`

	Template string `yaml:"template"`
}

func New(w io.Writer) error {
	article := Article{
		Tags: []string{},
		Template: `<tr>
		<td align="center">
			<h2>{{.Title}}</h2>
		</td>
	</tr>
	<tr>
		<td align="center">
			<em>{{.Author}}</em>
		</td>
	</tr>
	<tr>
		<td align="center">
			<img src="{{.Image}}" alt="{{.Title}}" width="380">
		</td>
	</tr>
	<tr>
		<td align="center">
			<a href="{{.Link}}">symbolic link</a>
		</td>
	</tr>
	<tr>
		<td>
			<p>{{.Content}}</p>
		</td>
	</tr>`,
	}
	if err := survey.Ask([]*survey.Question{
		{
			Name: "Title",
			Prompt: &survey.Input{
				Message: "Title:",
			},
		},
		{
			Name: "Author",
			Prompt: &survey.Input{
				Message: "Author:",
			},
		},
		{
			Name: "Image",
			Prompt: &survey.Input{
				Message: "Image URL:",
			},
		},
		{
			Name: "Link",
			Prompt: &survey.Input{
				Message: "Link:",
			},
		},
	}, &article); err != nil {
		return err
	}
	encoder := yaml.NewEncoder(w)
	encoder.SetIndent(2)
	if err := encoder.Encode(article); err != nil {
		return err
	}
	return nil
}

func Load(name string) (Article, error) {
	f, err := os.Open(name)
	if err != nil {
		return Article{}, err
	}
	defer f.Close()
	article := Article{}
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&article); err != nil {
		return Article{}, err
	}
	return article, nil
}

func (a *Article) Parse() string {
	template, err := template.New("article").Parse(a.Template)
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(nil)
	if err := template.Execute(buf, a); err != nil {
		panic(err)
	}
	return buf.String()
}
