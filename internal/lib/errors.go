package lib

import "fmt"

func DeletedError(groupName string, id int) error {
	return fmt.Errorf(
		"cannot find sprite (group: %s id: %d).\n It may already be deleted.",
		groupName, id,
	)
}
