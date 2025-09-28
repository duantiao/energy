//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package remotecfg

import (
	"encoding/json"
	assetRemoteCfg "github.com/energye/energy/v2/cmd/assets/remotecfg"
)

type VersionsUpgradeList map[string]TVersionsUpgrade

var upgradeList VersionsUpgradeList

type TVersionsUpgrade struct {
	Enable           int               `json:"enable"`
	Identical        string            `json:"identical"`
	DependenceModule TDependenceModule `json:"dependenceModule"`
}

type TDependenceModule struct {
	LCL map[string]string `json:"lcl"`
	CEF map[string]string `json:"cef"`
}

// 版本发布升级列表
func VersionUpgradeList() (VersionsUpgradeList, error) {
	if upgradeList == nil {
		//data, err := tools.Get(env.GlobalDevEnvConfig.RemoteURL(consts.VERSIONS_UPGRADE_URL), env.GlobalDevEnvConfig.Proxy)
		//if err != nil {
		//	return nil, err
		//}
		data := assetRemoteCfg.VersionUpgradeCfg
		var vu map[string]TVersionsUpgrade
		if err := json.Unmarshal(data, &vu); err != nil {
			return nil, err
		}
		upgradeList = vu
	}
	return upgradeList, nil
}

// 参数 ver: x.x.x
func (m VersionsUpgradeList) Item(ver string) *TVersionsUpgrade {
	// 找到相同版配置
	installVersion, ok := m[ver]
	if !ok {
		return nil
	}
	for {
		if installVersion.Identical != "" {
			if installVersion, ok = m[installVersion.Identical]; !ok {
				return nil
			}
		} else {
			break
		}
	}
	return &installVersion
}
