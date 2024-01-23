# Bottom up declaration
We don't want big files because they will be a pain to work with:

- It's harder to find where things are and harder to nagivate within the file.
- There are increased chance of multiple working versions modifying the same file and causing conflicts.

To do that, we should split big package into smaller packages, and then functions can be in its own file. This is not necessary, but it helps and there is pretty much no cost in doing so, as we can still user stuff from the same package.

However, multiple smaller packages can again lead to giant import list. Particularly, the router package might have to import one package for each endpoint if each end point are in their own package. To solve this, we use a bottom up method of declaring stuff:

- We implement the handler and declare its endpoint in a single file (using `func init()`).
- The files are imported into a single master file that will be imported into main. We can generate this import list using the file system, so there's no chance of missing things. We can also temporarily disable a feature if we want.

This type of bottom up declaration should also be used for other types of handler.

## Naming
File dedicated to a type / function should be named after that type / function, except we need to use snake_case for the file name and camelCase for the type/function. If we want the type to be exported outside the package, we also need to use CamelCase or PascalCase.