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

func main() {

  filename := flag.String("filename", "tables.csv", "a filename to parse")
  flag.Parse()

  filenameMd := strings.TrimSuffix(*filename, filepath.Ext(*filename)) + ".md"

  fmt.Println("Use CSV: " + *filename);
  fmt.Println("Use MD: " + filenameMd);

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

    i := 0

    for range record {
      fmt.Println(record[i])

      i++
    }

		fmt.Println(record)
	}

  mdFp.WriteString("test\n")

  csvFp.Close()
  mdFp.Close()
}
