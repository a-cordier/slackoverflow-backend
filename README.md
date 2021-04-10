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

## The slack app

To interact with the hook, the app must be granted the following scoped:
  - commands
  - emoji:read

Once you have configured scopes, enable interactivity in the `Interactivity & Shortcuts` menu
and fill the request URL form field as follow:

```
https://${SLACKOVERFLOW_HOST}/hook
```

Then add expected shortcuts on your messages:

| Name         | Location | Callback ID  |
|--------------|----------|--------------|
| Add Question | Messages | add_question |
| Add Answer   | Messages | add_answer   |

> Name can be overridden to a name your choice, but location and callback ID *must* be defined
as given.
