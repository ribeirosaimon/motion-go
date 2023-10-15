package emailSender

import "fmt"

func emailValue(code string) string {
	return fmt.Sprintf(`
    <!DOCTYPE html>
    <html>
    <head>
    <title>Binary Code: 101010</title>
    <style>
    body {
    font-family: Arial, sans-serif;
    background-color: #f0f0f0;
    text-align: center;
    margin: 0;
    padding: 0;
    }
    .container {
    background-color: #ffffff;
    max-width: 400px;
    margin: 100px auto;
    padding: 20px;
    border-radius: 5px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    }
    h1 {
    color: #333;
    }
    p {
    font-size: 24px;
    color: #333;
    }
    </style>
    </head>
    <body>
    <div class="container">
    <h1>Binary Code</h1>
    <p>%s</p>
    </div>
    </body>
    </html>`, code)
}
