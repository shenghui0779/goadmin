package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "user",
			Collation: "utf8mb4_general_ci",
		},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Comment("自增ID"),
		field.String("name").MaxLen(32).Default("").Comment("用户昵称"),
		field.String("email").MaxLen(32).Default("").Comment("邮箱"),
		field.String("password").MaxLen(256).Default("").Comment("用户手机号"),
		field.String("salt").MaxLen(16).Default("").Comment("加密盐"),
		field.Uint8("role").Default(0).Comment("角色：1 - 超级管理员；2 - 高级管理员；3 - 普通管理员"),
		field.Int64("registed_at").Immutable().Default(0).Comment("注册时间"),
		field.Int64("last_login_at").Default(0).Comment("最近一次登录时间"),
		field.Int64("created_at").Immutable().Default(0).Comment("创建时间"),
		field.Int64("updated_at").Default(0).Comment("更新时间"),
		field.Int64("deleted_at").Default(0).Comment("删除时间"),
	}
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").StorageKey("uniq_name").Unique(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}
