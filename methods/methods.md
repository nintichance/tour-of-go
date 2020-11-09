**Go** does not have classes; however, you can define methods on types.

A **method** is a function with a special *receiver* argument.

The receiver appears in its own argument list  between the ```func``` keyword and the method name.

"With receiver functions, you don't have to mess aroud with classes or deal with ingeritance. The Vertex type has no  knowledge of the receiver function (method). One advantage of using receiver function is when we couple it with interfaces."

### Remember, methods are just functions with receivers

A **method** can be declared on non-struct types as well
    * You can only declare a method with a receiver whose type is defined in the same package of the receiver.
    * even built-in packages such as int must be defined within the same package to be used with the receiver
