
+++
title = "#each in Elixir (pt 1)"
draft = true
date = "2016-09-04T22:03:00.000Z"

+++
###### Ruby
Rubyists use `#each` all over the place to iterate though collections.  Its super common to see code along these lines:
```
[1, 2, 3].each {|num| puts num}
```

Basically what `#each` does is execute the given block on every element of the given collection and then return the collection. Its usually used for side effects, rather than the return value.  `Array` implements this method, as do `Hash`, `Range`, and, even more interestingly, `CSV`.

>So how do I do the same thing in Elixir?
###### Elixir
This simple question gets at some of the difference between these two languages, and so I'll explore it in a little depth.

Here is some Elixir code roughly equivalent to the above:
```
Enum.each([1, 2, 3], fn num -> IO.puts num end)

# Or with function shorthand syntax

Enum.each([1, 2, 3], &(IO.puts &1))

# Or most compact

Enum.each([1, 2, 3], &IO.puts/1)

# Or perhaps most *Elixirishly*

[1, 2, 3] |> Enum.each(&IO.puts/1)
```

##### Look ma, no methods!

Notice that we're not invoking a method on the array itself (this is actually a list, rather than an array). Elixir has types (list is a type) and functions (`#each` is a function scoped to the `Enum` module).  It does __not__ have objects which implement methods. Basically we're feeding the array into `Enum.each` and also giving *that* function an anonymous function in the second parameter to invoke on every member of the array.

As for the versions above, that the second uses the shorthand inline function syntax (kind of similar Ruby's symbol to proc).  The `&1` is the first parameter passed to the function.

The third version is most compact, but the most foreign looking.  Basically this second parameter demonstrates another version of the shorthand syntax that wraps a function (for which you specify the arity with the `/`) and then passes the first argument implicitly to that function.

The fourth is the most Elixirish, as it uses the famous `|>` or pipe operator.  This operator basically functions like a pipeline, passing the data along and calling the functions with the data (or return value of a previous function) as the first parameter of the function call.  This allows you to build series of transformations such as:
```

iex> [1, 2, 3, 4] |> Enum.filter(&Integer.is_even/1) |> Enum.map(&(&1 * 3))
[6, 12]
```

So there you have it, `#each` in Elixir.  There a lot more to the story, actually, like how it's possible to enumerate through different kinds of collections that aren't lists (yes, you can pass in tuples, maps, etc into `Enum.each`).  I'll get into how this works in part two of this post, so stay tuned!


