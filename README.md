# go-assistant

## installation

```shell
cd cmd/ && go install
```

## command

### generator

- generate unittest from a function

```shell
goass -cmd=generate -opt=unittest -file=<FILE_PATH> -func=<FUNC_NAME> -mode=<TYPE_MODE>
```

- generate unittest from a generic function specified type args

```shell
goass -cmd=generator -opt=unittest -file=<FILE_REL_PATH> -func=<FUNC_NAME> -types=<TYPE_ARG>[,TYPE_ARG] -mode=<TYPE_MODE>
```

- generate benchmark from a function

```shell
goass -cmd=generate -opt=benchmark -file=<FILE_PATH> -func=<FUNC_NAME> -mode=<TYPE_MODE>
```

- generate benchmark from a generic function specified type args

```shell
goass -cmd=generator -opt=benchmark -file=<FILE_REL_PATH> -func=<FUNC_NAME> -types=<TYPE_ARG>[,TYPE_ARG] -mode=<TYPE_MODE>
```
