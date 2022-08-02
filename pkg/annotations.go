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

package pkg

const (
	// The swagger URI annotations tells the k8s looking glass which HTTP endpoint to hit on
	// the ingress service for acquiring the swagger.json spec for that API, if it exists.
	//
	// If an ingress is instantiated without this annotation, the controller will not look for
	// a swagger endpoint with the ingress, and will not advertise the hostname as having a
	// swagger spec to end users.
	SwaggerURIAnnotation string = "k8s-looking-glass.io/swagger"

	// The LookingGlassAllowed annotation explicitly marks an ingress and its corresponding
	// hosts as being allowed to be scraped by the looking glass. The endpoints for that
	// ingress will be advertised in the catalog.
	//
	// Can be either set to "true" or "false".
	LookingGlassAllowed string = "k8s-looking-glass.io/allowed"

	// The LookingGlassAllowed annotation explicitly marks an ingress and its corresponding
	// hosts as not being allowed to be scraped by the looking glass. If LookingGlassAllowed
	// is also set to true, then the service fails closed and that ingress resource is not
	// tracked by the service catalog.
	//
	// Can be either set to "true" or "false".
	LookingGlassBlocked string = "k8s-looking-glass.io/blocked"
)
