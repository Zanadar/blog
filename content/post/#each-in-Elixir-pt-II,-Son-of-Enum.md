
+++
title = "#each in Elixir pt II, Son of Enum"
draft = true
date = "2016-09-07T13:59:47.000Z"

+++

###### Pt. 2

I mentioned earlier that many Ruby classes implement `each`.  Ruby is an OO language, so that makes sense, but what happens when I want to call `#each` in Elixir on something other than a list.  I'll stress again, the list is not an object, its a type.  So we can't have different objects implementing the method in their particular way.

Lets delve a bit deeper by looking in the standard library for the [implementation](https://github.com/elixir-lang/elixir/blob/master/lib/elixir/lib/enum.ex#L632) of this function:
```
def each(enumerable, fun) when is_list(enumerable) and is_function(fun, 1) do
    :lists.foreach(fun, enumerable)
    :ok
  end

def each(enumerable, fun) when is_function(fun, 1) do
    reduce(enumerable, nil, fn(entry, _) ->
      fun.(entry)
      nil
    end)
    :ok
end
```

I've cleaned this up a bit, but the skeleton is there.  This code my look really weird if you're not used to reading Elixir, but I'll walk you through it slowly.  One of the cool things about this language is that much of Elixir is written in Elixir :fireworks: (as opposed to Ruby, which is mostly written in C). So if you lean the basic syntax, understanding the functioning of the rest of the language becomes straightforward.

At a high level, what we see are two definitions of `#each`.  Wait, what?  Two definitions?  Yup.  Pattern matching.  I won't go into that now, but this is where you'd usually see a case or switch statement in other languages.  Okay so reading the first `#each` definition we notice that clauses in the function head (those are called guard clauses).  This function will match on our earlier examples because the enumerable (the list in this case) is a list. The function then executes some other weird looking code. `:lists.forEach`.  This is actually Erlang code being called directly from the Elixir code, which .

https://imgflip.com/i/1a4h5w
https://imgflip.com/i/1a5c63

