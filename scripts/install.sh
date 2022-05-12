DIST_PATH=./dist/gogo

APPLICATION_PATH=~/.gogo
BINARY_PATH=$APPLICATION_PATH/bin
SETTINGS_FILE_PATH=$BINARY_PATH/settings.json

if [ ! -d $APPLICATION_PATH ]; then 
    mkdir $APPLICATION_PATH
    mkdir $BINARY_PATH
fi

mv $DIST_PATH $BINARY_PATH

if [ ! -f $SETTINGS_FILE_PATH ]; then 
    touch $SETTINGS_FILE_PATH

    echo '{ 
    "config-path": "/home/'$USER'/Documents/GitHub/GoGoProject/examples/config"
}' > $SETTINGS_FILE_PATH
fi

# Add to PATH
if ! grep -q "path+=('/home/$USER/.gogo/bin')" ~/.zshrc; then
    sed -i '$ d' ~/.zshrc

    echo "path+=('/home/$USER/.gogo/bin')
    
export PATH" >> ~/.zshrc
fi