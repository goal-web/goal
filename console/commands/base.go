package commands

type Base struct {
	Signature   string
	Description string
	Name        string
	Help        string
	args        []Arg
}

func BaseCommand(signature, description string) Base {
	name, args := ParseSignature(signature)
	return Base{
		Signature:   signature,
		Description: description,
		Name:        name,
		Help:        args.Help(),
		args:        args,
	}
}

func (this *Base) GetSignature() string {
	return this.Signature
}
func (this *Base) GetDescription() string {
	return this.Description
}
func (this *Base) GetName() string {
	return this.Name
}
func (this *Base) GetHelp() string {
	return this.Help
}
