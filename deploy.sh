GOOS=linux GOARCH=amd64 go build -ldflags='-w -s' -o ./main .
scf deploy -f
rm -f main