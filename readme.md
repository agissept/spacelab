# Spacelab

Personal lab to make [starship.rs](https://starship.rs) more fun for my self.


## Preview
![Peek 2025-01-03 10-25](https://github.com/user-attachments/assets/fa8332be-505a-4a0c-b587-1224694c442c)

## Feature
### Play poker hand
![Peek 2025-01-03 10-22](https://github.com/user-attachments/assets/bdf1647d-8c13-416a-b0b5-179c2d6b5dec)

### Now playing on Spotify
![2025-01-03_10-24](https://github.com/user-attachments/assets/b20f89b7-3685-4e15-be6f-a2700a4cd97d)


## Example usage
```toml
# starship.toml

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
