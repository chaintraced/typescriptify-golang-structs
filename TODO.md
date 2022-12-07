In order to properly add support for Generics the way we need them I will need to create:
1. config file where we can replace types, similar to how `ts_type:"string"` works
  ```
    replace ctjson.Field[string] string
  ```
2. config file where we can choose to skip and use the child as the primary type
  ```
    usechild ctjson.Field
  ```

If I enable the second point it should actually cover all of our use-cases, could be done with settings instead of config file.