package main

import (
	"log"
	"fmt"
	"flag"
	"strings"
	"github.com/dotcloud/beam"
)

// This is a placeholder for the beam command-line tool.
// It is meant as a convenience to expose capabilities of the protocol and library
// from the command-line.
func main() {
	flag.Parse()
	worker := beam.NewWorker(&beam.NetTransport{"tcp", flag.Arg(0)}, "/jobs")
	worker.RegisterJob("hello", JobHello)
	if err := worker.Work(); err != nil {
		log.Fatal(err)
	}
}

func JobHello(name string, args []string, env map[string]string, streams *beam.Streamer, db beam.DB) error {
	fmt.Printf("Hello, %s!\n", strings.Join(args, " "))
	return nil
}
