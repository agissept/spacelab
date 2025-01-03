# Spacelab

Personal lab to make [starhship.rs](https://starship.rs) more fun for my self.


## Preview

## Feature
### Play poker hand
### Now playing on Spotify

## Example usage
```
starship.toml

format = "${custom.poker}${custom.spotify}"

[custom.poker]
description = "Poker playing hands"
command = "spacelab pp"
when = "true"
format = "$output"

[custom.spotify]
command = 'spacelab np'
when = '''test -n "$(spacelab np)"'''
symbol = ""
style = "bg:#1DB954 fg:black"
format = ' - [](#1DB954)[$symbol ($output )]($style)[](#1DB954)'
```