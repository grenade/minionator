package main

import (
  "github.com/fatih/color"
  "gopkg.in/gcfg.v1"
  "fmt"
  "log"
  "strconv"
)

func main() {
  var cfg Config
  err := gcfg.ReadFileInto(&cfg, ".config")
  if err != nil {
    log.Fatal(err)
  }
  parent := GetBug(strconv.Itoa(cfg.Bugzilla.Parent), cfg)
  if parent.Id == cfg.Bugzilla.Parent {
    for _, alias := range cfg.Bugzilla.Child {
      child := GetBug(alias, cfg)
      if child.Alias == alias {
        fmt.Printf("bug %v found for %v\n", child.Id, child.Alias)
        if Contains(cfg.Bugzilla.Parent, child.DependsOn) {
          color.Blue(" - is linked to parent %v", cfg.Bugzilla.Parent)
        } else {
          color.Yellow(" - not linked to parent %v", cfg.Bugzilla.Parent)
        }
        if !child.IsOpen {
          if SetDepends(parent.Id, child.Id, fmt.Sprint("allocated to bug ", parent.Id), cfg) {
            color.Green(" - reopened %v (%v) and linked to parent %v", child.Id, child.Alias, parent.Id)
          }
        }
      } else {
        color.Yellow("no bug found for %v", alias)
        //todo: create the bug
      }
    }
  }
}

func Contains(needle int, haystack []int) bool {
  for _, straw := range haystack {
    if straw == needle {
      return true
    }
  }
  return false
}