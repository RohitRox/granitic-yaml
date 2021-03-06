package main

import (
	"bufio"
	"github.com/graniticio/granitic/cmd/grnc-project/generate"
	"path/filepath"
)

func main() {

	pg := new(generate.ProjectGenerator)
	pg.CompWriterFunc = writeComponentsFile
	pg.ConfWriterFunc = writeConfigFile
	pg.MainFileFunc = writeMainFile
	pg.ToolName = "grnc-yaml-project"
	pg.Generate()

}

func writeConfigFile(confDir string, pg *generate.ProjectGenerator) {

	compFile := filepath.Join(confDir, "config.yml")
	f := pg.OpenOutputFile(compFile)

	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString("# Configuration you want to make available to your components\n")
	w.Flush()

}

func writeComponentsFile(compDir string, pg *generate.ProjectGenerator) {

	compFile := filepath.Join(compDir, "components.yml")
	f := pg.OpenOutputFile(compFile)

	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString("packages:\n")
	w.WriteString("  # List of package names (e.g granitic.ws) referenced by components in this file.\n")
	w.WriteString("components:\n")
	w.WriteString("  # Definition of components you want to be managed by Granitic")
	w.Flush()

}

func writeMainFile(w *bufio.Writer, projectPackage string) {

	changePackageComment := "  //Change to a non-relative path if you want to use 'go install'"

	w.WriteString("package main\n\n")
	w.WriteString("import \"github.com/graniticio/granitic-yaml\"\n")
	w.WriteString("import \"")
	w.WriteString(projectPackage)
	w.WriteString("/bindings\"")
	w.WriteString(changePackageComment)
	w.WriteString("\n\n")
	w.WriteString("func main() {\n")
	w.WriteString("\tgranitic_yaml.StartGraniticWithYaml(bindings.Components())\n")
	w.WriteString("}\n")

}
