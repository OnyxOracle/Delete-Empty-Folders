# 🧹 Cleanup Utility

A powerful, cross-platform command-line tool to find and manage empty or large folders, helping you keep your filesystem tidy.

![Demo showing the cleanup utility in action](https://placehold.co/800x400/1e1e2e/dcdcdc?text=Animation+Showing+Script+Usage)
*(Image: A demonstration of the script identifying and deleting nested empty folders.)*

---

## ✨ Features

- **Extensive Cross-Platform Support**: Pre-compiled binaries are provided for over 20 combinations of operating systems and architectures (see below).
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

You can find all binaries conveniently packaged in a `.zip` file on the [**GitHub Releases**](https://github.com/your-username/your-repo/releases) page.

Alternatively, download the specific executable for your system from the table below.

| Operating System | Architecture | Download Link |
| :--------------- | :----------- | :------------ |
| **Windows** | 64-bit (x64) | [Download]()  |
|                  | 32-bit (x86) | [Download]()  |
| **macOS** | Apple Silicon| [Download]()  |
|                  | Intel        | [Download]()  |
| **Linux** | 64-bit (x64) | [Download]()  |
|                  | 32-bit (x86) | [Download]()  |
|                  | ARM 64-bit   | [Download]()  |
|                  | ARM 32-bit   | [Download]()  |
|                  | PowerPC 64-bit (Big Endian) | [Download]()  |
|                  | PowerPC 64-bit (Little Endian)| [Download]()  |
|                  | MIPS 64-bit  | [Download]()  |
|                  | MIPS 64-bit (Little Endian) | [Download]()  |
| **FreeBSD** | 64-bit (x64) | [Download]()  |
|                  | 32-bit (x86) | [Download]()  |
|                  | ARM          | [Download]()  |
| **OpenBSD** | 64-bit (x64) | [Download]()  |
|                  | 32-bit (x86) | [Download]()  |
|                  | ARM          | [Download]()  |
| **NetBSD** | 64-bit (x64) | [Download]()  |
|                  | 32-bit (x86) | [Download]()  |
|                  | ARM          | [Download]()  |
| **DragonFly BSD**| 64-bit (x64) | [Download]()  |
| **Solaris** | 64-bit (x64) | [Download]()  |
| **Plan 9** | 64-bit (x64) | [Download]()  |
|                  | 32-bit (x86) | [Download]()  |

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

1.  Clone this repository: `git clone https://github.com/your-username/your-repo.git`
2.  Navigate into the directory: `cd your-repo`
3.  Install dependencies: `go mod tidy`
4.  Run the build script:
    - On Windows: `.\build-all.bat` or `.\build-all.ps1`
    - On Linux/macOS: `chmod +x ./build-all.sh` then `./build-all.sh`

---

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Feel free to check the [issues page](https://github.com/your-username/your-repo/issues).