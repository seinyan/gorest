package logging

type Logger interface {
	Info(args ...interface{})
	Warning(args ...interface{})
	Fatal(args ...interface{})
	Error(args ...interface{})
}

type logger struct {

}


func New() Logger {
	return &logger{}
}

//func (l *logger) Info(args ...interface{})  {
//	l.Info(args)
//}

//func (l *Logger) Warning(args ...interface{})  {
//	l.Warning(args)
//}
//
//
//func (l *Logger) Fatal(args ...interface{})  {
//	l.Fatal(args)
//}
//
//func (l *Logger) Error(args ...interface{})  {
//	l.Error(args)
//}
//
