package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	inname, outname, err := filenamesFromCommandLine()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	infile, outfile := os.Stdin, os.Stdout
	if inname != "" {
		if infile, err = os.Open(inname); err != nil {
			log.Fatal(err)
		}
		defer infile.Close()
	}
	if outname != "" {
		if outfile, err = os.Create(outname); err != nil {
			log.Fatal(err)
		}
		defer outfile.Close()
	}

	if err = americanize(infile, outfile); err != nil {
		log.Fatal(err)
	}
}

// Returns the input and output file names from the command line; alternatively,
// dumps a helpful error message if the --help flags are used as the first
// command line argument to the program.
func filenamesFromCommandLine() (inname, outname string, err error) {
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		err = fmt.Errorf("usage: %s [<]infile.txt [>]outfile.txt",
			filepath.Base(os.Args[0]))
		return "", "", err
	}
	if len(os.Args) > 1 {
		inname = os.Args[1]
		if len(os.Args) > 2 {
			outname = os.Args[2]
		}
	}
	if inname != "" && inname == outname {
		log.Fatal("cannot overwrite the same file")
	}
	return inname, outname, nil
}

const BRITISH_AMERICAN = "british-american.txt"

type TransformFunc func(string) string

func americanize(infile io.Reader, outfile io.Writer) (err error) {
	reader := bufio.NewReader(infile)
	writer := bufio.NewWriter(outfile)
	defer func() {
		if err == nil {
			err = writer.Flush()
		}
	}()

	var transform TransformFunc
	if transform, err = makeTransformFunc(BRITISH_AMERICAN); err != nil {
		return err
	}

	wordRx := regexp.MustCompile("[A-Za-z]+")
	for eof := false; !eof; {
		var line string

		line, err = reader.ReadString('\n')
		if err == io.EOF {
			err = nil
			eof = true
		} else if err != nil {
			return err
		}
		line = wordRx.ReplaceAllStringFunc(line, transform)
		if _, err = writer.WriteString(line); err != nil {
			return err
		}
	}
	return nil
}

// makeTransformFunc accepts a file name and returns a function that uses the
// mapping generated from the file to provide a transformation between the
// british and the american versions of a word.
func makeTransformFunc(file string) (TransformFunc, error) {
	// Read the raw bytes from the file.
	rawBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// Generate the map from the lines in the text file, where each line has
	// fields for the key and value pair.
	text := string(rawBytes)
	usForBritish := make(map[string]string)
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) == 2 {
			usForBritish[fields[0]] = fields[1]
		}
	}

	// A function that uses the closure to return a replacement for a given word.
	return func(word string) string {
		if usWord, found := usForBritish[word]; found {
			return usWord
		}
		return word
	}, nil
}
