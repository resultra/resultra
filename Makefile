DEPTH = .
include $(DEPTH)/build/common.mk
	
realclean:
	$(RM_DIR) vendor
	
install:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure 
	
all: prebuild