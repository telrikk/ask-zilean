#/bin/bash

rm -rf dist
npm run webpack-production
npm run uglify
rm -rf node_modules
revel package github.com/telrikk/ask-zilean

mkdir ask-zilean
cp Procfile ask-zilean
cp Buildfile ask-zilean
tar -C ask-zilean -zxvf ask-zilean.tar.gz
zip -r ask-zilean.zip ask-zilean
