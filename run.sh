#!/bin/bash

src="$GOPATH/src"
vp="github.com/varunamachi/vaali"
sp="${src}/${vp}/"
dp="${src}/github.com/varunamachi/spac/vendor/${vp}/"

rsync -av --exclude='vendor/' --exclude='.git/' "${sp}" "${dp}" && go install 

spac "$@"
