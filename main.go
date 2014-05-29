package main

import (
	"flag"
	"path/filepath"
	"os"
	"fmt"
	"image/jpeg"
)

func WalkFunc(path string, info os.FileInfo, err error) error {
	if info.IsDir() == false {
		var ext = filepath.Ext(path)
		if ext == ".jpg" {
			f, fok := os.OpenFile(path, os.O_RDONLY, 0)
			defer f.Close()
			if fok == nil {
				_, ok := jpeg.Decode(f)
				if ok != nil {
					fmt.Println(path)
				}
			}
			
		}			
	}
	
	return nil
}

func main() {
	var path = flag.String("path", "./", "directory path")
	//var type = flag.String("type", "jpg", "jpg(default) or png")
	flag.Parse()
	
	fmt.Println("fake jpg searching...")
	filepath.Walk(*path, WalkFunc)
	fmt.Println("end...")
	
}