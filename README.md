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
goass -cmd=generate -opt=unittest -file=<FILE_REL_PATH> -func=<FUNC_NAME> -types=<TYPE_ARG>[,TYPE_ARG] -mode=<TYPE_MODE>
```

- generate benchmark from a function

```shell
goass -cmd=generate -opt=benchmark -file=<FILE_PATH> -func=<FUNC_NAME> -mode=<TYPE_MODE>
```

- generate benchmark from a generic function specified type args

```shell
goass -cmd=generate -opt=benchmark -file=<FILE_REL_PATH> -func=<FUNC_NAME> -types=<TYPE_ARG>[,TYPE_ARG] -mode=<TYPE_MODE>
```

- generate unittest from a struct method

```shell
goass -cmd=generate -opt=unittest -file=<FILE_PATH> -func=<FUNC_NAME> -struct=<STRUCT_NAME> -mode=<TYPE_MODE>
```

- generate unittest from a generic struct method specified type args

```shell
goass -cmd=generate -opt=unittest -file=<FILE_REL_PATH> -func=<FUNC_NAME> -struct=<STRUCT_NAME> -types=<TYPE_ARG>[,TYPE_ARG] -mode=<TYPE_MODE>
```

- generate benchmark from a struct method

```shell
goass -cmd=generate -opt=benchmark -file=<FILE_PATH> -func=<FUNC_NAME> -struct=<STRUCT_NAME> -mode=<TYPE_MODE>
```

- generate benchmark from a generic struct method specified type args

```shell
goass -cmd=generate -opt=benchmark -file=<FILE_REL_PATH> -func=<FUNC_NAME> -struct=<STRUCT_NAME> -types=<TYPE_ARG>[,TYPE_ARG] -mode=<TYPE_MODE>
```

- generate unittest from a interface method

```shell
goass -cmd=generate -opt=unittest -file=<FILE_PATH> -func=<FUNC_NAME> -interface=<INTERFACE_NAME> -mode=<TYPE_MODE>
```

- generate unittest from a generic interface method specified type args

```shell
goass -cmd=generate -opt=unittest -file=<FILE_REL_PATH> -func=<FUNC_NAME> -interface=<INTERFACE_NAME> -types=<TYPE_ARG>[,TYPE_ARG] -mode=<TYPE_MODE>
```

- generate benchmark from a interface method

```shell
goass -cmd=generate -opt=benchmark -file=<FILE_PATH> -func=<FUNC_NAME> -interface=<INTERFACE_NAME> -mode=<TYPE_MODE>
```

- generate benchmark from a generic interface method specified type args

```shell
goass -cmd=generate -opt=benchmark -file=<FILE_REL_PATH> -func=<FUNC_NAME> -interface=<INTERFACE_NAME> -types=<TYPE_ARG>[,TYPE_ARG] -mode=<TYPE_MODE>
```

- generate go AST

```shell
goass -cmd=generate -opt=ast -file=<FILE_REL_PATH> -meta=<META> -ident=<IDENT>[,IDENT]
```

META supports: func, method, struct, interface
note that method whichever from struct or interface needs two ident as struct/interface name and method name
IMPORTANT: not support type constraints