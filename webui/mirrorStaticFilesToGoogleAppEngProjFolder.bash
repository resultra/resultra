#!/usr/bin/env bash
# 
# This script mirrors all the static files which are maintained alongside and with the same
# folder structure as the golang files. The files are copied to the project folder for
# building the Google App Engine (GAE) project. See the README in ../googleAppEng/ for
# more details.
find -E . -type f -regex ".*\.(png|css|html|jpg|js)"  | cpio -p -d -v ../googleAppEng/static/
