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
$ go install github.com/paoloo/correios@latest
```
NOTE: make sure that your $GOPATH/bin is on your **PATH**!

## Execution

- `correios`  or `correios -h` To show syntax tips
- `correios -cep 00000000` To receive informations about the given CEP.
- `correios -pacote JR000000000BR` To verify tracking data from given code.

## To Do
- Add more services from correios

## License
Mit License
