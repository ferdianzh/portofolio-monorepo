package permission

type Permissions string

const (
	RoleCreate Permissions = "role.create"
	RoleRead   Permissions = "role.read"
	RoleUpdate Permissions = "role.update"
	RoleDelete Permissions = "role.delete"

	UserCreate Permissions = "user.create"
	UserRead   Permissions = "user.read"
	UserUpdate Permissions = "user.update"
	UserDelete Permissions = "user.delete"

	ProjectCreate Permissions = "project.create"
	ProjectRead   Permissions = "project.read"
	ProjectUpdate Permissions = "project.update"
	ProjectDelete Permissions = "project.delete"
)

// list for seeding
var All = []Permissions{
	RoleCreate, RoleRead, RoleUpdate, RoleDelete,
	UserCreate, UserRead, UserUpdate, UserDelete,
	ProjectCreate, ProjectRead, ProjectUpdate, ProjectDelete,
}
