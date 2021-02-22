package ishell

type ISheller interface {
	// Start (with some possible) arguments, must be able
	// to start the terminal with the arg parser.
	//
	// Start, brings up a terminal using our tp provider,
	// and the screen is in
	Start(func() Context) error
}
