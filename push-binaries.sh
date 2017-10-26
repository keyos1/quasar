#!/bin/sh
# Push binaries to quasar-binaries repo
make binaries
cd binaries && git remote add binaries https://$GH_TOKEN@github.com/keyos1/quasar-binaries
git push binaries develop
if [ $? -ne 0 ]; then
  cd .. && rm -rf binaries
  exit 1
fi
