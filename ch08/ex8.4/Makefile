server-start:
	go run reverve3/main.go

server-stop:
	# 以下はバグあり。子プロセスをKILLする必要あり。
	ps -f |grep reverve3/main.go | grep -v grep | awk '{print $$2}' | xargs kill -9 

client:
	go run ../ex8.3/netcat3.go
