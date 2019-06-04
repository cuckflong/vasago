package configure

import (
	"errors"
	"net/url"
	"strconv"
)

type ConfigParam struct {
	Key         string
	Description string
	Value       string
	Validate    func(value string) error
}

type ConfigDescription struct {
	Parameters []*ConfigParam
}

var TestConfig = &ConfigDescription{
	Parameters: []*ConfigParam{
		{
			Key:         "age",
			Description: "Your age here",
			Value:       "",
			Validate: func(value string) error {
				_, err := strconv.Atoi(value)
				if err != nil {
					return errors.New("you must enter a valid number")
				}

				return nil
			},
		},
		{
			Key:         "email",
			Description: "Your email here",
			Value:       "default@example.org",
			Validate: func(value string) error {
				u, err := url.Parse(value)
				if err != nil {
					return errors.New("your email is invalid")
				}

				if u.User.Username() == "" || u.Host == "" {
					return errors.New("your email is invalid")
				}

				return nil
			},
		},
		{
			Key:         "name",
			Description: "Your name here",
			Value:       "",
			Validate: func(value string) error {
				if value == "" {
					return errors.New("your name can't be empty")
				}

				return nil
			},
		},
	},
}
