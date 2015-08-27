package main

type Config struct {
  Minionator struct {
    Task []string
  }
  Bugzilla struct {
    Host string
    Key string
    Parent int
    Child []string
  }
}