package utils

import (
	"strings"

	"github.com/google/shlex"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	Name        string `short:"n" long:"name" description:"Name of your project" required:"true"`
	Url         string `short:"u" long:"url" description:"GitHub URL of project"`
	Description string `short:"d" long:"description" description:"A short description of the project"`
}

func GetArgs(command string) (Options, bool) {
	command = strings.ReplaceAll(command, "“", "\"")
	command = strings.ReplaceAll(command, "”", "\"")
	split, _ := shlex.Split(command)

	help := contains(split, "help")

	if !help {
		var opts Options
		_, err := flags.ParseArgs(&opts, split)
		if err != nil {
			panic(err)
		}

		return opts, false
	} else {
		return Options{}, true
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
