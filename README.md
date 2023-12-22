# Mechanical Drill


## Description
Mechanical Drill is a tool to enable a hardware aware Nomad cluster by discovering hardware equipment connected to a machine. This equipment information is given to a nomad node leveraging its metadata which can be used for a more dynamic scheduling control using placement constraints.


## Installation
```bash
wget https://github.com/cloudputation/mechanical-drill/releases/download/${RELEASE}/mechanical-drill -O ./drill
chmod +x ./drill
mv ./drill /usr/bin/
```

## Configuration
By default, Mechanical Drill looks for `/etc/mechanical-drill/config.hcl`

Here is an example server config file
```hcl
// Those configuration options are reserved for a future version
log_dir     = "log"
data_dir    = "drill-data"
listen      = "5177"
nomad_host  = "127.0.0.1"
//

drill {
  system {
    enabled   = true
    frequency = 60
  }
  network {
    enabled   = true
    frequency = 60
  }
  cpu {
    enabled   = true
    frequency = 60
  }
  battery {
    enabled   = true
    frequency = 30
  }
}
```

### Complete hardware classes avalaible
`System`<br>
`Battery`<br>
`Bridge`<br>
`Memory`<br>
`Processor`<br>
`Address`<br>
`Storage`<br>
`Disk`<br>
`Tape`<br>
`Bus`<br>
`Network`<br>
`Display`<br>
`Input`<br>
`Printer`<br>
`Multimedia`<br>
`Communication`<br>
`Power`<br>
`Volume`<br>
`Generic`<br>


## Usage
Just run the executable as is:
```bash
drill
```

## Classification
Device are collected as a list. The first occurrence is 0 meaning it will be labeled as `device0`. For example: if you want to schedule a job based on the percentage of the first battery of your system then `md.battery.device0.percentage` will be the correct constraint to use in your Nomad job.

#### Schedule jobs to manage a fleet of drones based on battery level
Fly to recharge pad:
```hcl
constraint {
  attribute = "${meta.md.battery.device0.percentage}"
  operator  = "<"
  value     = "30"
}
```

Return to work:
```hcl
constraint {
  attribute = "${meta.md.battery.device0.state}"
  value     = "fully-charged"
}
```

## Compatibility
For now Mechanical Drill is only compatible with Linux systems

## Specifications
Although the classes `battery` and `power` are similar, they are different in their purpose:
`battery` returns the battery status and charge percentage. Possible statuses are "charging" and "fully-charged".

`power` returns hardware specific information such as the manufacturer

## Contributing
Feel free to create an issue or propose a pull request.
Follow the [Code of Conduct](CODE_OF_CONDUCT.md).
