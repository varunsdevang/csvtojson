package converter

import "io"

type Converter interface {
	CsvToJson(io.Reader) ([]map[string]string, string, error)
}
