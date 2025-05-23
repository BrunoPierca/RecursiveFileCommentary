package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const tsNoCheckComment = "// @ts-nocheck // ts-nocheck automatically added by pleasenocheck script"

func main() {
	start := time.Now()
	// Get current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Processing TypeScript files in: %s\n", currentDir)
	
	var processedFiles []string
	var errorFiles []string

	// Walk through all files and directories
	err = filepath.Walk(currentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %s: %v\n", path, err)
			return nil // Continue walking despite errors
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Check if file has .ts or .tsx extension
		ext := strings.ToLower(filepath.Ext(path))
		if ext == ".ts" || ext == ".tsx" {
			if processFile(path) {
				processedFiles = append(processedFiles, path)
			} else {
				errorFiles = append(errorFiles, path)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
		os.Exit(1)
	}

	// Print summary
	fmt.Printf("\n=== Summary ===\n")
	fmt.Printf("Successfully processed: %d files\n", len(processedFiles))

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Finished processing in %v \n\n", elapsed)
	
	if len(processedFiles) > 0 {
		fmt.Println("\nProcessed files:")
		for _, file := range processedFiles {
			relPath, _ := filepath.Rel(currentDir, file)
			fmt.Printf("  ✓ %s\n", relPath)
		}
	}

	if len(errorFiles) > 0 {
		fmt.Printf("\nFailed to process: %d files\n", len(errorFiles))
		for _, file := range errorFiles {
			relPath, _ := filepath.Rel(currentDir, file)
			fmt.Printf("  ✗ %s\n", relPath)
		}
	}

	if len(processedFiles) == 0 && len(errorFiles) == 0 {
		fmt.Println("No TypeScript files (.ts or .tsx) found in the directory.")
	}
}

func processFile(filePath string) bool {
	// Read the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filePath, err)
		return false
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	
	// Read all lines
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %v\n", filePath, err)
		return false
	}

	// Check if @ts-nocheck already exists
	hasNoCheck := false
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == tsNoCheckComment {
			hasNoCheck = true
			break
		}
	}

	// If already has @ts-nocheck, skip
	if hasNoCheck {
		fmt.Printf("Skipping %s (already has @ts-nocheck)\n", filepath.Base(filePath))
		return true
	}

	// Find the right place to insert @ts-nocheck
	insertIndex := 0
	
	// Skip shebang and initial comments/empty lines to find the right insertion point
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		
		// Skip empty lines and comments at the beginning
		if trimmed == "" || strings.HasPrefix(trimmed, "//") || strings.HasPrefix(trimmed, "/*") {
			insertIndex = i + 1
			continue
		}
		
		// If we find actual code, insert before it
		break
	}

	// Insert @ts-nocheck at the determined position
	newLines := make([]string, 0, len(lines)+1)
	
	// Add lines before insertion point
	newLines = append(newLines, lines[:insertIndex]...)
	
	// Add @ts-nocheck comment
	newLines = append(newLines, tsNoCheckComment)
	
	// Add remaining lines
	newLines = append(newLines, lines[insertIndex:]...)

	// Write back to file
	outputFile, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", filePath, err)
		return false
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, line := range newLines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Printf("Error writing to file %s: %v\n", filePath, err)
			return false
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Printf("Error flushing file %s: %v\n", filePath, err)
		return false
	}

	fmt.Printf("Added @ts-nocheck to: %s\n", filepath.Base(filePath))
	return true
}