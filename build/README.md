# Building Resultra

# Top-level build script

build.py is the top-level build script for this project. This script runs a phased build. To run a debug build for development, invoke the following:

$ ./build.py

Or, to run a release build, invoke the following:

$ ./build.py --release

Other options are described by passing the --help option to the build script.

## Managing Golang Dependencies

Vendor packages depended upon from this project are managed through Go's dep tool.
To install the dep tool, do the following:

$ go get -u github.com/golang/dep/cmd/dep

After this was installed the first time, the following was used to initialized the 
dependencies:

$ dep init

Thereafter, the following can be performed to check the dependencies:

$ dep ensure