## Learning About GoLang Packages

Packages are a way for golang to modularised its code, create library and a way to build modules. In Golang, a folder is known as a package. 

To be able to import packages, golang will try to find the packages from its $GOPATH/src/ directory and if the packages is not found then you cannot use the packages. To be able to perform a local import of the packages, you need to do the following:

1. Initialised a go mod root directory by running `go mod init example.com`. Doing this will set the root directory of your cwd as example.com and will create a go.mod file.
2. After establishing the root directory, golang will understand the reference to the root and be able to import packages with a local reference. When initialising your code add the extensions of the packages being imported, in this case its: `import (example.com/tempconv)`
3. Golang will load all of the packages and sub .go command in the directory as a reference.