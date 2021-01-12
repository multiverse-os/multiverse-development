package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "github.com/stefaanc/golang-exec/runner"
    "github.com/stefaanc/golang-exec/script"
)

func main() {
    // define connection to the server
    c := map[string]string{
        "Type": "ssh",
        "Host": "localhost",
        "Port": "22",
        "User": "me",
        "Password": "my-password",
        "Insecure": "true",
    }

    // create script runner
    home := "C:\\Users\\" + c["User"]   // for "cmd" and "powershell"
//    home := "/home/" + c["User"]   // for "bash"

    r, err := runner.New(c, lsScript, lsArguments{
//        Path: home + "\\doesn't exist",
        Path: home,
    })
    if err != nil {
        log.Fatal(err)
    }
    defer r.Close()

    // get a stdout-reader
    stdout, err := r.StdoutPipe()
    if err != nil {
        log.Fatal(err)
    }

    // get a stderr-reader
    stderr, err := r.StderrPipe()
    if err != nil {
        log.Fatal(err)
    }

    // start script runner
    err = r.Start()
    if err != nil {
        log.Fatal(err)
    }

    // wait for stdout-reader to complete
    result, err := ioutil.ReadAll(stdout)
    if err != nil {
        log.Fatal(err)
    }

    // wait for stderr-reader to complete
    errors, err := ioutil.ReadAll(stderr)
    if err != nil {
        log.Fatal(err)
    }

    // wait for script runner to complete
    err = r.Wait()
    if err != nil {
        fmt.Printf("exitcode: %d\n", r.ExitCode())
        fmt.Printf("stdout: \n%s", string(result))
        fmt.Printf("stderr: \n%s\n", string(errors))
        log.Fatal(err)
    }

    // write the result
    fmt.Printf("exitcode: %d\n", r.ExitCode())
    fmt.Printf("result: \n%s", string(result))
}

type lsArguments struct{
    Path string
}

var lsScript = script.New("ls", "cmd", `
    @echo off
    set "dirpath={{.Path}}"
    dir %dirpath%
`)

// var lsScript = script.New("ls", "powershell", `
//     $ErrorActionPreference = 'Stop'
//
//     $dirpath = "{{.Path}}"
//     Get-ChildItem -Path $dirpath | Format-Table
//
//     exit 0
// `)

// var lsScript = script.New("ls", "bash", `
//     set -e -o pipefail
//
//     dirpath="{{.Path}}"
//     ls -la "$dirpath"
//
//     exit 0
// `)
