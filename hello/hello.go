package hello



const(
	helloPrefix = "Hello "
	spanishHelloPrefix = "Hola "
	spanish = "Spanish"
	french = "French"
	frenchHelloPrefix = "Bonjour "


)

func Hello(name,language string)string{
	if name==""{
		return helloPrefix+"World"
	}
	prefix := greeting(language)
	
	return prefix+name
}

func greeting(language string)(prefix string){
	switch language{
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix=helloPrefix
	}
	return
}