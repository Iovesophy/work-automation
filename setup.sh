#!/bin/sh -eu

printf "See this url=>[ chrome://settings/help ] with chrome app. \n"
printf "Please input chrome version (※ No omission: xx.x.xxx.xxx):"
read -r VERSION
open "https://chromedriver.storage.googleapis.com/index.html?path=${VERSION}/"

printf "Download ok?"
read -r

# install chrome driver
mv ~/Downloads/chromedriver_mac*.zip .
unzip chromedriver_mac*.zip
cp chromedriver /usr/local/bin
xattr -d com.apple.quarantine chromedriver
rm -f chromedriver chromedriver_mac*.zip

printf "Please input your email:"
read -r JOBEMAIL
printf "email: %s\n" "${JOBEMAIL}" > config/config.yml
printf "1" > config/touch.log

printf "Setup done\n"
