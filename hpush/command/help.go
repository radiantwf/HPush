package command

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

type HelpCommand struct {
	// Flag is a set of flags specific to this command.
	flag   flag.FlagSet
	topic  string
	Father *Commander `inject:""`
}

func (c *HelpCommand) Init() {
}

// Name returns the command's name: the first word in the usage line.
func (c *HelpCommand) Name() string {
	return "help"
}

// ShortUsage is the short description shown in the 'hpush help' output.
func (c *HelpCommand) ShortUsage() string {
	return "print hpush help"
}

func (c *HelpCommand) Description() string {
	return ""
}

func (c *HelpCommand) Example() string {
	return "hpush help"
}

func (c *HelpCommand) Usage() {
	c.flag.Usage = c.topUsage
	c.flag.Usage()
}

func (c *HelpCommand) Run(args []string) (err error) {
	c.flag.Parse(args)
	newArgs := c.flag.Args()

	c.flag.Usage = c.topUsage
	if len(newArgs) > 1 {
		c.flag.Usage = c.manyParasUsage
	} else if len(newArgs) == 1 {
		c.topic = newArgs[0]
		c.flag.Usage = c.subUsage
	} else {
	}
	c.flag.Usage()
	return
}
func (c *HelpCommand) manyParasUsage() {
	fmt.Fprintf(os.Stderr, "usage: hpush help command\n\nToo many arguments given.\n")
	c.topUsage()
}

func (c *HelpCommand) subUsage() {
	for _, cmd := range c.Father.Commands {
		if c.topic == cmd.Name() {
			cmd.Usage()
			return
		}
	}
	fmt.Fprintf(os.Stderr, "Unknown help topic %#q.  Run 'hpush help'.\n", c.topic)
}

func (c *HelpCommand) topUsage() {
	c.tmpl(os.Stderr, topUsageTemplate, c.Father.Commands)
}

const topUsageTemplate = `
HPush Service: High speed push service

Usage:

    hpush command [arguments]

The commands are:
{{range .}}
    {{.Name | printf "%-11s"}} {{.ShortUsage}}{{end}}

Use "hpush help [command]" for more information about a command.

`

func (c *HelpCommand) tmpl(w io.Writer, text string, data interface{}) {
	t := template.New("top")
	t.Funcs(template.FuncMap{"trim": strings.TrimSpace, "capitalize": c.capitalize})
	template.Must(t.Parse(text))
	if err := t.Execute(w, data); err != nil {
		panic(err)
	}
}

func (c *HelpCommand) capitalize(s string) string {
	if s == "" {
		return s
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToTitle(r)) + s[n:]
}
