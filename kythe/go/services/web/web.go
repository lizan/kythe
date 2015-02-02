/*
 * Copyright 2015 Google Inc. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package web defines utility functions for exposing services over HTTP.
package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"kythe/go/util/httpencoding"
)

// ReadJSONBody reads the entire body of r and unmarshals it from JSON into v.
func ReadJSONBody(r *http.Request, v interface{}) error {
	rec, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("body read error: %v", err)
	}
	return json.Unmarshal(rec, v)
}

// WriteJSONResponse encodes v as JSON and writes it to w.
func WriteJSONResponse(w http.ResponseWriter, r *http.Request, v interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	cw := httpencoding.CompressData(w, r)
	defer cw.Close()
	return json.NewEncoder(cw).Encode(v)
}

// Arg returns the first query value for the named parameter or "" if it was not
// set.
func Arg(r *http.Request, name string) string {
	args := r.URL.Query()[name]
	if len(args) == 0 {
		return ""
	}
	return args[0]
}

// ArgOrNil returns a pointer to first query value for the named parameter or
// nil if it was not set.
func ArgOrNil(r *http.Request, name string) *string {
	arg := Arg(r, name)
	if arg == "" {
		return nil
	}
	return &arg
}

// TrimPath returns the URL path of the given request with the given prefix
// trimmed from it.
func TrimPath(r *http.Request, prefix string) string {
	return strings.TrimPrefix(filepath.Clean(r.URL.Path), prefix)
}
