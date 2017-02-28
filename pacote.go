package main

import (
  "bytes"
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

func BuscaEncomenda(_code string) (string, string){
  apiUrl := "http://webservice.correios.com.br/service/rest/rastro/rastroMobile"
  data := "<rastroObjeto><usuario>MobileXect</usuario><senha>DRW0#9F$@0</senha><tipo>L</tipo><resultado>U</resultado><objetos>"+ _code  +"</objetos><lingua>101</lingua><token>QTXFMvu_Z-6XYezP3VbDsKBgSeljSqIysM9x</token></rastroObjeto>"
  client := &http.Client{}
  r, err := http.NewRequest("POST", apiUrl, bytes.NewBufferString(data)); if err != nil {
    return "-1 ERROR","cannot craft request!"
  }
  r.Header.Add("Accept", "application/json")
  r.Header.Add("Content-Type", "application/xml")
  r.Header.Add("User-Agent","Dalvik/1.6.0 (Linux; U; Android 4.2.1; LG-P875h Build/JZO34L)")
  resp, err := client.Do(r); if err != nil {
    return "-2 ERROR","cannot send request!"
  }
  _body, err := ioutil.ReadAll(resp.Body); if err != nil {
    return "-3 ERROR", "cannot decode body!"
  }
  return resp.Status, string(_body)
}

func PrintEncomenda(_jsonblob string){
  buf := new(bytes.Buffer)
  json.Indent(buf, []byte(_jsonblob), "", "  ")
  fmt.Println(buf)
}

