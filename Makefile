
# https://github.com/trstringer/cli-debugging-cheatsheets/blob/master/go.md
#
agenda.exe:
	go build agenda.go

clean:
	rm agenda.exe

test:
	go test
	go run agenda.go -path agenda_test_file.md