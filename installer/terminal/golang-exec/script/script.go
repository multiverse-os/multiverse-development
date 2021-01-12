//
// Copyright (c) 2019 Stefaan Coussement
// MIT License
//
// more info: https://github.com/stefaanc/golang-exec
//
package script

import (
    "bytes"
    "fmt"
    "io"
    "os"
    "math/rand"
    "strings"
    "text/template"
    "time"
)

//------------------------------------------------------------------------------

type Script struct {
    Name       string
    Shell      string   // "cmd", powershell", "bash", "sh", ...

    template   *template.Template

    Error      error    // error from New()
}

//------------------------------------------------------------------------------

func New(name string, shell string, code string) *Script {
    // remark that New() doesn't return any errors directly
    // instead, error are saved in the 'Error'-field of the returned script
    // this allows using New() in a package scope, while checking for errors in a function scope
    template, err := template.New(name).Parse(code)
    if err != nil {
        err = fmt.Errorf("[golang-exec/script/New()] cannot parse script: %#w\n", err)
    }

    s := new(Script)
    s.Name = name

    if err == nil {
        s.Shell = strings.ToLower(shell)
        s.template = template
    } else {
        s.Error = err
    }

    return s
}

func NewFromString(name string, shell string, code string) (*Script, error) {
    template, err := template.New(name).Parse(code)
    if err != nil {
        return nil, fmt.Errorf("[golang-exec/script/NewFromString()] cannot parse script: %#w\n", err)
    }

    s := new(Script)
    s.Name = name
    s.Shell = strings.ToLower(shell)
    s.template = template

    return s, nil
}

func NewFromFile(name string, shell string, file string) (*Script, error) {
    template, err := template.New(name).ParseFiles(file)
    if err != nil {
        return nil, fmt.Errorf("[golang-exec/script/NewFromFile()] cannot parse script: %#w\n", err)
    }

    s := new(Script)
    s.Name = name
    s.Shell = strings.ToLower(shell)
    s.template = template

    return s, nil
}

//------------------------------------------------------------------------------

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func (s *Script) Command() string {
    // returns the command(s) to execute a script that is read from stdin
    switch s.Shell {
    case "cmd":
        // for cmd, we cannot execute code directly from stdin
        // hence we save stdin (the rendered code) to a file and then execute that file
        //
        // the steps in the command are:
        // - run a cmd command, enable delayed expansion
        // - set the name of a temp file
        // - use "more" to save stdin to temp-file
        // - use "cmd" to execute temp-file
        // - save "%errorlevel%" because it will be overwritten by the next step
        // - delete temp-file
        // - exit with saved "%errorlevel%"
        wd, _ := os.Getwd()
        spath := fmt.Sprintf("%s\\_temp-%d.bat", wd, seededRand.Uint64())
        return fmt.Sprintf("cmd /E:ON /V:ON /C \"more > \"%s\" && cmd /C \"%s\" & set \"E=!errorlevel!\" & del /Q \"%s\" & exit !E!\"", spath, spath, spath)
    case "powershell":
        // for powershell, we can  execute code directly from stdin, returning "PowerShell -NoProfile -ExecutionPolicy ByPass -Command -"
        // however, it seems that fatal exceptions don't stop the script, and thus "$ErrorActionPreference = 'Stop'" also doesn't work properly
        // hence we save stdin (the rendered code) to a file using cmd and then execute that file using powershell
        //
        // the steps in the command are similar to the steps for the cmd shell
        wd, _ := os.Getwd()
        spath := fmt.Sprintf("%s\\_temp-%d.ps1", wd, seededRand.Uint64())
        return fmt.Sprintf("cmd /E:ON /V:ON /C \"more > \"%s\" && PowerShell -NoProfile -ExecutionPolicy ByPass -File \"%s\" & set \"E=!errorlevel!\" & del /Q \"%s\" & exit !E!\"", spath, spath, spath)
    default:
        // for bash,... we execute code directly from stdin
        return s.Shell + " -"
    }
}

func (s *Script) NewReader(arguments interface{}) (io.Reader, error) {
    // returns a reader for the parsed & rendered script
    var rendered bytes.Buffer
    if s.template != nil {
        err := s.template.Execute(&rendered, arguments)
        if err != nil {
            return nil, fmt.Errorf("[golang-exec/script/NewReader()] cannot render script: %#w\n", err)
        }
    }

    return &rendered, nil
}

//------------------------------------------------------------------------------
