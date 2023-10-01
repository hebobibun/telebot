# TELEBOT

Welcome to Telebot, a simple example Telegram bot built in Go and using [telegram-bot-api](github.com/go-telegram-bot-api/telegram-bot-api")

## Description

This telebot is a basic Telegram bot that responds to commands and messages. It's a starting point for building more complex Telegram bots with additional features.

## Features

- `/start`: Start the bot and receive a welcome message.
- `/hello`: Greet the bot and receive a friendly response.s

## Getting Started

### Prerequisites

- Go installed on your local machine (https://golang.org/doc/install)
- Telegram Bot API Token (Obtain one by talking to the BotFather on Telegram: https://core.telegram.org/bots#botfather)

### Installation

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/hebobibun/telebot.git
2. Navigate to the project directory:
    ```bash
    cd telebot
3. Create a .env file in the project root and add your Telegram Bot API Token (or pass when app is executed):
    ```bash
    TELEGRAM_BOT_API_TOKEN=your_api_token_here
4. Build and run the bot:
    ```bash
    go run main.go

### Usage
Start the bot by sending the /start command.
Greet the bot with /hello to receive a friendly response.

### Contributing
Contributions are welcome! If you have ideas for improvements or new features, please open an issue or submit a pull request.

#### License
This project is licensed under the MIT License - see the LICENSE file for details.

### Acknowledgments
Thanks to the Telegram team for providing the Telegram Bot API.
[Telegram Bot API Documentation.](https://core.telegram.org/bots/api)

### Contact
Author: [hebobibun](https://github.com/hebobibun)

Email: [hebobibun@algohaven.com](mailto:hebobibun@algohaven.com)

Feel free to reach out if you have any questions or feedback!
