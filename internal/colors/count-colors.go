package colors

import (
	"encoding/xml"
)

// CompetitorProfile XML top level
type CompetitorProfile struct {
	XMLName xml.Name   `xml:"competitor_profile"`
	Jerseys JerseyList `xml:"jerseys"`
}

// JerseyList XML jerseys tag
type JerseyList struct {
	XMLName    xml.Name `xml:"jerseys"`
	JerseyList []Jersey `xml:"jersey"`
}

// Jersey XML jersey tag
type Jersey struct {
	XMLName xml.Name `xml:"jersey"`
	Colors  string   `xml:"colors,attr"`
}

// Count returns number of occurrences of the given color in xml files in the dataDir
// TODO: implement
// TODO: think about performance improvements, error handling, ...
func Count(dataDir, color string) (int, error) {
	var total int

	fileNames, _ := getFileList(dataDir)

	for _, fName := range fileNames {
		getColorOccurrences(fName, color)
	}
	return total, nil
}

// getFileList returns list of XML files in the given directory
// TODO: implement
func getFileList(dataDir string) (fileNames []string, err error) {
	return []string{"my-file.xml"}, nil
}

// getColorOccurrences returns number of occurrences of the given color in the XML file
// TODO: implement
func getColorOccurrences(fName, color string) (int, error) {
	return 666, nil
}
