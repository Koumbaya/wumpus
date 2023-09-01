# Wumpus
This is a vanilla go implementation of the 1973 classic text-based computer game first developed by Gregory Yob [Hunt The Wumpus](https://en.wikipedia.org/wiki/Hunt_the_Wumpus).

### Gameplay:

1. **The Cave System**: The game is played in a series of interconnected caves arranged in different predefined. Each of the 20 caves is connected to three other caves.

2. **Hazards**:
   - **The Wumpus**: This is the main antagonist of the game. If you enter the Wumpus's cave, you may get eaten and lose the game. Or the Wumpus may get disturbed and move to another cave.
   - **Pits**: There are two pits in the cave system. Falling into a pit results in instant death.
   - **Giant Bats**: There are two caves with giant bats. If you enter a cave with a giant bat, you will be carried off to a random cave, which might be dangerous.

3. **Player Actions**: On each turn, you can choose to:
   - **Move**: To one of the three connected caves.
   - **Shoot**: You can shoot an arrow into one of the adjacent caves in hopes of killing the Wumpus. You start with 5 arrows. The arrow can travel up to 5 rooms. Shooting the Wumpus successfully means you win. Missing the Wumpus usually cause it to wake up and move to a random cave. Be careful of not hitting yourself with the arrow !
   
4. **Hints**: The game provides hints based on which cave the player is in:
   - **"An unmistakable scent hints at a Wumpus lurking nearby!"**: This means the Wumpus is in one of the adjacent caves.
   - **"A chilling draft suggests the presence of a pit nearby!"**: This indicates that one of the connected caves has a pit.
   - **"The distant fluttering of bats reaches your ears!"**: This suggests that giant bats are in an adjacent cave.

5. **Winning and Losing**:
   - **Winning**: The player wins by shooting the Wumpus without falling into a pit or being eaten.
   - **Losing**: The player loses by getting struck by an arrow, falling into a pit, or being eaten by the Wumpus.

6. **Levels Topology**:  
   Levels are often 3D shapes flattened to 2d and thus, can be hard to visualize. See the `levels` folder for 2D renders (but forget the numbers you see on the graph, they are randomized!)
   - **1**: Is a [flattened dodecahedron](https://people.math.sc.edu/Burkardt/data/grf/dodecahedron.png) with 20 caverns.
   - **2**: Is a moebius strip with 20 caverns.
   - **3**: Is a string of beads that loops around. ·:··:··:··:··:·
   - **4**: Is a serie of 5 hexagons connected in a torus shape.
   - **5**: Is a Dendrite-like cave with a branching pattern and a lot of dead-ends.
   - **6**: Is a grid of one-way tunnels. Each cavern is connected to 2 other cavern. Be careful or you might need to loop around !

7. **Advanced**:
   - In advanced mode, you must kill the wumpus then find your way out of the labyrinth by finding a key and its door!

**Notes**: Small difference with the original game :
It would only move 1 cave, which was kind of easier (to relocate) but less forgetfull (it would have a 1/3 chance of eating you unless you backtracked 1 room).  
Possible paths for the arrow weren't displayed, maybe because of computing requirements/space.

Pen & paper are recommended to take notes or draw the map !
While the layout remains fixed, the numbers by which the caverns are named do not follow a particular logic (i.e cavern 1 may not be connected to cavern 2).

### Running the game
```
go run .
```
Type `exit` any time to close the game.  
Type `reset` any time to restart from the start.

There a various flags available :  
`-advanced`: play in advanced mode ! On top of killing the wumpus, you'll need to escape the labyrinth !  
`-nodelay`: run the game without the fake terminal delay.  
`-arrows`: gives infinite arrows.  
`-debug`: cheat/debug mode.  
`-clean`: remove symbols/colors (in case your terminal is broken, or you want that oldschool feeling)

![cover](cover.png)

![screenshot1](screenshot1.png)

![screenshot2](screenshot2.png)

### TODO:
* Add a way to move directly instead of move + number
* Probably a few refactors of the state machine.
* Implement Wumpus 2 (6 different levels!) - in progress
* Checks & reroll of pits location on maps where that could be impossible to solve (`5` and `6` only ?).
* More features in advances (partial maps, clues, traps...)
* bubble-up the few possibles IO errors
* Tests

Original source code on [archive.org](https://archive.org/details/CreativeComputingv01n05SeptemberOctober1975/page/n51/mode/2up) !
![original.png](original.png)
