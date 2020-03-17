RUN_NAME="go_monitor"

run:
	cd ./output/ && ./${RUN_NAME}

build:
	gofmt -w .
	chmod +x build.sh
	sh build.sh

relog:
	rm -rf output/${RUN_NAME}_log
	mkdir output/${RUN_NAME}_log

clean:
	rm -rf output