package leetcode

/*
* 355. 设计推特
* https://leetcode-cn.com/problems/design-twitter/
 */

type Twitter struct {
	Users           map[int]*User
	curInnerTweetId int
}

type Tweet struct {
	Id      int
	innerId int //内部id用来排序
}

type User struct {
	Id int

	Follows   map[int]struct{} //关注了
	Followers map[int]struct{} //粉丝
	Tweets    []Tweet          //推文
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		Users: map[int]*User{},
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	user := this.getOrCreateUser(userId)

	tweet := Tweet{
		Id:      tweetId,
		innerId: this.getInnerTweetId(),
	}

	user.Tweets = append(user.Tweets, tweet)
}

func (this *Twitter) getInnerTweetId() int {
	this.curInnerTweetId++
	return this.curInnerTweetId
}

func (this *Twitter) getOrCreateUser(userId int) *User {
	if _, ok := this.Users[userId]; !ok {
		this.Users[userId] = &User{
			Id:        userId,
			Follows:   map[int]struct{}{},
			Followers: map[int]struct{}{},
		}
	}

	return this.Users[userId]
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	user := this.getOrCreateUser(userId)

	tweets := [][]Tweet{user.Tweets}

	for uid, _ := range user.Follows {
		tweets = append(tweets, this.getOrCreateUser(uid).Tweets)
	}

	var res []int
	pos := make([]int, len(tweets))
	num := 0
	allDone := true
	for num < 10 {
		minPos := -1
		max := Tweet{innerId: -1}
		for i := 0; i < len(tweets); i++ {
			idx := pos[i]
			tweet := tweets[i]
			l := len(tweet)
			if pos[i] < l {
				if tweet[l-idx-1].innerId > max.innerId {
					max = tweet[l-idx-1]
					minPos = i
				}
				allDone = false
			}
		}

		if minPos != -1 {
			pos[minPos]++
		}

		if allDone == true {
			break
		}

		res = append(res, max.Id)
		num++

		allDone = true
	}

	return res
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if followeeId == followerId {
		return //不能关注自己
	}
	follower := this.getOrCreateUser(followerId)
	followee := this.getOrCreateUser(followeeId)

	followee.Followers[followerId] = struct{}{}
	follower.Follows[followeeId] = struct{}{}
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	follower := this.getOrCreateUser(followerId)
	followee := this.getOrCreateUser(followeeId)

	delete(followee.Followers, followerId)
	delete(follower.Follows, followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
