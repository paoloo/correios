package main

import (
  "flag"
  "fmt"
)
func main(){
  flagCep    := flag.Bool("cep", false, "para verificar os dados de um cep")
  flagPacote := flag.Bool("pacote", false, "para verificar o status de um pacote")

  flag.Parse()

  if *flagCep && len(flag.Args())>0 {
    _,blob := BuscaCep(flag.Args()[0]) //"60040531")
    PrintCep(blob)
  } else if *flagPacote && len(flag.Args())>0{
    _,blob := BuscaEncomenda(flag.Args()[0]) //"JR804425413BR")
    PrintEncomenda(blob)
  } else {
    fmt.Println("        \\\\   ")
    fmt.Println("  // // // //")
    fmt.Println("     \\\\   Correios CLI 0.1")
    fmt.Println("          Uso ./correios flag VALOR")
    fmt.Println("          onde o VALOR e' obrigatorio e flag pode ser:")
    flag.PrintDefaults()
    fmt.Println("exemplos:")
    fmt.Println("./correios -cep 60040531")
    fmt.Println("./correios -pacote JR000000000BR")
  }

}
