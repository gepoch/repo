package rpm

// Checksum represents the checksum tags present in repomd files.
type Checksum struct {
	Type string `xml:"type,attr"`
	Sum  string `xml:",chardata"`
}

// Location is a hyperlink to the data file.
type Location struct {
	Href string `xml:"href,attr"`
}

// Data contains information about a file in the repodata.
type Data struct {
	Type            string   `xml:"type,attr"`
	Checksum        Checksum `xml:"checksum"`
	OpenChecksum    Checksum `xml:"open-checksum"`
	Location        Location `xml:"location"`
	DatabaseVersion int      `xml:"database_version"`
	OpenSize        int      `xml:"open-size"`
	Timestamp       int      `xml:"timestamp"`
	Size            int      `xml:"size"`
}

// Repomd is a struct designed to unmarshall repomd.xml files in XML.
type Repomd struct {
	Revision int    `xml:"revision"`
	Data     []Data `xml:"data"`
}
