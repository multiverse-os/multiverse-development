//
// Copyright (c) 2019 Stefaan Coussement
// MIT License
//
// more info: https://github.com/stefaanc/golang-exec
//
package local

import (
    "context"
    "errors"
    "fmt"
    "io"
    "os/exec"
    "strings"
    "syscall"

    "github.com/stefaanc/golang-exec/script"
)

//------------------------------------------------------------------------------

type Connection struct {
    Type string   // must be "local"
}

type Error struct {
    script   *script.Script
    command  string
    exitCode int
    err      error
}

type Runner struct {
    script  *script.Script
    command string
    cmd     *exec.Cmd
    cancel  context.CancelFunc

    exitCode int
}

//------------------------------------------------------------------------------

func (e *Error) Script()   *script.Script { return e.script }
func (e *Error) Command()  string         { return e.command }
func (e *Error) ExitCode() int            { return e.exitCode }
func (e *Error) Error()    string         { return e.err.Error() }
func (e *Error) Unwrap()   error          { return e.err }

//------------------------------------------------------------------------------

func New(connection interface{}, s *script.Script, arguments interface{}) (*Runner, error) {
    if s.Error != nil {
        return nil, &Error{
            script: s,
            exitCode: -1,
            err: fmt.Errorf("[golang-exec/runner/local/New()] script failed to parse: %#w\n", s.Error),
        }
    }

    r := new(Runner)
    r.script = s
    r.command = s.Command()

    stdin, err := s.NewReader(arguments)
    if err != nil {
        return nil, &Error{
            script: s,
            exitCode: -1,
            err: fmt.Errorf("[golang-exec/runner/local/New()] cannot create stdin reader: %#w\n", err),
        }
    }

    // create command, ready to start
    ctx, cancel := context.WithCancel(context.Background())
    args := strings.Split(r.command, " ")
    var cmd *exec.Cmd
    if args[0] == "cmd" {
        // cmd has argument-escaping rules that are different from other programs, so needs different treatment
        cmd = exec.CommandContext(ctx, args[0])
        cmd.SysProcAttr = &syscall.SysProcAttr{
            CmdLine: " " + strings.Join(args[1:], " "),
        }
    } else {
        cmd = exec.CommandContext(ctx, args[0], args[1:]...)
    }
    r.cmd = cmd
    r.cmd.Stdin  = stdin
    r.cancel = cancel

    return r, nil
}

//------------------------------------------------------------------------------

func (r *Runner) SetStdoutWriter(stdout io.Writer) {
    r.cmd.Stdout = stdout
}

func (r *Runner) SetStderrWriter(stderr io.Writer) {
    r.cmd.Stderr = stderr
}

func (r *Runner) StdoutPipe() (io.Reader, error) {
    reader, err := r.cmd.StdoutPipe()
    if err != nil {
        r.exitCode = -1
        return nil, &Error{
            script: r.script,
            exitCode: r.exitCode,
            err: fmt.Errorf("[golang-exec/runner/local/StdoutPipe()] cannot create stdout reader: %#w\n", err),
        }
    }

    return reader, nil
}

func (r *Runner) StderrPipe() (io.Reader, error) {
    reader, err := r.cmd.StderrPipe()
    if err != nil {
        r.exitCode = -1
        return nil, &Error{
            script: r.script,
            exitCode: r.exitCode,
            err: fmt.Errorf("[golang-exec/runner/local/StderrPipe()] cannot create stderr reader: %#w\n", err),
        }
    }

    return reader, nil
}

func (r *Runner) Run() error {
    err := r.cmd.Run()
    if err != nil {
        var exitErr *exec.ExitError
        if errors.As(err, &exitErr) {
            r.exitCode = exitErr.ProcessState.ExitCode()
            return &Error{
                script: r.script,
                command: r.command,
                exitCode: r.exitCode,
                err: fmt.Errorf("[golang-exec/runner/local/Run()] runner failed: %#w\n", err),
            }
        } else {
            r.exitCode = -1
            return &Error{
                script: r.script,
                command: r.command,
                exitCode: r.exitCode,
                err: fmt.Errorf("[golang-exec/runner/local/Run()] cannot execute runner: %#w\n", err),
            }
        }
    }

    return nil
}

func (r *Runner) Start() error {
    err := r.cmd.Start()
    if err != nil {
        r.exitCode = -1
        return &Error{
            script: r.script,
            command: r.command,
            exitCode: r.exitCode,
            err: fmt.Errorf("[golang-exec/runner/local/Start()] cannot start runner: %#w\n", err),
        }
    }

    return nil
}

func (r *Runner) Wait() error {
    err := r.cmd.Wait()
    if err != nil {
        var exitErr  *exec.ExitError
        if errors.As(err, &exitErr) {
            r.exitCode = exitErr.ProcessState.ExitCode()
        } else {
            r.exitCode = -1
        }
        return &Error{
            script: r.script,
            command: r.command,
            exitCode: r.exitCode,
            err: fmt.Errorf("[golang-exec/runner/local/Wait()] runner failed: %#w\n", err),
        }
    }

    r.exitCode = 0
    return nil
}

func (r *Runner) Close() error {
    if r.cancel != nil {
        r.cancel()
    }

    return nil
}

func (r *Runner) ExitCode() int {
    return r.exitCode
}

//------------------------------------------------------------------------------
