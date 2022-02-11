package converter

import "io"

type Converter interface {
	CsvToJson(io.Reader) ([]interface{}, string, error)
}
