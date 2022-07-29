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

// Documentation for Kubernetes Service Catalog.
//
// This service finds exposed services on your Kubernetes cluster and
// provides a simple catalog for listing all ingresses.
//
//	Version: 0.0.1
//	BasePath: /
// 	License: GPL V2.0 https://opensource.org/licenses/gpl-2.0.php
//	Contact: Kendall Tauser <kttpsy@gmail.com>
//
//	Consumes:
//	- application/json
// 	- application/yaml
//	- application/xml
//
//	Produces:
// 	- application/json
//	- application/yaml
//
// swagger:meta
package apitypes
