package pipes

import (
	"bytes"
	"io"
	"io/ioutil"
	"os/exec"
)

func Command(name string, args ...string) Cmd {
	return &cmd{
		Cmd: exec.Command(name, args...),
	}
}

type Cmd interface {
	Run() ([]byte, error)
	SetStdin(r io.Reader)
	PipeTo(name string, args ...string) Cmd
}

type cmd struct {
	*exec.Cmd

	next     *cmd
	outBuf   *bytes.Buffer
	whenDone func()
}

func (c *cmd) SetStdin(r io.Reader) {
	c.Cmd.Stdin = r
}

func (c *cmd) Run() ([]byte, error) {
	// ugly hack for now
	if err := c.start(); err != nil {
		return nil, err
	}
	return c.finish()
}

func (c *cmd) start() error {
	if err := c.Cmd.Start(); err != nil {
		return err
	}

	if c.next != nil {
		return c.next.start()
	}
	return nil
}

func (c *cmd) finish() ([]byte, error) {
	if err := c.Cmd.Wait(); err != nil {
		return nil, err
	}

	if c.next != nil {
		c.whenDone()
		return c.next.finish()
	}
	return ioutil.ReadAll(c.outBuf)
}

func (c *cmd) PipeTo(name string, args ...string) Cmd {
	c.next = &cmd{
		Cmd:    exec.Command(name, args...),
		outBuf: bytes.NewBuffer(nil),
	}

	r, w := io.Pipe()
	c.Cmd.Stdout = w
	c.whenDone = func() {
		w.Close()
	}
	c.next.Cmd.Stdin = r
	c.next.Cmd.Stdout = c.next.outBuf
	return c.next
}
