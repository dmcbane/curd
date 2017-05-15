@echo off
curd %* > %TEMP%\vv.tmp
set /p VV=<%TEMP%\vv.tmp
cd /D "%VV%"

