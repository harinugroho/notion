# Notion
Library for make notion database response more Simple

# Sample Usage
### Get Database Info And Properties
```
client, err := NewClient(
  "secret_xxxxxxxxxxxxxxxxx",
  "https://www.notion.so/your/page/url",
).GetDatabase()

if err != nil {
  panic(err)
}

fmt.Println("============    Info    ============")
info := client.GetInfo()
for key, value := range info {
  fmt.Printf("%s:%v\n", key, value)
}

fmt.Println("============ Properties ============")
listProperties := client.GetProperties()
for key, value := range listProperties {
  fmt.Printf("%s:%v\n", key, value)
}
```
### Get List Data
```
client, err := NewClient(
  "secret_xxxxxxxxxxxxxxxxx",
  "https://www.notion.so/your/page/url",
).Filters(properties.Filter{
  Logic: "and",
  Items: []properties.FilterItem{
    {
      Property: "Your Property",
      Type:     "checkbox",
      Logic:    "equals",
      Value:    true,
    },
  },
}).Sorts(map[string]string{
  "Your Property": "descending",
}).GetList()
if err != nil {
  panic(err)
}

fmt.Println("============    List    ============")
results := client.GetResults()
for _, result := range results {
  for key, value := range result {
    fmt.Printf("%s:%v\n", key, value)
  }
  fmt.Println("====================================")
}
```
