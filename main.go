package main

import (
	"gopkg.in/qml.v0"
	"path/filepath"
	"os"
	"image/jpeg"
	"strings"
	"container/list"
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
					fake_list.PushBack(path)
				}
			}
		}
	}

	return nil
}

var fake_list *list.List

func init() {
	fake_list = list.New()
}

func main() {
	qml.Init(nil)
	engine := qml.NewEngine()
	component, err := engine.LoadFile("main.qml")
	if err != nil {
		panic(err)
	}

	ctrl := Control{}
	
	context := engine.Context()
	context.SetVar("ctrl", &ctrl)

	window := component.CreateWindow(nil)

	ctrl.Root = window.Root()

	window.Show()
	window.Wait()
}


type Control struct {
	Root qml.Object
}

func (ctrl *Control) OnAcceptedBtnClicked(btn qml.Object, path string) {
	folder_path := strings.TrimPrefix(path, "file://")
	fake_list.Init()
	filepath.Walk(folder_path, WalkFunc)

	resultCtrl := ctrl.Root.ObjectByName("result")
	var list_value string

	for itm := fake_list.Front(); itm != nil; itm = itm.Next() {
		a, _:= itm.Value.(string)
		list_value += a
		list_value += "\n"
	}

	resultCtrl.Set("text", list_value)
}

