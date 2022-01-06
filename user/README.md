### 生成micro 目录
docker run --rm -v $(pwd):$(pwd) -w $(pwd) micro/micro new category





### 生成proto
- docker run --rm -v $(PWD):$(PWD) -w $(PWD) znly/protoc -I --go_out=./ --micro_out=./ ././*.proto


### 
