# 🧹 Cleanup Utility

A powerful, cross-platform command-line tool to find and manage empty or large folders, helping you keep your filesystem tidy.

![Demo showing the cleanup utility in action](https://github.com/user-attachments/assets/43a4646a-56d9-493f-aaae-e5378909818b)
*(Image: A demonstration of the script identifying and deleting nested empty folders.)*

---

## ✨ Features

- **Extensive Cross-Platform Support**: Pre-compiled binaries are provided for over 40 combinations of operating systems and architectures (see below).
- **No Installation Needed**: Download the executable for your OS, and it's ready to run.
- **Two Powerful Commands**:
    - `empty`: Finds and deletes empty folders, with support for recursive scanning and cascading deletion.
    - `large`: Scans and lists the largest directories to help you find what's taking up space.
- **Safe and Interactive**:
    - Always asks for confirmation before deleting.
    - `--dry-run` flag to preview changes without modifying any files.
    - `--trash` flag to move files to the system trash instead of permanently deleting them.
- **Highly Configurable**: Use flags to ignore specific files, exclude directories, filter by age, and more.

---

## 🚀 Downloads

Download the specific executable for your system from the table below, or find all binaries packaged in a `.zip` file on the [**GitHub Releases**](https://github.com/OnyxOracle/Delete-Empty-Folders/releases) page.

| Operating System | Architecture | Download Link |
| :--------------- | :----------- | :------------ |
| **Windows** | 64-bit (x64) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-windows-amd64.exe)  |
|                  | 32-bit (x86) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-windows-386.exe)  |
|                  | ARM 64-bit   | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-windows-arm64.exe)  |
|                  | ARM 32-bit   | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-windows-arm.exe)  |
| **macOS** | Apple Silicon (ARM 64-bit)| [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-darwin-arm64)  |
|                  | Intel 64-bit (x64)        | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-darwin-amd64)  |
| **Linux** | 64-bit (x64) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-linux-amd64)  |
|                  | 32-bit (x86) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-linux-386)  |
|                  | ARM 64-bit   | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-linux-arm64)  |
|                  | ARM 32-bit   | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-linux-arm)  |
|                  | LoongArch 64-bit | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-linux-loong64)  |
|                  | PowerPC 64-bit (Big Endian) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-linux-ppc64)  |
|                  | PowerPC 64-bit (Little Endian)| [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-linux-ppc64le)  |
|                  | RISC-V 64-bit| [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-linux-riscv64)  |
|                  | IBM Z (s390x)| [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-linux-s390x)  |
|                  | MIPS 64-bit (Big Endian) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-linux-mips64)  |
|                  | MIPS 64-bit (Little Endian) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-linux-mips64le)  |
|                  | MIPS 32-bit (Big Endian) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-linux-mips)  |
|                  | MIPS 32-bit (Little Endian) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-linux-mipsle)  |
| **FreeBSD** | 64-bit (x64) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-freebsd-amd64)  |
|                  | 32-bit (x86) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-freebsd-386)  |
|                  | ARM 64-bit   | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-freebsd-arm64)  |
|                  | ARM 32-bit   | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-freebsd-arm)  |
|                  | RISC-V 64-bit| [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-freebsd-riscv64)  |
| **OpenBSD** | 64-bit (x64) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-openbsd-amd64)  |
|                  | 32-bit (x86) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-openbsd-386)  |
|                  | ARM 64-bit   | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-openbsd-arm64)  |
|                  | ARM 32-bit   | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-openbsd-arm)  |
|                  | PowerPC 64-bit | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-openbsd-ppc64)  |
|                  | RISC-V 64-bit| [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-openbsd-riscv64)  |
| **NetBSD** | 64-bit (x64) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-netbsd-amd64)  |
|                  | 32-bit (x86) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-netbsd-386)  |
|                  | ARM 64-bit   | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-netbsd-arm64)  |
|                  | ARM 32-bit   | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-netbsd-arm)  |
| **DragonFly BSD**| 64-bit (x64) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-dragonfly-amd64)  |
| **illumos** | 64-bit (x64) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-illumos-amd64)  |
| **Solaris** | 64-bit (x64) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-solaris-amd64)  |
| **Plan 9** | 64-bit (x64) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-plan9-amd64)  |
|                  | 32-bit (x86) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-plan9-386)  |
|                  | ARM 32-bit   | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-plan9-arm)  |
| **AIX** | PowerPC 64-bit | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/latest/download/cleanup-aix-ppc64)  |


---

## 🔧 Installation (System-Wide Access)

To run the `cleanup` command from any directory, you need to place the executable in a location listed in your system's `PATH` environment variable.

### macOS & Linux

1.  Download the correct binary for your system (e.g., `cleanup-linux-amd64`).
2.  Open your terminal and make the binary executable:
    ```sh
    chmod +x ./cleanup-linux-amd64
    ```
3.  Move it to `/usr/local/bin`, which is a standard directory for user-installed executables. You can also rename it to just `cleanup` for convenience.
    ```sh
    sudo mv ./cleanup-linux-amd64 /usr/local/bin/cleanup
    ```
4.  You can now run the tool from anywhere by typing `cleanup`.

### Windows

1.  Download the executable (e.g., `cleanup-windows-amd64.exe`) and rename it to `cleanup.exe` for convenience.
2.  Create a folder where you will keep your command-line tools (e.g., `C:\ProgramFiles\tools`).
3.  Move `cleanup.exe` into that folder.
4.  Now, add that folder to your system's `PATH`:
    * Press the Windows key and type "environment variables".
    * Select "Edit the system environment variables".
    * Click the "Environment Variables..." button.
    * In the "System variables" section, find and select the `Path` variable, then click "Edit...".
    * Click "New" and add the path to your tools folder (e.g., `C:\ProgramFiles\tools`).
    * Click OK on all windows to save.
5.  Open a **new** terminal window. You can now run the tool from anywhere by typing `cleanup`.

---

## ⚙️ Usage

After installation, you can run the tool from any terminal window.

### Basic Command
```bash
cleanup [command] [OPTIONS]
```

### Examples

**1. Find empty folders recursively in the current directory.**
```bash
cleanup empty -r --dry-run
```

**2. Find the top 5 largest folders in your Downloads folder.**
```bash
cleanup large -n 5 "C:\Users\YourUser\Downloads"
```

**3. Get help for a specific command.**
```bash
cleanup empty --help
```

---

## 📦 Building from Source

If you prefer to compile the script yourself, you will need the [Go SDK](https://go.dev/dl/).

1.  Clone this repository: `git clone https://github.com/OnyxOracle/Delete-Empty-Folders.git`
2.  Navigate into the directory: `cd Delete-Empty-Folders`
3.  Install dependencies: `go mod tidy`

### Simple Build (For Your Current System)
To build a single executable for your current operating system and architecture, run:
```bash
go build -o cleanup .
```
Or, for Windows run:
```bash
go build -o cleanup.exe .
```

### Build All Binaries
To generate all the executables listed in the downloads table, use the provided build scripts.

* **On Windows:** Run `build-all.bat` or `build-all.ps1`
* **On Linux/macOS:** First, make the script executable with `chmod +x build-all.sh`, then run `./build-all.sh`

All generated binaries will be placed in a `builds` folder.

---

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Feel free to check the [issues page](https://github.com/OnyxOracle/Delete-Empty-Folders/issues).
