package main

import (
	"LeetCodeDocumentor/filebuilder" // Assuming this contains GenerateLeetCodeDocumentation
	"LeetCodeDocumentor/models"
	"image/color"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Path and language selection containers (global)
var (
	languageSelectionContainerLeft  = container.NewVBox()
	languageSelectionContainerRight = container.NewVBox()
	selectedPath                    string
	answerEntry                     *widget.Entry // Declare answerEntry as a global variable
)

func createGUIElements(myWindow fyne.Window, filename string, predefinedLanguages []models.Language) (*widget.Entry, *fyne.Container, *fyne.Container, *fyne.Container, *widget.Entry) {
	// Create the main text box for the question input with fixed size
	questionEntry := widget.NewMultiLineEntry()
	questionEntry.SetPlaceHolder("Enter your question here...")
	questionEntry.SetMinRowsVisible(10) // Fixed height for the question box

	// Create the "Topic" label as a title with bold style and white color
	topicTitle := canvas.NewText("Topic", color.White)
	topicTitle.TextStyle = fyne.TextStyle{Bold: true}
	topicTitle.TextSize = 18 // Make the title larger

	// Create two text boxes for Q# and Topic
	numberEntry := widget.NewEntry()
	numberEntry.SetPlaceHolder("Q#")
	topicEntry := widget.NewEntry()
	topicEntry.SetPlaceHolder("Enter topic here...")

	// Set the width for Q# (1/4) and Topic (3/4)
	grid := container.New(layout.NewGridWrapLayout(fyne.NewSize(200, numberEntry.MinSize().Height)), numberEntry, topicEntry)

	// Align the topic and question boxes by placing them in a vertical box
	topicContainer := container.NewVBox(
		topicTitle, // "Topic" title label
		grid,       // Row with Q# and Topic
	)

	// Create a folder path input section (for folder selection)
	folderTitle := canvas.NewText("Folder Location", color.White)
	folderTitle.TextStyle = fyne.TextStyle{Bold: true}
	folderTitle.TextSize = 18

	// Create a clickable text without an underline
	pathLabel := canvas.NewText("Select a Path", color.RGBA{R: 0, G: 122, B: 255, A: 255}) // Blue text
	pathLabel.TextStyle = fyne.TextStyle{Bold: true}                                       // Bold for emphasis

	// Wrap the clickable text in a container to simulate a hyperlink
	clickablePathContainer := widget.NewHyperlink(pathLabel.Text, nil)
	clickablePathContainer.OnTapped = func() {
		// Open folder dialog when clicked
		dialog.NewFolderOpen(func(list fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, myWindow)
				return
			}
			if list == nil {
				return
			}
			fullPath := list.Path()
			selectedPath = fullPath
			shortPath := ShortenPath(fullPath)
			clickablePathContainer.SetText(shortPath)
			// Remove the URL as it's no longer needed
			clickablePathContainer.SetURL(nil)
			clickablePathContainer.Refresh()
		}, myWindow).Show()
	}

	// Folder selection content aligned with proper padding
	folderSelectionContent := container.NewVBox(
		folderTitle,            // "Folder Location" title
		clickablePathContainer, // The clickable "Select a Path" text
	)

	// Create the default language selection checkboxes in two rows
	languageSelectionContainerLeft = container.NewVBox()
	languageSelectionContainerRight = container.NewVBox()
	languageContainers := container.NewHBox(languageSelectionContainerLeft, languageSelectionContainerRight) // Ensure proper layout without stretching

	populateLanguageCheckboxes(predefinedLanguages) // Display languages

	// Create the "+" button for adding more languages
	addLanguageButton := widget.NewButton(" + ", func() {
		showAddLanguageDialog(myWindow, filename)
	})

	// Create the "-" button for removing displayed languages
	removeLanguageButton := widget.NewButton(" - ", func() {
		showRemoveLanguageDialog(myWindow)
	})

	// Create the "Run" button for running the selected languages
	runButton := widget.NewButton("Run", func() {
		// On run, collect the question, title, solution, languages, and path
		questionNumber := numberEntry.Text
		topic := topicEntry.Text
		question := questionEntry.Text
		answer := answerEntry.Text

		// Validate input fields
		if questionNumber == "" || topic == "" || question == "" || answer == "" || selectedPath == "" {
			dialog.ShowInformation("Missing Information", "Please complete all fields before running.", myWindow)
			return
		}

		// Create the Question object from the entries
		questionObject := models.Question{
			Number:   stringToInt(questionNumber),
			Title:    topic,
			Question: question,
			Solution: answer,
		}

		// Get the selected languages
		selectedLanguages := getSelectedLanguages(predefinedLanguages)

		if len(selectedLanguages) == 0 {
			dialog.ShowInformation("No Languages Selected", "Please select at least one language.", myWindow)
			return
		}

		// Run the GenerateLeetCodeDocumentation function
		err := filebuilder.GenerateLeetCodeDocumentation(filepath.Join(selectedPath), questionObject, selectedLanguages, username)
		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}

		dialog.ShowInformation("Success", "Documentation generated successfully!", myWindow)
	})

	// Create a container for the language selection section
	languageSelectionTitle := canvas.NewText("Language Selection", color.White)
	languageSelectionTitle.TextStyle = fyne.TextStyle{Bold: true}
	languageSelectionTitle.TextSize = 18

	// The row of buttons, with "Run" aligned to the left and "+" and "-" aligned to the right
	buttonRow := container.NewHBox(
		runButton,            // "Run" button on the left
		layout.NewSpacer(),   // Spacer to push "+" and "-" to the right
		addLanguageButton,    // "+" button
		removeLanguageButton, // "-" button
	)

	// Language selection content with the button row at the bottom
	languageSelectionContent := container.NewVBox(
		languageSelectionTitle,
		languageContainers,
		widget.NewLabel("    "), // Extra padding
		buttonRow,               // Row with "Run" on the left, "+" and "-" on the right
	)

	// Create the answer entry section
	answerEntry = widget.NewMultiLineEntry() // Use the global answerEntry
	answerEntry.SetPlaceHolder("Enter your answer here...")
	answerEntry.SetMinRowsVisible(10) // Fixed height for the answer box

	return questionEntry, topicContainer, folderSelectionContent, languageSelectionContent, answerEntry
}
