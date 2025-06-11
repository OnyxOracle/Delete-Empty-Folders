# 🧹 Cleanup Utility

A powerful, cross-platform command-line tool to find and manage empty or large folders, helping you keep your filesystem tidy.

![Demo showing the cleanup utility in action](https://placehold.co/800x400/1e1e2e/dcdcdc?text=Animation+Showing+Script+Usage)
*(Image: A demonstration of the script identifying and deleting nested empty folders.)*

---

## ✨ Features

- **Extensive Cross-Platform Support**: Pre-compiled binaries are provided for over 30 combinations of operating systems and architectures (see below).
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

You can find all binaries conveniently packaged in a `.zip` file on the [**GitHub Releases**](https://github.com/varvaruk-v/Delete-Empty-Folders/releases) page.

Alternatively, download the specific executable for your system from the table below.

| Operating System | Architecture | Download Link |
| :--------------- | :----------- | :------------ |
| **Windows** | 64-bit (x64) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-windows-amd64.exe)  |
|                  | 32-bit (x86) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-windows-386.exe)  |
|                  | ARM 64-bit   | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-windows-arm64.exe)  |
|                  | ARM 32-bit   | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-windows-arm.exe)  |
| **macOS** | Apple Silicon (ARM 64-bit)| [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-darwin-arm64)  |
|                  | Intel 64-bit (x64)     | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-darwin-amd64)  |
| **Linux** | 64-bit (x64) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-linux-amd64)  |
|                  | 32-bit (x86) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-linux-386)  |
|                  | ARM 64-bit   | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-linux-arm64)  |
|                  | ARM 32-bit   | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-linux-arm)  |
|                  | RISC-V 64-bit| [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-linux-riscv64)  |
|                  | PowerPC 64-bit (Big Endian) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-linux-ppc64)  |
|                  | PowerPC 64-bit (Little Endian)| [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-linux-ppc64le)  |
|                  | MIPS 64-bit (Big Endian) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-linux-mips64)  |
|                  | MIPS 64-bit (Little Endian) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-linux-mips64le)  |
|                  | MIPS 32-bit (Big Endian) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-linux-mips)  |
|                  | MIPS 32-bit (Little Endian) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-linux-mipsle)  |
| **FreeBSD** | 64-bit (x64) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-freebsd-amd64)  |
|                  | 32-bit (x86) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-freebsd-386)  |
|                  | ARM 64-bit   | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-freebsd-arm64)  |
|                  | ARM 32-bit   | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-freebsd-arm)  |
| **OpenBSD** | 64-bit (x64) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-openbsd-amd64)  |
|                  | 32-bit (x86) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-openbsd-386)  |
|                  | ARM 64-bit   | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-openbsd-arm64)  |
|                  | ARM 32-bit   | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-openbsd-arm)  |
| **NetBSD** | 64-bit (x64) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-netbsd-amd64)  |
|                  | 32-bit (x86) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-netbsd-386)  |
|                  | ARM 64-bit   | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-netbsd-arm64)  |
|                  | ARM 32-bit   | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-netbsd-arm)  |
| **DragonFly BSD**| 64-bit (x64) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-dragonfly-amd64)  |
| **Solaris** | 64-bit (x64) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-solaris-amd64)  |
| **Plan 9** | 64-bit (x64) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-plan9-amd64)  |
|                  | 32-bit (x86) | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-plan9-386)  |
|                  | ARM 32-bit   | [Download](https://github.com/OnyxOracle/Delete-Empty-Folders/releases/download/v.1.1.0/cleanup-plan9-arm)  |

---

## ⚙️ Usage

After downloading, you may need to make the file executable. On macOS and Linux, run: `chmod +x ./cleanup-linux-amd64`

Open your terminal and run the program.

### Basic Command

```bash
# On Windows
.\cleanup-windows-amd64.exe [command] [OPTIONS]

# On macOS/Linux
./cleanup-linux-amd64 [command] [OPTIONS]
```

### Examples

**1. Find empty folders recursively in the current directory.**
```bash
./cleanup empty -r --dry-run
```

**2. Find the top 5 largest folders in your Downloads folder.**
```bash
./cleanup large -n 5 "C:\Users\YourUser\Downloads"
```

**3. Get help for a specific command.**
```bash
./cleanup empty --help
```

---

## 📦 Building from Source (Optional)

If you prefer to compile the script yourself, you will need the [Go SDK](https://go.dev/dl/).

1.  Clone this repository: `git clone https://github.com/varvaruk-v/Delete-Empty-Folders.git`
2.  Navigate into the directory: `cd Delete-Empty-Folders`
3.  Install dependencies: `go mod tidy`
4.  Run the appropriate build script for your system:
    - On Windows: `.\build-all.bat` or `.\build-all.ps1`
    - On Linux/macOS: `chmod +x ./build-all.sh` then `./build-all.sh`

---

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Feel free to check the [issues page](https://github.com/varvaruk-v/Delete-Empty-Folders/issues).
