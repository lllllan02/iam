#！/bin/bash

# 遍历 ./internal/$1 目录下的 go 文件，在对应的 test 目录下生成 _mock.go 文件
mock() {
    for file in `ls ./internal/$1`
        do
            if [ ! -d $1/$file ] && [ ${file##*.} == "go" ]; then
                `mockgen -source ./internal/$1/$file -destination ./test/mock/$1/${file%.*}_mock.go -package $1`
            fi
        done
}

mock data
mock service