package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	UserName string    `gorm:"column:username;unique" json:"username"`
	Picture  string    `gorm:"picture_url" json:"pictureUrl"`
}

func (User) TableName() string {
	return "users"
}

type Group struct {
	ID      uuid.UUID `gorm:"primaryKey"`
	Name    string    `gorm:"column:name" json:"name"`
	Desc    string    `gorm:"column:desc" json:"desc"`
	Picture string    `gorm:"column:picture_url" json:"pictureUrl"`
	Created time.Time `gorm:"column:created" json:"created"`
	Members []Member  `gorm:"foreignKey:GroupID"`
}

func (Group) TableName() string {
	return "groups"
}

type Member struct {
	ID               uuid.UUID `gorm:"primaryKey"`
	GroupID          uuid.UUID `gorm:"column:group_id;uniqueIndex:idx_first;size:191" json:"group_id"`
	UserID           uuid.UUID `gorm:"column:user_id;uniqueIndex:idx_first;size:191" json:"user_id"`
	User             User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Group            Group     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Nick             string    `gorm:"column:nick" json:"nick"`
	Adding           bool      `gorm:"column:adding" json:"adding"`
	DeletingMembers  bool      `gorm:"column:deleting_members" json:"deletingMembers"`
	DeletingMessages bool      `gorm:"column:deleting_messages" json:"deletingMessages"`
	Setting          bool      `gorm:"column:setting" json:"setting"`
	Creator          bool      `gorm:"column:creator" json:"creator"`
	Deleted          bool      `gorm:"column:deleted" json:"deleted"`
}

func (Member) TableName() string {
	return "members"
}

type Invite struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	IssId    uuid.UUID `gorm:"column:iss_id;size:191" json:"issID"`
	Iss      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"issuer"`
	TargetID uuid.UUID `gorm:"column:target_id;size:191" json:"targetID"`
	Target   User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	GroupID  uuid.UUID `gorm:"column:group_id;size:191" json:"groupID"`
	Group    Group     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"group"`
	Status   int       `gorm:"column:status" json:"status"`
	Created  time.Time `gorm:"column:created" json:"created"`
	Modified time.Time `gorm:"column:modified" json:"modified"`
}

func (Invite) TableName() string {
	return "invites"
}
