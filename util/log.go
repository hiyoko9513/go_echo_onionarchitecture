package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetLogFilePath(logPath string) (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	exePathFolder := filepath.Dir(exePath)

	return fmt.Sprintf("%s%s", exePathFolder, logPath), nil
}
