//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"fmt"
	"github.com/energye/energy/v2/consts"
)

// process
const (
	internalProcess          = "process"
	infraInternalProcess     = "infraProcess"
	internalProcessBrowserId = "browserId"
	internalProcessFrameId   = "frameId"
)

var _processObject *ICefV8Value

// makeProcess 进程扩展变量
func makeProcess(browser *ICefBrowser, frame *ICefFrame, context *ICefV8Context, enableInfraProcess bool) {
	if _processObject != nil {
		fmt.Println("[Debug Process] _processObject string value will free:", _processObject.GetStringValue())
		// 刷新时释放掉
		_processObject.Free()
	}
	// process
	_processObject = V8ValueRef.NewObject(nil, nil)
	_processObject.setValueByKey(internalProcessBrowserId, V8ValueRef.NewInt(browser.Identifier()), consts.V8_PROPERTY_ATTRIBUTE_READONLY)
	_processObject.setValueByKey(internalProcessFrameId, V8ValueRef.NewString(frame.Identifier()), consts.V8_PROPERTY_ATTRIBUTE_READONLY)

	// process key to v8 global
	processKey := internalProcess
	if enableInfraProcess {
		processKey = infraInternalProcess
	}
	fmt.Println("[Debug Process] processKey:", processKey)
	fmt.Println("[Debug Process] before setting, internalProcess string value:", context.Global().getValueByKey(internalProcess).GetStringValue())
	fmt.Println("[Debug Process] before setting, infraInternalProcess string value:", context.Global().getValueByKey(infraInternalProcess).GetStringValue())

	context.Global().setValueByKey(processKey, _processObject, consts.V8_PROPERTY_ATTRIBUTE_READONLY)

	fmt.Println("[Debug Process] after setting, internalProcess string value:", context.Global().getValueByKey(internalProcess).GetStringValue())
	fmt.Println("[Debug Process] after setting, infraInternalProcess string value:", context.Global().getValueByKey(infraInternalProcess).GetStringValue())
}
