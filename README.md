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
    description,
    applications {
      id,
      name,  
    }, 
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

* small cells
```graphql
query {
    smallCells{
      id,
      name,
    }
}
```

* slices
```graphql
query {
    slices{
      id,
      name,
    }
}
```
