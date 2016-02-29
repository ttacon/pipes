# pipes
--
    import "github.com/ttacon/pipes"

Package pipes is a utility package for drastically simplying piping input/output
between shell commands. For instance if we wanted to count the number of entries
in a directory:

    cmd := Command("ls", "-al")
    cmd.PipeTo("wc", "-l")
    output, err := cmd.Run()
    if err != nil {
      // deal with error...
    }
    fmt.Println("got: ", string(output))

## Usage

#### type Cmd

```go
type Cmd interface {
	// Run runs the connected commands and returns the output
	// of the final command if there are no errors.
	Run() ([]byte, error)
	// SetStdin sets the Stdin for the given Cmd. However,
	// this should only really be used ever on the first/original
	// Cmd structure.
	SetStdin(r io.Reader)
	// PipeTo is the magic of this package. Simply provide the command
	// and any given arguments, and PipeTo will make it so the current
	// command will pipe to the new one.
	PipeTo(name string, args ...string) Cmd
}
```

Cmd is a higher level wrapper around exec.Cmd that allows hooking together
multiple commands.

#### func  Command

```go
func Command(name string, args ...string) Cmd
```
Command creates a new command with the given name and arguments.
