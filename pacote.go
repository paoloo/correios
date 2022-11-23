package main

import (
  "bytes"
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  b64 "encoding/base64"
)

type tokenType struct {
  Token string `json:"token"`
}

func BuscaEncomenda(_code string) (string, string){
  // =============================================== phase#0 data
  validationUrl := "https://proxyapp.correios.com.br/v1/app-validation"
  data := []byte(`{"requestToken":"`+ b64.StdEncoding.EncodeToString([]byte("android;br.com.correios.preatendimento;F32E29976709359859E0B97F6F8A483B9B953578")) +`"}`)
  apiURL := "https://proxyapp.correios.com.br/v1/sro-rastro/"+_code
  // =============================================== phase#1 validation
  client := &http.Client{}
  r, err := http.NewRequest("POST", validationUrl, bytes.NewBuffer(data)); if err != nil {
    return "-1 ERROR","cannot craft request!"
  }
  r.Header.Add("Accept", "application/json")
  r.Header.Add("Content-Type", "application/json")
  r.Header.Add("User-Agent","Dart/2.18 (dart:io)")
  resp, err := client.Do(r); if err != nil {
    return "-2 ERROR","cannot send request!"
  }
  _body, err := ioutil.ReadAll(resp.Body); if err != nil {
    return "-3 ERROR", "cannot decode body!"
  }
  res := tokenType{}
  json.Unmarshal(_body, &res)
  // =============================================== phase#2 search
  client2 := &http.Client{}
  r2, err := http.NewRequest("GET", apiURL, nil); if err != nil {
    return "-4 ERROR","cannot craft request!"
  }
  r2.Header.Add("User-Agent","Dart/2.18 (dart:io)")
  r2.Header.Add("Accept", "application/json")
  r2.Header.Add("app-check-token", res.Token)
  resp2, err := client2.Do(r2); if err != nil {
    return "-5 ERROR", "cannot send request"
  }
  _body2, err := ioutil.ReadAll(resp2.Body); if err != nil {
    return "-6 ERROR", "cannot decode body!"
  }
  return resp2.Status, string(_body2)
}

func PrintEncomenda(_jsonblob string){
  buf := new(bytes.Buffer)
  json.Indent(buf, []byte(_jsonblob), "", "  ")
  fmt.Println(buf)
}
