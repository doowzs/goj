package template

import (
	"bufio"
	"fmt"
	"goj/compile"
	"goj/file"
	"gopkg.in/russross/blackfriday.v2"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	time2 "time"
)

func Generate(f *os.File, path string) error {
	t, v, err := NewTemplateWithViper(path)
	if err != nil {
		return err
	}

	title  := v.GetString("problem.title")
	time   := v.GetInt("limits.time")
	memory := v.GetInt("limits.memory")
	size   := v.GetInt("testdata.size")
	ow     := v.GetBool("testdata.overwrite")

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

	err = GenerateTests(t, ow, size, time, memory)
	if err != nil {
		return err
	}

	err = ParseTests(f, t, size)
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
	if err != nil {
		return err
	}
	err = ParseData(f, t, true, 0)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(f, `]]></sample_input>
		<sample_output><![CDATA[`)
	if err != nil {
		return err
	}
	err = ParseData(f, t, false, 0)
	if err != nil {
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

func GenerateTests(t Template, ow bool, size, time, memory int) error {
	var gen, std string
	log.Println("Compiling programs...")
	gen, err := compile.Compile(t["gen"].Path, t["gen"].Name, t["gen"].Ext)
	if err != nil {
		return err
	}
	log.Println(" - gen:", gen)

	std, err = compile.Compile(t["std"].Path, t["std"].Name, t["std"].Ext)
	if err != nil {
		return err
	}
	log.Println(" - std:", std)

	log.Println("Generating input files... overwrite", ow)
	for i := 1; i <= size; i++ {
		name := t["test-in"].Path + t["test-in"].Name + strconv.Itoa(i) + t["test-in"].Ext
		notExist, _ := file.NotExist(name)
		if ow || notExist {
			time2.Sleep(time2.Second)
			fo, err := file.OpenAndTruncate(name, os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return err
			}

			cmd := exec.Command(gen)
			cmd.Stdout = fo
			err = cmd.Run()
			if err != nil {
				return err
			}

			err = fo.Close()
			if err != nil {
				return err
			}
			log.Println(" -", name, "OK")
		} else {
			log.Println(" -", name, "skipped")
		}
	}

	log.Println("Generating output files...")
	for i := 0; i <= size; i++ {
		iname := t["test-in"].Path  + t["test-in"].Name  + strconv.Itoa(i) + t["test-in"].Ext
		oname := t["test-out"].Path + t["test-out"].Name + strconv.Itoa(i) + t["test-out"].Ext
		fi, err := os.Open(iname)
		if err != nil {
			return err
		}
		fo, err := file.OpenAndTruncate(oname, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}

		cmd := exec.Command(std)
		cmd.Stdin  = fi
		cmd.Stdout = fo
		err = cmd.Run()
		if err != nil {
			return err
		}

		err = fi.Close()
		if err != nil {
			return err
		}

		err = fo.Close()
		if err != nil {
			return err
		}
		log.Println(" -", oname, "OK")
	}

	err = os.Remove(gen)
	if err != nil {
		return err
	}
	return os.Remove(std)
}

func ParseData(f *os.File, t Template, isInput bool, no int) error {
	fi, err := os.Open(GetData(t, isInput, no))
	if err != nil {
		return err
	}
	_, err = io.CopyN(f, bufio.NewReader(fi), 1024)
	if err != io.EOF {
		return err
	}
	return fi.Close()
}

func ParseTests(f *os.File, t Template, size int) error {
	_, err := fmt.Fprintf(f, `<!--TEST DATA-->
`)
	if err != nil {
		return err
	}

	for i := 0; i <= size; i++ {
		_, err := fmt.Fprintf(f, `<!--TEST ` + strconv.Itoa(i) + `-->
		<test_input><![CDATA[`)
		if err != nil {
			return err
		}
		err = ParseData(f, t, true, i)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintf(f, `]]></test_input>
		<test_output><![CDATA[`)
		if err != nil {
			return err
		}
		err = ParseData(f, t, false, i)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintf(f, `]]></test_output>
`)
		if err != nil {
			return err
		}
	}
	return nil
}