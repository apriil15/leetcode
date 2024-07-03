package main

import (
	"cmp"

	"github.com/emirpasic/gods/queues/priorityqueue"
)

func main() {

}

type Twitter struct {
	timestamp           int                      // a global timestamp to mark every tweet's timestamp
	userIDToFolloweeIDs map[int]map[int]struct{} // use set, so that it would be better for [Twitter.Unfollow]
	userIDToTweets      map[int][]Tweet
}

type Tweet struct {
	timestamp int
	id        int
}

// Node for maxHeap usage in [Twitter.GetNewsFeed]
type Node struct {
	timestamp  int
	tweetID    int
	followeeID int
	nextIndex  int
}

func Constructor() Twitter {
	return Twitter{
		timestamp:           0,
		userIDToFolloweeIDs: make(map[int]map[int]struct{}),
		userIDToTweets:      make(map[int][]Tweet),
	}
}

func (this *Twitter) PostTweet(userID int, tweetID int) {
	this.userIDToTweets[userID] = append(this.userIDToTweets[userID], Tweet{
		timestamp: this.timestamp,
		id:        tweetID,
	})

	this.timestamp++
}

func (this *Twitter) GetNewsFeed(userID int) []int {
	this.Follow(userID, userID)

	// Enqueue each followee's last tweet
	maxHeap := priorityqueue.NewWith(func(a, b any) int {
		aa := a.(Node)
		bb := b.(Node)
		return -cmp.Compare(aa.timestamp, bb.timestamp)
	})
	for followeeID := range this.userIDToFolloweeIDs[userID] {
		tweets := this.userIDToTweets[followeeID]
		if len(tweets) == 0 {
			continue
		}

		lastIndex := len(tweets) - 1
		lastTweet := tweets[lastIndex]

		maxHeap.Enqueue(Node{
			timestamp:  lastTweet.timestamp,
			tweetID:    lastTweet.id,
			followeeID: followeeID,
			nextIndex:  lastIndex - 1,
		})
	}

	var res []int
	for maxHeap.Size() > 0 && len(res) < 10 {
		tmp, _ := maxHeap.Dequeue()
		node := tmp.(Node)

		res = append(res, node.tweetID)

		if node.nextIndex >= 0 {
			tweet := this.userIDToTweets[node.followeeID][node.nextIndex]

			maxHeap.Enqueue(Node{
				timestamp:  tweet.timestamp,
				tweetID:    tweet.id,
				followeeID: node.followeeID,
				nextIndex:  node.nextIndex - 1,
			})
		}
	}
	return res
}

func (this *Twitter) Follow(followerID int, followeeID int) {
	if this.userIDToFolloweeIDs[followerID] == nil {
		this.userIDToFolloweeIDs[followerID] = make(map[int]struct{})
	}
	this.userIDToFolloweeIDs[followerID][followeeID] = struct{}{}
}

func (this *Twitter) Unfollow(followerID int, followeeID int) {
	delete(this.userIDToFolloweeIDs[followerID], followeeID)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
