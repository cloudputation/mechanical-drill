// Those configuration options are reserved for a future version
log_dir     = "log"
data_dir    = "drill-data"
listen      = "5177"
nomad_host  = "127.0.0.1"
//


drill {
  network {
    enabled   = true
    frequency = 60
  }
  storage {
    enabled   = true
    frequency = 60
  }
  bus {
    enabled   = true
    frequency = 60
  }
  memory {
    enabled   = true
    frequency = 60
  }
  cpu {
    enabled   = true
    frequency = 60
  }
  bridge {
    enabled   = true
    frequency = 60
  }
  hub {
    enabled   = true
    frequency = 60
  }
  display {
    enabled   = true
    frequency = 10
  }
  multimedia {
    enabled   = true
    frequency = 60
  }
  communication {
    enabled   = true
    frequency = 60
  }
  volume {
    enabled   = true
    frequency = 60
  }
  disk {
    enabled   = true
    frequency = 60
  }
  generic {
    enabled   = true
    frequency = 60
  }
}
