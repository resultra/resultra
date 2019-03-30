
# This file is part of the Resultra project.
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.
DEPTH = ..
include $(DEPTH)/build/common.mk


localContextDockerSrcFiles:
	tar -czf static.tar.gz -C $(DEPTH)/build/dest/static .
	tar -czvf factoryTemplates.tar.gz -C $(DEPTH)/build/dest/factoryTemplates .
	cp ../build/dest/bin/resultraServer .

dockerImage: localContextDockerSrcFiles
	sudo docker build -t resultra-server:latest .

dockerdist: localContextDockerSrcFiles
	$(INSTALL_FILE) static.tar.gz $(DOCKER_IMAGE_DIR)
	$(INSTALL_FILE) factoryTemplates.tar.gz $(DOCKER_IMAGE_DIR)
	$(INSTALL_FILE) resultraServer $(DOCKER_IMAGE_DIR)
	$(INSTALL_FILE) Dockerfile $(DOCKER_IMAGE_DIR)
	$(INSTALL_FILE) docker-compose.yml $(DOCKER_IMAGE_DIR)

# TODO - Incorporate the latest version number into the file name, based upon
# the semantic version number stored in git. This can possibly be done using
# the `git describe` CLI. 
DOCKER_DIST_TAR_FILE = $(DOCKER_DIST_ROOT_DIR)/resultra.linux-amd64.tar.gz

# We want a tar file which includes all the files in the resultra
# top-level directory inside $(DOCKER_DIST_ROOT_DIR). This allows the
# whole tar file to be expanded into /usr/local for an installation,
# with the resultra top-level directory appended. This is the same
# scheme by which Golang is distributed.
dockerpkg:
	$(RM) $(DOCKER_DIST_TAR_FILE)
	tar -czvf $(DOCKER_DIST_TAR_FILE) -C $(DOCKER_DIST_ROOT_DIR) resultra

all:
	

clean:
	$(RM) *.gz resultraServer
