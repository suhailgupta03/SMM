PLUGINS=("djangoeol" "nodeeol" "pythoneol" "reacteol" "readme" "repovuln")
for plugin in "${PLUGINS[@]}"
do
  go build -buildmode=plugin -o plugins/"$plugin"/"$plugin".so plugins/"$plugin"/"$plugin".go
  echo "Building ${plugin} ✅"

done

source test.env
go run runner.go depchecker.go