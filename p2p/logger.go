/*
Copyright (C) BABEC. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package p2p

import (
	"chainmaker.org/chainmaker/logger/v2"
	liquid "chainmaker.org/chainmaker/net-liquid/liquidnet"
	"chainmaker.org/chainmaker/protocol/v2"
)

var GlobalNetLogger protocol.Logger

func init() {
	l := &logger.LogConfig{
		SystemLog: logger.LogNodeConfig{
			FilePath:        "./data/log/default.log",
			LogLevelDefault: "DEBUG",
			LogInConsole:    true,
			ShowColor:       false,
		},
	}
	logger.SetLogConfig(l)

	GlobalNetLogger = logger.GetLogger(logger.MODULE_NET)
	liquid.InitLogger(GlobalNetLogger, func(chainId string) protocol.Logger {
		return logger.GetLoggerByChain(logger.MODULE_NET, chainId)
	})
}
