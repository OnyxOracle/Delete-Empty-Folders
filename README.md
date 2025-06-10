# 🧹 Empty Folder Cleanup Utility

A powerful and easy-to-use Dart command-line utility to find and delete empty folders, helping you keep your directories tidy. It can search recursively and is smart enough to remove parent directories that only contain other empty folders.

![Demo showing the cleanup utility in action](https://placehold.co/800x400/1e1e2e/dcdcdc?text=Animation+Showing+Script+Usage)
*(Image: A demonstration of the script identifying and deleting nested empty folders.)*

---

## ⚙️ Usage

Open your terminal or command prompt and navigate to the directory where you saved `cleanup.dart`.

### Basic Command

```bash
dart run cleanup.dart [OPTIONS] [PATH]
```

### Arguments and Options

| Argument | Description |
| :--- | :--- |
| `[PATH]` | **(Optional)** The path to the directory you want to scan. If you don't provide a path, the script will use the current directory. |
| `-r` | **(Optional)** Enables **recursive mode**. The script will search all subdirectories and remove parent folders that only contain other empty folders. |
| `-h`, `--help` | **(Optional)** Displays the help message with all available commands and options. |

### Examples

**1. Scan the current directory (top-level only)**
```bash
dart run cleanup.dart
```

**2. Scan a specific directory (e.g., `C:\Users\YourUser\Downloads`)**
```bash
dart run cleanup.dart C:\Users\YourUser\Downloads
```

**3. Scan the current directory recursively**
```bash
dart run cleanup.dart -r
```

**4. Scan a specific directory recursively**
```bash
dart run cleanup.dart -r ./my_projects/
```

**5. Show the help message**
```bash
dart run cleanup.dart --help
```

---

## ✨ Features

- **Find Empty Folders**: Scans a target directory for folders that contain no files.
- **Recursive Search**: Use the `-r` flag to search through all subdirectories.
- **Cascading Deletion**: In recursive mode, it smartly removes parent directories that become empty after their children are deleted.
- **Interactive Prompt**: Always asks for confirmation before deleting any files, showing you exactly what will be removed.
- **Path Flexibility**: Run it on a specific directory (using an absolute or relative path) or simply in the current directory.
- **Helpful Instructions**: Use the `-h` or `--help` flag to see usage instructions at any time.
- **Cross-Platform**: Built with Dart, it runs on Windows, macOS, and Linux.

---

## 🚀 Getting Started

### Prerequisites

You need to have the [Dart SDK](https://dart.dev/get-dart) installed on your system.

### Installation

1.  Clone this repository or just download the `cleanup.dart` file to your local machine.

2.  That's it! The script is ready to run.

---

## 📦 Compiling to an Executable

You can compile the script into a standalone executable so you don't need the Dart SDK to run it. This is great for sharing or for using it as a system-wide command.

```bash
# This will create 'cleanup.exe' (Windows) or 'cleanup' (macOS/Linux)
dart compile exe cleanup.dart
```

You can now run the compiled file directly:

```bash
# On Windows
.\cleanup.exe -r C:\path\to\check

# On macOS/Linux
./cleanup -r /path/to/check
```

---

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Feel free to check the [issues page](https://github.com/your-username/your-repo/issues) if you want to contribute.
