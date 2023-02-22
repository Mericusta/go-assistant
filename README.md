# go-assistant

## installation

cd cmd/ && go install

## command

### generator

- generate unit test from a function
    goass -cmd=generator -opt=unittest -args="<FILE_REL_PATH>;<FUNC_NAME>"

- generate unit test from a generic function specified type args
    goass -cmd=generator -opt=unittest -args="<FILE_REL_PATH>;<FUNC_NAME>;<TYPE_ARG>,[TYPE_ARG]"