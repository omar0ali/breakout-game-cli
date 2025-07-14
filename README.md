# Breakout CLI Game

A simple terminal-based implementation of the classic **Breakout** game written in **Go**, using the [`tcell`](https://github.com/gdamore/tcell) library for handling terminal graphics and input.

## Features
- [x] Paddle movement using arrow keys
    - Added animation when the paddle is moved, to give it a little bit of smoothness while moving.
- [x] Ball animation and basic collision with walls and paddle
    - Ensure all objects on screen are smooth even when frame rate fluctuates.
- [x] Smooth rendering in the terminal
    - FPS is visible, the game fixed at 30 frames per second.
- [ ] Add bricks for the player to break by the ball.

## Getting Started

Clone repository

```bash
git clone https://github.com/omar0ali/breakout-game-cli.git
```
```
```

Run the game

```bash
go run main.go
```
