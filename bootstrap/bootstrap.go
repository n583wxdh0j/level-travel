package bootstrap

import "level-travel/config"

func BootstrapConfig() {
	// TODO сделай инитициалиацию оконфига из kubov
	config.Config.Database.User = "n583wxdh0j"
	config.Config.Database.Name = "n583wxdh0j"
	config.Config.Database.SSLMode = "disable"

	config.Config.GithubAPIKey = ""
}
