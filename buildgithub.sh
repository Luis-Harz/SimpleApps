VERSION=$(cat build/Linux/version.txt)
gh release create "$VERSION" \
  build/Linux/SimpleApps build/Windows/SimpleApps.exe \
  build/Linux.zip build/Windows.zip \
  --title "SimpleApps v$VERSION" \
  --notes "Release $VERSION: a new release of SimpleApps"
