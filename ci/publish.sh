#!/bin/bash
echo "$0: start copying packages"

TARGET=/export
if [ $# -eq 1 ]
then
    TARGET=$1
fi

if [ ! -d "$TARGET" ]
then
    echo "$0: Error: target dir $TARGET does not exist"
    exit 1
fi

cp -v  /tmp/*.deb "$TARGET"