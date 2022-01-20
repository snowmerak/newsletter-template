package newsletter

import (
	"bytes"
	"io"
	"os"
	"text/template"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"gopkg.in/yaml.v3"
)

type Newsletter struct {
	Title    string   `yaml:"title"`
	Date     string   `yaml:"date"`
	Articles []string `yaml:"articles"`
	Template string   `yaml:"template"`
}

func New(w io.Writer) error {
	newsletter := Newsletter{
		Date: time.Now().Format("2006-01-02"),
		Template: `<tr>
		<td align="center">
			<h1>{{.Title}}</h1>
		</td>
	</tr>
	<tr>
		<td align="center">
			<p>{{.Date}}</p>
		</td>
	</tr>`,
		Articles: []string{},
	}
	if err := survey.Ask([]*survey.Question{
		{
			Name: "Title",
			Prompt: &survey.Input{
				Message: "Title:",
			},
		},
	}, &newsletter); err != nil {
		return err
	}
	encoder := yaml.NewEncoder(w)
	encoder.SetIndent(2)
	if err := encoder.Encode(newsletter); err != nil {
		return err
	}
	return nil
}

func Load(name string) (*Newsletter, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	newsletter := Newsletter{}
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&newsletter); err != nil {
		return nil, err
	}
	return &newsletter, nil
}

func (n *Newsletter) Parse() string {
	template, err := template.New("newsletter").Parse(n.Template)
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(nil)
	if err := template.Execute(buf, n); err != nil {
		panic(err)
	}
	return buf.String()
}
