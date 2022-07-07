for id in 1 2 3 4 5 6 7 8 9 10 ; do go run client/main.go "$id" 3 localhost:8001 localhost:8002 localhost:8003 & done
