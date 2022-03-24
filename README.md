## lem-in

### Objectives

This project is meant to make you code a digital version of an ant farm.

Create a program `lem-in` that will read from a file (describing the ants and the colony) given in the arguments.

Upon successfully finding the quickest path, `lem-in` will display the content of the file passed as argument and each move the ants make from room to room.

How does it work?

- You make an ant farm with tunnels and rooms.
- You place the ants on one side and look at how they find the exit.

You need to find the quickest way to get `n` ants across a colony (composed of rooms and tunnels).

- At the beginning of the game, all the ants are in the room `##start`. The goal is to bring them to the room `##end` with as few moves as possible.
- The shortest path is not necessarily the simplest.
- Some colonies will have many rooms and many links, but no path between `##start` and `##end`.
- Some will have rooms that link to themselves, sending your path-search spinning in circles. Some will have too many/too few ants, no `##start` or `##end`, duplicated rooms, links to unknown rooms, rooms with invalid coordinates and a variety of other invalid or poorly-formatted input. In those cases the program will return an error message `ERROR: invalid data format`. If you wish, you can elaborate a more specific error message (example: `ERROR: invalid data format, invalid number of Ants` or `ERROR: invalid data format, no start room found`).

You must display your results on the standard output in the following format :

```console
number_of_ants
the_rooms
the_links

Lx-y Lz-w Lr-o ...
```

- x, z, r represents the ants numbers (going from 1 to number_of_ants) and y, w, o represents the rooms names.

- A room is defined by `"name coord_x coord_y"`, and will usually look like `"Room 1 2", "nameoftheroom 1 6", "4 6 7"`.

- The links are defined by `"name1-name2"` and will usually look like `"1-2", "2-5"`.

Here is an example of this in practice :

```
##start
1 23 3
2 16 7
#comment
3 16 3
4 16 5
5 9 3
6 1 5
7 4 8
##end
0 9 5
0-4
0-6
1-3
4-3
5-2
3-5
#another comment
4-2
2-1
7-6
7-2
7-4
6-5
```

Which corresponds to the following representation :

```console
        _________________
       /                 \
  ____[5]----[3]--[1]     |
 /            |    /      |
[6]---[0]----[4]  /       |
 \   ________/|  /        |
  \ /        [2]/________/
  [7]_________/
```

### Usage

Example 1 :

```
$ go run . example00.txt
3
##start
1 23 3
2 16 7
3 16 3
4 16 5
5 9 3
6 1 5
7 4 8
##end
0 
