package ex31

import (
	"errors"
)

type MergeDictsJob struct {
	Dicts      []map[string]string
	Merged     map[string]string
	IsFinished bool
}

var (
	ErrNotEnoughDicts = errors.New("at least 2 dictionaries are required")
	ErrNilDict        = errors.New("nil dictionary")
)

func ExecuteMergeDictsJob(job *MergeDictsJob) (*MergeDictsJob, error) {
	defer job.makeFinished()

	if len(job.Dicts) < 2 {
		return job, ErrNotEnoughDicts
	}

	job.Merged = make(map[string]string)
	for _, dict := range job.Dicts {
		if dict == nil {
			return job, ErrNilDict
		}

		for key, value := range dict {
			job.Merged[key] = value
		}
	}

	return job, nil
}

func (job *MergeDictsJob) makeFinished() {
	job.IsFinished = true
}
