#/bin/bash

export NODE_ENV=production
mkdir -p /tmp/deploy/src/github.com/telrikk/ask-zilean/
cd ${TRAVIS_BUILD_DIR}
rm -rf dist
npm run webpack-production
npm run uglify
rm -rf node_modules
cp -R ./ /tmp/deploy/src/github.com/telrikk/ask-zilean/
cd /tmp/deploy/
cp /tmp/deploy/src/github.com/telrikk/ask-zilean/Procfile ./
cp /tmp/deploy/src/github.com/telrikk/ask-zilean/Buildfile ./
zip -q ${TRAVIS_BUILD_DIR}/deploy.zip -x'src/github.com/telrikk/ask-zilean/node_modules/*' -r ./
cd ${TRAVIS_BUILD_DIR}
