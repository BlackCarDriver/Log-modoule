//描述日志目录，date为日期,times为次数列表
export class loglist{
	name:string;
	list:string[];
}
export class codeformat{
	text : string[];
}
// 操作日志目录
export class Orelog {
    public logid:       number;
    public logtime:     string;
    public admin:       string;
    public module:      string;
    public logsql:      string;
}
