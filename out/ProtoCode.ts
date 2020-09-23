export module pb{
	export const enum Cmds {
		C2S_EnterGame = 10001,
		S2C_EnterGame = 20001,
	}
	export var cmds:{ [key: number]: string }={
		10001:"C2S_EnterGame",
		20001:"S2C_EnterGame"
	}
	export var cfgs:{ [key: string]: string[][] }={
		"Bcst_EntityEnter":[["1","spaceid","1"],["2","entitys","Entity","1"]],
		"Bcst_EntityLeave":[["1","spaceid","1"],["2","eids","1","1"]],
		"Bcst_EntityMove":[["1","spaceid","1"],["2","eid","1"],["3","position","Vector3"]],
		"Entity":[["1","id","1"],["2","owner","1"],["3","position","Vector3"],["4","spaceid","1"],["5","name","1"],["6","etype","6"],["7","t","TT"]],
		"Vector3":[["1","x","3"],["2","y","3"],["3","z","3"]],
		"C2S_EnterGame":[["1","roleid","1"]],
		"S2C_EnterGame":[["1","error","8"],["2","self","Entity"],["3","entitys","Entity","1"]],
		"C2S_RegisterAccount":[["1","account","1"],["2","password","1"]],
		"S2C_RegisterAccount":[["1","error","8"]],
		"C2S_LoginGame":[["1","account","1"],["2","password","1"]],
		"S2C_LoginGame":[["1","error","8"]],
		"S2C_RoleList":[["1","error","8"],["2","role_list","Entity","1"]],
		"C2S_CreateRole":[["1","name","1"]],
		"S2C_CreateRole":[["1","error","8"],["2","role","Entity"]],
		"G2S_CreateSpace":[["1","spaceid","1"]],
		"S2G_CreateSpace":[["1","error","8"]],
		"G2L_RoleList":[["1","account","1"]],
		"G2L_CreateRole":[["1","name","1"],["2","account","1"]]
	}

	export const enum TT {
		A=0,
		B=1
	}
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
		t:TT;
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