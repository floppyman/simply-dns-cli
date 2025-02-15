package configs

import (
	"fmt"
	"os/user"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/umbrella-sh/um-common/configuration"
	"github.com/umbrella-sh/um-common/logging/ulog"
	"github.com/umbrella-sh/um-common/types"

	"github.com/umbrella-sh/simply-dns-cli/internal/api"
)

type BackupFile struct {
	Items []*api.SimplyProduct `json:"items"`
}

func diskEntityExists(entityPath string) bool {
	if p, _ := filepath.Glob(entityPath); len(p) > 0 {
		return true
	}
	return false
}

func LoadBackup(backupName string) (*BackupFile, error) {
	ulog.Console.Info().Msg("Loading backup...")

	fileName := fmt.Sprintf("%s.json", backupName)

	var usr *user.User
	var err error
	usr, err = user.Current()
	if err != nil {
		return nil, err
	}
	var configFolder = path.Join(usr.HomeDir, ".config", AppName)
	var homeFilePath = path.Join(usr.HomeDir, ".config", AppName, fileName)
	var localFilePath = path.Join("./", fileName)

	configFolderExists := diskEntityExists(configFolder)
	homeFileExists := diskEntityExists(homeFilePath)

	cfgDef := BackupFile{}
	var backupFile *BackupFile

	if configFolderExists && homeFileExists {
		backupFile, err = configuration.LoadJson(homeFilePath, &cfgDef, false, "")
		if err != nil {
			return nil, err
		}
		return backupFile, nil
	}

	if !configFolderExists {
		backupFile, err = configuration.LoadJson(localFilePath, &cfgDef, false, "")
		if err != nil {
			return nil, err
		}
	}

	ulog.Console.Info().Msg("Backup loaded")
	return backupFile, nil
}

func SaveBackup(data []*api.SimplyProduct, now time.Time) error {
	ulog.Console.Info().Msg("Saving backup...")

	fileName := strings.Replace(backupFileName, "{{ts}}", now.Format(types.TimeFormatIsoFullDateTimeCompact), 1)
	backupFile := BackupFile{Items: data}

	var usr *user.User
	var err error
	usr, err = user.Current()
	if err != nil {
		return err
	}
	var configFolder = path.Join(usr.HomeDir, ".config", AppName)
	var homeFilePath = path.Join(usr.HomeDir, ".config", AppName, fileName)
	var localFilePath = path.Join("./", fileName)

	configFolderExists := false
	if p, _ := filepath.Glob(configFolder); len(p) > 0 {
		configFolderExists = true
	}

	if configFolderExists {
		err = configuration.SaveJsonIndented(homeFilePath, backupFile, true)
		if err != nil {
			return err
		}
	}

	if !configFolderExists {
		err = configuration.SaveJsonIndented(localFilePath, backupFile, true)
		if err != nil {
			return err
		}
	}

	ulog.Console.Info().Msg("Backup saved")
	return nil
}
