package restore

import (
	"encoding/json"
	"os"
	"time"

	apio "github.com/umbrella-sh/simply-dns-cli/internal/api_objects"
	"github.com/umbrella-sh/simply-dns-cli/internal/styles"
)

//goland:noinspection GoNameStartsWithPackageName
type RestoreFile struct {
	TimeStamp time.Time                      `json:"time_stamp"`
	Items     map[string]*apio.SimplyProduct `json:"items"`
}

func LoadBackup(backupFilePath string) *RestoreFile {
	bytes, err := os.ReadFile(backupFilePath)
	if err != nil {
		styles.FailPrint("Failed to read backup file from provided path")
		styles.FailPrint("Error: %v", err)
		return nil
	}

	var backupFile *RestoreFile
	err = json.Unmarshal(bytes, &backupFile)
	if err != nil {
		styles.FailPrint("Failed to parse backup file, invalid json")
		styles.FailPrint("Error: %v", err)
		return nil
	}

	return backupFile
}
