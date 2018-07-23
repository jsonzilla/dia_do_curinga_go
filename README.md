# DiaDoCuringaGo

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/3885765f996243a5be0363757aa2d9f9)](https://www.codacy.com/app/0um/DiaDoCuringaGo?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=0um/DiaDoCuringaGo&amp;utm_campaign=Badge_Grade)
[![codebeat badge](https://codebeat.co/badges/754a5a89-52f3-46da-9bf5-12cec2c05f7f)](https://codebeat.co/projects/github-com-0um-diadocuringago-master)
[![Build Status](https://travis-ci.org/0um/DiaDoCuringaGo.png)](https://travis-ci.org/0um/DiaDoCuringaGo)

A Gregorian calendar converter to an annual calendar with 364 days, each representing a playing card. Divided into 52 weeks and 13 months, all with 28 days. Day 365 is an extra day, the day of the wildcard. Every four years to two extra days the double day of the wildcard.

The count of years begins in 1790

Taking into account the southern hemisphere! The seasons of the year are marked by different suits thus providing the suits of the months.

* Autumn is club 
* Summer and is golds
* Spring is cup 
* Winter is swords

## Compile
```
go build -o build/ddc ddc.go
```

## Run
```
> ./build/ddc
```

## Run test
```
go test
```

## Check coverage
```
go test -coverprofile=c.out && go tool cover -html=c.out
```
