# Engine Documentation

## Modules
Engine includes a variety of modules to easily use and interact with the various graphic, input, and audio APIs.

---

## Graphics
#### `Graphics.clear()`
#### `Graphics.draw(image: string, x: number, y: number)`
#### `Graphics.filledRectangle(x: number, y: number, width: number, height: number)`
#### `Graphics.line(x1: number, y1: number, x2: number, y2: number)`
#### `Graphics.pixel(x: number, y: number)`
#### `Graphics.rectangle(x: number, y: number, width: number, height: number)`
#### `Graphics.setColor(red: number, green: number, blue: number)`

---

## Keyboard
#### `Keyboard.isDown(key: string)`

---

## Mouse
#### `Mouse.hideCursor()`
#### `Mouse.showCursor()`

---

## Window
#### `Window.bordered()`
#### `Window.borderless()`
#### `Window.fullscreen()`
#### `Window.height(pixels: number)`
#### `Window.setSize(width: number, height: number)`
#### `Window.title(title: string)`
#### `Window.width(pixels: number)`