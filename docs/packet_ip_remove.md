## packet ip remove

Command to remove IP reservation.

### Synopsis

Example:	

packet ip remove --id [reservation-UUID]



```
packet ip remove [flags]
```

### Options

```
  -h, --help        help for remove
  -i, --id string   UUID of the reservation
```

### Options inherited from parent commands

```
      --config string     Path to JSON or YAML configuration file
      --exclude strings   Comma seperated Href references to collapse in results, may be dotted three levels deep
      --include strings   Comma seperated Href references to expand in results, may be dotted three levels deep
  -j, --json              JSON output
      --search string     Search keyword for use in 'get' actions. Search is not supported by all resources.
  -y, --yaml              YAML output
```

### SEE ALSO

* [packet ip](packet_ip.md)	 - IP operations

