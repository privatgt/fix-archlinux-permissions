package main

import (
   "os/exec"
   "fmt"
   "strings"
   "regexp"
   "os"
)

func main(){
    fmt.Println("Program started\nLaunching paccheck")
    cmd := exec.Command("/bin/bash", "-c",`paccheck --file-properties --quiet`)
    output, _ := cmd.CombinedOutput()
    r, _ := regexp.Compile(`(?im)(^.*)permission mismatch(.*$)`)
    out:=r.FindAllString(string(output),-1)
    path, _ := regexp.Compile(`(?im)(\d?)[0-9][0-9][0-9]`)
    for _, s := range out {
        cmd = exec.Command("chmod",path.FindString(strings.Split(s,"'")[2]),strings.Split(s,"'")[1])
        output, err := cmd.CombinedOutput()
        fmt.Println("chmod",path.FindString(strings.Split(s,"'")[2]),strings.Split(s,"'")[1])
        if err!=nil{
            fmt.Println("Command error")
            fmt.Println(string(output))
            os.Exit(1)
        }
    }
}
