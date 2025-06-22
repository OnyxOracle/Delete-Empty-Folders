// Package main defines the entry point for the cleanup command-line utility.
// Being in package 'main' means this file can be compiled into an executable program.
package main

import (
	// Standard library imports
	"bufio"         // For buffered I/O, used for reading user input from the console.
	"context"       // For managing cancellation and deadlines, crucial for graceful shutdown (e.g., on Ctrl+C).
	"crypto/md5"    // Implements the MD5 hash algorithm, an option for finding duplicates.
	"crypto/sha1"   // Implements the SHA-1 hash algorithm, an option for finding duplicates.
	"crypto/sha256" // Implements the SHA-256 hash algorithm, the default for finding duplicates.
	"encoding/csv"  // For writing output in CSV format when requested by the user.
	"encoding/json" // For writing output in JSON format when requested by the user.
	"errors"        // For creating and inspecting standard error values.
	"fmt"           // Provides functions for formatted I/O (like printing to the console).
	"hash"          // Provides a common interface for cryptographic hash functions (md5, sha1, sha256).
	"io"            // Provides basic I/O interfaces, like io.Writer for handling different output streams.
	"log"           // Provides simple logging capabilities.
	"os"            // Provides a platform-independent interface to operating system functionality.
	"os/signal"     // For capturing operating system signals, allowing the program to react to Ctrl+C.
	"path/filepath" // For manipulating filesystem paths in a way that is safe across different OSes (Windows, Linux, macOS).
	"regexp"        // For regular expression matching to exclude files/folders based on patterns.
	"runtime"       // Provides interaction with the Go runtime, used here to get the number of CPUs for parallel processing.
	"sort"          // Provides sorting algorithms for slices.
	"strconv"       // For string conversions (e.g., string to integer or float).
	"strings"       // For string manipulation functions, like splitting and replacing.
	"sync"          // Provides synchronization primitives, like Mutexes (for safe concurrent access) and WaitGroups (for managing goroutines).
	"syscall"       // Provides a low-level interface to the underlying operating system's calls (for signals).
	"time"          // For time-related operations, like parsing durations and checking file modification times.

	// Third-party library imports
	"github.com/hymkor/trash-go"        // A cross-platform library for moving files to the system's trash/recycle bin.
	"github.com/schollz/progressbar/v3" // A library for displaying progress bars in the terminal.
	"github.com/spf13/cobra"            // A powerful library for creating modern command-line applications with subcommands and flags.
	"github.com/spf13/viper"            // A library for application configuration, handling files, environment variables, and flags.
	"gopkg.in/yaml.v3"                  // A library for working with YAML files, used for the config file.
)

// --- Version Information ---
// These variables hold build-time information about the application.
// They can be set during compilation using ldflags to embed version details.
var (
	version    = "v2.0.0"                                             // Application version number.
	date       = "22.06.2025"                                         // Build date.
	appName    = "Cleanup Utility"                                    // Application name.
	sourceCode = "https://github.com/OnyxOracle/Delete-Empty-Folders" // Link to the source code.
	copyright  = "Copyright (c) 2025 OnyxOracle"                      // Copyright notice.
)

// --- Global Configuration ---
// Config struct holds all the configuration parameters for the application.
// - `mapstructure` tags are for Viper to read keys from the config file (e.g., 'use-trash' in YAML maps to UseTrash field).
// - `yaml` tags are for generating a clean config file with `config init` (e.g., UseTrash field is written as 'trash' in YAML).
type Config struct {
	Recursive       bool     `mapstructure:"recursive" yaml:"recursive"`
	DryRun          bool     `mapstructure:"dry-run" yaml:"dry-run"`
	Force           bool     `mapstructure:"force" yaml:"force"`
	Verbose         bool     `mapstructure:"verbose" yaml:"verbose"`
	Quiet           bool     `mapstructure:"quiet" yaml:"quiet"`
	UseTrash        bool     `mapstructure:"trash" yaml:"trash"`
	OutputFormat    string   `mapstructure:"output-format" yaml:"output-format"`
	LogFile         string   `mapstructure:"log-file" yaml:"log-file"`
	IgnoreFiles     []string `mapstructure:"ignore-files" yaml:"ignore-files"`
	ExcludeDirs     []string `mapstructure:"exclude-dirs" yaml:"exclude-dirs"`
	ExcludePattern  string   `mapstructure:"exclude-pattern" yaml:"exclude-pattern"`
	ExcludeGlob     string   `mapstructure:"exclude-glob" yaml:"exclude-glob"`
	ExcludeGlobPath string   `mapstructure:"exclude-glob-path" yaml:"exclude-glob-path"`
	OlderThanStr    string   `mapstructure:"older-than" yaml:"older-than"`
	FilesOverStr    string   `mapstructure:"files-over" yaml:"files-over"`
	TopN            int      `mapstructure:"top-n" yaml:"top-n"`
	DuplicateKeep   string   `mapstructure:"duplicate-keep" yaml:"duplicate-keep"`
	FindDuplicates  bool     `mapstructure:"find-duplicates" yaml:"find-duplicates"`
	HashAlgo        string   `mapstructure:"hash-algo" yaml:"hash-algo"`
	SortBy          string   `mapstructure:"sort-by" yaml:"sort-by"`
	ConfigFile      string   `mapstructure:"-" yaml:"-"` // This field is for internal use and should not be saved to or read from the config file.
}

// Global instance of the Config struct, accessible throughout the application.
var config Config

// fileResult is a helper struct to hold a file's path and its os.FileInfo.
// This is used to avoid repeated os.Stat calls when sorting results.
type fileResult struct {
	Path string
	Info os.FileInfo
}

// --- Global Variables ---
// These variables are used across different parts of the application.
var (
	errorList      []string    // A slice to collect all non-fatal errors encountered during execution.
	errorMutex     sync.Mutex  // A mutex to protect concurrent access to the errorList slice from multiple goroutines.
	logger         *log.Logger // The global logger instance, configured based on --quiet and --log-file flags.
	configFileUsed string      // Holds the path of the config file that was loaded, for display to the user.
)

