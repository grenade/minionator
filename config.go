package main

type Config struct {
  Bugzilla struct {
    Host string
    Key string
    Parent int
    Child []string
  }
}