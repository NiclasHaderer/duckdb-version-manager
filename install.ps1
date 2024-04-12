# Define the URL
$URL = "https://raw.githubusercontent.com/NiclasHaderer/duckdb-version-manager/main/versions/latest-vm.json"

# Function to download and install duckman
function Download-Duckman {
    # Define download directory
    $DownloadDir = "$env:USERPROFILE\.local\bin"
    New-Item -ItemType Directory -Path $DownloadDir -Force | Out-Null

    # Get JSON content
    $JsonContent = Invoke-RestMethod -Uri $URL

    # Get OS and Architecture
    $OS = $ENV:OS
    $Arch = $ENV:PROCESSOR_ARCHITECTURE

    switch -Regex ($Arch) {
        "AMD64" { $ArchKey = "ArchitectureX86" }
        "ARM64" { $ArchKey = "ArchitectureArm64" }
        default {
            Write-Host "Unsupported architecture: $Arch"
            exit 2
        }
    }

    switch ($OS) {
        "Windows_NT" { $OSKey = "PlatformWindows" }
        default {
            Write-Host "Unsupported OS: $OS"
            exit 3
        }
    }

    # Get download URL
    $DownloadUrl = $JsonContent.platforms.$OSKey.$ArchKey.downloadUrl

    if ($DownloadUrl -eq $null) {
        Write-Host "Failed to find a valid download URL for your platform."
        exit 4
    }

    Write-Host "Downloading from $DownloadUrl..."
    Invoke-WebRequest -Uri $DownloadUrl -OutFile "$DownloadDir\duckman.exe"
    Write-Host "Download complete. duckman is now available in $DownloadDir\duckman.exe"
}

# Function to append line to file if not present
function Append-IfNotPresent {
    param (
        [string]$File,
        [string]$Line
    )
    # Check if the line is already present in the file
    $match = Get-Content -Path $File | Where-Object { $_ -eq $Line }

    if (-not $match) {
        # If the line is not present, append it to the file
        Add-Content -Path $File -Value $Line
    }
}


# Function to setup shells
# Function to setup shells
function Setup-Shells {
    Write-Host "Setting up PATH..."

    # PowerShell
    $PowerShellPathLine = '$env:PATH += ";$env:USERPROFILE\.local\bin"'

    if (!(test-path $profile)) {
        New-Item -Path $profile -Type File -Force
    }

    Append-IfNotPresent $profile $PowerShellPathLine
}


# Function to print shell help
function Print-ShellHelp {
    if (-not ($env:PATH -like "*$env:USERPROFILE\.local\bin*")) {
        Write-Host ""
        Write-Host "$env:USERPROFILE\.local\bin is not in PATH."
    }
}

Download-Duckman

# Prompt user for setup
while ($true) {
    $yn = Read-Host "Do you want duckman to setup autocomplete and PATH for you? (y/n)"
    switch ($yn.ToLower()) {
        "y" { Setup-Shells; exit }
        "n" { Print-ShellHelp; exit }
        default { Write-Host "Please answer yes or no." }
    }
}
