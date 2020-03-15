package template

import (
	"fmt"
	"goj/file"
	"log"
	"os"
)

func Create(path string) error {
	template := NewTemplate(path)
	log.Println("Creating from template...")
	for _, index := range Order {
		if index[0] == '_' {
			/* create folder */
			log.Println(" - folder:", template[index].Path)
			err := os.Mkdir(template[index].Path, os.ModeDir|0755)
			if err != nil {
				return err
			}
		} else {
			/* create file */
			name := template[index].Path + template[index].Name + template[index].Ext
			log.Println(" - file:  ", name)

			f, err := file.OpenAndTruncate(name, os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return err
			}

			/* try to write the content to the new file */
			_, err = fmt.Fprintf(f, template[index].Content)
			if err != nil {
				return err
			}

			err = f.Close()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
