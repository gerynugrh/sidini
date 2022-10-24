package forum_membership

type ForumMembership struct {
	ForumID uint64
	UserID  uint64
	Roles   []string
}
