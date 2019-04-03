# Developing for Resultra on Ubuntu Linux

The following is a guide for setting up and building Resultra on a system running Ubuntu based Linux.

## System Requirements

The development system (or virtual machine), should have at least:

* Ubuntu 18.04 LTS
* 4GB Memory
* 2-4 Processors - builds will run on a single processor, but the build runs in parallel and will complete faster with more processors.
* 20GB disk space

## Update the System

Before installing tools needed to build the project, it is a good idea to update the system as a whole:

	$ sudo apt-get update
	$ sudo apt-get upgrade

## Installing Build Dependencies
	
Install packages needed to build and test the project:

	$ sudo apt-get install build-essential git python-pip sqlite3
	
Install the latest version of Node.js:

	$ sudo apt-get install curl
	$ curl -sL https://deb.nodesource.com/setup_8.x | sudo -E bash -
	$ sudo apt-get install -y nodejs
	
## Install Go (Golang) Development Tools

Download and install the Go software development tools:

	$ curl https://dl.google.com/go/go1.12.1.linux-amd64.tar.gz --output ~/Downloads/go1.12.1.linux-amd64.tar.gz
	$ sudo tar -C /usr/local -zxf ~/Downloads/go1.12.1.linux-amd64.tar.gz
	
Append the following to ~/.bashrc:

	export GOROOT=/usr/local/go
	export GOPATH=$HOME/go
	export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
	
Verify the installation:

	$ source .bashrc
	$ which go
	$ go version
	
## Install Python and Node.js Packages

Install the requests package for system testing:

	$ sudo pip install requests
	
Install the Gulp Node.js package for compiling resources in the build:

	$ sudo npm install -g gulp
	$ sudo npm install -g gulp-cli 

## Install Tools for Cross-compiling to Windows

Install 64 bit gcc-mingw for cross-compiling the Golang sources for Windows (specifically needed by the SQLite driver):

	$ sudo apt-get install gcc-mingw-w64-i686
	
Install Wine for Building the Electron Windows Client on Linux, following the winthe following instructions: https://linuxconfig.org/install-wine-on-ubuntu-18-04-bionic-beaver-linux.

Add 32 bit support:

	$ sudo dpkg --add-architecture i386
	
Add the repository:

	$ wget -qO- https://dl.winehq.org/wine-builds/winehq.key | sudo apt-key add -
	$ sudo apt-add-repository 'deb http://dl.winehq.org/wine-builds/ubuntu/ bionic main'
	
Update the repositories and install the latest/stable version of wine:

	$ sudo apt-get update
	$ sudo apt-get install --install-recommends winehq-stable

Verify the installation:

	$ wine --version
	
## Get the Source Code

Resultra's development includes Go (Golang) source code, and therefore needs to be downloaded in a specific location below ```$GOPATH```: 

	$ mkdir -p $GOPATH/src/github.com/resultra
	$ cd $GOPATH/src/github.com/resultra
	$ git clone https://github.com/resultra/resultra.git

	
## Build the System

	$ cd $GOPATH/src/github.com/resultra/resultra/build
	$ ./build.py
 
	
 


