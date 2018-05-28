## packet create volume

A brief description of your command

### Synopsis

A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

```
packet create volume [flags]
```

### Options

```
  -b, --billing-cycle string   --billing-cycle or -b [billing_cycle] (default "hourly")
  -d, --description string     --description or -d [description]
  -f, --facility string        --facility or -f [facility_code]
  -h, --help                   help for volume
  -l, --locked                 --locked or -l
  -P, --plan string            --plan or -P [plan_name]
  -p, --project-id string      --project-id or -p [UUID]
  -s, --size int               --size or -s [size_in_GB]
```

### Options inherited from parent commands

```
  -j, --json   -j or --json JSON output
  -y, --yaml   -y or --yaml YAML output
```

### SEE ALSO

* [packet create](packet_create.md)	 - Create command

###### Auto generated by spf13/cobra on 28-May-2018