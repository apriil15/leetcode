package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"
)

// example: go run . -folder=xxx
func main() {
	folderName := flag.String("folder", "", "Name of the folder to create")
	flag.Parse()

	CreateAndOpen(*folderName)
}

// CreateAndOpen create new folder and main.go, finally open main.go
func CreateAndOpen(folderName string) {
	folderPath := path.Join("..", "..", folderName)
	if err := os.Mkdir(folderPath, 0755); err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	filePath := path.Join(folderPath, "main.go")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	content := `package main

func main() {

}
`
	if _, err := file.WriteString(content); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Open the "main.go" file
	if err := openFile(filePath); err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
}

// openFile opens a file in the default editor based on the OS
func openFile(fileName string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", fileName)
	case "darwin":
		cmd = exec.Command("open", fileName)
	case "linux":
		cmd = exec.Command("xdg-open", fileName)
	default:
		return fmt.Errorf("unsupported platform")
	}

	return cmd.Start()
}
