package main

import "github.com/GleanCoder/podcast/auth"

func main() {
	auth.LoginCredentials("GleanCoder", "secreat")
}

/*
- we can reuse the code in other files by creating a package and importing it into other files
- we can speed up compilation, cause go compile the packages which are changed and not the whole project
- before to create a package we need to create a folder with the name of the package and then create a file with the same name of the package and then we can create our functions and variables in that file
- but now we can create a package in the same folder as the main.go file and then we can import it into the main.go file
- there is module system in go which allows us to create a package and import it into other files
- we can create modules in go by using the command "go mod init <module_name>" and then we can create a package in that module and import it into other files
- there is a convention in module_name give the github.com/<username>/<repository_name> and then we can import it into other files
*/

/*
- we can install third party packages in go by using the command "go get <package_name>" and then we can import it into other files
- we can use the command "go get -u <package_name>" to update the package to the latest version
- we can use the command "go get <package_name>@<version>" to install a specific version of the package
-we can use go tiddy to manage our dependencies and we can use the command "go mod tidy" to remove the unused dependencies from our go.mod file
- also we can use tiddy to add the missing dependencies to our go.mod file
*/
