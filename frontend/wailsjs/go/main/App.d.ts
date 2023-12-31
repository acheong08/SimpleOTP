// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {database} from '../models';

export function AddEntry(arg1:database.Entry):Promise<void>;

export function DeleteEntry(arg1:string):Promise<void>;

export function GenerateCode(arg1:string):Promise<string>;

export function List():Promise<Array<database.Entry>>;

export function Login(arg1:string):Promise<string>;

export function SaveState():Promise<void>;

export function Search(arg1:string):Promise<Array<database.Entry>>;
