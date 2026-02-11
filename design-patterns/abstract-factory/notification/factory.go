package notification

import (
	"abstract-factory/notification/aws"
	"abstract-factory/notification/contract"
	"abstract-factory/notification/gcp"
	"errors"
)

type Provider string

const (
	AWS Provider = "aws"
	GCP Provider = "gcp"
)

func NewFactory(p Provider) (contract.NotificationFactory, error) {
	switch p {
	case AWS:
		return &aws.Factory{}, nil
	case GCP:
		return &gcp.Factory{}, nil
	default:
		return nil, errors.New("unsupported provider")
	}
}
