package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

var reg *regexp.Regexp

func main() {
	// Make a Regex to say we only want letters and numbers
	var err error
	reg, err = regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	oldFolder := "PM_OLD"
	err = os.MkdirAll(oldFolder, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	message := ""

	for _, f := range files {
		fName := f.Name()
		if strings.Contains(fName, ".txt") && !strings.Contains(fName, "_old.txt") {

			lines, err := readLines(fName)
			if err != nil {
				log.Fatalf("readLines: %s", err)
			}
			if !needsFixing(lines) {
				message += f.Name() + " is already fixed\n"
				continue
			}
			message += f.Name() + " fixed\n"
			fOldName := strings.Replace(fName, ".txt", "_old.txt", 1)
			err = writeLines(lines, oldFolder+"/"+fOldName)
			if err != nil {
				log.Fatalf("old file writeLines: %s", err)
			}
			updateLines(lines)
			err = writeLines(lines, fName)
			if err != nil {
				log.Fatalf("new file writeLines: %s", err)
			}
		}
	}
	app := app.New()

	w := app.NewWindow("Format Changed")
	w.SetContent(widget.NewVBox(
		widget.NewLabel(message),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))

	w.ShowAndRun()
}

func needsFixing(lines []string) bool {
	if strings.Count(lines[1], "|") == 36 {
		return true
	}
	return false
}

func updateLines(lines []string) {
	s := strings.Split(lines[1], "|")
	storeName := s[6]
	lines[0] = strings.TrimSpace(lines[0]) + "|" + storeName
	for i := 1; i < len(lines); i++ {
		l := strings.Split(lines[i], "|")
		number := strings.TrimSpace(l[36])
		number = reg.ReplaceAllString(number, "")
		if len(number) == 10 {
			number = number[4:]
		} else if len(number) == 6 {
			// keep number
		} else {
			number = ""
		}
		lines[i] = strings.TrimSpace(lines[i])
		lines[i] += "|" + number + "|||||||"
	}
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\r\n")
	return lines, nil
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, filename string) error {
	data := []byte(strings.Join(lines, "\r\n"))
	err := ioutil.WriteFile(filename, data, 0666)
	if err != nil {
		return err
	}
	return nil
}
