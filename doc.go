// Package pipes is a utility package for drastically simplying piping
// input/output between shell commands. For instance if we wanted to count
// the number of entries in a directory:
//
//   cmd := Command("ls", "-al")
//   cmd.PipeTo("wc", "-l")
//   output, err := cmd.Run()
//   if err != nil {
//     // deal with error...
//   }
//   fmt.Println("got: ", string(output))
package pipes
