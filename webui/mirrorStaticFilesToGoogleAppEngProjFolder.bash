#!/usr/bin/env bash
# 
# This script mirrors all the static files which are maintained alongside and with the same
# folder structure as the golang files. The files are copied to the project folder for
# building the Google App Engine (GAE) project. See the README in ../googleAppEng/ for
# more details.

# The destination directory first needs to be removed. In the case where files have been moved
# around, this prevents old files from sticking around.
rm -rf ../serverMain/static/

# Mirror all the static files from the source code (development) tree to an equivalent
# structure for reference by the GAE executable. Notably, golang source files are not copied
# over, since they are built into the application, not referenced as static files in the 
# running GAE executable.
find -E . -type f -regex ".*\.(png|css|html|jpg|js)"  | cpio -p -d -v ../serverMain/static/
