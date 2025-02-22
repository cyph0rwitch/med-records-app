package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const filename = "records.txt"

func main() {
	for {
		fmt.Println("\nMedical Records App")
		fmt.Println("1. Add a new entry")
		fmt.Println("2. View all records")
		fmt.Println("3. Delete a record")
		fmt.Println("4. Move a record")
		fmt.Println("5. Update a record")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addEntry()
		case 2:
			viewRecords()
		case 3:
			deleteRecord()
		case 4:
			moveRecord()
		case 5:
			updateRecord()
		case 6:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func addEntry() {
	fmt.Print("Enter a new medical note: ")
	reader := bufio.NewReader(os.Stdin)
	note, _ := reader.ReadString('\n')
	note = strings.TrimSpace(note)

	// Append to the file
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(note + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}

	fmt.Println("Entry added successfully!")
}

func viewRecords() {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	if len(lines) == 0 || (len(lines) == 1 && lines[0] == "") {
		fmt.Println("No records found.")
		return
	}

	fmt.Println("\nStored Records:")
	for i, line := range lines {
		if line != "" {
			fmt.Printf("%d: %s\n", i+1, line)
		}
	}
}

func deleteRecord() {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	if len(lines) == 0 || (len(lines) == 1 && lines[0] == "") {
		fmt.Println("No records found.")
		return
	}

	// Display all records with index numbers
	fmt.Println("\nStored Records:")
	for i, line := range lines {
		if line != "" {
			fmt.Printf("%d: %s\n", i+1, line)
		}
	}

	fmt.Print("\nEnter the number of the record to delete: ")
	var num int
	fmt.Scanln(&num)

	if num < 1 || num > len(lines) || lines[num-1] == "" {
		fmt.Println("Invalid selection.")
		return
	}

	// Remove the selected entry
	lines = append(lines[:num-1], lines[num:]...)

	// Rewrite the file with the updated list
	err = os.WriteFile(filename, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		fmt.Println("Error updating file:", err)
		return
	}

	fmt.Println("Record deleted successfully!")
}

func moveRecord() {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	if len(lines) == 0 || (len(lines) == 1 && lines[0] == "") {
		fmt.Println("No records found.")
		return
	}

	// Display all records with index numbers
	fmt.Println("\nStored Records:")
	for i, line := range lines {
		if line != "" {
			fmt.Printf("%d: %s\n", i+1, line)
		}
	}

	fmt.Print("\nEnter the number of the record to move: ")
	var num int
	fmt.Scanln(&num)

	if num < 1 || num > len(lines) || lines[num-1] == "" {
		fmt.Println("Invalid selection.")
		return
	}

	// Ask for the new index to move the record to
	fmt.Print("\nEnter the index where you want to move the record: ")
	var newIndex int
	fmt.Scanln(&newIndex)

	if newIndex < 1 || newIndex > len(lines) || newIndex == num {
		fmt.Println("Invalid index.")
		return
	}

	// Remove the record from its current position
	record := lines[num-1]
	lines = append(lines[:num-1], lines[num:]...)

	// Insert the record at the new position (adjust for zero-based indexing)
	if newIndex > len(lines) {
		// If the new index is beyond the last position, append it at the end
		lines = append(lines, record)
	} else {
		// Insert the record at the desired index
		lines = append(lines[:newIndex-1], append([]string{record}, lines[newIndex-1:]...)...)
	}

	// Rewrite the file with the updated list
	err = os.WriteFile(filename, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		fmt.Println("Error updating file:", err)
		return
	}

	fmt.Println("Record moved successfully!")
}
func updateRecord() {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	if len(lines) == 0 || (len(lines) == 1 && lines[0] == "") {
		fmt.Println("No records found.")
		return
	}

	// Display all records with index numbers
	fmt.Println("\nStored Records:")
	for i, line := range lines {
		if line != "" {
			fmt.Printf("%d: %s\n", i+1, line)
		}
	}

	fmt.Print("\nEnter the number of the record to update: ")
	var num int
	fmt.Scanln(&num)

	if num < 1 || num > len(lines) || lines[num-1] == "" {
		fmt.Println("Invalid selection.")
		return
	}

	// Prompt the user for a new record
	fmt.Print("Enter the new value for the record: ")
	var newRecord string
	reader := bufio.NewReader(os.Stdin)
	newRecord, _ = reader.ReadString('\n')
	newRecord = strings.TrimSpace(newRecord)

	// Update the selected record
	lines[num-1] = newRecord

	// Rewrite the file with the updated list
	err = os.WriteFile(filename, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		fmt.Println("Error updating file:", err)
		return
	}

	fmt.Println("Record updated successfully!")
}
