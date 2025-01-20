package files

import (
	"fmt"
	"os"
)

func ReadFile()  {
    //os.ReadFile(name string) ([]byte, error)
    date, err :=  os.ReadFile("file.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(date))
}

func WritaFile(content []byte, name string)  {
   file, err :=  os.Create(name)
   if err != nil {
    fmt.Println(err)
   }
   defer file.Close()
   _, err = file.Write(content)
   if err != nil {
    fmt.Println(err)
    return
   }
   fmt.Println("Запись успешна!")
   
}