// loader.go
package loader

import (
	"LeetCodeDocumentor/models"
	"LeetCodeDocumentor/xmlhelper"
	"fmt"
	"os"
)

// getDefaultLanguages returns a list of default programming languages.
func getDefaultLanguages() []models.Language {
	return []models.Language{
		{Name: "Java", Extension: ".java", CommentStart: "/*", CommentEnd: "*/"},
		{Name: "Go", Extension: ".go", CommentStart: "/*", CommentEnd: "*/"},
		{Name: "Python", Extension: ".py", CommentStart: "\"\"\"", CommentEnd: "\"\"\""},
		{Name: "C", Extension: ".c", CommentStart: "/*", CommentEnd: "*/"},
		{Name: "Python 3", Extension: ".py", CommentStart: "\"\"\"", CommentEnd: "\"\"\""},
		{Name: "C++", Extension: ".cpp", CommentStart: "/*", CommentEnd: "*/"},
		{Name: "C#", Extension: ".cs", CommentStart: "/*", CommentEnd: "*/"},
		{Name: "JavaScript", Extension: ".js", CommentStart: "/*", CommentEnd: "*/"},
		{Name: "TypeScript", Extension: ".ts", CommentStart: "/*", CommentEnd: "*/"},
		{Name: "PHP", Extension: ".php", CommentStart: "/*", CommentEnd: "*/"},
		{Name: "Swift", Extension: ".swift", CommentStart: "/*", CommentEnd: "*/"},
		{Name: "Kotlin", Extension: ".kt", CommentStart: "/*", CommentEnd: "*/"},
		{Name: "Dart", Extension: ".dart", CommentStart: "/*", CommentEnd: "*/"},
		{Name: "Ruby", Extension: ".rb", CommentStart: "=begin", CommentEnd: "=end"},
		{Name: "Scala", Extension: ".scala", CommentStart: "/*", CommentEnd: "*/"},
		{Name: "Rust", Extension: ".rs", CommentStart: "/*", CommentEnd: "*/"},
		{Name: "Racket", Extension: ".rkt", CommentStart: "#|", CommentEnd: "|#"},
		{Name: "Erlang", Extension: ".erl", CommentStart: "%%", CommentEnd: "%%"},
		{Name: "Elixir", Extension: ".ex", CommentStart: "#", CommentEnd: "#"},
		{Name: "R", Extension: ".r", CommentStart: "#", CommentEnd: "#"},
	}
}

// InitializeLanguageXML initializes the Language XML file if it doesn't exist or adds missing languages.
func InitializeLanguageXML(filename string) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// File does not exist, initialize with all default languages
		languages := getDefaultLanguages()
		err := xmlhelper.InitializeWithLanguageList(filename, languages)
		if err != nil {
			return fmt.Errorf("failed to initialize language XML: %v", err)
		}
		fmt.Println("Language XML initialized successfully.")
	} else {
		// File exists, load existing languages and add missing ones
		languageList := models.LanguageList{}
		err := languageList.Load(filename)
		if err != nil {
			return fmt.Errorf("failed to load existing languages: %v", err)
		}

		defaultLanguages := getDefaultLanguages()
		missingLanguages := []models.Language{}

		// Create a map for easy lookup of existing languages
		existingMap := make(map[string]bool)
		for _, lang := range languageList.Languages {
			existingMap[lang.Name] = true
		}

		// Find missing languages
		for _, lang := range defaultLanguages {
			if !existingMap[lang.Name] {
				missingLanguages = append(missingLanguages, lang)
			}
		}

		// Add missing languages if any
		if len(missingLanguages) > 0 {
			for _, lang := range missingLanguages {
				languageList.Add(lang)
			}
			err = languageList.Save(filename)
			if err != nil {
				return fmt.Errorf("failed to add missing languages: %v", err)
			}
			fmt.Println("Missing languages added to Language XML.")
		} else {
			fmt.Println("No missing languages to add. Language XML is up to date.")
		}
	}
	return nil
}
