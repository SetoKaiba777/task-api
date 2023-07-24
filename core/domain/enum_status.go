package domain

type Status string



const (
	CREATED   Status = "Started"
	DOING Status = "Doing"
	FINISHED Status = "Finished"
)

func (s Status) fromString() error {
	switch s{
	case CREATED,DOING,FINISHED:
			return nil
	default:
			return ErrInvalidStatus
	}
}