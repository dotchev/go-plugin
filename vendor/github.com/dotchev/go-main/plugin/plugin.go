package plugin

import (
	uuid "github.com/satori/go.uuid"
)

type PluginRequest struct {
	Name string
	Size int
}

func UUID() string {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return id.String()
}
