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
    tab      string
    depth   int
}

func (iw *newWriter) Write(p []byte) (n int, err error) {
    tabs := ""
    for i := 0; i < iw.depth; i++ {
        tabs += iw.tab
    }
    
    p = append([]byte(tabs), p...)
    return iw.out.Write(p)
}

var cter int = 0

func dirTree(out io.Writer, path string, printFiles bool) error{
	writer := &newWriter{
		out:    os.Stdout,
		tab: "\t",
		depth: cter,
	  }
	
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

		
		if files[len(files)-1] == files[i] && i != 0 {
        fmt.Fprint(writer, "\t└───" + file.Name(), "\n")
        cter += 1
      }else{
        fmt.Fprint(writer, "├───" + file.Name(), "\n")
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
      
      var shit string
      for i := 0 ; i < cter ; i++ {
        shit += "\t│"
      }
      if files[len(files)-1] == files[i] && i != 0{
		cter += 1
		fmt.Println(shit, "\t└───", fileName, fileSize)
      }else if cter == 0{
		fmt.Println("├───", fileName, fileSize)
	  
	  }else {
        fmt.Println(shit, "\t├───", fileName, fileSize)
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
