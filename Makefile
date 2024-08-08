.PHONY: clean
clean:
	rm -f _main

target=main.go
build_cmd=go build -o _main
label_main: $(target)
	$(build_cmd)

.PHONY: clean
clean:
	tree

