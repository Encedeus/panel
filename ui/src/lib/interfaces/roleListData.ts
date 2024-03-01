import type {Role, User} from "@encedeus/js-api";

export interface RoleListData {
    rolesData: {
        role: Role
        userList: string[]
    }[]
}