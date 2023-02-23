# go-assistant

## installation

```shell
cd cmd/ && go install
```

## command

### generator

- generate unit test from a function

```shell
goass -cmd=generate -opt=unittest -file=<FILE_PATH> -func=<FUNC_NAME>
```

- generate unit test from a generic function specified type args

```shell
goass -cmd=generator -opt=unittest -file=<FILE_REL_PATH> -func=<FUNC_NAME> -types=<TYPE_ARG>[,TYPE_ARG]
```
