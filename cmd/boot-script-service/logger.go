/*
 * Boot Script Server
 *
 * The boot script server will collect all information required to produce an
 * iPXE boot script for each node of a system.  This script will the be
 * generated on demand and delivered to the requesting node during an iPXE
 * boot.  The main items the script will deliver are the kernel image URL/path,
 * boot arguments, and the initrd URL/path.  Note that the kernel and initrd
 * images are specified with a URL or path.  A plain path will result in a tfpt
 * download from this server.  If a URL is provided, it can be from any
 * available service which iPXE supports, and any location that the iPXE client
 * has access to. It is not restricted to a particular Cray provided service.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		debugf("%s %s %s %s", r.Method, r.RequestURI, name, time.Since(start))
	})
}
