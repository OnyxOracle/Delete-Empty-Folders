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
    @{OS="darwin";    Arch="amd64"},
    @{OS="darwin";    Arch="arm64"},
    @{OS="dragonfly"; Arch="amd64"},
    @{OS="freebsd";   Arch="386"},
    @{OS="freebsd";   Arch="amd64"},
    @{OS="freebsd";   Arch="arm"},
    @{OS="linux";     Arch="386"},
    @{OS="linux";     Arch="amd64"},
    @{OS="linux";     Arch="arm"},
    @{OS="linux";     Arch="arm64"},
    @{OS="linux";     Arch="ppc64"},
    @{OS="linux";     Arch="ppc64le"},
    @{OS="linux";     Arch="mips64"},
    @{OS="linux";     Arch="mips64le"},
    @{OS="netbsd";    Arch="386"},
    @{OS="netbsd";    Arch="amd64"},
    @{OS="netbsd";    Arch="arm"},
    @{OS="openbsd";   Arch="386"},
    @{OS="openbsd";   Arch="amd64"},
    @{OS="openbsd";   Arch="arm"},
    @{OS="plan9";     Arch="386"},
    @{OS="plan9";     Arch="amd64"},
    @{OS="solaris";   Arch="amd64"},
    @{OS="windows";   Arch="386"},
    @{OS="windows";   Arch="amd64"},
    @{OS="windows";   Arch="arm"},
    @{OS="windows";   Arch="arm64"}
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
    
    # Run the build command.
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

# Optional: Create a ZIP archive of all builds for easy distribution.
Write-Host "Creating ZIP archive of all builds..." -ForegroundColor Cyan
try {
    Compress-Archive -Path "$buildDir\*" -DestinationPath "cleanup-builds.zip" -Force
    Write-Host "Successfully created cleanup-builds.zip" -ForegroundColor Green
} catch {
    Write-Host "Warning: Could not create ZIP file. This feature requires PowerShell 5.1 or newer." -ForegroundColor Yellow
}

Write-Host ""
Write-Host "Process finished. Executables are in the 'builds' folder."
Read-Host -Prompt "Press Enter to exit"