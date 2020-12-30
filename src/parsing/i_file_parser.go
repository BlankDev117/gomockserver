package parsing

import (
	"path/filepath"

	"github.com/BlankDev117/gomockserver/src/utils"
)

// IFileParser is an object that parse a file that is listed in its supported extensions
type IFileParser interface {
	SupportedExtensions() []string
	ParseFile(filePath string) (map[string]interface{}, error)
}

// #region Helpers

// CanParse Deterimine whether the given parser is capable of handling the specified file
func CanParse(parser IFileParser, filePath string) bool {
	if filePath == "" {
		return false
	}

	extension := filepath.Ext(filePath)
	return utils.StringInSlice(extension, parser.SupportedExtensions())
}

// #endregion
