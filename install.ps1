$URL = "https://raw.githubusercontent.com/NiclasHaderer/duckdb-version-manager/main/versions/latest-vm.json"

function Download-Duckman {
    $DownloadDir = "$env:USERPROFILE\.local\bin"
    New-Item -ItemType Directory -Path $DownloadDir -Force | Out-Null

    $JsonContent = Invoke-RestMethod -Uri $URL

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

    $DownloadUrl = $JsonContent.platforms.$OSKey.$ArchKey.downloadUrl

    if ($DownloadUrl -eq $null) {
        Write-Host "Failed to find a valid download URL for your platform."
        exit 4
    }

    Write-Host "Downloading from $DownloadUrl..."
    Invoke-WebRequest -Uri $DownloadUrl -OutFile "$DownloadDir\duckman.exe"
    Write-Host "Download complete. duckman is now available in $DownloadDir\duckman.exe"
}

function Append-IfNotPresent {
    param (
        [string]$File,
        [string]$Line
    )
    $match = Get-Content -Path $File | Where-Object { $_ -eq $Line }

    if (-not $match) {
        Add-Content -Path $File -Value $Line
    }
}


function Setup-Shells {
    Write-Host "Setting up PATH..."

    # PowerShell
    $PowerShellPathLine = '$env:PATH += ";$env:USERPROFILE\.local\bin"'

    if (!(test-path $profile)) {
        New-Item -Path $profile -Type File -Force
    }

    Append-IfNotPresent $profile $PowerShellPathLine
}


function Print-ShellHelp {
    if (-not ($env:PATH -like "*$env:USERPROFILE\.local\bin*")) {
        Write-Host ""
        Write-Host "$env:USERPROFILE\.local\bin is not in PATH."
    }
}


while ($true) {
    $yn = Read-Host "Do you want duckman to setup autocomplete and PATH for you? (y/n)"
    switch ($yn.ToLower()) {
        "y" { Setup-Shells; break }
        "n" { Print-ShellHelp; break }
        default { Write-Host "Please answer yes or no." }
    }
    if ($yn.ToLower() -eq 'y' -or $yn.ToLower() -eq 'n') {
        break
    }
}