// --- Main Command Structure ---
// rootCmd is the root command for the entire application, managed by the Cobra library.
// All other commands are subcommands of this one.
var rootCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "A powerful utility to find and manage files and folders.",
	Long: `Cleanup is a command-line tool that helps you keep your filesystem tidy.
It can find and delete empty folders, find duplicate files, and identify large directories.`,
	// PersistentPreRunE runs before any command's main execution function (RunE).
	// It's used for setup tasks common to all subcommands.
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Initialize configuration from file, environment variables, and flags.
		if err := initConfig(cmd); err != nil {
			return err
		}
		// Set up the global logger based on the configuration.
		if err := configureLogger(); err != nil {
			// If the log file can't be opened, warn the user but don't crash.
			fmt.Fprintf(os.Stderr, "⚠️  Could not open log file: %v. Logging to stderr only.\n", err)
		}
		return nil
	},
	// PersistentPostRun runs after any command's main execution function.
	// It's used for teardown or summary tasks.
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// Print a summary of all non-fatal errors that occurred during the run.
		printErrorSummary()
	},
}

// init is a special Go function that runs once when the package is initialized.
// It's used here to define all the commands and their flags using Cobra.
func init() {
	// --- Persistent Flags ---
	// These flags are available on the root command and all of its subcommands.
	rootCmd.PersistentFlags().StringVar(&config.ConfigFile, "config", "", "Config file (default is ./.cleanup.yaml or $HOME/.cleanup.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&config.Verbose, "verbose", "v", false, "Enable verbose output.")
	rootCmd.PersistentFlags().BoolVarP(&config.Quiet, "quiet", "q", false, "Suppress all output except for errors.")
	rootCmd.PersistentFlags().StringVarP(&config.LogFile, "log-file", "l", "", "Path to a file to write logs to.")
	rootCmd.PersistentFlags().StringVarP(&config.OutputFormat, "output", "o", "", "Output results in a structured format: json|csv")
	rootCmd.PersistentFlags().StringVarP(&config.ExcludePattern, "exclude-pattern", "p", "", "Exclude paths matching regex pattern.")
	rootCmd.PersistentFlags().StringVarP(&config.ExcludeGlob, "exclude-glob", "g", "", "Exclude paths matching glob pattern on file/dir name.")
	rootCmd.PersistentFlags().StringVar(&config.ExcludeGlobPath, "exclude-glob-path", "", "Exclude paths matching glob pattern on the full path.")
	rootCmd.PersistentFlags().StringSliceVarP(&config.ExcludeDirs, "exclude-dirs", "x", []string{}, "Comma-separated list of directories to exclude by name.")

	// --- Add Subcommands ---
	// Each subcommand is initialized in its own function for better organization and clarity.
	addEmptyCmd()
	addFindCmd()
	addLargeCmd()
	addConfigCmd()
	addVersionCmd()
}

// main is the entry point of the application.
func main() {
	// Create a context that listens for interrupt (Ctrl+C) and termination signals.
	// This allows for graceful shutdown of long-running operations like file scanning.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop() // Ensure the signal notification is cleaned up when main exits.

	// Execute the root command with the cancellable context.
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		// If the error is due to cancellation, print a user-friendly message.
		if !errors.Is(err, context.Canceled) {
			// For other errors, Cobra typically prints them, so we just exit with a non-zero code.
			os.Exit(1)
		}
		fmt.Println("\n🚫 Operation cancelled by user.")
	}
}

// --- Command Definitions ---

// addEmptyCmd sets up the 'empty' subcommand and its specific flags.
func addEmptyCmd() {
	cmd := &cobra.Command{
		Use:   "empty [PATH]",
		Short: "Find and delete empty folders",
		Args:  cobra.MaximumNArgs(1), // Accepts zero or one argument (the path).
		RunE: func(cmd *cobra.Command, args []string) error {
			// The main execution logic for this command is in the runEmpty function.
			return runEmpty(cmd.Context(), args)
		},
	}
	// Flags specific to the 'empty' command.
	cmd.Flags().BoolVarP(&config.Recursive, "recursive", "r", false, "Search recursively.")
	cmd.Flags().BoolVarP(&config.DryRun, "dry-run", "d", false, "Perform a trial run without making any changes.")
	cmd.Flags().BoolVarP(&config.Force, "force", "f", false, "Skip confirmation prompts.")
	cmd.Flags().BoolVarP(&config.UseTrash, "trash", "t", false, "Move to system trash instead of deleting permanently.")
	cmd.Flags().StringSliceVarP(&config.IgnoreFiles, "ignore-files", "i", []string{}, "Files to ignore when determining if a folder is empty.")
	cmd.Flags().StringVarP(&config.OlderThanStr, "older-than", "O", "", "Only consider folders older than a duration (e.g., 30d, 4w, 12h, 90m).")
	rootCmd.AddCommand(cmd)
}

// addFindCmd sets up the 'find' subcommand and its specific flags.
func addFindCmd() {
	cmd := &cobra.Command{
		Use:   "find [PATH]",
		Short: "Find files by criteria (duplicates, size, age)",
		Args:  cobra.MaximumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			// This pre-run validation gives early feedback on invalid flag values.
			if !contains([]string{"path", "size", "age"}, config.SortBy) {
				return fmt.Errorf("invalid value for --sort: %q. Allowed values are: [path, size, age]", config.SortBy)
			}
			if !contains([]string{"sha256", "sha1", "md5"}, config.HashAlgo) {
				return fmt.Errorf("invalid value for --hash-algo: %q. Allowed values are: [sha256, sha1, md5]", config.HashAlgo)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return runFind(cmd.Context(), args)
		},
	}
	// Flags specific to the 'find' command.
	cmd.Flags().BoolVarP(&config.DryRun, "dry-run", "d", false, "Perform a trial run without making any changes.")
	cmd.Flags().BoolVarP(&config.Force, "force", "f", false, "Skip confirmation prompts for deletion.")
	cmd.Flags().BoolVarP(&config.UseTrash, "trash", "t", false, "Move files to system trash instead of deleting.")
	cmd.Flags().StringVarP(&config.OlderThanStr, "older-than", "O", "", "Find files older than a duration (e.g., 30d, 4w, 12h, 90m).")
	cmd.Flags().StringVarP(&config.FilesOverStr, "files-over", "S", "", "Find files larger than a size (e.g., 100MB, 2GB).")
	cmd.Flags().BoolVarP(&config.FindDuplicates, "find-duplicates", "D", false, "Find duplicate files by content hash.")
	cmd.Flags().StringVar(&config.DuplicateKeep, "keep", "prompt", "Duplicate handling strategy: prompt, newest, oldest, first (alphabetical).")
	cmd.Flags().StringVar(&config.SortBy, "sort", "path", "Sort results by: path|size|age")
	cmd.Flags().StringVar(&config.HashAlgo, "hash-algo", "sha256", "Hash algorithm for finding duplicates: sha256|sha1|md5")
	rootCmd.AddCommand(cmd)
}

