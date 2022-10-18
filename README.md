# Correios CLI
```
        \\
  // // // //
     \\   Correios CLI 0.3
```
A pure go implementation of some services for brazilian state post office: Correios.

## Build

Too simple!
```
$ go get github.com/paoloo/correios
$ cd $GOPATH/src/github.com/paoloo/correios
$ go build
```
## Execution

- `./correios`  or `./correios -h` To show syntax tips
- `./correios -cep 00000000` To receive informations about the given CEP.
- `./correios -pacote JR000000000BR` To verify tracking data from given code.

## To Do
- Fix the output of json and xml. I accept PR ;D
- Add more services from correios

## License
Mit License
