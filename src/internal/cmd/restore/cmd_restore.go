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

	"github.com/umbrella-sh/simply-dns-cli/internal/configs"
	"github.com/umbrella-sh/simply-dns-cli/internal/forms"
	"github.com/umbrella-sh/simply-dns-cli/internal/objects"
	"github.com/umbrella-sh/simply-dns-cli/internal/shared"
	"github.com/umbrella-sh/simply-dns-cli/internal/styles"
)

func cmdRun(_ *cobra.Command, _ []string) {
	styles.Println(styles.Info("Generating create or delete commands to restore DNS records for Domain"))
	styles.Blank()

	var file *objects.RestoreFile
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
	toCreate, toDelete := findChanges(fileDnsRecords, records)

	styles.Blank()
	generateCommands(domain, toCreate, toDelete)
}

func generateCommands(domain string, toCreate map[string]*objects.SimplyDnsRecord, toDelete map[string]*objects.SimplyDnsRecord) {
	styles.WaitPrint("Generating commands to execute to restore the domain from the backup selected")

	if len(toCreate) == 0 && len(toDelete) == 0 {
		styles.InfoPrint("No commands needed to restore the domain")
		return
	}

	styles.Blank()

	longestDomain := 0
	for _, v := range toCreate {
		domLen := len(fmt.Sprintf("%s.%s", v.Name, domain))
		if domLen > longestDomain {
			longestDomain = domLen
		}
	}
	for _, v := range toDelete {
		domLen := len(fmt.Sprintf("%s.%s", v.Name, domain))
		if domLen > longestDomain {
			longestDomain = domLen
		}
	}

	for _, v := range toDelete {
		styles.Printf("%s%s\n",
			styles.Graphic("%-*s | ", longestDomain, fmt.Sprintf("%s.%s", v.Name, domain)),
			styles.Error("%s remove -d %s -r %d",
				configs.AppName,
				domain,
				v.RecordId,
			),
		)
	}

	for _, v := range toCreate {
		if v.Type == objects.DnsRecTypeMX {
			styles.Printf("%s%s\n",
				styles.Graphic("%-*s | ", longestDomain, fmt.Sprintf("%s.%s", v.Name, domain)),
				styles.Success("%s create -d %s -t %s -l %d -n %s -v %s -p %d -c %s",
					configs.AppName,
					domain,
					v.Type,
					v.TTL,
					v.Name,
					v.Data,
					v.Priority,
					v.Comment,
				),
			)
			continue
		}

		styles.Printf("%s%s\n",
			styles.Graphic("%-*s | ", longestDomain, fmt.Sprintf("%s.%s", v.Name, domain)),
			styles.Success(`%s create -d %s -t %s -l %d -n %s -v %s -c "%s"`,
				configs.AppName,
				domain,
				v.Type,
				v.TTL,
				v.Name,
				v.Data,
				v.Comment,
			),
		)
	}
	styles.Blank()
	styles.SuccessPrint("Commands generated")
}

func findChanges(localDnsRecords []*objects.SimplyDnsRecord, remoteDnsRecords []*objects.SimplyDnsRecord) (toCreate map[string]*objects.SimplyDnsRecord, toDelete map[string]*objects.SimplyDnsRecord) {
	toCreate = make(map[string]*objects.SimplyDnsRecord)
	toDelete = make(map[string]*objects.SimplyDnsRecord)

	// detect new elements
	for _, l := range localDnsRecords {
		found := false
		for _, r := range remoteDnsRecords {
			if l.GetHash() == r.GetHash() {
				found = true
			}
		}
		if !found {
			toCreate[l.GetHash()] = l
		}
	}

	// detect delete elements
	for _, r := range remoteDnsRecords {
		found := true
		for _, l := range localDnsRecords {
			if l.GetHash() == r.GetHash() {
				found = false
			}
		}
		if found {
			toDelete[r.GetHash()] = r
		}
	}

	return toCreate, toDelete
}

func collectDomainFromBackup(file *objects.RestoreFile) (cancelled bool, domain string) {
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
