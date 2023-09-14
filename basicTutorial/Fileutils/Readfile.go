package fileutils

import "os"

func ReadTextFile(filename string) string {

              content,_ :=os.ReadFile(filename)//function used to read the file on given path

			return string(content)
}

