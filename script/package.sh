# compile for version
env CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/fileserver ./main.go

if [ $? -ne 0 ]; then
    echo "compile for version error"
    exit 1
fi

fileserver_version=$(./bin/fileserver --version)
echo "build version: $fileserver_version"

# cross_compiles
make -f ./script/Makefile.cross-compiles

rm -rf ./release/packages
mkdir -p ./release/packages

os_all='linux windows darwin freebsd'
arch_all='386 amd64 arm arm64 mips64 mips64le mips mipsle riscv64'

cd ./release

for os in $os_all; do
    for arch in $arch_all; do
        fileserver_dir_name="fileserver_${fileserver_version}_${os}_${arch}"
        fileserver_path="./packages/fileserver_${fileserver_version}_${os}_${arch}"

        if [ "x${os}" = x"windows" ]; then
            if [ ! -f "./fileserver_${os}_${arch}.exe" ]; then
                continue
            fi
            mkdir ${fileserver_path}
            mv ./fileserver_${os}_${arch}.exe ${fileserver_path}/fileserver.exe
        else
            if [ ! -f "./fileserver_${os}_${arch}" ]; then
                continue
            fi
            if [ ! -f "./fileserver_${os}_${arch}" ]; then
                continue
            fi
            mkdir ${fileserver_path}
            mv ./fileserver_${os}_${arch} ${fileserver_path}/fileserver
        fi
        cp ../LICENSE ${fileserver_path}

        # packages
        cd ./packages
        if [ "x${os}" = x"windows" ]; then
            zip -rq ${fileserver_dir_name}.zip ${fileserver_dir_name}
        else
            tar -zcf ${fileserver_dir_name}.tar.gz ${fileserver_dir_name}
        fi
        cd ..
        rm -rf ${fileserver_path}
    done
done

cd -
