package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func copyDir(src string, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		targetPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(targetPath, info.Mode())
		} else {
			srcFile, err := os.Open(path)
			if err != nil {
				return err
			}
			defer srcFile.Close()

			dstFile, err := os.Create(targetPath)
			if err != nil {
				return err
			}
			defer dstFile.Close()

			_, err = io.Copy(dstFile, srcFile)
			if err != nil {
				return err
			}

			// 保留文件权限
			return os.Chmod(targetPath, info.Mode())
		}
	})
}

func main() {
	targetPath := `\\MICHAEL\ctshippingapp\STICKER\PROCESS`
	reader := os.Stdin
	fmt.Print("Please scan the keyword: ")

	var keyword string
	fmt.Fscanln(reader, &keyword)
	keyword = strings.TrimSpace(keyword)
	if keyword == "" {
		fmt.Println("Keyword cannot be empty")
		return
	}
	keyword = strings.ToLower(keyword)

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
			break
		}
	}

	if firstFolder == "" {
		fmt.Println("No matching folders found")
		return
	}

	// 目标路径
	dstPath := `C:\ACT\数据对接Frank\小票服务器补发货`
	dstFolder := filepath.Join(dstPath, filepath.Base(firstFolder))

	err = copyDir(firstFolder, dstFolder)
	if err != nil {
		fmt.Println("Error copying folder:", err)
		return
	}

	fmt.Println("Folder copied successfully to:", dstFolder)
}
