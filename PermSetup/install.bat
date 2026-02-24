@echo off
echo ----SimpleApps Permanent----
echo This installer will permanently install SimpleApps
echo Press ENTER to start
pause

echo Making Directory...
mkdir "C:\SimpleApps"

echo Copying Files...
xcopy "..\*" "C:\SimpleApps\" /s /e /y

echo Creating Desktop Icon...
powershell -Command "$WshShell = New-Object -ComObject WScript.Shell; $Shortcut = $WshShell.CreateShortcut(\"$env:USERPROFILE\Desktop\SimpleApps.lnk\"); $Shortcut.TargetPath = 'C:\SimpleApps\SimpleLauncher.bat'; $Shortcut.WorkingDirectory = 'C:\SimpleApps'; $Shortcut.IconLocation = 'C:\SimpleApps\icon.ico'; $Shortcut.Save()"

echo Adding to PATH...
setx PATH "%PATH%;C:\SimpleApps"

echo All Done!
pause
