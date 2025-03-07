package lib

import "fmt"

func DeletedError(groupName string, id int) error {
	return fmt.Errorf(
		"cannot find sprite (group: %s id: %d).\n It may already be deleted.",
		groupName, id,
	)
}

func UnmarshalError(path string, err error) error {
	return fmt.Errorf("\"%s\", %w", path, err)
}
