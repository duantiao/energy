package remotecfg

import _ "embed"

var (
	//go:embed versions-upgrade.json
	VersionUpgradeCfg []byte
	//go:embed model-cef.json
	ModelCEFCfg []byte
	//go:embed model-base-config.json
	ModelBaseCfg []byte
	//go:embed latest-version.json
	LatestVersionCfg []byte
)
