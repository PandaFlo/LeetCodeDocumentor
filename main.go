package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Define a global variable for the username
var username = "Malik Maitland"

func main() {
	// Load predefined languages from the XML file
	filename := "languages.xml"
	predefinedLanguages := getPredefinedLanguagesFromFile(filename)

	myApp := app.NewWithID("com.maitland.LeetCodeDocumentor")
	myWindow := myApp.NewWindow("LeetCode Documentor")

	// Set the icon for the application
	iconFile := "./icon.png"
	if _, err := os.Stat(iconFile); err == nil {
		icon := fyne.NewStaticResource("icon.png", loadIcon(iconFile))
		myWindow.SetIcon(icon)
	} else {
		// Fallback to a built-in theme icon if the file doesn't exist
		myWindow.SetIcon(theme.FyneLogo())
	}

	// Create GUI elements for question, answer, folder, and language selection
	questionEntry, topicContainer, folderSelectionContent, languageSelectionContent, answerEntry := createGUIElements(myWindow, filename, predefinedLanguages)

	// Combine the content for the controls panel, adding padding for the controls section
	combinedContent := container.NewVBox(
		widget.NewLabel("    "),  // Left-side padding for the controls section
		topicContainer,           // Content aligned and padded
		widget.NewLabel("    "),  // Extra padding
		folderSelectionContent,   // Folder path
		widget.NewLabel("    "),  // Extra padding
		languageSelectionContent, // Language selection section
	)

	// Add left padding to the entire controls panel
	paddedContent := container.NewHBox(
		widget.NewLabel(""), // Left padding
		combinedContent,     // Controls panel with all content and padding
	)

	// Step 1: Create a horizontal split between the question and answer sections at a 50:50 ratio
	questionAnswerSplit := container.NewHSplit(
		container.NewVScroll(questionEntry), // Left panel (question input)
		container.NewVScroll(answerEntry),   // Middle panel (answer input)
	)
	questionAnswerSplit.SetOffset(0.5) // Split equally between question and answer

	// Step 2: Create the final split between question/answer and the padded controls panel
	finalSplit := container.NewHSplit(
		questionAnswerSplit, // Left panel (question and answer input)
		paddedContent,       // Right panel (controls with padding)
	)
	finalSplit.SetOffset(0.75) // 75% for question/answer, 25% for the controls

	// Set the final layout as the content of the window
	myWindow.SetContent(finalSplit)
	myWindow.Resize(fyne.NewSize(1000, 600))
	myWindow.ShowAndRun()
}

// Helper function to load an icon from a file
func loadIcon(path string) []byte {
	data, _ := os.ReadFile(path)
	return data
}
