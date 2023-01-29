package types

// Struct holding valid project task status values
var Status = struct {
	IN_PROGRESS string
	DONE        string
	SUSPENDED   string
}{
	IN_PROGRESS: "in_progress",
	DONE:        "done",
	SUSPENDED:   "suspended",
}

var StatusValues = []string{Status.IN_PROGRESS, Status.DONE, Status.SUSPENDED}

// Struct holding valid project task priority values
var Priority = struct {
	LOW    string
	MEDIUM string
	HIGH   string
}{
	LOW:    "low",
	MEDIUM: "medium",
	HIGH:   "high",
}

var PriorityValues = []string{Priority.LOW, Priority.MEDIUM, Priority.HIGH}

const (
	TIME_FORMAT = "2006-01-02"
)
