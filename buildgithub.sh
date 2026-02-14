#!/bin/bash
VERSION=$(cat build/Linux/version.txt)
echo "Adding Files"
git add .
echo "Making Commit"
git commit -m "Update: $VERSION" || echo "Nothing to commit"
echo "Pushing"
git push origin main
gh release create "$VERSION" \
  build/Linux/SimpleApps \
  build/Windows/SimpleApps.exe \
  build/Linux.zip \
  build/Windows.zip \
  --title "SimpleApps v$VERSION" \
  --notes "Release $VERSION: Includes Linux and Windows binaries and zips"
