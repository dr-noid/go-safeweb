// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// echo implements a simple echo server which listents on localhost:8080.
//
// Endpoints:
//   - /echo?message=YourMessage
//   - /uptime
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/google/safehtml"
	"github.com/google/safehtml/template"

	"github.com/google/go-safeweb/examples/echo/security/web"
	"github.com/google/go-safeweb/safehttp"
)

var start time.Time

func main() {
	port := 8080
	addr := fmt.Sprintf("localhost:%d", port)
	m := web.NewMuxConfigDev(port).Mux()

	// Before the program starts, create data structures that hold coverage information about the branches;
	safehttp.InitializeCoverageMap()

	m.Handle("/echo", safehttp.MethodGet, safehttp.HandlerFunc(echo))
	m.Handle("/uptime", safehttp.MethodGet, safehttp.HandlerFunc(uptime))

	start = time.Now()
	log.Printf("Visit http://%s\n", addr)
	log.Printf("Listening on %s...\n", addr)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// At the end of the program, write all information about the branches taken to a file or console.

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		for range c {
			for k, v := range safehttp.Coverage {
				if v {
					fmt.Printf("Branch %s was taken\n", k)
				} else {
					fmt.Printf("Branch %s was not taken\n", k)
				}
			}
			wg.Done()
		}
	}()
	wg.Wait()
	http.ListenAndServe(addr, m)
}

func echo(w safehttp.ResponseWriter, req *safehttp.IncomingRequest) safehttp.Result {
	q, err := req.URL().Query()
	if err != nil {
		safehttp.Coverage["echo-1"] = true
		return w.WriteError(safehttp.StatusBadRequest)
	}
	x := q.String("message", "")
	if len(x) == 0 {
		safehttp.Coverage["echo-2"] = true
		return w.WriteError(safehttp.StatusBadRequest)
	}
	safehttp.Coverage["echo-3"] = true
	return w.Write(safehtml.HTMLEscaped(x))
}

var uptimeTmpl *template.Template = template.Must(template.New("uptime").Parse(
	`<h1>Uptime: {{ .Uptime }}</h1>
{{- if .EasterEgg }}<h1>You've found an easter egg using "{{ .EasterEgg }}". Congrats!</h1>{{ end -}}`))

func uptime(w safehttp.ResponseWriter, req *safehttp.IncomingRequest) safehttp.Result {
	var x struct {
		Uptime    time.Duration
		EasterEgg string
	}
	x.Uptime = time.Since(start)

	// Easter egg handling.
	q, err := req.URL().Query()
	if err != nil {
		safehttp.Coverage["uptime-1"] = true
		return w.WriteError(safehttp.StatusBadRequest)
	}
	if egg := q.String("easter_egg", ""); len(egg) != 0 {
		safehttp.Coverage["uptime-2"] = true
		x.EasterEgg = egg
	}

	safehttp.Coverage["uptime-3"] = true
	return safehttp.ExecuteTemplate(w, uptimeTmpl, x)
}
