package weigo

//控制器接口，框架基类控制器继承类该接口
type ControllerInterface interface {
	Init(context *Context)
	Index()
	List()
	View()
	Delete()
	Update()
}
