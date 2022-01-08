package commonds

type Base struct {
	signature   string
	description string
	name        string
	help        string
}

func (this *Base) Signature() string {
	return this.signature
}
func (this *Base) Description() string {
	return this.description
}
func (this *Base) Name() string {
	return this.name
}
func (this *Base) Help() string {
	return this.help
}
