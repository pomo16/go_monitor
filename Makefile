RUN_NAME="go_monitor"

run:
	./output/${RUN_NAME}

build:
	gofmt -w .
	chmod +x build.sh
	sh build.sh

clean:
	rm -rf output