
+++
title = "Bootstrap + React = React-Bootstrap"
draft = false
date = "2016-06-01T18:51:49.000Z"

+++
Say you're building an app, and you want to put it in front of people. And you don't want it to look shitty. Bootstrap has long been a go to library for this. It does what it says on the tin: it lets you bootstrap your UI and move forward with building the core functionality of your app.

As an example, say you have a dropdown menu that calls some JavaScript code on a model object. With regular bootstrap, you have the functionality in two places: a set of HTML/CSS that implements the markup for the components and a set of javscript that implements the functionality. This creates a confusing situation (although that's kind of where the web is at these days).

```
<div class="dropdown">
  <button class="btn btn-default dropdown-toggle" type="button" id="dropdownMenu1" data-toggle="dropdown" aria-haspopup="true" aria-expanded="true">
    Dropdown
    <span class="caret"></span>
  </button>
  <ul class="dropdown-menu" aria-labelledby="dropdownMenu1">
    <li><a href="#" id="action1">One</a></li>
    <li><a href="#" id='action2'>Two</a></li>
    <li><a href="#" id='action3'>Three</a></li>
    <li role="separator" class="divider"></li>
    <li><a href="#">Separated link</a></li>
  </ul>
</div>

```
```
var model = {
  'one': function(){ alert('one')},
  'two': function(){ alert('two')},
  'three': function(){ alert('three')}
}

$('.dropdown-toggle').dropdown()
$('#action1').on('click', alerts.one)
$('#action2').on('click', alerts.two)
$('#action3').on('click', alerts.three)
```

Enter React. React is Facebook's library for building UI. It's designed around a component based architecture, where functionality and structure are implemented in tandem, in one place. And with this philosophy applied to bootstrap, you get nice clean code that truly gets at the idea of components. The above tab example becomes pretty clean:
```
function alertSelect(key) {
  alerts[key]()
}

var model = {
  'one': function(){ alert('one')},
  'two': function(){ alert('two')},
  'three': function(){ alert('three')}
}

const ToolBar = (
    <Dropdown id="dropdown-custom-1" open={true}>
      <Dropdown.Menu>
        <MenuItem eventKey="one" onSelect={alertSelect}>One</MenuItem>
        <MenuItem eventKey="two" onSelect={alertSelect}>Two action</MenuItem>
        <MenuItem eventKey="three" onSelect={alertSelect}>Three</MenuItem>
        <MenuItem divider />
        <MenuItem eventKey="4">Separated link</MenuItem>
      </Dropdown.Menu>
    </Dropdown>
)

ReactDOM.render(Toolbar, mountNode);
```
As you can see, we're using JSX syntax to simply instantiate these components (provided by the `React-Bootstrap` library and render them into our page. And things are much cleaner and self-contained.

As I see it, there are two wins you get here.  You've started down the path to using `React` in your app, but you've been able to punt on implementing everything by hand. In the future, if and when you start to build custom components, so you have everything in place. You've also got easy integration for custom Javascript.  The other win is that you've got Bootstrap in there, but it doesn't look like a cat barfed `divs` all over your page. 

Be warned: The react-bootstrap library is very much in flux.  Its under pretty active development, but hasn't hit `1.0.0` yet.  That said, its a pretty exciting project and maybe you could contribute?!



