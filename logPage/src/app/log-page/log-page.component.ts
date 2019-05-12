import { Component, OnInit } from '@angular/core';
import {ServerService } from '../server.service';
import {  loglist , codeformat, Orelog } from '../struct';
// import { ApiSerivice } from '../apiservice';
// import { ActivatedRoute, Router } from '@angular/router';
import * as $ from 'jquery'
// import { Orelog } from '../apiresponse';

@Component({
  selector: 'app-log-page',
  templateUrl: './log-page.component.html',
  styleUrls: ['./log-page.component.css']
})

export class LogPageComponent implements OnInit {
  constructor(private server: ServerService) { }
    ngOnInit() {
      this.Getorelog();
    this.Getloglist(0);
     this.Getlogtext("2019-05-12_1","info.log");

     for (let i = 0; i < 100; i++) {
      this.listOfAllData.push({
        id: i,
        name: `Edward King ${i}`,
        age: 32,
        address: `London, Park Lane no. ${i}`
      });
    }
  }
  filename = "";  //当前日志文本框显示的日志的文件名和路径
  page = 0; //当前浏览的是日志列表的第几页
  list = loglist[50]; //日志目录
  logtext = [];  //日志文本

  
  // tslint:disable-next-line:member-ordering
  public orelogs: Orelog[];

  public orelog: Orelog = {
    logid: 0,
    logtime: '',
    admin:'',
    module:'',
    logsql:''
  };

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

  updatelogtext(date:string,input:number, ty:string){
    var tag = date +"_"+input;
    this.Getlogtext(tag, ty);
  }

  // 请求操作日志内容
  public Getorelog() {           // 获取信息
    this.server.getorelog().subscribe(response => {
      this.orelogs = response;
    });
  }


}
