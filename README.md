# Wumpus
This is a vanilla go implementation of the 1973 classic text-based computer first developed by Gregory Yob [Hunt The Wumpus](https://en.wikipedia.org/wiki/Hunt_the_Wumpus).

### Gameplay:

1. **The Cave System**: The game is played in a series of interconnected caves arranged in a dodecahedron. Each of the 20 caves is connected to three other caves.

2. **Hazards**:
   - **The Wumpus**: This is the main antagonist of the game. If you enter the Wumpus's cave, you may get eaten and lose the game. Or the Wumpus may get disturbed and move to another cave.
   - **Pits**: There are two pits in the cave system. Falling into a pit results in instant death.
   - **Giant Bats**: There are two caves with giant bats. If you enter a cave with a giant bat, you will be carried off to a random cave, which might be dangerous.

3. **Player Actions**: On each turn, you can choose to:
   - **Move**: To one of the three connected caves.
   - **Shoot**: You can shoot an arrow into one of the adjacent caves in hopes of killing the Wumpus. The arrow can travel up to 5 rooms. Shooting the Wumpus successfully means you win. Missing the Wumpus might cause it to wake up and move to a random adjacent cave. Be careful of not hitting yourself with the arrow !
   
4. **Hints**: The game provides hints based on which cave the player is in:
   - **"You smell a Wumpus!"**: This means the Wumpus is in one of the adjacent caves.
   - **"You feel a strong draft!"**: This indicates that one of the connected caves has a pit.
   - **"You hear bats nearby!"**: This suggests that giant bats are in an adjacent cave.

5. **Winning and Losing**:
   - **Winning**: The player wins by shooting the Wumpus without falling into a pit or being eaten.
   - **Losing**: The player loses by getting struck by an arrow, falling into a pit, or being eaten by the Wumpus.

Pen & paper are recommended to take notes or draw the map ! (or you can print a [flattened dodecahedron](https://people.math.sc.edu/Burkardt/data/grf/dodecahedron.png))

### Running the game
```
go run .
```
Type `exit` any time to close the game.

![cover](cover.png)

### TODO:
* A prettier print would be great, maybe with a little bit of latency so as to feel more 1970s-ish.
* Probably a few refactors of the state machine.
* Provide a map with correctly numbered nodes.
* Unexport labyrinth.Rooms
* Maybe handle better the +1 offset ?
* Tests