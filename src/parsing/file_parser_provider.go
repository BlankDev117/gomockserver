package parsing

import "errors"

// FileParserProvider helps with providing an ifileparser for a given file extension
type FileParserProvider struct {
	Parsers []IFileParser
}

// #region Constructors

// DefaultFileParserProvider creates a parser provider that can help with parsing a given file. Default Provider supports the following formats: .json
func DefaultFileParserProvider() FileParserProvider {
	return FileParserProvider{[]IFileParser{DefaultJSONFileParser()}}
}

// #endregion

// #region Helpers

// GetParser retrieves an ifileparser that can handle parsing the given file
func (parserProvider FileParserProvider) GetParser(filePath string) (IFileParser, error) {
	for _, parser := range parserProvider.Parsers {
		if CanParse(parser, filePath) {
			return parser, nil
		}
	}

	return nil, errors.New("No parsers could parse the file extension at the following path: " + filePath)
}

// #endregion
