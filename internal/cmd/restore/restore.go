package restore

import (
	"time"

	"github.com/umbrella-sh/simply-dns-cli/internal/api"
	"github.com/umbrella-sh/simply-dns-cli/internal/styles"
)

//goland:noinspection GoNameStartsWithPackageName
type RestoreFile struct {
	TimeStamp time.Time            `json:"time_stamp"`
	Items     []*api.SimplyProduct `json:"items"`
}

func LoadBackup(backupName string) (*RestoreFile, error) {
	styles.WaitPrint("Loading backup")

	// fileName := fmt.Sprintf("%s.json", backupName)
	//
	// var usr *user.User
	// var err error
	// usr, err = user.Current()
	// if err != nil {
	// 	return nil, err
	// }
	// var homeConfigFolder = path.Join(usr.HomeDir, ".config", configs.AppName)
	// var homeConfigFolderBackups = path.Join(homeConfigFolder, "backups")
	// var homeFilePath = path.Join(homeConfigFolderBackups, fileName)
	//
	// var localFolderBackups = "./backups"
	// var localFilePath = path.Join(localFolderBackups, fileName)
	//
	// homeConfigFolderExists := ext.FolderExist(homeConfigFolder)
	// homeConfigFolderBackupsExists := ext.FolderExist(homeConfigFolderBackups)
	// homeFileExists := ext.FileExist(homeFilePath)
	//
	// localFolderBackupsExists := ext.FolderExist(localFolderBackups)
	//
	// cfgDef := RestoreFile{}
	var backupFile *RestoreFile

	// if homeConfigFolderExists && homeConfigFolderBackupsExists && homeFileExists {
	// 	backupFile, err = configuration.LoadJson(homeFilePath, &cfgDef, false, "")
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return backupFile, nil
	// }
	//
	// if !homeConfigFolderExists {
	// 	backupFile, err = configuration.LoadJson(localFilePath, &cfgDef, false, "")
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	styles.SuccessPrint("Backup loaded")
	return backupFile, nil
}
