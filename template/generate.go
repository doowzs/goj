package template

import (
	"bufio"
	"fmt"
	"gopkg.in/russross/blackfriday.v2"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func Generate(f *os.File, path string) error {
	t, v, err := NewTemplateWithViper(path)
	if err != nil {
		return err
	}

	title  := v.GetString("problem.title")
	time   := v.GetInt("limits.time")
	memory := v.GetInt("limits.memory")

	_, err = fmt.Fprintf(f, `<?xml version="1.0" encoding="UTF-8"?>
<!--BEGIN FPS XML-->
<fps version="1.2" url="https://github.com/zhblue/freeproblemset/">
	<generator name="GOJ-` + Version + `" url="https://git.doowzs.com/doowzs/goj"/>
	<item>
<!--INFORMATION-->
		<title><![CDATA[` + title + `]]></title>
		<time_limit unit="s"><![CDATA[` + strconv.Itoa(time) + `]]></time_limit>
		<memory_limit unit="mb"><![CDATA[` + strconv.Itoa(memory) + `]]></memory_limit>
`)
	if err != nil {
		return err
	}

	err = ParseDescriptions(f, t)
	if err != nil {
		return err
	}

	err = ParseSamples(f, t)
	if err != nil {
		return err
	}

	if v.GetBool("problem.hint") {
		err = ParseHint(f, t)
		if err != nil {
			return err
		}
	}

	err = ParseTests(f, t, time, memory)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(f, `<!--END OF FPS XML-->
	</item>
</fps>
`)
	return err
}

func ParseMarkdownFile(f *os.File, t Template, i string) error {
	data, err := ioutil.ReadFile(t[i].Path + t[i].Name + t[i].Ext)
	if err != nil {
		return err
	}
	html := string(blackfriday.Run([]byte(string(data))))
	_, err = fmt.Fprintf(f, `		<` + i + `><![CDATA[` + html + `]]></` + i + `>
`)
	return err
}

func ParseDescriptions(f *os.File, t Template) error {
	_, err := fmt.Fprintf(f, `<!--DESCRIPTIONS-->
`)
	if err != nil {
		return err
	}

	for _, index := range []string{"description", "input", "output"} {
		err = ParseMarkdownFile(f, t, index)
		if err != nil {
			return err
		}
	}
	return nil
}

func ParseSamples(f *os.File, t Template) error {
	_, err := fmt.Fprintf(f, `<!--SAMPLE DATA-->
		<sample_input><![CDATA[`)
	fi, err := os.Open(GetData(t, true, 0))
	if err != nil {
		return err
	}
	defer fi.Close()
	_, err = io.CopyN(f, bufio.NewReader(fi), 1024)
	if err != io.EOF {
		return err
	}
	_, err = fmt.Fprintf(f, `]]></sample_input>
		<sample_output><![CDATA[`)
	fo, err := os.Open(GetData(t, false, 0))
	if err != nil {
		return err
	}
	defer fo.Close()
	_, err = io.CopyN(f, bufio.NewReader(fo), 1024)
	if err != io.EOF {
		return err
	}
	_, err = fmt.Fprintf(f, `]]></sample_output>
`)
	return err
}

func ParseHint(f *os.File, t Template) error {
	_, err := fmt.Fprintf(f, `<!--HINT-->
`)
	if err != nil {
		return err
	}

	return ParseMarkdownFile(f, t, "hint")
}

func ParseTests(f *os.File, t Template, time, memory int) error {
	_, err := fmt.Fprintf(f, `<!--TEST DATA-->
`)
	return err
}