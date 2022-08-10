package bootstrap

var Target = "http://localhost/{{path}}"

func Init(target string) {
	Target = target
}
