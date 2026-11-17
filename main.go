package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/alecthomas/kong"
)

type ARGS struct {
	Addr string `default:"127.0.0.1:8000" short:"a" help:"Enter address to serve."`
	Path string `default:"." short:"p" help:"Enter folder to serve."`
	Log  string `default:"" short:"l" help:"Enter filename to write logs to. (leave empty for no log file)"`
}

func main() {
	exit_log()

	var a ARGS
	kong.Parse(&a)
	host(a.Addr, a.Path, a.Log)
}

func host(addr string, path string, logpath string) {
	fi, err := os.Stat(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s does not exist\n", path)
		os.Exit(1)
	}

	if !fi.IsDir() {
		fmt.Fprintf(os.Stderr, "%s is not a directory\n", path)
		os.Exit(1)
	}

	dir := http.Dir(path)
	fs := http.FileServer(dir)

	out := io.Discard
	if logpath != "" {
		f, err := os.OpenFile(logpath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		out = io.MultiWriter(os.Stderr, f)
		log.SetOutput(out)
	}

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s => %s\n", r.RemoteAddr, r.URL.Path)
		f, err := dir.Open(r.URL.Path)
		if err != nil {
			fs.ServeHTTP(w, r)
			return
		}
		defer f.Close()

		stat, err := f.Stat()
		if !stat.IsDir() {
			w.Header().Set("Content-Disposition", "attachment")
		}

		fs.ServeHTTP(w, r)
	}))

	log.Printf("Serving %s on http://%s", path, addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func exit_log() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		s := <-sigs
		log.Println("Exit:", s)
		os.Exit(0)
	}()
}
