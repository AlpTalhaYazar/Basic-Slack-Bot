# Basic Slack Bot

This is a simple Slack bot application built in Go, using the Slacker and Slack-Go libraries. It currently supports two commands: an age calculator command and a file upload command.

## Getting Started

To run this bot locally, you'll need to have Go installed on your machine. Also, you will need to create a Slack bot and have its bot token and app token.

### Installation

1. Clone the repository:
    ```
    git clone https://github.com/yourusername/Basic-Slack-Bot.git
    ```

2. Move into the project directory:
    ```
    cd Basic-Slack-Bot
    ```

3. Build and run the bot:
    ```
    go build
    ./Basic-Slack-Bot
    ```

The bot should now be running and connected to your specified Slack workspace.

## Usage

This Slack bot supports the following commands:

- `I born in <year>`: This command calculates the age based on the provided birth year. For example, if you input "I born in 1990", the bot will calculate the age and respond with "You are xx years old".

- `upload <sentence>`: This command allows you to upload a text file containing the sentence you've specified. The file will be titled "Text" and it will be uploaded to the channel where the command was executed.

### Example

Create a new book:

```bash
I born in 1990
```

Upload a text file:

```bash
upload "Hello, World!"
```

## Contributing

Contributions are welcome. Please feel free to open an issue or create a pull request.
