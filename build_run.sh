USE_LOCAL_ENV=$1
ONLY_BUILD=$2
PLUGINS=("djangoeol" "nodeeol" "pythoneol" "reacteol" "readme" "repovuln" "ecrvuln" "latestpatchdjango" "latestpatchnode" "latestpatchpython")
for plugin in "${PLUGINS[@]}"
do
  go build -buildmode=plugin -o plugins/"$plugin"/"$plugin".so plugins/"$plugin"/"$plugin".go
  echo "Building ${plugin} ✅"

done

if [ "$ONLY_BUILD" == "--only-build" ]; then
  exit 1
fi


if [ "$USE_LOCAL_ENV" == "--local-vars" ]; then
  source test.env
else
  echo "Pass --local-vars to source test.env .."
fi

go run runner.go depchecker.go