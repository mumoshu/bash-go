test:
	goyacc -o my.go my.go.y && go run main.go ~/.bash_profile
