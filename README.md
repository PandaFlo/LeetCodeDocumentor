# LeetCode Documentor

The LeetCode Documentor is a desktop application built with the **Fyne** GUI framework in Go. It assists developers in documenting LeetCode solutions in multiple programming languages by allowing them to input the problem, select relevant programming languages, and generate documentation files. Below are the core features, setup instructions, and future enhancements.

The goal of the LeetCode Documentor project is to provide a practical way to improve programming skills while contributing to your Git repository. By documenting solutions for each problem, you not only practice coding but also maintain a clear record of your progress and problem-solving approaches. This project ensures consistent contributions to your Git, enhancing your portfolio while helping you refine your skills in various languages, structuring code effectively, and managing a well-organized development process.

## Features

### 1. **User-friendly GUI**: 
   The application provides an intuitive interface where users can easily input their LeetCode questions, answers, and select programming languages for documentation. The interface is divided into sections for easy navigation, including panels for questions, answers, folder selection, and language settings.

### 2. **Language Selection**:
   Users can select multiple languages from a predefined list or add custom languages for documentation. The language selection is facilitated by two panels, and a "+" button allows adding new languages, while a "-" button removes any selected language.

### 3. **Path Selection**:
   A built-in file dialog allows users to choose where to save the generated documentation. This simplifies managing and organizing the LeetCode problem files into corresponding folders.

### 4. **Documentation Generation**:
   Upon providing all inputs (question number, topic, question, answer, and selected languages), users can generate LeetCode documentation in structured folders. The app automatically creates subfolders for each selected language and generates solution files in the appropriate format (e.g., `.java`, `.py`).

### 5. **Predefined Language Management**:
   The application allows adding or removing predefined programming languages from an XML configuration file. This ensures that the list of available languages stays up-to-date and customizable.

## Requirements

- Go language installed on your system.
- Fyne library for building the graphical user interface.

### Install Fyne:

```bash
go get fyne.io/fyne/v2
```

## Installation and Setup

1. **Clone the Repository**:
   Clone the LeetCode Documentor repository to your local machine.
   
   ```bash
   git clone https://github.com/MalikMaitland/LeetCodeDocumentor.git
   ```

2. **Navigate to the Project Directory**:
   ```bash
   cd LeetCodeDocumentor
   ```

3. **Run the Application**:
   Use the following command to run the application:

   ```bash
   go run .
   ```

## How to Use

1. **Input Your Question**: 
   - Enter the LeetCode problem number and topic in the provided fields.
   - Write the problem description and your solution in the respective sections.

2. **Select a Folder**: 
   - Click the "Select a Path" link to choose a folder where your documentation will be saved.

3. **Choose Programming Languages**: 
   - Use the checkboxes to select one or more programming languages from the predefined list.
   - Optionally, click the "+" button to add new languages or the "-" button to remove languages.

4. **Generate Documentation**:
   - Click the "Run" button to generate the documentation files in the selected folder.
   - A success message will appear once the documentation has been successfully generated.

## Customization

### Add New Programming Languages
You can modify the list of available programming languages by editing the `languages.xml` file. To add a new language:
- Open the XML file.
- Add a new `<Language>` element with the `Name`, `Extension`, `CommentStart`, and `CommentEnd` tags.

### Modify the GUI Layout
You can adjust the layout or functionality by modifying the `main.go` file, which uses the **Fyne** framework to manage the window, panels, and controls.

## Future Enhancements

- **Code Editor Integration**: Add a syntax-highlighted code editor for writing solutions directly within the application.
- **Auto-Save**: Automatically save progress periodically to avoid data loss.
- **Template Support**: Pre-fill templates for common programming languages with necessary boilerplate code.
- **File Export Options**: Add functionality to export generated files as zip archives for easy sharing.

---

This updated setup ensures that users can efficiently document their LeetCode solutions in multiple languages, making it easier to organize and track their coding progress. The use of the `go run .` command provides a simpler way to run the application without specifying individual files. The application remains lightweight and cross-platform due to the Fyne framework.
