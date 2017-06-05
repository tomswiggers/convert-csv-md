package main

import (
    "fmt"
    "flag"
    "os"
    "io"
    "strings"
    "path/filepath"
    "encoding/csv"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func countHeader(record []string) (int) {
  i := 0
  c := 0

  for range record {

    if strings.Contains(record[i], "--") {
      c++
      fmt.Println("head")
    }

    i++
  }

  return c
}

func main() {
  headerCounter := 0

  filename := flag.String("filename", "tables.csv", "a filename to parse")
  flag.Parse()

  filenameMd := strings.TrimSuffix(*filename, filepath.Ext(*filename)) + ".md"

  fmt.Println("Use CSV: " + *filename)
  fmt.Println("Use MD: " + filenameMd)

  csvFp, err := os.Open(*filename)
  check(err)

  mdFp, err := os.Create(filenameMd)

  r := csv.NewReader(csvFp)

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

    check(err)

    c := countHeader(record)

    if c > 0 {
      headerCounter = c
    }

    i := 0
    output := "|"

    for range record {

      // This needs to be refactored
      // doesn't work for the first line of the table
      if i < headerCounter {
        output = output + record[i] + "|"
      }

      i++
    }

    mdFp.WriteString(output + "\n")
		fmt.Println(output)
	}

  csvFp.Close()
  mdFp.Close()
}
