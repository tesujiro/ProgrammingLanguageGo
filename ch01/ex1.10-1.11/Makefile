all:
	go run main.go https://golang.org http://gopl.io https://godoc.org

long:
	curl http://s3.amazonaws.com/alexa-static/top-1m.csv.zip -o top-1m.csv.zip
	unzip top-1m.csv.zip
	cat top-1m.csv | awk -F, '{print $$2}' | xargs go run main.go
