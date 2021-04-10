# Slackoverflow Backend

A Q&A backend that acts as a hook for slack messages short cut integration.

Question and Answers in a slack thread will be saved for further reuse, with slack
emoji acting as tags (on both Questions and Answers).

## Usage

```shell
go build
slackoverflow -h # prints help and available commands
slackoverflow serve -p 8080
```