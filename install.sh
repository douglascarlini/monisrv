#!/bin/bash
os=$(uname -s)
arch=$(uname -m)

# ONLY WORKS IN LINUX X86
if [[ "$os" == "Linux" && "$arch" == "x86_64" ]]; then

    # INSTALL WGET
    apt update && apt install -y wget

    # DOWNLOAD GO 1.20
    wget https://go.dev/dl/go1.20.5.linux-amd64.tar.gz

    # EXTRACT GO FILES
    rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.5.linux-amd64.tar.gz

    # EXPORT GO TO PATH
    export PATH=$PATH:/usr/local/go/bin

    # CHECK GO VERSION
    go version

else

    echo "[ERROR] OS is not Linux AMD64"

fi