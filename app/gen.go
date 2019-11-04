package app

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strconv"
)

func runGen(path string) error {
	f, err := os.Open(path + "/dist.xml")
	if os.IsNotExist(err) {
		f, err = os.Create(path + "/dist.xml")
	}
	if err != nil {
		return err
	}
	defer f.Close()

	return genFPS(f, path)
}

func genFPS(f *os.File, path string) error {
	_, err := fmt.Fprintf(f, `
<?xml version="1.0" encoding="UTF-8"?>   
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

	viper.SetConfigType("toml")
	err = viper.ReadConfig(c)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(f, `
		<title><![CDATA[` + viper.GetString("problem.title") + `]]></title>
		<time_limit unit="s"><![CDATA[` + strconv.Itoa(viper.GetInt("limits.time")) + `]]></time_limit>
		<memory_limit unit="mb"><![CDATA[` + strconv.Itoa(viper.GetInt("limits.memory")) + `]]></memory_limit>
`)

	_, err = fmt.Fprintf(f, `
	</item>
</fps>
`)
	if err != nil {
		return err
	}
	return nil
}