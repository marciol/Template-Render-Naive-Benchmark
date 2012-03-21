package main

import (
    "os"
    "strconv"
    "html/template"
    "fmt"
    "runtime"
    "bytes"
//    "io/ioutil"
)

func init() {
    runtime.GOMAXPROCS(runtime.NumCPU())
}

type Person struct {
    Name string
    LastName string
}

func main() {

    filename := "tmpl/person_list.html"

    t := template.Must(template.ParseFiles(filename))

    persons := make([]Person, 100)

    for i := 0; i < 100; i++ {
        iStr := strconv.Itoa(i)
        persons[i] = Person{"Name " + iStr, "Last " + iStr}
    }

    ch_sig := make(chan *bytes.Buffer, 10)

    for i := 0; i < 10; i++ {
        go func(i int, ch chan *bytes.Buffer) {
            buf := new(bytes.Buffer)
            err := t.Execute(buf, persons)
            if err != nil {
                panic(err)
            }
            ch <- buf 
        }(i, ch_sig)
    }

    for i := 0; i < 10; i++ {
        fmt.Println("Executing ", i)
        buf := <-ch_sig
        buf.WriteTo(os.Stdout)
    }
}
