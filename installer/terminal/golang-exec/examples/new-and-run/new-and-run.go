package main

import (
    "bytes"
    "fmt"
    "log"
    "os"
    "github.com/stefaanc/golang-exec/runner"
    "github.com/stefaanc/golang-exec/script"
)

type myConnection struct {
    Type     string
    Host     string
    Port     uint16
    User     string
    Password string
    Insecure bool
}

func main() {
    // define connection to the server
    c := myConnection{
        Type: "ssh",
        Host: "localhost",
        Port: 22,
        User: "me",
        Password: "my-password",
        Insecure: true,
    }

    // create script runner
    wd, _ := os.Getwd()
    r, err := runner.New(c, lsScript, lsArguments{
//        Path: wd + "\\doesn't exist",
        Path: wd,
    })
    if err != nil {
        log.Fatal(err)
    }
    defer r.Close()

    // create buffer to capture stdout, set a stdout-writer
    var stdout bytes.Buffer
    r.SetStdoutWriter(&stdout)

    // create buffer to capture stderr, set a stderr-writer
    var stderr bytes.Buffer
    r.SetStderrWriter(&stderr)

    // run script runner
    err = r.Run()
    if err != nil {
        fmt.Printf("exitcode: %d\n", r.ExitCode())
        fmt.Printf("stdout: \n%s\n", stdout.String())
        fmt.Printf("stderr: \n%s\n", stderr.String())
        log.Fatal(err)
    }

    // write the result
    fmt.Printf("exitcode: %d\n", r.ExitCode())
    fmt.Printf("result: \n%s", stdout.String())
}

type lsArguments struct{
    Path string
}

var lsScript = script.New("ls", "powershell", `
    $ErrorActionPreference = 'Stop'

    $dirpath = "{{.Path}}"
    Get-ChildItem -Path $dirpath | Format-Table

    exit 0
`)
