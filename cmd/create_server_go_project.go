package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

func createServerGoProject(home string, packageName string) {
	var data = make(map[string]interface{})
	data["PackageName"] = packageName

	os.Mkdir(home, os.ModePerm)

	writeFile(path.Join(home, ".gitignore"), "asset/templates/server/go/.gitignore", data)
	writeFile(path.Join(home, "Makefile"), "asset/templates/server/go/Makefile", data)
	writeFile(path.Join(home, "cmd_root.go"), "asset/templates/server/go/cmd_root.go", data)
	writeFile(path.Join(home, "cmd_serve.go"), "asset/templates/server/go/cmd_serve.go", data)
	writeFile(path.Join(home, "cmd_version.go"), "asset/templates/server/go/cmd_version.go", data)

	writeFile(path.Join(home, "go.mod"), "asset/templates/server/go/go.mod", data)

	writeFile(path.Join(home, "go.sum"), "asset/templates/server/go/go.sum", data)
	writeFile(path.Join(home, "main.go"), "asset/templates/server/go/main.go", data)
	writeFile(path.Join(home, "route.go"), "asset/templates/server/go/route.go", data)

	os.Mkdir(path.Join(home, "version"), os.ModePerm)
	writeFile(path.Join(home, "version", "version.go"), "asset/templates/server/go/version/version.go", data)

	os.Mkdir(path.Join(home, "server"), os.ModePerm)
	writeFile(path.Join(home, "server", "http.go"), "asset/templates/server/go/server/http.go", data)

	os.Mkdir(path.Join(home, "flows"), os.ModePerm)
	os.Mkdir(path.Join(home, "flows", "flow1"), os.ModePerm)
	writeFile(path.Join(home, "flows", "flow1", "flow.go"), "asset/templates/server/go/flows/flow1/flow.go", data)
	writeFile(path.Join(home, "flows", "flow1", "graph.go"), "asset/templates/server/go/flows/flow1/graph.go", data)
	writeFile(path.Join(home, "flows", "flow1", "node1.go"), "asset/templates/server/go/flows/flow1/node1.go", data)
	writeFile(path.Join(home, "flows", "flow1", "node2.go"), "asset/templates/server/go/flows/flow1/node2.go", data)
	writeFile(path.Join(home, "flows", "flow1", "node3.go"), "asset/templates/server/go/flows/flow1/node3.go", data)
	writeFile(path.Join(home, "flows", "flow1", "node4.go"), "asset/templates/server/go/flows/flow1/node4.go", data)
	writeFile(path.Join(home, "flows", "flow1", "node5.go"), "asset/templates/server/go/flows/flow1/node5.go", data)
	writeFile(path.Join(home, "flows", "flow1", "node6.go"), "asset/templates/server/go/flows/flow1/node6.go", data)

	cmd := exec.Command("make")
	cmd.Dir = home
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Println(string(out))
}
