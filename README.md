if you use image processing on the server side on window
you must install gcc tesseract-ocr leptonica with msYS2 MINGW64
if you window not have a msYS2 MINGW64 plase install it first!

and after success installed
your should to use powershell to set env
like this

```
$env:CC="C:/msys64/ucrt64/bin/gcc.exe"
$env:CGO_CFLAGS="-IC:/msys64/ucrt64/include"
$env:CGO_LDFLAGS="-LC:/msys64/ucrt64/lib -ltesseract -lleptonica -llept"
$env:TESSDATA_PREFIX="C:/msys64/ucrt64/share/tessdata/"
$env:CGO_ENABLED="1"
$env:PATH = "C:\msys64\ucrt64\bin;" + $env:PATH
```

and if your wan't to config everytime your run this program
you can set with

```
[System.Environment]::SetEnvironmentVariable("CC", "C:\msys64\ucrt64\bin\gcc.exe", "User")
[System.Environment]::SetEnvironmentVariable("CXX", "C:\msys64\ucrt64\bin\g++.exe", "User")
[System.Environment]::SetEnvironmentVariable("CGO_ENABLED", "1", "User")
[System.Environment]::SetEnvironmentVariable("CGO_CFLAGS", "-IC:\msys64\ucrt64\include", "User")
[System.Environment]::SetEnvironmentVariable("CGO_LDFLAGS", "-LC:\msys64\ucrt64\lib -ltesseract -lleptonica -llept", "User")
[System.Environment]::SetEnvironmentVariable("TESSDATA_PREFIX", "C:\msys64\ucrt64\share\tessdata\", "User")
```

in your powershell
you can check env with
go env VARIABLE_NAME
example
env CGO_CFLAGS
if it return -O2 -g it mean not set a path then error when run
if it return -IC:/msys64/ucrt64/include maybe can work

for load test
sudo docker compose -f ./docker-compose.test.yml up --build
