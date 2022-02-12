# firebase-claims-explorer

## Abstract

Will be a tui tool which can be used to explore production claims for a firebase authentication.
A common use case I've notices is that people have to use amazon provided cli to manually check what claims are applied to a specific user at a point of time.

Instead of a cli, this aims to be a easy to use TUI to explore a firebase auth database with the intent to view the custom claims for a user.

## Screens 

### Main Screen

I'm imagining this would look like a two panel layout, the left layout pre-filled with some user's name. User may navigate using <kbd>j/down</kbd> to scroll down and <kbd>k/up</kbd> to scroll up. Hit an enter on any user's name and the right side should get populated with the claims for that user.

There should be a search box where user may input part of a name or uid to quicky skim through the users.


## Features to implement

- Cache user database locally for quick searching 
