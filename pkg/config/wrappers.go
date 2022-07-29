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

import "time"

func (conf *K8sSMConfig) RegisterDefault(key string, value interface{}) {
	conf.c.SetDefault(key, value)
}

func (conf *K8sSMConfig) InConfig(key string) bool {
	return conf.c.InConfig(key)
}

func (conf *K8sSMConfig) IsSet(key string) bool {
	return conf.c.IsSet(key)
}

// Wrapper functions around the viper config container.

func (conf *K8sSMConfig) GetString(key string) string {
	return conf.c.GetString(key)
}

func (conf *K8sSMConfig) GetBool(key string) bool {
	return conf.c.GetBool(key)
}

func (conf *K8sSMConfig) GetInt(key string) int {
	return conf.c.GetInt(key)
}

func (conf *K8sSMConfig) GetUint(key string) uint {
	return conf.c.GetUint(key)
}

func (conf *K8sSMConfig) GetTime(key string) time.Time {
	return conf.c.GetTime(key)
}

func (conf *K8sSMConfig) WatchConfig() {
	conf.c.WatchConfig()
}
