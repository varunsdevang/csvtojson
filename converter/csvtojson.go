package converter

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"math"
	"sync"

	"github.com/labstack/gommon/log"
)

const (
	maxNumberOfRoutines = 1000
)

var (
	ErrorInvalidRow = errors.New("invalid row")

	ErrorInvalidHeader = errors.New("invalid header")
)

type CsvToJsonConverter struct{}

func (c *CsvToJsonConverter) CsvToJson(reader io.Reader) ([]map[string]string, string, error) {
	log.Info("In CsvToJson...")
	resp := []map[string]string{}

	csvReader := csv.NewReader(reader)
	csvReader.TrimLeadingSpace = true
	header, err := csvReader.Read()
	if err != nil {
		return resp, err.Error(), ErrorInvalidHeader
	}

	records, err := csvReader.ReadAll()
	if err != nil {
		cerr := err.(*csv.ParseError)
		return resp, fmt.Sprintf("error in line %d column %d, error: %s", cerr.Line, cerr.Column, cerr.Err.Error()), ErrorInvalidRow
	}

	noOfRoutines := math.Min(maxNumberOfRoutines, float64((len(records))))
	in := make(chan []string, len(records))
	out := make(chan map[string]string)
	wg := sync.WaitGroup{}
	wg.Add(int(noOfRoutines))

	log.Info("no of workers: ", noOfRoutines)
	for i := 0; i < int(noOfRoutines); i++ {
		go func(i int) {
			for {
				c, ok := <-in
				if !ok {
					wg.Done()
					return
				}
				csvToJsonWorker(header, c, out)
			}
		}(i)
	}

	go func() {
		for i := 0; i < len(records); i++ {
			in <- records[i]
		}
		close(in)
	}()

	for i := 0; i < len(records); i++ {
		record := <-out
		resp = append(resp, record)
	}

	close(out)
	wg.Wait()
	return resp, "", nil
}

func csvToJsonWorker(header, record []string, output chan map[string]string) {
	m := make(map[string]string)
	for i, f := range record {
		m[header[i]] = f
	}
	output <- m
}
