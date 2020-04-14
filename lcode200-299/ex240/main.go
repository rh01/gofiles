package main

import (
	"sort"
	"time"
)

func main() {

}

type tweet struct {
	id   int
	time int64
}

type tweets []tweet

func (t tweets) Len() int {
	return len(t)
}

func (t tweets) Less(i, j int) bool {
	return t[i].time > t[j].time
}

func (t tweets) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

// Twitter 是一个推特用户
type Twitter struct {
	userTweets map[int]tweets
	follow     map[int][]int
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	t := make(map[int]tweets)
	f := make(map[int][]int)
	return Twitter{userTweets: t, follow: f}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.userTweets[userId] = append(this.userTweets[userId], tweet{
		id:   tweetId,
		time: time.Now().UnixNano(),
	})
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	// 获取本人的tweet
	temp := make(tweets, len(this.userTweets[userId]))
	copy(temp, this.userTweets[userId])

	// 获取follow的tweet
	for _, id := range this.follow[userId] {
		temp = append(temp, this.userTweets[id]...)
	}

	// 按照时间排序
	sort.Sort(temp)

	res := make([]int, 0, 10)
	// 获取最近10条或者更少的内容
	for i := 0; i < len(temp) && i < 10; i++ {
		res = append(res, temp[i].id)
	}
	return res
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	// 不能follow自己
	if followerId == followeeId {
		return
	}

	// 不能重复follow
	for _, id := range this.follow[followerId] {
		if id == followeeId {
			return
		}
	}

	this.follow[followerId] = append(this.follow[followerId], followeeId)
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	for i, id := range this.follow[followerId] {
		if id == followeeId {
			// 删除followee的记录
			this.follow[followerId] = append(this.follow[followerId][:i], this.follow[followerId][i+1:]...)
		}
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
