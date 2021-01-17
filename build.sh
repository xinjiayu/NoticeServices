#!/bin/sh

BuildVersion=`git describe --abbrev=0 --tags`
BuildTime=`date +%FT%T%z`
CommitID=`git rev-parse HEAD`


function help() {
    echo "$0 linux|windows|mac"
}

function linux(){

    CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=x86_64-linux-musl-gcc CGO_LDFLAGS="-static" go build -a -ldflags "-w -s -X main.BuildVersion=${BuildVersion} -X main.CommitID=${CommitID} -X main.BuildTime=${BuildTime}"

    copyFile
    cp curl.sh bin/

    cp NoticeServices bin/

    rm -f NoticeServices

}
function windows(){
    CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -ldflags "-w -s -X main.BuildVersion=${BuildVersion} -X main.CommitID=${CommitID} -X main.BuildTime=${BuildTime}"

    copyFile
    cp NoticeServices.exe bin/

    rm -f NoticeServices.exe

}
function mac(){
    go build -ldflags "-w -s -X main.BuildVersion=${BuildVersion} -X main.CommitID=${CommitID} -X main.BuildTime=${BuildTime}"

    copyFile
    cp curl.sh bin/

    cp NoticeServices bin/

    rm -f NoticeServices


}

function copyFile() {
  echo "主程序编译完成"

    rm -rf bin
    mkdir bin
    cp -r document/. bin/document/
    cp -r template/. bin/template/
    cp -r public/. bin/public/
    cp -r config/. bin/config/
    cp -r db/. bin/db/

        plugin

}

function plugin() {

    cd ./plugins

    ./build.sh

    cd ../

    mkdir bin/plugins
    mkdir bin/plugins/mail
    mkdir bin/plugins/webhook

    cp plugins/mail/mail.so bin/plugins/mail
    cp plugins/webhook/webhook.so bin/plugins/webhook

    cp plugins/mail/config.toml bin/plugins/mail
    cp plugins/webhook/config.toml bin/plugins/webhook

    rm -f plugins/mail/mail.so
    rm -f plugins/webhook/webhook.so

}


if [ "$1" == "" ]; then
    help
elif [ "$1" == "linux" ];then
    linux
elif [ "$1" == "windows" ];then
    windows
elif [ "$1" == "mac" ];then
    mac
fi