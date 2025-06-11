GOOS=windows GOARCH=amd64 go build -o cleanup_win_amd64.exe builds/
GOOS=windows GOARCH=arm64 go build -o cleanup_win_arm64.exe builds/
GOOS=windows GOARCH=arm go build -o cleanup_win_arm32.exe builds/
echo "Building for Windows (AMD64, ARM64, ARM32) completed."

GOOS=darvin GOARCH=amd64 go build -o cleanup_macos_amd64 builds/
GOOS=darvin GOARCH=arm64 go build -o cleanup_macos_arm64 builds/
echo "Building for MacOS (AMD64(Intel), ARM64(Apple Silicon)) completed."

GOOS=linux GOARCH=amd64 go build -o cleanup_linux_amd64 builds/
GOOS=linux GOARCH=arm64 go build -o cleanup_linux_arm64 builds/
GOOS=linux GOARCH=arm go build -o cleanup_linux_arm32 builds/
echo "Building for Linux (AMD64, ARM64, ARM32) completed."