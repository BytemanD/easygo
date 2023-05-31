
function logInfo() {
    echo "INFO:" $@ 1>&2
}

function main(){
    version=$(go run cmd/magic-pocket.go version)
    logInfo "set main.Version to ${version}"  1>&2
    go build  -ldflags "-X main.Version=${version}" cmd/magic-pocket.go
}
main
