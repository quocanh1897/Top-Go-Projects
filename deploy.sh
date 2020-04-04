#!/bin/bash

# Your repo
AUTO_REPO="https://thinhdanggroup:$OAuth@github.com/thinhdanggroup/Top-Go-Projects.git"
# Yourname for TravisCI
AUTO_USERNAME="Travis CI"
# Git Email
AUTO_EMAIL="dinhluchvc@gmail.com"
# Branch
AUTO_BRANCH="feature/add-travis-ci"
# Go file
SCRIPT_UPDATE="auto-update.go"


# execute deploy
set -o errexit -o nounset

CUR_TIME=$(date +"%Y-%m-%dT%H")
echo "Current update time: $CUR_TIME"

IS_UPDATE=$(cat README.md  | grep $CUR_TIME | wc -l)

if [ $IS_UPDATE = 1 ]
then 
  echo "README.md is updated for latest=$CUR_TIME"
  exit 0
fi

# push git
git config user.name $AUTO_USERNAME
git config user.email $AUTO_EMAIL
git remote set-url origin $AUTO_REPO
git fetch origin $AUTO_BRANCH
git reset origin/$AUTO_BRANCH
git checkout $AUTO_BRANCH

go run $SCRIPT_UPDATE

git add -A .
git commit -m "Auto Update at $CUR_TIME"
git push -q origin $AUTO_BRANCH