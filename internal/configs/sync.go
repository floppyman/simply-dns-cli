package configs

import (
	"errors"
	"os"
	"os/user"
	"path"
	"path/filepath"

	"github.com/umbrella-sh/um-common/configuration"
	"github.com/umbrella-sh/um-common/logging/ulog"

	"github.com/umbrella-sh/simply-dns-sync/internal/api"
)

const syncFileName = "sync-state.json"

type SyncState struct {
	Products []*api.SimplyProduct `json:"products"`
}

var Sync *SyncState

func InitSync() error {
	ulog.Console.Info().Msg("Loading sync-state...")

	var usr *user.User
	var err error
	usr, err = user.Current()
	if err != nil {
		return err
	}

	var homeConfigFolder = path.Join(usr.HomeDir, ".config", AppName)
	var homeConfigFilePath = path.Join(usr.HomeDir, ".config", AppName, syncFileName)

	configFolderExists := false
	if p, _ := filepath.Glob(homeConfigFolder); len(p) > 0 {
		configFolderExists = true
	}

	cfgDef := initDefaultSync()
	Sync, err = configuration.LoadJson(homeConfigFilePath, &cfgDef, false, "")
	if err != nil {
		if configFolderExists {
			if errors.Is(err, os.ErrNotExist) {
				err = configuration.SaveJson(homeConfigFilePath, cfgDef, true)
				if err != nil {
					ulog.Console.Err(err).Msgf("Failed to save sync-state file to '%s'", homeConfigFilePath)
					return err
				}
				return nil
			}

			ulog.Console.Err(err).Msgf("Failed to load sync-state file from '%s'", homeConfigFilePath)
			return err
		}

		var localConfigFilePath = path.Join("./", syncFileName)
		Sync, err = configuration.LoadJson(localConfigFilePath, &cfgDef, false, "")
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				err = configuration.SaveJson(localConfigFilePath, cfgDef, true)
				if err != nil {
					ulog.Console.Err(err).Msgf("Failed to save sync-state file to '%s'", localConfigFilePath)
					return err
				}
				return nil
			}

			ulog.Console.Err(err).Msgf("Failed to load sync-state file from '%s'", localConfigFilePath)
			return err
		}
	}

	ulog.Console.Info().Msg("Sync-state loaded")
	return nil
}

func SaveSyncState() error {
	ulog.Console.Info().Msg("Saving sync-state...")

	var usr *user.User
	var err error
	usr, err = user.Current()
	if err != nil {
		return err
	}
	var configFolder = path.Join(usr.HomeDir, ".config", AppName)
	var homeFilePath = path.Join(usr.HomeDir, ".config", AppName, syncFileName)
	var localFilePath = path.Join("./", syncFileName)

	configFolderExists := false
	if p, _ := filepath.Glob(configFolder); len(p) > 0 {
		configFolderExists = true
	}

	if configFolderExists {
		err = configuration.SaveJsonIndented(homeFilePath, Sync, true)
		if err != nil {
			return err
		}
	}

	if !configFolderExists {
		err = configuration.SaveJsonIndented(localFilePath, Sync, true)
		if err != nil {
			return err
		}
	}

	ulog.Console.Info().Msg("Sync-state saved")
	return nil
}

func initDefaultSync() SyncState {
	return SyncState{
		Products: make([]*api.SimplyProduct, 0),
	}
}
