import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, of } from 'rxjs';
import {  loglist, Orelog } from './struct';
@Injectable({
  providedIn: 'root'
})
export class ServerService {
  private addr: string  = "http://localhost:8090";
  constructor(private http: HttpClient) { }

  //获取系统日志列表，参数page指定第几页
  GetLogCatalog (page:number){
    var url = this.addr + "/log/getlogpage?page="+page;
    return this.http.get<loglist[]>(url);
  }
  //获取日志文本,参数为日期,第几次和文件名
  GetLogText(floder:string, name:string){
    var url = this.addr + "/log/getlogtext?floder="+floder+"&&name="+name;
    return this.http.get<any>(url);
  }
  //获取操作日志
  public getorelog(): Observable<Orelog[]> {
    return this.http.get<Orelog[]>('http://localhost:8090/log/logdisplay');
}
}
