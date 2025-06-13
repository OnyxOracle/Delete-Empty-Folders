package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	// Third-party libraries for advanced features.
	// User must run 'go mod tidy' to install these.
	"github.com/hymkor/trash-go" // Correct, native Go library for moving files to trash.
	"github.com/spf13/cobra"      // The best library for modern CLI applications.
)

// --- Global Configuration & State ---

// config holds all the flags that can be set by the user.
var config struct {
	isRecursive  bool
	dryRun       bool
	force        bool
	verbose      bool
	quiet        bool
	useTrash     bool
	ignoreFiles  []string
	excludeDirs  []string
	olderThanStr string
	logFile      string
	topN         int
}

// logger is our global logger. It will be configured in main().
var logger *log.Logger

// --- cobra Command Definitions ---

// rootCmd is the main entry point for the application.
var rootCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "A powerful utility to find and manage empty or large folders.",
	Long: `Cleanup is a command-line tool that helps you keep your filesystem tidy.
It can find and delete empty folders (recursively) and also find the largest folders to help you manage disk space.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// This function runs before any subcommand.
		// It's the perfect place to set up our global logger.
		configureLogger()
	},
}

// emptyCmd represents the 'empty' subcommand for finding empty folders.
var emptyCmd = &cobra.Command{
	Use:   "empty [PATH]",
	Short: "Find and delete empty folders",
	Long: `Scans a directory to find empty folders.
