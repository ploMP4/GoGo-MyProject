INCLUDEDIR = includes/

CC = g++

build: 
    @echo Compiling...
    $(CC) -o bin/main -O3 -I $(INCLUDEDIR) ./src/*.cpp

run: build
    ./bin/main