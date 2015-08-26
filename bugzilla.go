package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "encoding/json"
)

func GetBug(id_or_alias string, cfg Config) Bug {
  bzUrl := fmt.Sprint("http://", cfg.Bugzilla.Host, "/rest/bug/", id_or_alias, "?api_key=", cfg.Bugzilla.Key) 
  fmt.Printf("%v\n", bzUrl)

  //var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
  req, err := http.NewRequest("GET", bzUrl, nil/* bytes.NewBuffer(jsonStr)*/)
  req.Header.Set("Accept", "application/json")
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
      panic(err)
  }
  defer resp.Body.Close()

  fmt.Println("response Status:", resp.Status)
  fmt.Println("response Headers:", resp.Header)
  if resp.StatusCode == 200 {
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      log.Fatal(err)
    }
    //fmt.Println("response Body:", string(body))
    var bugsApiResponse BugsApiResponse
    err = json.Unmarshal(body, &bugsApiResponse)
    if err != nil {
      log.Fatal(err)
    }
    return bugsApiResponse.Bugs[0]
  }
  return Bug{}
}