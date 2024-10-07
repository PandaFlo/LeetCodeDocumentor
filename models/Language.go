package models

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

type Language struct {
	Name         string `xml:"name"`
	Extension    string `xml:"extension"`
	CommentStart string `xml:"commentStart"`
	CommentEnd   string `xml:"commentEnd"`
}

// LanguageList represents a list of languages.
type LanguageList struct {
	XMLName   xml.Name   `xml:"languages"`
	Languages []Language `xml:"language"`
}

// Add adds a new language to the list.
func (ll *LanguageList) Add(language Language) {
	ll.Languages = append(ll.Languages, language)
}

// Delete removes a language from the list by its name.
func (ll *LanguageList) Delete(name string) {
	for i, lang := range ll.Languages {
		if lang.Name == name {
			ll.Languages = append(ll.Languages[:i], ll.Languages[i+1:]...)
			return
		}
	}
}

// Load reads the XML file and loads the languages into the LanguageList.
func (ll *LanguageList) Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	return xml.Unmarshal(byteValue, ll)
}

// Save writes the LanguageList to an XML file.
func (ll *LanguageList) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	output, err := xml.MarshalIndent(ll, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(output)
	return err
}
