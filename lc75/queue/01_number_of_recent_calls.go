package queue

type RecentCounter struct {
    requests []int
}

func Constructor() RecentCounter {
    return RecentCounter{
        requests: []int{},
    }
}

func (this *RecentCounter) Ping(t int) int {
    this.requests = append(this.requests, t)

    // Remove requests older than t-3000 from the start of slice
    for len(this.requests) > 0 && this.requests[0] < t-3000 {
        this.requests = this.requests[1:]
    }

    return len(this.requests)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */