package main

import(
  "fmt"
  "os/exec"
  "os"
)

type workRequest struct {
  code string
}

func (w workRequest) String() string {
  return fmt.Sprintf("%s", w.code)
}

func (w workRequest) execute() {

  var (
    cmdOut []byte
    err    error
  )

  cmdOut, err = exec.Command("node", "-e", w.code).Output()

  if err != nil {
    fmt.Fprintln(os.Stderr, "There was an error running a bash command: ", err)
    os.Exit(1)
  }
  out := string(cmdOut)
  fmt.Println(out)

}
