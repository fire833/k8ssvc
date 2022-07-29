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

package routes

import (
	"encoding/json"
	"time"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func RegisterStatusHandlers(router *router.Group) {

	router.Handle(fasthttp.MethodGet, "/readyz", func(ctx *fasthttp.RequestCtx) {
		type ready struct {
			IsReady   bool      `json:"ready"`
			Timestamp time.Time `json:"timestamp"`
		}

		data, _ := json.Marshal(&ready{
			IsReady:   true,
			Timestamp: time.Now(),
		})
		ctx.Response.SetBody(data)
		ctx.Response.SetStatusCode(200)
	})

	router.Handle(fasthttp.MethodGet, "/livez", func(ctx *fasthttp.RequestCtx) {
		type live struct {
			IsLive    bool      `json:"ready"`
			Timestamp time.Time `json:"timestamp"`
		}

		data, _ := json.Marshal(&live{
			IsLive:    true,
			Timestamp: time.Now(),
		})
		ctx.Response.SetBody(data)
		ctx.Response.SetStatusCode(200)
	})

}
