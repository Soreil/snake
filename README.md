snake
=====

the area to draw on consists of a grid of x * y cells.
inside of this grid is another grid of x * y cells where the actual game is played. The padding around this grid is used to display statistics.
The game runs in a for  select  loop with two cases, in case of a pressed key the key is evaluated. in any other case the default case executes with all of the turn logic. The game logic waits for a certain amount of time before ending to control the framerate. in case key pressed is ESC the game ends. 

>#### The grid has a number of properties:
- Cell array
- Width Height int
	
The slice is a two dimensional slice of Width * Height
Cell[0][0] is in the top left corner.

>#### a cell has a number of properties:
- enum:
  - SNAKE
  - EMPTY
  - ITEM
- Ttl int

in case of SNAKE there is also a TTL, this is decremented by one every game tick. When the TTL hits 0 the square becomes EMPTY. A cell can be only one of these three basic types, When an item becomes snake the snake length is increased by one and the TTL is set. when a cell becomes empty it is eligible to become ITEM.

>#### DIRECTION in a enum of the following values:
- LEFT
- DOWN
- UP
- RIGHT

A direction is changed by pressing the corresponding arrow key. By default the direction is set to RIGHT. The DIRECTION value is checked every game tick, in case no arrow keys were hit the old direction is used. After this the snake is moved.

>#### SNAKE contains the following properties:
- X Y Length int

The snake starts from the center and moves into DIRECTION by one cell every turn. The cell it lands on becomes SNAKE. If the cell it lands on is ITEM every SNAKE square is incremented by one and the SNAKE LENGTH is incremented by one. If the snake lands on a SNAKE cell the game ends. If the snake lands on an out of bounds(in the padding) cell the game ends.

ITEM is spawned in a random cell not yet occupied by SNAKE. Whenever ITEM is changed into SNAKE a new ITEM is spawned. If there are no areas not yet SNAKE the game ends. Seeing as how the ITEM is only spawned when another item is taken out which increments SCORE SCORE is used as the seed.

PADDING has two elements, a score and a coordinate for snake. The score is printed at the top left of the padding and the coordinates are printed in the top right.

SCORE is incremented by the length of the snake every time an ITEM is changed into SNAKE.
