// Those configuration options are reserved for a future version
log_dir     = "log"
data_dir    = "drill-data"
listen      = "5177"
nomad_host  = "127.0.0.1"


drill {
  network {
    enabled   = true
    frequency = 60
  }
  storage {
    enabled   = true
    frequency = 60
  }
}
