package gcp

import "abstract-factory/notification/contract"

type Factory struct{}

func(f *Factory) CreateEmailSender() contract.EmailSender{
	return &EmailSender{}
}

func(f *Factory) CreateSMSSender() contract.SMSSender{
	return &SMSSender{}
} 