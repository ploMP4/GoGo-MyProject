DIST_PATH=./dist/gogo

APPLICATION_PATH=$HOME/.gogo
BINARY_PATH=$APPLICATION_PATH/bin
SETTINGS_FILE_PATH=$BINARY_PATH/settings.json

go build -o ./dist/gogo ./cmd/...

if [ ! -d $APPLICATION_PATH ]; then 
    mkdir $APPLICATION_PATH
    mkdir $BINARY_PATH
fi

mv $DIST_PATH $BINARY_PATH

rmdir ./dist

if [ ! -f $SETTINGS_FILE_PATH ]; then 
    touch $SETTINGS_FILE_PATH

    echo '{ 
    "config-path": "'$APPLICATION_PATH'/config"
}' > $SETTINGS_FILE_PATH
fi
