import 'dart:io';

// Function to print the help message.
void printHelpMessage() {
  print('┌───────────────────────────────────┐');
  print('│ 🧹 Empty Folder Cleanup Utility   │');
  print('└───────────────────────────────────┘');
  print('\nFinds and optionally deletes empty folders.');
  print('\nUSAGE:');
  print('  def <script.dart> [OPTIONS] [PATH]');
  print('\nARGUMENTS:');
  print('  [PATH]    The path to the directory to search.');
  print('            If omitted, the current directory is used.');
  print('\nOPTIONS:');
  print('  -r        Search recursively. This will also delete parent folders');
  print('            that only contain other empty folders.');
  print('  -h, --help  Show this help message and exit.');
}


void main(List<String> args) async {
  // --- Argument Parsing ---
  // Check for help argument first. If present, show help and exit.
  if (args.contains('--help') || args.contains('-h')) {
    printHelpMessage();
    exit(0);
  }

  // --- Pretty Header ---
  print('┌───────────────────────────────────┐');
  print('│ 🧹 Empty Folder Cleanup Utility   │');
  print('└───────────────────────────────────┘');

  // Check if the recursive flag '-r' is present.
  final isRecursive = args.contains('-r');
  
  // Find the path argument by removing flags, if any exist.
  final pathArgs = args.where((arg) => !arg.startsWith('-')).toList();
  final path = pathArgs.isNotEmpty ? pathArgs.first : null;

  Directory targetDirectory;

  // Determine the target directory based on the path argument.
  if (path == null) {
    targetDirectory = Directory.current;
    print('ℹ️ No path provided. Using current directory.');
  } else {
    targetDirectory = Directory(path);
    if (!await targetDirectory.exists()) {
      print('🚨 Error: The specified path does not exist or is not a directory.');
      print('   Path: "${targetDirectory.path}"');
      exit(1);
    }
  }

  // --- Search Execution ---
  final searchMode = isRecursive ? 'recursively' : 'at the top-level';
  print('🔍 Searching for empty folders in "${targetDirectory.path}" ($searchMode)...\n');

  final emptyDirectories = <Directory>[];

  try {
    if (isRecursive) {
      // --- Advanced Recursive Logic (Bottom-Up Approach) ---
      final allDirs = <Directory>[];
      // Pass 1: Collect all directories.
      await for (final entity in targetDirectory.list(recursive: true, followLinks: false)) {
          if (entity is Directory) {
              allDirs.add(entity);
          }
      }

      // Sort directories from deepest to shallowest by path length.
      // This is crucial for the bottom-up check.
      allDirs.sort((a, b) => b.path.length.compareTo(a.path.length));

      // Use a Set for efficient lookups of deletable paths.
      final Set<String> deletablePaths = {};

      // Pass 2: Iterate from the bottom up to identify empty directories.
      for (final dir in allDirs) {
          bool containsNonDeletableContent = false;
          // Synchronously list contents. This is acceptable as we expect these
          // directories to have few items if they are candidates for deletion.
          for (final entity in dir.listSync()) {
              if (entity is File) {
                  // If it contains a file, it's not deletable.
                  containsNonDeletableContent = true;
                  break;
              }
              if (entity is Directory) {
                  // If it contains a subdirectory that is NOT marked as deletable,
                  // then this parent directory is also not deletable.
                  if (!deletablePaths.contains(entity.path)) {
                      containsNonDeletableContent = true;
                      break;
                  }
              }
          }

          if (!containsNonDeletableContent) {
              // This directory is empty or only contains other deletable directories.
              emptyDirectories.add(dir);
              deletablePaths.add(dir.path);
          }
      }
    } else {
      // --- Simple Non-Recursive Logic ---
      // Only finds folders that are strictly empty at the top level.
      await for (final entity in targetDirectory.list(recursive: false)) {
        if (entity is Directory && await entity.list().isEmpty) {
          emptyDirectories.add(entity);
        }
      }
    }

    // --- Results and Deletion Prompt ---
    if (emptyDirectories.isEmpty) {
      print('✅ No empty folders were found.');
    } else {
      print('🔎 Found the following empty folders to delete:');
      // The list is already sorted for deletion (deepest first).
      for (final dir in emptyDirectories) {
        print('  📁 ${dir.path}');
      }

      stdout.write('\n❔ Would you like to delete them? (yes/no) ');
      final response = stdin.readLineSync()?.toLowerCase();

      if (response == 'yes' || response == 'y') {
        print('\n🔥 Deleting empty folders...');
        for (final dir in emptyDirectories) {
          try {
            await dir.delete(); // No 'recursive' flag needed due to bottom-up logic
            print('  🗑️ Deleted: ${dir.path}');
          } catch (e) {
            print('  ❌ Error deleting ${dir.path}: $e');
          }
        }
        print('\n✨ Deletion complete.');
      } else {
        print('\n👍 No folders were deleted.');
      }
    }
  } catch (e) {
    print('🚨 An error occurred: $e');
  }
}
