package audit

import (
	"encoding/json"
	"log"
	"time"
)

type AuditEvent struct {
	Service   string                 `json:"service"`
	Action    string                 `json:"action"`
	Entity    string                 `json:"entity"`
	EntityID  string                 `json:"entity_id"`
	UserID    string                 `json:"user_id"`
	Timestamp time.Time              `json:"timestamp"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

func LogEvent(event AuditEvent) {
	event.Timestamp = time.Now().UTC()
	data, err := json.Marshal(event)
	if err != nil {
		log.Printf("audit log marshal error: %v", err)
		return
	}

	log.Println(string(data))
}
