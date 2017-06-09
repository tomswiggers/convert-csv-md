package main

import (
    "fmt"
    "flag"
    "os"
    "io"
    "strings"
    "strconv"
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

func countColumns(record []string) (int) {
  i := 0
  c := 0

  for range record {

    if len(record[i]) > 0 {
      c++
    }

    i++
  }

  return c
}

func checkEndOfTable(record []string) (bool) {
  i := 0

  for range record {

    if len(record[i]) > 0 {
      return false
    }

    i++
  }

  return true
}

func main() {
  var filenameMd string
  var mdFp *os.File

  headerCounter := 0
  tableCounter := 0
  beginTable := true

  filename := flag.String("filename", "tables.csv", "a filename to parse")
  flag.Parse()

  fmt.Println("Use CSV: " + *filename)

  csvFp, err := os.Open(*filename)
  check(err)


  r := csv.NewReader(csvFp)

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

    check(err)

    if beginTable {
      headerCounter = countColumns(record)
      beginTable = false

      tableCounter++

      filenameMd = strings.TrimSuffix(*filename, filepath.Ext(*filename)) + strconv.Itoa(tableCounter) + ".md"
      mdFp, err = os.Create(filenameMd)
    }

    if checkEndOfTable(record) {
      mdFp.Close()
      beginTable = true
    }

    i := 0
    output := "|"

    for range record {

      if i < headerCounter {
        output = output + record[i] + "|"
      }

      i++
    }

    mdFp.WriteString(output + "\n")
		fmt.Println(output)
    defer mdFp.Close()
	}

  csvFp.Close()
}
