package options

func (s *ServerRunOptions) Validate() []error {
	var errors []error

	errors = append(errors, s.GenericOptions.Validate()...)
	errors = append(errors, s.DatabaseOptions.Validate()...)
	errors = append(errors, s.DetectorOptions.Validate()...)
	errors = append(errors, s.LoggerOptions.Validate()...)

	return errors
}
