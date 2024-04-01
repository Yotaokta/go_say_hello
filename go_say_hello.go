package go_say_hello

func SayHello(nama string) string {
	return "hello " + nama
}

func ConfirmMessage(nama string) string {
	return SayHello(nama) + " Mau Pesan Apa?"
}
