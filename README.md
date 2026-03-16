# Sodawater
Simple engine for rendering animations with Bubbletea

## To-Do
1. Rendering System
    - Surface cell style optimization
    - Camera system
    - ASCII sprites & procedural
2. Entity System
    - Simple Object Physics
3. Window Sizing Option (only fullscreen currently)

### Window Sizing
Should provide two options:
1. full screen by default
2. set option for width/height on frame instantiation

> [!NOTE]
> subsequent `WindowSizeMsg` are ignored after initialization

**WindowSizeMsg** &mdash; can occur after initial View() is called... 
set state in soda model [see here](https://github.com/charmbracelet/bubbletea/discussions/283)

