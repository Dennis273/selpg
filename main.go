package main

import (
	"bufio"
	"fmt"
	flag "github.com/spf13/pflag"
	"io"
	"os"
	"os/exec"
)

type cliArgs struct {
	startPage   int
	endPage     int
	pageLength  int
	pageBreak   bool
	source      string
	destination string
}

func main() {
	var args = parseArgs()
	var valid = checkArgs(args)
	if !valid {
		fmt.Printf("Invalid arguments\nAborted")
	}
	readFile(args)
}

func parseArgs() cliArgs {
	var args cliArgs
	var startPage = flag.IntP("startPage", "s", 0, "Start Page Number")
	var endPage = flag.IntP("endPage", "e", 0, "End Page Number")
	var pageLength = flag.IntP("length", "l", 72, "Page Length")
	var pageBreak = flag.BoolP("formFeed", "f", false, "Form feed")
	var destination = flag.StringP("destination", "d", "", "Target")

	flag.Parse()

	args.startPage = *startPage
	args.endPage = *endPage
	args.pageLength = *pageLength
	args.pageBreak = *pageBreak
	args.destination = *destination
	args.source = flag.Args()[0]
	return args
}

func checkArgs(args cliArgs) bool {
	var valid = true
	if args.endPage < 0 || args.startPage < 0 || args.startPage > args.endPage {
		fmt.Println("Page Argumnet not valid.")
		valid = false
	}
	return valid
}

func readFile(args cliArgs) {
	file, err := os.Open(args.source)
	check(err)
	reader := bufio.NewReader(file)
	if args.pageBreak == true {
		// reading file by page
		for wastePageCtr := 0; wastePageCtr < args.startPage; wastePageCtr++ {
			_, err := reader.ReadString('\f')
			if err == io.EOF {
				break
			} else {
				check(err)
			}
		}
		for wantedPageCtr := args.startPage; wantedPageCtr <= args.endPage; wantedPageCtr++ {
			for lineCtr := 0; lineCtr < args.pageLength; lineCtr++ {
				str, err := reader.ReadString('\n')
				if err == io.EOF {
					break
				} else {
					check(err)
				}
				pipe(str, args.destination)
			}
		}
	} else {
		// reading file by line
		for wastePageCtr := 0; wastePageCtr < args.startPage; wastePageCtr++ {
			for lineCtr := 0; lineCtr < args.pageLength; lineCtr++ {
				_, err := reader.ReadString('\n')
				if err == io.EOF {
					break
				} else {
					check(err)
				}
			}
		}
		for wantedPageCtr := args.startPage; wantedPageCtr <= args.endPage; wantedPageCtr++ {
			for lineCtr := 0; lineCtr < args.pageLength; lineCtr++ {
				str, err := reader.ReadString('\n')
				if err == io.EOF {
					break
				} else {
					check(err)
				}
				pipe(str, args.destination)
			}
		}
	}
}

func pipe(str string, destination string) {
	if destination == "" {
		fmt.Print(str)
	} else {
		cmd := exec.Command("lp", "-d"+str)
		err := cmd.Run()
		check(err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
