/*
*	Copyright (C) 2022  Kendall Tauser
*
*	This program is free software; you can redistribute it and/or modify
*	it under the terms of the GNU General Public License as published by
*	the Free Software Foundation; either version 2 of the License, or
*	(at your option) any later version.
*
*	This program is distributed in the hope that it will be useful,
*	but WITHOUT ANY WARRANTY; without even the implied warranty of
*	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
*	GNU General Public License for more details.
*
*	You should have received a copy of the GNU General Public License along
*	with this program; if not, write to the Free Software Foundation, Inc.,
*	51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
 */

package config

import (
	"os"

	"github.com/spf13/viper"
)

type k8sSMConfig struct {
	c *viper.Viper
}

func newSMConfig() *k8sSMConfig {
	return &k8sSMConfig{
		c: viper.New(),
	}
}

func (conf *k8sSMConfig) Initialize() error {
	cnf := conf.c

	cnf.AddConfigPath("/etc/k8ssm")
	cnf.AddConfigPath(os.Getenv("CONFIG"))
	cnf.AddConfigPath(".")

	// Default fail closed and disable tracking non-whitelisted ingresses.
	cnf.SetDefault(DefaultTrack, false)

	// Server configuration defaults
	cnf.SetDefault(ListenCert, "/etc/k8ssm/cert.crt")
	cnf.SetDefault(ListenKey, "/etc/k8ssm/cert.key")
	cnf.SetDefault(ListenPort, 443)
	cnf.SetDefault(DoTLS, true)

	return cnf.ReadInConfig()
}
