package web


type PublicController struct {
	BaseWebController
}



// @router /index [get]
func (this *PublicController) Index() {
	this.TplName = "web/index.tpl"
}





