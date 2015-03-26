package rpm

// Entry defines a provided package.
type Entry struct {
	Name  string
	Flags string
	Version
}

// Provides defines provided entries.
type Provides struct {
	Entries []Entry
}

// Format defines the package format
type Format struct {
	License     string
	Vendor      string
	Group       string
	BuildHost   string
	SourceRPM   string
	HeaderRange string
}

// Size defines a package size in various states.
type Size struct {
	Package   int `xml:"package"`
	Installed int `xml:"installed"`
	Archive   int `xml:"archive"`
}

// Time defines various times on a package.
type Time struct {
	File  int `xml:"file"`
	Build int `xml:"build"`
}

// Version defines a package version
type Version struct {
	Epoch   string `xml:"epoch,attr"`
	Version string `xml:"ver,attr"`
	Release string `xml:"rel,attr"`
	PkgID   string `xml:"pkgid,attr"`
}

// Package defines package metadata
type Package struct {
	Type        string   `xml:"type,attr"`
	Name        string   `xml:"name"`
	Arch        string   `xml:"arch"`
	Version     Version  `xml:"version"`
	Checksum    Checksum `xml:"checksum"`
	Summary     string   `xml:"summary"`
	Description string   `xml:"description"`
	Packager    string   `xml:"packager"`
	URL         string   `xml:"url"`
	Time        Time     `xml:"time"`
	Location    Location `xml:"location"`
}

// Metadata defines RPM repository metadata.
type Metadata struct {
	XMLNamespace    string    `xml:"xmlns,attr"`
	XMLNamespaceRPM string    `xml:"xmlnx:rpm,attr"`
	Length          int       `xml:"packages,attr"`
	Packages        []Package `xml:"package"`
}
