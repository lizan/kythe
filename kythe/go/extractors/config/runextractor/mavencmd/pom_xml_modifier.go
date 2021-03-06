/*
 * Copyright 2018 The Kythe Authors. All rights reserved.
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

package mavencmd

import (
	"fmt"
	"os"

	"github.com/beevik/etree"
)

// preProcessPomXML takes a pom.xml file and either verifies that it already has
// the bits necessary to specify a separate compiler on commandline, or adds
// functionality by dropping in a maven-compiler-plugin to the build.
//
// Note this potentially overwrites the input file, even if it returns an error,
// so make a copy beforehand if you need to keep the original.
func preProcessPomXML(pomXMLFile string) error {
	doc := etree.NewDocument()
	err := doc.ReadFromFile(pomXMLFile)
	if err != nil {
		return fmt.Errorf("reading XML file %s: %v", pomXMLFile, err)
	}
	if hasCompilerPlugin(doc) {
		return nil
	}
	if err := insertCompilerPlugin(doc); err != nil {
		return err
	}
	f, err := os.OpenFile(pomXMLFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("opening file: %v", err)
	}
	doc.Indent(2)
	doc.WriteTo(f)
	return f.Close()
}

func hasCompilerPlugin(doc *etree.Document) bool {
	for _, p := range doc.FindElements("//project/build/plugins/plugin/artifactId") {
		if p.Text() == "maven-compiler-plugin" {
			return true
		}
	}
	return false
}

// appendCompilerPlugin modifies a doc containing the root element of a pom.xml
// configuration, by injecting a custom Maven compiler <plugin> element into the
// plugin list for the top-level project.  This allows mvn install commands to
// specify --Dmaven.compiler.executable on commandline and use a separate
// compiler.
//
// Note this is unlike the gradle.build modification, where a separate
// commandline argument is not necessary.
//
// An example modification can be found in the tests.  We would expect
// testdata/other-pom.xml to be transformed into testdata/modified-pom.xml
func insertCompilerPlugin(doc *etree.Document) error {
	project := doc.SelectElement("project")
	if project == nil {
		return fmt.Errorf("no top level <project> element")
	}
	build := project.SelectElement("build")
	if build == nil {
		build = project.CreateElement("build")
	}
	plugins := build.SelectElement("plugins")
	if plugins == nil {
		plugins = build.CreateElement("plugins")
	}
	newPlugin := plugins.CreateElement("plugin")
	groupID := newPlugin.CreateElement("groupId")
	groupID.SetText("org.apache.maven.plugins")
	artifactID := newPlugin.CreateElement("artifactId")
	artifactID.SetText("maven-compiler-plugin")
	version := newPlugin.CreateElement("version")
	version.SetText("3.7.0")
	configuration := newPlugin.CreateElement("configuration")
	source := configuration.CreateElement("source")
	source.SetText("1.8")
	target := configuration.CreateElement("target")
	target.SetText("1.8")
	return nil
}
