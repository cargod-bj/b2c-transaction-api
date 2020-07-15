md github.com\cargod-bj\b2c-proto-common\common
curl -o github.com/cargod-bj/b2c-proto-common/common/common.proto https://raw.githubusercontent.com/cargod-bj/b2c-proto-common/master/common/common.proto
for /r %%i in (*.proto) do (
	protoc -I%~dp0 --go_out=paths=source_relative:.  --micro_out=paths=source_relative:. %%i
)
rd /s /q github.com