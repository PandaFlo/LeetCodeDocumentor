package xmlhelper

import (
	"LeetCodeDocumentor/models"
)

// InitializeWithLanguage creates an XML file with a single language.
func InitializeWithLanguage(filename string, language models.Language) error {
	languageList := models.LanguageList{
		Languages: []models.Language{language},
	}
	return languageList.Save(filename)
}

// InitializeWithLanguageList creates an XML file with a given list of languages.
func InitializeWithLanguageList(filename string, languages []models.Language) error {
	languageList := models.LanguageList{
		Languages: languages,
	}
	return languageList.Save(filename)
}
