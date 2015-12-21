package carbon

type Logger interface {
	SetPrefix(string)

	Debug(msg interface{})
	Info(msg interface{})
	Warning(msg interface{})
	Error(msg interface{})
	Fatal(msg interface{})

	Debugf(f string, msg interface{})
	Infof(f string, msg interface{})
	Warningf(f string, msg interface{})
	Errorf(f string, msg interface{})
	Fatalf(f string, msg interface{})
}
