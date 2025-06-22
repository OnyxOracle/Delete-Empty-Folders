# This script builds the cleanup utility for a wide range of platforms.
Write-Host "Starting the build process..." -ForegroundColor Cyan

# Create a clean 'builds' directory to store the executables.
$buildDir = ".\builds"
if (Test-Path $buildDir) {
    Write-Host "Cleaning up old build directory..."
    Remove-Item -Recurse -Force $buildDir
}
New-Item -ItemType Directory -Path $buildDir | Out-Null

# Define the list of build targets as an array of objects.
$targets = @(
    @{OS="aix";       Arch="ppc64"},
    @{OS="darwin";    Arch="amd64"},
    @{OS="darwin";    Arch="arm64"},
    @{OS="dragonfly"; Arch="amd64"},
    @{OS="freebsd";   Arch="386"},
    @{OS="freebsd";   Arch="amd64"},
    @{OS="freebsd";   Arch="arm"},
    @{OS="freebsd";   Arch="arm64"},
    @{OS="freebsd";   Arch="riscv64"},
    @{OS="illumos";   Arch="amd64"},
    @{OS="linux";     Arch="386"},
    @{OS="linux";     Arch="amd64"},
    @{OS="linux";     Arch="arm"},
    @{OS="linux";     Arch="arm64"},
    @{OS="linux";     Arch="loong64"},
    @{OS="linux";     Arch="mips"},
    @{OS="linux";     Arch="mips64"},
    @{OS="linux";     Arch="mips64le"},
    @{OS="linux";     Arch="mipsle"},
    @{OS="linux";     Arch="ppc64"},
    @{OS="linux";     Arch="ppc64le"},
    @{OS="linux";     Arch="riscv64"},
    @{OS="linux";     Arch="s390x"},
    @{OS="netbsd";    Arch="386"},
    @{OS="netbsd";    Arch="amd64"},
    @{OS="netbsd";    Arch="arm"},
    @{OS="netbsd";    Arch="arm64"},
    @{OS="openbsd";   Arch="386"},
    @{OS="openbsd";   Arch="amd64"},
    @{OS="openbsd";   Arch="arm"},
    @{OS="openbsd";   Arch="arm64"},
    @{OS="openbsd";   Arch="ppc64"},
    @{OS="openbsd";   Arch="riscv64"},
#   @{OS="plan9";     Arch="386"},    afero does not support Plan9
#   @{OS="plan9";     Arch="amd64"},
#   @{OS="plan9";     Arch="arm"},
    @{OS="solaris";   Arch="amd64"},
    @{OS="windows";   Arch="386"},
    @{OS="windows";   Arch="amd64"},
    @{OS="windows";   Arch="arm64"},
    @{OS="windows";   Arch="arm"}
)

$total = $targets.Count
$current = 1

# Loop through each target and build the executable.
foreach ($target in $targets) {
    $os = $target.OS
    $arch = $target.Arch
    
    Write-Host "($current/$total) Building for $os-$arch..." -ForegroundColor Yellow
    
    $env:GOOS = $os
    $env:GOARCH = $arch
    
    $outputName = "cleanup-$os-$arch"
    if ($os -eq "windows") {
        $outputName += ".exe"
    }
    
    $outputPath = Join-Path $buildDir $outputName
    
    # Run the simplified build command.
    go build -o $outputPath .
    
    # Check for build errors.
    if ($LASTEXITCODE -ne 0) {
        Write-Host "ERROR: Build failed for $os-$arch." -ForegroundColor Red
        Read-Host -Prompt "Build failed. Press Enter to exit"
        exit 1
    }
    
    $current++
}

Write-Host ""
Write-Host "All individual builds completed!" -ForegroundColor Green
Read-Host -Prompt "Press Enter to exit"