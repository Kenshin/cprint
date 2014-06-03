::===========================================================
:: CPrint: ANSI coloured text to the standard output on Windows
:: HOST  : https://github.com/kenshin/cprint
:: Author: Kenshin<kenshin@ksria.com>
::===========================================================

@ECHO off

IF "%1" == "doc" GOTO doc
IF "%1" == "test" GOTO test

:doc
@ECHO godoc -http=:6060 -server=:6060
godoc -http=:6060 -server=:6060
IF "%1" == "doc" GOTO exit

:test
@ECHO go test
go test
GOTO exit

:exit
@ECHO complete.