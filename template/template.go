package template

import (
	"github.com/spf13/viper"
	"os"
	"strconv"
	"strings"
)

type File struct {
	Path    string
	Name    string
	Ext     string
	Content string
}

type Template map[string]File

var (
	Version string
	Order = []string{
		"_root",
		"config",
		"gen",
		"std",
		"_problem",
		"description",
		"input",
		"output",
		"hint",
		"_data",
		"sample-in",
		"sample-out",
		// "test-in",
		// "test-out",
	}
)

func NewTemplate(path string) Template {
	template := make(Template)
	template["_root"]  = File{path, "", "", ``}
	template["config"] = File{path + "/", "config", ".toml", `[[problem]]
title  = "` + path[strings.LastIndexByte(path, '/') + 1 : ] + `"
hint   = false    # set to true to give hints defined in hint.md
source = ""

[[limits]]
time   = 1        # measured in seconds
memory = 256      # measured in MiB

[[testdata]]
size      = 10     # excluding the sample case (e.g. 10 means [0-10] => 11 tests)
overwrite = false  # set to true to force overwriting existing test data files
language  = "cpp"  # same as the source file extension, i.e. "c", "cpp", "java", etc.`}
	template["gen"] = File{path + "/", "gen", "testdata.language", `#include <bits/stdc++.h>
int main() {
    /* please use stdin and stdout */
    return 0;
}`}
	template["std"] = File{path + "/", "std", "testdata.language", `#include <bits/stdc++.h>
int main() {
    /* please use stdin and stdout */
    return 0;
}`}
	template["_problem"]    = File{path + "/problem",  "",            "",     ``}
	template["description"] = File{path + "/problem/", "description", ".md",  ``}
	template["input"]       = File{path + "/problem/", "input",       ".md",  ``}
	template["output"]      = File{path + "/problem/", "output",      ".md",  ``}
	template["hint"]        = File{path + "/problem/", "hint",        ".md",  ``}
	template["_data"]       = File{path + "/data",     "",            "",     ``}
	template["sample-in"]   = File{path + "/data/",    "test0",       ".in",  ``}
	template["sample-out"]  = File{path + "/data/",    "test0",       ".out", ``}
	template["test-in"]     = File{path + "/data/",    "test",        ".in",  ``}
	template["test-out"]    = File{path + "/data/",    "test",        ".out", ``}
	return template
}

func NewTemplateWithViper(path string) (Template, *viper.Viper, error) {
	t := NewTemplate(path)

	c, err := os.Open(t["config"].Path + t["config"].Name + t["config"].Ext)
	if err != nil {
		return nil, nil, err
	}
	defer c.Close()

	v := viper.New()
	v.SetConfigType("toml")
	err = v.ReadConfig(c)
	if err != nil {
		return nil, nil, err
	}

	for _, index := range []string{"gen", "std"} {
		t[index] = File{
			Path:    t[index].Path,
			Name:    t[index].Name,
			Ext:     "." + v.GetString(t[index].Ext),
			Content: ``,
		}
	}
	return t, v, nil
}

func GetData(t Template, isInput bool, no int) string {
	var direction string
	if isInput {
		direction = "in"
	} else {
		direction = "out"
	}
	return t["test-" + direction].Path + t["test-" + direction].Name + strconv.Itoa(no) + t["test-" + direction].Ext
}