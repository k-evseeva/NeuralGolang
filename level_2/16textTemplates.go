package main

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"strings"
)

func main_16() {

	// Define template
	tmpl, err := template.New("example").Parse("Welcome, {{.name}}! How are you doing?\n")
	if err != nil {
		panic(err)
	}

	// The same template (but 1 str instead of 4)
	// tmpl := template.Must(template.New("example").Parse("Welcome, {{.name}}! How are you doing?\n"))

	// Define data for template
	data := map[string]interface{}{
		"name": "John",
	}

	// Output data
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	// Create console app
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// Define named templates for different types
	templates := map[string]string{
		"welcome":      "Welcome, {{.name}}! We're glad you joined!",
		"notification": "{{.nm}}, you have a new notification: {{.ntf}}",
		"error":        "Oops! An error occured: {{.em}}",
	}

	// Parse and store templates
	parsedTemplates := make(map[string]*template.Template)
	for name, tmpl := range templates {
		parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}

	// Define app
	for {
		// Show menu
		fmt.Println("\nMenu:")
		fmt.Println("1. Join")
		fmt.Println("2. Get notification")
		fmt.Println("3. Get error")
		fmt.Println("4. Exit")
		fmt.Println("Choose an option: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]interface{}
		var tmpl *template.Template

		switch choice {
		case "1":
			tmpl = parsedTemplates["welcome"]
			data = map[string]interface{}{"name": name}
		case "2":
			fmt.Println("Enter your notification message: ")
			notification, _ := reader.ReadString('\n')
			notification = strings.TrimSpace(notification)
			tmpl = parsedTemplates["notification"]
			data = map[string]interface{}{"nm": name, "ntf": notification}
		case "3":
			fmt.Println("Enter your error message: ")
			errorMessage, _ := reader.ReadString('\n')
			errorMessage = strings.TrimSpace(errorMessage)
			tmpl = parsedTemplates["error"]
			data = map[string]interface{}{"nm": name, "em": errorMessage}
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please, select correct option.")
			continue
		}
		// Render and print template to console
		err := tmpl.Execute(os.Stdout, data)
		if err != nil {
			fmt.Println("Error executing template: ", err)
		}
	}

}
