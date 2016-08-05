// The MIT License (MIT)
//
// Copyright (c) 2015 Fredy Wijaya
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func show(path string, n int) {
	for i := 0; i < n; i++ {
		fmt.Print("|   ")
	}
	fmt.Println("|--", path)
}

func readDirNames(dirname string) ([]string, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	names, err := f.Readdirnames(-1)
	if err != nil {
		return nil, err
	}
	sort.Strings(names)
	return names, nil
}

func tree(path string, info os.FileInfo, n int) error {
	if !info.IsDir() {
		return nil
	}
	names, err := readDirNames(path)
	if err != nil {
		return err
	}
	for _, name := range names {
		filename := filepath.Join(path, name)
		fileInfo, err := os.Lstat(filename)
		if err != nil {
			return err
		}
		show(fileInfo.Name(), n)
		err = tree(filename, fileInfo, n+1)
		if err != nil {
			if !fileInfo.IsDir() {
				return err
			}
		}
	}
	return nil
}

// Tree displays the file tree structure.
func Tree(root string) error {
	info, err := os.Lstat(root)
	if err != nil {
		return err
	}
	fmt.Println(root)
	return tree(root, info, 0)
}

func validateArgs() error {
	if len(os.Args) != 2 {
		return errors.New("Usage: " + os.Args[0] + " <directory>")
	}
	return nil
}

func errorAndExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	err := validateArgs()
	if err != nil {
		errorAndExit(err)
	}
	err = Tree(os.Args[1])
	if err != nil {
		errorAndExit(err)
	}
}
