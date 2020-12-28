all:
	quark web/
	go build

clean:
	rm -f app calculator.app/Contents/MacOS/Calculator

macOS:
	make all
	cp app calculator.app/Contents/MacOS/Calculator
