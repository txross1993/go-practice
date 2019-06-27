package main

import (
	"github.com/fsnotify/fsnotify" //Dependency on go get -u golang.org/x/sys/...
	log "github.com/txross1993/go-practice/WatchFolderForFiles/logwrapper"
)

func main() {
	li := log.NewLogger()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		li.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				li.Println("event:", event)
				if event.Op&fsnotify.Create == fsnotify.Create {
					li.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				li.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("foldertoWatch")
	if err != nil {
		li.Fatal(err)
	}
	<-done
}
