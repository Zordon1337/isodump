package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/kdomanski/iso9660"
)

func main() {
	fmt.Println("Enter partition where your CD/DVD is mounted:")
	var partition string
	fmt.Scanln(&partition)
	fmt.Println("Enter output iso name(without .iso):")
	var output string
	fmt.Scanln(&output)

	if strings.Contains(output, ".iso") {
		fmt.Println("Output name contains .iso, removing...")
		output = strings.Replace(output, ".iso", "", -1)
	}

	fmt.Println("Creating iso...")
	// Create writer
	writer, err := iso9660.NewWriter()
	if err != nil {
		fmt.Println("Error creating writer")
		return
	}
	defer writer.Cleanup()
	// add disk content to writer
	writer.AddLocalDirectory(partition, "/")
	// Open file handle
	target, err := os.OpenFile(output+".iso", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error creating output file")
		return
	}
	defer target.Close()
	// Write to target
	writer.WriteTo(target, partition)
	fmt.Println("Done!")
}
