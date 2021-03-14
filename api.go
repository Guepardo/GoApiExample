package main

func main() {
	app := App{}

	app.Initialize()
	app.Start(":8080")
}
