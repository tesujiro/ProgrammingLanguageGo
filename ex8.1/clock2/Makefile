build:
	go build -o clock2

test-start:
	TZ=US/Eastern    ./clock2 -port 8010 &
	TZ=Asia/Tokyo    ./clock2 -port 8020 &
	TZ=Europe/London ./clock2 -port 8030 &

test-stop:
	ps -ef|grep clock2| grep -v grep | awk '{print $$2}' | xargs kill -9
