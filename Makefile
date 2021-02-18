OSList = darwin linux #windows
# 目标平台的操作系统（darwin、freebsd、linux、windows）
ARCHList = amd64 arm64
# 目标平台的体系架构（386、amd64、arm）
default:
	bash autogen.sh
	@echo "Build executables..."
	@for os in $(OSList) ; do            \
		for arch in $(ARCHList) ; do     \
			GOOS=$$os GOARCH=$$arch go build -o builds/keepmego-$$os-$$arch .;    \
		done                              \
	done


clean:
	@rm -rf builds
	@rm -f test.db

test:
	@go test -v -cover
	@rm -f test.db

