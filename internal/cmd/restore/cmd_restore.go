package restore

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"slices"

	"github.com/spf13/cobra"
	"github.com/umbrella-sh/um-common/ext"

	"github.com/umbrella-sh/simply-dns-cli/internal/api"
	"github.com/umbrella-sh/simply-dns-cli/internal/configs"
	"github.com/umbrella-sh/simply-dns-cli/internal/forms"
	"github.com/umbrella-sh/simply-dns-cli/internal/shared"
	"github.com/umbrella-sh/simply-dns-cli/internal/styles"
)

func cmdRun(_ *cobra.Command, _ []string) {
	styles.Println(styles.Info("Restore DNS records from a backup"))
	styles.Blank()

	var file *RestoreFile
	if options.BackupFilePath != "" {
		styles.WaitPrint("Loading backup from provided backup file")
		file = LoadBackup(options.BackupFilePath)
		if file == nil {
			return
		}
		styles.SuccessPrint("Backup loaded")
	}

	var cancelled bool
	if file == nil {
		styles.InfoPrint("No backup file provided, listing backups in default locations")
		styles.Blank()

		var backupFilePath string
		var noFiles bool
		cancelled, backupFilePath, noFiles = collectBackupName()
		if cancelled {
			printCancelText()
			return
		}
		if noFiles {
			printNoFilesText()
			return
		}

		styles.Blank()
		styles.SuccessPrint("Backup selected")

		styles.WaitPrint("Loading backup from selection")
		file = LoadBackup(backupFilePath)
		if file == nil {
			return
		}
		styles.SuccessPrint("Backup loaded")
		styles.Blank()
	}

	var domain string
	cancelled, domain = collectDomainFromBackup(file)
	if cancelled {
		printCancelText()
		return
	}

	styles.Blank()
	records := shared.PullDnsRecords(domain, "")
	if records == nil {
		styles.FailPrint("Failed to get DNS entries from Domain")
		return
	}

	fileDnsRecords := file.Items[domain].DnsRecords
	generateCommands(fileDnsRecords, records)
}

func generateCommands(localDnsRecords []*api.SimplyDnsRecord, remoteDnsRecords []*api.SimplyDnsRecord) (toCreate map[string]*api.SimplyDnsRecord, toUpdate map[string]*api.SimplyDnsRecord, toDelete map[string]*api.SimplyDnsRecord) {
	noChanges := make(map[string]*api.SimplyDnsRecord)
	hasChanges := make(map[string]*api.SimplyDnsRecord)
	toCreate = make(map[string]*api.SimplyDnsRecord)
	toUpdate = make(map[string]*api.SimplyDnsRecord)
	toDelete = make(map[string]*api.SimplyDnsRecord)

	for _, l := range localDnsRecords {
		for _, r := range remoteDnsRecords {
			if l.GetHash() == r.GetHash() {
				found = true
			}
		}
	}

	// detect new elements
	for _, r := range remoteDnsRecords {
		found := false
		for _, l := range localDnsRecords {
			if l.GetHash() == r.GetHash() {
				found = true
			}
		}
		if !found {
			toDelete[r.GetHash()] = r
		}
	}

	return toCreate, toUpdate, toDelete
}

func collectDomainFromBackup(file *RestoreFile) (cancelled bool, domain string) {
	var objNames = make([]string, 0)
	for k := range file.Items {
		objNames = append(objNames, k)
	}
	slices.Sort(objNames)

	cancelled, domain = forms.RunDomainSelect(objNames)
	if cancelled {
		return cancelled, ""
	}

	return
}

func collectBackupName() (cancelled bool, backupFilePath string, noFiles bool) {
	usr, usrErr := user.Current()

	if usrErr != nil {
		return true, "", false
	}

	var folder string
	var fileNames []string
	var err error
	var fullFileNames []any

	folder, fileNames, err = listFromHomeFolder(usr)
	if err == nil {
		if len(fileNames) == 0 {
			return false, "", true
		}

		fullFileNames = createFullFileNames(folder, fileNames)
		cancelled, backupFilePath = forms.RunBackupNameSelect(fileNames, fullFileNames)
		return cancelled, backupFilePath, false
	}

	folder, fileNames, err = listFromLocalFolder()
	if err == nil {
		if len(fileNames) == 0 {
			return false, "", true
		}

		fullFileNames = createFullFileNames(folder, fileNames)
		cancelled, backupFilePath = forms.RunBackupNameSelect(fileNames, fullFileNames)
		return cancelled, backupFilePath, false
	}

	return false, "", true
}

func listFromHomeFolder(usr *user.User) (string, []string, error) {
	backupFolder := path.Join(usr.HomeDir, ".config", configs.AppName, "backups")
	if !ext.FolderExist(backupFolder) {
		return "", nil, fmt.Errorf("'%s' folder does not exist", backupFolder)
	}

	return getFolderAndFiles(backupFolder)
}

func listFromLocalFolder() (string, []string, error) {
	backupFolder := "./backups"
	if !ext.FolderExist(backupFolder) {
		return "", nil, fmt.Errorf("'%s' folder does not exist", backupFolder)
	}

	return getFolderAndFiles(backupFolder)
}

func getFolderAndFiles(backupFolder string) (folder string, fileNames []string, err error) {
	folder = backupFolder
	var files []os.DirEntry
	files, err = os.ReadDir(folder)
	if err != nil {
		return "", nil, err
	}

	fileNames = make([]string, 0)
	for _, v := range files {
		if v.IsDir() {
			continue
		}
		fileNames = append(fileNames, v.Name())
	}

	return
}

func createFullFileNames(folder string, fileNames []string) []any {
	var fullFileNames = make([]any, 0)
	for _, v := range fileNames {
		fullFileNames = append(fullFileNames, filepath.Join(folder, v))
	}
	return fullFileNames
}

func printCancelText()  { styles.Println(styles.Warn("\nRestore was cancelled\n")) }
func printNoFilesText() { styles.Println(styles.Warn("\nNo backup files found to restore\n")) }
