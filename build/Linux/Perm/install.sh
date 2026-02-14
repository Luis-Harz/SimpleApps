#!/bin/bash
echo "----SimpleApps Permanent----"
echo "Press ENTER to start"
read
SCRIPT_PATH="$(readlink -f "$0")"
SCRIPT_DIR="$(dirname "$SCRIPT_PATH")"
cd "$SCRIPT_DIR" || exit 1
echo "This program includes st as bundled Terminal it's under the MIT License"
echo "Terminal LICENSE START"
cat Terminal-LC
echo "LICENSE END"
read
echo "Making Directory..."
sudo mkdir -p /opt/simpleapps
mkdir -p ~/.local/share/applications
echo "Copying Files..."
sudo cp ./Terminal /opt/simpleapps/
sudo cp ./icon.png /opt/simpleapps/
sudo cp ./wrapper.sh /opt/simpleapps/
sudo cp ./wrapper2.sh /opt/simpleapps/
sudo rsync -av --exclude 'Perm' ../ /opt/simpleapps/
echo "Making Symlinks"
sudo ln -sf /opt/simpleapps/wrapper.sh /usr/bin/simpleapps
sudo ln -sf /opt/simpleapps/wrapper2.sh /usr/bin/simpleapps-gui
echo "Making Desktop File"
DESKTOP_FILE="$HOME/Desktop/SimpleApps.desktop"
cat > "$DESKTOP_FILE" <<EOL
[Desktop Entry]
Type=Application
Name=SimpleApps
Comment=Run SimpleApps in bundled st terminal
Exec=simpleapps-gui
Icon=/opt/simpleapps/icon.png
Terminal=false
Categories=Utility;
EOL
cp "$DESKTOP_FILE" ~/.local/share/applications/
echo "Settings Permissions"
chmod +x "$DESKTOP_FILE"
sudo chmod +x /opt/simpleapps/SimpleApps
sudo chmod +x /opt/simpleapps/Terminal
echo "Updating Desktop Database"
update-desktop-database ~/.local/share/applications
echo "Setting Version to PERM"
echo "All Done!"
echo "Commands: simpleapps, simpleapps-gui"
