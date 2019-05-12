import { Component, OnInit } from '@angular/core';
import {ServerService } from '../server.service';
import {  loglist ,codeformat } from '../struct'
import * as $ from 'jquery'

@Component({
  selector: 'app-log-page',
  templateUrl: './log-page.component.html',
  styleUrls: ['./log-page.component.css']
})

export class LogPageComponent implements OnInit {
  constructor(private server :ServerService ) { }
    ngOnInit() {
    this.Getloglist(0);
     this.Getlogtext("2019-05-12_1","error.log");
  }
  filename = "";  //当前日志文本框显示的日志的文件名和路径
  page = 0; //当前浏览的是日志列表的第几页
  list = loglist[50]; //日志目录
  logtext = [];  //日志文本

  //请求获取系统日志的目录列表
  Getloglist(page :number){
    this.server.GetLogCatalog(page).subscribe(result=>{
        this.list = result;
        console.log(this.list);
    });
  }

  //请求获取特定日志的文本内容
  Getlogtext(tag:string, type:string){
    this.server.GetLogText(tag,type).subscribe(result=>{
      this.logtext = result;
    });
  }

}
