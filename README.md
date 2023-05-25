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

### asdf

If you installed go with asdf, make sure to reshim it (`asdf reshim golang`) after installing this package.
Refer to the asdf [plugin](https://github.com/kennyp/asdf-golang) for more information.

## Execution

- `correios`  or `correios -h` To show syntax tips
- `correios -cep 00000000` To receive informations about the given CEP.
- `correios -pacote JR000000000BR` To verify tracking data from given code.

### OR, if you decide to use docker:

- `docker pull paoloo/correios`
- `docker run --rm paoloo/correios:latest -cep 00000000` or
- `docker run --rm paoloo/correios:latest -pacote JR000000000BR`


## To Do
- Add more services from correios

## License
Mit License
