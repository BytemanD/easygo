
function logInfo() {
    echo "INFO:" $@ 1>&2
}
function getVersionByGit(){
    tagRage=HEAD
    lastTag=$(git tag --sort=-taggerdate |tail -n1)
    if [[ "${lastTag}" != "" ]]; then
        tagRange="${tagRange...${lastTag}}"
    else
        lastTag="0.0.0"
    fi

    commitNum=$(git log --pretty=oneline ${tagRange} |wc -l)
    logInfo "last tag is ${lastTag}, commit num is ${commitNum}"

    if [[ ${commitNum} -eq 0 ]]; then
        version="${lastTag}"
    else
        version="${lastTag}.dev${commitNum}"
    fi
    echo ${version}
}

function main(){
    version=$(getVersionByGit)
    logInfo "set main.Version to ${version}"  1>&2
    go build  -ldflags "-X main.Version=${version}" cmd/magic-pocket.go
}

main
