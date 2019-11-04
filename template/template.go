package template

import "strings"

type File struct {
	Path    string
	Name    string
	Ext     string
	Content string
}

var Order = []string{
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
	"test-in",
	"test-out",
}

func NewTemplate(path string) map[string]File {
	template := make(map[string]File)
	template["_root"]  = File{path, "", "", ``}
	template["config"] = File{path + "/", "config", ".toml", `[[problem]]
title  = "` + path[strings.LastIndexByte(path, '/') + 1 : ] + `"
hint   = false    # set to true to give hints defined in hint.md
source = ""

[[limits]]
time   = 1        # measured in seconds
memory = 256      # measured in MiB

[[testdata]]
size     = 10     # excluding the sample case (e.g. 10 means [0-10] => 11 tests)
generate = true   # set to false for custom data instead of generating some
language = "cpp"  # same as the source file extension, i.e. "c", "cpp", "java", etc.`}
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
	template["_problem"]    = File{path + "/problem",  "",            "",    ``}
	template["description"] = File{path + "/problem/", "description", ".md", ``}
	template["input"]       = File{path + "/problem/", "input",       ".md", ``}
	template["output"]      = File{path + "/problem/", "output",      ".md", ``}
	template["hint"]        = File{path + "/problem/", "hint",        ".md", ``}
	template["_data"]       = File{path + "/data",     "",            "",    ``}
	template["test-in"]     = File{path + "/data/",    "test",        ".in", ``}
	template["test-out"]    = File{path + "/data/",    "test",        ".out", ``}
	return template
}