In recursive mode, it can also find parent directories that only contain other empty folders.
You will be prompted before any deletion occurs unless the --force flag is used.`,
	Args: cobra.MaximumNArgs(1), // Accepts zero or one argument for the path.
	Run:  runEmpty,
}

// largeCmd represents the 'large' subcommand for finding large folders.
var largeCmd = &cobra.Command{
	Use:   "large [PATH]",
	Short: "Find the largest folders in a directory",
	Long:  `Scans a directory recursively to find the folders that consume the most disk space.`,
	Args:  cobra.MaximumNArgs(1),
	Run:   runLarge,
}

// --- Initialization ---

// init() is a special Go function that runs when the program starts.
// We use it to define all our flags and set up the command structure.
func init() {
	// Add flags to the 'empty' subcommand.
	emptyCmd.Flags().BoolVarP(&config.isRecursive, "recursive", "r", false, "Search recursively and delete parent folders that become empty.")
	emptyCmd.Flags().BoolVar(&config.dryRun, "dry-run", false, "Perform a trial run without making any changes.")
	emptyCmd.Flags().BoolVarP(&config.force, "force", "y", false, "Skip the confirmation prompt and delete all found folders.")
	emptyCmd.Flags().BoolVar(&config.useTrash, "trash", false, "Move folders to system trash instead of permanently deleting.")
	emptyCmd.Flags().StringSliceVar(&config.ignoreFiles, "ignore-files", []string{".DS_Store", "Thumbs.db"}, "Comma-separated list of files to ignore (e.g., .DS_Store,Thumbs.db).")
	emptyCmd.Flags().StringSliceVar(&config.excludeDirs, "exclude-dirs", []string{".git", "node_modules"}, "Comma-separated list of directories to exclude from scanning.")
	emptyCmd.Flags().StringVar(&config.olderThanStr, "older-than", "", "Only consider folders older than a duration (e.g., 30d, 48h, 2w).")

	// Add flags to the 'large' subcommand.
	largeCmd.Flags().IntVarP(&config.topN, "top", "n", 10, "Number of largest folders to show.")
	largeCmd.Flags().StringSliceVar(&config.excludeDirs, "exclude-dirs", []string{".git", "node_modules"}, "Comma-separated list of directories to exclude from scanning.")

	// Add global flags that apply to all subcommands.
	rootCmd.PersistentFlags().BoolVarP(&config.verbose, "verbose", "v", false, "Enable verbose output.")
	rootCmd.PersistentFlags().BoolVarP(&config.quiet, "quiet", "q", false, "Suppress all output except for errors and final prompts.")
	rootCmd.PersistentFlags().StringVar(&config.logFile, "log-file", "", "Path to a file to write logs to.")

	// Add the subcommands to the root command.
	rootCmd.AddCommand(emptyCmd)
	rootCmd.AddCommand(largeCmd)
}

// --- Main Execution Logic ---

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// runEmpty is the main function for the 'empty' subcommand.
func runEmpty(cmd *cobra.Command, args []string) {
	// --- Setup ---
	olderThanDuration, err := parseDuration(config.olderThanStr)
	if err != nil {
		logError("💥 Invalid duration string for --older-than: %v", err)
		os.Exit(1)
	}

	targetDir := getTargetDir(args)
	printEmptyModeSummary(targetDir, olderThanDuration)

	var emptyDirs []string
	if config.isRecursive {
		emptyDirs, err = findEmptyRecursive(targetDir, olderThanDuration)
	} else {
		emptyDirs, err = findEmptyTopLevel(targetDir, olderThanDuration)
	}

	if err != nil {
		logError("💥 An error occurred during scan: %v", err)
		os.Exit(1)
	}

	// --- Results and Deletion ---
	if len(emptyDirs) == 0 {
		logInfo("\n🎉 Success! No empty folders were found.")
		return
	}

	logInfo("\n🔎 Found %d empty folders to process:", len(emptyDirs))
	for _, dir := range emptyDirs {
		logInfo("  📁 %s", dir)
	}

	if config.dryRun {
		logInfo("\n🧪 --dry-run enabled. No changes will be made.")
		return
	}

	if !config.force {
		fmt.Printf("\n\033[34m🤔 Would you like to %s them? [Yes/No] \033[0m", getActionString())
		reader := bufio.NewReader(os.Stdin)
		response, _ := reader.ReadString('\n')
		response = strings.ToLower(strings.TrimSpace(response))
		if response != "yes" && response != "y" {
			logInfo("\n👍 OK. No folders were changed.")
			return
		}
	}

	logInfo("\n🔥 Processing folders...")
	processedCount := 0
	for _, dir := range emptyDirs {
		var opErr error
		if config.useTrash {
			opErr = trash.Throw(dir)
		} else {
			opErr = os.RemoveAll(dir)
		}

		if opErr != nil {
			logError("  ❌ Error %s %s: %v", getActionStringPast(), dir, opErr)
		} else {
			logInfo("  %s %s: %s", getActionIcon(), getActionStringPast(), dir)
			processedCount++
		}
	}
	logInfo("\n✨ All done! %s %d folder(s).", getActionStringPast(), processedCount)
}

// runLarge is the main function for the 'large' subcommand.
func runLarge(cmd *cobra.Command, args []string) {
	// --- Setup ---
	targetDir := getTargetDir(args)
	logInfo("--- 📊 Find Large Folders Mode ---")
	logInfo("🎯 Target Directory: %s", targetDir)
	logInfo("📈 Showing Top: %d folders", config.topN)
	if len(config.excludeDirs) > 0 {
		logInfo("🚫 Excluding: %s", strings.Join(config.excludeDirs, ", "))
	}
	logInfo("----------------------------------\n")
	logInfo("⏳ Scanning for largest folders... this may take a while.")

	// --- Execution ---
	dirSizes, err := calculateDirectorySizes(targetDir)
	if err != nil {
		logError("💥 Failed to calculate directory sizes: %v", err)
		os.Exit(1)
	}

	// --- Results ---
	type dirInfo struct {
		Path string
		Size int64
	}

	var sortedDirs []dirInfo
	for path, size := range dirSizes {
		sortedDirs = append(sortedDirs, dirInfo{path, size})
	}

	sort.Slice(sortedDirs, func(i, j int) bool {
		return sortedDirs[i].Size > sortedDirs[j].Size
	})

	logInfo("\n🔎 Top %d largest folders:", config.topN)
	for i := 0; i < len(sortedDirs) && i < config.topN; i++ {
		logInfo("  %s - %s", formatBytes(sortedDirs[i].Size), sortedDirs[i].Path)
	}
}

// --- Core Logic & Helper Functions ---

func printEmptyModeSummary(targetDir string, olderThan time.Duration) {
	logInfo("--- 🧹 Find Empty Folders Mode ---")
	logInfo("🎯 Target Directory: %s", targetDir)
	if config.isRecursive {
		logInfo("🌲 Mode: Recursive Scan Enabled")
	} else {
		logInfo("📂 Mode: Top-Level Scan Only")
	}
	if config.dryRun {
		logInfo("🧪 Action: Dry Run (no changes will be made)")
	} else if config.useTrash {
		logInfo("♻️  Action: Move to System Trash")
	} else {
		logInfo("🗑️  Action: Permanent Deletion")
	}
	if config.force {
		logInfo("❗ Confirmation: Skipped (--force enabled)")
	}
	if olderThan > 0 {
		logInfo("⏳ Age Filter: Only folders older than %s", config.olderThanStr)
	}
	if len(config.ignoreFiles) > 0 {
		logInfo("🙈 Ignoring Files: %s", strings.Join(config.ignoreFiles, ", "))
	}
	if len(config.excludeDirs) > 0 {
		logInfo("🚫 Excluding Dirs: %s", strings.Join(config.excludeDirs, ", "))
	}
	logInfo("----------------------------------\n")
}

// findEmptyRecursive performs a bottom-up search.
func findEmptyRecursive(path string, olderThan time.Duration) ([]string, error) {
	var allDirs []string
	deletablePaths := make(map[string]bool)
	ignoreFileSet := stringSliceToSet(config.ignoreFiles)
	excludeDirSet := stringSliceToSet(config.excludeDirs)

	err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if _, excluded := excludeDirSet[info.Name()]; excluded {
				logVerbose("Excluding directory: %s", p)
				return filepath.SkipDir
			}
			allDirs = append(allDirs, p)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Slice(allDirs, func(i, j int) bool {
		return len(strings.Split(allDirs[i], string(os.PathSeparator))) > len(strings.Split(allDirs[j], string(os.PathSeparator)))
	})

	for _, dir := range allDirs {
		if dir == path {
			continue
		}

		info, err := os.Stat(dir)
		if err != nil {
			continue
		}

		if olderThan > 0 && time.Since(info.ModTime()) < olderThan {
			logVerbose("Skipping recent directory: %s", dir)
			continue
		}

		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}

		isEmpty := true
		for _, entry := range entries {
			if !entry.IsDir() {
				if _, ignored := ignoreFileSet[entry.Name()]; !ignored {
					isEmpty = false
					break
				}
			} else {
				if !deletablePaths[filepath.Join(dir, entry.Name())] {
					isEmpty = false
					break
				}
			}
		}

		if isEmpty {
			deletablePaths[dir] = true
		}
	}

	var emptyDirs []string
	for dir := range deletablePaths {
		emptyDirs = append(emptyDirs, dir)
	}
	sort.Strings(emptyDirs)
	return emptyDirs, nil
}

// findEmptyTopLevel searches only the immediate children of the target directory.
func findEmptyTopLevel(path string, olderThan time.Duration) ([]string, error) {
	var emptyDirs []string
	ignoreFileSet := stringSliceToSet(config.ignoreFiles)
	excludeDirSet := stringSliceToSet(config.excludeDirs)

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		if _, excluded := excludeDirSet[entry.Name()]; excluded {
			logVerbose("Excluding directory: %s", entry.Name())
			continue
		}

		dirPath := filepath.Join(path, entry.Name())
		info, err := os.Stat(dirPath)
		if err != nil {
			continue
		}

		if olderThan > 0 && time.Since(info.ModTime()) < olderThan {
			logVerbose("Skipping recent directory: %s", dirPath)
			continue
		}

		subEntries, err := os.ReadDir(dirPath)
		if err != nil {
			continue
		}

		isEmpty := true
		for _, subEntry := range subEntries {
			if !subEntry.IsDir() {
				if _, ignored := ignoreFileSet[subEntry.Name()]; !ignored {
					isEmpty = false
					break
				}
			} else {
				isEmpty = false
				break
			}
		}

		if isEmpty {
			emptyDirs = append(emptyDirs, dirPath)
		}
	}

	sort.Strings(emptyDirs)
	return emptyDirs, nil
}

// calculateDirectorySizes walks the filesystem and calculates the size of each directory.
func calculateDirectorySizes(root string) (map[string]int64, error) {
	dirSizes := make(map[string]int64)
	excludeDirSet := stringSliceToSet(config.excludeDirs)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if _, excluded := excludeDirSet[info.Name()]; excluded {
				return filepath.SkipDir
			}
			return nil
		}

		size := info.Size()
		for p := filepath.Dir(path); strings.HasPrefix(p, root) || p == root; p = filepath.Dir(p) {
			dirSizes[p] += size
		}
		return nil
	})

	return dirSizes, err
}

// --- Utility Functions ---

func configureLogger() {
	var out io.Writer = os.Stdout
	if config.quiet {
		out = io.Discard
	} else if config.logFile != "" {
		file, err := os.OpenFile(config.logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}
		if config.verbose {
			out = io.MultiWriter(os.Stdout, file)
		} else {
			out = file
		}
	}
	logger = log.New(out, "", 0)
}

func logInfo(format string, v ...interface{}) {
	if !config.quiet {
		logger.Printf(format, v...)
	}
}

func logInfoBlue(format string, v ...interface{}) {
	if !config.quiet {
		logger.Printf("\033[34m"+format+"\033[0m", v...)
	}
}

func logVerbose(format string, v ...interface{}) {
	if config.verbose && !config.quiet {
		logger.Printf("[VERBOSE] "+format, v...)
	}
}

func logError(format string, v ...interface{}) {
	if !config.quiet {
		logger.Printf("\033[31m[ERROR] "+format+"\033[0m", v...)
	}
}

func getTargetDir(args []string) string {
	pathArg := "."
	if len(args) > 0 {
		pathArg = args[0]
	}
	targetDir, err := filepath.Abs(pathArg)
	if err != nil {
		logError("Could not resolve path '%s': %v", pathArg, err)
		os.Exit(1)
	}
	info, err := os.Stat(targetDir)
	if os.IsNotExist(err) || !info.IsDir() {
		logError("The specified path does not exist or is not a directory: %s", targetDir)
		os.Exit(1)
	}
	return targetDir
}

func stringSliceToSet(slice []string) map[string]struct{} {
	set := make(map[string]struct{}, len(slice))
	for _, item := range slice {
		set[item] = struct{}{}
	}
	return set
}

func parseDuration(durationStr string) (time.Duration, error) {
	if durationStr == "" {
		return 0, nil
	}
	durationStr = strings.ToLower(durationStr)
	var multiplier time.Duration
	if strings.HasSuffix(durationStr, "d") {
		multiplier = 24 * time.Hour
		durationStr = strings.TrimSuffix(durationStr, "d")
	} else if strings.HasSuffix(durationStr, "w") {
		multiplier = 7 * 24 * time.Hour
		durationStr = strings.TrimSuffix(durationStr, "w")
	} else if strings.HasSuffix(durationStr, "h") {
		multiplier = time.Hour
		durationStr = strings.TrimSuffix(durationStr, "h")
	} else {
		return 0, fmt.Errorf("unknown duration suffix. Use 'd', 'w', or 'h'")
	}

	var num int
	_, err := fmt.Sscanf(durationStr, "%d", &num)
	if err != nil {
		return 0, err
	}

	return time.Duration(num) * multiplier, nil
}

func formatBytes(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}

func getActionString() string {
	if config.useTrash {
		return "move to trash"
	}
	return "delete"
}

func getActionStringPast() string {
	if config.useTrash {
		return "Moved to trash"
	}
	return "Deleted"
}

func getActionIcon() string {
	if config.useTrash {
		return "♻️"
	}
	return "🗑️"
}