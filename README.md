# Connect Four Game 

This is a simple implementation of the classic Connect Four game in Go. Players take turns dropping colored discs into a grid, aiming to connect four of their own discs vertically, horizontally, or diagonally before their opponent does.

## How to Play

1. Clone the repository to your local machine.
2. Navigate to the project directory.
3. Run the game using the Go compiler:

    ```
    go run main.go
    ```

4. Follow the on-screen instructions to input your moves. Players alternate turns, dropping their colored discs into the grid.
5. The game ends when a player successfully connects four of their own discs or when the grid is full with no winner.

## Controls

- Use the numeric keys (1-7) to select the column where you want to drop your disc.
- Press 'q' to quit the game at any time.

## Features

- ANSI escape codes are used to provide a simple terminal-based UI.
- Players' scores are displayed along with their symbols.
- The game board is dynamically updated with each move.
- Winning combinations are detected and the game ends when a player wins.

## Contributing

Contributions are welcome! If you find any bugs or have suggestions for improvements, feel free to open an issue or submit a pull request.

## Credits

This game was developed by qothman as a learning project.

