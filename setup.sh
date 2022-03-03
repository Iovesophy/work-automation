#!/bin/sh -eu

printf "See this url=>[ chrome://settings/help ] with chrome app. \n"
printf "Please input chrome version (※ No omission: xx.x.xxx.xxx):"
read -r VERSION
echo "$VERSION"
open "https://chromedriver.storage.googleapis.com/index.html?path=${VERSION}/"

printf "download ok?"
read

# install chrome driver
mv ~/Downloads/chromedriver_mac*.zip .
unzip chromedriver_mac*.zip
cp chromedriver /usr/local/bin
xattr -d com.apple.quarantine chromedriver
rm -f chromedriver chromedriver_mac*.zip

echo 1 > config/touch.log
