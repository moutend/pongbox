# PongBox

Shell command launcher inspired by MIDI drum pads.

## Install

```console
go install github.com/moutend/pongbox/cmd/pongbox@latest
```

## Configuration

Before starting, create a configuration file named `.pongbox` at your home directory.

The configuration file looks like this:

```json
{
  "commands": {
    "key_<key name>": {
      "name": "<command name>",
      "args": [
        "<arg1>",
        "<arg2>",
        "<argN>"
      ]
    }
  }
}
```

### Example

NOTE: This configuration targets macOS.

```json
{
  "commands": {
    "key_1": {
      "name": "say",
      "args": [
        "-r",
        "100",
        "One"
      ]
    },
    "key_2": {
      "name": "say",
      "args": [
        "-r",
        "100",
        "Two"
      ]
    },
    "key_3": {
      "name": "say",
      "args": [
        "-r",
        "100",
        "Three"
      ]
    },
    "key_arrow_left": {
      "name": "afplay",
      "args": [
        "/System/Library/Sounds/Glass.aiff"
      ]
    },
    "key_arrow_right": {
      "name": "afplay",
      "args": [
        "/System/Library/Sounds/Tink.aiff"
      ]
    }
  }
}
```

## Usage

To launch the pongbox, just type:

```console
pongbox
```

To quit pongbox, press ESC key or Ctrl-C.

## LICENSE

MIT
