package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	oldName := ""
	newName := ""
	dirPath := "."

	//User input
	fmt.Print("Name of the folder you wish to recursively rename (case sensitive): ")
	fmt.Scan(&oldName)
	fmt.Print("Name that folder should be recursively renamed to (case sensitive): ")
	fmt.Scan(&newName)
	fmt.Print("Path to root folder to scan. Use '.' for this dir (case sensitive): ")
	fmt.Scan(&dirPath)

	fmt.Print()

	var oldPaths []string
	var newPaths []string
	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() == true {
			dirList, err := wut_files(path)
			if err != nil {
				fmt.Println(err)
			}
			for k := 0; k < len(dirList); k++ {
				//fmt.Println(k) //debugging
				if dirList[k] == oldName {
					//debugging
					//fmt.Println(dirList[i])
					//fmt.Println("/" + path + "/" + dirList[k] + " is a match") // replace the older folder name with the new folder
					//fmt.Println("/" + path + "/" + newName + " is the new path")

					oldPaths = append(oldPaths, "./"+path+"/"+dirList[k])
					newPaths = append(newPaths, "./"+path+"/"+newName)

				}
			}
		}
		return nil
	})

	//tell me what's gonna change
	fmt.Println("The paths that will change: ")
	for i := 0; i < len(oldPaths); i++ {
		fmt.Println(oldPaths[i])
	}
	fmt.Println()

	fmt.Println("What the paths will become: ")
	for i := 0; i < len(newPaths); i++ {
		fmt.Println(newPaths[i])
	}
	fmt.Println()

	//ask if that's okay
	// var then variable name then variable type
	var okeyDokey string
	fmt.Print("Are these changes okay? Y/n: ")
	fmt.Scan(&okeyDokey)

	if okeyDokey == "Y" {
		//do the changes
		for i := 0; i < len(oldPaths); i++ {
			fileErr := os.Rename(oldPaths[i], newPaths[i])
			if fileErr != nil {
				fmt.Println("Error renaming:", oldPaths, fileErr.Error())
			}
		}
	} else {
		fmt.Println("Cancelling name change operation")
	}

	return
}

// what files are in a given directory
func wut_files(dirPath string) (the_files []string, err error) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	} else {
		for _, file := range files {
			//fmt.Println(file.Name())
			the_files = append(the_files, file.Name())
		}
	}
	return the_files, err
}

// is directory?, outputs 1, 0, error
func directoryCheck(dirPath string) (bool, error) {
	fileInfo, err := os.Stat(dirPath)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), err
}

func directoryListCheck(dirPath string, dirList []string) ([]bool, []error) {
	var dirBoolForDir []bool
	var errList []error
	for i := 0; i < len(dirList); i++ {
		one, _ := directoryCheck(dirPath + "/" + dirList[i]) //catch the boolean
		_, two := directoryCheck(dirPath + "/" + dirList[i]) //catch the error
		dirBoolForDir = append(dirBoolForDir, one)

		//append errors if they occur
		if two != nil {
			errList = append(errList, two)
		}
	}
	return dirBoolForDir, errList
}
