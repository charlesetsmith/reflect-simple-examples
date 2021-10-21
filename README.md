# testreflect
Simple examples on using go Reflect for interface decoding of structs, slices and slices of structs

So this is very simple but it gets the gist of reflect and how to create an interface{} and how they
are handled by reflect.

Examples include
Simple Struct of uint32's and uint64's.

Simple Struct with members now including slices of uint's.

Getting more complex demonstrating a struct embedded within another struct of slices.

Yes there are much better ways of handling these using recursion but you need to crawl before you walk.
I will keep on adding more as needed.
I am using these to implement a "frame" interface to handle the differing frame types in the Saratoga Protocol.
There is one interface which can now handle all of the operations on "Beacon, Request, Metadata, Data, Status"
frame types. have have a look at my Saratoga repository and the "frames" package.

Cheers

