package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

const (
	setFlag          = "set"
	dashPrefix       = "-"
	doubleDashPrefix = "--"
)

var (
	// ErrWrongFlags is returned when arguments are wrong
	ErrWrongFlags = fmt.Errorf("Wrong arg")
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You must pass at least the file path.")
		fmt.Println("Example: ")
		fmt.Println("> gen -foo -bar=42 template.txt")
		os.Exit(-1)
	}
	vals, err := parseFlags(os.Args[1 : len(os.Args)-1])
	checkError(err)
	content, err := ioutil.ReadFile(os.Args[len(os.Args)-1])
	checkError(err)
	renderedTemplate, err := renderTemplate(string(content), vals)
	checkError(err)
	fmt.Print(renderedTemplate)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

// parseFlags flags from cli input
func parseFlags(args []string) (map[string]string, error) {
	flags := make(map[string]string)

	var prevFlag = ""
	for _, arg := range args {
		if strings.HasPrefix(arg, doubleDashPrefix) || strings.HasPrefix(arg, dashPrefix) {
			flag := strings.Split(arg, "=")
			if len(flag) > 1 {
				flags[strings.TrimLeft(flag[0], dashPrefix)] = flag[1]
			} else {
				flags[strings.TrimLeft(flag[0], dashPrefix)] = setFlag
			}
			prevFlag = flag[0]
		} else {
			if strings.HasPrefix(prevFlag, doubleDashPrefix) || strings.HasPrefix(prevFlag, dashPrefix) {
				flags[strings.TrimLeft(prevFlag, dashPrefix)] = arg
			} else {
				return nil, ErrWrongFlags
			}
		}
	}
	return flags, nil
}

// renderTemplate renders a template with the given data
func renderTemplate(tmpl string, data map[string]string) (string, error) {
	t := template.Must(template.New("tmp").Parse(tmpl))
	var b bytes.Buffer
	err := t.Execute(&b, data)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}
