#!/bin/bash
# This script builds the cleanup utility for a comprehensive list of platforms.
echo -e "\033[36mStarting the build process...\033[0m"

# Create a clean 'builds' directory.
if [ -d "builds" ]; then
    echo "Cleaning up old build directory..."
    rm -r builds
fi
mkdir builds

# Define the list of target platforms in an array.
# Format: "OS/ARCH"
targets=(
    "aix/ppc64"
    "darwin/amd64"
    "darwin/arm64"
    "dragonfly/amd64"
    "freebsd/386"
    "freebsd/amd64"
    "freebsd/arm"
    "freebsd/arm64"
    "freebsd/riscv64"
    "illumos/amd64"
    "linux/386"
    "linux/amd64"
    "linux/arm"
    "linux/arm64"
    "linux/loong64"
    "linux/mips"
    "linux/mips64"
    "linux/mips64le"
    "linux/mipsle"
    "linux/ppc64"
    "linux/ppc64le"
    "linux/riscv64"
    "linux/s390x"
    "netbsd/386"
    "netbsd/amd64"
    "netbsd/arm"
    "netbsd/arm64"
    "openbsd/386"
    "openbsd/amd64"
    "openbsd/arm"
    "openbsd/arm64"
    "openbsd/ppc64"
    "openbsd/riscv64"
#   "plan9/386"        afero does not support Plan9
#   "plan9/amd64"
#   "plan9/arm"
    "solaris/amd64"
    "windows/386"
    "windows/amd64"
    "windows/arm64"
    "windows/arm"
)

total=${#targets[@]}
current=1

# Loop through each target and build the executable.
for target in "${targets[@]}"; do
    # Split the target string into OS and ARCH.
    GOOS=$(echo $target | cut -d'/' -f1)
    GOARCH=$(echo $target | cut -d'/' -f2)
    
    echo -e "\033[33m($current/$total) Building for $GOOS-$GOARCH...\033[0m"
    
    # Set the output name and add .exe for Windows.
    output_name="builds/cleanup-$GOOS-$GOARCH"
    if [ $GOOS = "windows" ]; then
        output_name+=".exe"
    fi
    
    # Run the simplified build command.
    env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name .
    
    # Check if the build failed.
    if [ $? -ne 0 ]; then
        echo -e "\033[31mERROR: Build failed for $GOOS-$GOARCH.\033[0m"
        exit 1
    fi
    
    current=$((current + 1))
done

echo ""
echo -e "\033[32mAll builds completed successfully!\033[0m"
echo "Executables are in the 'builds' folder."