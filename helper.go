// helpers.go
package main

import (
	"LeetCodeDocumentor/models"
	"LeetCodeDocumentor/xmlhelper"
	"fmt"
	"os"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// Helper function to load an icon from a file
func loadIcon(path string) []byte {
	data, _ := os.ReadFile(path)
	return data
}

// ShortenPath reduces the length of a file path for display purposes
func ShortenPath(path string) string {
	const maxLength = 50

	if len(path) <= maxLength {
		return path
	}

	// Replace backslashes with forward slashes for uniform processing
	path = strings.ReplaceAll(path, "\\", "/")

	parts := strings.Split(path, "/")
	if len(parts) <= 3 {
		return path
	}

	// Keep the first part (usually drive letter or root) and last two parts
	shortened := parts[0] + " .../" + strings.Join(parts[len(parts)-2:], "/")

	// If it's still too long, truncate the last part
	if len(shortened) > maxLength {
		lastPart := parts[len(parts)-1]
		if len(lastPart) > 10 {
			lastPart = lastPart[:7] + "..."
		}
		shortened = parts[0] + " .../" + parts[len(parts)-2] + "/" + lastPart
	}

	return shortened
}

// stringToInt converts a string to an integer safely
func stringToInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0 // Return 0 or handle error appropriately
	}
	return val
}

// getSelectedLanguages returns a list of selected languages
// getSelectedLanguages returns a list of selected languages
func getSelectedLanguages(predefinedLanguages []models.Language) []models.Language {
	selectedLanguages := []models.Language{}

	// Iterate over checkboxes in both left and right containers
	for _, obj := range append(languageSelectionContainerLeft.Objects, languageSelectionContainerRight.Objects...) {
		if check, ok := obj.(*widget.Check); ok && check.Checked {
			// Find the corresponding language based on the checkbox name
			for _, lang := range predefinedLanguages {
				if lang.Name == check.Text {
					selectedLanguages = append(selectedLanguages, lang)
				}
			}
		}
	}

	return selectedLanguages
}

// getPredefinedLanguagesFromFile loads predefined languages from an XML file
func getPredefinedLanguagesFromFile(filename string) []models.Language {
	var predefinedLanguages models.LanguageList
	err := predefinedLanguages.Load(filename)
	if err != nil {
		fmt.Printf("Failed to load predefined languages: %v\n", err)
		return []models.Language{}
	}
	return predefinedLanguages.Languages
}

// populateLanguageCheckboxes populates the left and right containers with language checkboxes
func populateLanguageCheckboxes(predefinedLanguages []models.Language) {
	languageSelectionContainerLeft.Objects = nil  // Clear existing checkboxes in left container
	languageSelectionContainerRight.Objects = nil // Clear existing checkboxes in right container

	// Only display the first 4 languages initially
	for i, lang := range predefinedLanguages {
		if i >= 4 {
			break
		}
		languageCheck := widget.NewCheck(lang.Name, nil)
		languageSelectionContainerLeft.Add(languageCheck)
	}

	// Refresh the containers after adding checkboxes
	languageSelectionContainerLeft.Refresh()
	languageSelectionContainerRight.Refresh()
}

// addLanguageToContainers adds a new language checkbox to the appropriate container
// addLanguageToContainers adds a new language checkbox to the appropriate container
func addLanguageToContainers(newLanguage models.Language) {
	languageCheck := widget.NewCheck(newLanguage.Name, nil) // Initialize with unchecked state

	// Add the new language checkbox to the correct container (left or right)
	if len(languageSelectionContainerLeft.Objects) <= len(languageSelectionContainerRight.Objects) {
		languageSelectionContainerLeft.Add(languageCheck)
	} else {
		languageSelectionContainerRight.Add(languageCheck)
	}

	// Set the checkbox change listener
	languageCheck.OnChanged = func(checked bool) {
		// Now we handle the checked state:
		if checked {

		}
	}

	// Refresh the containers after adding the new language
	languageSelectionContainerLeft.Refresh()
	languageSelectionContainerRight.Refresh()
}

// showRemoveLanguageDialog allows you to remove a displayed language by selecting from a list of the currently displayed ones.
func showRemoveLanguageDialog(window fyne.Window) {
	currentLanguages := []string{}
	// Collect all currently displayed languages (in both left and right containers)
	for _, obj := range append(languageSelectionContainerLeft.Objects, languageSelectionContainerRight.Objects...) {
		if check, ok := obj.(*widget.Check); ok {
			currentLanguages = append(currentLanguages, check.Text)
		}
	}

	if len(currentLanguages) == 0 {
		dialog.ShowInformation("No Languages", "No languages available to remove.", window)
		return
	}

	// Create a drop-down selection for displayed languages
	languageSelect := widget.NewSelect(currentLanguages, nil)
	languageSelect.PlaceHolder = "Select language to remove"

	dialog.ShowCustomConfirm("Remove Language", "Remove", "Cancel", container.NewVBox(languageSelect), func(confirmed bool) {
		if confirmed {
			selected := languageSelect.Selected
			if selected != "" {
				found := false

				// Search in left container and remove if found
				for i, obj := range languageSelectionContainerLeft.Objects {
					if check, ok := obj.(*widget.Check); ok && check.Text == selected {
						languageSelectionContainerLeft.Objects = append(languageSelectionContainerLeft.Objects[:i], languageSelectionContainerLeft.Objects[i+1:]...)
						found = true
						break
					}
				}

				// If not found in left, search and remove from the right container
				if !found {
					for i, obj := range languageSelectionContainerRight.Objects {
						if check, ok := obj.(*widget.Check); ok && check.Text == selected {
							languageSelectionContainerRight.Objects = append(languageSelectionContainerRight.Objects[:i], languageSelectionContainerRight.Objects[i+1:]...)
							break
						}
					}
				}

				// Adjust the layout after removal
				for len(languageSelectionContainerLeft.Objects) < 4 && len(languageSelectionContainerRight.Objects) > 0 {
					obj := languageSelectionContainerRight.Objects[0]
					languageSelectionContainerRight.Objects = languageSelectionContainerRight.Objects[1:]
					languageSelectionContainerLeft.Add(obj)
				}

				languageSelectionContainerLeft.Refresh()
				languageSelectionContainerRight.Refresh()
			} else {
				dialog.ShowInformation("No Selection", "Please select a language to remove.", window)
			}
		}
	}, window)
}

