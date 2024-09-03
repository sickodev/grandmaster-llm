# Grandmaster LLM

This project aims to study the capabilities of large language models (LLMs) in playing chess, with a focus on evaluating their prophylactic thinking and strategic skills.

## Project Structure

```
.
├── main.go
├── .env.example
├── cmd/
├── black/
├── white/
└── engine/
    ├── engine.go
    └── manager.go
```

## Description

The Grandmaster LLM Project is designed to pit large language models against each other in chess matches. The primary goal is to assess whether these models are capable of prophylactic thinking and to evaluate their strategic capabilities in a complex game environment.

## Setup

1. Clone the repository:

    ```
    git clone https://github.com/sickodev/grandmaster-llm.git
    cd chess-llm-project
    ```

2. Copy the `.env.example` file to `.env` and fill in the necessary environment variables:

    ```
    cp .env.example .env
    ```

3. Install dependencies:
    ```
    go mod tidy
    ```

## Usage

To run the project:

```
go run main.go default
# This command loads the simulation in it's default settings.
```

## Project Components

-   `main.go`: The entry point of the application.
-   `cmd/`: Contains command-line interface related code.
-   `black/`: Handles the black player's moves and strategy.
-   `white/`: Handles the white player's moves and strategy.
-   `engine/`:
    -   `engine.go`: Implements the chess engine logic.
    -   `manager.go`: Manages the overall flow of the chess games.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Contact

[sickodev] - [kalyanbishwa03@gmail.com]

Project Link: [https://github.com/sickodev/grandmaster-llm](https://github.com/sickodev/grandmaster-llm)
