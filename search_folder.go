package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Target network shared folder to search
	targetPath := `\\MICHAEL\ctshippingapp\STICKER\PROCESS`
	reader := bufio.NewReader(os.Stdin)

	// Prompt user to scan a barcode
	fmt.Print("Please scan the keyword: ")
	// Read input from scanner (usually acts as keyboard, sends Enter automatically)
	keyword, _ := reader.ReadString('\n')
	keyword = strings.ToLower(strings.TrimSpace(keyword))

	if keyword == "" {
		fmt.Println("Keyword cannot be empty")
		return
	}

	fmt.Println("Searching, please wait...")

	// Read entries in the top-level directory (one layer only)
	entries, err := os.ReadDir(targetPath)
	if err != nil {
		fmt.Println("Cannot access directory:", err)
		return
	}

	matches := []string{}
	for _, entry := range entries {
		// Check if entry is a directory and matches the keyword
		if entry.IsDir() && strings.Contains(strings.ToLower(entry.Name()), keyword) {
			fullPath := filepath.Join(targetPath, entry.Name())
			matches = append(matches, fullPath)
			fmt.Println("Found:", fullPath) // Print matching folder immediately
		}
	}

	// Print summary
	if len(matches) == 0 {
		fmt.Println("No matching folders found")
	} else {
		fmt.Printf("Search completed, found %d matching folder(s)\n", len(matches))
	}
}
