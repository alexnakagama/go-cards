# Go Course - Card Deck

A simple Go project that simulates a deck of playing cards. Built as part of learning fundamentals of Go.

## Overview

This project implements a custom `Deck` type (a slice of strings) with various operations: creating a deck, shuffling, dealing hands, printing, and persisting to disk.

## Project Structure

```
.
├── main.go         # Entry point
├── deck.go         # Deck type and its methods
├── deck_test.go    # Unit tests
├── go.mod          # Go module definition
└── my_cards        # Saved deck file
```

## Features

- **New Deck** — generates a 48-card deck (4 suits × 12 values)
- **Print** — prints all cards with their index
- **Deal** — splits the deck into a hand and the remaining cards
- **Shuffle** — randomizes card order in place
- **Save to File** — persists the deck as a comma-separated string
- **Load from File** — restores a deck from a saved file

## Suits & Values

| Suits    | Values                                                   |
|----------|----------------------------------------------------------|
| Spades   | Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Eleven, Twelve |
| Diamonds |                                                          |
| Hearts   |                                                          |
| Club     |                                                          |

## Getting Started

### Prerequisites

- Go 1.25+

### Run

```bash
go run main.go
```

### Test

```bash
go test ./...
```

## How It Works

`main.go` loads a deck from the `my_cards` file, shuffles it, and prints the result.

```go
cards := newDeckFromFile("my_cards")
cards.shuffle()
cards.print()
```
