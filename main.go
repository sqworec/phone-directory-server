package main

func main() {

	config := Config{
		DSN:  "host=localhost port=5432 user=postgres password=Tenpos2005 dbname=phone-directory",
		Port: 8080,
	}

	app := NewApp(config)
	app.Start(config.Port)
}
