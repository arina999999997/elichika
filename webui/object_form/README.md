# "Pretty" object form for http
Generate a form from a type.

Each fields must be:

- bool or *bool
- string or *string
- integer types or pointer to integer types

## Generated page format
Fields are generated from top to bottom. There will be a label on the left with the field name, and an `<input>` or `<select>` on the right.

Label default to the field name, but can be set using `of_label:""`

Each `<input>` will have id and name both set to the field name.

After that, the default value is set based on an object (if one was provided).

Finally, user can tag `of_attrs:""` for the rest of the attributes.

The type of the input is based on the type of the field, but it can be overridden by `of_type:""`. `<select>` will also be accessed this way.

## Default input types

### bool or *bool
Boolean always use a checkbox input, checked means true. There is no tag modifier for this.

```html
<input type="checkbox">
```

### string or *string
Strings by default use a text input with no specified restriction.

### Integers
Integer types use a number input with no limit.

## Custom input types
The following input types can be set using `of_type:""`:

### `select`
Use a `<select>` element for the input, user will need to supply the options using tag `of_options:""`:

- The options are multiple "lines", the line delimeter is `\n`
- Each options will be on 2 "lines".
- The first line is the key, the value that will be shown in the drop down menu.
- The second line is the value, the values that will be sent. 

### `time`
Use a `time` element for a time of day input.

- If the type is string, the time is passed in as is.
- If the type is integer, the time is converted to # of second since midnight.

<!-- ### `email`, `url`
These are alternative text input, treated like normal strings -->

## TODO
These would be nice feature to have, but they are currently not supported due to one problem or another:

- Generate `of_options` using an enum with enum tag.
- Support more input types
