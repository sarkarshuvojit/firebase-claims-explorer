# firebase-claims-explorer

Is a tui tool which can be used to explore production claims for a firebase authentication. A common use case I've notices is that people have to use firebase provided cli to manually check what claims are applied to a specific user at a point of time.

Instead of a cli, this aims to be a easy to use TUI to explore a firebase auth database with the intent to view the custom claims for a user.

## firebase-claims-exporer

A tui application to manage claims in your firebase app

### Synopsis

A tui application to manage claims in your firebase app

### Options

```
  -c, --config string   Pass config file from firebase admin
  -h, --help            help for firebase-claims-exporer
```


## firebase-claims-exporer explore

Runs tui app to list users and view claims

### Synopsis

Runs tui app to list users and view claims

```
firebase-claims-exporer explore [flags]
```

### Options

```
  -h, --help   help for explore
```

### Options inherited from parent commands

```
  -c, --config string   Pass config file from firebase admin
```

## firebase-claims-exporer seed

Used to seed firebase with random users.

### Synopsis

Used to seed firebase with random users.

```
firebase-claims-exporer seed [flags]
```

### Options

```
  -h, --help       help for seed
  -s, --size int   Number of users to seed. Default is 10. (default 10)
```

### Options inherited from parent commands

```
  -c, --config string   Pass config file from firebase admin
```

###### Readme Auto generated by spf13/cobra on 23-Feb-2022

## Features to implement

- Cache user database locally for quick searching 
