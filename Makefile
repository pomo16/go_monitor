RUN_NAME="go_monitor"

run:
	./output/${RUN_NAME}

build:
	gofmt -w .
	chmod +x build.sh
	sh build.sh

relog:
	rm -rf output/monitor_log
	mkdir output/monitor_log

clean:
	rm -rf output