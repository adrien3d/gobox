@echo off 
for /r "%cd%\" %%i in (*.bat) do set VrTest=%%~nxi
cls
go run %VrTest:~0,-4%.go
pause