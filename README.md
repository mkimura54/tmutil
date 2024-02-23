# tmutil

時間を扱うためのユーティリティ

## Examples

```
func main() {
	tt, _ := tmutil.NewTimeTermJust(
		time.Date(2024, 2, 22, 22, 22, 22, 0, time.Local),
		time.Date(2024, 2, 22, 22, 33, 33, 0, time.Local),
	)
	fmt.Println(tmutil.Contains(tt, time.Date(2024, 2, 22, 22, 33, 11, 0, time.Local))) // true
}
```

```
func main() {
	tt1, _ := tmutil.NewTimeTermJust(
		time.Date(2024, 2, 22, 10, 0, 0, 0, time.Local),
		time.Date(2024, 2, 22, 11, 0, 0, 0, time.Local),
	)
	tt2, _ := tmutil.NewTimeTermJust(
		time.Date(2024, 2, 22, 10, 30, 0, 0, time.Local),
		time.Date(2024, 2, 22, 13, 30, 0, 0, time.Local),
	)
	
	d, ok := tmutil.DuplicateTerm(tt1, tt2)
	if !ok {
		fmt.Println("no duplicate")
		return
	}
	fmt.Println(d.Start().Format(time.DateTime)) // 2024-02-22 10:30:00
	fmt.Println(d.End().Format(time.DateTime))   // 2024-02-22 11:00:00
}
```
