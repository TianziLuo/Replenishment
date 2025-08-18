package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	targetPath := `\\MICHAEL\ctshippingapp\STICKER\PROCESS`
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please scan the keyword: ")
	keyword, _ := reader.ReadString('\n')
	keyword = strings.ToLower(strings.TrimSpace(keyword))

	if keyword == "" {
		fmt.Println("Keyword cannot be empty")
		return
	}

	fmt.Println("Searching, please wait...")

	entries, err := os.ReadDir(targetPath)
	if err != nil {
		fmt.Println("Cannot access directory:", err)
		return
	}

	var firstFolder string
	for _, entry := range entries {
		if entry.IsDir() && strings.Contains(strings.ToLower(entry.Name()), keyword) {
			firstFolder = filepath.Join(targetPath, entry.Name())
			fmt.Println("Found folder:", firstFolder)

			// Just record the first matched folder
			break
		}
	}

	if firstFolder == "" {
		fmt.Println("No matching folders found")
		return
	}

	// Call the Python script to process sku.TXT
	pyExe := `C:\Users\monica\AppData\Local\Programs\Python\Python313\python.exe`
	pyScript := `C:\Users\monica\Desktop\Seagull\Replenishment\utils_func\readSKU.py`
	cmd := exec.Command(pyExe, pyScript, firstFolder)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running Python:", err)
	}
	fmt.Println(string(output))
}
