# go-nsi-example

This project is an example project to show how to call the NSI API from golang. 

To build this project:
```
go build

```
To run this project:
```
./go-nsi-example
```
expected output:
```
https://nsi.sec.usace.army.mil/nsiapi/structures?fips=15005&fmt=fs
GetByFips(15005) yeilded 58 structures; expected 58
GetByFips(15005) yeilded population of 115 across all structures; expected 115
GetByFips(15005) yeilded total value of 44632201.845300 across all structures; expected 44632201
```