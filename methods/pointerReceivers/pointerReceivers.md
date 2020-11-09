You can declare methods with **pointer receivers**

The recever type has the literal syntax `*T` for some type `T`.

With a value receiver, the method will operate on a copy of the original type; however, with a pointer receiver, it operates on the original (mutates).

Notice that function with a pointer argument, must be called with a pointer; however, methods with a pointer argument can be received by a type either as a value or pointer.

This is due to Go interpreting the statements `v.Scale(5)`  and `(&v).Scale(5)` the same.

The difference is one is a copy of the original reference; while the other is the original.