package parsing

import (
	"io/ioutil"

	"github.com/BlankDev117/gomockserver/src/utils"
)

// JsonFileParser is a file parser that handles files using json format
type JsonFileParser struct {
}

// #region Constructors

// DefaultJSONFileParser creates a parser capable of handling json file extensions
func DefaultJSONFileParser() JsonFileParser {
	return JsonFileParser{}
}

// #endregion

// #region ifileparser

// SupportedExtensions returns the extensions this parser supports
func (parser JsonFileParser) SupportedExtensions() []string {
	return []string{".json"}
}

// ParseFile will parse the given file, if it is a supported extension, and return the parsed file as a map
func (parser JsonFileParser) ParseFile(path string) (map[string]interface{}, error) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	dataMap, err := utils.ConvertBytesToMap(data)

	if err == nil {
		return dataMap, nil
	}

	return nil, err
}

// #endregion
