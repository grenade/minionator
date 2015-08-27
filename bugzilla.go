package main

import (
  "github.com/fatih/color"
  "bytes"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "encoding/json"
)

func GetBug(id_or_alias string, cfg Config) Bug {
  bzUrl := fmt.Sprint("https://", cfg.Bugzilla.Host, "/rest/bug/", id_or_alias, "?api_key=", cfg.Bugzilla.Key) 
  req, err := http.NewRequest("GET", bzUrl, nil)
  req.Header.Set("Accept", "application/json")
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()
  if resp.StatusCode == 200 {
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      log.Fatal(err)
    }
    var bugsApiResponse BugsApiResponse
    err = json.Unmarshal(body, &bugsApiResponse)
    if err != nil {
      log.Fatal(err)
    }
    return bugsApiResponse.Bugs[0]
  }
  return Bug{}
}

func SetDepends(blocker int, depends int, comment string, cfg Config) bool {
  bzUrl := fmt.Sprint("https://", cfg.Bugzilla.Host, "/rest/bug/", depends, "?api_key=", cfg.Bugzilla.Key)
  message := ReOpenChildMessage {
    []int { depends },
    "REOPENED",
    DependsOnAppender { []int { blocker } },
    Comment { comment, false, false } }
  //color.Cyan("request Payload: %+v", message)
  payload, err := json.Marshal(message)
  if err != nil {
    log.Fatal(err)
  }
  req, err := http.NewRequest("PUT", bzUrl, bytes.NewBuffer(payload))
  if err != nil {
    log.Fatal(err)
  }
  req.Header.Set("Content-Type", "application/json")
  req.Header.Set("Accept", "application/json")
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal(err)
  }
  if resp.StatusCode == 200 {
    var bugsApiResponse BugsApiResponse
    err = json.Unmarshal(body, &bugsApiResponse)
    if err != nil {
      log.Fatal(err)
    }
    //color.Cyan("response Status: %v", resp.Status)
    //color.Cyan("response Headers: %v", resp.Header)
    //color.Cyan("response Body: %v", string(body))
    return true
  }
  color.Red("response Status: %v", resp.Status)
  color.Red("response Headers: %v", resp.Header)
  color.Red("response Body: %v", string(body))
  return false
}
