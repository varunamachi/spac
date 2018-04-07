#!/bin/bash

RHOST=${REMOTE_HOST:="raspberrypi"}
RUSER=${REMOTE_USER:="pi"}

echo "Building..."
GOARCH=arm go build -o "$GOBIN"/linux_arm/spoon || exit -1

echo "Copying..."
scp "$GOPATH"/bin/linux_arm/spoon "$RUSER"@"$RHOST":/home/"$RUSER" || exit -2

echo "Executing..."
ssh "${RUSER}@${RHOST}" /home/"$RUSER"/spoon || exit -3