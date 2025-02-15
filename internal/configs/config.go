package configs

import (
	"errors"
	"os"
	"os/user"
	"path"
	"path/filepath"

	"github.com/umbrella-sh/um-common/configuration"
	log "github.com/umbrella-sh/um-common/logging/basic"
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
	log.WaitPrint("loading config")

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
			log.Errorf("failed to load config file from '%s'\n", homeConfigPath)
			log.Errorln(err)
			return err
		}

		if errors.Is(err, os.ErrNotExist) {
			var localConfigPath = path.Join("./", configFileName)
			Main, err = configuration.LoadJson(localConfigPath, &cfgDef, false, "")
			if err != nil {
				log.Errorf("failed to load config file from '%s'", localConfigPath)
				log.Errorln(err)
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
	log.BlankPrint("testing config")

	if Main.SimplyApi.Url == "" {
		hasErr = true
		log.Errorln("'simply_api.url' must be set to a valid url for Simply.com API (https://www.simply.com/dk/docs/api/)")
	}

	if Main.SimplyApi.AccountNumber == "" {
		hasErr = true
		log.Errorln("'simply_api.account_number' must be set to account number retrieved from Simply.com")
	}

	if Main.SimplyApi.AccountApiKey == "" {
		hasErr = true
		log.Errorln("'simply_api.account_api_key' must be set to the account Api Key retrieved from Simply.com")
	}

	if hasErr {
		log.FailPrint("config loaded but testing failed")
		return errors.New("")
	}
	log.SuccessPrint("config loaded and testing success")
	return nil
}
