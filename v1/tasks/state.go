package tasks

import "time"

const (
	// StatePending - initial state of a task
	StatePending = "PENDING"
	// StateReceived - when task is received by a worker
	StateReceived = "RECEIVED"
	// StateStarted - when the worker starts processing the task
	StateStarted = "STARTED"
	// StateRetry - when failed task has been scheduled for retry
	StateRetry = "RETRY"
	// StateSuccess - when the task is processed successfully
	StateSuccess = "SUCCESS"
	// StateFailure - when processing of the task fails
	StateFailure = "FAILURE"
	// StateFailure - when processing of the task fails
	StateCancelled = "CANCELLED"
)

// TaskState represents a state of a task
type TaskState struct {
	TaskUUID  string        `bson:"_id"`
	TaskName  string        `bson:"task_name"`
	State     string        `bson:"state"`
	Results   []*TaskResult `bson:"results"`
	Error     string        `bson:"error"`
	CreatedAt time.Time     `bson:"created_at"`
}

// GroupMeta stores useful metadata about tasks within the same group
// E.g. UUIDs of all tasks which are used in order to check if all tasks
// completed successfully or not and thus whether to trigger chord callback
type GroupMeta struct {
	GroupUUID      string    `bson:"_id"`
	TaskUUIDs      []string  `bson:"task_uuids"`
	ChordTriggered bool      `bson:"chord_triggered"`
	Lock           bool      `bson:"lock"`
	CreatedAt      time.Time `bson:"created_at"`
	UpdatedAt      time.Time `bson:"updated_at"`
}

// NewCancelledTaskState ...
func NewCancelledTaskState(signature *Signature) *TaskState {
	return &TaskState{
		TaskUUID:  signature.UUID,
		TaskName:  signature.Name,
		State:     StateCancelled,
		CreatedAt: time.Now().UTC(),
	}
}

// NewPendingTaskState ...
func NewPendingTaskState(signature *Signature) *TaskState {
	return &TaskState{
		TaskUUID:  signature.UUID,
		TaskName:  signature.Name,
		State:     StatePending,
		CreatedAt: time.Now().UTC(),
	}
}

// NewReceivedTaskState ...
func NewReceivedTaskState(signature *Signature) *TaskState {
	return &TaskState{
		TaskUUID:  signature.UUID,
		State:     StateReceived,
		CreatedAt: time.Now().UTC(),
	}
}

// NewStartedTaskState ...
func NewStartedTaskState(signature *Signature) *TaskState {
	return &TaskState{
		TaskUUID:  signature.UUID,
		State:     StateStarted,
		CreatedAt: time.Now().UTC(),
	}
}

// NewSuccessTaskState ...
func NewSuccessTaskState(signature *Signature, results []*TaskResult) *TaskState {
	return &TaskState{
		TaskUUID:  signature.UUID,
		State:     StateSuccess,
		Results:   results,
		CreatedAt: time.Now().UTC(),
	}
}

// NewFailureTaskState ...
func NewFailureTaskState(signature *Signature, err string) *TaskState {
	return &TaskState{
		TaskUUID:  signature.UUID,
		State:     StateFailure,
		Error:     err,
		CreatedAt: time.Now().UTC(),
	}
}

// NewRetryTaskState ...
func NewRetryTaskState(signature *Signature) *TaskState {
	return &TaskState{
		TaskUUID:  signature.UUID,
		State:     StateRetry,
		CreatedAt: time.Now().UTC(),
	}
}

// IsCompleted returns true if state is SUCCESS or FAILURE,
// i.e. the task has finished processing and either succeeded or failed.
func (taskState *TaskState) IsCompleted() bool {
	return taskState.IsSuccess() || taskState.IsFailure()
}

// IsSuccess returns true if state is SUCCESS
func (taskState *TaskState) IsSuccess() bool {
	return taskState.State == StateSuccess
}

// IsFailure returns true if state is FAILURE
func (taskState *TaskState) IsFailure() bool {
	return taskState.State == StateFailure
}

// IsCancelled returns true if state is CANCELLED
func (taskState *TaskState) IsCancelled() bool {
	return taskState.State == StateCancelled
}

// IsPending returns true if state is PENDING
func (taskState *TaskState) IsPending() bool {
	return taskState.State == StatePending
}
