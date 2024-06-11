#!/bin/sh

install_dev_binaries() {
    echo "Installing dependent package binaries..."
    go install github.com/cespare/reflex@v0.3.1
    echo "Finished."
}

if [[ -z "$1" ]]; then
	if [[ $ENVIRONMENT == "local" ]]; then
        install_dev_binaries
        go mod tidy
        go mod vendor
        cp reflex.conf cmd/
        cd cmd
        reflex -c reflex.conf
	else
        echo "bulid the src files"
    	# build the src files
	fi
fi