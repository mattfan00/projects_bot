package utils

import (
	"strings"

	"github.com/google/shlex"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	Name        string `short:"n" long:"name" description:"Name of your project" required:"true"`
	Url         string `short:"u" long:"url" description:"GitHub URL of project"`
	Description string `short:"d" long:"description" description:"A short description of the project"`
}

func GetArgs(command string) map[string]string {
	command = strings.ReplaceAll(command, "“", "\"")
	command = strings.ReplaceAll(command, "”", "\"")
	split, _ := shlex.Split(command)

	_, err := flags.ParseArgs(&opts, split)
	if err != nil {
		panic(err)
	}

	return map[string]string{
		"Name":        opts.Name,
		"Url":         opts.Url,
		"Description": opts.Description,
	}
}
