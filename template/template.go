package template

type File struct {
	Path    string
	Name    string
	Ext     string
	Content string
}

func NewTemplate(path string) map[string]File {
	template := make(map[string]File)
	template["config"] = File{path + "/", "config", ".toml", `[[problem]]
title  = "example"
hint   = false
source = ""

[[limits]]
time   = 1   # in seconds
memory = 256 # in MiB

[[testdata]]
size     = 10     # excluding sample
generate = true
language = "C++"`}
	template["std"] = File{path + "/", "std", "testdata.language", `#include <bits/stdc++.h>
int main() {
    /* please use stdin and stdout */
    return 0;
}`}
	template["gen"] = File{path + "/", "gen", "testdata.language", `#include <bits/stdc++.h>
int main() {
    /* please use stdin and stdout */
    return 0;
}`}
	template["description"] = File{path + "/problem/", "description", ".md", ``}
	template["input"]       = File{path + "/problem/", "input",       ".md", ``}
	template["output"]      = File{path + "/problem/", "output",      ".md", ``}
	template["hint"]        = File{path + "/problem/", "hint",        ".md", ``}
	template["test-in"]     = File{path + "/data/",    "test",        ".in", ``}
	template["test-out"]    = File{path + "/data/",    "test",        ".out", ``}
	return template
}