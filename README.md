# 🧹 Empty Folder Cleanup Utility

A powerful and easy-to-use command-line utility to find and delete empty folders, helping you keep your directories tidy. It can search recursively and is smart enough to remove parent directories that only contain other empty folders.

![Demo showing the cleanup utility in action](https://placehold.co/800x400/1e1e2e/dcdcdc?text=Animation+Showing+Script+Usage)
*(Image: A demonstration of the script identifying and deleting nested empty folders.)*

---

## ✨ Features

- **No Installation Needed**: Download the executable for your OS and run it directly.
- **Cross-Platform**: Pre-compiled binaries provided for Windows, macOS, and Linux.
- **Recursive Search**: Use the `-r` flag to search through all subdirectories.
- **Cascading Deletion**: In recursive mode, it smartly removes parent directories that become empty after their children are deleted.
- **Interactive Prompt**: Always asks for confirmation before deleting any files, showing you exactly what will be removed.
- **Helpful Instructions**: Use the `-h` or `--help` flag to see usage instructions at any time.

---

## 🚀 Installation

1.  Go to the [**Releases**](https://github.com/your-username/your-repo/releases) page.
2.  Download the latest executable for your operating system (Windows, macOS, or Linux).

### For macOS and Linux

You may need to make the file executable after downloading it. Open your terminal and run:

```bash
chmod +x ./cleanup
```

---

## ⚙️ Usage

Open your terminal or command prompt in the directory where you downloaded the executable.

### Basic Command

```bash
# On Windows
.\cleanup.exe [OPTIONS] [PATH]

# On macOS/Linux
./cleanup [OPTIONS] [PATH]
```

### Arguments and Options

| Argument | Description |
| :--- | :--- |
| `[PATH]` | **(Optional)** The path to the directory you want to scan. If you don't provide a path, the script will use the current directory. |
| `-r` | **(Optional)** Enables **recursive mode**. The script will search all subdirectories and remove parent folders that only contain other empty folders. |
| `-h`, `--help` | **(Optional)** Displays the help message with all available commands and options. |

### Examples

**1. Scan the current directory recursively on Windows**
```bash
.\cleanup.exe -r
```

**2. Scan a specific directory on macOS or Linux**
```bash
./cleanup /home/user/documents/projects
```

---

## 📦 Building from Source (Optional)

If you prefer to compile the script yourself, you will need the [Dart SDK](https://dart.dev/get-dart).

1.  Clone this repository or download the `cleanup.dart` file.
2.  Run the following command to compile:

```bash
# This will create 'cleanup.exe' (Windows) or 'cleanup' (macOS/Linux)
dart compile exe cleanup.dart
```

---

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Feel free to check the [issues page](https://github.com/your-username/your-repo/issues) if you want to contribute.