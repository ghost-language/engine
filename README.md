# Engine
Engine is a simple framework to make 2D games in Ghost.

## Status
Currently messing around with SDL and experimenting on what the API should look like. Not in a usable state (Ghost itself isn't even complete for the scripting aspect to be built).

## Notes

### Built In Functions

#### `load()`
Called only once before the game loop begins. This is where you would want to pre-load any resources, initialize variables, and configure settings. While its possible to perform these actions elsewhere in your code, its generally a better idea to do so here as its not called at every frame.

#### `update()`
Called continuously where calculations and other deterministic factors should be performed. Will pass through the `delta time` allowing you to factor in the time since the last time this function was called. `delta time` is usually a very small number, like `0.025714`.

#### `draw()`
Much like `update`, this function is called continuously, allowing you to draw and update the screen.