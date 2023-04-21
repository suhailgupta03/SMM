USE_LOCAL_ENV=$1
ONLY_BUILD=$2

make

if [ "$ONLY_BUILD" == "--only-build" ]; then
  exit 0
fi


if [ "$USE_LOCAL_ENV" == "--local-vars" ]; then
  source test.env
else
  echo "Pass --local-vars to source test.env .."
fi

go run runner.go depchecker.go