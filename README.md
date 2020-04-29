# Reconstruct binary search trees

I subscribe to a "daily coding problem" email list.
Every day, they send out one interview question from a technical interview of
a desirable company to work for.
The sent me this:

    Given the sequence of keys visited by a postorder traversal of a binary
    search tree, reconstruct the tree.
    
    For example, given the sequence 2, 4, 3, 8, 7, 5, you should construct
    the following tree:
    
        5
       / \
      3   7
     / \   \
    2   4   8

I would have done poorly in an interview where they
expected success on this. It took me 2 days and 7 pieces of paper
covered with binary search tree graphs to hit on a solution.

Then it took me several hours of experimenting to get to working code.
I even had other binary search tree programs saved, so I didn't have
to start from a zero-length file.

I found several solutions on the web:

* [Stack overflow](https://stackoverflow.com/questions/13167536/how-to-construct-bst-given-post-order-traversal#13168162)
* [Geeks for geeks](https://www.geeksforgeeks.org/construct-a-binary-search-tree-from-given-postorder/)
gives an **O**(n<sup>2</sup>) and an **O**(n) solution.

My solution differed radically from those.

## Design

The input to the desired algorithm is an array or list of
integers, ordered as a post-order traverse of an existing tree.
The final node of the list ends up as the root of the finished
binary search tree.
I observed that if you build a binary search tree from
that list one element at a time,
every value gets inserted as the new root of a tree.
Because the original tree has the binary search tree property,
and the keys are integers,
you can decide where to place the node that was the tree's root.
Without that property, you can't recover a binary tree from
just its post-order traversal.

With this observation, it becomes possible to write a function
with a binary search tree and a value to insert
that returns a binary search tree.
The returned binary search tree has a post-order traversal
identical to the order of the values inserted so far.

I wrote just such [a program](postorder.go) to try things out.
It accepts a list of integers on the command line.
The first stage of the program does a traditional binary search tree insert
on each of the integers from the command line, in order.
I believe you can produce a binary search tree of any desired configuration
by listing tree node values breadth-first.
Next, the program performs a post-order traversal of that binary search tree,
giving back a slice (I wrote these programs in Go) of ints.

The slice of ints from the post-order traverse gets used to create
a second binary search tree by inserting the ints in array order.

The program prints out a [GraphViz](http://graphviz.org/)
language representation of
both trees, the tree created by ordinary insertion, and the tree
created by post-order insertion.

The tree in the problem statement could be created by a command
line like this:

    ./postorder 5 3 7 2 4 8 > example.dot

You get the order of the numbers on the command line 
by doing a breadth-first traversal of the desired binary search tree.
After using [a script](comp) to run the command line,
and process using GraphViz `dot`, you get an image like this:

![tree comparison](https://github.com/bediger4000/postorder-tree-traversal/raw/master/example.png)

Sure, I could be faking that image.
I am including the source of the program that created it,
so you can look at the source,
and try it yourself.

I believe the example is carefully crafted to avoid
at least one case.
My algorithm has to have a for-loop to traverse all-left or all-right
branches to find where to insert a root node.
Trees exist where the insertion point is some arbitrary number of links
down the left or right branches from the current root node.
You can create an example tree with breadth-first order of insertion
`5 3 7 6`.
A binary search tree created like that has post-order traversal
array of `3 6 7 5`.  Inserting the final value of 5 causes a traverse
down the left branches of the existing (linear) tree using my algorithm.
I believe you can construct binary search trees where that traverse can
have an arbitrary (yet finite) value.

### Build it

    $ go build postorder.go
    $ ./comp 5 9 2 0 3 8 7 6 1
 
The `comp` script invokes [feh](https://feh.finalrewind.org/) to immediately display the
two binary search trees it creates.

## Verification - fuzzing

I wrote a [version of the program](randtree.org) that accepts a single number on its command line,
generates that many pseudo-random integers, and inserts them into a binary search tree.
A post-order traversal of that tree returns an array of integers.
The program inserts that array's values in a second tree.

More-or-less random trees help flush out any bugs in the algorithm.
I generated the image below like this: 

    $ go build randtree.go
    $ ./randt 11

![tree comparison](https://github.com/bediger4000/postorder-tree-traversal/raw/master/example_random.png)

I use `rand.Intn()` to generate pseudo-random values for these trees.
`rand.Intn()` generates uniform distribution numbers,
so the resulting trees tend to be "filled in",
rather than have long scraggly side branches.

## Performance

The web pages I found that gave solutions to the problem above
had complexity estimates for their algorithms ( **O**(N) and **O**(N<sup>2</sup>) )
My algorithm is quite different,
and it's not that much code.
I decided to count the number of comparisons for trees of different sizes to get an empirical estimete
of time complexity.
Also, I couldn't figure out how to estimate the time complexity of this algorithm.
I thought it was possibly sub-linear, **O**(lg N)

I write [another version](compcnt.go) that actually counts the number of comparisons
of node's key and insert value, and gives back the total number of comparisons made
to recreate a binary search tree from the post-order traverse values.

    $ go build compcnt.go
    $ ./runcomp
    $ ./process
    $ gnuplot < comp.load

![complexity](https://github.com/bediger4000/postorder-tree-traversal/raw/master/complexity.png)

Wow, it's linear! And this particular implementation has a factor of almost exactly 2.5!
If I counted the tests for node against nil, I would get something like a factor of 7,
but it would still be linear with the number of values inserted.
When I see this, and think about it, my early **O**(lg N) guess is completely wrong.
Insertion of a value into a tree means that the value, and some other nodes' values,
get inspected.
Linear is the best you can do in this case.
