package rpm

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/kr/pretty"
)

var repomdXML = `<?xml version="1.0" encoding="UTF-8"?>
<repomd>
  <revision>1427333070</revision>
  <data type="filelists">
    <checksum type="sha256">401dc19bda88c82c403423fb835844d64345f7e95f5b9835888189c03834cc93</checksum>
    <open-checksum type="sha256">bf9808b81cb2dbc54b4b8e35adc584ddcaa73bd81f7088d73bf7dbbada961310</open-checksum>
    <location href="repodata/401dc19bda88c82c403423fb835844d64345f7e95f5b9835888189c03834cc93-filelists.xml.gz"></location>
    <database_version>0</database_version>
    <open-size>125</open-size>
    <timestamp>1427333070</timestamp>
    <size>123</size>
  </data>
  <data type="primary">
    <checksum type="sha256">dabe2ce5481d23de1f4f52bdcfee0f9af98316c9e0de2ce8123adeefa0dd08b9</checksum>
    <open-checksum type="sha256">e1e2ffd2fb1ee76f87b70750d00ca5677a252b397ab6c2389137a0c33e7b359f</open-checksum>
    <location href="repodata/dabe2ce5481d23de1f4f52bdcfee0f9af98316c9e0de2ce8123adeefa0dd08b9-primary.xml.gz"></location>
    <database_version>0</database_version>
    <open-size>167</open-size>
    <timestamp>1427333070</timestamp>
    <size>134</size>
  </data>
  <data type="primary_db">
    <checksum type="sha256">5573105e3b55ac682b5741d82dba137d48f4957ac767955d6bbd9c71bd5e7299</checksum>
    <open-checksum type="sha256">7adbfb3208efc174230a8b0faf87f7d6041ef0c3d1cce782bec29c83fc9ad0b9</open-checksum>
    <location href="repodata/5573105e3b55ac682b5741d82dba137d48f4957ac767955d6bbd9c71bd5e7299-primary.sqlite.bz2"></location>
    <database_version>10</database_version>
    <open-size>21504</open-size>
    <timestamp>1427333070</timestamp>
    <size>1132</size>
  </data>
  <data type="other_db">
    <checksum type="sha256">208e5ef9572eea5835b70527c5e7cc85cc5c09cf05a4e0c7ad9b427d7f3501c8</checksum>
    <open-checksum type="sha256">b0f01a349e902aa0a038574fcd74bcfeb877c7d49278356be134a8f858cf0290</open-checksum>
    <location href="repodata/208e5ef9572eea5835b70527c5.Fe7cc85cc5c09cf05a4e0c7ad9b427d7f3501c8-other.sqlite.bz2"></location>
    <database_version>10</database_version>
    <open-size>6144</open-size>
    <timestamp>1427333070</timestamp>
    <size>572</size>
  </data>
  <data type="other">
    <checksum type="sha256">6bf9672d0862e8ef8b8ff05a2fd0208a922b1f5978e6589d87944c88259cb670</checksum>
    <open-checksum type="sha256">e0ed5e0054194df036cf09c1a911e15bf2a4e7f26f2a788b6f47d53e80717ccc</open-checksum>
    <location href="repodata/6bf9672d0862e8ef8b8ff05a2fd0208a922b1f5978e6589d87944c88259cb670-other.xml.gz"></location>
    <database_version>0</database_version>
    <open-size>121</open-size>
    <timestamp>1427333070</timestamp>
    <size>123</size>
  </data>
  <data type="filelists_db">
    <checksum type="sha256">6174899a2d38869b23cab3b6308eb0c32ab3480b20aaca93226695c4684c4bf9</checksum>
    <open-checksum type="sha256">93e39c97d742867e8187e51deb3a7cdfe95a7e8dacdadde835a2ccd61ddf7e40</open-checksum>
    <location href="repodata/6174899a2d38869b23cab3b6308eb0c32ab3480b20aaca93226695c4684c4bf9-filelists.sqlite.bz2"></location>
    <database_version>10</database_version>
    <open-size>7168</open-size>
    <timestamp>1427333070</timestamp>
    <size>591</size>
  </data>
</repomd>`

