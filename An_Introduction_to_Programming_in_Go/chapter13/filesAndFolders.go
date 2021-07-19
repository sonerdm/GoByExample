package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		// handle the error here
		fmt.Println("No such file")
		return
	}
	//We use defer file.Close() right after opening the file to make sure the file is closed as soon as the function completes
	defer file.Close()

	// // get the file size
	// stat, err := file.Stat()
	// if err != nil {
	// 	return
	// }
	// // read the file
	// bs := make([]byte, stat.Size())
	// _, err = file.Read(bs)
	// if err != nil {
	// 	return
	// }
	//Reading files is very common, so there's a shorter way to do this:
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("No such file")
		return
	}

	str := string(bs)
	fmt.Println(str)
	//Here is how we can create a file:

	file2, err := os.Create("test2.txt")
	if err != nil {
		fmt.Println("err")
		return
	}
	defer file2.Close()

	file2.WriteString("test-soner")

	//To get the contents of a directory we use the same os.Open function but give it a directory path instead of a file name. Then we call the Readdir method:

	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name(), fi.Size(), fi.IsDir(), fi.Mode())
	}

	//Often we want to recursively walk a folder (read the folder's contents, all the sub-folders, all the sub-sub-folders, …). To make this easier there's a Walk function provided in the path/filepath package:
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
	//The function you pass to Walk is called for every file and folder in the root folder. (in this case .)
}
