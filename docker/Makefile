
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

all:
	

clean:
	$(RM) *.gz resultraServer
