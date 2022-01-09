package commonds

type Base struct {
	signature   string
	description string
	name        string
	help        string
}

func (this *Base) GetSignature() string {
	return this.signature
}
func (this *Base) GetDescription() string {
	return this.description
}
func (this *Base) GetName() string {
	return this.name
}
func (this *Base) GetHelp() string {
	return this.help
}
