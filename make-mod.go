package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func getDescriptionMod() []string {
	filename := "descriptor.mod"
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return []string{"Error reading file"}
	}

	content := string(data)
	return []string{content}
}

func getCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
	}

	nameOfDir := filepath.Base(dir)
	return nameOfDir
}

func writeMod(modFileName string, modContent []string) {
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Get the parent directory
	parentDir := filepath.Dir(currentDir)

	// Create the mod file in the parent directory
	file, err := os.Create(filepath.Join(parentDir, modFileName))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Write the mod file
	for _, line := range modContent {
		_, err = file.WriteString(line + "\n")
	}
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Append 'path="mod/"' + 'currentDir' (excluding the full path) to the file
	_, err = file.WriteString("path=\"mod/" + path.Base(getCurrentDirectory()) + "\"\n")

	// Check if the file was written successfully
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func main() {
	fmt.Println(path.Base(getCurrentDirectory()))

	modFileName := getCurrentDirectory() + ".mod"
	modContent := getDescriptionMod()
	writeMod(modFileName, modContent)
}
