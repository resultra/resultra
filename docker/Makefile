
localContextDockerSrcFiles:
	tar -czvf dest.tar.gz -C ../build/dest .

dockerImage: localContextDockerSrcFiles
	sudo docker build -t resultra-server:latest .

all: dockerImage
	

clean:
	rm dest.tar.gz
