package main

import(
  "encoding/json"
  "net/http"
  "fmt"
  "io/ioutil"
)

type Article struct {
  Id   int    `json:"id"`
  Link  string `json:"link"`
  Title struct {
    Rendered string `json:"rendered"`
  }
}

func main(){
  resp, _ := http.Get("")
  defer resp.Body.Close()
  byteArray, _ := ioutil.ReadAll(resp.Body)
  jsonBytes := ([]byte)(byteArray)
  data := new(Article)
  if err := json.Unmarshal(jsonBytes, data); err != nil {
    fmt.Println("JSON Unmarshal error:", err)
    return
  }
  fmt.Println(data.Title.Rendered)
}

