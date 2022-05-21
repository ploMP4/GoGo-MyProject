INCLUDEDIR = includes/

CC = g++

build: 
    @echo Compiling...
    $(CC) -o bin/main -I $(INCLUDEDIR) ./src/*.cpp

run: build
    ./bin/main