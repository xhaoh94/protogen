export module pb{
	export interface Bcst_EntityEnter {
		spaceid:string;
		entitys:Entity[];
	}
	export interface Bcst_EntityLeave {
		spaceid:string;
		eids:string[];
	}
	export interface Bcst_EntityMove {
		spaceid:string;
		eid:string;
		position:Vector3;
	}
	export interface Entity {
		id:string;
		owner:string;
		position:Vector3;
		spaceid:string;
		name:string;
		etype:number;
	}
	export interface Vector3 {
		x:number;
		y:number;
		z:number;
	}
	export interface C2S_EnterGame {
		roleid:string;
	}
	export interface S2C_EnterGame {
		error:number;
		self:Entity;
		entitys:Entity[];
	}
	export interface C2S_RegisterAccount {
		account:string;
		password:string;
	}
	export interface S2C_RegisterAccount {
		error:number;
	}
	export interface C2S_LoginGame {
		account:string;
		password:string;
	}
	export interface S2C_LoginGame {
		error:number;
	}
	export interface S2C_RoleList {
		error:number;
		role_list:Entity[];
	}
	export interface C2S_CreateRole {
		name:string;
	}
	export interface S2C_CreateRole {
		error:number;
		role:Entity;
	}
	export interface G2S_CreateSpace {
		spaceid:string;
	}
	export interface S2G_CreateSpace {
		error:number;
	}
	export interface G2L_RoleList {
		account:string;
	}
	export interface G2L_CreateRole {
		name:string;
		account:string;
	}
}