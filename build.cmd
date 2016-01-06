@echo off
set GOPATH=%~dp0..\..\..\..\

if "%1" == "arm" goto armb
goto build

:armb
set GOOS=linux
set GOARCH=arm
set GOARM=6

:build
go get
go install

if "%1" == "start" goto run
if "%2" == "start" goto run
goto:eof

:run
start ./../../../../bin/hello-rest.exe