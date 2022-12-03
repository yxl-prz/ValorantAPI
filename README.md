# Valorant API
This is a wrapper for "api.henrikdev.xyz" in Go.

## Get Started
```
go get github.com/yxl-prz/ValorantAPI
```
### Print all Players
```go
res, err := ValorantAPI.GetUserMatches(ValorantAPI.SERVER_TYPE_NA, "user", "tag")
if err != nil {
	panic(err)
}

if res.Status == 400 {
	return
}

fmt.Println("Players in Match:")
for _, player := range res.Matches[0].Players.AllPlayers {
	fmt.Printf("%v#%v | %v | Played for %v min.\n", player.Name, player.Tag, player.CurrentTierPatched, player.SessionPlaytime.Minutes)
}
```
### Saving Match Data
```go
res, err := ValorantAPI.GetUserMatches(ValorantAPI.SERVER_TYPE_NA, "user", "tag")
if err != nil {
	panic(err)
}

if res.Status == 400 {
	return
}

err = res.Save("match.json")
if err != nil {
	panic(err)
}
```

### Loading Match Data
```go
res, err := ValorantAPI.ReadMatchesFromFile("match.json")
if err != nil {
	panic(err)
}

if res.Status == 400 {
	return
}

// Here you can use res
```

## Legal
This project is not affiliated with Riot Games or any of its employees and therefore does not reflect the views of said parties. This is purely a fan-made API wraper made for public use.

Riot Games does not endorse or sponsor this project. Riot Games, and all associated properties are trademarks or registered trademarks of Riot Games, Inc.

# LICENSE
![gnu-logo](./media/gplv3-88x31.png)

This program is free software: you can redistribute it and/or modify
it under the terms of the [GNU General Public License](https://github.com/NeutronX-dev/ws.js/blob/main/LICENSE) as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.
