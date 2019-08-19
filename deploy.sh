#!/bin/sh
git pull
go run auto-update.go

rm README.md
mv README2.md README.md

git add .
git commit -m "Auto update" -a
git push origin