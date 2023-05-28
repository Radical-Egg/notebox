TARGET=build/notebox

build:
	go build -trimpath -o $(TARGET) main.go

.PHONY: build

clean:
	rm -rf build/
run:
	go build -trimpath -o $(TARGET) main.go && ./build/notebox
