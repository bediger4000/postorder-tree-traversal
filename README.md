# Reconstruct binary search trees

I subscribe to a "daily coding problem" email list.
They send out one interview question from a technical interview of
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

I'm certain I would have done poorly in an interview where they
expected success on this. It took me 2 days and 7 pieces of paper
covered with binary search tree graphs to hit on a solution.

Then it took me several hours of experimenting to get to working code.
I even had other binary search tree programs saved, so I didn't have
to start from a zerol-length file.

I found several solutions on the web:

* [Stack overflow](https://stackoverflow.com/questions/13167536/how-to-construct-bst-given-post-order-traversal#13168162)
* [Geeks for geeks](https://www.geeksforgeeks.org/construct-a-binary-search-tree-from-given-postorder/) gives an O(n^2) and an O(n) solution.

## Design

My design is different.
I observe that if you build a binary search tree from a list of integers
ordered as if they were derived from a post-order traverse of another
binary search tree,
not only is the last node the root of the finished tree,
but that every node inserted into the tree becomes the root of the
tree when it is inserted.

With this observation, it becomes possible to write a function
with a binary search tree and a value to insert
that returns a binary search tree.
The returned binary search tree has a pos-order traversal
identical to the values inserted so far.

I wrote just such [a program](postorder.go) to try things out.
It accepts a list of integers on the command line.
The first stage of the program does a traditional binary search tree insert
on each of the integers from the command line, in order.
I believe you can produce a binary search tree of any desired configuration
by an appropriate command line list of integers.
Next, the program performas a post-order traversal of that binary search tree,
giving back a slice (program is in Go) of ints.

The slice of ints from the post-order traverse gets used to create
a second binary search tree by inserting the ints in array order.

The program prints out a [GraphViz](http://graphviz.org/)
language representation of
both trees, the tree created by ordinary insertion, and the tree
created by post-order insertion.

The tree in the problem statement could be created by a command
line like this:

   ./postorder 5 3 7 2 4 8 > example.dot

After using [a script](comp) to run the command line,
and process using GraphViz `dot`, you get an image like this:

![tree comparison](https://raw.githubusercontent.com/bediger4000/postorder-traversal/master/example.png)

Sure, I could be faking that image.
But I am including the source of the program that created it.
But you can look at the source, and even try it yourself.

### Build it

    $ go build postorder.go
    $ ./comp 5 9 2 0 3 8 7 6 1
 
The `comp` script invokes [feh](https://feh.finalrewind.org/) to immediately display the
two binary search trees it creates.

## Verification - fuzzing

I build a [version of the program](randtree.org) that accepts a single number on its command line,
generates that many pseudo-random integers, and inserts them into a binary search tree.
A post-order traversal of that tree generates an array of integers, which gets used
in the array order to create a second tree.

More-or-less random trees help flush out any bugs in the algoirthm.

## Performance
