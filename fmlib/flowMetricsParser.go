package fmlib

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

// WorkItem represents a single work item with dynamic state transitions
type WorkItem struct {
	ID            string
	Name          string
	CreatedDate   time.Time
	StartedDate   time.Time
	FinishedDate  time.Time
	StateChanges  map[string]time.Time // Key: State Name, Value: Date Entered
	CycleTimeDays float64
}

// ReadCSV reads a CSV file and returns a slice of WorkItems
func ReadCSV(filename string) ([]WorkItem, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) < 2 {
		return nil, fmt.Errorf("CSV file must have at least a header and one row of data")
	}

	// Identify columns dynamically
	headers := records[0]
	workItems := []WorkItem{}

	for _, row := range records[1:] {
		item := WorkItem{
			ID:           row[0],
			Name:         row[1],
			StateChanges: make(map[string]time.Time),
		}

		// Convert dates
		for i, col := range headers {
			if date, err := parseDate(row[i]); err == nil {
				switch col {
				case "CreatedDate":
					item.CreatedDate = date
				case "StartedDate":
					item.StartedDate = date
				case "FinishedDate":
					item.FinishedDate = date
				default:
					item.StateChanges[col] = date
				}
			}
		}

		// Calculate Cycle Time (if FinishedDate exists)
		if !item.CreatedDate.IsZero() && !item.FinishedDate.IsZero() {
			item.CycleTimeDays = item.FinishedDate.Sub(item.StartedDate).Hours() / 24
			if item.CycleTimeDays == 0 {
				item.CycleTimeDays = 1 //Cycle time for a completed item is always no less than 1 day
			}
		}

		workItems = append(workItems, item)
	}

	return workItems, nil
}

// parseDate converts string to time.Time
func parseDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr) // Adjust format as per CSV
}

func Parser(fileName string) []WorkItem {
	WorkItems, err := ReadCSV(fileName)
	if err != nil {
		log.Fatal("Failed to read CSV: %v", err)
	}

	return WorkItems
}
