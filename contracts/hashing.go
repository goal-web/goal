package contracts

type HasherProvider func(config Fields) Hasher

type HasherFactory interface {
	Hasher
	Driver(driver string) Hasher
	Extend(driver string, hasherProvider HasherProvider)
}

type Hasher interface {
	Info(hashedValue string) Fields
	Make(value string, options Fields) string
	Check(value, hashedValue string, options Fields) bool
}
