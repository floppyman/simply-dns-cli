package restore

import (
	"encoding/json"
	"os"

	"github.com/floppyman/simply-dns-cli/internal/configs"
	"github.com/floppyman/simply-dns-cli/internal/mocks"
	"github.com/floppyman/simply-dns-cli/internal/objects"
	"github.com/floppyman/simply-dns-cli/internal/styles"
)

func LoadBackup(backupFilePath string) *objects.RestoreFile {
	if configs.IsMocking {
		return mocks.LoadBackup()
	}

	bytes, err := os.ReadFile(backupFilePath)
	if err != nil {
		styles.FailPrint("Failed to read backup file from provided path")
		styles.FailPrint("Error: %v", err)
		return nil
	}

	var restoreFile *objects.RestoreFile
	err = json.Unmarshal(bytes, &restoreFile)
	if err != nil {
		styles.FailPrint("Failed to parse backup file, invalid json")
		styles.FailPrint("Error: %v", err)
		return nil
	}

	return restoreFile
}
