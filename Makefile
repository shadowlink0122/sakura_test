FILE = main

default:
	go build ${FILE}.go
	mv main main.cgi

clean:
	rm -rf main.cgi
