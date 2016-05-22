#!/bin/bash

echo "mode: set" > acc.out
fail=0

# get coverage for unit tests
for dir in $(find . -maxdepth 10 -not -path './tests/functional*' -not -path './Godeps**' -not -path './.git*' -not -path '*/_*' -type d);
do
  if ls $dir/*.go &> /dev/null; then
    go test -coverprofile=profile.out $dir || fail=1
    if [ -f profile.out ]
    then
      cat profile.out | grep -v "mode: set" >> unit.out
      rm profile.out
    fi
  fi
done

# get coverage for functional tests
go test -coverprofile=functional.out ./tests/functional || fail=1

# merge
cat functional.out unit.out | grep -v "mode: set" >> acc.out

# Failures have incomplete results, so don't send
if [ -n "$COVERALLS" ] && [ "$fail" -eq 0 ]
then
  goveralls -v -coverprofile=acc.out $COVERALLS
fi

rm -f acc.out
rm -f unit.out
rm -f functional.out

exit $fail
