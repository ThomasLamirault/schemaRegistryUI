package errors

import "fmt"

type AppConfigError struct {
	Address string
	Port    string
}

func (e *AppConfigError) Error() string {
	return fmt.Sprintf("appconfig uncomplete, address : %s, port : %s", e.Address, e.Port)
}
