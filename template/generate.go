package template

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strconv"
)

func Generate(f *os.File, path string) error {
	_, err := fmt.Fprintf(f, `<?xml version="1.0" encoding="UTF-8"?>
<!--BEGIN FPS XML-->
<fps version="1.2" url="https://github.com/zhblue/freeproblemset/">
	<generator name="GOJ-` + Version + `" url="https://git.doowzs.com/doowzs/goj"/>
	<item>
`)
	if err != nil {
		return err
	}

	c, err := os.Open(path + "/config.toml")
	if err != nil {
		return err
	}
	defer c.Close()

	v := viper.New()
	v.SetConfigType("toml")
	err = v.ReadConfig(c)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(f, `<!--INFORMATION-->
		<title><![CDATA[` + v.GetString("problem.title") + `]]></title>
		<time_limit unit="s"><![CDATA[` + strconv.Itoa(v.GetInt("limits.time")) + `]]></time_limit>
		<memory_limit unit="mb"><![CDATA[` + strconv.Itoa(v.GetInt("limits.memory")) + `]]></memory_limit>
`)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(f, `<!--DESCRIPTIONS-->
`)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(f, `<!--SAMPLE INPUT/OUTPUT-->
`)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(f, `<!--OPTIONAL HINT-->
`)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(f, `<!--TEST INPUT/OUTPUT-->
`)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(f, `<!--END OF FPS XML-->
	</item>
</fps>
`)
	if err != nil {
		return err
	}
	return nil
}