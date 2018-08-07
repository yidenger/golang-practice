package main

import (
	"fmt"
	"strings"
	"time"
)

type logProcess struct {
	read Reader
	write Writer
	rc chan string
	wc chan string
}

type Reader interface {
	Read(rc chan string)
}

type Writer interface {
	Write(rc chan string)
}

type ReadFromFile struct {
	path string
}

type WriteToDb struct {
	dbUrl string
}

func (r *ReadFromFile) Read(rc chan string) {
	rc <- "this is a message"
}

func (w *WriteToDb) Write(wc chan string) {
	fmt.Println(<-wc)
}

func (l *logProcess) process() {
	l.wc <- strings.ToUpper(<-l.rc)
}

func main() {
	r := &ReadFromFile{
		path: "./log.txt",
	}

	w := &WriteToDb{
		dbUrl: "username:password@host",
	}

	l := &logProcess{
		read: r,
		write: w,
		rc: make(chan string),
		wc: make(chan string),
	}

	go l.read.Read(l.rc)
	go l.process()
	go l.write.Write(l.wc)

	time.Sleep(time.Second)
}