// addLargeCmd sets up the 'large' subcommand.
func addLargeCmd() {
	cmd := &cobra.Command{
		Use:   "large [PATH]",
		Short: "Find the largest folders in a directory",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runLarge(cmd.Context(), args)
		},
	}
	cmd.Flags().IntVarP(&config.TopN, "top", "n", 10, "Number of largest folders to show.")
	rootCmd.AddCommand(cmd)
}

// addConfigCmd sets up the 'config' subcommand for managing the configuration file.
func addConfigCmd() {
	var isGlobal bool
	configCmd := &cobra.Command{Use: "config", Short: "Manage configuration"}
	initCmd := &cobra.Command{
		Use: "init", Short: "Create a default config file",
		Long: `Creates a default .cleanup.yaml configuration file.
By default, it creates the file in the current directory.
Use the --global flag to create it in your home directory instead.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var configDir string
			var err error

			// Allow user to choose config file location with the --global flag.
			if isGlobal {
				configDir, err = os.UserHomeDir()
				if err != nil {
					return err
				}
			} else {
				configDir, err = os.Getwd()
				if err != nil {
					return err
				}
			}

			configFile := filepath.Join(configDir, ".cleanup.yaml")
			// Check if the file already exists to avoid overwriting.
			if _, err := os.Stat(configFile); err == nil {
				return fmt.Errorf("config file already exists at %s", configFile)
			}

			// Define a complete default configuration struct.
			// This ensures all user-configurable fields are written to the file.
			defaultConfig := Config{
				Recursive:       false,
				DryRun:          false,
				Force:           false,
				Verbose:         false,
				Quiet:           false,
				UseTrash:        false,
				OutputFormat:    "",
				LogFile:         "",
				IgnoreFiles:     []string{".DS_Store", "Thumbs.db"},
				ExcludeDirs:     []string{".git", "node_modules", "vendor", "tmp"},
				ExcludePattern:  "",
				ExcludeGlob:     "",
				ExcludeGlobPath: "",
				OlderThanStr:    "",
				FilesOverStr:    "",
				TopN:            10,
				DuplicateKeep:   "prompt",
				FindDuplicates:  false,
				HashAlgo:        "sha256",
				SortBy:          "path",
			}

			// Marshal the struct into YAML format.
			yamlData, err := yaml.Marshal(&defaultConfig)
			if err != nil {
				return err
			}
			// Write the YAML data to the file.
			if err = os.WriteFile(configFile, yamlData, 0644); err != nil {
				return err
			}
			logInfo("✅ Default config file created at %s", configFile)
			return nil
		},
	}
	initCmd.Flags().BoolVar(&isGlobal, "global", false, "Create the config file in the user's home directory")
	configCmd.AddCommand(initCmd)
	rootCmd.AddCommand(configCmd)
}

// addVersionCmd sets up the 'version' subcommand.
func addVersionCmd() {
	versionCmd := &cobra.Command{
		Use: "version", Short: "Print the version number and build info",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("🧹 %s\n", appName)
			fmt.Printf("   Version: %s (%s)\n", version, date)
			fmt.Printf("   Go Version: %s\n", runtime.Version())
			fmt.Printf("   Source Code: %s\n", sourceCode)
			fmt.Printf("   %s\n", copyright)
		},
	}
	rootCmd.AddCommand(versionCmd)
}

// --- Command Execution Logic ---

// runEmpty contains the core logic for the 'empty' command.
func runEmpty(ctx context.Context, args []string) error {
	runCtx, err := newRunContext()
	if err != nil {
		return err
	}

	targetDir, err := getTargetDir(args)
	if err != nil {
		return err
	}
	printEmptyModeSummary(targetDir)

	var emptyDirs []string
	if config.Recursive {
		emptyDirs, err = findEmptyRecursive(ctx, targetDir, runCtx)
	} else {
		emptyDirs, err = findEmptyTopLevel(ctx, targetDir, runCtx)
	}
	if err != nil {
		addError(fmt.Errorf("error during scan: %w", err))
		return err
	}
	if ctx.Err() != nil {
		return ctx.Err()
	}

	if len(emptyDirs) > 0 {
		logInfo("\n🔎 Found %d empty folders to process:", len(emptyDirs))
		var outputData []map[string]interface{}
		for _, dir := range emptyDirs {
			outputData = append(outputData, map[string]interface{}{"path": dir})
		}
		outputResults(outputData, []string{"path"})
		handleDeletion("empty folders", emptyDirs, 0)
	} else {
		logInfo("\n🎉 Success! No empty folders were found.")
	}
	return nil
}

// runFind contains the core logic for the 'find' command.
func runFind(ctx context.Context, args []string) error {
	runCtx, err := newRunContext()
	if err != nil {
		return err
	}

	targetDir, err := getTargetDir(args)
	if err != nil {
		return err
	}
	printFindModeSummary(targetDir)

	if config.FindDuplicates {
		return findDuplicates(ctx, targetDir, runCtx)
	}
	return findFilesByCriteria(ctx, targetDir, runCtx)
}

// runLarge contains the core logic for the 'large' command.
func runLarge(ctx context.Context, args []string) error {
	runCtx, err := newRunContext()
	if err != nil {
		return err
	}

	targetDir, err := getTargetDir(args)
	if err != nil {
		return err
	}
	printLargeModeSummary(targetDir)
	logInfo("⏳ Scanning... this may take a while. Press Ctrl+C to cancel.")

	dirSizes, err := calculateDirectorySizes(ctx, targetDir, runCtx)
	if err != nil {
		return err
	}
	if ctx.Err() != nil {
		return ctx.Err()
	}

	type dirInfo struct {
		Path string `json:"path"`
		Size int64  `json:"size"`
	}
	var sortedDirs []dirInfo
	for path, size := range dirSizes {
		sortedDirs = append(sortedDirs, dirInfo{path, size})
	}

	sort.Slice(sortedDirs, func(i, j int) bool { return sortedDirs[i].Size > sortedDirs[j].Size })

	limit := config.TopN
	if len(sortedDirs) < limit {
		limit = len(sortedDirs)
	}
	results := sortedDirs[:limit]
	logInfo("\n🔎 Top %d largest folders:", limit)

	var outputData []map[string]interface{}
	for _, dir := range results {
		outputData = append(outputData, map[string]interface{}{"path": dir.Path, "size": dir.Size, "size_formatted": formatBytes(dir.Size)})
	}
	outputResults(outputData, []string{"path", "size", "size_formatted"})
	return nil
}

// --- Core Logic ---

// findFilesByCriteria scans for files based on size and age filters.
func findFilesByCriteria(ctx context.Context, targetDir string, runCtx *runContext) error {
	var foundFiles []fileResult
	var mu sync.Mutex

	processFile := func(path string, info os.FileInfo) {
		if runCtx.shouldExclude(path) || !info.Mode().IsRegular() {
			return
		}

		matchSize := runCtx.filesOverBytes > 0 && info.Size() >= runCtx.filesOverBytes
		matchAge := !runCtx.olderThan.IsZero() && info.ModTime().Before(runCtx.olderThan)
		noCriteria := runCtx.filesOverBytes == 0 && runCtx.olderThan.IsZero()

		if matchSize || matchAge || noCriteria {
			mu.Lock()
			foundFiles = append(foundFiles, fileResult{Path: path, Info: info})
			mu.Unlock()
		}
	}

	if err := scanFilesParallel(ctx, targetDir, processFile); err != nil {
		return err
	}
	sortResults(foundFiles)
	logInfo("\n🔎 Found %d files matching criteria.", len(foundFiles))

	var pathsToDelete []string
	var totalSize int64
	var outputData []map[string]interface{}
	for _, file := range foundFiles {
		pathsToDelete = append(pathsToDelete, file.Path)
		totalSize += file.Info.Size()
		outputData = append(outputData, map[string]interface{}{"path": file.Path, "size": file.Info.Size(), "modified": file.Info.ModTime().Format(time.RFC3339)})
	}

	outputResults(outputData, []string{"path", "size", "modified"})
	handleDeletion("matching files", pathsToDelete, totalSize)
	return nil
}

// findDuplicates scans for files with identical content hashes.
func findDuplicates(ctx context.Context, targetDir string, runCtx *runContext) error {
	hashes := make(map[string][]string)
	var mu sync.Mutex

	processFile := func(path string, info os.FileInfo) {
		if runCtx.shouldExclude(path) || info.Size() == 0 || !info.Mode().IsRegular() {
			return
		}
		hash, err := hashFile(path)
		if err != nil {
			addError(err)
			return
		}
		logVerbose("Hashed %s (%s): %s", path, config.HashAlgo, hash)
		mu.Lock()
		hashes[hash] = append(hashes[hash], path)
		mu.Unlock()
	}

	if err := scanFilesParallel(ctx, targetDir, processFile); err != nil {
		return err
	}

	var duplicatesToProcess [][]string
	for hash, files := range hashes {
		if len(files) > 1 {
			logVerbose("Found duplicate set for hash %s: %v", hash, files)
			duplicatesToProcess = append(duplicatesToProcess, files)
		}
	}

	logInfo("\n👯 Found %d sets of duplicate files.", len(duplicatesToProcess))
	if len(duplicatesToProcess) == 0 {
		return nil
	}

	var pathsToDelete []string
	var totalSizeDeleted int64
	for i, set := range duplicatesToProcess {
		toDelete, err := processDuplicateSet(set, i+1)
		if err != nil {
			addError(err)
			continue
		}
		pathsToDelete = append(pathsToDelete, toDelete...)
		for _, p := range toDelete {
			totalSizeDeleted += getFileSize(p)
		}
	}
	handleDeletion("duplicate files", pathsToDelete, totalSizeDeleted)
	return nil
}

// calculateDirectorySizes walks the filesystem and sums up file sizes.
func calculateDirectorySizes(ctx context.Context, targetDir string, runCtx *runContext) (map[string]int64, error) {
	dirSizes := make(map[string]int64)
	var mu sync.Mutex

	processFile := func(path string, info os.FileInfo) {
		if runCtx.shouldExclude(path) {
			return
		}
		size := info.Size()
		for p := filepath.Dir(path); strings.HasPrefix(p, targetDir) || p == targetDir; p = filepath.Dir(p) {
			mu.Lock()
			dirSizes[p] += size
			mu.Unlock()
		}
	}
	err := scanFilesParallel(ctx, targetDir, processFile)
	if err != nil {
		addError(fmt.Errorf("directory size calculation failed: %w", err))
	}
	return dirSizes, err
}

// findEmptyRecursive finds all empty folders in a directory tree.
func findEmptyRecursive(ctx context.Context, path string, runCtx *runContext) ([]string, error) {
	var allDirs []string
	deletablePaths := make(map[string]bool)
	ignoreFileSet := stringSliceToSet(config.IgnoreFiles)

	logVerbose("Phase 1: Walking filesystem to collect all directories...")
	walkErr := filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		if err != nil {
			addError(fmt.Errorf("cannot access path %s: %w", p, err))
			return nil
		}
		if runCtx.shouldExclude(p) {
			if d.IsDir() {
				logVerbose("Skipping excluded directory: %s", p)
				return filepath.SkipDir
			}
			return nil
		}
		if d.IsDir() {
			allDirs = append(allDirs, p)
		}
		return nil
	})
	if walkErr != nil {
		return nil, walkErr
	}

	sort.Slice(allDirs, func(i, j int) bool {
		return len(strings.Split(allDirs[i], string(os.PathSeparator))) > len(strings.Split(allDirs[j], string(os.PathSeparator)))
	})

	logVerbose("Phase 2: Evaluating %d directories from deepest to shallowest...", len(allDirs))
	for _, dir := range allDirs {
		if dir == path {
			continue
		}
		isDirEmpty, err := isDirectoryEmpty(dir, ignoreFileSet, deletablePaths, runCtx)
		if err != nil {
			addError(err)
			continue
		}
		if isDirEmpty {
			logVerbose("    ✅ Marked as empty: %s", dir)
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

// findEmptyTopLevel finds empty folders only in the top level of the given path.
func findEmptyTopLevel(ctx context.Context, path string, runCtx *runContext) ([]string, error) {
	var emptyDirs []string
	ignoreFileSet := stringSliceToSet(config.IgnoreFiles)

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		if !entry.IsDir() {
			continue
		}

		fullPath := filepath.Join(path, entry.Name())
		logVerbose("-> Evaluating top-level directory: %s", fullPath)
		if runCtx.shouldExclude(fullPath) {
			continue
		}

		isDirEmpty, err := isDirectoryEmpty(fullPath, ignoreFileSet, nil, runCtx)
		if err != nil {
			addError(fmt.Errorf("could not evaluate dir %s: %w", fullPath, err))
			continue
		}

		if isDirEmpty {
			emptyDirs = append(emptyDirs, fullPath)
		}
	}
	sort.Strings(emptyDirs)
	return emptyDirs, nil
}

// --- Parallel Scanner ---
func scanFilesParallel(ctx context.Context, targetDir string, processFunc func(string, os.FileInfo)) error {
	var fileCount int64
	if !config.Quiet {
		logVerbose("Pre-scanning to count files for progress bar...")
		_ = filepath.WalkDir(targetDir, func(path string, d os.DirEntry, err error) error {
			if err == nil && !d.IsDir() {
				fileCount++
			}
			return nil
		})
	}

	bar := progressbar.NewOptions64(fileCount,
		progressbar.OptionSetDescription("Scanning files..."),
		progressbar.OptionSetWriter(os.Stderr), progressbar.OptionShowBytes(false),
		progressbar.OptionShowCount(), progressbar.OptionSetWidth(30),
		progressbar.OptionOnCompletion(func() {
			if !config.Quiet && !config.Verbose {
				fmt.Fprint(os.Stderr, "\n")
			}
		}),
		progressbar.OptionSpinnerType(14), progressbar.OptionFullWidth(),
		progressbar.OptionSetRenderBlankState(true),
		progressbar.OptionSetVisibility(!config.Quiet && !config.Verbose),
	)

	paths := make(chan string, 2*runtime.NumCPU())
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Walker Goroutine
	go func() {
		defer close(paths)
		_ = filepath.WalkDir(targetDir, func(path string, d os.DirEntry, err error) error {
			if ctx.Err() != nil {
				return ctx.Err()
			}
			if err != nil {
				addError(fmt.Errorf("access error on %s: %w", path, err))
				return nil
			}
			if !d.IsDir() {
				select {
				case paths <- path:
				case <-ctx.Done():
					return ctx.Err()
				}
			}
			return nil
		})
	}()

	// Worker Goroutines
	numWorkers := runtime.NumCPU()
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case path, ok := <-paths:
					if !ok {
						return
					}
					info, err := os.Lstat(path)
					if err != nil {
						addError(err)
					} else {
						processFunc(path, info)
					}
					_ = bar.Add(1)
				case <-ctx.Done():
					return
				}
			}
		}()
	}
	wg.Wait()
	return ctx.Err()
}

// --- Helper & Utility Functions ---

// handleDeletion manages the user confirmation and deletion process.
func handleDeletion(itemType string, paths []string, totalSize int64) {
	if len(paths) == 0 {
		return
	}

	if config.DryRun {
		logInfo("\n--- 🧪 Dry Run Summary ---")
		logInfo("Would have %s %d %s.", getActionStringPast(), len(paths), itemType)
		if totalSize > 0 {
			logInfo("Total size that would be freed: %s", formatBytes(totalSize))
		}
		for _, p := range paths {
			logInfo("  - %s", p)
		}
		logInfo("--------------------------")
		logInfo("No changes were made.")
		return
	}

	if !config.Force {
		action := getActionString()
		logInfo("\n\033[33m--- WARNING: About to %s %d %s ---", action, len(paths), itemType)
		if !config.UseTrash {
			logInfo("\033[31mThis action is PERMANENT and CANNOT be undone.\033[0m")
		}
		fmt.Printf("\033[34m❓ Are you sure you want to proceed? [yes/No] \033[0m")
		reader := bufio.NewReader(os.Stdin)
		response, _ := reader.ReadString('\n')
		if strings.ToLower(strings.TrimSpace(response)) != "yes" && strings.ToLower(strings.TrimSpace(response)) != "y" {
			logInfo("\n👍 OK. No changes were made.")
			return
		}
	}

	logInfo("\n🔥 Processing...")
	processedCount := 0
	bar := progressbar.NewOptions(len(paths),
		progressbar.OptionSetDescription(fmt.Sprintf("%s items...", getActionStringPresent())),
		progressbar.OptionSetVisibility(!config.Quiet && !config.Verbose),
	)

	for _, path := range paths {
		var opErr error
		if config.UseTrash {
			opErr = trash.Throw(path)
		} else {
			opErr = os.RemoveAll(path)
		}
		if opErr != nil {
			addError(fmt.Errorf("error %s %s: %w", getActionStringPast(), path, opErr))
		} else {
			logVerbose("  %s %s: %s", getActionIcon(), getActionStringPast(), path)
			processedCount++
		}
		_ = bar.Add(1)
	}
	logInfo("\n✨ All done! %s %d %s.", getActionStringPast(), processedCount, itemType)
}

// initConfig reads configuration from file, env vars, and flags, establishing a clear precedence.
func initConfig(cmd *cobra.Command) error {
	v := viper.New()

	// This logic handles the precedence: Flags > Env > Config File.
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.AutomaticEnv()

	// Bind all flags from the current command and all parent commands to Viper.
	// This is the crucial step that makes flags override all other settings.
	if err := v.BindPFlags(cmd.Flags()); err != nil {
		return err
	}
	if err := v.BindPFlags(cmd.Root().PersistentFlags()); err != nil {
		return err
	}

	// Tell Viper where to look for a config file.
	// If the --config flag was used, use that path.
	if specifiedConfigFile := v.GetString("config"); specifiedConfigFile != "" {
		v.SetConfigFile(specifiedConfigFile)
	} else {
		// Otherwise, search in standard locations.
		v.SetConfigName(".cleanup")
		v.SetConfigType("yaml")
		v.AddConfigPath(".")
		if home, err := os.UserHomeDir(); err == nil {
			v.AddConfigPath(home)
		}
	}

	// Attempt to read the config file. It's okay if it doesn't exist.
	if err := v.ReadInConfig(); err == nil {
		configFileUsed = v.ConfigFileUsed()
	} else if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		// Only return an error if a specific file was requested but could not be read.
		if v.GetString("config") != "" {
			return err
		}
	}

	// Finally, unmarshal all the resolved values (from flags, env, or config) into our struct.
	return v.Unmarshal(&config)
}

// getTargetDir resolves the input argument to an absolute, validated directory path.
func getTargetDir(args []string) (string, error) {
	pathArg := "."
	if len(args) > 0 {
		pathArg = args[0]
	}
	targetDir, err := filepath.Abs(pathArg)
	if err != nil {
		return "", fmt.Errorf("could not resolve path '%s': %w", pathArg, err)
	}
	info, err := os.Stat(targetDir)
	if os.IsNotExist(err) || !info.IsDir() {
		return "", fmt.Errorf("the specified path does not exist or is not a directory: %s", targetDir)
	}
	return targetDir, nil
}

// parseDuration converts a human-readable duration string (e.g., "30d", "4w") into a time.Duration.
func parseDuration(s string) (time.Duration, error) {
	s = strings.ToLower(strings.TrimSpace(s))
	if s == "" {
		return 0, errors.New("duration string cannot be empty")
	}

	re := regexp.MustCompile(`^(\d+)\s*(w|d|h|m|s)$`)
	matches := re.FindStringSubmatch(s)
	if len(matches) != 3 {
		return 0, fmt.Errorf("invalid duration format: %q. Use formats like '30d', '4w', '12h'", s)
	}

	num, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, err
	}

	unit := matches[2]
	var multiplier time.Duration
	switch unit {
	case "w":
		multiplier = 7 * 24 * time.Hour
	case "d":
		multiplier = 24 * time.Hour
	case "h":
		multiplier = time.Hour
	case "m":
		multiplier = time.Minute
	case "s":
		multiplier = time.Second
	}
	return time.Duration(num) * multiplier, nil
}

// parseSize converts a human-readable size string (e.g., "1.5GB") into bytes.
func parseSize(s string) (int64, error) {
	s = strings.ToUpper(strings.TrimSpace(s))
	if s == "" {
		return 0, errors.New("size string cannot be empty")
	}

	re := regexp.MustCompile(`^([\d.]+)\s*(G|M|K)?B?$`)
	matches := re.FindStringSubmatch(s)
	if len(matches) < 2 {
		return 0, fmt.Errorf("invalid size format: %q", s)
	}

	num, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number in size string: %q", matches[1])
	}

	var multiplier float64 = 1
	if len(matches) > 2 {
		switch matches[2] {
		case "G":
			multiplier = 1024 * 1024 * 1024
		case "M":
			multiplier = 1024 * 1024
		case "K":
			multiplier = 1024
		}
	}
	return int64(num * multiplier), nil
}

// formatBytes converts a byte count into a human-readable string.
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

// hashFile computes the hash of a file's content using the configured algorithm.
func hashFile(path string) (string, error) {
	var h hash.Hash
	switch config.HashAlgo {
	case "md5":
		h = md5.New()
	case "sha1":
		h = sha1.New()
	default:
		h = sha256.New()
	}

	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("could not open file %s for hashing: %w", path, err)
	}
	defer file.Close()

	if _, err := io.Copy(h, file); err != nil {
		return "", fmt.Errorf("could not copy file content for hashing %s: %w", path, err)
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// addError safely adds an error to the global error list and logs it.
func addError(err error) {
	if err == nil {
		return
	}
	errorMutex.Lock()
	defer errorMutex.Unlock()
	errorList = append(errorList, err.Error())
	logError(err.Error())
}

// printErrorSummary prints all collected errors at the end of execution.
func printErrorSummary() {
	errorMutex.Lock()
	defer errorMutex.Unlock()
	if len(errorList) > 1 {
		logInfo("\n--- ⚠️ Encountered %d Error(s) Summary ---", len(errorList))
		for i, e := range errorList {
			logInfo("%d: %s", i+1, e)
		}
		logInfo("---------------------------------------")
	}
}

// outputResults handles writing results to console or a file in the specified format.
func outputResults(data interface{}, headers []string) {
	if config.OutputFormat == "" {
		switch v := data.(type) {
		case []map[string]interface{}:
			for _, row := range v {
				if path, ok := row["path"]; ok {
					line := fmt.Sprintf("  • %s", path)
					if size, ok := row["size_formatted"]; ok {
						line += fmt.Sprintf(" (%s)", size)
					}
					logInfo("%s", line)
				}
			}
		}
		return
	}

	var writer io.Writer = os.Stdout
	if config.Quiet {
		if config.LogFile == "" {
			return
		}
		file, err := os.OpenFile(config.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			addError(fmt.Errorf("failed to open log file for results: %w", err))
			return
		}
		defer file.Close()
		writer = file
	}

	switch config.OutputFormat {
	case "json":
		outputJSON(data, writer)
	case "csv":
		outputCSV(data, headers, writer)
	}
}

func outputJSON(data interface{}, writer io.Writer) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		addError(fmt.Errorf("failed to generate JSON output: %w", err))
		return
	}
	fmt.Fprintln(writer, string(jsonData))
}

func outputCSV(data interface{}, headers []string, writer io.Writer) {
	csvWriter := csv.NewWriter(writer)
	defer csvWriter.Flush()

	if err := csvWriter.Write(headers); err != nil {
		addError(fmt.Errorf("failed to write CSV header: %w", err))
		return
	}
	switch v := data.(type) {
	case []map[string]interface{}:
		for _, rowMap := range v {
			var record []string
			for _, h := range headers {
				if val, ok := rowMap[h]; ok {
					record = append(record, fmt.Sprintf("%v", val))
				} else {
					record = append(record, "")
				}
			}
			if err := csvWriter.Write(record); err != nil {
				addError(fmt.Errorf("failed to write CSV record: %w", err))
			}
		}
	default:
		addError(fmt.Errorf("unsupported data type for CSV output: %T", data))
	}
}

// configureLogger sets up the global logger based on config flags.
func configureLogger() error {
	var writers []io.Writer
	if !config.Quiet {
		writers = append(writers, os.Stderr)
	}
	if config.LogFile != "" {
		file, err := os.OpenFile(config.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return err
		}
		writers = append(writers, file)
	}

	out := io.Discard
	if len(writers) > 0 {
		out = io.MultiWriter(writers...)
	}
	logger = log.New(out, "", 0)
	return nil
}

// --- Logging Wrappers ---
func logInfo(format string, v ...interface{}) {
	if logger != nil {
		logger.Printf(format, v...)
	}
}
func logVerbose(format string, v ...interface{}) {
	if config.Verbose && logger != nil {
		logger.Printf("[VERBOSE] "+format, v...)
	}
}
func logError(format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, "\033[31m[ERROR] "+format+"\033[0m\n", v...)
}

// isDirectoryEmpty checks if a directory contains no files and no non-deletable subdirectories.
func isDirectoryEmpty(dir string, ignoreFileSet map[string]struct{}, deletableSubDirs map[string]bool, runCtx *runContext) (bool, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false, err
	}
	if !runCtx.olderThan.IsZero() {
		info, err := os.Stat(dir)
		if err != nil {
			return false, err
		}
		if info.ModTime().After(runCtx.olderThan) {
			logVerbose("    - Dir '%s' is too new, skipping check.", dir)
			return false, nil
		}
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			if _, ignored := ignoreFileSet[entry.Name()]; !ignored {
				logVerbose("    - Contains non-ignored file: %s. Marking as NOT empty.", entry.Name())
				return false, nil
			}
		} else {
			if deletableSubDirs == nil {
				return false, nil
			}
			subDirPath := filepath.Join(dir, entry.Name())
			if !deletableSubDirs[subDirPath] {
				logVerbose("    - Contains non-empty subdirectory: %s. Marking as NOT empty.", subDirPath)
				return false, nil
			}
		}
	}
	return true, nil
}

// processDuplicateSet applies the chosen strategy to a set of duplicate files.
func processDuplicateSet(set []string, setIndex int) ([]string, error) {
	if len(set) < 2 {
		return nil, nil
	}
	type fileInfo struct {
		Path    string
		ModTime time.Time
	}
	var files []fileInfo
	for _, p := range set {
		info, err := os.Stat(p)
		if err != nil {
			return nil, fmt.Errorf("could not stat file %s: %w", p, err)
		}
		files = append(files, fileInfo{Path: p, ModTime: info.ModTime()})
	}

	var fileToKeep string
	var filesToDelete []string

	strategy := config.DuplicateKeep
	if config.Force && strategy == "prompt" {
		strategy = "first"
	}

	switch strings.ToLower(strategy) {
	case "newest":
		sort.Slice(files, func(i, j int) bool { return files[i].ModTime.After(files[j].ModTime) })
		fileToKeep = files[0].Path
	case "oldest":
		sort.Slice(files, func(i, j int) bool { return files[i].ModTime.Before(files[j].ModTime) })
		fileToKeep = files[0].Path
	case "first":
		sort.Slice(files, func(i, j int) bool { return files[i].Path < files[j].Path })
		fileToKeep = files[0].Path
	case "prompt":
		fmt.Println()
		logInfo("--- Set %d (%s) ---", setIndex, formatBytes(getFileSize(files[0].Path)))
		for i, f := range files {
			logInfo("  [%d] %s (%s ago)", i+1, f.Path, time.Since(f.ModTime).Round(time.Second))
		}
		fmt.Printf("\033[34m❓ For set %d, enter the number of the file to KEEP [1-%d], or 's' to skip: \033[0m", setIndex, len(files))
		reader := bufio.NewReader(os.Stdin)
		for {
			response, _ := reader.ReadString('\n')
			response = strings.TrimSpace(response)
			if strings.ToLower(response) == "s" {
				return nil, nil
			}
			choice, err := strconv.Atoi(response)
			if err == nil && choice >= 1 && choice <= len(files) {
				fileToKeep = files[choice-1].Path
				break
			}
			fmt.Printf("\033[31mInvalid input. Please enter a number between 1 and %d, or 's' to skip: \033[0m", len(files))
		}
	default:
		return nil, fmt.Errorf("unknown duplicate keep strategy: '%s'", config.DuplicateKeep)
	}

	for _, f := range files {
		if f.Path != fileToKeep {
			filesToDelete = append(filesToDelete, f.Path)
		}
	}
	if fileToKeep != "" {
		logVerbose("  -> For set %d, keeping: %s", setIndex, fileToKeep)
	}
	return filesToDelete, nil
}

func getFileSize(path string) int64 {
	info, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return info.Size()
}

// --- String Helpers & Sorters ---
func getActionString() string {
	if config.UseTrash {
		return "move to trash"
	}
	return "delete"
}
func getActionStringPresent() string {
	if config.UseTrash {
		return "Moving to trash"
	}
	return "Deleting"
}
func getActionStringPast() string {
	if config.UseTrash {
		return "Moved to trash"
	}
	return "Deleted"
}
func getActionIcon() string {
	if config.UseTrash {
		return "♻️"
	}
	return "🗑️"
}

func stringSliceToSet(slice []string) map[string]struct{} {
	set := make(map[string]struct{}, len(slice))
	for _, item := range slice {
		set[item] = struct{}{}
	}
	return set
}

func sortResults(files []fileResult) {
	switch strings.ToLower(config.SortBy) {
	case "size":
		sort.Slice(files, func(i, j int) bool { return files[i].Info.Size() > files[j].Info.Size() })
	case "age":
		sort.Slice(files, func(i, j int) bool { return files[i].Info.ModTime().Before(files[j].Info.ModTime()) })
	default:
		sort.Slice(files, func(i, j int) bool { return files[i].Path < files[j].Path })
	}
}

// --- Summary Printers ---
func printEmptyModeSummary(targetDir string) {
	logInfo("--- 🧹 Find Empty Folders Mode ---")
	logInfo("🎯 Target Directory: %s", targetDir)
	if config.Recursive {
		logInfo("🌲 Mode: Recursive Scan Enabled")
	} else {
		logInfo("📂 Mode: Top-Level Scan Only")
	}
	printCommonSummary(false)
	if len(config.IgnoreFiles) > 0 {
		logInfo("🙈 Ignoring Files: %s", strings.Join(config.IgnoreFiles, ", "))
	}
	if config.OlderThanStr != "" {
		logInfo("⏳ Age Filter: Only folders older than %s", config.OlderThanStr)
	}
	logInfo("----------------------------------\n")
}

func printFindModeSummary(targetDir string) {
	logInfo("--- 🔎 Find Files Mode ---")
	logInfo("🎯 Target Directory: %s", targetDir)
	if config.FindDuplicates {
		logInfo("🔎 Mode: Find Duplicates by Hash (%s)", config.HashAlgo)
		logInfo("🤔 Keep Strategy: %s", config.DuplicateKeep)
	} else {
		logInfo("📜 Mode: Find by Size/Age")
		if config.FilesOverStr != "" {
			logInfo("📏 Size Filter: Files over %s", config.FilesOverStr)
		}
		if config.OlderThanStr != "" {
			logInfo("⏳ Age Filter: Files older than %s", config.OlderThanStr)
		}
	}
	printCommonSummary(false)
	logInfo("----------------------------------\n")
}

func printLargeModeSummary(targetDir string) {
	logInfo("--- 📊 Find Large Folders Mode ---")
	logInfo("🎯 Target Directory: %s", targetDir)
	logInfo("📈 Showing Top: %d folders", config.TopN)
	printCommonSummary(true)
	logInfo("----------------------------------\n")
}

func printCommonSummary(isReadOnly bool) {
	if configFileUsed != "" {
		logInfo("⚙️ Using Config: %s", configFileUsed)
	}
	if !isReadOnly {
		if config.DryRun {
			logInfo("🧪 Action: Dry Run")
		} else if config.UseTrash {
			logInfo("♻️ Action: Move to System Trash")
		} else {
			logInfo("🗑️ Action: Permanent Deletion")
		}
		if config.Force {
			logInfo("❗ Confirmation: Skipped (--force enabled)")
		}
	}
	if len(config.ExcludeDirs) > 0 {
		logInfo("🚫 Excluding Dirs by Name: %s", strings.Join(config.ExcludeDirs, ", "))
	}
	if config.ExcludeGlob != "" {
		logInfo("🚫 Excluding by Glob (Name): %s", config.ExcludeGlob)
	}
	if config.ExcludeGlobPath != "" {
		logInfo("🚫 Excluding by Glob (Path): %s", config.ExcludeGlobPath)
	}
	if config.ExcludePattern != "" {
		logInfo("🚫 Excluding by Regex: %s", config.ExcludePattern)
	}
}

// --- State Isolation for Command Runs ---
// runContext isolates state for a single command execution to prevent conflicts.
type runContext struct {
	olderThan      time.Time
	filesOverBytes int64
	excludeRegex   *regexp.Regexp
	excludeDirSet  map[string]struct{}
}

// newRunContext creates a new isolated context from the global config.
func newRunContext() (*runContext, error) {
	ctx := &runContext{
		excludeDirSet: stringSliceToSet(config.ExcludeDirs),
	}
	var err error

	if config.OlderThanStr != "" {
		dur, err := parseDuration(config.OlderThanStr)
		if err != nil {
			return nil, fmt.Errorf("invalid duration for --older-than: %w", err)
		}
		ctx.olderThan = time.Now().Add(-dur)
	}

	if config.FilesOverStr != "" {
		ctx.filesOverBytes, err = parseSize(config.FilesOverStr)
		if err != nil {
			return nil, fmt.Errorf("invalid size for --files-over: %w", err)
		}
	}

	if config.ExcludePattern != "" {
		ctx.excludeRegex, err = regexp.Compile(config.ExcludePattern)
		if err != nil {
			return nil, fmt.Errorf("invalid regex pattern for --exclude-pattern: %w", err)
		}
	}

	return ctx, nil
}

// shouldExclude checks if a path should be skipped based on the isolated run context.
func (rc *runContext) shouldExclude(path string) bool {
	pathForMatching := filepath.ToSlash(path)

	// It checks if any component of the path is in the exclusion set.
	parts := strings.Split(pathForMatching, "/")
	for _, part := range parts {
		if _, isExcluded := rc.excludeDirSet[part]; isExcluded {
			logVerbose("Excluding '%s' because path component '%s' is in exclude list", path, part)
			return true
		}
	}

	if rc.excludeRegex != nil && rc.excludeRegex.MatchString(pathForMatching) {
		logVerbose("Excluding '%s' due to --exclude-pattern", path)
		return true
	}

	baseName := filepath.Base(path)
	if config.ExcludeGlob != "" {
		if matched, err := filepath.Match(config.ExcludeGlob, baseName); err != nil {
			addError(fmt.Errorf("invalid glob pattern for --exclude-glob: %w", err))
		} else if matched {
			logVerbose("Excluding '%s' due to --exclude-glob on name '%s'", path, baseName)
			return true
		}
	}
	if config.ExcludeGlobPath != "" {
		if matched, err := filepath.Match(config.ExcludeGlobPath, pathForMatching); err != nil {
			addError(fmt.Errorf("invalid glob pattern for --exclude-glob-path: %w", err))
		} else if matched {
			logVerbose("Excluding '%s' due to --exclude-glob-path", path)
			return true
		}
	}
	return false
}

// contains is a simple helper function to check for string presence in a slice.
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
