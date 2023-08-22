#!/bin/bash

DIST_PATH=./dist/gogo

APPLICATION_PATH=$HOME/.gogo
BINARY_PATH=$APPLICATION_PATH/bin
SETTINGS_FILE_PATH=$BINARY_PATH/settings.yaml
GADGET_PATH=$APPLICATION_PATH/gadgets
TEMPLATES_PATH=$APPLICATION_PATH/templates

go build -o ./dist/gogo ./cmd/...

if [ ! -d "$APPLICATION_PATH" ]; then 
    mkdir "$APPLICATION_PATH"
    mkdir "$BINARY_PATH"
    mkdir "$GADGET_PATH"
    mkdir "$TEMPLATES_PATH"
fi

mv $DIST_PATH "$BINARY_PATH"

rm -rf ./dist

if [ ! -f "$SETTINGS_FILE_PATH" ]; then 
    touch "$SETTINGS_FILE_PATH"

    echo 'gadget-path: "'$APPLICATION_PATH'/gadgets"' > "$SETTINGS_FILE_PATH"
    echo 'template-path: "'$APPLICATION_PATH'/templates"' >> "$SETTINGS_FILE_PATH"
fi
