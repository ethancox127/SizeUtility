package utils

import (
	"fmt"
	"path/filepath"
	"os"
	"strings"
)

// Function for finding the size of a directory non recursive.
func DirSize(dirname string, human bool) (int64, error) {

	var size int64

	// Add the size of all files in the directory if there are no errors reading the directory items.
    items, err := os.ReadDir(dirname)
	if err != nil {
		return 0, err
	}
    for _, item := range items {
		if !item.IsDir() {
            entry, err := item.Info()
			if err != nil {
				fmt.Println(err)
				
			}
            size += entry.Size()
        }
	}

	// Print the size of the directory.
	if human {
		HumanPrint(dirname, size)
	} else {
		Print(dirname, size)
	}

	// Return the size of the directory.
    return size, nil
}

func RecursiveDirSize(dirname string, human bool, size *int64) (int64, error) {

	var curSize int64 = 0
	dirs := []string{}

	// Add the size of all files in the directory if there are no errors reading the directory items.
	items, err := os.ReadDir(dirname)
	if err != nil {
		return 0, err
	}
    for _, item := range items {
		if !item.IsDir() {
			entry, err := item.Info()
			if err != nil {
				fmt.Println(err)
				
			}
            curSize += entry.Size()
        } else if item.Name() != "." && item.Name() != ".." && item.Name() != dirname {
			// Add the subdirectories to the the slice of directories.
			dirs = append(dirs, item.Name())
		}
	}

	// Print the size of the directory.
	if human {
		HumanPrint(dirname, curSize)
	} else {
		Print(dirname, curSize)
	}

	// Update the size of all the directory and subdirectories.
	*size += curSize

	// For each subdirectory add it to the current path, recursively call this function, and remove the subdirectory from the path after completion.
	for i := 0; i < len(dirs); i++ {
		suffix := "/" + dirs[i]
		dirname += suffix
		RecursiveDirSize(dirname, human, size)
		dirname = strings.TrimSuffix(dirname, suffix)
	}

	// Return the size.
    return *size, nil
}

// Function for finding the absolute path of a directory.
func AbsPath(dirname string) string {
	path, err := filepath.Abs(dirname)
	if err != nil {
		fmt.Println(err)
	}
	return path
}

// Function for printing the size of a directory.
func Print(dirname string, size int64) {
	fmt.Print("Directory " + dirname + " is size: ")
	fmt.Print(size)
	fmt.Println(" bytes")
}

// Function for printing the size of a directory in a human friendly format.
func HumanPrint(dirname string, size int64) {
	fmt.Print("Directory " + dirname + " is size: ")
	PrettyPrint(size)
}

// Function for determing how to print the size of the directory in human friendly format.
func PrettyPrint(size int64) {
	var kBytes, mBytes, gBytes, tBytes int64 = 1000, 1000000, 1000000000, 1000000000000
	
	// Checks if the directory size is best represented by kilobytes, megabytes, gigabytes, or terrabytes.
	if size < kBytes {
		fmt.Print(size)
		fmt.Println(" bytes")
		return
	} else if size / tBytes > 0 {
		fmt.Print(size / tBytes)
		fmt.Println("T")
		return
	} else if size / gBytes > 0 {
		fmt.Print(size / gBytes)
		fmt.Println("G")
		return
	} else if size / mBytes > 0 {
		fmt.Print(size / mBytes)
		fmt.Println("M")
		return
	} else if size / kBytes > 0 {
		fmt.Print(size / kBytes)
		fmt.Println("K")
		return
	}
}