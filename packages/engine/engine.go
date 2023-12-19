package engine


func StartEngine() {
    scheduler := NewDrillScheduler()
    for {
        scheduler.RunController()
    }
}
