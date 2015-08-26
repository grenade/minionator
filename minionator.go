package main

import (
  "gopkg.in/gcfg.v1"
  "fmt"
  "log"
)

func main() {
  var cfg Config
  err := gcfg.ReadFileInto(&cfg, ".config")
  if err != nil {
    log.Fatal(err)
  }
  bug := GetBug("b-2008-ix-0034", cfg)
  fmt.Printf("%+v\n", bug)
}