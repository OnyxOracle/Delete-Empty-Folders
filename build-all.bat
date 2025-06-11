@echo off
rem This script builds the cleanup utility for a comprehensive list of platforms.
echo Starting the build process...

rem Create a clean 'builds' directory.
if exist builds (
    echo Cleaning up old build directory...
    rmdir /s /q builds
)
mkdir builds

rem Define all targets and call the build function for each one.
call :build darwin amd64
call :build darwin arm64
call :build dragonfly amd64
call :build freebsd 386
call :build freebsd amd64
call :build freebsd arm
call :build linux 386
call :build linux amd64
call :build linux arm
call :build linux arm64
call :build linux ppc64
call :build linux ppc64le
call :build linux mips64
call :build linux mips64le
call :build netbsd 386
call :build netbsd amd64
call :build netbsd arm
call :build openbsd 386
call :build openbsd amd64
call :build openbsd arm
call :build plan9 386
call :build plan9 amd64
call :build solaris amd64
call :build windows 386
call :build windows amd64
call :build windows arm64
call :build windows arm

echo.
echo All builds completed successfully!
echo Executables are in the 'builds' folder.
goto :eof

:build
set "GOOS=%1"
set "GOARCH=%2"
set "EXT="
if "%GOOS%"=="windows" set "EXT=.exe"

echo Building for %GOOS%-%GOARCH%...
go build -o builds/cleanup-%GOOS%-%GOARCH%%EXT% .

rem Check if the last command failed.
if errorlevel 1 (
    echo ERROR: Build failed for %GOOS%-%GOARCH%.
    pause
    exit /b 1
)
goto :eof

:eof
echo.
pause