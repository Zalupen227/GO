package main

import (
	"fmt"
 	"io"
 	"os"
 	// "path/filepath"
 	// "strings"
	"strconv"
)

type newWriter struct {
    out      io.Writer
    tab   string
}

func (iw *newWriter) Write(p []byte) (n int, err error) {
    // Добавляем отступ перед выводом
    p = append([]byte(iw.tab), p...)
    return iw.out.Write(p)
}


func dirTree(out io.Writer, path string, printFiles bool) error{
	dir, err := os.Open(path)

	if err != nil {
		fmt.Errorf(err.Error())
	}
	defer dir.Close()
	files, err := dir.ReadDir(-1)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	

	for i, file := range files{
		if file.IsDir()  {
			fmt.Fprint(out, "├───" + file.Name(), "\n")
			writer := &newWriter{
				out:    os.Stdout,
				tab: "\t",
			}
			
			dirTree(writer, path + "\\" + file.Name(), printFiles)
		}else if printFiles == true {
			fileInfo, err := file.Info()
			if err != nil{
			fmt.Println("Ошибка при получении информации о файле:", err)
				continue
			}
			
			fileName := fileInfo.Name()
			fileSize := "(" + strconv.FormatInt(fileInfo.Size(), 10) + "b)"
			if files[0] != files[i]{
				fmt.Println("├───", fileName, fileSize)
			} else {
				fmt.Println("└───", fileName, fileSize)
			}
		}
	}

	

	return nil
}



func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
