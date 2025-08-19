package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Keyword cannot be empty")
		return
	}
	keyword := strings.ToLower(os.Args[1])

	targetPath := `\\MICHAEL\ctshippingapp\STICKER\PROCESS`

	entries, err := os.ReadDir(targetPath)
	if err != nil {
		fmt.Println("Cannot access directory:", err)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() && strings.Contains(strings.ToLower(entry.Name()), keyword) {
			firstFolder := filepath.Join(targetPath, entry.Name())
			fmt.Print(firstFolder) // ⚠️ 只输出路径，Python 好解析
			return
		}
	}

	fmt.Print("") // 没找到输出空字符串
}