var repomdStruct = &repomd{
	Revision: 1427333070,
	Data: []Data{
		{
			Type:            "filelists",
			Checksum:        Checksum{Type: "sha256", Sum: "401dc19bda88c82c403423fb835844d64345f7e95f5b9835888189c03834cc93"},
			OpenChecksum:    Checksum{Type: "sha256", Sum: "bf9808b81cb2dbc54b4b8e35adc584ddcaa73bd81f7088d73bf7dbbada961310"},
			Location:        Location{Href: "repodata/401dc19bda88c82c403423fb835844d64345f7e95f5b9835888189c03834cc93-filelists.xml.gz"},
			DatabaseVersion: 0,
			OpenSize:        125,
			Timestamp:       1427333070,
			Size:            123,
		},
		{
			Type:            "primary",
			Checksum:        Checksum{Type: "sha256", Sum: "dabe2ce5481d23de1f4f52bdcfee0f9af98316c9e0de2ce8123adeefa0dd08b9"},
			OpenChecksum:    Checksum{Type: "sha256", Sum: "e1e2ffd2fb1ee76f87b70750d00ca5677a252b397ab6c2389137a0c33e7b359f"},
			Location:        Location{Href: "repodata/dabe2ce5481d23de1f4f52bdcfee0f9af98316c9e0de2ce8123adeefa0dd08b9-primary.xml.gz"},
			DatabaseVersion: 0,
			OpenSize:        167,
			Timestamp:       1427333070,
			Size:            134,
		},
		{
			Type:            "primary_db",
			Checksum:        Checksum{Type: "sha256", Sum: "5573105e3b55ac682b5741d82dba137d48f4957ac767955d6bbd9c71bd5e7299"},
			OpenChecksum:    Checksum{Type: "sha256", Sum: "7adbfb3208efc174230a8b0faf87f7d6041ef0c3d1cce782bec29c83fc9ad0b9"},
			Location:        Location{Href: "repodata/5573105e3b55ac682b5741d82dba137d48f4957ac767955d6bbd9c71bd5e7299-primary.sqlite.bz2"},
			DatabaseVersion: 10,
			OpenSize:        21504,
			Timestamp:       1427333070,
			Size:            1132,
		},
		{
			Type:            "other_db",
			Checksum:        Checksum{Type: "sha256", Sum: "208e5ef9572eea5835b70527c5e7cc85cc5c09cf05a4e0c7ad9b427d7f3501c8"},
			OpenChecksum:    Checksum{Type: "sha256", Sum: "b0f01a349e902aa0a038574fcd74bcfeb877c7d49278356be134a8f858cf0290"},
			Location:        Location{Href: "repodata/208e5ef9572eea5835b70527c5.Fe7cc85cc5c09cf05a4e0c7ad9b427d7f3501c8-other.sqlite.bz2"},
			DatabaseVersion: 10,
			OpenSize:        6144,
			Timestamp:       1427333070,
			Size:            572,
		},
		{
			Type:            "other",
			Checksum:        Checksum{Type: "sha256", Sum: "6bf9672d0862e8ef8b8ff05a2fd0208a922b1f5978e6589d87944c88259cb670"},
			OpenChecksum:    Checksum{Type: "sha256", Sum: "e0ed5e0054194df036cf09c1a911e15bf2a4e7f26f2a788b6f47d53e80717ccc"},
			Location:        Location{Href: "repodata/6bf9672d0862e8ef8b8ff05a2fd0208a922b1f5978e6589d87944c88259cb670-other.xml.gz"},
			DatabaseVersion: 0,
			OpenSize:        121,
			Timestamp:       1427333070,
			Size:            123,
		},
		{
			Type:            "filelists_db",
			Checksum:        Checksum{Type: "sha256", Sum: "6174899a2d38869b23cab3b6308eb0c32ab3480b20aaca93226695c4684c4bf9"},
			OpenChecksum:    Checksum{Type: "sha256", Sum: "93e39c97d742867e8187e51deb3a7cdfe95a7e8dacdadde835a2ccd61ddf7e40"},
			Location:        Location{Href: "repodata/6174899a2d38869b23cab3b6308eb0c32ab3480b20aaca93226695c4684c4bf9-filelists.sqlite.bz2"},
			DatabaseVersion: 10,
			OpenSize:        7168,
			Timestamp:       1427333070,
			Size:            591,
		},
	},
}

func TestRepomdUnmarshal(t *testing.T) {
	r := new(repomd)
	err := xml.Unmarshal([]byte(repomdXML), &r)

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(r, repomdStruct) {
		t.Errorf("%v", pretty.Diff(r, repomdStruct))
	}

}

func TestRepomdMarshal(t *testing.T) {
	output, err := xml.MarshalIndent(repomdStruct, "", "  ")

	soutput := xml.Header + string(output)

	if err != nil {
		t.Error(err)
	}
	if soutput != repomdXML {
		t.Errorf("Incorrect marshall to XML.")
	}
}
