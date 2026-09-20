package auth // its a convention to give the name of the package as the name of the folder

import "fmt" // if you want to import the package outside of the auth scope you need to make the function public by capitalizing the first letter of the function name or variable name
func LoginCredentials(username string, password string) {
	fmt.Printf("Logging with the credentials of %s and %s", username, password)
}
