package configs

import (
	"errors"
	"os"
	"os/user"
	"path"
	"path/filepath"

	"github.com/umbrella-sh/um-common/configuration"

	"github.com/umbrella-sh/simply-dns-cli/internal/styles"
)

type Config struct {
	SimplyApi ConfigSimplyApi `json:"simply_api"`
}

type ConfigSimplyApi struct {
	Url           string `json:"url"`
	AccountNumber string `json:"account_number"`
	AccountApiKey string `json:"account_api_key"`
}

var Main *Config

func InitConfig() error {
	styles.WaitPrint("Loading config")

	var usr *user.User
	var err error
	usr, err = user.Current()
	if err != nil {
		return err
	}

	var homeConfigFolder = path.Join(usr.HomeDir, ".config", AppName)
	var homeConfigPath = path.Join(usr.HomeDir, ".config", AppName, configFileName)

	configFolderExists := false
	if p, _ := filepath.Glob(homeConfigFolder); len(p) > 0 {
		configFolderExists = true
	}

	cfgDef := initDefaultConfig()
	Main, err = configuration.LoadJson(homeConfigPath, &cfgDef, false, "")
	if err != nil {
		if configFolderExists {
			styles.FailPrint("Failed to load config file from '%s'", homeConfigPath)
			styles.FailPrint("Error: %v", err)
			return err
		}

		if errors.Is(err, os.ErrNotExist) {
			var localConfigPath = path.Join("./", configFileName)
			Main, err = configuration.LoadJson(localConfigPath, &cfgDef, false, "")
			if err != nil {
				styles.FailPrint("Failed to load config file from '%s'", localConfigPath)
				styles.FailPrint("Error: %v", err)
				return err
			}
		}
	}

	return testConfig()
}

func initDefaultConfig() Config {
	return Config{
		SimplyApi: ConfigSimplyApi{
			Url:           "",
			AccountNumber: "",
			AccountApiKey: "",
		},
	}
}

func testConfig() error {
	hasErr := false
	styles.BlankPrint("Testing config")

	if Main.SimplyApi.Url == "" {
		hasErr = true
		styles.FailPrint("'simply_api.url' must be set to a valid url for Simply.com API (https://www.simply.com/dk/docs/api/)")
	}

	if Main.SimplyApi.AccountNumber == "" {
		hasErr = true
		styles.FailPrint("'simply_api.account_number' must be set to account number retrieved from Simply.com")
	}

	if Main.SimplyApi.AccountApiKey == "" {
		hasErr = true
		styles.FailPrint("'simply_api.account_api_key' must be set to the account Api Key retrieved from Simply.com")
	}

	if hasErr {
		styles.FailPrint("Config loaded but testing failed")
		return errors.New("")
	}
	styles.SuccessPrint("Config loaded and testing success")
	return nil
}
