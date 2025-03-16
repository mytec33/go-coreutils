//go:build darwin || linux || openbsd

package platform

import (
	"os"
	"os/user"
	"strconv"
)

func GetUserId() (string, error) {
	uid := os.Geteuid()
	user, err := user.LookupId(strconv.Itoa(uid))
	if err != nil {
		return "", err
	}

	return user.Username, nil
}
