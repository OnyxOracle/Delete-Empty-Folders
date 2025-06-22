# 🧹 Cleanup Utility

A powerful, cross-platform command-line tool to find and manage empty folders, duplicate files, and large directories, helping you keep your filesystem tidy.

![Demo showing the cleanup utility in action](https://github.com/user-attachments/assets/4e6a5b38-8a53-4f59-bed1-93eedb1edfe5)

*(Image: A demonstration of the script identifying and deleting nested empty folders.)*

---

## ✨ Features

- **Extensive Cross-Platform Support**: Pre-compiled binaries are provided for nearly 40 combinations of operating systems and architectures.
- **No Installation Needed**: Download the executable for your OS, and it's ready to run.
- **Three Powerful Commands**:
    - `empty`: Finds and deletes empty folders, with support for recursive scanning and cascading deletion.
    - `large`: Scans and lists the largest directories to help you find what's taking up space.
    - `find`: A versatile tool to find files by various criteria:
        - **Duplicates**: Finds files with identical content using SHA-256, SHA-1, or MD5 hashing.
        - **Size**: Finds files larger than a specified size (e.g., `100MB`, `2GB`).
        - **Age**: Finds files older than a specified duration (e.g., `30d`, `4w`, `12h`).
- **Safe and Interactive**:
    - Always asks for confirmation before deleting.
    - `--dry-run` flag to preview changes without modifying any files.
    - `--trash` flag to move files to the system trash instead of permanently deleting them.
    - Interactive prompts for handling duplicate files, with strategies like `--keep newest` or `--keep oldest`.
- **Highly Configurable**:
    - Use a `.cleanup.yaml` file for persistent settings.
    - Searches for config in the current directory first, then the home directory.
    - Command-line flags always override config file settings for maximum flexibility.
    - Exclude files and directories by name, full path, glob pattern, or regular expression.

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

## ⚙️ Configuration

You can configure the tool using a `.cleanup.yaml` file to avoid typing flags repeatedly.

### How It Works

- **Search Path**: The tool looks for `.cleanup.yaml` in the **current directory first**. If not found, it checks your **home directory**. This allows for project-specific configs.
- **Precedence**: Settings are applied in the following order (highest priority first):
    1.  Command-line flags (e.g., `-r`)
    2.  Environment variables
    3.  Config file (`.cleanup.yaml`)
    4.  Default values

### Creating a Config File

To generate a complete, default configuration file, use the `config init` command.

- **Create in current directory:**
  ```sh
  cleanup config init
  ```
- **Create in your home directory (global):**
  ```sh
  cleanup config init --global
  ```

This will create a `.cleanup.yaml` file with all available settings that you can edit.

---

## 📖 Usage

### Basic Command
```bash
cleanup [command] [OPTIONS]
```

### Examples

**1. Find and delete empty folders recursively (using trash)**
```bash
cleanup empty -r -t
```

**2. Find the top 5 largest folders, excluding build artifacts**
```bash
cleanup large -n 5 --exclude-dirs "build,dist,target" "C:\Users\YourUser\Projects"
```

**3. Find all duplicate images in a directory and delete the oldest of each set**
```bash
# --dry-run is recommended first!
cleanup find --find-duplicates --keep oldest --exclude-glob "*.txt" /path/to/pictures --dry-run
```

**4. Find and permanently delete all `.tmp` files older than 90 days**
```bash
cleanup find --older-than 90d --exclude-glob "*.tmp" --force
```

**5. Get help for a specific command**
```bash
cleanup find --help
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
