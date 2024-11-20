package collector

import (
	"net/http"
	"os"
	"time"	
)

// LogEntry represents a log message with timestamp and content(metadata).
type LogEntry struct {
	Timestamp  time.Time `json:"timestamp"`
	ServeiceID string    `json:"serviceID"`
	Level      string    `json:"level"`
	Message    string    `json:"message`
}

// Collector represents the log collector server.
type Collector struct {
	logFile *os.File
}


func NewCollector(logFilePath string) (*Collector, error) {
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o666) 
	if err != nil {
		return nil, err
	}
	// Explain how line 23 functions be detailed // 
	return &Collector{logFile: logFile}, nil
}

// WriteLog parses incoming JSON log data and writes the entry to the log file.
func (c * Collector) HandleLog(w http.ResponseWriter, r *http.Request) {
	var logEnty LogEntry

	// Decode the JSON body
	if err := json.NewDecoder(r.Body).Decode(&LogEntry); err != nil {
		http.Error(w, "Invalid log format", http.StatusBadRequest)
		return
	}
}


