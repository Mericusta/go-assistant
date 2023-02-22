package generate

type generateBenchmarkArgs struct {
}

func (a *generateBenchmarkArgs) SliceGenerateBenchmark(arg1, arg2 string) (*generateBenchmarkArgs, []*generateBenchmarkArgs) {
	return nil, nil
}

func SliceGenerateBenchmark(arg1, arg2 string) (*generateBenchmarkArgs, []*generateBenchmarkArgs) {
	return nil, nil
}

func (a *generateBenchmarkArgs) NonReturnGenerateBenchmark(arg1, arg2 string) {

}

func NonReturnGenerateBenchmark(arg1, arg2 string) {

}

func TemplateFunc[T any](arg1, arg2 string, arg3 *T) *T {
	return nil
}

func (a *generateBenchmarkArgs) NonArgsGenerateBenchmark() {

}

func NonArgsGenerateBenchmark() {

}
