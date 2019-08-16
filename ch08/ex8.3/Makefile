client:
	go run netcat3.go

server-start:
	go run reverve2/main.go &

server-stop:
	ps -f|grep reverve2/main.go | grep -v grep | awk '{print $$2}' | xargs kill -9
