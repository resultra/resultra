Vendor packages depended upon from this project are managed through Go's dep tool.
To install the dep tool, do the following:

$ go get -u github.com/golang/dep/cmd/dep

After this was installed the first time, the following was used to initialized the 
dependencies:

$ dep init

Thereafter, the following can be performed to check the dependencies:

$ dep ensure