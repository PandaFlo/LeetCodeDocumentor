package filebuilder

import (
	"LeetCodeDocumentor/models"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GenerateLeetCodeDocumentation creates the necessary folder structure and files for a LeetCode question
func GenerateLeetCodeDocumentation(basePath string, question models.Question, languages []models.Language, username string) error {
	// Step 1: Create the main folder for the LeetCode question at the given file path
	leetCodeFolderPath := filepath.Join(basePath, fmt.Sprintf("LeetCode#%d", question.Number))
	err := os.MkdirAll(leetCodeFolderPath, 0755)
	if err != nil {
		return fmt.Errorf("failed to create LeetCode folder: %w", err)
	}

	// Step 2: Generate the result.doc file with question and solution info
	err = CreateResultDoc(leetCodeFolderPath, question, languages, username)
	if err != nil {
		return fmt.Errorf("failed to create result.doc: %w", err)
	}

	// Step 3: Create the Languages subfolder and individual language folders
	for _, lang := range languages {
		// Ensure that the extension does not include an extra dot
		if !strings.HasPrefix(lang.Extension, ".") {
			lang.Extension = "." + lang.Extension // Ensure the extension starts with a dot
		}

		// Create the folder for each language
		languageFolderPath := filepath.Join(leetCodeFolderPath, "Languages", lang.Name)
		err := os.MkdirAll(languageFolderPath, 0755)
		if err != nil {
			return fmt.Errorf("failed to create language folder for %s: %w", lang.Name, err)
		}

		// Step 4: Generate the question file in the corresponding folder
		fileName := fmt.Sprintf("question_%d%s", question.Number, lang.Extension) // File name with correct extension
		filePath := filepath.Join(languageFolderPath, fileName)

		err = GenerateFile(filePath, lang, question)
		if err != nil {
			return fmt.Errorf("failed to create file for language %s: %w", lang.Name, err)
		}
	}

	return nil
}