// showAddLanguageDialog displays a dialog to add a new language
// showAddLanguageDialog displays a dialog to add a new language
func showAddLanguageDialog(window fyne.Window, filename string) {
	// Create entries for language details
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Language Name")

	extEntry := widget.NewEntry()
	extEntry.SetPlaceHolder("File Extension (e.g., .java)")

	commentOpenEntry := widget.NewEntry()
	commentOpenEntry.SetPlaceHolder("Multiline Comment Start (e.g., /*)")

	commentCloseEntry := widget.NewEntry()
	commentCloseEntry.SetPlaceHolder("Multiline Comment End (e.g., */)")

	// Load predefined languages from the XML file
	predefinedLanguages := getPredefinedLanguagesFromFile(filename)

	// Remove already displayed languages from predefinedLanguages
	currentDisplayedLanguages := make(map[string]bool)
	for _, obj := range append(languageSelectionContainerLeft.Objects, languageSelectionContainerRight.Objects...) {
		if check, ok := obj.(*widget.Check); ok {
			currentDisplayedLanguages[check.Text] = true
		}
	}

	filteredLanguages := []models.Language{}
	for _, lang := range predefinedLanguages {
		if !currentDisplayedLanguages[lang.Name] {
			filteredLanguages = append(filteredLanguages, lang)
		}
	}

	// Create a list of language names for predefined languages that aren't displayed
	languageNames := []string{}
	for _, lang := range filteredLanguages {
		languageNames = append(languageNames, lang.Name)
	}

	var languagePopup *widget.PopUp
	languageList := widget.NewList(
		func() int { return len(languageNames) },
		func() fyne.CanvasObject { return widget.NewLabel("template") },
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(languageNames[id])
		},
	)

	// Track if the language is selected from the predefined list
	var selectedFromPredefined bool = false

	// Handle item selection for predefined languages
	languageList.OnSelected = func(id widget.ListItemID) {
		if id >= 0 && id < len(filteredLanguages) {
			selectedLang := filteredLanguages[id]
			nameEntry.SetText(selectedLang.Name)
			extEntry.SetText(selectedLang.Extension)
			commentOpenEntry.SetText(selectedLang.CommentStart)
			commentCloseEntry.SetText(selectedLang.CommentEnd)
			languagePopup.Hide()
			selectedFromPredefined = true // Mark as selected from predefined list
		}
	}

	// Wrap the list in a scroll container with a fixed size
	listScroll := container.NewVScroll(languageList)
	listScroll.SetMinSize(fyne.NewSize(200, 200))

	// Create the popup to display the list
	languagePopup = widget.NewModalPopUp(
		listScroll,
		window.Canvas(),
	)

	// Create a button to open the popup
	languageButton := widget.NewButton("Select Predefined Language", func() {
		if len(languageNames) == 0 {
			dialog.ShowInformation("No Languages", "No predefined languages available to select.", window)
			return
		}
		languagePopup.Show()
	})

	// Create the form with the button for selecting a predefined language or manual entry
	form := widget.NewForm(
		widget.NewFormItem("Predefined Language", languageButton),
		widget.NewFormItem("Name", nameEntry),
		widget.NewFormItem("Extension", extEntry),
		widget.NewFormItem("Comment Start", commentOpenEntry),
		widget.NewFormItem("Comment End", commentCloseEntry),
	)

	// Optionally, wrap the form in a Scroll container if the dialog content is large
	scroll := container.NewScroll(form)
	scroll.SetMinSize(fyne.NewSize(400, 300))

	// Show the custom confirm dialog with the scrollable form
	dialog.ShowCustomConfirm("Add New Language", "Add", "Cancel", scroll, func(confirmed bool) {
		if confirmed {
			name := nameEntry.Text
			ext := extEntry.Text
			commentStart := commentOpenEntry.Text
			commentEnd := commentCloseEntry.Text

			if name != "" && ext != "" && commentStart != "" && commentEnd != "" {
				// Check if the language is already displayed
				if currentDisplayedLanguages[name] {
					dialog.ShowInformation("Duplicate Language", "This language is already displayed.", window)
					return
				}

				// Create a new language object
				newLanguage := models.Language{
					Name:         name,
					Extension:    ext,
					CommentStart: commentStart,
					CommentEnd:   commentEnd,
				}

				// Add the new language to the displayed UI (checkboxes)
				addLanguageToContainers(newLanguage)

				// If the language was not selected from the predefined list, add it to the XML
				if !selectedFromPredefined {
					// Add the new language to the XML
					predefinedLanguages = append(predefinedLanguages, newLanguage)
					err := xmlhelper.InitializeWithLanguageList(filename, predefinedLanguages)
					if err != nil {
						dialog.ShowError(err, window)
					}
				}

			} else {
				dialog.ShowInformation("Incomplete Data", "Please fill in all fields.", window)
			}
		}
	}, window)
}
