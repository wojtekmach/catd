package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/howeyc/fsnotify"
)

func run(dir string, file string) {
	cmd := fmt.Sprintf("cat %s/* > %s", dir, file)
	_, err := exec.Command("sh", "-c", cmd).Output()

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	once := flag.Bool("once", false, "Run once and quit")
	dir := flag.String("dir", "", "Directory to poll")
	file := flag.String("file", "", "File to concatenate contents of dir to")
	flag.Parse()

	if _, err := os.Stat(*dir); os.IsNotExist(err) {
		fmt.Println("No such file or directory", *dir)
		os.Exit(1)
	}

	if *once {
		run(*dir, *file)
		os.Exit(0)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)

	// Process events
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				log.Println("event:", ev)
				run(*dir, *file)

			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Watch(*dir)
	if err != nil {
		log.Fatal(err)
	}

	<-done

	watcher.Close()
}
