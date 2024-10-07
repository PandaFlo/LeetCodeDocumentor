package filebuilder

import (
	"LeetCodeDocumentor/models"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// GenerateFile creates a file for each question with the given format at the specified file destination
func GenerateFile(filePath string, lang models.Language, question models.Question) error {
	// Create the file at the specified file path
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Write Question Title with Language in comments
	_, err = file.WriteString(fmt.Sprintf("%s Question #%d: %s (%s) %s\n",
		lang.CommentStart, question.Number, question.Title, lang.Name, lang.CommentEnd))
	if err != nil {
		return fmt.Errorf("failed to write title: %w", err)
	}

	// Write Question Description in comments
	_, err = file.WriteString(fmt.Sprintf("%s %s %s\n\n", lang.CommentStart, question.Question, lang.CommentEnd))
	if err != nil {
		return fmt.Errorf("failed to write question: %w", err)
	}

	// Space for an answer
	_, err = file.WriteString("\n\n")
	if err != nil {
		return fmt.Errorf("failed to write space for answer: %w", err)
	}

	// Write Solution in comments
	_, err = file.WriteString(fmt.Sprintf("%s Solution: %s %s\n", lang.CommentStart, question.Solution, lang.CommentEnd))
	if err != nil {
		return fmt.Errorf("failed to write solution: %w", err)
	}

	return nil
}

// CreateLeetCodeFolder creates a folder titled "LeetCode#Number" and a "Languages" subfolder
func CreateLeetCodeFolder(question models.Question) (string, error) {
	folderName := fmt.Sprintf("LeetCode#%d", question.Number)
	err := os.Mkdir(folderName, 0755)
	if err != nil {
		return "", fmt.Errorf("failed to create folder: %w", err)
	}

	// Create the "Languages" subfolder
	languagesFolder := filepath.Join(folderName, "Languages")
	err = os.Mkdir(languagesFolder, 0755)
	if err != nil {
		return "", fmt.Errorf("failed to create Languages folder: %w", err)
	}

	return folderName, nil
}

// CreateLanguageFolder creates a folder for a specific language in the "Languages" folder
func CreateLanguageFolder(basePath string, language models.Language) error {
	languageFolder := filepath.Join(basePath, "Languages", language.Name)
	err := os.Mkdir(languageFolder, 0755)
	if err != nil {
		return fmt.Errorf("failed to create language folder: %w", err)
	}
	return nil
}

// CreateResultDoc generates a "result.doc" file with details like username, date, question, solution, and completed languages
// CreateResultDoc generates a "result.doc" file with details like username, date, question, solution, and completed languages
func CreateResultDoc(folderName string, question models.Question, languages []models.Language, username string) error {
	// Create the result.doc file in the main folder
	resultFilePath := filepath.Join(folderName, "result.doc")
	file, err := os.Create(resultFilePath)
	if err != nil {
		return fmt.Errorf("failed to create result.doc: %w", err)
	}
	defer file.Close()

	// Get the current date in MM/DD/YYYY format
	currentDate := time.Now().Format("01/02/2006")

	// Write the information to result.doc
	_, err = file.WriteString(fmt.Sprintf("Username: %s\n", username))
	if err != nil {
		return fmt.Errorf("failed to write username: %w", err)
	}

	_, err = file.WriteString(fmt.Sprintf("Date: %s\n\n", currentDate))
	if err != nil {
		return fmt.Errorf("failed to write date: %w", err)
	}

	_, err = file.WriteString(fmt.Sprintf("LeetCode Question #%d\n", question.Number))
	if err != nil {
		return fmt.Errorf("failed to write question number: %w", err)
	}

	_, err = file.WriteString(fmt.Sprintf("Question: %s\n\n", question.Question))
	if err != nil {
		return fmt.Errorf("failed to write question: %w", err)
	}

	_, err = file.WriteString(fmt.Sprintf("Solution: %s\n\n", question.Solution))
	if err != nil {
		return fmt.Errorf("failed to write solution: %w", err)
	}

	// Write the list of completed languages
	languageNames := []string{}
	for _, lang := range languages {
		languageNames = append(languageNames, lang.Name)
	}
	_, err = file.WriteString(fmt.Sprintf("Completed Languages: %s\n", strings.Join(languageNames, ", ")))
	if err != nil {
		return fmt.Errorf("failed to write languages: %w", err)
	}

	return nil
}
