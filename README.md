# tmutil

時間を扱うためのユーティリティ

## Examples

```
func main() {
	tt, _ := tmutil.NewTimeTermJust(
		time.Date(2024, 2, 22, 22, 22, 22, 0, time.Local),
		time.Date(2024, 2, 22, 22, 33, 33, 0, time.Local),
	)
	fmt.Println(tmutil.Contains(tt, time.Date(2024, 2, 22, 22, 33, 44, 0, time.Local)))
}
```
