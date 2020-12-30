package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/BlankDev117/gomockserver/src/parsing"
	"github.com/BlankDev117/gomockserver/src/routes"
	"github.com/BlankDev117/gomockserver/src/schemas"
)

// ReadConfigurationFiles Reads a single file or entire directory and generates manageable route maps from the configuration files
func ReadConfigurationFiles(filePath string) ([]routes.RouteMap, error) {

	maps, err := parseConfigurationFiles(filePath)

	if err != nil {
		return nil, err
	}

	return parseSchemaMaps(maps)
}

func parseConfigurationFiles(filePath string) ([]map[string]interface{}, error) {

	log.Printf("Attempting to parse configuration files at %s", filePath)
	fileInfo, err := os.Stat(filePath)

	if err != nil {
		return nil, err
	}

	fileInfos := []os.FileInfo{}
	mode := fileInfo.Mode()
	switch mode.IsDir() {
	case true:
		log.Printf("File path is for a directory. Reading All files in directory...")
		f, err := os.Open(filePath)

		if err != nil {
			return nil, err
		}

		fileInfos, err = f.Readdir(-1)
		f.Close()

		if err != nil {
			return nil, err
		}

		log.Printf("Files have been loaded. A total of %d files were found.", len(fileInfos))
	case false:
		log.Printf("File path is for a single file.")
		fileInfos = append(fileInfos, fileInfo)
	}

	parserProvider := parsing.DefaultFileParserProvider()
	maps := []map[string]interface{}{}
	for _, fileInfo := range fileInfos {

		// Only parse one level deep from input path
		if fileInfo.Mode().IsDir() {
			continue
		}

		configSettingPath := filepath.Join(filePath, fileInfo.Name())
		parser, err := parserProvider.GetParser(configSettingPath)

		if err != nil {
			return nil, err
		}

		parsedMap, err := parser.ParseFile(configSettingPath)

		if err != nil {
			return nil, err
		}

		maps = append(maps, parsedMap)
	}

	return maps, nil
}

func parseSchemaMaps(schemaMaps []map[string]interface{}) ([]routes.RouteMap, error) {
	schemaProvider := schemas.DefaultSchemaProvider()

	routeMaps := []routes.RouteMap{}
	for _, schemaMap := range schemaMaps {
		schema, err := schemaProvider.GetSchema(schemaMap)

		if err != nil {
			return nil, err
		}

		routes, err := schema.ReadMap(schemaMap)

		if err != nil {
			return nil, err
		}

		for _, route := range routes {
			routeMaps = append(routeMaps, route)
		}
	}

	log.Printf("A total of %d routes were configured for the mock server.", len(routeMaps))

	return routeMaps, nil
}
