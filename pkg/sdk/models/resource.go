package models

type StreamSender func(Resource) error

type Resource struct {
	// ARN uniquely identifies an AWS resource across regions, accounts and types.
	ARN string
	// ID doesn't uniquely identifies a resource. It will be used to create a
	// unique identifier by concating PARTITION|REGION|ACCOUNT|TYPE|ID
	ID          string
	Description interface{}

	Name      string
	Account   string
	Region    string
	Partition string
	Type      string
}

func (r Resource) UniqueID() string {
	return r.ID
}
