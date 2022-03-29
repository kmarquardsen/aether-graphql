# aether graphql

## Queries

* simcards
```graphql
query {
    simCards{
    id,
    name,
  }
}
```

* sites
```graphql
query {
    sites{
    id,
    name,
    devices{
      id,
      name,
    }
  }
}
```

* devices
```graphql
query {
    devices{
      id,
      name,
    }
}
```

* enterprises
```graphql
query {
    enterprises{
    id,
    name,
    description
    sites{
      id,
      name,
      devices{
        id,
        name,
      }
    },
  }
}
```
