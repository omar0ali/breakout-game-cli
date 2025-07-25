# Breakout CLI Game

A simple terminal-based implementation of the classic **Breakout** game written in **Go**, using the [`tcell`](https://github.com/gdamore/tcell) library for handling terminal graphics and input.

![game-screenshot](https://github.com/omar0ali/breakout-game-cli/blob/main/screenshots/breakout-game-screenshot.png)

## Features
- [x] Paddle movement using arrow keys
    - Added animation when the paddle is moved, to give it a little bit of smoothness while moving.
- [x] Ball animation and basic collision with walls and paddle
    - Ensure all objects on screen are smooth even when frame rate fluctuates.
    - Added additional brick collision, when the ball hits the sides of the brick, the ball changes
        direction.
- [x] Smooth rendering in the terminal
    - FPS is visible, the game fixed at 30 frames per second.
- [x] Add bricks for the player to break by the ball.
    - Level / Height of bricks can be modified through the configuration file.

- [x] FPS can be configured through the configuration file.

- [x] Add debug information
    - To enable debug mode - can be enabled by adding `debug=true` under `core` in the 
        configuration file.
- [x] Added mouse support, to enable `mouse=true`
- [x] Added multi ball game, hitting the space bar will shoot them, *limited to 10 balls a game*.
- [ ] Add times the ball fell over the paddle.
- [ ] Status Bar to show player details.

## Getting Started

Clone repository

```bash
git clone https://github.com/omar0ali/breakout-game-cli.git
```

Run the game

```bash
go run .
```

### Configuration
The game uses `toml` file for the user to configure the game, paddle, and ball speed can be changed as well as 
the paddle width to fit the player need.

Example

```bash
[core]
duration_ticker = 33 # Target frame/update interval in milliseconds (actual FPS may vary)
debug = true
mouse = true
[player]
speed = 60
jump_by = 8
paddle_wdith = 10
[ball]
speed = 20
[bricks]
level = 3
```

File is saved as `config.toml`

### Controls
- The paddle can be moved left and right using the `Arrow Keys`.
- Mouse control is now supported.
