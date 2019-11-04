package template

import (
	"fmt"
	"os"
)

func Create(path string) error {
	template := NewTemplate(path)
	for index := range template {
		var name string
		if template[index].Ext[0] != '.' {
			name = template[index].Path + template[index].Name + ".cpp"
		} else {
			name = template[index].Path + template[index].Name + template[index].Ext
		}

		f, err := os.Create(name)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintf(f, template[index].Content)
		if err != nil {
			return err
		}

		err = f.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
