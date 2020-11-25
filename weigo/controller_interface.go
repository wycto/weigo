package weigo

type ControllerInterface interface {
	Init(context *Context)
	Index()
	List()
	View()
	Delete()
	Update()
}
