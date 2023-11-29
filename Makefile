# Change this to your own Go module

# Note
# $ make build => menghapus folder protogen, compile file proto, & download dependensi go
# $ make run => menjalankan file main.go
# $ make $execute => execute build dan menjalankan aplikasi

GO_MODULE := my-protobuf

.PHONY: tidy
tidy:
	go mod tidy


.PHONY: clean
clean:
ifeq ($(OS), Windows_NT)
#	if exist "protogen" rd /s /q protogen
else
	rm -fR ./protogen
endif


.PHONY: protoc
# (un)comment or add folder that contains proto as required
protoc:
	protoc --go_opt=module=${GO_MODULE} --go_out=. \
	./proto/basic/*.proto \
	./proto/dummy/*.proto \
	./proto/jobsearch/*.proto \


.PHONY: build
build: clean protoc tidy


.PHONY: run
run:
	go run main.go


.PHONY: execute
execute: build run

# jalankan make protoc-validate
.PHONY: protoc-validate
protoc-validate:
	protoc --validate_out="lang=go:./generated" --go_opt=module=${GO_MODULE} --go_out=. ./proto/car/*.proto
