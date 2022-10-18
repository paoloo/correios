package main

import (
  "bytes"
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/xml"
  "encoding/json"
)

type createEnvelope struct {
	Body createBody `xml:"Body"`
}

type createBody struct {
	CepResponse create `xml:"consultaCEPResponse"`
}

type create struct {
	Return workItem `xml:"return"`
}

type workItem struct {
  Bairro string `xml:"bairro" json:"bairro"`
  Cep string `xml:"cep" json:"cep"`
  Cidade string `xml:"cidade" json:"cidade"`
  Complemento string `xml:"complemento" json:"complemento"`
  Complemento2 string `xml:"complemento2" json:"complemento2"`
  End string `xml:"end" json:"end"`
  Id int `xml:"id" json:"id"`
  Uf string `xml:"uf" json:"UF"`
}

func toUtf8(iso8859_1_buf []byte) []byte {
    buf := make([]rune, len(iso8859_1_buf))
    for i, b := range iso8859_1_buf {
        buf[i] = rune(b)
    }
    return []byte(string(buf))
}

func PrintCep(_blob []byte) {
  b := createEnvelope{}
  xml.Unmarshal(toUtf8(_blob), &b)
  data, _ := json.MarshalIndent(b.Body.CepResponse.Return, "", "\t")
  fmt.Println(string(data))
}

func BuscaCep(_cep string) (string, []byte){
  apiUrl := "https://apps.correios.com.br/SigepMasterJPA/AtendeClienteService/AtendeCliente?wsdl"
  data := `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cli="http://cliente.bean.master.sigep.bsb.correios.com.br/"><soapenv:Header/><soapenv:Body><cli:consultaCEP><cep>`+_cep+`s</cep></cli:consultaCEP></soapenv:Body></soapenv:Envelope>`
  client := &http.Client{}
  r, err := http.NewRequest("POST", apiUrl, bytes.NewBufferString(data)); if err != nil {
    return "-1 ERROR",[]byte("cannot craft request!")
  }
  r.Header.Add("Content-Type", "application/soap+xml; charset=iso-8859-1")
  r.Header.Add("User-Agent","Dalvik/1.6.0 (Linux; U; Android 4.2.1; LG-P875h Build/JZO34L)")
  resp, err := client.Do(r); if err != nil {
    return "-2 ERROR",[]byte("cannot send request!")
  }
  _body, err := ioutil.ReadAll(resp.Body); if err != nil {
    return "-3 ERROR", []byte("cannot decode body!")
  }
  return resp.Status, _body
}
