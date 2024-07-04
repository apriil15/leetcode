package main

import (
	"cmp"

	"github.com/emirpasic/gods/queues/priorityqueue"
)

func main() {

}

type Twitter struct {
	timestamp           int                      // a global timestamp to mark every tweet's timestamp
	userIDToFollowIDSet map[int]map[int]struct{} // use set, so that it would be better for [Twitter.Unfollow]
	userIDToTweets      map[int][]Tweet
}

type Tweet struct {
	timestamp int
	id        int
}

func Constructor() Twitter {
	return Twitter{
		timestamp:           0,
		userIDToFollowIDSet: make(map[int]map[int]struct{}),
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
	// follow myself, so that I can get my feed
	this.Follow(userID, userID)

	type Node struct {
		timestamp int
		tweetID   int
		followID  int
		nextIndex int
	}
	maxHeap := priorityqueue.NewWith(func(a, b any) int {
		return -cmp.Compare(a.(Node).timestamp, b.(Node).timestamp)
	})

	for followID := range this.userIDToFollowIDSet[userID] {
		tweets := this.userIDToTweets[followID]
		if len(tweets) == 0 {
			continue
		}

		// Enqueue each follow's last tweet
		lastIndex := len(tweets) - 1
		lastTweet := tweets[lastIndex]

		maxHeap.Enqueue(Node{
			timestamp: lastTweet.timestamp,
			tweetID:   lastTweet.id,
			followID:  followID,
			nextIndex: lastIndex - 1,
		})
	}

	// merge k sorted array
	var res []int
	for maxHeap.Size() > 0 && len(res) < 10 {
		tmp, _ := maxHeap.Dequeue()
		node := tmp.(Node)

		res = append(res, node.tweetID)

		if node.nextIndex >= 0 {
			tweet := this.userIDToTweets[node.followID][node.nextIndex]

			maxHeap.Enqueue(Node{
				timestamp: tweet.timestamp,
				tweetID:   tweet.id,
				followID:  node.followID,
				nextIndex: node.nextIndex - 1,
			})
		}
	}
	return res
}

func (this *Twitter) Follow(followerID int, followeeID int) {
	if this.userIDToFollowIDSet[followerID] == nil {
		this.userIDToFollowIDSet[followerID] = make(map[int]struct{})
	}
	this.userIDToFollowIDSet[followerID][followeeID] = struct{}{}
}

func (this *Twitter) Unfollow(followerID int, followeeID int) {
	delete(this.userIDToFollowIDSet[followerID], followeeID)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
