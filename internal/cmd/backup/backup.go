package backup

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"strings"
	"time"

	"github.com/umbrella-sh/um-common/configuration"
	"github.com/umbrella-sh/um-common/ext"
	log "github.com/umbrella-sh/um-common/logging/basic"
	"github.com/umbrella-sh/um-common/types"

	"github.com/umbrella-sh/simply-dns-cli/internal/api"
	"github.com/umbrella-sh/simply-dns-cli/internal/configs"
)

const backupFileName = "backup_{{ts}}.json"

//goland:noinspection GoNameStartsWithPackageName
type BackupFile struct {
	TimeStamp time.Time            `json:"time_stamp"`
	Items     []*api.SimplyProduct `json:"items"`
}

func SaveBackup(data []*api.SimplyProduct, now time.Time) (string, error) {
	fileName := strings.Replace(backupFileName, "{{ts}}", now.Format(types.TimeFormatIsoFullDateTimeCompact), 1)
	backupFile := BackupFile{TimeStamp: now, Items: data}

	usr, usrErr := user.Current()
	if usrErr != nil {
		return "", usrErr
	}

	homeErr := saveInHomeFolder(usr, fileName, backupFile)
	if homeErr == nil {
		return fileName, nil
	}

	localErr := saveInLocalFolder(fileName, backupFile)
	if localErr == nil {
		return fileName, nil
	}

	log.FailPrint(fmt.Sprint(homeErr))
	log.FailPrint(fmt.Sprint(localErr))
	return "", fmt.Errorf("")
}

func saveInHomeFolder(usr *user.User, fileName string, backupFile BackupFile) error {
	homeConfigFolder := path.Join(usr.HomeDir, ".config", configs.AppName)
	if !ext.FolderExist(homeConfigFolder) {
		return fmt.Errorf("'%s' folder does not exist", homeConfigFolder)
	}

	var err error
	homeConfigFolderBackups := path.Join(homeConfigFolder, "backups")
	if !ext.FolderExist(homeConfigFolderBackups) {
		err = os.MkdirAll(homeConfigFolderBackups, 0750)
		if err != nil {
			return err
		}
	}

	homeFilePath := path.Join(homeConfigFolderBackups, fileName)
	err = configuration.SaveJsonIndented(homeFilePath, backupFile, true)
	if err != nil {
		return err
	}

	return nil
}

func saveInLocalFolder(fileName string, backupFile BackupFile) error {
	var err error

	localFolderBackups := "./backups"
	if !ext.FolderExist(localFolderBackups) {
		err = os.MkdirAll(localFolderBackups, 0750)
		if err != nil {
			return err
		}
	}

	localFilePath := path.Join(localFolderBackups, fileName)
	err = configuration.SaveJsonIndented(localFilePath, backupFile, true)
	if err != nil {
		return err
	}

	return nil
}
