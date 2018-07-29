package log

// Hlog log libaray
type Hlog interface {
	Debug(f interface{},v ...interface{})
	Info(f interface{},v ...interface{})
	Warn(f interface{},v ...interface{})
	Error(f interface{},v ...interface{})
}

var globalLog Hlog

func init(){
	globalLog,_=NewBeelog("",true,true)
}

// SetHlogger reg a logger in global
func SetHlogger(logger Hlog){
	globalLog=logger
}

// Info log info
func Info(f interface{},v ...interface{}){
	globalLog.Info(f,v)
}

// Error log error
func Error(f interface{},v ...interface{}){
	globalLog.Error(f,v)
}

// Debug log debug
func Debug(f interface{},v ...interface{}){
	globalLog.Debug(f,v)
}